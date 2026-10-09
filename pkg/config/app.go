package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		user := os.Getenv("MYSQL_USER")
		if user == "" {
			user = "root"
		}
		pass := os.Getenv("MYSQL_PASSWORD")
		if pass == "" {
			pass = "password" // placeholder; set MYSQL_PASSWORD or MYSQL_DSN in your environment
		}
		host := os.Getenv("MYSQL_HOST")
		if host == "" {
			host = "127.0.0.1"
		}
		port := os.Getenv("MYSQL_PORT")
		if port == "" {
			port = "3306"
		}
		dbname := os.Getenv("MYSQL_DBNAME")
		if dbname == "" {
			dbname = "simplerest"
		}
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, dbname)
	}
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
