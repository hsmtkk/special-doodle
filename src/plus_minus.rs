#[derive(Debug, PartialEq)]
struct Ratio  {
    positive_ratio: f32,
    negative_ratio: f32,
    zero_ratio: f32,
}

impl Ratio {
    #[allow(dead_code)]
    fn new(positive_ratio: f32, negative_ratio: f32, zero_ratio: f32) -> Ratio {
        Ratio{positive_ratio, negative_ratio, zero_ratio}
    }
}

#[allow(dead_code)]
fn get_ratio(arr: &Vec<i32>) -> Ratio {
    let mut pos_count = 0;
    let mut neg_count = 0;
    let mut zero_count = 0;
    for n in arr {
        if n > &0 {
            pos_count += 1;
        } else if n < &0 {
            neg_count += 1;
        } else {
            zero_count += 1;
        }
    }
    let arr_len: f32 = arr.len() as f32;
    let positive_ratio = pos_count as f32 / arr_len;
    let negative_ratio = neg_count as f32 / arr_len;
    let zero_ratio = zero_count as f32  / arr_len;
    Ratio{positive_ratio, negative_ratio, zero_ratio}
}

#[cfg(test)]
mod tests {
    #[test]
    fn test1(){
        let arr = vec![-4,3,-9,0,4,1];
        let want = super::Ratio::new(1.0/2.0, 1.0/3.0, 1.0/6.0);
        let got = super::get_ratio(&arr);
        assert_eq!(want, got);
    }
}