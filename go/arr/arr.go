package arr

import "fmt"

func CopyArray(src []int, dst []int, size int) {
	for i := 0; i < size; i++ {
		dst[i] = src[i]
	}
}

func PrintArray(arr []int, size int) {
	fmt.Print("[")
	for i := 0; i < size; i++ {
		if i < size-1 {
			fmt.Printf("%d ", arr[i])
		} else {
			fmt.Printf("%d", arr[i])
		}
	}
	fmt.Print("]")
}
