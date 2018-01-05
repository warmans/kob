package google

// https://skarlso.github.io/2016/06/12/google-signin-with-go/

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	"github.com/warmans/kob/server/pkg/rpc"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

func NewClient(credsPath string, redirectURI string) (*Client, error) {

	file, err := ioutil.ReadFile(credsPath)
	if err != nil {
		return nil, err
	}
	cred := &Credentials{}
	if err := json.Unmarshal(file, &cred); err != nil {
		return nil, err
	}

	conf := &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  redirectURI,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}

	return &Client{conf: conf}, nil
}

type Client struct {
	conf *oauth2.Config
}

func (c *Client) CreateLoginURL() (string, string) {
	code := CreateCode()
	return c.conf.AuthCodeURL(code), code
}

func (c *Client) GetAuthor(code string) (*rpc.Author, error) {

	tok, err := c.conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	client := c.conf.Client(oauth2.NoContext, tok)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

    author := &rpc.Author{}
	if err := json.NewDecoder(resp.Body).Decode(author); err != nil {
		return nil, err
	}
	return author, nil
}

func CreateCode() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
