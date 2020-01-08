package main


import(
	"fmt"
)

func main(){

	staircase(5)
}

func staircase(n int32) {

	var i,j,k int32
	for k=1; k<n+1; k++{
		for i=1; i<=n-i; i++{
			fmt.Print("\t")
		}
	
		for j=k; j>0;j--{
        	fmt.Print("*")
		}
		
    
       
	
	fmt.Print("\n")
}
	


}
