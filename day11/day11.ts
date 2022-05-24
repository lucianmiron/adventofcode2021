import { readFileSync } from 'fs';
let countFlashes = 0;
const main = () => {
  const input = readFileSync('input-easy.txt', 'utf-8');
  let map: number[][] = [];
  input.split('\r\n').forEach((row, index) => {
    map[index] = [];
    for (let i = 0; i < row.length; i++) {
      map[index][i] = parseInt(row.charAt(i));
    }
  });

  let iterationFullFlash: boolean;

  for (let i = 0; i < 2000; i++) {
    if (i === 195) {
      debugger;
    }
    iterationFullFlash = true;
    for (let y = 0; y < map.length; y++) {
      for (let x = 0; x < map[y].length; x++) {
        if (map[y][x] !== 0) {
          iterationFullFlash = false;
          break;
        }
      }
      if (!iterationFullFlash) break;
    }
    if (iterationFullFlash) {
      console.log('hit');
      console.log(`Iteration full flash ${i + 1}`);
    }

    for (let y = 0; y < map.length; y++) {
      for (let x = 0; x < map[y].length; x++) {
        flashOctopus(map, y, x);
      }
    }

    for (let y = 0; y < map.length; y++) {
      for (let x = 0; x < map[y].length; x++) {
        if (map[y][x] > 9) map[y][x] = 0;
      }
    }
    i++;
  }

  //console.log(countFlashes);
};

const flashOctopus = (map: number[][], y: number, x: number) => {
  map[y][x]++;
  if (map[y][x] !== 10) return;
  countFlashes++;

  for (let i = y - 1; i <= y + 1; i++) {
    for (let j = x - 1; j <= x + 1; j++) {
      if (i == y && j == x) continue;

      try {
        flashOctopus(map, i, j);
      } catch (err) {}
    }
  }
};

main();
