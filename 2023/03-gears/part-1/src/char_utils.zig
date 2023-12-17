const std = @import("std");

pub fn char_is_numeric(c: u8) bool {
    return c >= '0' and c <= '9';
}