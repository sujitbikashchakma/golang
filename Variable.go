package main
import "fmt"

func main()  {
	var Name string
	Name = "Sujit"
	fmt.Println(Name)
	
	result:=Hello("Sujit")
	fmt.Println(result)

	fmt.Println(MathOperation(100,10))

	
}

func Hello(First string) string{
	return "Hello " + First

}

func MathOperation(A int, B int)(int, int, int, int){
	Add:=A+B
	Sub:=A-B
	Multi:=A*B
	Div:=A/B
	return Add, Sub, Multi, Div
}