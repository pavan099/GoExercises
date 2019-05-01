package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dogFunc)
	log.Panicln(http.ListenAndServe(":8080", nil))
}

func dogFunc(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("content-type", "text/html;charset=utf8")
	_, _ = io.WriteString(writer, `<img src="dog.jpg"" />`)
}
