use std::{
    fs::File,
    io::{self, BufRead},
};

pub fn run() {
    pt1();
    pt2();
}

fn pt1() {
    let data_pth = "data/day4_pt1.txt";
    let file = File::open(data_pth).expect(format!("{data_pth} not found").as_str());
    let input_buffer = io::BufReader::new(file);

    let mut ttl = 0;
    for line in input_buffer.lines() {
        let line = line.expect("trouble reading the line");
        let line = line.split(",").collect::<Vec<&str>>();
        let (a, b) = get_assignments(line);

        if assignments_overlap(a, b) {
            ttl += 1;
        }
    }

    println!("pt1: {}", ttl);
}

fn pt2() {
    let data_pth = "data/day4_pt1.txt";
    let file = File::open(data_pth).expect(format!("{data_pth} not found").as_str());
    let input_buffer = io::BufReader::new(file);

    let mut ttl = 0;
    for line in input_buffer.lines() {
        let line = line.expect("trouble reading the line");
        let line = line.split(",").collect::<Vec<&str>>();
        let (a, b) = get_assignments(line);

        if assignments_overlap_at_all(a, b) {
            ttl += 1;
        }
    }

    println!("pt2: {}", ttl);
}

fn assignments_overlap(a: [i32; 2], b: [i32; 2]) -> bool {
    //             b0 a0  b1                            b0 a1 b1
    let a_in_b = (b[0] <= a[0] && a[0] <= b[1]) && (b[0] <= a[1] && a[1] <= b[1]);
    //             a0 b0 a1                             a0 b1 a1
    let b_in_a = (a[0] <= b[0] && b[0] <= a[1]) && (a[0] <= b[1] && b[1] <= a[1]);

    if a_in_b || b_in_a {
        return true;
    }
    return false;
}

fn get_assignments(line: Vec<&str>) -> ([i32; 2], [i32; 2]) {
    let a = line[0].split("-").collect::<Vec<&str>>();
    let a = [a[0].parse::<i32>().unwrap(), a[1].parse::<i32>().unwrap()];

    let b = line[1].split("-").collect::<Vec<&str>>();
    let b = [b[0].parse::<i32>().unwrap(), b[1].parse::<i32>().unwrap()];

    return (a, b);
}

fn assignments_overlap_at_all(a: [i32; 2], b: [i32; 2]) -> bool {
    //             b0 a0  b1                            b0 a1 b1
    let a_in_b = (b[0] <= a[0] && a[0] <= b[1]) || (b[0] <= a[1] && a[1] <= b[1]);
    //             a0 b0 a1                             a0 b1 a1
    let b_in_a = (a[0] <= b[0] && b[0] <= a[1]) || (a[0] <= b[1] && b[1] <= a[1]);

    if a_in_b || b_in_a {
        return true;
    }
    return false;
}
#[cfg(test)]
mod tests {
    use super::assignments_overlap;

    #[test]
    fn test_make_alpha_array() {
        let a = [7, 7];
        let b = [8, 70];
        let x = assignments_overlap(a, b);
        assert!(!x);

        let a = [7, 7];
        let b = [7, 70];
        let x = assignments_overlap(a, b);
        assert!(x);
    }
}
