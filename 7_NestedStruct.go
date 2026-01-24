package main
import "fmt"

type Address struct{
	City string
	HouseNumber string
	
}

type Customer struct{
	Name string
	Email string
	Location Address
}
func main(){

	user1:=Customer{
		Name: "GuruSharif",
		Email: "gurusharif@sharif.com",
		Location: Address{
			City: "Khulna",
			HouseNumber: "B/4",
		},
	}
	fmt.Println(user1)
	fmt.Println()
	fmt.Println("Details of Customer")
	fmt.Println()
	fmt.Println("Name:"+ user1.Name)
	fmt.Println("Email:"+user1.Email)
	fmt.Println("House Number:"+user1.Location.HouseNumber)
	fmt.Println("City:"+user1.Location.City)
}