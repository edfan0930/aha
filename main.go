package main

import (
	"github.com/edfan0930/aha/common/email"
	"github.com/edfan0930/aha/common/oauth2"
	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/env"
	"github.com/edfan0930/aha/router"
)

func main() {

	//set env
	env.SetENV()
	//email set global
	email.InitEmail(env.Email, env.ServerDomain+router.PathVerfication)
	//database init
	db.InitDB(env.DBAccount, env.DBPassword, env.SocketDir, env.DBConnectName, env.DBName)
	//set oauth2 provider callback URI
	oauth2.SetProvider(env.ServerDomain + "/callback")
	//set session key
	oauth2.SetStore(storage.Store.Key)
	//router init
	router.InitRouter()
}
