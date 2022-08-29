package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"sort_algorithms/arr"
	"sort_algorithms/sorts"
)

const MAX int = 100000

var nums [MAX]int
var numsBubble [MAX]int
var numsSelection [MAX]int
var numsInsertion [MAX]int
var numsMerge [MAX]int
var numsHeap [MAX]int
var numsQuick [MAX]int

var arraySize int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("input file missing")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("error opening file: %s\n", inputFile)
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for i := 0; i < len(nums); i++ {
		num, err := reader.ReadString(',')
		if err != nil {
			n, err := strconv.Atoi(num)
			if err == nil {
				nums[i] = n
				arraySize++
			}
			break
		}
		n, err := strconv.Atoi(num[:len(num)-1])
		if err == nil {
			nums[i] = n
			arraySize++
		}
	}

	arr.CopyArray(nums[:], numsBubble[:], arraySize)
	arr.CopyArray(nums[:], numsSelection[:], arraySize)
	arr.CopyArray(nums[:], numsInsertion[:], arraySize)
	arr.CopyArray(nums[:], numsMerge[:], arraySize)
	arr.CopyArray(nums[:], numsHeap[:], arraySize)
	arr.CopyArray(nums[:], numsQuick[:], arraySize)

	var start time.Time
	var diff time.Duration

	// bubble sort
	start = time.Now()
	sorts.BubbleSort(numsBubble[:], arraySize)
	diff = time.Since(start)
	fmt.Printf("bubble sort:    %.6fs\n", diff.Seconds())

	// selection sort
	start = time.Now()
	sorts.SelectionSort(numsSelection[:], arraySize)
	diff = time.Since(start)
	fmt.Printf("selection sort: %.6fs\n", diff.Seconds())

	// insertion sort
	start = time.Now()
	sorts.InsertionSort(numsInsertion[:], arraySize)
	diff = time.Since(start)
	fmt.Printf("insertion sort: %.6fs\n", diff.Seconds())

	// merge sort
	start = time.Now()
	sorts.MergeSort(numsMerge[:], 0, arraySize-1)
	diff = time.Since(start)
	fmt.Printf("merge     sort: %.6fs\n", diff.Seconds())

	// heap sort
	start = time.Now()
	sorts.HeapSort(numsHeap[:], arraySize)
	diff = time.Since(start)
	fmt.Printf("heap      sort: %.6fs\n", diff.Seconds())

	// quick sort
	start = time.Now()
	sorts.QuickSort(numsQuick[:], 0, arraySize-1)
	diff = time.Since(start)
	fmt.Printf("quick     sort: %.6fs\n", diff.Seconds())
}
