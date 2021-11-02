#[allow(dead_code)]
const WEIRD: &str = "Weird";
const NOT_WEIRD: &str = "Not Weird";

#[allow(dead_code)]
fn get_message(n: u32) -> String {
    if n % 2 == 1 {
        return WEIRD.to_string();
    }
    if 2 <= n && n <= 5 {
        return NOT_WEIRD.to_string();
    }
    if 6 <= n && n <= 20 {
        return WEIRD.to_string();
    }
    NOT_WEIRD.to_string()
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0(){
        assert_eq!(super::WEIRD, super::get_message(3));
        assert_eq!(super::NOT_WEIRD, super::get_message(24));
    }
}