use crate::pt1::Coord;

const DIRECTIONS: [(i32, i32); 8] = [
    (-1, -1), // UL
    (-1, 0),  //U
    (-1, 1),  //UR
    (0, -1),  //L
    (0, 1),   // R
    (1, -1),  //DL
    (1, 0),   //D
    (1, 1),   //DR
];

pub fn run(dat: &str, i: usize, j: usize, nums: &mut Vec<i32>, visited: &mut Vec<Coord>) {
    let dat = dat.lines().collect::<Vec<&str>>();
    let cols = dat[0].len();
    let mut found_nums = vec![];

    // bfs to look around the symbol
    for d in DIRECTIONS {
        let row = (i as i32 + d.0) as usize;
        let col = (j as i32 + d.1) as usize;
        let coord = Coord(row, col);
        let line = dat.get(row);
        if line.is_none() {
            // handles addition overflow
            continue;
        }
        let line = line.unwrap();

        // if the char at the coord is a digit
        if row < dat.len()
                    && col < cols
                    && line.chars().nth(col).unwrap().is_digit(10) // char is digit
        && !visited.contains(&coord)
        // idx hasn't been checked already
        {
            // visit the coord
            visited.push(coord);

            // find the rest of the number
            let mut l = if col > 0 { col - 1 } else { 0 };
            let mut r = col + 1;
            // find leftmost limit
            while l > 0 && line.chars().nth(l).unwrap().is_digit(10) {
                visited.push(Coord(row, l));
                l -= 1;
            }
            if !line.chars().nth(l).unwrap().is_digit(10) {
                l += 1;
            }
            // find rightmost limit
            while r < line.len() && line.chars().nth(r).unwrap().is_digit(10) {
                visited.push(Coord(row, r));
                r += 1;
            }
            let n: i32 = line[l..r].parse().unwrap();
            found_nums.push(n);
        }
    }

    if found_nums.len() == 2 {
        nums.push(found_nums[0] * found_nums[1])
    }
}
