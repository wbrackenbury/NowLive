package data

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/google/uuid"

)

const db_fname string = "nowlive.db"

func GetUUID() string {
	return uuid.New().String()
}


func Conn() (*gorm.DB, error){

	db, err := gorm.Open(sqlite.Open(db_fname), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, err
}
