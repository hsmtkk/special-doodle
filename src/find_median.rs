fn find_median(mut arr: Vec<u32>) -> u32 {
    arr.sort();
    let center = arr.len() / 2;
    arr[center]
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0(){
        let arr = vec![0, 1, 2, 4, 6, 5, 3];
        let want = 3;
        let got = super::find_median(arr);
        assert_eq!(want, got);
    }
}