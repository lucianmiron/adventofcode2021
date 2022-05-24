"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var paranthesisError = 3;
var straightParError = 57;
var acoladeError = 1197;
var greaterError = 25137;
var main = function () {
    var input = (0, fs_1.readFileSync)('input.txt', 'utf-8');
    let autocompletionScores = [];
    var errors = input.split('\r\n').map(function (row, index) {
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
        autocompletionScores[autocompletionScores.length] = 0;
        for (var i = row.length - 1; i >= 0; i--) {
            if (!openingProcessed.includes(i)) {
                //it's an opening which needs to be closed
                let symbolScore = 0;
                switch (row.charAt(i)) {
                    case "(":
                        symbolScore = 1;
                        break;
                    case "[":
                        symbolScore = 2;
                        break;
                    case "{":
                        symbolScore = 3;
                        break;
                    case "<":
                        symbolScore = 4;
                        break;
                }
                autocompletionScores[autocompletionScores.length - 1] = autocompletionScores[autocompletionScores.length - 1] * 5 + symbolScore;
            }
        }
        return 0;
    });
    //var sum = errors.reduce(function (previous, current) { return previous + current; }, 0);
    insertionSort(autocompletionScores, autocompletionScores.length);

    console.log(autocompletionScores[(autocompletionScores.length - 1) / 2]);
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
