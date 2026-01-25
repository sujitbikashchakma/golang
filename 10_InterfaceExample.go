package main
import	"fmt"

type Vehicle interface{
	Start() string
	Stop() string
}

type Car struct{}
func (c Car) Start()string{
	return "Car started"
}

func (c Car) Stop() string{
	return "Car stopped"
}
func main(){
	var v Vehicle = Car{}
	fmt.Println(v.Start())
	fmt.Println(v.Stop())

}