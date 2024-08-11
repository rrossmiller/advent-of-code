use std::collections::HashMap;

pub fn run(dat: &str) {
    let mut ans = 0;
    let bag = HashMap::from([("red", 12), ("blue", 14), ("green", 13)]);

    // loop through games
    for line in dat.lines() {
        let spl: Vec<&str> = line.split(":").collect();

        let game_id: i32 = spl[0].split_whitespace().nth(1).unwrap().parse().unwrap();
        let mut valid_game = true;
        // loop through sets
        for set in spl[1].split(";").map(|x| x.trim()) {
            // loop through pics
            for c in set.split(",").map(|x| x.trim()) {
                // check tat the pick is relevant
                let spl: Vec<_> = c.split_whitespace().collect();
                let color = spl[1];
                let n: i32 = spl[0].parse().unwrap();
                if *bag.get(color).unwrap() < n {
                    valid_game = false;
                    break;
                }
            }
        }
        if valid_game {
            ans += game_id;
        }
    }
    println!("pt1: {}", ans);
}
