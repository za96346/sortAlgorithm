package main

import (
	"fmt"
	"math/rand"
)

var arr = make([]int, 0)
func main() {
	for i := 0; i < 20; i++ {
		x := i
		v := rand.Intn(30)
		if x % 2 == 0 {
			x -= v
		} else {
			x += v
		}

		arr = append(arr, x)
	}
	fmt.Println("原始陣列", arr)
	// sorting.SelectSort(&arr)
	fmt.Print("演算後陣列", arr, "\n")
	
}