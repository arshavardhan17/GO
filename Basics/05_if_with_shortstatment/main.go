package main

import (
	"fmt"
)

func main(){
	items:=40
	price:=20

	if total := items*price;total>100{
		fmt.Println(total)
	}
}
