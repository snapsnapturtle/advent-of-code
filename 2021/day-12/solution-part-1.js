const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

let foundPaths = 0;

let mapIndex = 0;
let startIndex;
let endIndex;

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

function getPaths(start, destination, visitedCaves, currentPath) {
    if (start == destination) {
        foundPaths++;
        return;
    }

    if (singleUseCaveIndices.includes(start)) {
        visitedCaves[start] = true;
    }

    for (let i = 0; i < board[start].length; i++) {
        if (!visitedCaves[board[start][i]]) {
            currentPath.push(board[start][i]);
            getPaths(board[start][i], destination, visitedCaves, currentPath);
            currentPath.splice(currentPath.indexOf(board[start][i]), 1);
        }
    }

    visitedCaves[start] = false;
}

function countAllPaths(s, d) {
    const v = Object.keys(caveToIndexMap).length;

    let isVisited = new Array(v);
    for (let i = 0; i < v; i++) isVisited[i] = false;
    let pathList = [];

    pathList.push(s);

    getPaths(s, d, isVisited, pathList);
}

inputStream.on('line', (line) => {
    const [start, end] = line.split('-');

    addConnection(start, end);
    addConnection(end, start);
});

inputStream.on('close', () => {
    countAllPaths(startIndex, endIndex);

    console.log(foundPaths);
});
