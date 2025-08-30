package ioc

import (
	"github.com/spf13/viper"
	"github.com/webook-project-go/webook-comment/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&dao.Comment{})
	if err != nil {
		panic(err)
	}
	return db
}
