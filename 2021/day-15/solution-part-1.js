const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

const riskLevelMap = [];

inputStream.on('line', (line) => {
    const riskRow = line.split('').map((it) => parseInt(it));
    riskLevelMap.push(riskRow);
});

function doesNodeExist(row, col) {
    return row >= 0 && col >= 0 && row < riskLevelMap.length && col < riskLevelMap[0].length;
}

function hasNotVisited(visitedNodes, row, col) {
    return !visitedNodes.has(row + ',' + col);
}

function processNode(pathMap, visitedNodes, startRow, startCol, targetRow, targetCol) {
    // left, top, right, bottom
    const neighbours = [
        [startRow, startCol - 1],
        [startRow - 1, startCol],
        [startRow, startCol + 1],
        [startRow + 1, startCol],
    ];

    neighbours
        .filter((it) => doesNodeExist(it[0], it[1]) && hasNotVisited(visitedNodes, it[0], it[1]))
        .forEach(([row, col]) => {
            const startRisk = pathMap.get(startRow + ',' + startCol).risk;
            const newRisk = riskLevelMap[row][col];
            const newTotalRisk = startRisk + newRisk;
            const existingTotalRisk = pathMap.get(row + ',' + col);

            if (existingTotalRisk == undefined || newTotalRisk < existingTotalRisk.risk) {
                pathMap.set(row + ',' + col, {
                    risk: newTotalRisk,
                    current: row + ',' + col,
                    from: startRow + ',' + startCol,
                });
            }
        });

    visitedNodes.add(startRow + ',' + startCol);

    if (startRow == targetRow && startCol == targetCol) {
        return pathMap;
    }

    const nextNeighbour = Array.from(pathMap.values())
        .filter((it) => {
            const [row, col] = it.current.split(',');
            return hasNotVisited(visitedNodes, row, col);
        })
        .sort((a, b) => a - b)[0]
        .current.split(',');

    return processNode(
        pathMap,
        visitedNodes,
        parseInt(nextNeighbour[0]),
        parseInt(nextNeighbour[1]),
        targetRow,
        targetCol
    );
}

function print(pathMap, end) {
    const chosenPath = new Set();
    let pathInComplete = true;
    let current = end[0] + ',' + end[1];

    while (pathInComplete) {
        const next = pathMap.get(current);
        if (next.from == null) {
            pathInComplete = false;
        }

        chosenPath.add(current);
        current = next.from;
    }

    riskLevelMap.forEach((rowValues, row) => {
        let line = '';

        rowValues.forEach((risk, col) => {
            if (chosenPath.has(row + ',' + col)) {
                line += '\x1b[34m' + risk + '\x1b[0m';
            } else {
                line += risk;
            }
        });

        console.log(line);
    });
}

inputStream.on('close', () => {
    const end = [riskLevelMap.length - 1, riskLevelMap[0].length - 1];
    const map = new Map();
    map.set('0,0', { risk: 0, current: '0,0', from: null });

    const n = processNode(map, new Set(), 0, 0, end[0], end[1]);
    print(n, end);
    console.log(n.get(end[0] + ',' + end[1]));
});
