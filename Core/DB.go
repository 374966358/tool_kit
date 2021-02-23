package Core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var Config = struct {
	DB struct {
		Name         string
		Host         string
		User         string
		Pass         string
		Port         string        `default:"3306"`
		MaxIdleConns int           `default:8`
		MaxOpenConns int           `default:16`
		MaxLifetime  time.Duration `default:120`
	}
}{}

var (
	DB  *gorm.DB
	err error
)

func GetDB(name string) {
	configor.Load(&Config, "config.yml")
	if name != "" {
		Config.DB.Name = name
	}
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.DB.User, Config.DB.Pass, Config.DB.Host, Config.DB.Port, Config.DB.Name))
	if err != nil {
		log.Println(err)
	}
	// 是否开启调试模式
	// DB.LogMode(true)
	DB.DB().SetMaxIdleConns(Config.DB.MaxIdleConns)
	DB.DB().SetMaxOpenConns(Config.DB.MaxOpenConns)
	DB.DB().SetConnMaxLifetime(time.Second * Config.DB.MaxLifetime)
}
