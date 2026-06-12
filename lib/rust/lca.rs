#[allow(dead_code)]
struct BinaryLiftingLCA {
    up: Vec<Vec<usize>>,
    depth: Vec<usize>,
    log: usize,
}

#[allow(dead_code)]
impl BinaryLiftingLCA {
    fn new(root: usize, graph: &Vec<Vec<usize>>) -> Self {
        let n = graph.len();
        let log = usize::BITS as usize - n.leading_zeros() as usize;
        let mut up = vec![vec![usize::MAX; log]; n];
        let mut depth = vec![0; n];

        Self::dfs(root, usize::MAX, 0, graph, &mut up, &mut depth);
        for k in 1..log {
            for v in 0..up.len() {
                let p = up[v][k - 1];
                if p != usize::MAX {
                    up[v][k] = up[p][k - 1];
                }
            }
        }

        Self { up, depth, log }
    }

    fn dfs(
        i: usize,
        parent: usize,
        d: usize,
        graph: &Vec<Vec<usize>>,
        up: &mut Vec<Vec<usize>>,
        depth: &mut Vec<usize>,
    ) {
        up[i][0] = parent;
        depth[i] = d;

        for &j in &graph[i] {
            if j == parent {
                continue;
            }
            Self::dfs(j, i, d + 1, graph, up, depth);
        }
    }

    fn get(&self, mut u: usize, mut v: usize) -> usize {
        if self.depth[u] < self.depth[v] {
            std::mem::swap(&mut u, &mut v);
        }

        let diff = self.depth[u] - self.depth[v];

        for k in 0..self.log {
            if (diff >> k) & 1 == 1 {
                u = self.up[u][k];
            }
        }

        if u == v {
            return u;
        }
        for k in (0..self.log).rev() {
            if self.up[u][k] != self.up[v][k] {
                u = self.up[u][k];
                v = self.up[v][k];
            }
        }

        self.up[u][0]
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn lower() {
        let graph = vec![vec![], vec![2, 3], vec![], vec![4, 5], vec![], vec![]];
        let lca = BinaryLiftingLCA::new(1, &graph);
        assert_eq!(lca.get(2, 4), 1)
    }

    #[test]
    fn deeper() {
        let graph = vec![vec![], vec![2, 3], vec![], vec![4, 5], vec![], vec![]];
        let lca = BinaryLiftingLCA::new(1, &graph);
        assert_eq!(lca.get(5, 2), 1)
    }

    #[test]
    fn in_the_path() {
        let graph = vec![vec![], vec![2, 3], vec![], vec![4, 5], vec![], vec![]];
        let lca = BinaryLiftingLCA::new(1, &graph);
        assert_eq!(lca.get(5, 3), 3)
    }

    #[test]
    fn same_depth() {
        let graph = vec![vec![], vec![2, 3], vec![], vec![4, 5], vec![], vec![]];
        let lca = BinaryLiftingLCA::new(1, &graph);
        assert_eq!(lca.get(4, 5), 3)
    }

    #[test]
    fn same_node() {
        let graph = vec![vec![], vec![2, 3], vec![], vec![4, 5], vec![], vec![]];
        let lca = BinaryLiftingLCA::new(1, &graph);
        assert_eq!(lca.get(1, 1), 1)
    }
}

