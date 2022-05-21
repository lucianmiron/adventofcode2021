import { readFileSync } from 'fs';

var main = function () {
  const input = readFileSync('./input.txt', 'utf-8');
  const inputs = input.split('\n');
  const numberSequence = inputs[0];
  const boards = fillBoards(inputs);
  let [winingBoard, lastNumber] = drawNumbersLooser(numberSequence, boards);
  console.log(lastNumber * computeScore(boards[winingBoard]));
};

function drawNumbersLooser(numbers: String, boards: number[][][]) {
  const winners: number[] = [];
  const numbersSplit = numbers.split(',');
  for (let s = 0; s < numbersSplit.length; s++) {
    var number = parseInt(numbersSplit[s]);
    for (let i = 0; i < boards.length; i++) {
      if (winners.includes(i)) continue;
      for (let j = 0; j < 5; j++) {
        for (let k = 0; k < 5; k++) {
          if (boards[i][j][k] === number) {
            boards[i][j][k] = -1;
            // determine if this makes a win condition
            let won = true;
            for (let x = 0; x < 5; x++) {
              if (boards[i][j][x] !== -1) {
                won = false;
              }
            }
            if (won) {
              winners.push(i);
            } else {
              won = true;
              for (let x = 0; x < 5; x++) {
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
      const lastWinningBoard = winners[winners.length - 1];
      return [lastWinningBoard, number];
    }
  }
}

function computeScore(board: number[][]) {
  let sum: number = 0;
  // board.forEach((row) => {
  //   row.forEach((value) => {
  //     if (value !== -1) sum += value;
  //   });
  // });
  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board.length; j++) {
      if (board[i][j] !== -1) {
        sum += board[i][j];
      }
    }
  }
  return sum;
}

function drawNumbers(numbers: String, boards: number[][][]) {
  const numbersSplit = numbers.split(',');
  for (let s = 0; s < numbersSplit.length; s++) {
    var number = parseInt(numbersSplit[s]);
    for (let i = 0; i < boards.length; i++) {
      for (let j = 0; j < 5; j++) {
        for (let k = 0; k < 5; k++) {
          if (boards[i][j][k] === number) {
            boards[i][j][k] = -1;
            // determine if this makes a win condition
            let won = true;
            for (let x = 0; x < 5; x++) {
              if (boards[i][j][x] !== -1) {
                won = false;
              }
            }
            if (won) {
              return { i, number };
            }
            won = true;
            for (let x = 0; x < 5; x++) {
              if (boards[i][x][k] !== -1) {
                won = false;
              }
            }
            if (won) {
              return { i, number };
            }
          }
        }
      }
    }
  }
}

function fillBoards(input: String[]): number[][][] {
  let boards: number[][][] = [];
  let boardCounter = 0,
    rowCounter = 0;
  for (let i = 2; i < input.length; i++) {
    if (input[i] == '\r') continue;
    const elements = input[i]
      .replace('\r', '')
      .split(' ')
      .filter((element) => !(element === ' ' || element === ''));
    if (!boards[boardCounter]) boards[boardCounter] = Array.from(Array(1));
    for (let j = 0; j < elements.length; j++) {
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
function lastNumber(winingBoard: any, lastNumber: any) {
  throw new Error('Function not implemented.');
}
