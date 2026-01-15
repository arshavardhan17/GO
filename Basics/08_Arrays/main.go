package main

import "fmt"

func main(){
	//Size of array is fixed it can not grow
	var arr[5]int  //--->one method
	arr[0]=10
	arr[1]=20
	arr[3]=30

	arr[3]=40
	fmt.Println(arr)

	//Array literal
	marks :=[3]int{2,4,6}  // ---> Second method
	fmt.Println(len(marks))
}
// output = [10 20 0 30 0]
