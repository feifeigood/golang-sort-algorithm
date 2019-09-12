package main

import (
	"bufio"
	"golang-sort-algorithm/diskmerge"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func main() {
	// generateFiles
	files := generateFiles()
	// sort file
	diskmerge.DiskMergeSort(files)
	// remove files and move outfile to tmp
	for _, file := range files {
		os.Remove(file)
	}
	os.Rename("output.txt", "/tmp/output.txt")
}

func generateFiles() []string {
	files := []string{}

	for i := 0; i < 10; i++ {
		filename := "input-" + strconv.Itoa(i) + ".txt"
		files = append(files, filename)
		// sort

		numbers := []int{}
		for k := 0; k < 100; k++ {
			numbers = append(numbers, rand.Intn(1000))
		}

		sort.Ints(numbers)

		// write random numbers to input files
		file, _ := os.Create(filename)
		wr := bufio.NewWriter(file)
		for _, number := range numbers {
			wr.WriteString(strconv.Itoa(number) + "\n")
			wr.Flush()
		}
		file.Close()
	}

	return files
}
