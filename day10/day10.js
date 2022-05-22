"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var paranthesisError = 3;
var straightParError = 57;
var acoladeError = 1197;
var greaterError = 25137;
var main = function () {
    var input = (0, fs_1.readFileSync)('input.txt', 'utf-8');
    var errors = input.split('\r\n').map(function (row) {
        var openingProcessed = [];
        for (var i = 0; i < row.length; i++) {
            if (row.charAt(i).match(/\)|]|}|>/)) {
                if (i == 0)
                    return getErrorFromChar(row.charAt(i));
                //find it's beginning
                for (var j = i; 0 <= j; j--) {
                    let processed = false;
                    for (let k = 0; k < openingProcessed.length; k++) {
                        if (openingProcessed[k] == j) {
                            processed = true;
                            break;
                        }
                    }
                    if (processed) continue;

                    if (row.charAt(i) === ')' &&
                        row.charAt(j).match(/\[|{|</)) {
                        return getErrorFromChar(')');
                    }
                    if (row.charAt(i) == ']' &&
                        row.charAt(j).match(/\(|{|</)) {
                        return getErrorFromChar(']');
                    }
                    if (row.charAt(i) == '}' &&
                        row.charAt(j).match(/\(|\[|</)) {
                        return getErrorFromChar('}');
                    }
                    if (row.charAt(i) == '>' &&
                        row.charAt(j).match(/\(|\[|{/)) {
                        return getErrorFromChar('>');
                    }

                    if (row.charAt(j) == '(' && row.charAt(i) == ')' ||
                        row.charAt(j) == '[' && row.charAt(i) == ']' ||
                        row.charAt(j) == '{' && row.charAt(i) == '}' ||
                        row.charAt(j) == '<' && row.charAt(i) == '>') {
                        openingProcessed.push(j);
                        openingProcessed.push(i);
                        break;
                    }

                }
            }
        }
        return 0;
    });
    var sum = errors.reduce(function (previous, current) { return previous + current; }, 0);
    console.log(sum);
};
var getErrorFromChar = function (character) {
    if (character == ')')
        return paranthesisError;
    if (character == ']')
        return straightParError;
    if (character == '}')
        return acoladeError;
    if (character == '>')
        return greaterError;
    return 0;
};
main();
