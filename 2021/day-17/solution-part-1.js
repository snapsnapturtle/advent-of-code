const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

function updatePositionY(position, velocity) {
    return [position[0], position[1] + velocity[1]];
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

    previousPositions = [];

    setTargetArea(targetArea) {
        this.targetArea = targetArea;
    }

    getHighestYValue() {
        // get the lowest point of target area and use it as starting velocity
        const startingYVelocity = Math.abs(Math.min(this.targetArea.y1, this.targetArea.y2) + 1);
        this.velocity = [0, startingYVelocity];

        let maxHeight = -1;

        while (this.position[1] > maxHeight) {
            maxHeight = this.position[1];

            this.position = updatePositionY(this.position, this.velocity);
            this.velocity = updateVelocity(this.velocity);
        }

        return maxHeight;
    }
}

const probeCalculator = new ProbeCalculator();

inputStream.on('line', (line) => {
    if (line.startsWith('target area: ')) {
        const [_, x1, x2, y1, y2] = line.match(/x=(-?\d+)\.+(-?\d+).+y=(-?\d+)\.+(-?\d+)/);

        probeCalculator.setTargetArea({
            x1: parseInt(x1),
            x2: parseInt(x2),
            y1: parseInt(y1),
            y2: parseInt(y2),
        });
    }
});

inputStream.on('close', () => {
    const highestYValue = probeCalculator.getHighestYValue();
    console.log(highestYValue);
});
