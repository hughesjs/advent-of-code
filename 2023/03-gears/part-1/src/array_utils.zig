const std = @import("std");

pub fn array_contains(comptime T: type, haystack: []const T, needle: T) bool {
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