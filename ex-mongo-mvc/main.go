package main

import (
	"log"
	"net/http"

	"github.com/chandanghosh/goexercises/ex-mongo-mvc/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	usrController := controllers.NewUserController()
	router.GET("/", usrController.GetUsers)
	router.GET("/user/:id", usrController.GetUser)
	router.POST("/user/create", usrController.CreateUser)
	router.PUT("/user/update", usrController.UpdateUser)
	router.DELETE("/user/delete/:id", usrController.DeleteUser)
	log.Fatalln(http.ListenAndServe(":8080", router))
}
