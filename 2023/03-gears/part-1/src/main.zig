const std = @import("std");
pub const array_utils = @import("./array_utils.zig");
pub const file_utils = @import("./file_utils.zig");
pub const debug_utils = @import("./debug_utils.zig");

const ArrayList = std.ArrayList;
const allocator = std.heap.page_allocator;

pub fn main() !void {
    const fPath = try get_file_path_from_args();
    defer allocator.free(fPath);

    const engine_schematic = try file_utils.get_file_contents(allocator, fPath);
    defer allocator.free(engine_schematic);

    const answer = try calculate_answer(engine_schematic);
    try output_answer(answer);
}

fn output_answer(answer: u16) !void {
    const out = std.io.getStdOut().writer();
    try out.print("Answer: {d}\n", .{answer});
}


// TODO - Simplify this and use [:0] const
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


    std.log.debug("Symbol mask: \n{s}", .{try debug_utils.bool_arr_to_bitfield(allocator, symbol_mask, 10)});

    return 12;
}

fn generate_symbol_mask(engine_schematic: []const u8) ![]const bool {
    const non_symbol_chars = [_]u8{ '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', '\n' };
    var mask = try allocator.alloc(bool, engine_schematic.len);
    for (engine_schematic, mask) |c, *m| {
        if (!array_utils.array_contains(u8, &non_symbol_chars, c)) {
            m.* = true;
        }
    }
    return mask;
}


test "Test dependencies" {
   std.testing.refAllDeclsRecursive(@This());
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
