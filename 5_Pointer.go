package main
import "fmt"

func main(){
	x:=1020
	pointer := &x

	fmt.Println(pointer)
	fmt.Println(*pointer)
}