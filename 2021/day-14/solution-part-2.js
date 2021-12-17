const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

let polymerString = [];

const characterCounts = {};
const memo = new Map();
const replacementMap = new Map();

let readReplacements = false;

inputStream.on('line', (line) => {
    if (line == '') {
        readReplacements = true;
        return;
    }

    if (readReplacements) {
        const [from, to] = line.split(' -> ');
        replacementMap.set(from, to);
    } else {
        line.split('').forEach((letter) => {
            polymerString.push(letter);
            increaseCharacterCount(characterCounts, letter);
        });
    }
});

function increaseCharacterCount(counts, character, increase = 1) {
    if (counts[character] == undefined) {
        counts[character] = 0;
    }

    counts[character] += increase;

    return counts;
}

function sumCounts(base, objectToAdd) {
    Object.keys(objectToAdd).forEach((key) => {
        increaseCharacterCount(base, key, objectToAdd[key]);
    });
}

function handlePair(characterA, characterB, currentCharacters, currentHeight) {
    if (currentHeight <= 0) {
        return currentCharacters;
    }

    if (memo.has(characterA + characterB + '-' + currentHeight)) {
        return memo.get(characterA + characterB + '-' + currentHeight)
    }

    const characterPair = characterA + characterB;
    const replacementCharacter = replacementMap.get(characterPair);

    if (replacementCharacter) {
        increaseCharacterCount(currentCharacters, replacementCharacter);

        const a = handlePair(characterA, replacementCharacter, {}, currentHeight - 1);
        const b = handlePair(replacementCharacter, characterB, {}, currentHeight - 1);

        sumCounts(currentCharacters, a);
        sumCounts(currentCharacters, b);
    }

    memo.set(characterA + characterB + '-' + currentHeight, currentCharacters);

    return currentCharacters;
}

function replaceCharacters(maxDepth) {
    polymerString.forEach((character, index) => {
        if (polymerString[index + 1]) {
            const counts = handlePair(character, polymerString[index + 1], {}, maxDepth);
            sumCounts(characterCounts, counts);
        }
    });
}

inputStream.on('close', () => {
    replaceCharacters(40);

    let max, min;

    Object.values(characterCounts).forEach((letterCount) => {
        if (!max) {
            max = letterCount;
        } else {
            max = Math.max(max, letterCount);
        }

        if (!min) {
            min = letterCount;
        } else {
            min = Math.min(min, letterCount);
        }
    });

    console.log(max - min);
});
