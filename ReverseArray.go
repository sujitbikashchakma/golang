package main
import "fmt"

func reverseByKth(arr [] int, k int){
	n:=len(arr)-1
	if n==0{
		return
	}

	k=k%n

	reverse(arr, 0, n)
	reverse(arr, 0, k-1)
	reverse(arr, k, n)
}

func reverse(arr[] int, start int, end int){
	for start<end{
	arr[start], arr[end]= arr[end], arr[start]
	start++
	end--
	}

}

func main(){
	arr:=[]int{1,2,3,4,5,6,7,8}

	reverseByKth(arr, 4)
	fmt.Println(arr)
}