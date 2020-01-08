package main


import(
	"fmt"
)

func main(){


	var array [5][5]int
	var odd []int
	var even []int
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			fmt.Scan(&array[i][j])
			if array[i][j] % 2 != 0 {
				odd = append(odd,array[i][j])
			}else{
				even = append(even,array[i][j])
			}
		}
	}

	for i:=0; i<len(odd); i++ {
		fmt.Printf("The odd elements are %d\n",odd[i])
	}

	for j:=0; j<len(even); j++ {
		fmt.Printf("The even elements are %d\n",even[j])
	}

}