use std::{fs, io};
mod pt1;
mod pt2;

#[derive(Debug)]
pub struct Mapping {
    pub dest_start: i64,
    pub src_start: i64,
    pub rng_len: i64,
}

fn main() -> std::io::Result<()> {
    println!("Day 1");
    let (seeds, data) = data()?;
    // let (seeds, data) = test_data();

    let (_, p1) = pt1::run(&seeds, &data);
    println!("Pt1: {}", p1);
    let (_, p2) = pt2::run(&seeds, &data);
    println!("Pt2: {}", p2);

    Ok(())
}

fn data() -> io::Result<(Vec<i64>, Vec<Vec<Mapping>>)> {
    let dat = fs::read_to_string("../../data/5.txt")?;
    Ok(parse(dat))
}

// fn test_data(pt: i64) -> String {
pub fn test_data() -> (Vec<i64>, Vec<Vec<Mapping>>) {
    // destination rng start, src rng start, rng len
    let dat = "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4"
        .to_string();
    parse(dat)
}

fn parse(dat: String) -> (Vec<i64>, Vec<Vec<Mapping>>) {
    let seeds = dat
        .lines()
        .nth(0)
        .unwrap()
        .split(": ")
        .nth(1)
        .unwrap()
        .split(" ")
        .map(|x| x.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();

    //             to soil, to fertilizer, to water, to light, to temp, to humidity, to location
    let mut maps = vec![vec![], vec![], vec![], vec![], vec![], vec![], vec![]];
    let mut i = 0;
    for l in dat.lines().skip(3) {
        if l.len() == 0 {
            continue;
        }

        if l.contains(":") {
            i += 1;
        } else {
            let nums = l
                .split(" ")
                .map(|x| x.parse::<i64>().unwrap())
                .collect::<Vec<i64>>();

            maps[i].push(Mapping {
                dest_start: nums[0],
                src_start: nums[1],
                rng_len: nums[2],
            });
        }
    }
    (seeds, maps)
}
