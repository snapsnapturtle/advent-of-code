const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const heightMap = [];

inputStream.on('line', (line) => {
    heightMap.push(line.split('').map((it) => Number(it)));
});

function getAbove(row, col) {
    if (heightMap[row - 1] !== undefined && heightMap[row - 1][col] !== undefined) {
        return {
            value: heightMap[row - 1][col],
            position: [row - 1, col],
        };
    } else {
        return {
            value: undefined,
            position: undefined,
        };
    }
}

function getBelow(row, col) {
    if (heightMap[row + 1] !== undefined && heightMap[row + 1][col] !== undefined) {
        return {
            value: heightMap[row + 1][col],
            position: [row + 1, col],
        };
    } else {
        return {
            value: undefined,
            position: undefined,
        };
    }
}

function getLeft(row, col) {
    if (heightMap[row] !== undefined && heightMap[row][col - 1] !== undefined) {
        return {
            value: heightMap[row][col - 1],
            position: [row, col - 1],
        };
    } else {
        return {
            value: undefined,
            position: undefined,
        };
    }
}
function getRight(row, col) {
    if (heightMap[row] !== undefined && heightMap[row][col + 1] !== undefined) {
        return {
            value: heightMap[row][col + 1],
            position: [row, col + 1],
        };
    } else {
        return {
            value: undefined,
            position: undefined,
        };
    }
}

function isLowPoint(row, col) {
    const height = heightMap[row][col];
    const adjacentPositions = [
        getAbove(row, col).value,
        getLeft(row, col).value,
        getRight(row, col).value,
        getBelow(row, col).value,
    ].filter((it) => it !== undefined);

    return Math.min(...adjacentPositions) > height;
}

function checkDirection(newPosition, height) {
    return newPosition.value !== 9 && newPosition.value > height;
}

function calculateBasinSize(startingPosition, ignoreDirection, basinMap) {
    const [row, col] = startingPosition;
    basinMap.add(startingPosition.join(','));

    if (ignoreDirection !== 0 && checkDirection(getLeft(row, col), heightMap[row][col])) {
        calculateBasinSize(getLeft(row, col).position, 2, basinMap);
    }

    if (ignoreDirection !== 1 && checkDirection(getAbove(row, col), heightMap[row][col])) {
        calculateBasinSize(getAbove(row, col).position, 3, basinMap);
    }

    if (ignoreDirection !== 2 && checkDirection(getRight(row, col), heightMap[row][col])) {
        calculateBasinSize(getRight(row, col).position, 0, basinMap);
    }

    if (ignoreDirection !== 3 && checkDirection(getBelow(row, col), heightMap[row][col])) {
        calculateBasinSize(getBelow(row, col).position, 1, basinMap);
    }

    return basinMap;
}

// directions for ignoreDirection: right(0) top(1) left(2) bottom(3)

function printInput(basins) {
    const points = new Set(basins.map((it) => [...it]).flat());

    heightMap.forEach((rowValues, row) => {
        let line = '';
        rowValues.forEach((colValue, col) => {
            if (points.has(`${row},${col}`)) {
                line = line + '\x1b[36m' + colValue + '\x1b[0m';
            } else {
                line = line + colValue;
            }
        });

        console.log(line);
    });
}

inputStream.on('close', () => {
    const lowPoints = [];

    heightMap.forEach((rowValues, row) => {
        rowValues.forEach((_, col) => {
            if (isLowPoint(row, col)) {
                lowPoints.push([row, col]);
            }
        });
    });

    const basins = lowPoints.map((point) => {
        return calculateBasinSize(point, undefined, new Set());
    });

    const resultOfThreeLargest = basins
        .sort((a, b) => b.size - a.size)
        .slice(0, 3)
        .reduce((acc, it) => {
            return acc * it.size;
        }, 1);

    printInput(
        lowPoints.map((point) => {
            return calculateBasinSize(point, undefined, new Set());
        })
    );

    console.log(resultOfThreeLargest);
});
