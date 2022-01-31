const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

class Reactor {
    coreMap = new Map();
    dimensions = {
        x: [],
        y: [],
        z: [],
    };

    constructor(dimensions) {
        this.dimensions = dimensions;
    }

    handleInstruction(input) {
        const [_, instruction, xStartString, xEndString, yStartString, yEndString, zStartString, zEndString] =
            input.match(/(on|off)\sx=(-?\d+)\.\.(-?\d+)\,y=(-?\d+)\.\.(-?\d+)\,z=(-?\d+)\.\.(-?\d+)/);

        const xRange = [parseInt(xStartString), parseInt(xEndString)];
        const yRange = [parseInt(yStartString), parseInt(yEndString)];
        const zRange = [parseInt(zStartString), parseInt(zEndString)];

        if (
            !this._isInDimensions(xRange[0], yRange[0], zRange[0]) ||
            !this._isInDimensions(xRange[1], yRange[1], zRange[1])
        ) {
            console.log('skip instruction set:', input);
            return;
        }

        for (let x = xRange[0]; x <= xRange[1]; x++) {
            for (let y = yRange[0]; y <= yRange[1]; y++) {
                for (let z = zRange[0]; z <= zRange[1]; z++) {
                    this._modifyCube(instruction, x, y, z);
                }
            }
        }
    }

    getEnabledCubes() {
        let enabledCubeCount = 0;
        for (const [_, enabled] of this.coreMap) {
            if (enabled) {
                enabledCubeCount++;
            }
        }

        return enabledCubeCount;
    }

    _modifyCube(instruction, x, y, z) {
        this.coreMap.set(`${x},${y},${z}`, instruction == 'on' ? true : false);
    }

    _isInDimensions(x, y, z) {
        return (
            (x >= this.dimensions.x[0] && x <= this.dimensions.x[1]) &&
            (y >= this.dimensions.y[0] && y <= this.dimensions.y[1]) &&
            (z >= this.dimensions.z[0] && z <= this.dimensions.z[1])
        );
    }
}
const reactor = new Reactor({
    x: [-50, 50],
    y: [-50, 50],
    z: [-50, 50],
});

inputStream.on('line', (line) => {
    reactor.handleInstruction(line);
});

inputStream.on('close', () => {
    console.log(reactor.getEnabledCubes());
});
