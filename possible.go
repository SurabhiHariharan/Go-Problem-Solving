package main

import(
	"fmt"
	"strings"
)

func main(){
	status :=findPresence("TEERPIONUMS","PERMIT")
	fmt.Println(status)
}

func findPresence(sample string, element string)(bool){
	charCounter :=0
	string1:=strings.Split(sample,"")
	string2:=strings.Split(element,"")
	for i:=0; i<len(string2); i++ {
		for j:=0; j<len(string1); j++ {
			if string1[i] == string2[j]{
				charCounter+=1
				break
			}
		}
	}
	if charCounter == len(element){
		return true
	}
	return false
}