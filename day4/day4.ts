import { readFileSync } from 'fs';

var main = function () {
  const input = readFileSync('./input-easy.txt', 'utf-8');
  const inputs = input.split('\n');
  const numberSequence = inputs[0];

  let boards: String[][][] = [];
  let boardCounter = 0,
    rowCounter = 0;
  for (let i = 2; i < inputs.length; i++) {
    const elements = inputs[i].split(' ');
    if (elements.length == 0) continue;
    boards[boardCounter] = [];
    for (let j = 0; j < elements.length; j++) {
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
