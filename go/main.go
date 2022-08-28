package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX int = 100000
const NUMBERS_FILE string = "../numbers.txt"

var nums [MAX]int

func main() {
	f, err := os.Open(NUMBERS_FILE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	count := 0
	for i := 0; i < len(nums); i++ {
		num, err := reader.ReadString(',')
		if err != nil {
			n, err := strconv.Atoi(num)
			if err == nil {
				nums[i] = n
				count++
			}
			break
		}
		n, err := strconv.Atoi(num[:len(num)-1])
		if err == nil {
			nums[i] = n
			count++
		}
	}

	fmt.Println(nums)
	fmt.Printf("len(nums): %d\n", count)
}
