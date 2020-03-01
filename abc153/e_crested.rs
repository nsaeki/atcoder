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
    let h = v[0] as usize;
    let n = v[1] as usize;

    let mut m = Vec::<(i32, i32)>::new();
    for _i in 0..n {
        let v = read_line();
        m.push((v[0], v[1]));
    }

    let a_max = match m.iter().max_by(|x, y| x.0.cmp(&y.0)) {
        Some(x) => x.0,
        _ => 0,
    };
    let a_max = a_max as usize;

    let mut dp = vec![std::i32::MAX; h+a_max];

    let d_max = h + a_max as usize;

    dp[0] = 0;
    for i in 0..h {
        for x in m.iter() {
            let ni = i + x.0 as usize;
            let mp = if dp[i] == std::i32::MAX { dp[i] } else { dp[i] + x.1 };
            if mp < dp[ni] {
                dp[ni] = mp;
            }
        }
    }

    let na = 0;
    let m_min = match dp[h..d_max].iter().min() {
        Some(x) => x,
        _ => &na,
    };
    println!("{}", m_min);
}