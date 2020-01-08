package main


import(
	"fmt"
	"strings"
)


//Checks equal number of x and o in the input

func main(){
	var array string
	fmt.Scan(&array)
	status := find(array)
	fmt.Println("There is a match in 0's and X's",status)
}

func find(array string)bool{
	countx:=0
	counto:=0
		
	countx=strings.Count(strings.ToLower(array),"x")
	counto=strings.Count(strings.ToLower(array),"o")
		
	if counto == countx {
		return true
	}
	return false
}
