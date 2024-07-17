use std::sync::atomic::{AtomicU32, Ordering};

use crate::pt1::{self};
use rayon::prelude::*;

pub fn run(data: &str) -> usize {
    // if row is 0 or last, all. else first and last
    // [....]
    // [.  .]
    // [.  .]
    // [....]
    let tiles = data
        .lines()
        .map(|x| x.chars().collect::<Vec<char>>()) // split each line into chars
        .collect::<Vec<Vec<char>>>(); // indexible tile array
                                      // make list of starting beams

    let start_beams = get_starts(tiles);

    let i = AtomicU32::new(0);
    let max = start_beams
        .par_iter()
        // .iter()
        .map(|s| {
            i.fetch_add(1, Ordering::Relaxed);
            println!("{:?}/{}", i, start_beams.len());
            // pt1::run(data, Some(*s), Some(1000_000))
            pt1::run(data, Some(*s), None)
        })
        .max()
        .expect("This should be a usize");
    if max <= 7276 {
        println!("too low");
    }
    max
}

fn get_starts(tiles: Vec<Vec<char>>) -> Vec<pt1::Beam> {
    let (h, w) = (tiles.len(), tiles[0].len());
    let mut starting_beams = Vec::new();

    // top left
    starting_beams.push(pt1::Beam {
        x: 0,
        y: 0,
        dir: pt1::Dir::RIGHT,
    });
    starting_beams.push(pt1::Beam {
        x: 0,
        y: 0,
        dir: pt1::Dir::DOWN,
    });

    // top row
    for i in 1..w - 1 {
        starting_beams.push(pt1::Beam {
            x: i,
            y: 0,
            dir: pt1::Dir::DOWN,
        });
    }

    // top right
    starting_beams.push(pt1::Beam {
        x: w - 1,
        y: 0,
        dir: pt1::Dir::LEFT,
    });
    starting_beams.push(pt1::Beam {
        x: w - 1,
        y: 0,
        dir: pt1::Dir::DOWN,
    });

    // intermediate rows
    for i in 1..h - 1 {
        for j in vec![0, w - 1] {
            starting_beams.push(pt1::Beam {
                x: j,
                y: i,
                dir: pt1::Dir::DOWN,
            });
        }
    }

    // bottom left
    starting_beams.push(pt1::Beam {
        x: 0,
        y: h - 1,
        dir: pt1::Dir::RIGHT,
    });
    starting_beams.push(pt1::Beam {
        x: 0,
        y: h - 1,
        dir: pt1::Dir::UP,
    });
    // bottom row
    for i in 1..w - 1 {
        starting_beams.push(pt1::Beam {
            x: i,
            y: h - 1,
            dir: pt1::Dir::DOWN,
        });
    }

    // bottom right
    starting_beams.push(pt1::Beam {
        x: w - 1,
        y: h - 1,
        dir: pt1::Dir::LEFT,
    });
    starting_beams.push(pt1::Beam {
        x: w - 1,
        y: h - 1,
        dir: pt1::Dir::UP,
    });
    starting_beams
}

#[cfg(test)]
mod tests {
    use crate::pt2::get_starts;

    #[test]
    fn start_eq_perim_size() {
        let f = String::from(
            r".|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....",
        );
        let tiles = f
            .lines()
            .map(|x| x.chars().collect::<Vec<char>>()) // split each line into chars
            .collect::<Vec<Vec<char>>>(); // indexible tile array
                                          // make list of starting beams

        let (h, w) = (tiles.len(), tiles[0].len());
        let x = get_starts(tiles);
        assert_eq!(x.len(), 2 * (h + w));
        println!("{} == {}", x.len(), 2 * (h + w));
    }
}
