package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"clean-arch/app/config"
	_userModel "clean-arch/features/users/data"
	_bookModel "clean-arch/features/books/data"
)

func InitDB(cfg config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	fmt.Println(connectionString)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to DB", err.Error())
		return nil
	}

	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userModel.UserGorm{}, &_bookModel.BookGorm{})
}
