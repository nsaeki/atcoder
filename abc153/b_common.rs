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
    let h = v[0];
    let _n = v[1];
    let a = read_line();

    let sum = a.iter().fold(0, |sum, x| sum + x);

    if sum >= h {
        println!("Yes");
    } else {
        println!("No");
    }
}