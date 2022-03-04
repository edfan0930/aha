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

	env.SetENV()
	email.InitEmail(env.Email)
	db.InitDB(env.DBAccount, env.DBPassword, env.DBHost, env.DBName)
	oauth2.SetProvider(env.ServerDomain + "/callback")
	oauth2.SetStore(storage.Store.Key)

	router.InitRouter()

}
