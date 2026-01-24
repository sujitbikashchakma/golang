package main
import (
	"fmt"
	"errors"
)

func main(){

fmt.Println(Divide(10,0))
fmt.Println(Divide(10,2))
fmt.Println(Divide(10,3))

}

func Divide(a float64, b float64) (float64, error){
if b==0 {
	return 0, errors.New("Divided by Zero")
}
return a/b, nil
}
