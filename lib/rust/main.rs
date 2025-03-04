use std::io;
use std::str::FromStr;

#[macro_export]
macro_rules! read_tuple {
    ( $( $t:ty ),+ ) => {{
        use std::io::{self, BufRead};

        let stdin = io::stdin();
        let mut iterator = stdin.lock().lines();
        let line = iterator
            .next()
            .expect("input error")
            .expect("input error");

        let mut iter = line.split_whitespace();

        (
            $(
                iter.next()
                    .expect("few input")
                    .parse::<$t>()
                    .expect("parse error")
            ),+
        )
    }};
}

#[allow(dead_code)]
fn read_line() -> String {
    let mut buf = String::new();
    io::stdin().read_line(&mut buf).expect("failed to read");
    buf.trim().to_string()
}

#[allow(dead_code)]
fn parse_line<T: FromStr>() -> Vec<T> {
    let mut buf = String::new();
    io::stdin().read_line(&mut buf).expect("failed to read");
    buf.split_whitespace()
        .map(|v| {
            v.parse::<T>()
                .map_err(|_| "parse error")
                .expect("failed to parse")
        })
        .collect()
}

fn main() {
    // let t = read_tuple!(i32, char, usize);
    // let line = parse_line::<i32>();
    // let n = line[0];
    // let s = read_line();
}
