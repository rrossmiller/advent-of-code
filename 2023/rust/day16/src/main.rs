use std::{fs, io};

mod pt1;
mod pt2;
fn main() -> Result<(), io::Error> {
    let test = false;
    let data = read(test)?;
    let ans = pt1::run(&data, None, None);
    println!("pt1: {}", ans);
    let ans = pt2::run(&data);

    println!("pt2: {}", ans);

    Ok(())
}

fn read(test: bool) -> io::Result<String> {
    let f;
    if test {
        f = String::from(
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
    } else {
        f = fs::read_to_string("../../data/16.txt")?;
    }
    Ok(f)
}
