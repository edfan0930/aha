package env

import (
	"log"
	"os"

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
	DBConnectName string
	SocketDir     string
)

//ReadConfig read in config
func ReadConfig() error {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

//SetENV set env
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
	DBAccount = mustGetenv("DB_USER")
	DBPassword = mustGetenv("DB_PASS")
	DBHost = emptyPanic("db.host")
	DBConnectName = mustGetenv("INSTANCE_CONNECTION_NAME")
	DBName = mustGetenv("DB_NAME")

	var isSet bool
	SocketDir, isSet = os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		SocketDir = "/cloudsql"
	}
}

//emptyPanic get config value
func emptyPanic(key string) string {
	v := viper.GetString(key)
	if v == "" {
		panic("config doesn't has:" + key)
	}
	return v
}

//mustGetenv get os env value
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
