import * as fs from "fs";

// data
let f = fs.readFileSync("../data/9.txt", "utf8").split("\n");
// f = `0 3 6 9 12 15
// 1 3 6 10 15 21
// 10 13 16 21 30 45`.split("\n");
// make each line a list of numbers
const data = f.map((e) => e.split(" ").map((num) => parseInt(num)));

// p1
let sum = 0n;
for (let i = 0; i < data.length; i++) {
  let txt = data[i].join("  ");
  console.log(txt);
  const v = processLine(data[i], 0n, 1);

  console.log(typeof v, typeof BigInt(data[i][data[i].length - 1]));
  console.log(v, v + BigInt(data[i][data[i].length - 1]));
  console.log();
  sum += v + BigInt(data[i][data[i].length - 1]);
}
console.log("pt1:", sum);

function processLine(line: number[], prevEnd: bigint, lvl: number): bigint {
  const zeroes = line.every((e) => e === 0);
  if (zeroes) {
    return prevEnd;
  }

  let diffTxt = "  ".repeat(lvl);
  const diffs = [];
  for (let i = 1; i < line.length; i++) {
    const d = line[i] - line[i - 1];
    diffs.push(d);

    // logging
    if (line[i - 1] > 9 && d < 10) {
      diffTxt += " ";
    }
    diffTxt += `${d}`;
    diffTxt += "  ";
    //
  }
  console.log(diffTxt);
  return prevEnd + processLine(diffs, BigInt(diffs[diffs.length - 1]), lvl + 1);
}
