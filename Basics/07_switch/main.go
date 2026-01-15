package main

import (
	"fmt"
)

func main(){
	option:=8
	switch(option){
	case 1: fmt.Println("option 1")
	case 2: fmt.Println("option 2")
	case 3: fmt.Println("option 3")
	case 4: fmt.Println("option 4")
	case 5: fmt.Println("option 5")
	default:fmt.Println("Invalid Option")
	}
}
