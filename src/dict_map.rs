use anyhow::{anyhow, Result};
use std::collections::HashMap;

fn match_kvs(kvs: Vec<&str>, keys: Vec<&str>) -> Result<Vec<String>> {
    let mut kvmap: HashMap<&str, &str> = HashMap::new();
    for kv in kvs {
        let elems: Vec<&str> = kv.split(' ').collect();
        if elems.len() != 2 {
            return Err(anyhow!("invalid format {}", kv));
        }
        let k = elems[0];
        let v = elems[1];
        kvmap.insert(k, v);
    }
    let mut result: Vec<String> = Vec::new();
    for k in keys {
        let v = match kvmap.get(k) {
            Some(v) => format!("{}={}", k, v),
            None => "Not found".to_string(),
        };
        result.push(v.to_string());
    }
    Ok(result)
}

#[cfg(test)]
mod tests {
    #[test]
    fn test0() {
        let kvs = vec!["sam 99912222", "tom 11122222", "harry 12299933"];
        let keys = vec!["sam", "edward", "harry"];
        let want = vec!["sam=99912222".to_string(), "Not found".to_string(), "harry=12299933".to_string()];
        let got = super::match_kvs(kvs, keys).unwrap();
        assert_eq!(want, got);
    }
}
