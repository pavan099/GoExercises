package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", readCookies)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func readCookies(writer http.ResponseWriter, request *http.Request) {
	c1, err := request.Cookie("visited")
	if err != nil {
		_, _ = fmt.Fprintln(writer, "Could not retrieve cookie details", err)
	}
	_, _ = fmt.Fprintln(writer, "The site visited:", c1.Value)
}

var visitedCount int

func set(writer http.ResponseWriter, _ *http.Request) {
	visitedCount++

	http.SetCookie(writer, &http.Cookie{Name: "visited", Value: strconv.Itoa(visitedCount)})

	_, _ = fmt.Fprintln(writer, "Visitor cookie is set")

}
