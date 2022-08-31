import sys
import csv
from arr import *
from sorts import *
import time

MAX = 100000

nums = [0]*MAX
numsBubble = [0]*MAX
numsSelection = [0]*MAX
numsInsertion = [0]*MAX
numsMerge = [0]*MAX
numsHeap = [0]*MAX
numsQuick = [0]*MAX

arraySize = 0

if len(sys.argv) < 2:
    print("input file missing")
    exit(1)

inputFile = sys.argv[1]

try:
    with open(inputFile, mode='r') as f:
        reader = csv.reader(f, delimiter=',')
        for row in reader:
            i = 0
            while (arraySize < MAX and i < len(row)):
                nums[arraySize] = int(row[i])
                arraySize += 1
                i += 1        
    
except FileNotFoundError:
    print("FileNotFoundError")
except IOError:
    print(IOError)

copyArray(nums, numsBubble, arraySize)
copyArray(nums, numsSelection, arraySize)
copyArray(nums, numsInsertion, arraySize)
copyArray(nums, numsMerge, arraySize)
copyArray(nums, numsHeap, arraySize)
copyArray(nums, numsQuick, arraySize)

# bubble sort
start = time.process_time()
bubbleSort(numsBubble, arraySize)
elapsed_time = time.process_time()
print("bubble sort:    %fs" % elapsed_time)

# selection sort
start = time.process_time()
selectionSort(numsSelection, arraySize)
elapsed_time = time.process_time()
print("selection sort: %fs" % elapsed_time)

# insertion sort
start = time.process_time()
insertionSort(numsInsertion, arraySize)
elapsed_time = time.process_time() - start
print("insertion sort: %fs" % elapsed_time)

# merge sort
start = time.process_time()
mergeSort(numsMerge, 0, arraySize - 1)
elapsed_time = time.process_time() - start
print("merge sort:     %fs" % elapsed_time)

# heap sort
start = time.process_time()
heapSort(numsHeap, arraySize)
elapsed_time = time.process_time() - start
print("heap sort:      %fs" % elapsed_time)

# quick sort
start = time.process_time()
quickSort(numsQuick, 0, arraySize - 1)
elapsed_time = time.process_time() - start
print("quick sort:     %fs" % elapsed_time)
