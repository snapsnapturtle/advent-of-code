const fs = require("fs");

const lines = fs
    .readFileSync("./input.txt", "utf8")
    .split("\n")
    .filter((it) => it !== "");

const possibleGames = []

lines.forEach((line) => {
    // input example "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
    const [gameString, gameId] = line.match(/Game (\d+):\s/);

    const draws = line.replace(gameString, "").split("; ").map((drawString) => {
        return drawString.split(", ").reduce((acc, countAndColorString) => {
            const [count, color] = countAndColorString.split(" ");

            acc[color] = parseInt(count);

            return acc;
        }, {})
    });

    let possibleDraw = true;

    draws.forEach((draw) => {
        if (draw.red > 12 || draw.green > 13 || draw.blue > 14) {
            possibleDraw = false;
        }
    });

    if (possibleDraw) {
        possibleGames.push(parseInt(gameId));
    }
});

let sum = possibleGames.reduce((partialSum, a) => partialSum + a, 0);

console.log(sum);
