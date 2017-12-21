package oauth

// https://skarlso.github.io/2016/06/12/google-signin-with-go/

import (
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
    "os"
    "net/http"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

// Credentials which stores google ids.
type Credentials struct {
    Cid     string `json:"cid"`
    Csecret string `json:"csecret"`
}

// User is a retrieved and authentiacted user.
type User struct {
    Sub string `json:"sub"`
    Name string `json:"name"`
    GivenName string `json:"given_name"`
    FamilyName string `json:"family_name"`
    Profile string `json:"profile"`
    Picture string `json:"picture"`
    Email string `json:"email"`
    EmailVerified string `json:"email_verified"`
    Gender string `json:"gender"`
}

func NewClient(credsPath string, redirectURI string) *Client {

    file, err := ioutil.ReadFile(credsPath)
    if err != nil {
		return nil, err
    }
    json.Unmarshal(file, &cred)

    conf :=  &oauth2.Config{
        ClientID:     cred.Cid,
        ClientSecret: cred.Csecret,
        RedirectURL:  redirectURI,
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
        },
        Endpoint: google.Endpoint,
    }

    return &Client{conf: conf}
}

type Client struct {
    conf *oauth2.Config
}

func (c *Client) CreateLoginURL() (string, string) {
    code := CreateCode()
    return c.conf.AuthCodeURL(code), code
}

func CreateCode() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

