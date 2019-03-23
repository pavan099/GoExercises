package main

import "fmt"

// User depicts the base struct
type User struct {
	ID             int
	Name, Location string
}

// Greetings greet a user
func (u *User) Greetings() {
	fmt.Println("Hello ", u.Name)
}

// Player uses composition by reference
// uses User as struct pointer
type Player struct {
	*User
	GameID int
}

func main() {
	p := &Player{
		User:   &User{ID: 1, Name: "Player1", Location: "India"},
		GameID: 1234,
	}

	p.Greetings()
}
