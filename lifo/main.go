package main

import "fmt"

func push(arr *[]int, n int) {
	*arr = append(*arr, n)
}
func pop(arr []int) []int {
	return arr[:len(arr)-1]
}
func print1() {

}
func main() {
	var arr []int
	push(&arr, 5)
	push(&arr, 3)
	push(&arr, 2)
	push(&arr, 4)
	arr = pop(arr)
	arr = pop(arr)
	arr = pop(arr)
	arr = pop(arr)

	fmt.Println(arr)

}
