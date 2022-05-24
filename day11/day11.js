"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var countFlashes = 0;
var main = function () {
    var input = (0, fs_1.readFileSync)('input.txt', 'utf-8');
    var map = [];
    input.split('\r\n').forEach(function (row, index) {
        map[index] = [];
        for (var i = 0; i < row.length; i++) {
            map[index][i] = parseInt(row.charAt(i));
        }
    });
    var iterationFullFlash;
    for (var i = 0; i < 2000; i++) {
        iterationFullFlash = true;
        for (var y = 0; y < map.length; y++) {
            for (var x = 0; x < map[y].length; x++) {
                if (map[y][x] !== 0) {
                    iterationFullFlash = false;
                    break;
                }
            }
            if (!iterationFullFlash)
                break;
        }
        if (iterationFullFlash) {
            console.log("Iteration full flash ".concat(i));
            break;
        }
        for (var y = 0; y < map.length; y++) {
            for (var x = 0; x < map[y].length; x++) {
                flashOctopus(map, y, x);
            }
        }
        for (var y = 0; y < map.length; y++) {
            for (var x = 0; x < map[y].length; x++) {
                if (map[y][x] > 9)
                    map[y][x] = 0;
            }
        }
    }
    //console.log(countFlashes);
};
var flashOctopus = function (map, y, x) {
    map[y][x]++;
    if (map[y][x] !== 10)
        return;
    countFlashes++;
    for (var i = y - 1; i <= y + 1; i++) {
        for (var j = x - 1; j <= x + 1; j++) {
            if (i == y && j == x || i < 0 || i >= map.length || j < 0 || j >= map[i].length)
                continue;
            try {
                flashOctopus(map, i, j);
            }
            catch (err) { }
        }
    }
};
main();
