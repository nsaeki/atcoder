fn make_multi_choose(n: usize, k: usize, m: i64) -> impl Fn(usize, usize) -> i64 {
    let mut mc = vec![vec![1; k + 1]; n + 1];
    mc[0].fill(0);
    for i in 2..=n {
        for j in 1..=k {
            mc[i][j] = (mc[i - 1][j] + mc[i][j - 1]) % m;
        }
    }
    return move |n, k| mc[n][k]
}

fn main() {
    let n = 5;
    let k = 6;
    let multi_choose = make_multi_choose(n, k, 1_000_000_007);
    assert_eq!(multi_choose(0, 3), 0);
    assert_eq!(multi_choose(2, 0), 1);
    assert_eq!(multi_choose(1, 3), 1);
    assert_eq!(multi_choose(5, 3), 35);
    assert_eq!(multi_choose(5, 6), 210);
}
