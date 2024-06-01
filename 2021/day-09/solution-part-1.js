const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const heightMap = [];

inputStream.on('line', (line) => {
    heightMap.push(line.split('').map((it) => Number(it)));
});

function getAbove(x, y) {
    return heightMap[x][y - 1] !== undefined ? heightMap[x][y - 1] : Number.POSITIVE_INFINITY;
}

function getLeft(x, y) {
    if (heightMap[x - 1] !== undefined && heightMap[x - 1][y] !== undefined) {
        return heightMap[x - 1][y];
    } else {
        return Number.POSITIVE_INFINITY;
    }
}
function getRight(x, y) {
    if (heightMap[x + 1] !== undefined && heightMap[x + 1][y] !== undefined) {
        return heightMap[x + 1][y];
    } else {
        return Number.POSITIVE_INFINITY;
    }
}
function getBelow(x, y) {
    return heightMap[x][y + 1] !== undefined ? heightMap[x][y + 1] : Number.POSITIVE_INFINITY;
}

function isLowPoint(x, y) {
    const height = heightMap[x][y];
    const adjacentPositions = [getAbove(x, y), getLeft(x, y), getRight(x, y), getBelow(x, y)];

    return Math.min(...adjacentPositions) > height;
}

inputStream.on('close', () => {
    let riskLevel = 0;

    heightMap.forEach((yValues, x) => {
        yValues.forEach((_, y) => {
            if (isLowPoint(x, y)) {
                riskLevel += 1 + heightMap[x][y];
            }
        });
    });

    console.log(riskLevel);
});
