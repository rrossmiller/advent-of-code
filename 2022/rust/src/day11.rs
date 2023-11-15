use std::{collections::VecDeque, fs, usize};

#[derive(Debug, Clone)]
enum OpVal {
    Old,
    Val(i32),
}

#[derive(Debug, Clone)]
struct Monkey {
    items: VecDeque<i32>, //Vec<i32>,
    operation: String,
    op_val: OpVal,
    test_val: i32,
    pass_to: Vec<usize>,
    inspections: i32,
}

impl Monkey {
    fn operate(&self, a: i32) -> i32 {
        let op_val = match self.op_val {
            OpVal::Old => a,
            OpVal::Val(x) => x,
        };
        match self.operation.as_str() {
            "+" => a + op_val,
            "-" => a - op_val,
            "*" => a * op_val,
            "/" => a / op_val,
            _ => 0,
        }
    }
    fn test_and_throw(&mut self, a: i32) -> (usize, i32) {
        self.inspections += 1;
        let a = self.operate(a) / 3;

        if a % self.test_val == 0 {
            // tmonkey.items.push_back(a);
            return (0, a);
        } else {
            // tmonkey.items.push_back(a);
            return (0, a);
        }
    }
}

fn get_data() -> Vec<Monkey> {
    let data_path = "../python/test.txt";
    let data = fs::read_to_string(data_path).expect(format!("{data_path} not found").as_str());

    let x: Vec<&str> = data.split("\n\n").collect();
    let x: Vec<Vec<&str>> = x.iter().map(|x| x.split("\n").collect()).collect();
    let mut monkeys = vec![];
    for i in 0..x.len() {
        let m = x.get(i).unwrap();
        // Items
        let l: &str = m
            .get(1)
            .unwrap()
            .split(":")
            .collect::<Vec<&str>>()
            .get(1)
            .unwrap()
            .trim();
        let items: VecDeque<i32> = l
            .split(",")
            .collect::<Vec<&str>>()
            .iter()
            .map(|x| x.trim().parse().unwrap())
            .collect();

        // Operation
        let l: &str = m
            .get(2)
            .unwrap()
            .split(":")
            .collect::<Vec<&str>>()
            .get(1)
            .unwrap()
            .trim();
        let op: Vec<&str> = l.split(" ").collect();
        let operation = op.get(op.len() - 2).unwrap().to_string();

        let op_val = op.get(op.len() - 1).unwrap();
        let op_val = match *op_val {
            "old" => OpVal::Old,
            _ => OpVal::Val(op_val.parse().unwrap()),
        };

        // Test
        let test_val = m.get(3).unwrap().split(" ").collect::<Vec<&str>>();
        let test_val = test_val.get(test_val.len() - 1).unwrap().parse().unwrap();
        let t = m.get(4).unwrap().split(" ").collect::<Vec<&str>>();
        let f = m.get(5).unwrap().split(" ").collect::<Vec<&str>>();
        let pass_to: Vec<usize> = vec![
            t.get(t.len() - 1).unwrap().parse().unwrap(),
            f.get(f.len() - 1).unwrap().parse().unwrap(),
        ];

        let inspections = 0;
        let monkey = Monkey {
            items,
            operation,
            op_val,
            test_val,
            pass_to,
            inspections,
        };
        println!("{:?}", monkey);
        monkeys.push(monkey);
    }
    return monkeys;
}

pub fn run() {
    let mut monkeys = get_data();
    println!("");

    for _ in 0..20 {
        for m in monkeys.iter_mut() {
            while m.items.len() > 0 {
                let a = m.items.pop_front().unwrap();
                let (i, a) = m.test_and_throw(a);
                monkeys.get(i).unwrap().items.push_back(a);
            }
        }
    }
    println!("");
    let mut ans = vec![];
    for m in monkeys.iter() {
        ans.push(m.inspections);
    }
    ans.sort();
    ans.reverse();
    println!("{:?}", ans);

    let a = ans.get(0).unwrap() * ans.get(1).unwrap();
    println!("{}", a);
}
