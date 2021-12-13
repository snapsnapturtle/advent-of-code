const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

function isOpening(char) {
    return '([{<'.includes(char);
}

function isClosing(char) {
    return ')]}>'.includes(char);
}

function isMatch(opening, closing) {
    return (
        (opening == '(' && closing == ')') ||
        (opening == '[' && closing == ']') ||
        (opening == '{' && closing == '}') ||
        (opening == '<' && closing == '>')
    );
}

function getScoreValue(invertedClosingChar) {
    if (invertedClosingChar == '(') {
        return 1;
    }
    if (invertedClosingChar == '[') {
        return 2;
    }
    if (invertedClosingChar == '{') {
        return 3;
    }
    if (invertedClosingChar == '<') {
        return 4;
    }
}

const scores = [];

inputStream.on('line', (line) => {
    if (line == '') {
        process.exit();
    }

    const current = [];
    let corrupt = false;
    const charsFromInput = line.split('');

    for (const char of charsFromInput) {
        if (isOpening(char)) {
            current.push(char);
        }

        if (isClosing(char)) {
            if (isMatch(current[current.length - 1], char)) {
                current.splice(current.length - 1, 1);
            } else {
                corrupt = true;
                break;
            }
        }
    }

    if (!corrupt && current.length > 0) {

        scores.push(
            current.reverse().reduce((acc, it) => {
                acc = acc * 5;
                acc = acc + getScoreValue(it);

                return acc;
            }, 0)
        );
    }
});

inputStream.on('close', () => {
    const sortedScores = scores.sort((a, b) => a - b);
    console.log(sortedScores[Math.floor(sortedScores.length / 2)]);
});
