use std::iter::Iterator;

struct DPad {
    i: usize,
    j: usize,
    m: usize,
    n: usize,
    c: usize,
}

impl Iterator for DPad {
    type Item = (usize, usize);
    fn next(&mut self) -> Option<Self::Item> {
        if self.c <= 0 && self.i > 0 {
            self.c = 1;
            Some((self.i - 1, self.j))
        } else if self.c <= 1 && self.i < self.m - 1 {
            self.c = 2;
            Some((self.i + 1, self.j))
        } else if self.c <= 2 && self.j > 0 {
            self.c = 3;
            Some((self.i, self.j - 1))
        } else if self.c <= 3 && self.j < self.n - 1 {
            self.c = 4;
            Some((self.i, self.j + 1))
        } else {
            None
        }
    }
}

#[macro_export]
macro_rules! dpad {
    ( $i:expr, $j:expr ) => {
        DPad {
            i: $i,
            j: $j,
            m: usize::MAX,
            n: usize::MAX,
            c: 0,
        }
    };
    ( $i:expr, $j:expr, $n:expr ) => {
        DPad {
            i: $i,
            j: $j,
            m: $n,
            n: $n,
            c: 0,
        }
    };
    ( $i:expr, $j:expr, $m:expr, $n:expr ) => {
        DPad {
            i: $i,
            j: $j,
            m: $m,
            n: $n,
            c: 0,
        }
    };
}

fn main() {
    assert_eq!(
        dpad!(0, 9).collect::<Vec<_>>(),
        vec![(1, 9), (0, 8), (0, 10)]
    );
    assert_eq!(
        dpad!(0, 9, 10).collect::<Vec<_>>(),
        vec![(1, 9), (0, 8)]
    );
    assert_eq!(
        dpad!(0, 9, 100, 10).collect::<Vec<_>>(),
        vec![(1, 9), (0, 8)]
    );
    assert_eq!(
        dpad!(10, 10, 100, 100).collect::<Vec<_>>(),
        vec![(9, 10), (11, 10), (10, 9), (10, 11)]
    );
}
