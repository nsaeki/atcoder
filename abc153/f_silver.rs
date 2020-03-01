use std::io::stdin;
use std::collections::VecDeque;

fn read_line() -> Vec<i64> {
    let mut a_str = String::new();
    stdin().read_line(&mut a_str).expect("read error");
    a_str.split_whitespace()
        .map(|x| x.parse::<i64>().expect("parse error"))
        .collect::<Vec<i64>>()
}

fn main() {
    let v = read_line();
    let n = v[0] as usize;
    let d = v[1] as usize;
    let a = v[2];

    let mut m = Vec::<(i64, i64)>::new();
    for _i in 0..n as usize {
        let v = read_line();
        m.push((v[0], v[1]));
    }

    m.sort_by(|x, y| x.0.cmp(&y.0));

    let mut ret = 0;
    let mut total = 0;
    let mut queue = VecDeque::<(i64, i64)>::new();

    for i in 0..n {
        while queue.len() > 0 {
            let x = queue[0];
            if x.0 < m[i].0 {
                total -= queue[0].1;
                queue.pop_front();
            } else {
                break;
            }
        }

        m[i].1 -= total;

        if m[i].1 > 0 {
            let c = (m[i].1 + a - 1) / a;
            ret += c;
            total += c * a;
            queue.push_back((m[i].0 + 2 * d as i64, c * a));
        }
    }
    
    println!("{}", ret);
}