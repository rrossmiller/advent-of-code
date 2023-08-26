use std::fs;

use itertools::Itertools;

pub fn run() {
    pt1();
}

#[derive(Debug)]
struct Direction {
    cnt: usize,
    from: usize,
    to: usize,
}
fn pt1() {
    let data_path = "data/day5_pt1.txt";
    let data = fs::read_to_string(data_path).expect(format!("{data_path} not found").as_str());

    let (mut stacks, directions) = read_input(data);

    run_directions(&mut stacks, directions);
    for s in stacks.iter() {
        print!("{}", s.last().unwrap());
    }
    println!("");
}

fn run_directions(stacks: &mut Vec<Vec<char>>, directions: Vec<Direction>) {
    for d in directions.iter() {
        for _ in 0..d.cnt {
            if let Some(val) = stacks[d.from - 1].pop() {
                stacks[d.to - 1].push(val);
            } else {
                break;
            }
        }
    }
}

fn read_input(data: String) -> (Vec<Vec<char>>, Vec<Direction>) {
    let mut stacks: Vec<Vec<char>> = Vec::new();
    // this doesn't have to be hardcoded
    for _ in 0..9 {
        stacks.push(vec![]);
    }

    let mut directions: Vec<Direction> = Vec::new();
    let mut is_directions = false;

    for row in data.lines() {
        // if the row is empty, stacks are done. directions starting
        if row.is_empty() {
            is_directions = true;
        } else if is_directions {
            let spl: Vec<&str> = row.split(" ").collect();
            directions.push(Direction {
                cnt: spl[1].parse().unwrap(),
                from: spl[3].parse().unwrap(),
                to: spl[5].parse().unwrap(),
            });
        }
        // otherwise, it's the starting stack state. parse it
        else {
            parse_row(row, &mut stacks);
        }
    }

    // let mut n = 1;
    // for x in stacks.iter() {
    //     println!("{n}: {:?}", x);
    //     n += 1;
    // }
    //
    // println!("");
    // let mut n = 1;
    // for x in directions.iter() {
    //     println!("{n}: {:?}", x);
    //     n += 1;
    //     if n > 4{
    //         break;
    //     }
    // }
    // println!("");

    stacks.iter_mut().for_each(|s| s.reverse());

    return (stacks, directions);
}

/// scanner/lexer for consuming rows
fn parse_row(row: &str, stacks: &mut Vec<Vec<char>>) {
    let mut col = 0;
    let mut i = 0;
    let chars = row.chars().collect_vec();

    while i < chars.len() {
        let c = chars[i];
        i += 1;
        match c {
            '[' => {
                //next char is the char
                stacks[col].push(chars[i]);
                // consume char, closing bracket and space
                i += 3;
                // advance column
                col += 1;
            }
            ' ' => {
                // it's empty space above a column
                // consume spaces
                i += 3;
                //advance column
                col += 1;
            }
            _ => {
                // it's a column number
            }
        };
    }
}

#[cfg(test)]
mod tests {
    use crate::day5::{self, Direction};
    use std::fs;

    #[test]
    fn t1() {
        let data_path = "data/day5_pt1.txt";
        let data = fs::read_to_string(data_path).expect(format!("{data_path} not found").as_str());
        let (stacks, _) = day5::read_input(data);
        let mut ans = vec!['T', 'R', 'D', 'H', 'Q', 'N', 'P', 'B'];
        ans.reverse();
        assert_eq!(stacks[0], ans);

        let mut ans = vec!['Z', 'W', 'C', 'V'];
        ans.reverse();
        assert_eq!(stacks[5], ans);
    }

    #[test]
    fn t2() {
        let data_path = "data/day5_pt1.txt";
        let data = fs::read_to_string(data_path).expect(format!("{data_path} not found").as_str());
        let (_, directions) = day5::read_input(data);

        assert_eq!(directions[0].cnt, 5);
        assert_eq!(directions[0].from, 4);
        assert_eq!(directions[0].to, 9);
    }
    #[test]
    fn t3() {
        let data_path = "data/day5_pt1.txt";
        let data = fs::read_to_string(data_path).expect(format!("{data_path} not found").as_str());
        let (mut stacks, _) = day5::read_input(data);

        day5::run_directions(
            &mut stacks,
            vec![Direction {
                cnt: 1,
                from: 4,
                to: 9,
            }],
        );

        assert_eq!(stacks[8], vec!['T', 'P', 'M', 'F', 'Z', 'C', 'G', 'C']);
    }
}
