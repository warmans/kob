package oauth

import (
	"go.uber.org/zap"
	"github.com/warmans/kob/server/pkg/db"
	"net/url"
	"time"
	"net/http"
	"github.com/patrickmn/go-cache"
)

func NewOauthHandler(client *Client, stateCache *cache.Cache, store *db.Store, logger *zap.Logger) *Handler {
	return &Handler{client: client, stateCache: stateCache, store: store, logger: logger}
}

type Handler struct {
	client *Client
	stateCache *cache.Cache
	store *db.Store
	logger *zap.Logger
}

// HandleLoginRequest performs redirect to google login page.
func (h *Handler) HandleLoginRequest(rw http.ResponseWriter, r *http.Request) {
	redirectURL, state := h.client.CreateLoginURL()
	h.stateCache.Set(state, state, time.Minute * 10)
	http.Redirect(rw, r, redirectURL, http.StatusFound)
}

func (h *Handler) HandleLoginResponse(rw http.ResponseWriter, r *http.Request) {
	
	returnURL := "/login/complete"
	returnErr := ""
	returnToken := ""

	defer func() {
		qs := url.Values{}
		qs.Set("err", returnErr)
		qs.Set("token", returnToken)
		http.Redirect(rw, r, returnURL+"?"+qs.Encode(), http.StatusFound)
	}()

	if err := r.ParseForm(); err != nil {
		returnErr = "Invalid response from Google. Missing verification code"
		h.logger.Error(returnErr, zap.Error(err))
		return
	}
	if _, found := h.stateCache.Get(r.Form.Get("state")); !found {
		returnErr = "Auth code was invalid. Codes are only valid for up to 10 minutes. "
		h.logger.Error(returnErr)
		return
	}

	author, err := h.client.GetAuthor(r.Form.Get("code"))
	if err != nil {
		returnErr = "Failed to retrieve user data from Google."
		h.logger.Error(returnErr, zap.Error(err), zap.String("code", r.Form.Get("code")))
		return
	}

	err = h.store.WithSession(func(s *db.Session) error {
		upserted, err := s.UpsertAuthor(author)
		if err != nil {
			h.logger.Error(returnErr, zap.Error(err))
			return err
		}
		returnToken = "jwt"+upserted.Email
		return nil
	});
	if err != nil {
		returnErr = "Failed to register local user."
		h.logger.Error(returnErr, zap.Error(err))
	}
}
