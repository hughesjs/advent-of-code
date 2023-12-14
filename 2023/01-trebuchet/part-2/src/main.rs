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

fn parse_args(mut args: Args) -> Result<String, Box<dyn Error>> {
    match args.nth(1) {
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

fn get_spelled_num_char_from_buffer(buf: &str) -> Option<char> {
    static DIGIT_MAP: phf::Map<&'static str, char> = phf_map! {
        "one" => '1',
        "two" => '2',
        "three" => '3',
        "four" => '4',
        "five" => '5',
        "six" => '6',
        "seven" => '7',
        "eight" => '8',
        "nine" => '9'
    };

    for fix in DIGIT_MAP.entries() {
        if buf.contains(fix.0) {
            return Some(DIGIT_MAP[fix.0]);
        }
    }

    return None;
}

fn outside_digit_val(buf: &str) -> Result<u16, ParseIntError> {
    let mut line = buf.trim().to_string();
    if line.is_empty() // Skip empty lines
    {
        return Ok(0);
    }

    let mut new_buf: String = String::with_capacity(2);

    // forward pass
    let mut spelled_buf: String = String::with_capacity(line.len());
    for c in line.chars() {
        if c.is_numeric() {
            new_buf.push(c);
            break;
        }

        spelled_buf.push(c);
        if let Some(c) = get_spelled_num_char_from_buffer(&spelled_buf) {
            new_buf.push(c);
            break;
        }
    }

    // backward pass
    spelled_buf = String::with_capacity(line.len());
    for c in line.chars().rev() {
        if c.is_numeric() {
            new_buf.push(c);
            break;
        }

        spelled_buf = c.to_string() + &spelled_buf;
        if let Some(c) = get_spelled_num_char_from_buffer(&spelled_buf) {
            new_buf.push(c);
            break;
        }
    }

    return u16::from_str(&*new_buf);
}

mod tests {
    use super::*;

    #[test]
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

    #[test]
    fn test_overlap() {
        let test_data = r#"
                            twone
                            "#;

        let res = calculate_puzzle(test_data).unwrap();
        assert_eq!(21, res);
    }
}
