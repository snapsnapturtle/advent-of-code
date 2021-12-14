const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

let octoputEnergy = [];

inputStream.on('line', (line) => {
    octoputEnergy.push(
        line.split('').map((it) => ({
            energy: Number(it),
            hasFlashed: false,
        }))
    );
});

function increaseEnergy(grid, row, col) {
    if (grid[row] != undefined && grid[row][col] != undefined && grid[row][col].hasFlashed == false) {
        grid[row][col].energy += 1;
    }
}

function performStep() {
    const newOctopusEnergy = [];

    octoputEnergy.forEach((rowValues, row) => {
        newOctopusEnergy.push([]);

        rowValues.forEach((octopus, col) => {
            newOctopusEnergy[row][col] = {
                energy: octopus.energy + 1,
                hasFlashed: false,
            };
        });
    });

    let hasAtLeastOneFlash = true;
    while (hasAtLeastOneFlash == true) {
        hasAtLeastOneFlash = false;

        newOctopusEnergy.forEach((rowValues, row) => {
            rowValues.forEach((octopus, col) => {
                if (octopus.energy > 9 && octopus.hasFlashed == false) {
                    // it flashes
                    hasAtLeastOneFlash = true;

                    newOctopusEnergy[row][col] = {
                        energy: 0,
                        hasFlashed: true,
                    };

                    // update surrounding octopus

                    // top left
                    increaseEnergy(newOctopusEnergy, row - 1, col - 1);
                    // top
                    increaseEnergy(newOctopusEnergy, row - 1, col);
                    // top right
                    increaseEnergy(newOctopusEnergy, row - 1, col + 1);
                    // right
                    increaseEnergy(newOctopusEnergy, row, col + 1);
                    // bottom right
                    increaseEnergy(newOctopusEnergy, row + 1, col + 1);
                    // bottom
                    increaseEnergy(newOctopusEnergy, row + 1, col);
                    // bottom left
                    increaseEnergy(newOctopusEnergy, row + 1, col - 1);
                    // left
                    increaseEnergy(newOctopusEnergy, row, col - 1);
                }
            });
        });
    }

    octoputEnergy = [...newOctopusEnergy];
}

function isNotSynchronized() {
    return octoputEnergy.flat().some((it) => it.energy > 0);
}

inputStream.on('close', () => {
    let steps = 0;

    while (isNotSynchronized()) {
        performStep();
        steps++;
    }

    console.log(steps);
});
