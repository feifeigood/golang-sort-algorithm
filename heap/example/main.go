package main

import (
	"fmt"
	"golang-sort-algorithm/heap"
)

func main() {
	data := []int{3, 7, 1, 5, 2, 6, 4, 8}
	heap.Sort(data)
	fmt.Println(fmt.Sprintf("%v", data))
}
