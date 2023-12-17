const std = @import("std");
const ArrayList = std.ArrayList;

const slindex_with_adjacent = @import("./main.zig").slindex_with_adjacents;

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

pub fn generate_adjacency_mask(allocator: std.mem.Allocator, slindices: ArrayList(slindex_with_adjacent), symbol_mask: []const bool, line_len: usize) ![]const u8 {
    const lines: usize =  symbol_mask.len / line_len;
    std.log.debug("1d buf len {d}", .{symbol_mask.len});
    // Create a 1d version of the mask
    var encoded_schema: []u8 = try allocator.alloc(u8, symbol_mask.len);
    for (0..encoded_schema.len) |i| {
        encoded_schema[i] = '.';
    }

    for (0..symbol_mask.len) |i| {
        if (symbol_mask[i]) {
            encoded_schema[i]='s';
        }
    }

    for (slindices.items) |s| {
        for (s.slindex.index..(s.slindex.index + s.slindex.slice.len)) |i| {
            encoded_schema[i] = s.slindex.slice[i - s.slindex.index];
        }
        for (s.adjacents.items) |a| {
            if (symbol_mask[a]) {
                encoded_schema[a]='X';
                continue;
            }
            encoded_schema[a] = 'a';
        }
    }

     std.log.debug("1d buf {s}", .{encoded_schema});

    const buf: []u8 = try allocator.alloc(u8, symbol_mask.len + lines); // We're adding in newlines

    var carat = buf.ptr;
    for (encoded_schema, 0..) |v, n| {
     carat[0] = v;
     carat += 1;
     if ((n + 1) % line_len == 0) {
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