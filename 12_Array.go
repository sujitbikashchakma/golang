package main
import (
	"fmt"
)


func main(){
	//Array declaration methods
 var numbers [5]int = [5]int {1,2,3,4,5}
 Roll:=[3]int{1,2,3}       //Direct initialization
 Votter:=[...]int{1,2,3}   //Dynamic declaration or Inferred length
fmt.Println(numbers)
fmt.Println(Roll)
fmt.Println(Votter)

//Slice
 slices:=numbers[1:3]
 fmt.Println(slices)
 fmt.Println("Length: ",len(slices))
 fmt.Println("Capacity: ",cap(slices))
}