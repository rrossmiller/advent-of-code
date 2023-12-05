use crate::Mapping;
use indicatif::{ProgressBar, ProgressStyle};

pub fn run(seeds: &Vec<i64>, mappings: &Vec<Vec<Mapping>>) -> (Vec<i64>, i64) {
    let bar = ProgressBar::new((seeds.len() / 2) as u64);
    let spinner_style = ProgressStyle::with_template(
        "{pos}/{len}({percent}%){wide_bar}{elapsed}/{duration}({eta})",
    )
    .unwrap();
    bar.set_style(spinner_style);

    let mut all_seeds = vec![];
    for (n, _) in seeds.iter().enumerate().skip(1).step_by(2) {
        bar.inc(1);
        // bar.set_message(bar.duration().as_secs());
        let start = seeds[n - 1];
        let finish = seeds[n - 1] + seeds[n];
        for i in start..finish {
            all_seeds.push(i);
        }
    }

    let bar = ProgressBar::new(all_seeds.len() as u64);
    let spinner_style = ProgressStyle::with_template(
        "{pos}/{len}({percent}%){wide_bar}{elapsed}/{duration}({eta})",
    )
    .unwrap();
    bar.set_style(spinner_style);

    // bellow is same as p1
    let ans: Vec<i64> = all_seeds
        .iter()
        .map(|s| {
            bar.inc(1);
            let mut x = *s;
            for m in mappings.iter() {
                x = apply_mapping(&x, m);
            }
            x
        })
        .collect();
    let min = *ans.iter().min().expect("this shouldn't happen");
    (ans, min)
}

/// run the seed through each mapping
fn apply_mapping(seed: &i64, mapping: &Vec<Mapping>) -> i64 {
    let mut ans = *seed;
    let mut applied = false;

    for map in mapping.iter() {
        if applied {
            break;
        }
        if seed >= &(map.src_start) && seed <= &(map.src_start + map.rng_len) {
            ans = ans - map.src_start + map.dest_start;
            applied = true;
        }
    }
    ans
}

#[cfg(test)]
mod tests {
    use crate::test_data;

    use super::*;

    #[test]
    fn test_p2() {
        let (seeds, data) = test_data();
        let (_, min_res) = run(&seeds, &data);
        let min = 46;

        assert_eq!(min, min_res);
    }
}
