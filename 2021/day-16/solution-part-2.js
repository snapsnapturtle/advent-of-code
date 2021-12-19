const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});
const ACTIONS = {
    SUM: 0,
    PRODUCT: 1,
    MINIMUM: 2,
    MAXIMUM: 3,
    GREATER_THAN: 4,
    LESS_THAN: 5,
    EQUAL_TO: 6,
};

let bits = [];

class Package {
    version;
    children = [];

    constructor(version) {
        this.version = version;
    }

    calculateValue() {
        throw Error('interface method');
    }
}

class Operator extends Package {
    action;

    constructor(version, action) {
        super(version);
        this.action = action;
    }

    addChild(packet) {
        this.children.push(packet);
    }

    calculateValue() {
        switch (this.action) {
            case ACTIONS.SUM:
                return this.children.reduce((acc, it) => {
                    acc += it.calculateValue();
                    return acc;
                }, 0);
            case ACTIONS.PRODUCT:
                return this.children.reduce((acc, it) => {
                    acc *= it.calculateValue();
                    return acc;
                }, 1);
            case ACTIONS.MINIMUM:
                return Math.min(...this.children.map((it) => it.calculateValue()));
            case ACTIONS.MAXIMUM:
                return Math.max(...this.children.map((it) => it.calculateValue()));
            case ACTIONS.GREATER_THAN:
                return this.children[0].calculateValue() > this.children[1].calculateValue() ? 1 : 0;
            case ACTIONS.LESS_THAN:
                return this.children[0].calculateValue() < this.children[1].calculateValue() ? 1 : 0;
            case ACTIONS.EQUAL_TO:
                return this.children[0].calculateValue() == this.children[1].calculateValue() ? 1 : 0;
        }
    }
}

class Literal extends Package {
    value;

    constructor(version, value) {
        super(version);
        this.value = value;
    }

    calculateValue() {
        return this.value;
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

function mapTypeIdToAction(typeId) {
    switch (typeId) {
        case 0:
            return ACTIONS.SUM;
        case 1:
            return ACTIONS.PRODUCT;
        case 2:
            return ACTIONS.MINIMUM;
        case 3:
            return ACTIONS.MAXIMUM;
        case 5:
            return ACTIONS.GREATER_THAN;
        case 6:
            return ACTIONS.LESS_THAN;
        case 7:
            return ACTIONS.EQUAL_TO;
    }
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
        packet = new Operator(version, mapTypeIdToAction(typeId));

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

inputStream.on('close', () => {
    const package = readPackage(bits);

    console.log(package.calculateValue());
});
