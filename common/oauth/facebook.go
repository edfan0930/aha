package oauth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

const (
	FBInfoURL           = "https://graph.facebook.com/me?access_token="
	FacebookRedirectURL = "http://localhost:3000/callback/facebook"
)

type (
	FacebookOauth2 struct {
		Config   *oauth2.Config
		Token    *oauth2.Token
		Response struct {
		}
	}
)

var facebookConfig = &oauth2.Config{
	//憑證的 client_id
	ClientID: "844823353051244",
	//憑證的 client_secret
	ClientSecret: "83fd5fd6bf47f4f0808fc6109427519d",
	//當 Google auth server 驗證過後，接收從 Google auth server 傳來的資訊
	RedirectURL: FacebookRedirectURL,
	//告知 Google auth server 授權範圍，在這邊是取得用戶基本資訊和Email，Scopes 為 Google 提供
	Scopes: []string{
		"email",
		"public_profile",
	},
	//指的是 Google auth server 的 endpoint，用 lib 預設值即可
	Endpoint: facebook.Endpoint,
}

//NewFacebookOauth2
func NewFacebookOauth2() *FacebookOauth2 {
	return &FacebookOauth2{
		Config: facebookConfig,
	}
}

func FackbookOAuthURL() string {
	return facebookConfig.AuthCodeURL(UUID)
}

func (f *FacebookOauth2) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) error {

	var err error
	f.Token, err = Exchange(f.Config, ctx, code, opts...)
	fmt.Println("fb token:", f.Token)
	fmt.Println("facebook exchange", err)
	return err
}

func (f *FacebookOauth2) Request(ctx context.Context) error {

	client := f.Config.Client(ctx, f.Token)
	res, getErr := client.Get(GoogleUserInfoURL)
	if getErr != nil {

		return fmt.Errorf("client request error: %w", getErr)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {

		return fmt.Errorf(fmt.Sprintf("request http status code: %d", res.StatusCode))
	}

	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io reader error: %w", err)
	}
	fmt.Println("rawData", string(rawData))
	return nil
}
func FacebookClient(ctx context.Context, token *oauth2.Token) {
	client := facebookConfig.Client(ctx, token)
	res, getErr := client.Get(FBInfoURL + url.QueryEscape(token.AccessToken))
	if getErr != nil {
		fmt.Println("getErr:", getErr)
		return
	}
	fmt.Println("fb client")
	fmt.Println("res", res)
	defer res.Body.Close()
	rawData, _ := ioutil.ReadAll(res.Body)
	fmt.Println("rawData:", string(rawData))
}
