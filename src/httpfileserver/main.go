package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
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

func main() {
	var port = flag.Int("port", 9090, "port to be listened on")
	var tgtPath = flag.String("target", "tigergraph", "Path to write file to")
	flag.Parse()
	targetPath = *tgtPath

	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
