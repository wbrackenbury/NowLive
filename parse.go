package main

import (
	"strings"
)


// The below is pretty awful, and would likely fail under
// even modest constraints (e.g., what if your message has an ampersand?)
func GetBody(req_body string) (string) {

	s_split := strings.Split(req_body, "&")
	var sub_l []string
	var attr, val string
	for _, s := range s_split {
		sub_l = strings.Split(s, "=")
		attr = sub_l[0]
		val = sub_l[1]

		if attr == "Body" {
			return val
		}
	}

	return ""

}
