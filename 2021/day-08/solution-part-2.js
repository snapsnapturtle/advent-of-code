const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

const entries = [];

inputStream.on('line', (line) => {
    const [signalPatterns, outputPatterns] = line.split('|').map((it) => it.trim().split(' '));
    entries.push({ signalPatterns, outputPatterns });
});

function removeUnknownLetters(list, knownLetters) {
    return list.filter((letter) => knownLetters.includes(letter));
}

function isResolved(possibleValues) {
    return (
        possibleValues.a.length == 1 ||
        possibleValues.b.length == 1 ||
        possibleValues.c.length == 1 ||
        possibleValues.d.length == 1 ||
        possibleValues.e.length == 1 ||
        possibleValues.f.length == 1 ||
        possibleValues.g.length == 1
    );
}

function cleanUpPattern(possibleValues) {
    Object.keys(possibleValues).forEach((key) => {
        if (possibleValues[key].length == 1) {
            // it's unique, remove from others
            Object.keys(possibleValues)
                .filter((it) => it != key)
                .forEach((keyToBeRemoved) => {
                    possibleValues[keyToBeRemoved] = possibleValues[keyToBeRemoved].filter(
                        (it) => it != possibleValues[key][0]
                    );
                });
        }
    });

    return possibleValues;
}

function decodePatterns(patterns) {
    const possibleValues = {
        a: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
        b: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
        c: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
        d: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
        e: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
        f: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
        g: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
    };

    // known patterns for numbers: 1, 4, 7
    patterns
        .filter((it) => [2, 3, 4].includes(it.length))
        .forEach((pattern) => {
            const lettersInPattern = pattern.split('');

            if (lettersInPattern.length == 2) {
                // it's number 1
                possibleValues.c = removeUnknownLetters(possibleValues.c, lettersInPattern);
                possibleValues.f = removeUnknownLetters(possibleValues.c, lettersInPattern);
            }

            if (lettersInPattern.length == 3) {
                // it's number 7
                possibleValues.a = removeUnknownLetters(possibleValues.a, lettersInPattern);
                possibleValues.c = removeUnknownLetters(possibleValues.c, lettersInPattern);
                possibleValues.f = removeUnknownLetters(possibleValues.c, lettersInPattern);
            }

            if (lettersInPattern.length == 4) {
                // it's number 4
                possibleValues.b = removeUnknownLetters(possibleValues.b, lettersInPattern);
                possibleValues.c = removeUnknownLetters(possibleValues.c, lettersInPattern);
                possibleValues.d = removeUnknownLetters(possibleValues.d, lettersInPattern);
                possibleValues.f = removeUnknownLetters(possibleValues.f, lettersInPattern);
            }

            cleanUpPattern(possibleValues);
        });

    patterns
        .filter((it) => [5, 6].includes(it.length))
        .forEach((pattern) => {
            const lettersInPattern = pattern.split('');
            if (lettersInPattern.length == 5) {
                // possible values are 2, 3, 5

                const possibleValuesForBarOfThree = [...possibleValues.c, ...possibleValues.f];

                if (possibleValuesForBarOfThree.every((v) => lettersInPattern.includes(v))) {
                    // it's a three!
                    possibleValues.a = removeUnknownLetters(possibleValues.a, lettersInPattern);
                    possibleValues.c = removeUnknownLetters(possibleValues.c, lettersInPattern);
                    possibleValues.d = removeUnknownLetters(possibleValues.d, lettersInPattern);
                    possibleValues.f = removeUnknownLetters(possibleValues.f, lettersInPattern);
                    possibleValues.g = removeUnknownLetters(possibleValues.g, lettersInPattern);
                } else if (
                    possibleValues.b
                        .filter((it) => !possibleValuesForBarOfThree.includes(it))
                        .every((v) => lettersInPattern.includes(v))
                ) {
                    // it's a five!
                    possibleValues.a = removeUnknownLetters(possibleValues.a, lettersInPattern);
                    possibleValues.b = removeUnknownLetters(possibleValues.b, lettersInPattern);
                    possibleValues.d = removeUnknownLetters(possibleValues.d, lettersInPattern);
                    possibleValues.f = removeUnknownLetters(possibleValues.f, lettersInPattern);
                    possibleValues.g = removeUnknownLetters(possibleValues.g, lettersInPattern);
                } else {
                    // it's a two!
                    possibleValues.a = removeUnknownLetters(possibleValues.a, lettersInPattern);
                    possibleValues.c = removeUnknownLetters(possibleValues.c, lettersInPattern);
                    possibleValues.d = removeUnknownLetters(possibleValues.d, lettersInPattern);
                    possibleValues.e = removeUnknownLetters(possibleValues.e, lettersInPattern);
                    possibleValues.g = removeUnknownLetters(possibleValues.g, lettersInPattern);
                }
            }

            if (lettersInPattern.length == 6) {
                // possibleValues are 0, 6, 9
            }

            cleanUpPattern(possibleValues);
        });

    if (!isResolved(possibleValues)) {
        console.log(possibleValues);
        throw Error('unresolved', possibleValues);
    }

    // return the lookup map (wrong => correct wiring)
    return {
        [possibleValues.a[0]]: 'a',
        [possibleValues.b[0]]: 'b',
        [possibleValues.c[0]]: 'c',
        [possibleValues.d[0]]: 'd',
        [possibleValues.e[0]]: 'e',
        [possibleValues.f[0]]: 'f',
        [possibleValues.g[0]]: 'g',
    };
}

function decodeNumber(pattern, patternMap) {
    const numberMap = {
        abcefg: '0',
        cf: '1',
        acdeg: '2',
        acdfg: '3',
        bcdf: '4',
        abdfg: '5',
        abdefg: '6',
        acf: '7',
        abcdefg: '8',
        abcdfg: '9',
    };

    const resolvedPattern = pattern
        .split('')
        .map((it) => patternMap[it])
        .sort()
        .join('');

    if (numberMap[resolvedPattern]) {
        return numberMap[resolvedPattern];
    } else {
        console.log(pattern, patternMap);
        throw Error('number pattern not found for code');
    }
}

inputStream.on('close', () => {
    let sum = 0;
    entries.forEach((it) => {
        const mapping = decodePatterns(it.signalPatterns);
        const decodedNumbers = it.outputPatterns.map((p) => decodeNumber(p, mapping));

        sum += Number(decodedNumbers.join(''));
    });

    console.log(sum);
});
