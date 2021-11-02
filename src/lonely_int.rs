use std::collections::HashSet;

#[allow(dead_code)]
fn lonely_int(ns: Vec<i32>) -> i32 {
    let mut lonely = HashSet::new();
    for n in ns {
        if lonely.contains(&n){
            lonely.remove(&n);
        } else {
            lonely.insert(n);
        }
    }
    for n in lonely.iter(){
        return *n;
    }
    0
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0(){
        let want = 1;
        let got = super::lonely_int(vec![1]);
        assert_eq!(want, got);
    }

    #[test]
    fn test1(){
        let want = 2;
        let got = super::lonely_int(vec![1, 1, 2]);
        assert_eq!(want, got);
    }

    #[test]
    fn test2(){
        let want = 2;
        let got = super::lonely_int(vec![0, 0, 1, 2, 1]);
        assert_eq!(want, got);
    }
}
