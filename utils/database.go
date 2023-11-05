package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "attendance-is/models"
)

var DB *gorm.DB

func Connection(user string, pass string, host string, port string, dbname string) error {
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		DB = db
		return err
	} else {
		return nil
	}
}

func Migrate() {
	DB.Debug().AutoMigrate(
		&model.Absent{},
		&model.Class{},
		&model.Event{},
		&model.Excuse{},
		&model.Kompen{},
		&model.Lecturer{},
		&model.SP{},
		&model.Student{},
		&model.User{},
	)
}
