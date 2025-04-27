const std = @import("std");
const types = @import("./types.zig");

pub fn rpcCall(
    client: *std.http.Client,
    allocator: std.mem.Allocator,
    url: []const u8,
    method: []const u8,
) !std.json.Parsed(types.SimpleResponse) {
    const payload = try std.fmt.allocPrint(allocator, "{{\"jsonrpc\":\"2.0\",\"method\":\"{s}\",\"params\":{s},\"id\":{}}}", .{ method, "[]", 1 });

    var response_body = std.ArrayList(u8).init(allocator);

    _ = try client.fetch(.{ .method = .POST, .location = .{ .url = url }, .extra_headers = &[_]std.http.Header{
        .{ .name = "Content-Type", .value = "application/json" },
    }, .response_storage = .{ .dynamic = &response_body }, .payload = payload });

    return try std.json.parseFromSlice(types.SimpleResponse, allocator, response_body.items, .{ .ignore_unknown_fields = true });
}

pub fn latestHeader(client: *std.http.Client, allocator: std.mem.Allocator, url: []const u8) !types.Header {
    const payload = try std.fmt.allocPrint(allocator, "{{\"jsonrpc\":\"2.0\",\"method\":\"{s}\",\"params\":{s},\"id\":{}}}", .{ "eth_getBlockByNumber", "[\"latest\",true]", 1 });

    var response_body = std.ArrayList(u8).init(allocator);

    _ = try client.fetch(.{ .method = .POST, .location = .{ .url = url }, .extra_headers = &[_]std.http.Header{
        .{ .name = "Content-Type", .value = "application/json" },
    }, .response_storage = .{ .dynamic = &response_body }, .payload = payload });

    const parsed = try std.json.parseFromSlice(struct { result: types.Header }, allocator, response_body.items, .{ .ignore_unknown_fields = true });

    return parsed.value.result;
}
