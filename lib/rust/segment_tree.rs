#[derive(Debug)]
struct SegmentTree {
    n: usize,
    data: Vec<i32>,
}

#[allow(dead_code)]
impl SegmentTree {
    fn new(n: usize) -> Self {
        Self {
            n: n,
            data: vec![0; n * 2],
        }
    }

    fn len(&self) -> usize {
        self.n
    }

    fn modify(&mut self, mut i: usize, v: i32) {
        i += self.n;
        self.data[i] = v;
        while i > 0 {
            self.data[i >> 1] = self.data[i] + self.data[i ^ 1];
            i >>= 1;
        }
    }

    fn query(&self, mut left: usize, mut right: usize) -> i32 {
        let mut res = 0;
        left += self.n;
        right += self.n;
        while left < right {
            if left & 1 == 1 {
                res += self.data[left];
                left += 1;
            }
            if right & 1 == 1 {
                right -= 1;
                res += self.data[right];
            }
            left >>= 1;
            right >>= 1;
        }
        res
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn query() {
        let mut s = SegmentTree::new(10);
        for i in 0..10 {
            s.modify(i, i as i32 + 1);
        }
        assert!(s.query(0, 10) == 55);
        assert!(s.query(4, 8) == 26);
    }

    #[test]
    fn len() {
        let s = SegmentTree::new(10);
        assert!(s.len() == 10);
    }
}
