use std::collections::HashMap;

pub fn run(dat: &str) {
    let mut ans = 0;

    // loop through games
    for line in dat.lines() {
        let spl: Vec<&str> = line.split(":").collect();

        let mut mins: HashMap<&str, i32> = HashMap::new();

        // loop through sets
        for set in spl[1].split(";").map(|x| x.trim()) {
            // loop through pics
            for c in set.split(",").map(|x| x.trim()) {
                // check tat the pick is relevant
                let spl: Vec<_> = c.split_whitespace().collect();
                let color = spl[1];
                let n: i32 = spl[0].parse().unwrap();
                // if *bag.get(color).unwrap() < n {
                //     valid_game = false;
                //     break;
                // }

                // update min
                if !mins.contains_key(color) || *mins.get(color).unwrap() < n {
                    mins.insert(color, n);
                }
            }
        }
            ans += mins.values().product::<i32>();
    }
    println!("pt2: {}", ans);
}
