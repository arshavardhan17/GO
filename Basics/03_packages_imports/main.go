package main

// imports brings external packages into the file that u r working
// where actually u need it

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// packageName.Functionname -> call a function from a package
	name:="ArShA"
	fmt.Println("sqrt(25)", math.Sqrt(25))
	 fmt.Println(strings.ToLower(name))
}
