const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

let numbersToDrawFrom = [];
const boards = [];

inputStream.on('line', (line) => {
    if (numbersToDrawFrom.length === 0) {
        numbersToDrawFrom = line.split(',').map((it) => Number(it));
        return;
    }

    if (line === '') {
        boards.push([]);
    } else {
        boards[boards.length - 1].push(
            line
                .split(' ')
                .filter((it) => it !== '')
                .map((it) => Number(it.trim()))
        );
    }
});

function countNumbersToWin(board) {
    const checks = board.map((row) => ({ numbers: row, matches: 0 }));

    for (let col = 0; col < board[0].length; col++) {
        checks.push({
            numbers: board.map((it) => it[col]),
            matches: 0,
        });
    }

    let lastDrawnNumber;
    let moves;

    numbersToDrawFrom.some((drawnNumber, index) => {
        return checks.some((check) => {
            if (check.numbers.includes(drawnNumber)) {
                check.matches = check.matches + 1;
            }

            if (check.matches === check.numbers.length) {
                lastDrawnNumber = drawnNumber;
                moves = index + 1;
                return true;
            }
        });
    });

    if (lastDrawnNumber) {
        const drawnNumbers = numbersToDrawFrom.slice(0, moves);

        return {
            lastDrawnNumber,
            moves,
            sumUnmarked: board.flat().reduce((acc, number) => {
                if (!drawnNumbers.includes(number)) {
                    return (acc += number);
                }

                return acc;
            }, 0),
        };
    }

    return null;
}

inputStream.on('close', () => {
    const winningBoards = [];

    boards.forEach((board) => {
        const b = countNumbersToWin(board);

        if (b) {
            winningBoards.push(b);
        }
    });

    const firstWinner = winningBoards.sort((a, b) => (a.moves > b.moves ? 1 : -1))[0];

    console.log(firstWinner.lastDrawnNumber * firstWinner.sumUnmarked);
});
