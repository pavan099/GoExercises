package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/html/home", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("/public" + r.URL.Path + ".html")
		if err != nil {
			w.Write([]byte("File not found"))
		}
		defer f.Close()

	})
}
