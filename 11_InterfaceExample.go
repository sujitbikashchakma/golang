package main
import	"fmt"

type Man interface{
		Eat() string
		Sleep() string
		Think() string
}

type Sujit struct{}
func (e Sujit) Eat()string{
	return "Sujit is Eating"
} 
func (s Sujit) Sleep() string{
	return "Sujit is Sleeping"
}

func (t Sujit) Think() string{
	return "Sharif is Thinking"
}

func main(){
	var a Man = Sujit{}
	fmt.Println(a.Eat())
	fmt.Println(a.Sleep())
	fmt.Println(a.Think())

}