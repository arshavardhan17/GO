package main

import (
	"fmt"
)

func main(){
	//While using for
	i:=1
	for i<=3{
		fmt.Println(i)
		i++
	}

	// Sum of numbers till n
	num:=10
	sum:=0
	for i:=0;i<num;i++{
		sum=sum+i
	}
	 fmt.Println(sum)

	// for each type loop using range
	
	 // var a =[5]int{1,2,3,4,5}
	 arr := []int{10, 20, 30}
for index, value := range arr {
    fmt.Println(index, value)
}


}
