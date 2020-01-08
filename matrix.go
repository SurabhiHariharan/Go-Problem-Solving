package main


import(
	"fmt"
)


func main(){


	var array [5][5]int
	
	for i:=0; i<5; i++ {
		for j:=0; j<5; j++{
			if i == j {
				array[i][j] = 0
			}else if i < j {
				array[i][j] = -1
			}else {
				array[i][j] = 1
			}
		}
	}
	for i:=0; i<5; i++ {
		fmt.Println("\n")
		for j:=0;j<5; j++ {
			fmt.Print(array[i][j])
		}
	}
}