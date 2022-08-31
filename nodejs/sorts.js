function bubbleSort(arr, size) {
    let i, j;
    let temp;

    for (i = 0; i < size - 1; i++) {
        for (j = 0; j < size - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
}

function selectionSort(arr, size) {
    let i, j, min_idx;
    let temp;

    for (i = 0; i < size - 1; i++) {
        min_idx = i;

        for (j = i + 1; j < size; j++) {
            if (arr[j] < arr[min_idx]) {
                min_idx = j;
            }
        }

        if (min_idx != i) {
            temp = arr[i];
            arr[i] = arr[min_idx];
            arr[min_idx] = temp;
        }
    }
}

function insertionSort(arr, size) {
    let i, key, j;

    for (i = 1; i < size; i++) {
        key = arr[i];
        j = i - 1;

        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j = j - 1;
        }

        arr[j + 1] = key;
    }
}

function merge(arr, beg, mid, end) {
    let i, j, k;
    let n1 = Math.floor(mid - beg) + 1;
    let n2 = Math.floor(end - mid);

    let leftArray = new Array(n1);
    let rightArray = new Array(n2);

    for (i = 0; i < n1; i++) {
        leftArray[i] = arr[beg + i];
    }

    for (j = 0; j < n2; j++) {
        rightArray[j] = arr[mid + 1 + j];
    }

    i = 0;
    j = 0;
    k = beg;

    while (i < n1 && j < n2) {
        if (leftArray[i] <= rightArray[j]) {
            arr[k] = leftArray[i];
            i++;
        } else {
            arr[k] = rightArray[j];
            j++;
        }
        k++;
    }

    while (i < n1) {
        arr[k] = leftArray[i];
        i++;
        k++;
    }

    while (j < n2) {
        arr[k] = rightArray[j];
        j++;
        k++;
    }
}

function mergeSort(arr, beg, end) {
    if (beg < end) {
        let mid = Math.floor(Math.floor(beg + end) / 2);
        mergeSort(arr, beg, mid);
        mergeSort(arr, mid + 1, end);
        merge(arr, beg, mid, end);
    }
}

function heapify(arr, n, i) {
    let largest = i;
    let left = 2 * i + 1;
    let right = 2 * i + 2;
    let temp;

    if (left < n && arr[left] > arr[largest]) {
        largest = left;
    }

    if (right < n && arr[right] > arr[largest]) {
        largest = right;
    }

    if (largest != i) {
        temp = arr[i];
        arr[i] = arr[largest];
        arr[largest] = temp;

        heapify(arr, n, largest);
    }
}

function heapSort(arr, size) {
    let i;
    let temp;

    for (i = size / 2 - 1; i >= 0; i--) {
        heapify(arr, size, i);
    }

    for (i = size - 1; i >= 0; i--) {
        temp = arr[0];
        arr[0] = arr[i];
        arr[i] = temp;

        heapify(arr, i, 0);
    }
}

function partition(arr, beg, end) {
    let pivot = arr[end];
    let i = beg - 1;
    let j;
    let temp;

    for (j = beg; j <= end - 1; j++) {
        if (arr[j] < pivot) {
            i++;
            temp = arr[i];
            arr[i] = arr[j];
            arr[j] = temp;
        }
    }

    temp = arr[i + 1];
    arr[i + 1] = arr[end];
    arr[end] = temp;
    return i + 1;
}

function quickSort(arr, beg, end) {
    if (beg < end) {
        let p = partition(arr, beg, end);
        quickSort(arr, beg, p - 1);
        quickSort(arr, p + 1, end);
    }
}

module.exports = {
    bubbleSort: bubbleSort,
    selectionSort: selectionSort,
    insertionSort: insertionSort,
    mergeSort: mergeSort,
    heapSort: heapSort,
    quickSort: quickSort,

};