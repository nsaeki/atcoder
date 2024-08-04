// WIP
// BIT (A Binary indexed tree or Fenwick tree)
struct BIT {
    n: usize,
    data: Vec<i32>,
}

impl BIT {
    fn new(n: usize) -> Self {
        Self {
            n: n,
            data: vec![0; n + 1],
        }
    }

    fn add(&mut self, mut i: usize, v: i32) {
        i += 1;
        while i <= self.n {
            self.data[i] += v;
            i += i & i.wrapping_neg();
        }
    }

    fn sum(&self, mut i: usize) -> i32 {
        let mut v = 0;
        i += 1;
        while i > 0 {
            v += self.data[i];
            i -= i & i.wrapping_neg();
        }
        v
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sum() {
        let mut bit = BIT::new(3);
        for i in 0..3 {
            bit.add(i, i as i32);
        }
        assert!(bit.sum(3) == 6);
    }
}
