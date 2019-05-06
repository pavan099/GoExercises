package main

import (
	"encoding/json"
	"github.com/chandanghosh/goexercises/ex-mongo-mvc/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	log.Fatalln(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	usr := models.User{Name: "Chandan Ghosh", Age: 35, Gender: "M", Id: "1"}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(usr)
	if err != nil {
		log.Println(err.Error())
	}
}
