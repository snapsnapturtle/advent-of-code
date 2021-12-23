const inputStream = require('readline').createInterface({
    input: require('fs').createReadStream(__dirname + '/input.txt'),
});

class Image {
    image = [];
    algorithm = [];
    width = 0;
    infinityValue = '0';

    setAlgorithm(algorithmLine) {
        this.algorithm = algorithmLine;
    }

    updateInfinityValue() {
        this.infinityValue = this.algorithm[parseInt(this.infinityValue.repeat(9), 2)] == '#' ? '1' : '0';
    }

    getImageHeight() {
        return this.image.length;
    }

    addToImage(pixels) {
        this.width = pixels.length;
        this.image.push(pixels);
    }

    getEnhancedPixelValue(row, col) {
        return this.algorithm[this._getEnhancementIndexForPixel(row, col)];
    }

    getLitPixelCount() {
        return this.image.flat().reduce((sum, pixel) => {
            const pixelValue = pixel == '#' ? 1 : 0;

            return sum + pixelValue;
        }, 0);
    }

    _getBinaryFromPixelValue(row, col) {
        if (this.image[row] == undefined) {
            return this.infinityValue;
        }

        const imageValue = this.image[row][col];

        if (imageValue) {
            if (imageValue == '#') {
                return '1';
            } else {
                return '0';
            }
        } else {
            return this.infinityValue;
        }
    }

    _getEnhancementIndexForPixel(row, col) {
        // all pixels in a 3x3 grid have to be taken into account
        const binaryString =
            this._getBinaryFromPixelValue(row - 1, col - 1) +
            this._getBinaryFromPixelValue(row - 1, col) +
            this._getBinaryFromPixelValue(row - 1, col + 1) +
            this._getBinaryFromPixelValue(row, col - 1) +
            this._getBinaryFromPixelValue(row, col) +
            this._getBinaryFromPixelValue(row, col + 1) +
            this._getBinaryFromPixelValue(row + 1, col - 1) +
            this._getBinaryFromPixelValue(row + 1, col) +
            this._getBinaryFromPixelValue(row + 1, col + 1);

        return parseInt(binaryString, 2);
    }

    print() {
        for (let row = 0; row < this.image.length; row++) {
            let currentLine = '';

            for (let col = 0; col < this.width; col++) {
                currentLine += this.image[row][col];
            }

            console.log(currentLine);
        }
    }
}

function enhance(inputImage) {
    const newImage = new Image();
    newImage.infinityValue = inputImage.infinityValue;
    newImage.setAlgorithm(inputImage.algorithm);

    for (let row = -1; row < inputImage.getImageHeight() + 1; row++) {
        const currentRow = [];

        for (let col = -1; col < inputImage.width + 1; col++) {
            currentRow.push(inputImage.getEnhancedPixelValue(row, col));
        }

        newImage.addToImage(currentRow);
    }

    newImage.updateInfinityValue();

    return newImage;
}

const initialInputImage = new Image();

inputStream.on('line', (line) => {
    const lineCharacters = line.split('');
    if (initialInputImage.algorithm.length == 0) {
        initialInputImage.setAlgorithm(lineCharacters);
    } else if (line != '') {
        initialInputImage.addToImage(lineCharacters);
    }
});

inputStream.on('close', () => {
    let image = initialInputImage;

    for (let _ = 0; _ < 50; _++) {
        image = enhance(image);
        // image.print();
    }

    console.log(image.getLitPixelCount());
});
