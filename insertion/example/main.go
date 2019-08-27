package main

import "fmt"
import "golang-sort-algorithm/insertion"

func main() {
	data := []int{7, 1, 9, 2, 13, 6, 8, 5, 5, 3}
	sort, _ := insertion.Sort(data)
	fmt.Println(fmt.Sprintf("%v", sort))
}
