import math

def bubbleSort(arr, size):
    for i in range(0, size - 1):
        for j in range(0, size - i - 1):
            if arr[j] > arr[j + 1]:
                temp = arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = temp


def selectionSort(arr, size):
    for i in range(0, size - 1):
        min_idx = i
        
        for j in range(i + 1, size):
            if arr[j] < arr[min_idx]:
                min_idx = j
        
        if min_idx != i:
            temp = arr[i]
            arr[i] = arr[min_idx]
            arr[min_idx] = temp


def insertionSort(arr, size):
    for i in range(1, size):
        key = arr[i]
        j = i - 1

        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]
            j = j - 1

        arr[j + 1] = key


def _merge(arr, beg, mid, end):
    n1 = math.floor(mid - beg) + 1
    n2 = math.floor(end - mid)

    leftArray = [0]*n1
    rightArray = [0]*n2

    for i in range(0, n1):
        leftArray[i] = arr[beg + i]

    for j in range(0, n2):
        rightArray[j] = arr[mid + 1 + j]

    i = 0
    j = 0
    k = beg

    while i < n1 and j < n2:
        if leftArray[i] <= rightArray[j]:
            arr[k] = leftArray[i]
            i += 1
        else:
            arr[k] = rightArray[j]
            j += 1
        k += 1

    while i < n1:
        arr[k] = leftArray[i]
        i += 1
        k += 1

    while j < n2:
        arr[k] = rightArray[j]
        j += 1
        k += 1


def mergeSort(arr, beg, end):
    if beg < end:
        mid = math.floor(math.floor(beg + end) / 2)
        mergeSort(arr, beg, mid)
        mergeSort(arr, mid + 1, end)
        _merge(arr, beg, mid, end)


def _heapify(arr, n, i):
    largest = i
    left = 2 * i + 1
    right = 2 * i + 2

    if left < n and arr[left] > arr[largest]:
        largest = left

    if right < n and arr[right] > arr[largest]:
        largest = right

    if largest != i:
        temp = arr[i]
        arr[i] = arr[largest]
        arr[largest] = temp

        _heapify(arr, n, largest)


def heapSort(arr, size):
    for i in range(math.floor(size/2)-1, -1, -1):
        _heapify(arr, size, i)

    for i in range(size-1, -1, -1):
        temp = arr[0]
        arr[0] = arr[i]
        arr[i] = temp

        _heapify(arr, i, 0)


def _partition(arr, beg, end) -> int:
    pivot = arr[end]
    i = beg - 1

    for j in range(beg, end):
        if arr[j] < pivot:
            i += 1
            temp = arr[i]
            arr[i] = arr[j]
            arr[j] = temp

    temp = arr[i + 1]
    arr[i + 1] = arr[end]
    arr[end] = temp
    return i + 1


def quickSort(arr, beg, end):
    if beg < end:
        p = _partition(arr, beg, end)
        quickSort(arr, beg, p - 1)
        quickSort(arr, p + 1, end)
