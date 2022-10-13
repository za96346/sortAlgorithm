package main

var arr = make([]int, 20)
func main() {
	for i :=0; i < 20; i++ {
		arr = append(arr, i)
	}
	insertionSort()
}