const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

let polymerString = [];

const characterCounts = {};
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
            increaseCharacterCount(letter);
        });
    }
});

function increaseCharacterCount(character) {
    if (characterCounts[character] == undefined) {
        characterCounts[character] = 0;
    }

    characterCounts[character]++;
}

function handlePair(characterA, characterB, maxDepth) {
    if (maxDepth <= 0) {
        return;
    }

    const characterPair = characterA + characterB;
    const replacementCharacter = replacementMap.get(characterPair);

    if (replacementCharacter) {
        increaseCharacterCount(replacementCharacter);
        handlePair(characterA, replacementCharacter, maxDepth - 1);
        handlePair(replacementCharacter, characterB, maxDepth - 1);
    }
}

function replaceCharacters(maxDepth) {
    polymerString.forEach((character, index) => {
        if (polymerString[index + 1]) {
            handlePair(character, polymerString[index + 1], maxDepth);
        }
    });
}

inputStream.on('close', () => {
    replaceCharacters(10);

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
