const std = @import("std");

pub fn get_file_contents(allocator: std.mem.Allocator, fPath: []const u8) ![]const u8 {
    const input_data_file = try std.fs.cwd().openFile(fPath, .{ });
    defer input_data_file.close();

    const file_size = try input_data_file.getEndPos();
    std.log.debug("File Size: {d}", .{file_size});

    const buffer: []u8 = try input_data_file.readToEndAlloc(allocator, file_size);
    std.log.debug("Schematic:\n{s}", .{buffer});
    return buffer;
}

pub fn get_file_contents_as_single_line(allocator: std.mem.Allocator, fPath: []const u8, line_length_out: *usize ) ![]const u8 {
        const initial_buffer = try get_file_contents(allocator, fPath);
        defer allocator.free(initial_buffer);
        line_length_out.* = get_line_length(initial_buffer) + 1;

        const new_buf = try strip_new_lines(allocator, initial_buffer);
        return new_buf;
}

fn strip_new_lines(allocator: std.mem.Allocator, buf: []const u8) ![]const u8 {
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

fn get_line_length(buf: []const u8) usize {
        for (buf, 0..) |c, i| {
            if (c == '\n') {
               return i;
            }
        }
        return buf.len;
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

test "Can get line length single line" {
        const test_data = "12345";
        const expected: usize = 5;

        const res = get_line_length(test_data);

        try std.testing.expectEqual(expected, res);
}

test "Can get line length multi line" {
            const test_data = "123\n123\n123";
            const expected: usize = 3;

            const res = get_line_length(test_data);

            try std.testing.expectEqual(expected, res);
}