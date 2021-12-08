const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const points = {};

function increasePoint(x, y) {
    const point = points[`${x},${y}`];

    if (point) {
        points[`${x},${y}`] = point + 1;
    } else {
        points[`${x},${y}`] = 1;
    }
}

inputStream.on('line', (line) => {
    const [_, x1s, y1s, x2s, y2s] = /(\d+),(\d+)\s->\s(\d+),(\d+)/.exec(line);
    const x1 = parseInt(x1s);
    const y1 = parseInt(y1s);
    const x2 = parseInt(x2s);
    const y2 = parseInt(y2s);

    let xPath = [];
    let yPath = [];

    if (x1 === x2) {
        xPath.push(Number(x1));
    }

    if (x1 < x2) {
        for (let currentX = x1; currentX <= x2; currentX++) {
            xPath.push(currentX);
        }
    } else if (x1 > x2) {
        for (let currentX = x1; currentX >= x2; currentX--) {
            xPath.push(currentX);
        }
    }

    if (y1 === y2) {
        yPath.push(Number(y1));
    }

    if (y1 < y2) {
        for (let currentY = y1; currentY <= y2; currentY++) {
            yPath.push(currentY);
        }
    } else if (y1 > y2) {
        for (let currentY = y1; currentY >= y2; currentY--) {
            yPath.push(currentY);
        }
    }

    // assuming only 45 degree diagonals
    for (let index = 0; index < Math.max(xPath.length, yPath.length); index++) {
        if (xPath.length === 1) {
            increasePoint(xPath[0], yPath[index]);
        } else if (yPath.length === 1) {
            increasePoint(xPath[index], yPath[0]);
        } else {
            increasePoint(xPath[index], yPath[index]);
        }
    }
});

function drawBoard() {
    for (let yIndex = 0; yIndex <= 9; yIndex++) {
        const currentLine = [];
        for (let xIndex = 0; xIndex <= 9; xIndex++) {
            if (points[`${xIndex},${yIndex}`]) {
                currentLine.push(points[`${xIndex},${yIndex}`]);
            } else {
                currentLine.push('.');
            }
        }

        console.log(currentLine.join(''));
    }
}

inputStream.on('close', () => {
    drawBoard();
    console.log(Object.values(points).filter((it) => it >= 2).length);
});
