package main

import (
	"fmt"
)

// User type
type User struct {
	ID              int
	Name, Localtion string
}

// Player type with simple composition(embedding)
// the User type
type Player struct {
	User
	GameID int
}

func main() {
	p := Player{}
	p.ID = 1
	p.Name = "Player1"
	p.Localtion = "LA"
	p.GameID = 1234

	fmt.Printf("%s is from %s", p.Name, p.Localtion)
}
