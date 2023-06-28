use std::{env, process::exit};
mod day3;
mod day4;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        println!("Add the number day to run");
        exit(1);
    }

    let day = args[1].clone();
    if day == "3" {
        day3::run();
    } else if day == "4" {
        day4::run();
    }
}
