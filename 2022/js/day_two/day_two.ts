import * as fs from 'fs';
import {Col1, Col2, RPS, Score} from './types';

const DATA_PATH = '../data/2.txt';
export function day_two() {
    // get the score from the startegy guid
    getScore();
    ptTwo();
}

async function getScore() {
    const data = await fs.promises.readFile(DATA_PATH).then((dat) => {
        return dat
            .toString()
            .split('\n')
            .map((d) => {
                return d.split(' ');
            });
    });

    let score = 0;
    data.forEach((game: string[]) => {
        let k = Col2[game[1] as keyof typeof Col2];
        const res = gameResult(game);
        score += Score[k] + res;
        console.log(game[1], k, score, Score[k], res);
    });
    console.log();
    console.log('******\n', score);
}

function gameResult(game: string[]): number {
    const a = Col1[game[0] as keyof typeof Col1];
    const b = Col2[game[1] as keyof typeof Col2];
    if (a === b) {
        return Score.draw;
    }
    // rock beats scissors                              scissors beats paper                     paper beats rock
    else if ((b === RPS.ROCK && a === RPS.SCISSORS) || (b === RPS.SCISSORS && a === RPS.PAPER) || (b === RPS.PAPER && a === RPS.ROCK)) {
        return Score.win;
    }

    return Score.lose;
}

async function ptTwo() {
    const data = await fs.promises.readFile(DATA_PATH).then((dat) => {
        return dat
            .toString()
            .split('\n')
            .map((d) => {
                return d.split(' ');
            });
    });

    let score = 0;
    data.forEach((game: string[]) => {
        let k = Col2[game[1] as keyof typeof Col2];
        const res = desiredResult(game);
        score += Score[k] + res;
        console.log(game[1], k, score, Score[k], res);
    });
    console.log();
    console.log('******\n', score);
}

function desiredResult(game: string[]): number {
    return 0
}