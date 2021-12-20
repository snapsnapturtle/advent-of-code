const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

class SingleNumber {
    numericalValue;
    startIndex;
    length;

    constructor(startIndex, value) {
        this.numericalValue = parseInt(value);
        this.startIndex = startIndex;
        this.length = value.length;
    }
}

class Pair {
    startIndex;

    value = '';
    length = 0;

    constructor(startIndex, value) {
        this.startIndex = startIndex;
        this.value = value;
        this.length = value.length;
    }

    getRightNumericalValue() {
        const [_, numericString] = this.value.match(/\d+,(\d+)/);

        return parseInt(numericString);
    }

    getLeftNumericalValue() {
        const [_, numericString] = this.value.match(/(\d+),\d+/);

        return parseInt(numericString);
    }
}

function readPairAtIndex(snailNumber, index) {
    let snailPairString = '';

    for (let currentIndex = index; currentIndex < snailNumber.length; currentIndex++) {
        const element = snailNumber[currentIndex];
        snailPairString += element;

        if (element == ']') {
            break;
        }
    }

    return new Pair(index, snailPairString);
}

function findPreviousNumber(snailNumber, startIndex) {
    let numberString = '';
    let numberStartIndex;

    for (let index = startIndex; index >= 0; index--) {
        const element = snailNumber[index];

        if (!'[],'.includes(element)) {
            // not a control character
            numberString += element;
        } else if (numberString != '') {
            numberStartIndex = index + 1;
            break;
        }
    }

    if (numberStartIndex != undefined) {
        return new SingleNumber(numberStartIndex, numberString.split('').reverse().join(''));
    }
    return null;
}

function findNextNumber(snailNumber, startIndex) {
    let numberString = '';
    let numberStartIndex;

    for (let index = startIndex; index < snailNumber.length; index++) {
        const element = snailNumber[index];

        if (!'[],'.includes(element)) {
            // not a control character
            numberString += element;
            if (numberStartIndex == undefined) {
                numberStartIndex = index;
            }
        } else if (numberString != '') {
            break;
        }
    }

    if (numberStartIndex != undefined) {
        return new SingleNumber(numberStartIndex, numberString);
    }

    return null;
}

function getNextExplosionIndex(snailNumber) {
    let openCount = 0;

    for (let index = 0; index < snailNumber.length; index++) {
        const character = snailNumber[index];

        if (character == '[') {
            openCount++;
        } else if (character == ']') {
            openCount--;
        }

        if (openCount >= 5) {
            return index;
        }
    }

    return -1;
}

function getNextSplitIndex(snailNumber) {
    let currentNumberString = '';
    let splitNumberIndex;

    for (let index = 0; index < snailNumber.length; index++) {
        const element = snailNumber[index];

        if ('[],'.includes(element)) {
            currentNumberString = '';
        } else {
            currentNumberString += element;
        }

        if (currentNumberString.length > 1) {
            splitNumberIndex = index - 1;
            break;
        }
    }

    return splitNumberIndex;
}

function replace(snailNumber, start, length, replacement) {
    snailNumber.splice(start, length, ...replacement);
}

function sum(snailNumberA, snailNumberB) {
    return reduce(['[', ...snailNumberA, ',', ...snailNumberB, ']']);
}

function explode(snailNumber, explosionIndex) {
    const explodingPair = readPairAtIndex(snailNumber, explosionIndex);
    let newExplodingPairIndex = explodingPair.startIndex;

    const nextRightNumber = findNextNumber(snailNumber, explodingPair.startIndex + explodingPair.length);

    if (nextRightNumber) {
        const rightNumberOfPair = explodingPair.getRightNumericalValue();

        replace(
            snailNumber,
            nextRightNumber.startIndex,
            nextRightNumber.length,
            (nextRightNumber.numericalValue + rightNumberOfPair).toString().split('')
        );
    }

    const nextLeftNumber = findPreviousNumber(snailNumber, explodingPair.startIndex);

    if (nextLeftNumber) {
        const leftNumberOfPair = explodingPair.getLeftNumericalValue();
        const newLeftNumber = (nextLeftNumber.numericalValue + leftNumberOfPair).toString().split('');

        newExplodingPairIndex -= nextLeftNumber.length;
        newExplodingPairIndex += newLeftNumber.length;

        replace(snailNumber, nextLeftNumber.startIndex, nextLeftNumber.length, newLeftNumber);
    }

    replace(snailNumber, newExplodingPairIndex, explodingPair.length, ['0']);
}

function split(snailNumber, splitIndex) {
    const splitNumber = findNextNumber(snailNumber, splitIndex);
    const newLeft = Math.floor(splitNumber.numericalValue / 2);
    const newRight = Math.ceil(splitNumber.numericalValue / 2);

    const newValue = `[${newLeft},${newRight}]`;

    replace(snailNumber, splitNumber.startIndex, splitNumber.length, newValue.split(''));
}

function reduce(snailNumber) {
    let needsReduction = true;

    while (needsReduction) {
        const explosionIndex = getNextExplosionIndex(snailNumber);

        if (explosionIndex >= 0) {
            explode(snailNumber, explosionIndex);
            continue;
        }

        const splitIndex = getNextSplitIndex(snailNumber);

        if (splitIndex >= 0) {
            split(snailNumber, splitIndex);
            continue;
        }

        needsReduction = false;
    }

    return snailNumber;
}

function getPair(snailNumber) {
    let open = 0;

    for (let index = 0; index < snailNumber.length; index++) {
        const element = snailNumber[index];

        if (element == '[') {
            open++;
        }

        if (element == ']') {
            open--;
        }

        if (element == ',' && open == 1) {
            return [snailNumber.slice(1, index), snailNumber.slice(index + 1, snailNumber.length - 1)];
        }
    }
}

function calculateMagnitudeForPair(left, right) {
    let currentLeft = left;
    let currentRight = right;

    if (left[0] == '[') {
        const [newLeft, newRight] = getPair(left);

        currentLeft = calculateMagnitudeForPair(newLeft, newRight);
    }

    if (right[0] == '[') {
        const [newLeft, newRight] = getPair(right);

        currentRight = calculateMagnitudeForPair(newLeft, newRight);
    }

    return parseInt(currentLeft) * 3 + parseInt(currentRight) * 2;
}

let maxMagnitude = 0;
const availableNumbers = [];

inputStream.on('line', (line) => {
    availableNumbers.push(line.split(''));
});

inputStream.on('close', () => {
    availableNumbers.forEach((firstNumber, firstIndex) => {
        availableNumbers.forEach((secondNumber, secondIndex) => {
            if (firstIndex != secondIndex) {
                const magForward = calculateMagnitudeForPair(...getPair(sum(firstNumber, secondNumber)));
                const magBackward = calculateMagnitudeForPair(...getPair(sum(secondNumber, firstNumber)));

                maxMagnitude = Math.max(maxMagnitude, magForward, magBackward);
            }
        });
    });

    console.log(maxMagnitude);
});
