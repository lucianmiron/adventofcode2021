"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var main = function () {
    var heatmap = buildHeatmap();
    var lowestPoints = getLowestPoints(heatmap);
    // let sum = 0;
    // lowestPoints.forEach((point) => (sum += point));
    // sum += lowestPoints.length;
    // console.log(`sum is ${sum}`);
    var basinsSizes = findBasinsSize(heatmap, lowestPoints);
    insertionSort(basinsSizes, basinsSizes.length);
    console.log(basinsSizes[basinsSizes.length - 1] *
        basinsSizes[basinsSizes.length - 2] *
        basinsSizes[basinsSizes.length - 3]);
};
var insertionSort = function (arr, n) {
    var i, key, j;
    for (i = 1; i < n; i++) {
        key = arr[i];
        j = i - 1;
        /* Move elements of arr[0..i-1], that are
        greater than the key,to one position ahead
        of their current position
        */
        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j = j - 1;
        }
        arr[j + 1] = key;
    }
};
function getLowestPoints(heatmap) {
    var lowestPoints = [];
    for (var i = 0; i < heatmap.length; i++) {
        for (var j = 0; j < heatmap[i].length; j++) {
            //check adjacent points clockwise
            if (i - 1 >= 0 && heatmap[i][j] >= heatmap[i - 1][j])
                continue;
            if (j + 1 <= heatmap[i].length - 1 && heatmap[i][j] >= heatmap[i][j + 1])
                continue;
            if (i + 1 <= heatmap.length - 1 && heatmap[i][j] >= heatmap[i + 1][j])
                continue;
            if (j - 1 >= 0 && heatmap[i][j] >= heatmap[i][j - 1])
                continue;
            lowestPoints.push({
                x: i,
                y: j
            });
        }
    }
    return lowestPoints;
}
var getSizeOfBasin = function (heatmap, basin, visitedLocations) {
    var size = 0;
    for (var i = 0; i < visitedLocations.length; i++) {
        if (visitedLocations[i].x == basin.x && visitedLocations[i].y == basin.y) {
            return 0;
        }
    }
    if (basin.x < 0 ||
        basin.x >= heatmap.length ||
        basin.y < 0 ||
        basin.y >= heatmap[0].length ||
        heatmap[basin.x][basin.y] == 9) {
        return 0;
    }
    size++;
    visitedLocations.push({ x: basin.x, y: basin.y });
    size += getSizeOfBasin(heatmap, { x: basin.x + 1, y: basin.y }, visitedLocations);
    visitedLocations.push({ x: basin.x + 1, y: basin.y });
    size += getSizeOfBasin(heatmap, { x: basin.x, y: basin.y + 1 }, visitedLocations);
    visitedLocations.push({ x: basin.x, y: basin.y + 1 });
    size += getSizeOfBasin(heatmap, { x: basin.x - 1, y: basin.y }, visitedLocations);
    visitedLocations.push({ x: basin.x - 1, y: basin.y });
    size += getSizeOfBasin(heatmap, {
        x: basin.x,
        y: basin.y - 1
    }, visitedLocations);
    visitedLocations.push({ x: basin.x, y: basin.y - 1 });
    return size;
};
var findBasinsSize = function (heatmap, lowestPoints) {
    var basinSizes = [];
    lowestPoints.forEach(function (coord) {
        basinSizes.push(getSizeOfBasin(heatmap, coord, []));
    });
    return basinSizes;
};
function buildHeatmap() {
    var heatmap = [];
    (0, fs_1.readFileSync)('./input.txt', 'utf-8')
        .replace('s', '')
        .split('\r\n')
        .forEach(function (row, index) {
        heatmap[index] = Array.from(Array(row.length));
        for (var i = 0; i < row.length; i++) {
            heatmap[index][i] = parseInt(row.charAt(i));
        }
    });
    return heatmap;
}
main();
