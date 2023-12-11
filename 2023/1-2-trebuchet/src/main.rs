use std::error::Error;
use std::fs;
use std::env::Args;
use std::num::ParseIntError;
use std::str::FromStr;
use phf::phf_map;
fn main() -> Result<(), Box<dyn Error>> {
    let file_path = parse_args(std::env::args())?;
    let data = fs::read_to_string(file_path)?;
    let res = calculate_puzzle(&data)?;
    print!("Result: {}", res);
    Ok(())
}

fn parse_args(mut args: Args) -> Result<String, Box<dyn Error>>{
    match args.nth(1)     {
        Some(s) => Ok(s),
        None => Err("No arguments provided".into())
    }
}

fn calculate_puzzle(data: &str) -> Result<u16, Box<dyn Error>> {
    let acc = data.split('\n').try_fold(0, |acc, token| {
        let line_val = outside_digit_val(token)?;
        Ok(acc + line_val)
    });
    return acc;
}

fn outside_digit_val(buf: &str) -> Result<u16, ParseIntError> {
    if buf.trim().is_empty() // Skip empty lines
    {
        return Ok(0);
    }

    static DIGIT_MAP: phf::Map<&'static str, &'static str> = phf_map! {
        "one" => "1",
        "two" => "2",
        "three" => "3",
        "four" => "4",
        "five" => "5",
        "six" => "6",
        "seven" => "7",
        "eight" => "8",
        "nine" => "9"
    };

    let mut fixed_line = buf.trim().to_string();

    for fix in DIGIT_MAP.entries() {
        fixed_line = fixed_line.replace(fix.0, fix.1)
    }

    let mut new_buf: String = String::with_capacity(2);
    // forward pass
    for c in buf.chars() {
        if c.is_numeric() {
            new_buf.push(c);
            break;
        }
    }
    // backward pass
    for c in buf.chars().rev() {
        if c.is_numeric() {
            new_buf.push(c);
            break;
        }
    }

    return u16::from_str(&*new_buf);
}

mod tests {
    use super::*;
    # [test]
    fn run_test_data() {
        let test_data = r#"
                            two1nine
                            eightwothree
                            abcone2threexyz
                            xtwone3four
                            4nineeightseven2
                            zoneight234
                            7pqrstsixteen
                            "#;

        let res = calculate_puzzle(test_data).unwrap();
        assert_eq!(281, res);
    }
}
