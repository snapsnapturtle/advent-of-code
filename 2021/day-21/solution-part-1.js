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

let diceRolls = 0;
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

inputStream.on('close', () => {
    let winningPlayer;
    while (winningPlayer == undefined) {
        for (let index = 0; index < players.length; index++) {
            const player = players[index];
            player.play();
            if (player.score >= 1000) {
                winningPlayer = player;
                break;
            }
        }
    }

    const losingPlayer = players.filter(it => it.name != winningPlayer.name)[0];

    console.log({winningPlayer, losingPlayer, diceRolls});

    console.log(losingPlayer.score * diceRolls);
});
