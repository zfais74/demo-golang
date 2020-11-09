package models

//user service
//package main is the entry point

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users []*User
	//slice holding pointer to users
	//pointers can let you manipulate users across the project
	nextID = 1
)

//var blocks!
