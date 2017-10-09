package main

import (
	"flag"
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
	var port = flag.Int("port", 9090, "port to be listened on")
	flag.Parse()

	loadSignups()
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
