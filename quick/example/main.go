package main

import (
	"fmt"
	"golang-sort-algorithm/quick"
)

func main() {
	data := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quick.Sort(data, 0, len(data)-1)
	fmt.Println(fmt.Sprintf("%v", data))
}
