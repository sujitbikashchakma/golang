package main
import "fmt"

//Animal Interface

type Animal interface{
		Speak() string
}

// Dog struct and method
type Dog struct{}

func (d Dog) Speak()string{
	return "Barking"
}

// Cat struct and method
type Cat struct{}

func (c Cat) Speak()string{
	return "Meoww"
}

func main(){
	var a Animal

	a= Dog{}
	fmt.Println(a.Speak())

	a = Cat{}
	fmt.Println(a.Speak())

}