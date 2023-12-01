const fs = require("fs");

const lines = fs
    .readFileSync("./input.txt", "utf8")
    .split("\n")
    .filter((it) => it !== "");


const firstAndLastNumbers = [];

const thingsToFind = {
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9
}

lines.forEach((line) => {
    const matches = [];

    Object.keys(thingsToFind).forEach((thingToFind) => {
        let index = -1;

        while ((index = line.indexOf(thingToFind, index + 1)) >= 0) {
            matches.push({
                indexAt: index,
                number: thingsToFind[thingToFind]
            })
        }
    });

    matches.sort((a, b) => a.indexAt - b.indexAt)

    const firstNumber = matches[0].number;
    const lastNumber = matches[matches.length - 1].number;

    firstAndLastNumbers.push(parseInt(firstNumber + "" + lastNumber));
});

const sum = firstAndLastNumbers.reduce((partialSum, a) => partialSum + a, 0);

console.log(sum);
