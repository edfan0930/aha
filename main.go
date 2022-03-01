package main

import (
	"github.com/edfan0930/aha/common/oauth2"
	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/router"
)

func main() {
	oauth2.SetProvider("http://localhost:3000/callback")
	oauth2.SetStore(storage.Store.Key)

	router.InitRouter()

}
