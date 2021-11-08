fn lets_review(s:&str) -> String {
    let mut first = String::new();
    let mut second = String::new();
    for (i, ch) in s.chars().enumerate() {
        if i % 2 == 0{
            first.push(ch);
        } else {
            second.push(ch);
        }
    }
    format!("{} {}", first, second)
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0(){
        let want = "Hce akr";
        let got = super::lets_review("Hacker");
        assert_eq!(want, got);
    }

    #[test]
    fn test1(){
        let want = "Rn ak";
        let got = super::lets_review("Rank");
        assert_eq!(want, got);
    }
}