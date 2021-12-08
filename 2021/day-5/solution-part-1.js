const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const points = {
    'x,y': 0,
};

function isStraightLine(x1, y1, x2, y2) {
    return x1 === x2 || y1 === y2;
}

function increasePoint(x, y) {
    const point = points[`${x},${y}`];

    if (point) {
        points[`${x},${y}`] = point + 1;
    } else {
        points[`${x},${y}`] = 1;
    }
}

inputStream.on('line', (line) => {
    const [_, x1, y1, x2, y2] = /(\d+),(\d+)\s->\s(\d+),(\d+)/.exec(line);

    if (isStraightLine(x1, y1, x2, y2)) {
        if (x1 === x2) {
            // increase y positions
            for (let currentY = Math.min(y1, y2); currentY <= Math.max(y1, y2); currentY++) {
                increasePoint(x1, currentY);
            }
        }

        if (y1 === y2) {
            // increase x positions
            for (let currentX = Math.min(x1, x2); currentX <= Math.max(x1, x2); currentX++) {
                increasePoint(currentX, y1);
            }
        }
    }
});

inputStream.on('close', () => {
    console.log(Object.values(points).filter(it => it >= 2).length)
});
