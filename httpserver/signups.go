package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type Signup struct {
	Name string
}

const signupPath = "/data/signups"

var gSignups = make([]*Signup, 0, 5)

func loadSignups() {
	if _, err := os.Stat(signupPath); err == nil {
		if data, err := ioutil.ReadFile(signupPath); err != nil {
			allnames := string(data)
			names := strings.Split(allnames, "\n")
			for _, name := range names {
				gSignups = append(gSignups, &Signup{Name: name})
			}
		}
	}
}

func getSignups() []*Signup {
	return gSignups
}

func addSignup(name string) {
	gSignups = append(gSignups, &Signup{Name: name})

	f, err := os.OpenFile(signupPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(name + "\n")
}
