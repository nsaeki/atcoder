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
    let n = v[0] as usize;
    let k = v[1] as usize;
    let mut h = read_line();

    if n <= k {
        println!("0");
        return;
    }

    h.sort();
    let c = h.iter().take(n-k).fold(0, |c, x| c + *x as u64);
    println!("{}", c);
}