pub fn build_z_array(s: &[u8]) -> Vec<usize> {
    let n = s.len();
    let mut z = vec![0usize; n];
    if n == 0 {
        return z;
    }

    let (mut l, mut r) = (0usize, 0usize); // [l, r)

    for i in 1..n {
        if i < r {
            z[i] = z[i - l].min(r - i);
        }
        while i + z[i] < n && s[z[i]] == s[i + z[i]] {
            z[i] += 1;
        }
        if i + z[i] > r {
            l = i;
            r = i + z[i];
        }
    }

    z
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_build_z_array() {
        let z = build_z_array("aabxaab".as_bytes());
        assert_eq!(z, vec![0, 1, 0, 0, 3, 1, 0]);
    }
}
