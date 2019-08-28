package main

import "fmt"
import "golang-sort-algorithm/selection"

func main() {
	data := []int{7, 1, 9, 2, 13, 6, 8, 5, 5, 3}
	sort := selection.Sort(data)
	fmt.Printf(fmt.Sprintf("%v", sort))
}
