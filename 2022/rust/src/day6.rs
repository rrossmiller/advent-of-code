use std::{collections::HashSet, fs};

pub fn run() {
    //pt1
    decode(4);
    //pt2
    decode(14);
}

fn decode(n: usize) {
    let data_path = "data/day6.txt";
    let data = fs::read_to_string(data_path).expect(format!("{data_path} not found").as_str());

    let chars: Vec<char> = data.chars().collect();
    for i in n..chars.len() {
        if is_different(&chars[i - n..i]) {
            println!("{}", i);
            break;
        }
    }
}

fn is_different(buf: &[char]) -> bool {
    let mut uniq = HashSet::new();
    buf.into_iter().all(move |x| uniq.insert(x))
}
