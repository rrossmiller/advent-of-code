use std::{fs, io};
mod pt1;
mod pt2;
fn main() -> std::io::Result<()> {
    let data = read()?;
    // let data = test_data(1);
    pt1::run(&data);

    // let data = test_data(2);
    pt2::run(&data);

    Ok(())
}

fn read() -> io::Result<String> {
    let f = fs::read_to_string("../../data/1.txt")?;
    Ok(f)
}

// fn test_data(pt: i32) -> String {
//     match pt {
//         1 => "1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// "
//         .to_string(),
//         2 => "two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen"
//             .to_string(),
//         _ => "".to_string(),
//     }
// }
