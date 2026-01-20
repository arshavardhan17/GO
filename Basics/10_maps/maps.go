package main

import "fmt"

//maps -> hashmap,objects,dicteronires etc
func main(){
// crearing a map
m:=make(map[string]string)

//setting elements
m["name"]="golangs"
m["area"]="backend"
//geting an element
fmt.Println(m["name"],m["area"])
fmt.Println(m)
fmt.Println(len(m))

// deleting
delete(m,"area")
fmt.Println(m)

// clear map
// clear(m)
// fmt.Println(m)
}
