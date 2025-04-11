const std = @import("std");

pub const Response = struct {
    result: []u8,
};

pub fn rpcCall(
    client: *std.http.Client,
    allocator: std.mem.Allocator,
    url: []const u8,
    method: []const u8,
) !std.json.Parsed(Response) {
    const payload = try std.fmt.allocPrint(allocator, "{{\"jsonrpc\":\"2.0\",\"method\":\"{s}\",\"params\":{s},\"id\":{}}}", .{ method, "[]", 1 });

    var response_body = std.ArrayList(u8).init(allocator);

    _ = try client.fetch(.{ .method = .POST, .location = .{ .url = url }, .extra_headers = &[_]std.http.Header{
        .{ .name = "Content-Type", .value = "application/json" },
    }, .response_storage = .{ .dynamic = &response_body }, .payload = payload });

    return try std.json.parseFromSlice(Response, allocator, response_body.items, .{ .ignore_unknown_fields = true });
}
