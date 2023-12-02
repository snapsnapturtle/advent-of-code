const fs = require("fs");

const lines = fs
    .readFileSync("./input.txt", "utf8")
    .split("\n")
    .filter((it) => it !== "");

const minimumPowers = []

lines.forEach((line) => {
    const draws = line.split("; ").map((drawString) => {
        return drawString.split(", ").reduce((acc, countAndColorString) => {
            const [count, color] = countAndColorString.split(" ");

            acc[color] = parseInt(count);

            return acc;
        }, {})
    });

    const minimums = {
        red: 0,
        green: 0,
        blue: 0
    };

    draws.forEach((draw) => {
        minimums.red = Math.max(minimums.red, draw.red || 0);
        minimums.green = Math.max(minimums.green, draw.green || 0);
        minimums.blue = Math.max(minimums.blue, draw.blue || 0);
    });


    minimumPowers.push(minimums.red * minimums.green * minimums.blue);
});

let sum = minimumPowers.reduce((partialSum, a) => partialSum + a, 0);

console.log(sum);
