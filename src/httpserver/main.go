package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		addSignup(r.Form["name"][0])
	}
	signups := getSignups()
	signupStr := ""
	for _, s := range signups {
		signupStr = signupStr + s.Name + "<br/>"
	}

	fmt.Fprintf(w, index_html, signupStr)
}

func main() {
	loadSignups()
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":9090", nil)
}
