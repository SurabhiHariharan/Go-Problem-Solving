package main


import(
	"fmt"
)

var stack []int
func main(){

	push(10,12,30,40)
	pop()
	fmt.Println(stack)
}

func push(elements ...int){

	for _,i:= range elements{
		stack = append(stack,i)
	}
	fmt.Println("Items pushed")

}

func pop(){

	n:= len(stack)-1
	stack = stack[:n]

	fmt.Println("Item popped")

}