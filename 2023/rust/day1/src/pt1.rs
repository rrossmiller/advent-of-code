pub fn run(data: &String) {
    let radix = 10;
    // get nums in each line
    let mut line_nums = Vec::new();
    for line in data.lines() {
        let mut nums = Vec::new();
        for c in line.chars() {
            if c.is_numeric() {
                nums.push(c.to_digit(radix).unwrap())
            }
        }
        line_nums.push(nums);
    }

    // convert the nums in each line into numbers [1-99]
    let mut sum = 0;
    for nums in line_nums.iter() {
        let tens = 10 * nums[0];
        let ones = nums.last().unwrap();
        sum += tens + ones;
    }
    println!("pt1: {}", sum);
}
// fn pt1(data: &String) {
//     let radix = 10;
//     // get nums in each line
//     let mut line_nums = Vec::new();
//     for line in data.lines() {
//         let mut nums = Vec::new();
//         for c in line.chars() {
//             if c.is_numeric() {
//                 nums.push(c.to_digit(radix).unwrap());
//                 break;
//             }
//         }
//         for c in line.chars().rev() {
//             if c.is_numeric() {
//                 nums.push(c.to_digit(radix).unwrap());
//                 break;
//             }
//         }
//         line_nums.push(nums);
//     }
//
//     // convert the nums in each line into numbers [1-99]
//     let mut sum = 0;
//     for nums in line_nums.iter() {
//         let tens = 10 * nums[0];
//         let ones = nums.last().unwrap();
//         sum += tens + ones;
//     }
//     println!("pt1: {}\n", sum);
// }
