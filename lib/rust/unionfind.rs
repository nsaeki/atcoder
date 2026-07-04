#[derive(Debug)]
struct UnionFind {
    root: Vec<usize>,
    size: Vec<usize>,
}

#[allow(dead_code)]
impl UnionFind {
    fn new(n: usize) -> Self {
        UnionFind {
            root: (0..n).collect(),
            size: vec![1; n],
        }
    }

    fn union(&mut self, mut x: usize, mut y: usize) -> bool {
        x = self.root(x);
        y = self.root(y);
        if x == y {
            return false;
        }
        if self.size[x] < self.size[y] {
            std::mem::swap(&mut x, &mut y);
        }
        self.root[y] = x;
        self.size[x] += self.size[y];
        true
    }

    fn root(&mut self, x: usize) -> usize {
        if self.root[x] == x {
            x
        } else {
            self.root[x] = self.root(self.root[x]);
            self.root[x]
        }
    }

    fn size(&mut self, x: usize) -> usize {
        let r = self.root(x);
        self.size[r]
    }
}
