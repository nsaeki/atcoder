use std::io::stdin;

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
    for i in 0..n {
        if m[i].1 <= 0 {
            continue;
        }
        let c = (m[i].1 + a - 1) / a;
        ret += c;
        for j in i+1..n {
            if m[j].0 > m[i].0 + 2 * d as i64 {
                break;
            }
            m[j].1 -= a * c;
        }
    }
    
    println!("{}", ret);
}