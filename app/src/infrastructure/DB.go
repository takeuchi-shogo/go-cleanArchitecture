package infrastructure

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName:   c.DB.Production.DBName,
	})
}

func newDB(d *DB) *DB {
	count := 0
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err.Error())
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 10 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
		}
	}
	fmt.Println("DB接続成功")

	db.LogMode(true)

	d.Connection = db

	//db.AutoMigrate()
	return d
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
