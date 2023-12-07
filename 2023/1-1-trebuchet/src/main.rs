use std::io;
use std::str::FromStr;

fn main() {
    let mut buff: String = String::new();
    let mut acc: u16 = 0;
    while io::stdin().read_line(&mut buff).is_ok_and(|s| s > 0) {
        let mut first_number: Option<char> = None;
        let mut last_number: Option<char> = None;
        // forward pass
        for c in buff.chars() {
            if c.is_numeric() {
                first_number = Some(c);
                break;
            }
        }
        // backward pass
        for c in buff.chars().rev() {
            if c.is_numeric() {
                last_number = Some(c);
                break;
            }
        }

        let first_char = match first_number {
            None => '\0',
            Some(c) => c
        };

        let last_char = match last_number {
            None => '\0',
            Some(c) => c
        };

        let char_string = format!("{first_char}{last_char}");
        let val = u16::from_str(&*char_string);

        acc += match val {
            Ok(v) => v,
            Err(_) => 0
        };

        buff = String::new();
    }

    dbg!(acc);
}
