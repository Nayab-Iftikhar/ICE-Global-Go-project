package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB(config *gorm.Config) *gorm.DB {
	dsn := "root:password@tcp(mysql:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
