const std = @import("std");

pub fn bool_arr_to_bitfield(allocator: std.mem.Allocator, arr: []const bool, line_length: usize) ![]const u8 {
    const lines: usize =  arr.len / line_length;
    const buf: []u8 = try allocator.alloc(u8, arr.len + (lines - 1)); // We're adding in newlines
    var carat = buf.ptr;
    for (arr, 0..) |v, n| {
        if (v) {
            carat[0] = '1';
        }
        else {
            carat[0] = '0';
        }
        carat += 1;
        if ((n + 1) % line_length == 0) {
            carat[0] = '\n';
            carat+=1;
        }
    }

    return buf;
}

test "Can generate single-line bitfield string" {
    const test_data = [_]bool {true, true, false, false, false};
    const expected: []const u8 = "11000";

    const res = try bool_arr_to_bitfield(std.testing.allocator, &test_data, 5);
    defer std.testing.allocator.free(res);

    try std.testing.expectEqualStrings(expected, res);
}

test "Can generate multiline bitfield string" {
    const test_data = [_]bool { true, true, false, false, false,
                                false, true, false, true, false};
    const expected: []const u8 = "11000\n01010";

    const res = try bool_arr_to_bitfield(std.testing.allocator, &test_data, 5);
    defer std.testing.allocator.free(res);

    try std.testing.expectEqualStrings(expected, res);
}