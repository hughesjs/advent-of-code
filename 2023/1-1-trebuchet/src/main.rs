use std::error::Error;
use std::io;
use std::num::ParseIntError;
use std::str::FromStr;

fn main() -> Result<(), Box<dyn Error>> {
    let mut buf: String = String::new();
    let mut acc: u16 = 0;
    while io::stdin().read_line(&mut buf).is_ok_and(|s| s > 0) {
        acc += outside_digit_val(&*buf)?;
        buf = String::new();
    }

    dbg!(acc);
    Ok(())
}

fn outside_digit_val(buf: &str) -> Result<u16, ParseIntError> {
    let mut new_buf: String = String::new();
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