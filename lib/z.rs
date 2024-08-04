struct ZArray {
    arr: Vec<usize>,
}

impl ZArray {
    pub fn new(s: &[char]) -> Self {
        let n = s.len();
        let mut arr = vec![0; n];
        let mut l = 0;
        let mut r = 0;
        let mut k = 0;
        for i in 0..n {
            if i > r {
                l = i;
                r = i;
                while r < n && s[r] == s[r-1] {
                    r += 1;
                }
                arr[i] = r - l;
                r--
            } else {
                k = i - 1;
                if arr[k] < r - i + 1 {
                    arr[i] = arr[k];
                } else {
                    l = i;
                    while r < n && s[r] == s[r-1] {
                        r += 1;
                    }
                    arr[i] = r - 1;
                    r -= 1;
                }
            }
        }

        ZArray{
            arr
        }
    }
}

