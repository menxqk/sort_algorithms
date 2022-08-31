import sys

def copyArray(src, dst, size):
    for i in range(0, size):
        dst[i] = src[i]

def printArray(arr, size):
    sys.stdout.write("[")
    
    for i in range(0, size):
        if i < size - 1:
            sys.stdout.write(str(arr[i]) + " ")
        else:
            sys.stdout.write(str(arr[i]))

    sys.stdout.write("]")
