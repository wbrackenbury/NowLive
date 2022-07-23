package main

import (
	"fmt"
	"net/http"

	data "github.com/wbrackenbury/NowLive/m/v2/data"
)




func initDb() {
	db, err := data.Conn()
	if err != nil {
		panic("Couldn't connect to DB")
	}
	db.AutoMigrate(&data.Adjustment{})
	db.AutoMigrate(&data.Transact{})
	db.AutoMigrate(&data.Show{})
	db.AutoMigrate(&data.User{})

}


func exampData() (error) {

	db, err := data.Conn()
	if err != nil {
		panic("Couldn't connect to DB")
	}

	var users []File
	db.Find(&users)

	if len(users) > 0 {
		return nil
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

	sid := GetUUID()

	s := Show{
		Id: sid,

		Name: "Three Sisters",
		PreviewPrice: 25.00,
		WeekendPrice: 50.00,
		WeekDayPrice: 40.00,
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
		Id: tid,

		Credit: 5,
		CreditType: data.WEEKEND,

		UserId: uid,
	}

	db.Create(&t2)


}


func home(w http.ResponseWriter, r * http.Request) {

	fmt.Fprintf(w, "Hi\n")
}



func main() {

	initDb()
	exampData()

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	//mux.HandleFunc("/", home)

	http.ListenAndServe(":8080",mux)

}
