package db

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MainSession *MySQL

func InitDB(account, password, socketDir, instanceConnectionName, dbName string) {

	var err error
	//MainSession, err = NewMySQL(account, password, host, dbName)
	MainSession, err = NewMySQLGCP(account, password, socketDir, instanceConnectionName, dbName)
	if err != nil {
		panic("sql init: " + err.Error())
	}
}

const (
	//DIALECT
	DIALECT = "mysql"
)

//Maria ...
type MySQL struct {
	Gorm *gorm.DB
}

//NewMaria instance Maria struct
//initialize a new db connection
//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
//host ip:port
func NewMySQL(user, password, host, dbName string) (*MySQL, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=UTC&timeout=30s", user, password, host, dbName)
	//"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=UTC&timeout=30s"
	//dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	fmt.Println("db args", args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	return &MySQL{
		db,
	}, err
}

//
func NewMySQLGCP(user, password, socketDir, instanceConnectionName, dbName string) (*MySQL, error) {
	//	args := fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8&parseTime=true&loc=UTC&timeout=30s", user, password, socketDir, instanceConnectionName, dbName)
	//"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=UTC&timeout=30s"
	//dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	args := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", user, password, socketDir, instanceConnectionName, dbName)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	return &MySQL{
		db,
	}, err
}
