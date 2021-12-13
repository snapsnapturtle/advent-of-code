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

let score = 0;

function updateScore(invalidCharacter) {
    let score = 0;

    switch (invalidCharacter) {
        case ')':
            score = 3;
            break;
        case ']':
            score = 57;
            break;
        case '}':
            score = 1197;
            break;
        case '>':
            score = 25137;
            break;
        default:
            break;
    }

    score += score
}

inputStream.on('line', (line) => {
    if (line == '') {
        process.exit();
    }

    const current = [];
    const charsFromInput = line.split('');

    for (const char of charsFromInput) {
        if (isOpening(char)) {
            current.push(char);
        }

        if (isClosing(char)) {
            if (isMatch(current[current.length - 1], char)) {
                current.splice(current.length - 1, 1);
            } else {
                updateScore(char);
                console.log('no match! expected %s found %s', current[current.length - 1], char);
                break;
            }
        }
    }
});

inputStream.on('close', () => {
    console.log(score);
});
