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