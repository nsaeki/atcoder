use std::io::stdin;

fn read_line() -> Vec<i32> {
    let mut a_str = String::new();
    stdin().read_line(&mut a_str).expect("read error");
    a_str.split_whitespace()
        .map(|x| x.parse::<i32>().expect("parse error"))
        .collect::<Vec<i32>>()
}

fn main() {
    let v = read_line();
    let mut h = v[0];
    let a = v[1];
    let mut c = 0;
    while h > 0 {
        c += 1;
        h -= a;
    }

    println!("{}", c);
}