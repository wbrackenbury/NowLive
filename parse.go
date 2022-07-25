package main

import (
	"strings"
)


// The below is pretty awful, and would likely fail under
// even modest constraints (e.g., what if your message has an ampersand?)
func GetBody(req_body string, get_num bool) (string) {

	s_split := strings.Split(req_body, "&")
	var sub_l []string
	var attr, val string
	for _, s := range s_split {
		sub_l = strings.Split(s, "=")
		attr = sub_l[0]
		val = sub_l[1]

		if attr == "Body" && !get_num {
			return val
		} else if attr == "From" && get_num {
			return val
		}
	}

	return ""

}
