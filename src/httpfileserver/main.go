package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	targetPath = "tigergraph"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)
		file, fHeader, err := r.FormFile("uploadfile")
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			log.Printf("Writing file %s to %s...\n", fHeader.Filename, targetPath)
			defer file.Close()
			if f, err := os.OpenFile(targetPath+"/"+fHeader.Filename, os.O_WRONLY|os.O_CREATE, 0600); err != nil {
				w.Write([]byte(err.Error()))
			} else {
				defer f.Close()
				if len, err := io.Copy(f, file); err != nil {
					w.Write([]byte(err.Error()))
				} else {
					w.Write([]byte(fmt.Sprintf("File uploaded successfully. %d bytes written", len)))
				}
			}
		}
	} else {
		w.Write([]byte("That's tickle."))
	}
}

var commands = map[string][]string{
	"gpe_reinit": []string{"./ginspect_cmd.sh", "./statushubcli", "poc_gpe_server", "gpe_reinit", "20"},
	"ls":         []string{"ls"},
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if _, ok := r.Form["cmd"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("cmd is not specified."))
	} else {
		cmds := r.Form["cmd"]
		cmdlines := make([][]string, 0, 10)
		for _, cmd := range cmds {
			if cmdline, ok := commands[cmd]; !ok {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("cmd (%s) is not specified.", cmd)))
				return
			} else {
				cmdlines = append(cmdlines, cmdline)
			}
		}
		for _, cmdln := range cmdlines {
			log.Printf("Executing %s...\n", cmdln)
			cmd := exec.Command(cmdln[0], cmdln[1:]...)
			if output, err := cmd.CombinedOutput(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			} else {
				w.WriteHeader(http.StatusAccepted)
				w.Write(output)
			}
		}
	}
}

func main() {
	var port = flag.Int("port", 9090, "port to be listened on")
	var tgtPath = flag.String("target", "tigergraph", "Path to write file to")
	flag.Parse()
	targetPath = *tgtPath

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/cmd", cmdHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Printf("ListenAndServe failed with %v", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
