package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(array[i])
	}
}
