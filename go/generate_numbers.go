package main

import (
	"fmt"
	"math/rand"
	"os"
)

const MAX int = 100000

func main() {

	f, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	for i := 0; i < MAX; i++ {
		n := rand.Intn(MAX)
		s := fmt.Sprintf("%d,", n)
		if i+1 == MAX {
			s = fmt.Sprintf("%d", n)
		}
		f.WriteString(s)
	}
}
