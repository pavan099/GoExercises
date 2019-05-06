package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chandanghosh/goexercises/ex-mongo-mvc/models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

var users = map[string]models.User{}

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err.Error())
	}
}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if usr, ok := users[id]; ok {
		_ = json.NewEncoder(w).Encode(&usr)
		return
	}
}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if _, ok := users[id]; ok {
		delete(users, id)
	}
	fmt.Fprintln(w, "User deleted")
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	usr := models.User{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		log.Println(err.Error())
	}
	sid, _ := uuid.NewV4()
	usr.Id = sid.String()
	usr.Name = usr.Name + " Ghosh"
	users[usr.Id] = usr
	mj, _ := json.Marshal(usr)
	_, _ = fmt.Fprintf(w, "%s\n", mj)
}
