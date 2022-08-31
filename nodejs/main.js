let fs = require('fs');
let os = require('os');
let arr = require('./arr.js');
let sorts = require('./sorts.js');

const MAX = 100000;

let nums = new Array(MAX);
let numsBubble = new Array(MAX);
let numsSelection = new Array(MAX);
let numsInsertion = new Array(MAX);
let numsMerge = new Array(MAX);
let numsHeap = new Array(MAX);
let numsQuick = new Array(MAX);

let arraySize = 0;

if (process.argv.length < 3) {
    console.log("input file missing");
    process.exit(1);
}

let inputFile = process.argv.slice(2)[0];

fs.readFile(inputFile, 'utf8', function (err, data) {
    if (err) {
        console.log("error opening file:", inputFile);
        process.exit(1);
    }

    data = String(data.split(os.EOL)).split(',');

    while (arraySize < MAX && arraySize < data.length) {
        nums[arraySize] = parseInt(data[arraySize]);
        arraySize++;
    }

    copyArrays();
    sortArrays();
});



function copyArrays() {
    arr.copyArray(nums, numsBubble, arraySize);
    arr.copyArray(nums, numsSelection, arraySize);
    arr.copyArray(nums, numsInsertion, arraySize);
    arr.copyArray(nums, numsMerge, arraySize);
    arr.copyArray(nums, numsHeap, arraySize);
    arr.copyArray(nums, numsQuick, arraySize);
}

function sortArrays() {
    let start = performance.now();
    sorts.bubbleSort(numsBubble, arraySize);
    let end = performance.now();
    let diff = end - start;
    console.log("bubble sort:      %fs", (diff / 1000).toFixed(6));

    start = performance.now();
    sorts.selectionSort(numsSelection, arraySize);
    end = performance.now();
    diff = end - start;
    console.log("selection sort:   %fs", (diff / 1000).toFixed(6));

    start = performance.now();
    sorts.insertionSort(numsInsertion, arraySize);
    end = performance.now();
    diff = end - start;
    console.log("insertion sort:   %fs", (diff / 1000).toFixed(6));

    start = performance.now();
    sorts.mergeSort(numsMerge, 0, arraySize - 1);
    end = performance.now();
    diff = end - start;
    console.log("merge sort:       %fs", (diff / 1000).toFixed(6));

    start = performance.now();
    sorts.heapSort(numsHeap, arraySize);
    end = performance.now();
    diff = end - start;
    console.log("heap sort:        %fs", (diff / 1000).toFixed(6));

    start = performance.now();
    sorts.quickSort(numsQuick, 0, arraySize - 1);
    end = performance.now();
    diff = end - start;
    console.log("quick sort:       %fs", (diff / 1000).toFixed(6));
}
