package data

import (

	"os"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"github.com/google/uuid"

)

const db_fname string = "nowlive.db"

func GetUUID() string {
	return uuid.New().String()
}


func sqliteConn() (*gorm.DB, error){

	db, err := gorm.Open(sqlite.Open(db_fname), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, err
}


func pgConn() (*gorm.DB, error){


	// conf_info, err := loadDbInfo()
	// if err != nil {
	// 	panic(err)
	// }

	s := dsnString(conf_info)
	db, err := gorm.Open(sql.Open("postgres", os.Getenv("DATABASE_URL")), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(s), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, err
}

func Conn() (*gorm.DB, error) {

	return pgConn()
	//return sqliteConn()
}
