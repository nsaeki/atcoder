trait Monoid {
    type S;

    fn identity() -> Self::S;
    fn op(a: &Self::S, b: &Self::S) -> Self::S;
}

#[derive(Debug)]
struct SegmentTree<M: Monoid> {
    n: usize,
    data: Vec<M::S>,
}

#[allow(dead_code)]
impl<M> SegmentTree<M>
where
    M: Monoid,
{
    fn new(n: usize) -> Self {
        Self {
            n: n,
            data: (0..n * 2).map(|_| M::identity()).collect(),
        }
    }

    fn len(&self) -> usize {
        self.n
    }

    fn update(&mut self, mut i: usize, v: M::S) {
        i += self.n;
        self.data[i] = v;
        i >>= 1;
        while i > 0 {
            self.data[i] = M::op(&self.data[i << 1], &self.data[i << 1 | 1]);
            i >>= 1;
        }
    }

    fn query(&self, mut left: usize, mut right: usize) -> M::S {
        let mut res_l = M::identity();
        let mut res_r = M::identity();
        left += self.n;
        right += self.n;
        while left < right {
            if left & 1 == 1 {
                res_l = M::op(&res_l, &self.data[left]);
                left += 1;
            }
            if right & 1 == 1 {
                right -= 1;
                res_r = M::op(&self.data[right], &res_r);
            }
            left >>= 1;
            right >>= 1;
        }

        M::op(&res_l, &res_r)
    }
}

#[allow(dead_code)]
struct Sum;

impl Monoid for Sum {
    type S = i32;
    fn identity() -> Self::S {
        0
    }
    fn op(a: &Self::S, b: &Self::S) -> Self::S {
        *a + *b
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn query_sum() {
        let mut s = SegmentTree::<Sum>::new(10);
        for i in 0..10 {
            s.update(i, i as i32 + 1);
        }
        assert_eq!(s.query(0, 10), 55);
        assert_eq!(s.query(4, 8), 26);
    }

    #[test]
    fn query_max() {
        struct Max;

        impl Monoid for Max {
            type S = i32;

            fn identity() -> Self::S {
                i32::MIN
            }

            fn op(a: &Self::S, b: &Self::S) -> Self::S {
                *a.max(b)
            }
        }

        let mut s = SegmentTree::<Max>::new(10);
        for i in 0..10 {
            s.update(i, i as i32 + 1);
        }
        assert_eq!(s.query(0, 10), 10);
        assert_eq!(s.query(4, 8), 8);
    }

    #[test]
    fn query_concat() {
        struct Concat;

        impl Monoid for Concat {
            type S = String;

            fn identity() -> Self::S {
                String::new()
            }

            fn op(a: &Self::S, b: &Self::S) -> Self::S {
                format!("{a}{b}")
            }
        }

        let mut s = SegmentTree::<Concat>::new(5);
        for i in 0..5 {
            s.update(i, ((b'a' + i as u8) as char).into());
        }
        assert_eq!(s.query(0, 5), "abcde");
        assert_eq!(s.query(1, 4), "bcd");
    }

    #[test]
    fn len() {
        let s = SegmentTree::<Sum>::new(10);
        assert_eq!(s.len(), 10);
    }
}
