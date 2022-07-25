package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"

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

	fmt.Printf("Initialized tables")

}


func exampData() {

	db, err := data.Conn()
	if err != nil {
		panic("Couldn't connect to DB")
	}

	var users []data.User
	db.Find(&users)

	if len(users) > 0 {
		return
	}

	uid := data.GetUUID()

	u := data.User{
		Id: uid,

		Name: "Ama Threeple",
		Email: "ama@example.com",
		Phone: "+13333333333",
		Address: "",
	}


	db.Create(&u)

	u_t := data.User{
		Id: data.GetUUID(),

		Name: "WB",
		Email: "wb@example.com",
		Phone: "+1***REMOVED***",
		Address: "",

		PreviewCredits: 2,
		WeekendCredits: 5,
		WeekdayCredits: 1,
	}

	db.Create(&u_t)

	sid := data.GetUUID()

	s := data.Show{
		Id: sid,

		Name: "Three Sisters",
		PreviewPrice: 25.00,
		WeekendPrice: 50.00,
		WeekdayPrice: 40.00,
	}

	db.Create(&s)

	tid := data.GetUUID()

	t := data.Transact{
		Id: tid,

		Quantity: 2,
		Rate: 25.00,

		UserId: uid,
		ShowId: &sid,
	}

	db.Create(&t)

	a1 := data.Adjustment{
		Id: data.GetUUID(),

		DiscountCode: "SIBLING10",

		Multiplier: 0.9,
		TransactId: tid,
	}

	db.Create(&a1)

	a2 := data.Adjustment{
		Id: data.GetUUID(),

		DiscountCode: "MOSCOW2",

		Additive: -2.50,
		TransactId: tid,
	}

	db.Create(&a2)

	tid2 := data.GetUUID()

	t2 := data.Transact{
		Id: tid2,

		Credit: 5,
		CreditType: data.WEEKEND,

		UserId: uid,
	}

	db.Create(&t2)


}


func home(w http.ResponseWriter, r * http.Request) {

	fmt.Fprintf(w, "Hi\n")
}


func HelloWord(w http.ResponseWriter, r * http.Request) {

	phone_l := r.URL.Query()["phone"]

	var phone string
	if len(phone_l) < 1 {
		phone = "+1***REMOVED***"
	} else {
		phone = phone_l[0]
	}

	err := SendHello(phone)

	if err != nil {
		panic(err)
	}

}


func CallbackCheck(w http.ResponseWriter, r * http.Request) {

	phone_l := r.URL.Query()["phone"]

	var phone string
	if len(phone_l) < 1 {
		phone = "+1***REMOVED***"
	} else {
		phone = phone_l[0]
	}

	_, err := Send(phone, "Do you want a callback?", true)

	if err != nil {
		panic(err)
	}

}

func SmsCallback(w http.ResponseWriter, r * http.Request) {

	orig_msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	vals, err := ParseBody(string(orig_msg))
	if err != nil {
		panic(err)
	}

	fmt.Println(vals)

	ret_msg, err := BasicResp(vals["Body"][0], vals["From"][0])

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, ret_msg)
}



func main() {

	initDb()
	exampData()


	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/hello", HelloWord)
	mux.HandleFunc("/callback", CallbackCheck)
	mux.HandleFunc("/sms", SmsCallback)

	port := os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}

	portstr := fmt.Sprintf(":%s", port)

	http.ListenAndServe(portstr,mux)

}
