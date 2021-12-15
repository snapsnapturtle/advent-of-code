const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

let mapIndex = 0;
let startIndex;
let endIndex;
let specialCaveIndex;

const uniquePaths = new Set();
const caveToIndexMap = {};
const singleUseCaveIndices = [];

const board = [];

function addCave(cave) {
    if (caveToIndexMap[cave] == undefined) {
        caveToIndexMap[cave] = mapIndex;
        board.push([]);

        if (cave[0] !== cave[0].toUpperCase()) {
            singleUseCaveIndices.push(mapIndex);
        }

        if (cave.toLowerCase() == 'start') {
            startIndex = mapIndex;
        }

        if (cave.toLowerCase() == 'end') {
            endIndex = mapIndex;
        }

        mapIndex++;
    }
}

function addConnection(start, end) {
    addCave(start);
    addCave(end);

    board[caveToIndexMap[start]].push(caveToIndexMap[end]);
}

function canVisitCave(caveIndex, visits) {
    if (singleUseCaveIndices.includes(caveIndex)) {
        if (caveIndex == specialCaveIndex) {
            return visits <= 1;
        } else {
            return visits == 0;
        }
    }

    return true;
}

function getPaths(start, destination, visitedCaves, currentPath) {
    if (start == destination) {
        uniquePaths.add(currentPath.join(','));
        return;
    }

    visitedCaves[start]++;

    for (let i = 0; i < board[start].length; i++) {
        if (canVisitCave(board[start][i], visitedCaves[board[start][i]])) {
            currentPath.push(board[start][i]);
            getPaths(board[start][i], destination, visitedCaves, currentPath);
            currentPath.splice(currentPath.length - 1, 1);
        }
    }

    visitedCaves[start]--;
}

function countAllPaths(s, d) {
    const v = Object.keys(caveToIndexMap).length;

    let visited = new Array(v);
    for (let i = 0; i < v; i++) visited[i] = 0;
    let pathList = [];

    pathList.push(s);

    getPaths(s, d, visited, pathList);
}

inputStream.on('line', (line) => {
    const [start, end] = line.split('-');

    addConnection(start, end);
    addConnection(end, start);
});

inputStream.on('close', () => {
    singleUseCaveIndices
        .filter((it) => it != startIndex && it != endIndex)
        .forEach((specialCave) => {
            specialCaveIndex = specialCave;
            countAllPaths(startIndex, endIndex);
        });

    console.log(uniquePaths.size);
});
