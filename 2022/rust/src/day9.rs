use std::collections::HashSet;

use std::fs;
#[derive(Debug)]
enum Direction {
    Up(i32),
    Down(i32),
    Left(i32),
    Right(i32),
}

#[derive(Hash, Eq, PartialEq, Debug)]
struct Point {
    x: i32,
    y: i32,
}

pub fn run() {
    let dirs = get_data();
    // let dirs = get_test_data();
    let mut h = Point { x: 0, y: 0 };
    let mut t = Point { x: 0, y: 0 };
    let mut visited = HashSet::from([Point { x: 0, y: 0 }]);
    for d in dirs.iter() {
        match d {
            Direction::Up(mut i) => {
                while i != 0 {
                    h.y += 1;
                    if (h.y - t.y).abs() == 2 {
                        // check if t needs to move diagonally
                        if h.x != t.x {
                            t.x = h.x;
                        }
                        t.y += 1;
                        visited.insert(Point { x: t.x, y: t.y });
                    }
                    i -= 1;
                }
            }
            Direction::Down(mut i) => {
                while i != 0 {
                    h.y -= 1;
                    if (h.y - t.y).abs() == 2 {
                        // check if t needs to move diagonally
                        if h.x != t.x {
                            t.x = h.x;
                        }
                        t.y -= 1;
                        visited.insert(Point { x: t.x, y: t.y });
                    }
                    i -= 1;
                }
            }
            Direction::Right(mut i) => {
                while i != 0 {
                    h.x += 1;
                    if (h.x - t.x).abs() == 2 {
                        // check if t needs to move diagonally
                        if h.y != t.y {
                            t.y = h.y;
                        }
                        t.x += 1;
                        visited.insert(Point { x: t.x, y: t.y });
                    }
                    i -= 1;
                }
            }
            Direction::Left(mut i) => {
                while i != 0 {
                    h.x -= 1;
                    if (h.x - t.x).abs() == 2 {
                        // check if t needs to move diagonally
                        if h.y != t.y {
                            t.y = h.y;
                        }
                        t.x -= 1;
                        visited.insert(Point { x: t.x, y: t.y });
                    }
                    i -= 1;
                }
            }
        }
    }

    println!("{:?}", visited.len());
}

fn get_test_data() -> Vec<Direction> {
    let data = "R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2";
    let mut dirs = vec![];
    for l in data.lines() {
        let spl: Vec<&str> = l.split(" ").collect();
        let l = spl[0];
        let i: i32 = spl[1].parse().expect("issue parsing int");
        match l {
            "U" => dirs.push(Direction::Up(i)),
            "D" => dirs.push(Direction::Down(i)),
            "R" => dirs.push(Direction::Right(i)),
            "L" => dirs.push(Direction::Left(i)),
            _ => unreachable!(),
        }
    }

    dirs
}
fn get_data() -> Vec<Direction> {
    let data_path = "../data/9.txt";
    let data = fs::read_to_string(data_path).expect(format!("{data_path} not found").as_str());
    // let file = fs::File::open(data_path).unwrap();
    // let reader = BufReader::new(file);

    let mut dirs = vec![];
    // for l in reader.lines() {
    for l in data.lines() {
        // let l = l.unwrap
        let spl: Vec<&str> = l.split(" ").collect();
        let l = spl[0];
        let i: i32 = spl[1].parse().expect("issue parsing int");
        match l {
            "U" => dirs.push(Direction::Up(i)),
            "D" => dirs.push(Direction::Down(i)),
            "R" => dirs.push(Direction::Right(i)),
            "L" => dirs.push(Direction::Left(i)),
            _ => unreachable!(),
        }
    }

    dirs
}
