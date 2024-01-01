use std::{fs, io};

mod pt1;
fn main() -> Result<(), io::Error> {
    let test = false;
    let data = read(test)?;
    pt1::run(&data);

    Ok(())
}

fn read(test: bool) -> io::Result<String> {
    let f;
    if test {
        f = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7".to_string();
    } else {
        f = fs::read_to_string("../../data/15.txt")?;
    }
    Ok(f)
}
