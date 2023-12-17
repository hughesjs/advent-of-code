const std = @import("std");

pub fn strip_new_lines(allocator: std.mem.Allocator, buf: []const u8) ![]const u8 {
    const num_new_lines = count_new_lines(buf);
    const new_size = buf.len - num_new_lines;
    const new_buf: []u8 = try allocator.alloc(u8, new_size);
    var carat = new_buf.ptr;

    for (buf) |c| {
        if (c != '\n') {
            carat[0] = c;
            carat += 1;
        }
    }
    return new_buf;
}

fn count_new_lines(buf: []const u8) usize {
    var n: usize = 0;
    for (buf) |c| {
        if (c == '\n') {
            n += 1;
        }
    }
    return n;
}

test "Can strip new lines multiple" {
    const test_data = "asdasd\ndsadsa\nasdasd";
    const expected = "asdasddsadsaasdasd";

    const res = try strip_new_lines(std.testing.allocator, test_data);
    defer std.testing.allocator.free(res);

    try std.testing.expectEqualStrings(expected, res);
}

test "Can strip new lines single" {
    const test_data = "asdasd\n";
    const expected = "asdasd";

    const res = try strip_new_lines(std.testing.allocator, test_data);
    defer std.testing.allocator.free(res);

    try std.testing.expectEqualStrings(expected, res);
}



test "Can count new lines" {
    const test_data = "asdasd\nadsasd\nasdasda";
    const expected: usize = 2;

    const res = count_new_lines(test_data);

    try std.testing.expectEqual(expected, res);
}