package main

import (
	"fmt"
)

type i1 int
type i2 int

func (i i1) String() string {
	return fmt.Sprintf("the value is %d", i)
}
func main() {
	var i i1 = 1
	fmt.Println(i)
	fmt.Printf("%d", i)
}
