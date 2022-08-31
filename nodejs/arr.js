exports.copyArray = function (src, dst, size) {
    for (let i = 0; i < size; i++) {
        dst[i] = src[i];
    }
}

exports.printArray = function (arr, size) {
    process.stdout.write("[");
    for (let i = 0; i < size; i++) {
        if (i < size - 1) {
            process.stdout.write(String(arr[i] + " "));
        } else {
            process.stdout.write(String(arr[i]));
        }
    }
    process.stdout.write("]");
}