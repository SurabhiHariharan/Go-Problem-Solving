package main

import(
    "fmt"
)

//declare parking space
var lots = []int{0,0,0,0,0}

func main(){

	var status string
	status=park(4220)
	status = park(1234)
	status = park(12345)
	status = park(1111)
	status = park(2222)
	status = unpark(2222)
	status = park(3333)
	status = park(4444)

	//status= unpark(4220)
	
	fmt.Println(status)
	printLot()
}


func park(element int)(string){
	fullcounter:=0
	if len(lots) == 0 {
		lots = append(lots,element)
	}else{
	for i:=0; i<5; i++ {
		if lots[i] == 0{
			lots[i]= element
			//print()	
			return "Car Parked"
		}
		fullcounter+=1
	
	}
	}
	if fullcounter == 5{
		//print()
		return "Parking Lot Full"
	}
	//print()
	return "Car Parked"
}

func unpark(element int)(string){
	for i:=0; i<5; i++ {
		if lots[i] == element {
			lots[i] = 0
			//print()
			return "Car Unparked"
		}
	}
	//print()
	return ""
}

func print(){
	for i:=0; i<5; i++ {
		fmt.Printf("Lot no %d and Car no %d \n",i,lots[i])
	}
}


func printLot(){

	fmt.Println("Parking Lot")


	for i:=0; i<5; i++ {
		fmt.Printf("|  %d  |",lots[i])
	}
}