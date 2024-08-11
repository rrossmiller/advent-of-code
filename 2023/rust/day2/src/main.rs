use std::{fs, io};

mod pt1;
mod pt2;
const TEST: bool = false;
fn main() -> std::io::Result<()> {
    let dat = read()?;
    pt1::run(&dat);
    pt2::run(&dat);

    Ok(())
}

fn read() -> io::Result<String> {
    return if !TEST {
        let f = fs::read_to_string("../../data/2.txt")?;
        Ok(f)
    } else {
        let rtn = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
            .to_string();
        Ok(rtn)
    };
}
