package main



import(
    "fmt"
)

var stack []int

func main(){

	
	
	push(5)
	push(4)
	push(3)
	push(2)
	pop()
	pop()
	print()
}

func push(data int){
	stack = append(stack,data)
}

func print(){
	for i:=0; i<len(stack); i++ {
		fmt.Println(stack[i])
	}
}

func pop(){
	n := len(stack) - 1 
	stack = stack[:n]
}