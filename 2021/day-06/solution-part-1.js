const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream('./input.txt'),
});

// days to spawn -> fish count
const daysToSpawnMap = new Map([
    [0, 0],
    [1, 0],
    [2, 0],
    [3, 0],
    [4, 0],
    [5, 0],
    [6, 0],
    [7, 0],
    [8, 0],
    [9, 0],
]);

inputStream.on('line', (line) => {
    line.split(',').forEach((ds) => {
        daysToSpawnMap.set(Number(ds), daysToSpawnMap.get(Number(ds)) + 1);
    });
});

function simulateDay() {
    for (let dayToSpawn of daysToSpawnMap.keys()) {
        if (dayToSpawn == 0) {
            // reproduce 🎉

            // increase values so they get reduced in the same cycle (values should be 8 and 6)
            const currentAdults = daysToSpawnMap.get(7);
            const newBabies = daysToSpawnMap.get(0);

            daysToSpawnMap.set(dayToSpawn, 0);
            daysToSpawnMap.set(9, newBabies);
            daysToSpawnMap.set(7, currentAdults + newBabies);
        } else {
            // set amount of fishes to one day before
            daysToSpawnMap.set(dayToSpawn - 1, daysToSpawnMap.get(dayToSpawn));
            daysToSpawnMap.set(dayToSpawn, 0);
        }
    }
}

inputStream.on('close', () => {
    for (let index = 0; index < 80; index++) {
        simulateDay();
    }

    let sum = 0;
    daysToSpawnMap.forEach((v) => {
        sum += v;
    });

    console.info(sum);
});
