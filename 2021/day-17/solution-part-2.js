const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

function updatePosition(position, velocity) {
    return [position[0] + velocity[0], position[1] + velocity[1]];
}

function updateVelocity(velocity) {
    let x = velocity[0];

    if (x > 0) {
        x--;
    } else if (x < 0) {
        x++;
    }

    return [x, velocity[1] - 1];
}

class ProbeCalculator {
    position = [0, 0];
    velocity = [];
    targetArea = {
        x1: 0,
        x2: 0,
        y1: 0,
        y2: 0,
    };

    acceptedVelocities = [];

    setTargetArea(targetArea) {
        this.targetArea = targetArea;
    }

    findAvailableVelocities() {
        const lowestVector = probeCalculator.getLowestVector();

        const vectorsX = Array.from({ length: lowestVector[0] }, (_, i) => i + 1);
        const vectorsY = Array.from({ length: Math.abs(lowestVector[1] - 1) * 2 }, (_, i) => lowestVector[1] + i);

        const possibleVectors = [];

        vectorsX.forEach((x) => {
            vectorsY.forEach((y) => {
                possibleVectors.push([x, y]);
            });
        });

        possibleVectors.forEach((it) => this.processThrow(it));

        return this.acceptedVelocities;
    }

    processThrow(velocity) {
        this.position = [0, 0];
        this.velocity = velocity;

        while (true) {
            this.position = updatePosition(this.position, this.velocity);

            if (this.isInTargetArea()) {
                this.acceptedVelocities.push(velocity);
                break;
            }

            if (this.isBelowTargetArea() || this.isRightOfTargetArea()) {
                break;
            }

            this.velocity = updateVelocity(this.velocity);
        }
    }

    isInTargetArea() {
        const xInTarget = this.position[0] >= this.targetArea.x1 && this.position[0] <= this.targetArea.x2;
        const yInTarget = this.position[1] <= this.targetArea.y1 && this.position[1] >= this.targetArea.y2;

        return xInTarget && yInTarget;
    }

    isRightOfTargetArea() {
        return this.position[0] > this.targetArea.x2;
    }

    isBelowTargetArea() {
        return this.position[1] < this.targetArea.y2;
    }

    getLowestVector() {
        return [this.targetArea.x2, this.targetArea.y2];
    }
}

const probeCalculator = new ProbeCalculator();

inputStream.on('line', (line) => {
    if (line.startsWith('target area: ')) {
        const [_, x1, x2, y1, y2] = line.match(/x=(-?\d+)\.+(-?\d+).+y=(-?\d+)\.+(-?\d+)/);

        probeCalculator.setTargetArea({
            x1: parseInt(Math.min(x1, x2)),
            x2: parseInt(Math.max(x1, x2)),
            y1: parseInt(Math.max(y1, y2)),
            y2: parseInt(Math.min(y1, y2)),
        });
    }
});

inputStream.on('close', () => {
    const vectors = probeCalculator.findAvailableVelocities();

    console.log(vectors.length);
});
