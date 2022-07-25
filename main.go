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
	data.ExampData()

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
