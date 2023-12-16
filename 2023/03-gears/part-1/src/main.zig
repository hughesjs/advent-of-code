const std = @import("std");
const allocator = std.heap.page_allocator;
const ArrayList = std.ArrayList;

pub fn main() !void {
    const fPath = try get_file_path_from_args();
    defer allocator.free(fPath);

    const engine_schematic = try get_schematic_from_file_path(fPath);
    defer allocator.free(engine_schematic);

    const answer = try calculate_answer(engine_schematic);
    try output_answer(answer);
}

fn output_answer(answer: u16) !void {
    const out = std.io.getStdOut().writer();
    try out.print("Answer: {d}\n", .{answer});
}

fn get_schematic_from_file_path(fPath: []const u8) ![]const u8 {
    const input_data_file = try std.fs.cwd().openFile(fPath, .{ });
    defer input_data_file.close();

    const file_size = try input_data_file.getEndPos();
    std.log.debug("File Size: {d}", .{file_size});

    const buffer: []u8 = try input_data_file.readToEndAlloc(allocator, file_size);
    std.log.debug("Schematic:\n{s}", .{buffer});
    return buffer;
}

fn get_file_path_from_args() ![]const u8 {
        const args = try std.process.argsAlloc(allocator);
        defer std.process.argsFree(allocator, args);

        std.debug.print("Number of arguments: {d}\n", .{args.len});
        for (args) |arg| {
            std.debug.print("Argument: {s}\n", .{arg});
        }

        const pathBuf = try allocator.dupe(u8, args[1]);
        std.debug.print("fpath: {s}\n", .{pathBuf});
        return pathBuf;
}

fn calculate_answer(engine_schematic: []const u8) !u16 {
    const symbol_mask = try generate_symbol_mask(engine_schematic);
    defer allocator.free(symbol_mask);


    std.log.debug("Symbol indices: \n", .{ });
    for (symbol_mask) |i| {
        std.log.debug("{}, ", .{i});
    }

    return 12;
}

fn generate_symbol_mask(engine_schematic: []const u8) ![]const bool {
    const non_symbol_chars = [_]u8{ '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', '\n' };
    var mask = try allocator.alloc(bool, @sizeOf(u8) * engine_schematic.len);
    for (engine_schematic, 0..) |c, i| {
        if (!array_contains(u8, non_symbol_chars.len,non_symbol_chars, c)) {
            mask[i] = true;
        }
    }
    return mask;
}

fn array_contains(comptime T: type, haystack: []const T, needle: T) bool {
    for (haystack) |e| {
        if (e == needle) {
            return true;
        }
    }
    return false;
}


test "array contains true works" {
    const test_array = [_]u8{ '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.' };

    for (test_array) |c| {
        const found = array_contains(u8, &test_array, c);
        try std.testing.expectEqual(true, found);
    }
}

test "array contains false works" {
    const test_array = [_]u8{ '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.' };
    const other_array = [_]u8{ 'a', 'b', '#', '/', '_', '=' };
    for (other_array) |c| {
        const found = array_contains(u8, &test_array, c);
        try std.testing.expectEqual(false, found);
    }
}

// test "can get symbol indices for line with no symbols" {
//     const test_data = ".......123....*";
// }
//
// test "can get symbol indices for single line" {
//     const test_data = "*.../..123....*";
// }

// test "can get symbol indices for multi line" {
//
// }
//
// test "array contains false works" {
//     const test_array = [_]u8{ '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.' };
//
// }

// test "provided_test_case" {
//     const testData =
//         \\467..114..
//         \\...*......
//         \\..35..633.
//         \\......#...
//         \\617*......
//         \\.....+.58.
//         \\..592.....
//         \\......755.
//         \\...$.*....
//         \\.664.598..
//     ;
//     const expected: u16 = 4361;
//
//     const res = try calculate_answer(testData);
//
//     try std.testing.expectEqual(expected, res);
// }
