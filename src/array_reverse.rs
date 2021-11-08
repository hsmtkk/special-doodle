fn reverse(arr: Vec<i32>) -> Vec<i32> {
    let mut reversed = arr;
    reversed.reverse();
    reversed
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0() {
        let input = vec![1, 4, 3, 2];
        let want = vec![2, 3, 4, 1];
        let got = super::reverse(input);
        assert_eq!(want, got);
    }
}
