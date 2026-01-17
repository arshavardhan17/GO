package main

import (
	"fmt"
)


func main(){
	// // slice is a Dynamic arrays like array list
	// var arr []int            //declared not initilized
	// arr =[]int{1,2,3,4,5}
	// fmt.Println(arr)

	// b:=[]int{1,2,3,4,5,6}    // declare and initilize
	// b = append(b, 5)
	// fmt.Println(b)

	// //copy function
	// nums:=[]int{1,2,3,4,5}
	// var nums2 =make([]int, len(nums))

	// fmt.Println("nums",nums)
	// fmt.Println("nums2 before copy",nums2)
	// copy(nums2,nums)
	// fmt.Println("nums2 After copy",nums2)

	// Slice operator

		var numbers =[]int{1,2,3}
		fmt.Println(numbers[0:2])
}
