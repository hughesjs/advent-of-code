const std = @import("std");
pub const array_utils = @import("./array_utils.zig");
pub const file_utils = @import("./file_utils.zig");
pub const debug_utils = @import("./debug_utils.zig");
pub const char_utils = @import("./char_utils.zig");

const ArrayList = std.ArrayList;
const allocator = std.heap.page_allocator;

pub fn main() !void {
    const fPath = try get_file_path_from_args();
    defer allocator.free(fPath);

    var line_length: usize = undefined;
    const engine_schematic = try file_utils.get_file_contents_as_single_line(allocator, fPath, &line_length);
    std.log.debug("Line Length: {d}", .{line_length});
    defer allocator.free(engine_schematic);

    const answer = try calculate_answer(engine_schematic, line_length);
    try output_answer(answer);
}

fn output_answer(answer: u16) !void {
    const out = std.io.getStdOut().writer();
    try out.print("Answer: {d}\n", .{answer});
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

fn calculate_answer(engine_schematic: []const u8, line_length: usize) !u16 {
    const symbol_mask = try generate_symbol_mask(engine_schematic);
    defer allocator.free(symbol_mask);

    std.log.debug("Symbol mask: \n{s}", .{try debug_utils.bool_arr_to_bitfield(allocator, symbol_mask, line_length)});

    const code_slices: ArrayList(slice_with_index) = try get_id_code_slices(engine_schematic);
    defer code_slices.deinit();

    for (code_slices.items) |cs| {
        const adjacents = try get_adjacent_indices(cs, line_length, engine_schematic.len);
        for (adjacents.items) |i| {
            if (symbol_mask[i]) {
                std.log.debug("CODE: {s}", .{cs.slice});
                break;
            }
            std.log.debug("NOT CODE: {s}", .{cs.slice});
        }
    }

    return 12;
}

fn get_adjacent_indices(slindex: slice_with_index, line_length: usize, buf_len: usize) !ArrayList(usize) {
    // TODO - Deal with wraparound later
    const first_index = slindex.index;
    const last_index = slindex.index + slindex.slice.len - 1;
    const max_adjacents = 2 * slindex.slice.len + 6;

    var index_list = try ArrayList(usize).initCapacity(allocator, max_adjacents);
    errdefer index_list.deinit();
    std.log.debug("Getting adjacents for: {s} @ {d} -> {d}", .{slindex.slice, slindex.index, slindex.index + slindex.slice.len});
    // Top row
    if (first_index >= line_length + 1) {
        for ((first_index - (line_length + 1))..(last_index - (line_length - 1))) |i| {
            if (i < 0 or i > buf_len) {
                continue;
            }
            try index_list.append(i);
        }
    }
    else {
        std.log.debug("Going down awkward path.\nLast: {d}\nLen: {d}", .{last_index, line_length});
        // Deal with negative usize
        if (last_index >= line_length - 1) { // We have at least "some" above the line... This might be removable after wrap is sorted
                for (0..(last_index - (line_length - 1))) |i| {
                    if (i < 0 or i > buf_len) {
                        continue;
                    }
                    try index_list.append(i);
                }
        }
    }


    // Bottom row
    for ((first_index + (line_length - 1))..(last_index + (line_length + 1))) |i| {
        if (i >= buf_len) {
            continue;
        }
        try index_list.append(i);
    }

    // Left
    if (first_index - 1 >= 0) {
        try index_list.append(first_index - 1);
    }


    // Right
    if (last_index < buf_len) {
       try index_list.append(last_index + 1);
    }

    return index_list;
}

fn get_id_code_slices(engine_schematic: []const u8) !ArrayList(slice_with_index) {
    var symbol_list = ArrayList(slice_with_index).init(allocator);
    errdefer symbol_list.deinit();

    var current_start: usize = 0;
    var current_end: usize = 0;
    var in_code = false;
    for (engine_schematic, 0..) |c, i| {
        const is_numeric: bool = char_utils.char_is_numeric(c);

        // Yes, we could do some fewer checks with nesting but this is easier to grok...
        if (!in_code and !is_numeric) {
            continue;
        }

        if (!in_code and is_numeric) {
            current_start = i;
            current_end = i;
            in_code = true;
            continue;
        }

        if (is_numeric) {
            continue;
        }

        current_end = i; // Slice end is exclusive
        in_code = false;
        const slice = engine_schematic[current_start..current_end];
        std.log.debug("Slice: {d} -> {d} -- {s}", .{current_start, current_end, slice});
        try symbol_list.append(slice_with_index{.slice = slice, .index = current_start});
    }

    return symbol_list;
}

const slice_with_index = struct {
    slice: []const u8,
    index: usize //Slice start index in outer buffer
};

const slindex_with_adjacents = struct {
    slindex: slice_with_index,
    adjacents: ArrayList(u8)
};

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
