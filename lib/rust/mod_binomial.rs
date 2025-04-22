fn mod_pow(mut x: i64, mut exp: i64, m: i64) -> i64 {
    let mut res = 1;
    x %= m;
    while exp > 0 {
        if exp & 1 == 1 {
            res = res * x % m;
        }
        x = x * x % m;
        exp >>= 1;
    }
    res
}

struct ModBinomial {
    n: usize,
    fact: Vec<i64>,
    inv: Vec<i64>,
    m: i64,
}

impl ModBinomial {
    fn new(n: usize, m: i64) -> Self {
        let mut fact = vec![1; n + 1];
        for i in 1..=n {
            fact[i] = fact[i - 1] * i as i64 % m;
        }
        let mut inv = vec![0; n + 1];
        inv[n] = mod_pow(fact[n], m - 2, m);
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
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn pow() {
        assert_eq!(mod_pow(3, 4, 1_000_000_007), 81);
        assert_eq!(mod_pow(2, 5, 1_000_000_007), 32);
        assert_eq!(mod_pow(3, 4, 7), 4);
    }

    #[test]
    fn permutation() {
        let f = ModBinomial::new(10, 1_000_000_007);
        assert_eq!(f.permutation(10, 3), 720);
        assert_eq!(f.permutation(5, 2), 20);

        let f = ModBinomial::new(10, 13);
        assert_eq!(f.permutation(10, 3), 5);
        assert_eq!(f.permutation(5, 2), 7);
    }

    #[test]
    fn choose() {
        let f = ModBinomial::new(10, 1_000_000_007);
        assert_eq!(f.choose(10, 3), 120);
        assert_eq!(f.choose(5, 2), 10);

        let f = ModBinomial::new(10, 13);
        assert_eq!(f.choose(10, 3), 3);
        assert_eq!(f.choose(5, 2), 10);
    }
}
