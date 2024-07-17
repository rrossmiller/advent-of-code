use std::collections::{HashSet, VecDeque};

use indicatif::ProgressBar;

#[derive(Debug, Clone, Copy)]
pub enum Dir {
    UP,
    DOWN,
    LEFT,
    RIGHT,
}

#[derive(Debug, Hash, PartialEq, Eq)]
struct Tile(usize, usize);

#[derive(Debug, Clone, Copy)]
pub struct Beam {
    pub x: usize,
    pub y: usize,
    pub dir: Dir,
}

pub fn run(data: &str, start: Option<Beam>, runs: Option<u64>) -> usize {
    let tiles = data
        .lines()
        .map(|x| x.chars().collect::<Vec<char>>()) // split each line into chars
        .collect::<Vec<Vec<char>>>(); // indexible tile array

    let mut active_tiles = HashSet::new();
    let b = match start {
        Some(b) => b,
        None => Beam {
            x: 0,
            y: 0,
            dir: Dir::RIGHT,
        },
    };
    let mut beams = VecDeque::from([b]);

    let n = if runs.is_none() {
        100_000_000
    } else {
        runs.unwrap()
    };
    let use_pb = if start.is_none() { true } else { false };
    let pb = ProgressBar::new(n);
    for _ in 0..n {
        if use_pb {
            pb.inc(1);
        }
        if beams.is_empty() {
            break;
        }
        let mut beam = beams.pop_front().unwrap();
        // move the beams
        match beam.dir {
            Dir::RIGHT => {
                // if the beam is in bounds, peek next tile
                // else, it flys off to inf
                if beam.x + 1 <= tiles[0].len() {
                    let t = tiles[beam.y][beam.x];

                    // add to active tiles
                    active_tiles.insert(Tile(beam.y, beam.x));

                    // beam result of tile
                    match t {
                        '|' => {
                            // split vertical
                            beams.push_back(Beam {
                                x: beam.x,
                                y: beam.y + 1,
                                dir: Dir::DOWN,
                            });

                            if beam.y > 0 {
                                beam.dir = Dir::UP;
                                beam.y -= 1;
                                beams.push_back(beam);
                            }
                        }
                        '\\' => {
                            // reflect down
                            beam.dir = Dir::DOWN;
                            beam.y += 1;
                            beams.push_back(beam);
                        }
                        '/' => {
                            // reflect up
                            if beam.y > 0 {
                                beam.dir = Dir::UP;
                                beam.y -= 1;
                                beams.push_back(beam);
                            }
                        }
                        '.' | '-' | _ => {
                            // move
                            beam.x += 1;
                            beams.push_back(beam);
                        }
                    }
                }
            }
            Dir::LEFT => {
                // // if the beam is in bounds, peek next tile
                let t = tiles[beam.y][beam.x];

                // add to active tiles
                active_tiles.insert(Tile(beam.y, beam.x));

                // beam result of tile
                match t {
                    '|' => {
                        // split vertical
                        beams.push_back(Beam {
                            x: beam.x,
                            y: beam.y + 1,
                            dir: Dir::DOWN,
                        });
                        if beam.y > 0 {
                            beam.dir = Dir::UP;
                            beam.y -= 1;
                            beams.push_back(beam);
                        }
                    }
                    '\\' => {
                        // reflect up
                        if beam.y > 0 {
                            beam.dir = Dir::UP;
                            beam.y -= 1;
                            beams.push_back(beam);
                        }
                    }
                    '/' => {
                        // reflect down
                        beam.dir = Dir::DOWN;
                        beam.y += 1;
                        beams.push_back(beam);
                    }
                    '.' | '-' | _ => {
                        // move
                        if beam.x > 0 {
                            beam.dir = Dir::LEFT;
                            beam.x -= 1;
                            beams.push_back(beam);
                        }
                    }
                }
            }
            Dir::UP => {
                let t = tiles[beam.y][beam.x];

                // add to active tiles
                active_tiles.insert(Tile(beam.y, beam.x));

                // beam result of tile
                match t {
                    '-' => {
                        // split horizontal
                        if beam.x > 0 {
                            beams.push_back(Beam {
                                x: beam.x - 1,
                                y: beam.y,
                                dir: Dir::LEFT,
                            });
                        }
                        beam.dir = Dir::RIGHT;
                        beam.x += 1;
                        beams.push_back(beam);
                    }
                    '\\' => {
                        // reflect LEFT
                        if beam.x > 0 {
                            beam.dir = Dir::LEFT;
                            beam.x -= 1;
                            beams.push_back(beam);
                        }
                    }
                    '/' => {
                        // reflect right
                        beam.dir = Dir::RIGHT;
                        beam.x += 1;
                        beams.push_back(beam);
                    }
                    '.' | '|' | _ => {
                        // move
                        if beam.y > 0 {
                            beam.dir = Dir::UP;
                            beam.y -= 1;
                            beams.push_back(beam);
                        }
                    }
                }
            }
            Dir::DOWN => {
                // if the beam is in bounds, peek next tile
                if beam.y + 1 <= tiles.len() {
                    let t = tiles[beam.y][beam.x];

                    // add to active tiles
                    active_tiles.insert(Tile(beam.y, beam.x));

                    // beam result of tile
                    match t {
                        '-' => {
                            // split horizontal
                            if beam.x > 0 {
                                beams.push_back(Beam {
                                    x: beam.x - 1,
                                    y: beam.y,
                                    dir: Dir::LEFT,
                                });
                            }
                            beam.dir = Dir::RIGHT;
                            beam.x += 1;
                            beams.push_back(beam);
                        }
                        '\\' => {
                            // reflect right
                            beam.dir = Dir::RIGHT;
                            beam.x += 1;
                            beams.push_back(beam);
                        }
                        '/' => {
                            // reflect left
                            if beam.x > 0 {
                                beam.dir = Dir::LEFT;
                                beam.x -= 1;
                                beams.push_back(beam);
                            }
                        }
                        '.' | '|' | _ => {
                            // move
                            beam.y += 1;
                            beams.push_back(beam);
                        }
                    }
                }
            }
        }
    }

    active_tiles.len()
}
//
// #[cfg(test)]
// mod tests {
//     use super::*;
//
//     #[test]
//     fn test() {
//         let f = String::from(
//             r".|...\....
// |.-.\.....
// .....|-...
// ........|.
// ..........
// .........\
// ..../.\\..
// .-.-/..|..
// .|....-|.\
// ..//.|....",
//         );
//
//         let ans = run(&f, None);
//         assert_eq!(ans, 46)
//     }
// }
