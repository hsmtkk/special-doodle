#[allow(dead_code)]
fn calculate_total_cost(meal_cost:f32, tip_percent:u32, tax_percent:u32) -> u32 {
    let tip = meal_cost * tip_percent as f32 * 0.01;
    let tax = meal_cost * tax_percent as f32 * 0.01;
    let total = meal_cost + tip + tax;
    total.round() as u32
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0(){
        let want = 15;
        let got = super::calculate_total_cost(12.0, 20, 8);
        assert_eq!(want, got)
    }

    #[test]
    fn test1(){
        let want = 19;
        let got = super::calculate_total_cost(15.5, 15, 10);
        assert_eq!(want, got)
    }

    #[test]
    fn test3(){
        let want = 13;
        let got = super::calculate_total_cost(10.25, 17, 5);
        assert_eq!(want, got)
    }
}