#[allow(dead_code)]
fn weighted_mean(x:Vec<u32>, w:Vec<u32>) -> f32 {
    let mut sum = 0;
    let mut sum_of_w = 0;
    for i in 0..x.len() {
        sum += x[i] * w[i];
        sum_of_w += w[i];
    }
    sum as f32 / sum_of_w as f32
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0(){
        let x = vec![10,40,30,50,20];
        let w = vec![1,2,3,4,5];
        let want = 32.0;
        let got = super::weighted_mean(x, w);
        assert_eq!(want, got);
    }
}