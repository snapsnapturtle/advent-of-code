const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const positionCounts = [];

inputStream.on('line', (line) => {
    line.split('').forEach((bit, index) => {
        if (!positionCounts[index]) {
            positionCounts[index] = [];
        }

        if (!positionCounts[index][Number(bit)]) {
            positionCounts[index][Number(bit)] = 0;
        }

        positionCounts[index][Number(bit)]++;
    });
});

function findComplement(number) {
    const mask = Math.pow(2, number.toString(2).length) - 1;
    return number ^ mask;
}

inputStream.on('close', () => {
    const gammaRateBinary = positionCounts.reduce((acc, position) => {
        if (position[0] > position[1]) {
            return acc + 0;
        } else {
            return acc + 1;
        }
    }, '');
    const gammaRate = parseInt(gammaRateBinary, 2);
    const epsilonRate = findComplement(parseInt(gammaRateBinary, 2));
    console.log({
        gammaRate: gammaRate.toString(2),
        epsilonRate: epsilonRate.toString(2),
    });

    console.log(gammaRate * epsilonRate);
});
