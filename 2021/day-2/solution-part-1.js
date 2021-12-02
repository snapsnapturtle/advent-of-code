const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

let horizontalPosition = 0;
let depth = 0;

inputStream.on('line', (line) => {
    switch (line[0]) {
        case 'f':
            horizontalPosition += Number(line.replace('forward ', ''));
            break;
        case 'u':
            depth -= Number(line.replace('up ', ''));
            break;
        case 'd':
            depth += Number(line.replace('down ', ''));
            break;
    }
});

inputStream.on('close', () => {
    console.log(horizontalPosition * depth);
});
