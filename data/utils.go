package data

import (

	//"database/sql"
	"time"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"github.com/google/uuid"
	"gorm.io/datatypes"

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


	//db, err := gorm.Open(sql.Open("postgres", os.Getenv("DATABASE_URL")), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, err
}

func Conn() (*gorm.DB, error) {

	return pgConn()
	//return sqliteConn()
}


func ExampData() {

	db, err := Conn()
	if err != nil {
		panic("Couldn't connect to DB")
	}

	var users []User
	db.Find(&users)

	if len(users) > 0 {
		return
	}

	uid := GetUUID()

	u := User{
		Id: uid,

		Name: "Ama Threeple",
		Email: "ama@example.com",
		Phone: "+13333333333",
		Address: "",
	}


	db.Create(&u)

	u_t := User{
		Id: GetUUID(),

		Name: "WB",
		Email: "wb@example.com",
		Phone: "+1***REMOVED***",
		Address: "",

		PreviewCredits: 2,
		WeekendCredits: 5,
		WeekdayCredits: 1,
	}

	db.Create(&u_t)

	sid := GetUUID()

	start, _ := time.Parse("01-02-06", "3-05-22")
	end, _ := time.Parse("01-02-06", "11-21-22")

	s := Show{
		Id: sid,

		Name: "Three Sisters",
		PreviewPrice: 25.00,
		WeekendPrice: 50.00,
		WeekdayPrice: 40.00,

		StartDate: datatypes.Date(start),
		EndDate: datatypes.Date(end),

	}

	db.Create(&s)

	tid := GetUUID()

	t := Transact{
		Id: tid,

		Quantity: 2,
		Rate: 25.00,

		UserId: uid,
		ShowId: &sid,
	}

	db.Create(&t)

	a1 := Adjustment{
		Id: GetUUID(),

		DiscountCode: "SIBLING10",

		Multiplier: 0.9,
		TransactId: tid,
	}

	db.Create(&a1)

	a2 := Adjustment{
		Id: GetUUID(),

		DiscountCode: "MOSCOW2",

		Additive: -2.50,
		TransactId: tid,
	}

	db.Create(&a2)

	tid2 := GetUUID()

	t2 := Transact{
		Id: tid2,

		Credit: 5,
		CreditType: WEEKEND,

		UserId: uid,
	}

	db.Create(&t2)


}
