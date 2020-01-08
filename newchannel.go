package main


import(
	"fmt"
	"sync"
	"math/rand"
)


type Random struct{
	Id         int
	RandomNo   int
	Sum        int
}


var random = make(chan Random,10)

func main(){

	var wg sync.WaitGroup

	var num int
	fmt.Printf("Enter the total number of jobs")
	fmt.Scanf("%d",&num)
	wg.Add(1)
	go getSum(random, &wg, num)
	done := make(chan bool)

	go print(done)

	<-done
	wg.Wait()
}


func getSum(random chan Random, wg *sync.WaitGroup,n int){

	var no int
	sum:=0

	for i:=0; i<n; i++{
	number := rand.Intn(999)
	no = number
	for number != 0 {
		digits := number % 10
		sum += digits
		number /= 10
	}

	ch:= Random{i,no, sum}
	random <- ch
	}

	close(random)
	wg.Done()
}


func print(done chan bool){

	for i:= range random{
		fmt.Printf("The ID is %d random number is %d and the sum of digits is %d\n",i.Id,i.RandomNo,i.Sum)
	}

	done <- true
}