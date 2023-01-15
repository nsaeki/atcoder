#[derive(Debug)]
struct UnionFind {
    root: Vec<usize>,
}

impl UnionFind {
    fn new(n: usize) -> Self {
        UnionFind {
            root: (0..n).map(|i| i).collect(),
        }
    }

    fn unite(&mut self, mut x: usize, mut y: usize) {
        x = self.root(x);
        y = self.root(y);
        if x == y {
            return;
        }
        self.root[x] = y;
    }

    fn root(&mut self, x: usize) -> usize {
        if self.root[x] == x {
            return x;
        }
        self.root[x] = self.root(self.root[x]);
        self.root[x]
    }
}

