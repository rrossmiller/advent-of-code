use pt1::Coord;
use std::{fs, io};

mod pt1;
mod pt2;

enum Part {
    P1,
    P2,
}
const TEST: bool = false;
fn main() -> std::io::Result<()> {
    let dat = read()?;
    run(&dat, Part::P1);
    run(&dat, Part::P2);

    Ok(())
}

fn run(dat: &str, p: Part) {
    let mut nums: Vec<i32> = Vec::new();
    let mut visited: Vec<Coord> = Vec::new();
    // for every line
    for (i, line) in dat.lines().enumerate() {
        // for every char in the line!()
        for (j, c) in line.chars().enumerate() {
            // if the char is a symbol, look around it for numbers
            if !c.is_digit(10) && c != '.' {
                match p {
                    Part::P1 => pt1::run(dat, i, j, &mut nums, &mut visited),
                    Part::P2 => pt2::run(dat, i, j, &mut nums, &mut visited),
                }
            }
        }
    }
    match p {
        Part::P1 => println!(
            "p1: {} ({})",
            nums.iter().sum::<i32>(),
            nums.iter().sum::<i32>() == 536202
        ),
        Part::P2 => println!(
            "p2: {} ({})",
            nums.iter().sum::<i32>(),
            nums.iter().sum::<i32>() == 78272573
        ),
    }
}
fn read() -> io::Result<String> {
    return if !TEST {
        let f = fs::read_to_string("../../data/3.txt")?;
        Ok(f)
    } else {
        let rtn = "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."
            .to_string();
        Ok(rtn)
    };
}
