pub fn run(data: &str) {
    let mut sum: u32 = 0;
    // for each entry
    data.split(",").for_each(|s| {
        let mut x = 0;
        // incr sum given the hash of 's'
        s.chars().for_each(|c| {
            if c != '\n' {
                x += c as u32;
                x = (x * 17) % 256;
            }
        });
        println!("{s}{x}");
        sum += x;
    });

    if sum <= 515840 {
        println!("too low");
    }
    println!("pt1: {}", sum);
}
