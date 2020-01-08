package main


import(
	"fmt"
)


func main(){

	var array =[]int{5,0,15,7,3,9,1,2,6,8}
	numberOne,numberTwo:= determinePresence(array,5)
	if ((numberOne !=0 ) && (numberTwo != 0)) {
	fmt.Println("the two numbers are",numberOne,numberTwo)
	}else{
		fmt.Println("No combinations exists")
	}
}


func determinePresence(array []int, element int)(int,int){
	for i:=0; i<len(array); i++{
		for j:= i+1; j<len(array)-1; j++ {
			if array[i]+array[j] == element {
				return array[i],array[j]
			}
		}
	}
	return 0,0
}