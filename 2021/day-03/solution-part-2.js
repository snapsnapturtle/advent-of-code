const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const allNumbers = [];
let binaryStringLength = 0;

inputStream.on('line', (line) => {
    allNumbers.push(parseInt(line, 2));
    binaryStringLength = Math.max(binaryStringLength, line.length);
});

function filterNumbers(numbers, index, chooseNumbers) {
    if (numbers.length === 1) {
        return numbers[0];
    }

    // create a mask to get the "index"th bit
    const mask = 1 << index;

    const bitAtPosition = [[], []];

    numbers.forEach((n) => {
        bitAtPosition[Number((n & mask) != 0)].push(n);
    });

    return filterNumbers(chooseNumbers(bitAtPosition[0], bitAtPosition[1]), index - 1, chooseNumbers);
}

inputStream.on('close', () => {
    const oxygenRating = filterNumbers(allNumbers, binaryStringLength - 1, (zeros, ones) => {
        if (zeros.length === ones.length) {
            return ones;
        } else if (zeros.length < ones.length) {
            return ones;
        } else {
            return zeros;
        }
    });

    const scrubberRating = filterNumbers(allNumbers, binaryStringLength - 1, (zeros, ones) => {
        if (zeros.length === ones.length) {
            return zeros;
        } else if (zeros.length > ones.length) {
            return ones;
        } else {
            return zeros;
        }
    });

    console.log(oxygenRating * scrubberRating);
});
