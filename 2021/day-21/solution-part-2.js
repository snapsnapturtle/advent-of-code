const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

class Player {
    name;
    position;
    score = 0;

    constructor(name, startPosition) {
        this.name = name;
        this.position = startPosition;
    }

    setPosition(position) {
        this.position = position;
    }

    play() {
        const diceResult = die.next().value + die.next().value + die.next().value;
        let nextPosition = this.position + diceResult;

        while (nextPosition > 10) {
            nextPosition -= 10;
        }

        this._increaseScore(nextPosition);
        this.position = nextPosition;
    }

    _increaseScore(increase) {
        this.score += increase;
    }
}

const players = [];

const die = (function* diceRoll() {
    let current = 1;

    while (true) {
        diceRolls++;
        yield current;
        current++;
        if (current > 100) {
            current = 1;
        }
    }
})();

inputStream.on('line', (line) => {
    const [_, name, playerPosition] = line.match(/(Player\s\d).+(\d+)/);

    players.push(new Player(name, parseInt(playerPosition)));
});

const dieResults = {
    3: 1,
    4: 3,
    // 5: 6,
    // 6: 7,
    // 7: 6,
    // 8: 3,
    // 9: 1,
};

let scores = {};

function calculateTree(node) {
    if (node.score >= 21) {
        if (!scores[node.moves]) {
            scores[node.moves] = 0;
        }

        scores[node.moves] += node.possibilities;
        return node;
    }

    Object.keys(dieResults).forEach((result) => {
        let nextPosition = node.position + parseInt(result);

        while (nextPosition > 10) {
            nextPosition -= 10;
        }

        node.children.push(
            calculateTree({
                position: nextPosition,
                score: node.score + nextPosition,
                moves: node.moves + 1,
                possibilities: node.possibilities * dieResults[result],
                children: [],
            })
        );
    });

    return node;
}

inputStream.on('close', () => {
    calculateTree({ score: 0, moves: 0, possibilities: 1, position: 4, children: [] });
    const player1 = scores;

    scores = {};

    calculateTree({ score: 0, moves: 0, possibilities: 1, position: 8, children: [] });
    const player2 = scores;

    console.log({ player1, player2 });
});
