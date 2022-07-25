package main

import (
	"fmt"
	"os"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
	"strings"
	"strconv"

	data "github.com/wbrackenbury/NowLive/m/v2/data"

)


const TwilFile string = "creds.json"
const UrlBase string = "https://api.twilio.com/2010-04-01/Accounts/"

type TwilInfo struct{

	AccountSid string `json:"account_sid"`
	AuthToken string `json:"auth_token"`
	TwilPhone string `json:"twil_phone"`

}

type TwimlResp struct {
	XMLName xml.Name `xml:"Response"`
	Message string `xml:"Message"`
}


func local_loadTwilInfo() (*TwilInfo, error) {

	conf, err := ioutil.ReadFile(TwilFile)
	if err != nil {
		panic(err)
	}

	conf_info := TwilInfo{}
	err = json.Unmarshal([]byte(conf), &conf_info)
	if err != nil {
		return nil, err
	}

	return &conf_info, nil

}

func loadTwilInfo() (*TwilInfo, error) {


	tf := TwilInfo{
		AccountSid: os.Getenv("ACCOUNT_SID"),
		AuthToken: os.Getenv("AUTH_TOKEN"),
		TwilPhone: os.Getenv("TWIL_PHONE"),
	}


	return &tf, nil

}


func decorateReq(req *http.Request, tf *TwilInfo) {

	(*req).SetBasicAuth(tf.AccountSid, tf.AuthToken)
	(*req).Header.Add("Accept", "application/json")
	(*req).Header.Add("Content-Type", "application/x-www-form-urlencoded")

}

func sendUrl(tf * TwilInfo) string {
	return UrlBase + tf.AccountSid + "/Messages.json"
}

func Send(num, message string, callback bool) (*http.Response, error) {

	tf, err := loadTwilInfo()
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Set("To", num)
	v.Set("From", tf.TwilPhone)
	v.Set("Body", message)

	if callback {
		v.Set("StatusCallback", "https://theatre-now-live.herokuapp.com/sms")
	}

	rb := *strings.NewReader(v.Encode())

	fmt.Printf("Rb: %s\n", v)

	req, err := http.NewRequest("POST", sendUrl(tf), &rb)
	if err != nil {
		return nil, err
	}
	decorateReq(req, tf)

	fmt.Printf("%s", req)

	client := &http.Client{}
	resp, err := client.Do(req)

	fmt.Println(resp)
	fmt.Println(resp.Status)

	return resp, err

}


func SendHello(num string) (error) {

	_, err := Send(num, "Hello!", false)
	return err
}


func handleCheck(num string) (string, error) {

	m, err := ioutil.ReadFile("messages/getcredit")
	if err != nil {
		return "", nil
	}

	prev, weekd, weekend := data.NumCredits(num)

	return fmt.Sprintf(string(m), prev, weekd, weekend), nil

}

func handleRunShows() (string, error) {

	m, err := ioutil.ReadFile("messages/shows")
	if err != nil {
		return "", nil
	}

	shows := data.RunningShows()

	if len(shows) < 1 {
		return "No shows are currently playing.", nil
	}

	ret_str := string(m)

	for _, s := range shows {
		ret_str += s
	}

	return ret_str, nil

}

func handleAddCredits(orig_msg, num string) (string, error) {


	vals := strings.Split(orig_msg, " ")
	if len(vals) < 2 {

		m, err := ioutil.ReadFile("messages/add_error")
		return string(m), err
	}

	nc, err := strconv.Atoi(vals[1])
	if err != nil {
		m, err := ioutil.ReadFile("messages/add_error")
		return string(m), err

	}

	ctype := vals[0]
	err = data.AddCredits(num, ctype, nc)

	if err != nil {
		return "", err
	}

	m, err := ioutil.ReadFile("messages/add_success")
	if err != nil {
		return "Credits added", nil
	}

	return fmt.Sprintf(string(m), nc, ctype), nil

}


func BasicResp(orig_msg, num string) (string, error) {

	tr := &TwimlResp{}

	switch  {

	case orig_msg == "CHECK":

		s, err := handleCheck(num)
		if err != nil {
			panic(err)
		}

		tr.Message = s

	case orig_msg == "SHOWS":

		s, err := handleRunShows()
		if err != nil {
			panic(err)
		}

		tr.Message = s

	case orig_msg == "ADD":

		m, err := ioutil.ReadFile("messages/add")
		if err != nil {
			panic(err)
		}

		tr.Message = string(m)

	// The below is not best practice. Technically, we would
	// want to store some state in the User table, and then
	// check whether there's a live session for this part, but
	// this is sufficient for a hobby project
	case (strings.HasPrefix(orig_msg, "PREVIEW") ||
		strings.HasPrefix(orig_msg, "WEEKDAY") ||
		strings.HasPrefix(orig_msg, "WEEKEND")):

		s, err := handleAddCredits(orig_msg, num)
		if err != nil {
			panic(err)
		}

		tr.Message = s

	default:

		m, err := ioutil.ReadFile("messages/help")
		if err != nil {
			panic(err)
		}

		tr.Message = string(m)
	}

	s, err := xml.Marshal(tr)
	if err != nil {
		return "", err
	}

	return xml.Header + string(s), nil


}
