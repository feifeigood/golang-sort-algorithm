package main

import "fmt"
import "golang-sort-algorithm/shell"

func main() {
	data := []int{1, 4, 2, 7, 9, 8, 3, 6}
	sorted := shell.Sort(data)
	fmt.Println(fmt.Sprintf("%v", sorted))
}
