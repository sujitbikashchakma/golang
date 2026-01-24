package main

import "fmt"

type User struct {
	Name       string
	Email      string
	Age        int
	isVerified bool
}

func main() {
	user1 := User{
		Name:       "Guru Sharif",
		Email:      "gurusharif@gmail.com",
		Age:        24,
		isVerified: true,
	}
	fmt.Println(user1)
	fmt.Println()	
	fmt.Println(user1.Name)
	fmt.Println(user1.Email)
	fmt.Println(user1.Age)
	fmt.Println(user1.isVerified)
}
