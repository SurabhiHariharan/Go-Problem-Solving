package main


import(
	"fmt"
	"sort"
	"strings"
)

func main(){

	var arraylength int
	
	fmt.Scanf("%d",&arraylength)

	slice := make([]string,arraylength)

	for i:=0; i<arraylength; i++{
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice)
	getBiggest(slice)
}

func getBiggest(slice []string){

	sort.Strings(slice)
	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	fmt.Println(strings.Join(slice,""))

}