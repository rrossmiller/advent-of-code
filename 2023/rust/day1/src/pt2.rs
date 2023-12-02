use std::collections::HashMap;

// pub fn run(data: &String) {
//     let radix = 10;
//     let map = HashMap::from([
//         ("one", 1),
//         ("two", 2),
//         ("three", 3),
//         ("four", 4),
//         ("five", 5),
//         ("six", 6),
//         ("seven", 7),
//         ("eight", 8),
//         ("nine", 9),
//     ]);
//
//     // get nums in each line
//     let mut line_nums = Vec::new();
//     for line in data.lines() {
//         let mut nums = Vec::new();
//         let mut word = String::new();
//         // println!("{}", line);
//
//         // look l->r on the line for the first nu
//         for c in line.chars() {
//             word.push(c);
//             // check if one of the map keys is in the word
//             for k in map.keys() {
//                 if word.contains(k) {
//                     let n = *map.get(k).unwrap();
//                     nums.push(n);
//                     // println!("{}: {}", word, n);
//                     word.clear();
//                     break;
//                 }
//             }
//             if c.is_ascii_digit() {
//                 nums.push(c.to_digit(radix).unwrap());
//                 // println!("{}: {}", c, c);
//                 word.clear();
//             }
//         }
//         line_nums.push(nums);
//         // break;
//     }
//     // convert the nums in each line into numbers [1-99]
//     let mut sum = 0;
//     for nums in line_nums.iter() {
//         let tens = 10 * nums[0];
//         let ones = nums.last().unwrap();
//         sum += tens + ones;
//     }
//     println!("pt2: {}\n", sum);
// }

pub fn run(data: &String) {
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

        // look l->r on the line for the first nu
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
            if nums.len() == 1 {
                word.clear();
                break;
            }
        }

        // look r->l on line for last num
        for c in line.chars().rev() {
            word = c.to_string() + &word;
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

            if nums.len() == 2 {
                break;
            }
        }
        line_nums.push(nums);
        // break;
    }
    // convert the nums in each line into numbers [1-99]
    let mut sum = 0;
    for nums in line_nums.iter() {
        let tens = 10 * nums[0];
        let ones = nums.last().unwrap();
        sum += tens + ones;
    }
    println!("pt2: {}\n", sum);
}
