#[allow(dead_code)]
fn std_dev(ns : Vec<i32>) -> f32 {
    let mn = mean(&ns);
    let mut sum: f32 = 0.0;
    for n in &ns {
        sum += pow(*n as f32 - mn);
    }
    (sum / ns.len() as f32).sqrt()
}

#[allow(dead_code)]
fn mean(ns: &Vec<i32>) -> f32 {
    let mut sum = 0;
    for n in ns {
        sum += *n;
    }
    sum as f32 / ns.len() as f32
}

#[allow(dead_code)]
fn pow(f: f32) -> f32 {
    f * f
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0() {
        let arr = vec![10, 40, 30, 50, 20];
        let want = 14.1 as u32;
        let got = super::std_dev(arr) as u32;
        assert_eq!(want, got);
    }
    
}