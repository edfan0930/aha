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
	FBInfoURL = "https://graph.facebook.com/me?access_token="
)

var facebookConfig = &oauth2.Config{
	//憑證的 client_id
	ClientID: "633076937906328",
	//憑證的 client_secret
	ClientSecret: "872403eae1782bb895f2e4460608e660",
	//當 Google auth server 驗證過後，接收從 Google auth server 傳來的資訊
	RedirectURL: "http://localhost:3000/main",
	//告知 Google auth server 授權範圍，在這邊是取得用戶基本資訊和Email，Scopes 為 Google 提供
	Scopes: []string{
		"email",
		"public_profile",
	},
	//指的是 Google auth server 的 endpoint，用 lib 預設值即可
	Endpoint: facebook.Endpoint,
}

func FackbookOAuthURL() string {
	return facebookConfig.AuthCodeURL("random")
}

func FacebookExchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {

	return facebookConfig.Exchange(ctx, code, opts...)
}

func FacebookClient(ctx context.Context, token *oauth2.Token) {
	fmt.Println("in fb")
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
