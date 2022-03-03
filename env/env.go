package env

import (
	"github.com/spf13/viper"
)

var (
	ServerDomain  string
	GoogleKey     string
	GoogleSecret  string
	FBKey         string
	FBSecret      string
	Email         string
	EmailPassword string
	DBAccount     string
	DBPassword    string
	DBHost        string
	DBName        string
)

func ReadConfig() error {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func SetENV() {

	if err := ReadConfig(); err != nil {

		panic("read config failed: " + err.Error())
	}

	ServerDomain = emptyPanic("host")
	GoogleKey = emptyPanic("oauth2.googleKey")

	GoogleSecret = emptyPanic("oauth2.googleSecret")
	FBKey = emptyPanic("oauth2.fbKey")
	FBSecret = emptyPanic("oauth2.fbSecret")
	Email = emptyPanic("email.account")
	EmailPassword = emptyPanic("email.password")
	DBAccount = emptyPanic("db.account")
	DBPassword = emptyPanic("db.password")
	DBHost = emptyPanic("db.host")
	DBName = emptyPanic("db.name")
}

//emptyPanic
func emptyPanic(key string) string {
	v := viper.GetString(key)
	if v == "" {
		panic("config doesn't has:" + key)
	}
	return v
}
