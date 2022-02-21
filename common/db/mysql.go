package db

import (
	"fmt"

	//	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MainSession *MySQL

func init() {
	var err error
	MainSession, err = NewMySQL("root", "", "127.0.0.1", "aha")
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
	*gorm.DB
}

//NewMaria instance Maria struct
//initialize a new db connection
//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
//host ip:port
func NewMySQL(user, password, host, dbName string) (*MySQL, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=UTC&timeout=30s", user, password, host, dbName)
	//"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=UTC&timeout=30s"
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	return &MySQL{
		db,
	}, err
}
