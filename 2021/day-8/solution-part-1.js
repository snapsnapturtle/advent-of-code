const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

let entries = [];

inputStream.on('line', (line) => {
    const [_, outputPatterns] = line.split('|').map(it => it.trim().split(' '));

    entries = entries.concat(outputPatterns);
});

function hasKnownLength(pattern) {
    return [2, 3, 4, 7].includes(pattern.length);
}

inputStream.on('close', () => {
    const patternsWithKnownLength = entries.filter((it) => hasKnownLength(it)).length;

    console.log(patternsWithKnownLength);
});
