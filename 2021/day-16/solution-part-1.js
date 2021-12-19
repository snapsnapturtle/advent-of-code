const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

let bits = [];

class Package {
    version;
    children = [];

    constructor(version) {
        this.version = version;
    }
}

class Operator extends Package {
    constructor(version) {
        super(version);
    }

    addChild(packet) {
        this.children.push(packet);
    }
}

class Literal extends Package {
    value;

    constructor(version, value) {
        super(version);
        this.value = value;
    }
}

inputStream.on('line', (line) => {
    bits = line
        .split('')
        .map((it) => parseInt(it, 16).toString(2).padStart(4, '0'))
        .join('')
        .split('');
});

function readBits(bits, numberOfBits) {
    const readBits = bits.slice(0, numberOfBits).join('');
    bits.splice(0, numberOfBits);

    return readBits;
}

function getHeader(bits) {
    const versionBits = readBits(bits, 3);
    const version = parseInt(versionBits, 2);

    const typeIdBits = readBits(bits, 3);
    const typeId = parseInt(typeIdBits, 2);

    return { version, typeId };
}

function getLiteralValue(bits) {
    let finalBits = '';
    let shouldContinue = true;

    while (shouldContinue) {
        const literalBits = readBits(bits, 5);

        if (literalBits[0] == '0') {
            shouldContinue = false;
        }

        finalBits += literalBits[1] + literalBits[2] + literalBits[3] + literalBits[4];
    }

    return parseInt(finalBits, 2);
}

function getLengthForChildren(bits) {
    // determine length type
    const lengthTypeIdBit = readBits(bits, 1);
    let childCountType;
    let lengthBits;

    switch (lengthTypeIdBit) {
        case '0':
            childCountType = 'LENGTH';
            lengthBits = readBits(bits, 15);
            break;
        case '1':
            childCountType = 'PACKET';
            lengthBits = readBits(bits, 11);
            break;
    }

    return {
        countType: childCountType,
        value: parseInt(lengthBits, 2),
    };
}

function readPackage(bits) {
    let packet;
    const { version, typeId } = getHeader(bits);

    if (typeId == 4) {
        // literal package
        const literalValue = getLiteralValue(bits);

        packet = new Literal(version, literalValue);
    } else {
        // operator package
        packet = new Operator(version);

        const { countType, value: length } = getLengthForChildren(bits);

        if (countType == 'LENGTH') {
            const childrenBits = bits.splice(0, length);

            while (childrenBits.length >= 1) {
                packet.addChild(readPackage(childrenBits));
            }
        }

        if (countType == 'PACKET') {
            for (let index = 0; index < length; index++) {
                packet.addChild(readPackage(bits));
            }
        }
    }

    return packet;
}

function calculateVersionSum(package, sum) {
    sum += package.version;

    if (package.children.length > 0) {
        package.children.forEach((child) => {
            sum += calculateVersionSum(child, 0);
        });
    }

    return sum;
}

inputStream.on('close', () => {
    const package = readPackage(bits);
    const versionSum = calculateVersionSum(package, 0);

    console.log(versionSum);
});
