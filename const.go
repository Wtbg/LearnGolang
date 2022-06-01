package main

import "fmt"
import "unsafe"

func main(){
	const a = "asd"
	const b string = "fgh"
	c := ""
	c = a + b
	fmt.Printf("面积是%s",c)
	fmt.Println();
	fmt.Println(len(c));
	fmt.Println(unsafe.Sizeof(c))
}
