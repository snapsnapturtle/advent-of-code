const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

class Instruction {
    type;
    xRange = [];
    yRange = [];
    zRange = [];

    constructor(input) {
        const [_, type, xStartString, xEndString, yStartString, yEndString, zStartString, zEndString] = input.match(
            /(on|off)\sx=(-?\d+)\.\.(-?\d+)\,y=(-?\d+)\.\.(-?\d+)\,z=(-?\d+)\.\.(-?\d+)/
        );

        this.xRange = [parseInt(xStartString), parseInt(xEndString)].sort((a, b) => a - b);
        this.yRange = [parseInt(yStartString), parseInt(yEndString)].sort((a, b) => a - b);
        this.zRange = [parseInt(zStartString), parseInt(zEndString)].sort((a, b) => a - b);

        this.type = type;
    }

    getRanges() {
        return [this.xRange, this.yRange, this.zRange];
    }
}

class Reactor {
    cubeOnCount = 0;
    processedInstructions = [];

    handleInstruction(input) {
        const instruction = new Instruction(input);
        console.log('> processing', input);

        let cubeCount = this._getCubeCountForArea(...instruction.getRanges());

        if (instruction.type == 'off') {
            cubeCount = 0;
        }

        this.processedInstructions.forEach((previousInstruction) => {
            cubeCount = this._adjustCubeCount(cubeCount, instruction, previousInstruction);
        });

        this.cubeOnCount += cubeCount;
        console.log('> change cube count by', cubeCount);

        this.processedInstructions.push(instruction);
    }

    _getCubeCountForArea(xRange, yRange, zRange) {
        return (xRange[1] - xRange[0] + 1) * (yRange[1] - yRange[0] + 1) * (zRange[1] - zRange[0] + 1);
    }

    _adjustCubeCount(cubeCount, currentInstruction, previousInstruction) {
        const overlappingRanges = this._getOverlappingArea(currentInstruction, previousInstruction);

        if (overlappingRanges) {
            const overlappingCubes = this._getCubeCountForArea(...overlappingRanges);

            if (currentInstruction.type == 'on' && previousInstruction.type == 'on') {
                console.log('  adjust cube count, both "on" to', cubeCount - overlappingCubes);
                cubeCount -= overlappingCubes;
            }

            if (currentInstruction.type == 'off' && previousInstruction.type == 'on') {
                console.log('  adjust cube count, both "off" to', cubeCount - overlappingCubes);
                cubeCount -= overlappingCubes;
            }
        }

        return cubeCount;
    }

    _getOverlappingArea(instructionA, instructionB) {
        const noIntersect =
            instructionB.xRange[0] > instructionA.xRange[1] ||
            instructionB.xRange[1] < instructionA.xRange[0] ||
            instructionB.yRange[0] > instructionA.yRange[1] ||
            instructionB.yRange[1] < instructionA.yRange[0] ||
            instructionB.zRange[0] > instructionA.zRange[1] ||
            instructionB.zRange[1] < instructionA.zRange[0];

        if (noIntersect) {
            return false;
        }

        return [
            [
                Math.max(instructionA.xRange[0], instructionB.xRange[0]),
                Math.min(instructionA.xRange[1], instructionB.xRange[1]),
            ],
            [
                Math.max(instructionA.yRange[0], instructionB.yRange[0]),
                Math.min(instructionA.yRange[1], instructionB.yRange[1]),
            ],
            [
                Math.max(instructionA.zRange[0], instructionB.zRange[0]),
                Math.min(instructionA.zRange[1], instructionB.zRange[1]),
            ],
        ];
    }
}
const reactor = new Reactor();

inputStream.on('line', (line) => {
    reactor.handleInstruction(line);
});

inputStream.on('close', () => {
    console.log('>> final cubes turned on:', reactor.cubeOnCount);
});
