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


func BasicResp(orig_msg string) (string, error) {

	tr := &TwimlResp{}

	switch orig_msg {

	case "hello":
		tr.Message = "hi"
	case "bye":
		tr.Message = "bye"
	default:
		tr.Message = "No message"
	}

	s, err := xml.Marshal(tr)
	if err != nil {
		return "", err
	}

	return xml.Header + string(s), nil


}
