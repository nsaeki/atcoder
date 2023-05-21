#[macro_export]
macro_rules! dpad {
    ( $i:expr, $j:expr ) => {{
        let mut temp_vec = Vec::new();
        if $i > 0 {
            temp_vec.push(($i - 1, $j));
        }
        temp_vec.push(($i + 1, $j));
        if $j > 0 {
            temp_vec.push(($i, $j - 1));
        }
        temp_vec.push(($i, $j + 1));
        temp_vec
    }};
    ( $i:expr, $j:expr, $n:expr ) => {{
        let mut temp_vec = Vec::new();
        if $i > 0 {
            temp_vec.push(($i - 1, $j));
        }
        if $i < $n - 1 {
            temp_vec.push(($i + 1, $j));
        }
        if $j > 0 {
            temp_vec.push(($i, $j - 1));
        }
        if $j < $n - 1 {
            temp_vec.push(($i, $j + 1));
        }
        temp_vec
    }};
    ( $i:expr, $j:expr, $m:expr, $n:expr ) => {{
        let mut temp_vec = Vec::new();
        if $i > 0 {
            temp_vec.push(($i - 1, $j));
        }
        if $j < $m - 1 {
            temp_vec.push(($i + 1, $j));
        }
        if $j > 0 {
            temp_vec.push(($i, $j - 1));
        }
        if $j < $n - 1 {
            temp_vec.push(($i, $j + 1));
        }
        temp_vec
    }};
}

fn main() {
    assert_eq!(dpad!(0, 9), vec![(1, 9), (0, 8), (0, 10)]);
    assert_eq!(dpad!(0, 9, 10), vec![(1, 9), (0, 8)]);
    assert_eq!(dpad!(0, 9, 100, 10), vec![(1, 9), (0, 8)]);
    assert_eq!(dpad!(0, 9, 100, 100), vec![(1, 9), (0, 8), (0, 10)]);
    assert_eq!(
        dpad!(10, 10, 100, 100),
        vec![(9, 10), (11, 10), (10, 9), (10, 11)]
    );
}
