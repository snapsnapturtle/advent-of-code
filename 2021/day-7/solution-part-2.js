const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const currentPositions = [];

inputStream.on('line', (line) => {
    line.split(',').forEach((position) => {
        currentPositions.push(Number(position));
    });
});

function calculateFuel(proposedPosition) {
    return currentPositions.reduce((acc, pos) => {
        const distance = Math.abs(proposedPosition - pos);
        acc += (distance * (distance + 1)) / 2;

        return acc;
    }, 0);
}

inputStream.on('close', () => {
    let lowestFuel;
    let position;

    for (let index = Math.min(...currentPositions); index <= Math.max(...currentPositions); index++) {
        const fuel = calculateFuel(index);

        if (fuel < lowestFuel || lowestFuel === undefined) {
            lowestFuel = fuel;
            position = index;
        } else if (fuel === lowestFuel) {
            console.log('current lowest ', position, lowestFuel);
            console.log('found secondary match at ', index, fuel);
        }
    }

    console.log({ lowestFuel, position });
});
