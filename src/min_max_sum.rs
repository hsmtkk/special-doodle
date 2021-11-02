#[derive(Debug, PartialEq)]
struct SumMinMax {
    min: i64,
    max: i64,
}

impl SumMinMax {
    #[allow(dead_code)]
    fn new(min: i64, max: i64) -> SumMinMax {
        SumMinMax { min, max }
    }
}

#[allow(dead_code)]
fn min_max_sum(arr: Vec<i64>) -> SumMinMax {
    let mut min: i64 = 0;
    let mut max: i64 = 0;
    for skip in 0..5 {
        let mut s: i64 = 0;
        for i in 0..5 {
            if i == skip {
                continue;
            }
            s += arr[i];
        }
        if min == 0 {
            min = s;
        }
        if max == 0 {
            max = s;
        }
        if s < min {
            min = s;
        }
        if s > max {
            max = s;
        }
    }
    SumMinMax::new(min, max)
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_a() {
        let arr = vec![1, 3, 5, 7, 9];
        let want = super::SumMinMax::new(16, 24);
        let got = super::min_max_sum(arr);
        assert_eq!(want, got);
    }

    #[test]
    fn test_b() {
        let arr = vec![1, 2, 3, 4, 5];
        let want = super::SumMinMax::new(10, 14);
        let got = super::min_max_sum(arr);
        assert_eq!(want, got);
    }

    #[test]
    fn test_1() {
        let arr = vec![256741038, 623958417, 467905213, 714532089, 938071625];
        let want = super::SumMinMax::new(2063136757, 2744467344);
        let got = super::min_max_sum(arr);
        assert_eq!(want, got);
    }
}
