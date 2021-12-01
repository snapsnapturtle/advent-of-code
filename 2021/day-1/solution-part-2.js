const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

let increases = 0;
let previousSum;
let window = [];

inputStream.on('line', (line) => {
    if (window.length < 2) {
        window.push(Number(line));

        return;
    }

    window.push(Number(line));

    const currentSum = window[0] + window[1] + window[2];

    if (previousSum && previousSum < currentSum) {
        increases++;
    }

    previousSum = currentSum;
    window.shift();
});

inputStream.on('close', () => {
    console.log(increases);
});
