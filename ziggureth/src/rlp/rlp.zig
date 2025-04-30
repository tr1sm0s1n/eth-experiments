const std = @import("std");
const Allocator = std.mem.Allocator;

/// RLP value types
pub const RlpItem = union(enum) {
    bytes: []const u8,
    list: []RlpItem,

    /// Free the memory of this RlpItem and any nested RlpItems
    pub fn deinit(self: RlpItem, allocator: Allocator) void {
        switch (self) {
            .bytes => |b| allocator.free(b),
            .list => |items| {
                for (items) |item| {
                    item.deinit(allocator);
                }
                allocator.free(items);
            },
        }
    }
};

/// RLP encoding error
pub const EncodeError = error{
    InputTooLarge,
    OutOfMemory,
};

/// RLP decoding error
pub const DecodeError = error{
    InvalidData,
    InputTooShort,
    UnexpectedLength,
    OutOfMemory,
};

/// Encode a single byte
fn encodeByte(b: u8) [1]u8 {
    return [_]u8{b};
}

/// Encode a length with a given offset
fn encodeLength(length: usize, offset: u8) !std.ArrayList(u8) {
    const allocator = std.heap.page_allocator;
    var result = std.ArrayList(u8).init(allocator);

    if (length < 56) {
        try result.append(offset + @as(u8, @intCast(length)));
    } else {
        // Convert length to bytes representation
        var length_bytes = std.ArrayList(u8).init(allocator);
        defer length_bytes.deinit();

        var len_val = length;
        while (len_val > 0) {
            try length_bytes.insert(0, @as(u8, @intCast(len_val & 0xFF)));
            len_val >>= 8;
        }

        if (length_bytes.items.len > 8) {
            return EncodeError.InputTooLarge;
        }

        try result.append(offset + 55 + @as(u8, @intCast(length_bytes.items.len)));
        try result.appendSlice(length_bytes.items);
    }

    return result;
}

/// Encode data as RLP bytes
pub fn encodeBytes(data: []const u8) !std.ArrayList(u8) {
    const allocator = std.heap.page_allocator;
    var result = std.ArrayList(u8).init(allocator);

    if (data.len == 1 and data[0] < 0x80) {
        // Single byte below 0x80 is self-represented
        try result.append(data[0]);
        return result;
    }

    // String of bytes
    var length_prefix = try encodeLength(data.len, 0x80);
    defer length_prefix.deinit();

    try result.appendSlice(length_prefix.items);
    try result.appendSlice(data);
    return result;
}

/// Encode an RLP list
pub fn encodeList(items: []const RlpItem) !std.ArrayList(u8) {
    const allocator = std.heap.page_allocator;
    var encoded_items = std.ArrayList(u8).init(allocator);
    defer encoded_items.deinit();

    // First, encode all items in the list
    for (items) |item| {
        var encoded = switch (item) {
            .bytes => |b| try encodeBytes(b),
            .list => |l| try encodeList(l),
        };
        defer encoded.deinit();
        try encoded_items.appendSlice(encoded.items);
    }

    // Then encode the list with the 0xC0 offset
    var result = std.ArrayList(u8).init(allocator);
    var length_prefix = try encodeLength(encoded_items.items.len, 0xC0);
    defer length_prefix.deinit();

    try result.appendSlice(length_prefix.items);
    try result.appendSlice(encoded_items.items);
    return result;
}

/// Encode an RLP item (either bytes or list)
pub fn encode(item: RlpItem) !std.ArrayList(u8) {
    return switch (item) {
        .bytes => |b| try encodeBytes(b),
        .list => |l| try encodeList(l),
    };
}

/// Decode RLP data
pub fn decode(allocator: Allocator, data: []const u8) !RlpItem {
    var offset: usize = 0;
    return try decodeItem(allocator, data, &offset);
}

/// Decode a single RLP item and advance the offset
fn decodeItem(allocator: Allocator, data: []const u8, offset: *usize) !RlpItem {
    if (offset.* >= data.len) {
        return DecodeError.InputTooShort;
    }

    const first_byte = data[offset.*];
    offset.* += 1;

    // Single byte case (0x00 - 0x7F)
    if (first_byte < 0x80) {
        const result = try allocator.dupe(u8, &[_]u8{first_byte});
        return RlpItem{ .bytes = result };
    }
    // Short string (0x80 - 0xB7)
    else if (first_byte < 0xB8) {
        const length = first_byte - 0x80;

        if (offset.* + length > data.len) {
            return DecodeError.InputTooShort;
        }

        const str_data = data[offset.* .. offset.* + length];
        offset.* += length;

        const result = try allocator.dupe(u8, str_data);
        return RlpItem{ .bytes = result };
    }
    // Long string (0xB8 - 0xBF)
    else if (first_byte < 0xC0) {
        const length_bytes_count = first_byte - 0xB7;

        if (offset.* + length_bytes_count > data.len) {
            return DecodeError.InputTooShort;
        }

        var length: usize = 0;
        for (0..length_bytes_count) |i| {
            length = length << 8 | data[offset.* + i];
        }
        offset.* += length_bytes_count;

        if (offset.* + length > data.len) {
            return DecodeError.InputTooShort;
        }

        const str_data = data[offset.* .. offset.* + length];
        offset.* += length;

        const result = try allocator.dupe(u8, str_data);
        return RlpItem{ .bytes = result };
    }
    // Short list (0xC0 - 0xF7)
    else if (first_byte < 0xF8) {
        const total_length = first_byte - 0xC0;

        if (offset.* + total_length > data.len) {
            return DecodeError.InputTooShort;
        }

        const end_offset = offset.* + total_length;
        var items = std.ArrayList(RlpItem).init(allocator);

        while (offset.* < end_offset) {
            const item = try decodeItem(allocator, data, offset);
            try items.append(item);
        }

        if (offset.* != end_offset) {
            // Clean up any items we've already decoded
            for (items.items) |item| {
                item.deinit(allocator);
            }
            items.deinit();
            return DecodeError.UnexpectedLength;
        }

        return RlpItem{ .list = try items.toOwnedSlice() };
    }
    // Long list (0xF8 - 0xFF)
    else {
        const length_bytes_count = first_byte - 0xF7;

        if (offset.* + length_bytes_count > data.len) {
            return DecodeError.InputTooShort;
        }

        var total_length: usize = 0;
        for (0..length_bytes_count) |i| {
            total_length = total_length << 8 | data[offset.* + i];
        }
        offset.* += length_bytes_count;

        if (offset.* + total_length > data.len) {
            return DecodeError.InputTooShort;
        }

        const end_offset = offset.* + total_length;
        var items = std.ArrayList(RlpItem).init(allocator);

        while (offset.* < end_offset) {
            const item = try decodeItem(allocator, data, offset);
            try items.append(item);
        }

        if (offset.* != end_offset) {
            // Clean up any items we've already decoded
            for (items.items) |item| {
                item.deinit(allocator);
            }
            items.deinit();
            return DecodeError.UnexpectedLength;
        }

        return RlpItem{ .list = try items.toOwnedSlice() };
    }
}

// Test encoding and decoding
test "Test single byte" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const test_byte = RlpItem{ .bytes = &[_]u8{0x42} };
    const encoded = try encode(test_byte);
    defer encoded.deinit();

    try std.testing.expectEqualSlices(u8, &[_]u8{0x42}, encoded.items);

    const decoded = try decode(allocator, encoded.items);
    defer decoded.deinit(allocator);

    try std.testing.expectEqualSlices(u8, &[_]u8{0x42}, decoded.bytes);
}

test "Test short string" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const test_str = RlpItem{ .bytes = "zig" };
    const encoded = try encode(test_str);
    defer encoded.deinit();

    try std.testing.expectEqualSlices(u8, &[_]u8{ 0x83, 'z', 'i', 'g' }, encoded.items);

    const decoded = try decode(allocator, encoded.items);
    defer decoded.deinit(allocator);

    try std.testing.expectEqualSlices(u8, "zig", decoded.bytes);
}

test "Test empty list" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const empty_list = RlpItem{ .list = &[_]RlpItem{} };
    const encoded = try encode(empty_list);
    defer encoded.deinit();

    try std.testing.expectEqualSlices(u8, &[_]u8{0xC0}, encoded.items);

    const decoded = try decode(allocator, encoded.items);
    defer decoded.deinit(allocator);

    try std.testing.expectEqual(@as(usize, 0), decoded.list.len);
}

test "Test nested list" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var items = [_]RlpItem{
        RlpItem{ .bytes = "hello" },
        RlpItem{ .bytes = "world" },
    };

    const list = RlpItem{ .list = &items };
    const encoded = try encode(list);
    defer encoded.deinit();

    const expected = [_]u8{ 0xCC, 0x85, 'h', 'e', 'l', 'l', 'o', 0x85, 'w', 'o', 'r', 'l', 'd' };
    try std.testing.expectEqualSlices(u8, &expected, encoded.items);

    const decoded = try decode(allocator, encoded.items);
    defer decoded.deinit(allocator);

    try std.testing.expectEqual(@as(usize, 2), decoded.list.len);
    try std.testing.expectEqualSlices(u8, "hello", decoded.list[0].bytes);
    try std.testing.expectEqualSlices(u8, "world", decoded.list[1].bytes);
}
