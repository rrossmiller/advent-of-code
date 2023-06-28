use std::{fs, iter};

use itertools::izip;
const UPPERCASE_OFFSET: u8 = 65 - 26;
const LOWERCASE_OFFSET: u8 = 97;

pub fn run() {
    part1();
    part2();
}

fn part1() {
    let mut ttl: usize = 0;
    let sacks = fs::read_to_string("data/day3_pt1.txt").expect("err reading file");
    for sack in sacks.lines() {
        // split the sack into each of its compartments
        let (c1, c2) = sack.split_at(sack.len() / 2);

        // split into chars
        let c1 = c1.chars().collect::<Vec<char>>();
        let c2 = c2.chars().collect::<Vec<char>>();

        // make boolean arrays
        let c1 = make_alpha_array(c1);
        let c2 = make_alpha_array(c2);

        for (idx, (i, j)) in iter::zip(c1, c2).enumerate() {
            if i && j {
                ttl += idx + 1;
            }
        }
    }
    println!("part 1 total: {}\n", ttl);
}
fn part2() {
    let mut ttl = 0;
    let sacks = fs::read_to_string("data/day3_pt2.txt").expect("err reading file");
    let mut group: [&str; 2] = [""; 2];
    for (i, sack) in sacks.lines().enumerate() {
        // every 3 sacks
        if i % 3 == 2 {
            let g1 = make_alpha_array(group[0].chars().collect::<Vec<char>>());
            let g2 = make_alpha_array(group[1].chars().collect::<Vec<char>>());
            let g3 = make_alpha_array(sack.chars().collect::<Vec<char>>());

            for (idx, (i, j, k)) in izip!(g1, g2, g3).enumerate() {
                if (i && j) && (i && k) {
                    ttl += idx + 1;
                }
            }
        } else {
            let idx = i % 3;
            group[idx] = sack;
        }
    }
    println!("part 2 total: {}\n", ttl);
}

fn make_alpha_array(chars: Vec<char>) -> [bool; 52] {
    let mut rtn: [bool; 52] = [false; 52];

    for c in chars {
        let idx = match c.is_lowercase() {
            true => c as u8 - LOWERCASE_OFFSET,
            false => c as u8 - UPPERCASE_OFFSET,
        } as usize;

        rtn[idx] = true;
    }
    rtn
}

#[cfg(test)]
mod tests {
    use super::make_alpha_array;

    #[test]
    fn test_make_alpha_array() {
        let mut ans: [bool; 52] = [false; 52];
        ans[0] = true;
        ans[25] = true;
        ans[26] = true;
        ans[51] = true;
        let i = vec!['a', 'z', 'A', 'Z'];
        let x = make_alpha_array(i);
        assert_eq!(x, ans);
    }
}
