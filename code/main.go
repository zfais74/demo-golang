package main

import (
	"fmt"

	"github.com/zfais74/demo-golang/models"
)

//availbale to the whole package
//have constant blocks!

func main() {
	u := models.User{
		ID:        2,
		FirstName: "doug",
		LastName:  "dougstein",
	}
	fmt.Println(u)
}
