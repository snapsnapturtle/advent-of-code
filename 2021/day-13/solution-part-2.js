const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const paper = [];
const folds = [];

let readFoldInstructions = false;

inputStream.on('line', (line) => {
    if (line == '') {
        readFoldInstructions = true;
        return;
    }

    if (readFoldInstructions) {
        const [_, axis, position] = line.match(/fold\salong\s([x,y])=(\d+)/);
        folds.push({ axis, position: parseInt(position) });
        if (axis == 'x') {
            foldX(position);
        }

        if (axis == 'y') {
            foldY(position);
        }
    } else {
        const [col, row] = line.split(',');

        if (paper[row] == undefined) {
            paper[row] = [];
        }

        if (paper[row][col] == undefined) {
            paper[row][col] = true;
        }
    }
});

function printPaper(paper) {
    for (let row = 0; row < paper.length; row++) {
        let line = ' ';

        if (paper[row]) {
            for (let col = 0; col < paper[row].length; col++) {
                if (paper[row][col]) {
                    line += '#';
                } else {
                    line += ' ';
                }
            }
        }

        console.log(line);
    }
}

function foldY(position) {
    // slice off first item to ignore the fold line
    const rowsToFold = paper.splice(position, paper.length).slice(1);

    for (let row = 0; row < rowsToFold.length; row++) {
        const newRow = position - row - 1;

        if (rowsToFold[row]) {
            for (let col = 0; col < rowsToFold[row].length; col++) {
                if (rowsToFold[row][col]) {
                    if (!paper[newRow]) {
                        paper[newRow] = [];
                    }

                    paper[newRow][col] = true;
                }
            }
        }
    }
}

function foldX(position) {
    for (let row = 0; row < paper.length; row++) {
        if (paper[row]) {
            const colsToFold = paper[row].splice(position, paper[row].length).slice(1);

            for (let col = 0; col < colsToFold.length; col++) {
                const newCol = position - col - 1;

                if (colsToFold[col]) {
                    paper[row][newCol] = true;
                }
            }
        }
    }
}

inputStream.on('close', () => {
    printPaper(paper);
});
