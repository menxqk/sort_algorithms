package sorts

func BubbleSort(arr []int, size int) {
	var i, j int
	var temp int

	for i = 0; i < size-1; i++ {
		for j = 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func SelectionSort(arr []int, size int) {
	var i, j, min_idx int
	var temp int

	for i = 0; i < size-1; i++ {
		min_idx = i

		for j = i + 1; j < size; j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}

		if min_idx != i {
			temp = arr[i]
			arr[i] = arr[min_idx]
			arr[min_idx] = temp
		}
	}
}

func InsertionSort(arr []int, size int) {
	var i, key, j int

	for i = 1; i < size; i++ {
		key = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = key
	}
}

func merge(arr []int, beg int, mid int, end int) {
	var i, j, k int
	var n1 int = mid - beg + 1
	var n2 int = end - mid

	leftArray := make([]int, n1)
	rightArray := make([]int, n2)

	for i = 0; i < n1; i++ {
		leftArray[i] = arr[beg+i]
	}

	for j = 0; j < n2; j++ {
		rightArray[j] = arr[mid+1+j]
	}

	i = 0
	j = 0
	k = beg

	for i < n1 && j < n2 {
		if leftArray[i] <= rightArray[j] {
			arr[k] = leftArray[i]
			i++
		} else {
			arr[k] = rightArray[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = leftArray[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = rightArray[j]
		j++
		k++
	}
}

func MergeSort(arr []int, beg int, end int) {
	if beg < end {
		var mid int = (beg + end) / 2
		MergeSort(arr, beg, mid)
		MergeSort(arr, mid+1, end)
		merge(arr, beg, mid, end)
	}
}

func heapify(arr []int, n int, i int) {
	var largest int = i
	var left int = 2*i + 1
	var right int = 2*i + 2
	var temp int

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		temp = arr[i]
		arr[i] = arr[largest]
		arr[largest] = temp

		heapify(arr, n, largest)
	}
}

func HeapSort(arr []int, size int) {
	var i int
	var temp int

	for i = size/2 - 1; i >= 0; i-- {
		heapify(arr, size, i)
	}

	for i = size - 1; i >= 0; i-- {
		temp = arr[0]
		arr[0] = arr[i]
		arr[i] = temp

		heapify(arr, i, 0)
	}
}

func partition(arr []int, beg int, end int) int {
	var pivot int = arr[end]
	var i int = beg - 1
	var j int
	var temp int

	for j = beg; j <= end-1; j++ {
		if arr[j] < pivot {
			i++
			temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	temp = arr[i+1]
	arr[i+1] = arr[end]
	arr[end] = temp
	return i + 1
}

func QuickSort(arr []int, beg int, end int) {
	if beg < end {
		var p int = partition(arr, beg, end)
		QuickSort(arr, beg, p-1)
		QuickSort(arr, p+1, end)
	}
}
