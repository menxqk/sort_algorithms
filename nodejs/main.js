fs = require('fs');
os = require('os');

const MAX = 100000;
const NUMBERS_FILE = "../numbers.txt";

let nums = new Array(MAX);

fs.readFile(NUMBERS_FILE, 'utf8', function (err, data) {
    if (err) {
        return console.log(err);
    }

    data = String(data.split(os.EOL)).split(',');

    if (!Array.isArray(data) || data.length < MAX) {
        console.log("data < MAX");
        process.exit(1);
    }

    let count = 0;
    for (let i = 0; i < nums.length; i++) {
        nums[i] = parseInt(data[i]);
        count++;
    }

    console.log(nums);
    console.log("len(nums):", count);
});
