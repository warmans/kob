package bot

import (
	"fmt"
	"strings"
	"github.com/nlopes/slack"
	"go.uber.org/zap"
	"github.com/warmans/kob/server/pkg/db"
	"github.com/warmans/kob/server/pkg/search"
	"github.com/blevesearch/bleve"
)

func NewSlackBot(token string, logger *zap.Logger, store *db.Store, index *search.Index) *SlackBot {
	api := slack.New(token)
	return &SlackBot{logger: logger, rtm: api.NewRTM(), store: store, index: index}
}

type SlackBot struct {
	logger *zap.Logger
	rtm    *slack.RTM
	store  *db.Store
	index  *search.Index
}

func (b *SlackBot) Run() error {

	go b.rtm.ManageConnection()
	defer b.rtm.Disconnect()

	for {
		select {
		case msg := <-b.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				b.handleMessage(ev)

			case *slack.RTMError:
				b.logger.Error("slackbot error", zap.String("error", ev.Error()))

			case *slack.InvalidAuthEvent:
				return fmt.Errorf("invalid slack credentials")
			}
		}
	}
}

func (b *SlackBot) handleMessage(ev *slack.MessageEvent) {
	fmt.Printf("Message: %+v\n", ev)
	info := b.rtm.GetInfo()
	prefix := fmt.Sprintf("<@%s> ", info.User.ID)

	if ev.User == info.User.ID || ev.BotID != "" {
		return //don't respond to myself or bots (todo: the myself part possibly doesn't work?!)
	}

	if strings.HasPrefix(ev.Text, prefix) {
		b.handleDM(ev)
		return
	}

	searchRequest := bleve.NewSearchRequest(bleve.NewMatchQuery(ev.Msg.Text))
	searchRequest.Highlight = bleve.NewHighlight()

	result, err := b.index.Search(searchRequest)
	if err != nil {
		b.logger.Error("failed search for bot message", zap.Error(err))
		return
	}

	if result.Total > 0 {
		topResult := result.Hits[0]
		URL := "http://localhost:4200/entry/" + topResult.ID
		preview := search.MakePreview(topResult)
		params := slack.PostMessageParameters{
			Attachments: []slack.Attachment{{
				Fallback: "View Entry at " + URL,
				Text:     preview,
				Actions: []slack.AttachmentAction{{
					Type: "button",
					Name: "view_entry_"+topResult.ID,
					Text: "View Entry",
					Value: URL, Style: "primary"},
				},
			}},
		}
		b.rtm.PostMessage(ev.Channel, "I found some stuff!", params)
	}
}

func (b *SlackBot) handleDM(ev *slack.MessageEvent) {
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage("sorry, I am not implemented", ev.Channel))
}
