#[allow(dead_code)]
struct Combination {
    n: usize,
    fact: Vec<i64>,
    inv: Vec<i64>,
    m: i64,
}

#[allow(dead_code)]
impl Combination {
    fn new(n: usize, m: i64) -> Self {
        let mut fact = vec![1; n + 1];
        for i in 1..=n {
            fact[i] = fact[i - 1] * i as i64 % m;
        }
        let mut inv = vec![0; n + 1];
        inv[n] = Self::pow(fact[n], m - 2, m);
        for i in (1..=n).rev() {
            inv[i - 1] = inv[i] * i as i64 % m;
        }

        Self { n, fact, inv, m }
    }

    fn permutation(&self, n: usize, k: usize) -> i64 {
        self.fact[n] * self.inv[n - k] % self.m
    }

    fn choose(&self, n: usize, k: usize) -> i64 {
        self.fact[n] * self.inv[k] % self.m * self.inv[n - k] % self.m
    }

    fn pow(x: i64, exp: i64, m: i64) -> i64 {
        if exp == 0 {
            return 1;
        }
        let r = Self::pow(x, exp / 2, m);
        if exp & 1 == 0 {
            r * r % m
        } else {
            r * r % m * x % m
        }
    }
}
