use crate::Mapping;

pub fn run(seeds: &Vec<i64>, mappings: &Vec<Vec<Mapping>>) -> (Vec<i64>, i64) {
    let ans: Vec<i64> = seeds
        .iter()
        .map(|s| {
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
    fn test_p1() {
        let (seeds, data) = test_data();
        let (res, min_res) = run(&seeds, &data);
        let ans = vec![82, 43, 86, 35];
        let min = 35;

        assert_eq!(min, min_res);
        res.into_iter()
            .enumerate()
            .for_each(|(i, r)| assert_eq!(ans[i], r));
    }
}
