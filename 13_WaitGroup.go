package main
import (
		"fmt"
		"sync")

func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()		//Wait untill worker is finished
	fmt.Println("Worker ", id, "Started")

}

func main(){
	var wg sync.WaitGroup

	for i:=1; i<=5; i++{
		wg.Add(1) //Initialize same wait 1
		go worker(i, &wg)
	}
	wg.Wait()
}