"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var main = function () {
    var input = (0, fs_1.readFileSync)('./input.txt', 'utf-8');
    var inputs = input.split('\n');
    var numberSequence = inputs[0];
    var boards = fillBoards(inputs);
    var _a = drawNumbersLooser(numberSequence, boards), winingBoard = _a[0], lastNumber = _a[1];
    console.log(lastNumber * computeScore(boards[winingBoard]));
};
function drawNumbersLooser(numbers, boards) {
    var winners = [];
    var numbersSplit = numbers.split(',');
    for (var s = 0; s < numbersSplit.length; s++) {
        var number = parseInt(numbersSplit[s]);
        for (var i = 0; i < boards.length; i++) {
            if (winners.includes(i))
                continue;
            for (var j = 0; j < 5; j++) {
                for (var k = 0; k < 5; k++) {
                    if (boards[i][j][k] === number) {
                        boards[i][j][k] = -1;
                        // determine if this makes a win condition
                        var won = true;
                        for (var x = 0; x < 5; x++) {
                            if (boards[i][j][x] !== -1) {
                                won = false;
                            }
                        }
                        if (won) {
                            winners.push(i);
                        }
                        else {
                            won = true;
                            for (var x = 0; x < 5; x++) {
                                if (boards[i][x][k] !== -1) {
                                    won = false;
                                }
                            }
                            if (won) {
                                winners.push(i);
                            }
                        }
                    }
                }
            }
        }
        if (winners.length === boards.length) {
            var lastWinningBoard = winners[winners.length - 1];
            return [lastWinningBoard, number];
        }
    }
}
function computeScore(board) {
    var sum = 0;
    // board.forEach((row) => {
    //   row.forEach((value) => {
    //     if (value !== -1) sum += value;
    //   });
    // });
    for (var i = 0; i < board.length; i++) {
        for (var j = 0; j < board.length; j++) {
            if (board[i][j] !== -1) {
                sum += board[i][j];
            }
        }
    }
    return sum;
}
function drawNumbers(numbers, boards) {
    var numbersSplit = numbers.split(',');
    for (var s = 0; s < numbersSplit.length; s++) {
        var number = parseInt(numbersSplit[s]);
        for (var i = 0; i < boards.length; i++) {
            for (var j = 0; j < 5; j++) {
                for (var k = 0; k < 5; k++) {
                    if (boards[i][j][k] === number) {
                        boards[i][j][k] = -1;
                        // determine if this makes a win condition
                        var won = true;
                        for (var x = 0; x < 5; x++) {
                            if (boards[i][j][x] !== -1) {
                                won = false;
                            }
                        }
                        if (won) {
                            return { i: i, number: number };
                        }
                        won = true;
                        for (var x = 0; x < 5; x++) {
                            if (boards[i][x][k] !== -1) {
                                won = false;
                            }
                        }
                        if (won) {
                            return { i: i, number: number };
                        }
                    }
                }
            }
        }
    }
}
function fillBoards(input) {
    var boards = [];
    var boardCounter = 0, rowCounter = 0;
    for (var i = 2; i < input.length; i++) {
        if (input[i] == '\r')
            continue;
        var elements = input[i]
            .replace('\r', '')
            .split(' ')
            .filter(function (element) { return !(element === ' ' || element === ''); });
        if (!boards[boardCounter])
            boards[boardCounter] = Array.from(Array(1));
        for (var j = 0; j < elements.length; j++) {
            if (!boards[boardCounter][rowCounter])
                boards[boardCounter][rowCounter] = Array.from(Array(5));
            boards[boardCounter][rowCounter][j] = parseInt(elements[j]);
        }
        rowCounter++;
        if (rowCounter % 5 == 0) {
            boardCounter++;
            rowCounter = 0;
        }
    }
    return boards;
}
main();
function lastNumber(winingBoard, lastNumber) {
    throw new Error('Function not implemented.');
}
