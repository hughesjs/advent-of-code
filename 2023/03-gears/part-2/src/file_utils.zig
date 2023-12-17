const std = @import("std");
pub const string_utils = @import("string_utils.zig");

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
        line_length_out.* = get_line_length(initial_buffer);

        const new_buf = try string_utils.strip_new_lines(allocator, initial_buffer);
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