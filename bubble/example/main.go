package main

import "fmt"
import "golang-sort-algorithm/bubble"

func main() {
	data := []int{9, 8, 7, 6, 5, 4}
	sorted := bubble.Sort(data)
	fmt.Println(fmt.Sprintf("%v", sorted))
}
