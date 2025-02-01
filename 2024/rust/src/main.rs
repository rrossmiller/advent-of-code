use std::{fs, path::PathBuf};
use clap::Parser;

use glob::glob;
mod day8;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    /// Name of the person to greet
    #[arg(short, long)]
    name: String,

    /// Number of times to greet
    #[arg(short, long, default_value_t = 1)]
    count: u8,
}
fn main() {
    let test = true;
    let day = get_latest_day();
    println!("Running day {}", day);

    let data = get_data(day + ".txt", test);
    for l in data.iter() {
        println!("{}", l);
    }
}

fn get_latest_day() -> String {
    let mut matches = glob("../data/*.txt")
        .expect("Failed to read glob pattern")
        .map(|e| e.unwrap())
        .collect::<Vec<PathBuf>>();
    matches.sort();
    let day_num = matches
        .last()
        .unwrap()
        .file_name()
        .unwrap()
        .to_str()
        .unwrap()
        .split(".")
        .nth(0)
        .unwrap();

    day_num.to_string()
}
fn get_data(day: String, test: bool) -> Vec<String> {
    let mut p = String::from("../data/");
    p.push_str(&day);

    let mut test_line = 0;
    let content = fs::read_to_string(p).unwrap();
    for (i, l) in content.lines().enumerate() {
        if l == "-----TEST-----" {
            test_line = i;
            break;
        }
    }
    let lines = content
        .lines()
        .map(|e| e.to_string())
        .collect::<Vec<String>>();
    if test {
        return lines[..test_line].to_vec();
    } else {
        return lines[test_line + 1..].to_vec();
    }
}
