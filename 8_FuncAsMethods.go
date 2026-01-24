package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func main() {
	user1 := User{
		Name:  "GuruSharif",
		Email: "guru@guru.com",
		Age:   40,
	}

	fmt.Println(user1.IsAdult())
}
