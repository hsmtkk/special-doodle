use anyhow::Result;
use std::collections::HashMap;
use text_io::read;

fn main(){
    let _: String = read!("{}\n");
    let ns_str: String = read!("{}\n");
    let ns = split_string(&ns_str).unwrap();
    println!("{}", mean(ns));
    println!("{}", median(ns));
    println!("{}", mode(ns));
}

fn split_string(s:&str) -> Result<Vec<i64>> {
    let mut result: Vec<i64> = Vec::new();
    for s in s.split(" ") {
        let n = s.parse::<i64>()?;
        result.push(n);
    }
    Ok(result)
}

fn mean(ns: Vec<i64>) -> f64 {
    let mut s = 0;
    for n in &ns {
        s += n;
    }
    s as f64 / ns.len() as f64
}

fn median(ns: Vec<i64>) -> f64 {
    let mut ns = ns.clone();
    ns.sort();
    let center = ns.len() / 2;
    if ns.len() %2 == 0 {
        (ns[center-1] as f64+ ns[center]as f64) / 2.0
    } else {
        ns[center] as f64
    }
}

fn mode(ns: Vec<i64>) -> i64 {
    let mut counter: HashMap<i64, i64> = HashMap::new();
    for n in &ns {
        match counter.get(n){
            Some(count) => { counter.insert(*n, count + 1);},
            None => { counter.insert(*n, 1);},
        }
    }
    let mut max_count = 1;
    for count in counter.values() {
        if count > &max_count {
            max_count= *count;
        }
    }
    let mut modes: Vec<i64> = Vec::new();
    for (n, count) in counter {
        if count == max_count {
            modes.push(n);
        }
    }
    modes.sort();
    modes[0]    
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_split_string(){
        let s = "1 2 3";
        let want = vec![1,2,3];
        let got = super::split_string(s).unwrap();
        assert_eq!(want, got);
    }

    #[test]
    fn test_mean(){
        let ns = vec![64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060];
        let want = 43900.6;
        let got = super::mean(ns);
        assert_eq!(want, got);
    }

    #[test]
    fn test_median(){
        let ns = vec![64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060];
        let want = 44627.5;
        let got = super::median(ns);
        assert_eq!(want, got);
    }

    #[test]
    fn test_mode(){
        let ns = vec![64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060];
        let want = 4978;
        let got = super::mode(ns);
        assert_eq!(want, got);
    }
}
