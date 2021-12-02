const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

let horizontalPosition = 0;
let depth = 0;
let aim = 0;

inputStream.on('line', (line) => {
    switch (line[0]) {
        case 'f':
            const forwardChange = Number(line.replace('forward ', ''));
            horizontalPosition += forwardChange;
            depth += forwardChange * aim;
            break;
        case 'u':
            aim -= Number(line.replace('up ', ''));
            break;
        case 'd':
            aim += Number(line.replace('down ', ''));
            break;
    }
});

inputStream.on('close', () => {
    console.log(horizontalPosition * depth);
});
