const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

let increases = 0;
let previousLine;

inputStream.on('line', (line) => {
    if (previousLine && previousLine < Number(line)) {
        increases++;
    }

    previousLine = line;
});

inputStream.on('close', () => {
    console.log(increases);
});
