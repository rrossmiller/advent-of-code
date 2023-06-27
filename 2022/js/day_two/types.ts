// A for 'Rock', B for 'Paper', and C for Scissors
export enum RPS {
    ROCK="Rock",
    PAPER="Paper",
    SCISSORS="Scissors",
}
export const Col1 = {
    A: RPS.ROCK,
    B: RPS.PAPER,
    C: RPS.SCISSORS,
};
// X for Rock, Y for Paper, and Z for Scissors. Winning every time would be
export const Col2 = {
    X: RPS.ROCK,
    Y: RPS.PAPER,
    Z: RPS.SCISSORS,
};

export const Score = {
    Rock: 1,
    Paper: 2,
    Scissors: 3,
    win: 6,
    draw: 3,
    lose: 0,
};
