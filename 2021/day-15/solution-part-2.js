const { search, Graph } = require('./test');

const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

const riskLevelMap = [];

inputStream.on('line', (line) => {
    const riskRow = line.split('').map((it) => parseInt(it));
    riskLevelMap.push(riskRow);
});

function expandMap() {
    const map = [];
    const newColumns = [];

    riskLevelMap.forEach((rowValues, row) => {
        const newRow = [];
        for (let colIncrease = 0; colIncrease <= 4; colIncrease++) {
            rowValues.forEach((risk) => {
                if (risk + colIncrease >= 10) {
                    newRow.push(risk + colIncrease - 9);
                } else {
                    newRow.push(risk + colIncrease);
                }
            });
        }

        newColumns.push(newRow);
    });

    for (let rowIncrease = 0; rowIncrease <= 4; rowIncrease++) {
        newColumns.forEach((rowValues) => {
            map.push(
                rowValues.map((it) => {
                    if (it + rowIncrease >= 10) {
                        return it + rowIncrease - 9;
                    } else {
                        return it + rowIncrease;
                    }
                })
            );
        });
    }

    return map;
}

inputStream.on('close', () => {
    const riskLevels = expandMap();

    var graph = new Graph(riskLevels);
    var start = graph.grid[0][0];
    var end = graph.grid[riskLevels.length - 1][riskLevels[0].length - 1];

    var result = search(graph, start, end);

    console.log(result.reduce((acc, it) => (acc += it.weight), 0));
});
