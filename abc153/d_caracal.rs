use std::io::stdin;

fn read_line() -> Vec<i64> {
    let mut a_str = String::new();
    stdin().read_line(&mut a_str).expect("read error");
    a_str.split_whitespace()
        .map(|x| x.parse::<i64>().expect("parse error"))
        .collect::<Vec<i64>>()
}

fn main() {
    let mut h = read_line()[0];
    let mut c = 0;
    while h > 0 {
        h /= 2;
        c += 1;
    }

    let x: i64 = 2;
    println!("{}", x.pow(c) - 1);
}