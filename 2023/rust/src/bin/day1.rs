use std::{collections::HashMap, fs, io};

fn main() -> std::io::Result<()> {
    let data = read()?;
    // let data = test_data(1);
    pt1(&data);

    // let data = test_data(2);
    pt2(&data);

    Ok(())
}

fn pt1(data: &String) {
    let radix = 10;
    // get nums in each line
    let mut line_nums = Vec::new();
    for line in data.lines() {
        let mut nums = Vec::new();
        for c in line.chars() {
            if c.is_numeric() {
                nums.push(c.to_digit(radix).unwrap())
            }
        }
        line_nums.push(nums);
    }

    // convert the nums in each line into numbers [1-99]
    let mut sum = 0;
    for nums in line_nums.iter() {
        let tens = 10 * nums[0];
        let ones = nums.last().unwrap();
        sum += tens + ones;
    }
    println!("pt1: {}\n", sum);
}

fn pt2(data: &String) {
    let radix = 10;
    let map = HashMap::from([
        ("one", 1),
        ("two", 2),
        ("three", 3),
        ("four", 4),
        ("five", 5),
        ("six", 6),
        ("seven", 7),
        ("eight", 8),
        ("nine", 9),
    ]);

    // get nums in each line
    let mut line_nums = Vec::new();
    for line in data.lines() {
        let mut nums = Vec::new();
        let mut word = String::new();
        // println!("{}", line);
        for c in line.chars() {
            word.push(c);
            // check if one of the map keys is in the word
            for k in map.keys() {
                if word.contains(k) {
                    let n = *map.get(k).unwrap();
                    nums.push(n);
                    // println!("{}: {}", word, n);
                    word.clear();
                    break;
                }
            }
            if c.is_ascii_digit() {
                nums.push(c.to_digit(radix).unwrap());
                // println!("{}: {}", c, c);
                word.clear();
            }
        }
        // println!();
        line_nums.push(nums);
        // break;
    }

    // convert the nums in each line into numbers [1-99]
    let mut sum = 0;
    for (mut i, nums) in line_nums.iter().enumerate() {
        i += 1;
        let tens = 10 * nums[0];
        let ones = nums.last().unwrap();
        println!("{i}: {:?} -> {}", nums, tens + ones);

        sum += tens + ones;
    }

    if sum <= 55061 {
        println!("too low");
    }
    println!("pt2: {}\n", sum);
}

fn read() -> io::Result<String> {
    let f = fs::read_to_string("../data/1.txt")?;
    Ok(f)
}

fn test_data(pt: i32) -> String {
    match pt {
        1 => "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
"
        .to_string(),
        2 => "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen"
            .to_string(),
        _ => "".to_string(),
    }
}
