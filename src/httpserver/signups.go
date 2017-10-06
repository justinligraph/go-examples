package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Signup struct {
	Name string
}

const signupPath = "data/signups"

var gSignups = make([]*Signup, 0, 5)

func loadSignups() {
	if _, err := os.Stat(signupPath); err == nil {
		if data, err := ioutil.ReadFile(signupPath); err == nil {
			allnames := string(data)
			names := strings.Split(allnames, "\n")
			for _, name := range names {
				gSignups = append(gSignups, &Signup{Name: name})
			}
		} else {
			log.Printf("Failed to read file %s:%v", signupPath, err)
		}
	} else {
		log.Printf("Faile to stat %s:%v", signupPath, err)
	}
}

func getSignups() []*Signup {
	return gSignups
}

func addSignup(name string) {
	gSignups = append(gSignups, &Signup{Name: name})

	f, err := os.OpenFile(signupPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Cannot open file %s: %v\n", signupPath, err)
		return
	}
	defer f.Close()
	if n, err := f.WriteString(name + "\n"); err != nil {
		log.Printf("Cannot write to file %s:%v", signupPath, err)
	} else {
		log.Printf("Write to file %s %d bytes", signupPath, n)
	}
}
