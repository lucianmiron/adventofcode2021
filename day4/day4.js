"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var main = function () {
    var input = (0, fs_1.readFileSync)('./input-easy.txt', 'utf-8');
    var inputs = input.split('\n');
    var numberSequence = inputs[0];
    var boards = [];
    var boardCounter = 0, rowCounter = 0;
    for (var i = 2; i < inputs.length; i++) {
        var elements = inputs[i].split(' ');
        if (elements.length == 0)
            continue;
        boards[boardCounter] = [];
        for (var j = 0; j < elements.length; j++) {
            boards[boardCounter][rowCounter][j] = elements[j];
        }
        rowCounter++;
        if (rowCounter / 5 == 0) {
            boardCounter++;
            rowCounter = 0;
        }
    }
    console.log(boards);
};
main();
