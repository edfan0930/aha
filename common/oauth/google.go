package oauth

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	jsoniter "github.com/json-iterator/go"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	GoogleUserInfoURL = "https://www.googleapis.com/oauth2/v3/userinfo"
	GoogleRedirectURL = "http://localhost:3000/callback/google"
)

type (
	GoogleOauth2 struct {
		Config   *oauth2.Config
		Token    *oauth2.Token
		Response struct {
			Sub           string `json:"sub"`
			Name          string `json:"name"`
			GiveName      string `json:"give_name"`
			FamilyName    string `json:"family_name"`
			Picture       string `json:"picture"`
			Email         string `json:"email" validate:"email"`
			EmailVerified bool   `json:"email_verified"`
			Locale        string `json:"locale"`
			HD            string `json:"hd"`
		}
	}

	Token struct {
		// AccessToken is the token that authorizes and authenticates
		// the requests.
		AccessToken string `json:"access_token"`

		// TokenType is the type of token.
		// The Type method returns either this or "Bearer", the default.
		TokenType string `json:"token_type,omitempty"`

		// RefreshToken is a token that's used by the application
		// (as opposed to the user) to refresh the access token
		// if it expires.
		RefreshToken string `json:"refresh_token,omitempty"`

		// Expiry is the optional expiration time of the access token.
		//
		// If zero, TokenSource implementations will reuse the same
		// token forever and RefreshToken or equivalent
		// mechanisms for that TokenSource will not be used.
		Expiry time.Time `json:"expiry,omitempty"`

		// raw optionally contains extra metadata from the server
		// when updating a token.
		raw interface{}
	}
)

var googleConfig = &oauth2.Config{
	//憑證的 client_id t
	ClientID: "479205503773-r674qa8u7b186hbupe43oimrrc9mrhga.apps.googleusercontent.com",
	//憑證的 client_secret
	ClientSecret: "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",
	//當 Google auth server 驗證過後，接收從 Google auth server 傳來的資訊
	RedirectURL: GoogleRedirectURL,
	//告知 Google auth server 授權範圍，在這邊是取得用戶基本資訊和Email，Scopes 為 Google 提供
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	//指的是 Google auth server 的 endpoint，用 lib 預設值即可
	Endpoint: google.Endpoint,
}

//NewGoogleOauth2
func NewGoogleOauth2() *GoogleOauth2 {
	return &GoogleOauth2{
		Config: googleConfig,
	}
}

func GoogleOAuthURL() string {

	return googleConfig.AuthCodeURL(UUID)
}

//Exchange
func (g *GoogleOauth2) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) error {

	opts = append(opts, oauth2.AccessTypeOffline)
	return Exchange(g.Config, ctx, code, g.Token, opts...)
}

//Client
func (g *GoogleOauth2) Request(ctx context.Context) error {

	client := g.Config.Client(ctx, g.Token)
	res, getErr := client.Get(GoogleUserInfoURL)
	if getErr != nil {

		return fmt.Errorf("client request error: %w", getErr)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {

		return errors.New(fmt.Sprintf("request http status code: %d", res.StatusCode))
	}

	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io reader error: %w", err)
	}

	return jsoniter.Unmarshal(rawData, &g.Response)
}

/* func GoogleClient(ctx context.Context, token *oauth2.Token) {
	client := googleConfig.Client(ctx, token)
	res, getErr := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if getErr != nil {
		fmt.Println("getErr:", getErr)
		return
	}
	fmt.Println("res", res)
	defer res.Body.Close()
	rawData, _ := ioutil.ReadAll(res.Body)
	fmt.Println("rawData:", string(rawData))
}
*/
