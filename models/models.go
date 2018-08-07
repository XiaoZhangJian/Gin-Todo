package models

import (
	"Gin-Todo/pkg/setting"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedAt  int `json:"created_at"`
	ModifiedAt int `json:"modified_at"`
}

var db *gorm.DB

func init() {

	var err error
	sec, err := setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatalf("fail to get section 'database' : %v", err)
	}

	dbType := sec.Key("TYPE").String()
	user := sec.Key("USER").String()
	pwd := sec.Key("PASSWORD").String()

	host := sec.Key("HOST").String()
	dbName := sec.Key("NAME").String()
	tablePrefix := sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		dbName,
	))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

func Clone() {
	defer db.Close()
}
