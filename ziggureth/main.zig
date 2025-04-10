const std = @import("std");
const writer = std.io.getStdOut().writer();

pub fn main() !void {
    const alloc = std.heap.page_allocator;
    var arena = std.heap.ArenaAllocator.init(alloc);
    const allocator = arena.allocator();

    defer arena.deinit();

    var client = std.http.Client{
        .allocator = allocator,
    };

    const headers = &[_]std.http.Header{
        .{ .name = "Content-Type", .value = "application/json" },
    };

    const response = try clientVersion("http://127.0.0.1:8545", headers, &client, alloc);

    const result = try std.json.parseFromSlice(Result, allocator, response.items, .{ .ignore_unknown_fields = true });

    try writer.print("Client: {s}\n", .{result.value.result});
}

const Result = struct {
    result: []u8,
};

fn clientVersion(
    url: []const u8,
    headers: []const std.http.Header,
    client: *std.http.Client,
    allocator: std.mem.Allocator,
) !std.ArrayList(u8) {
    var response_body = std.ArrayList(u8).init(allocator);

    const response = try client.fetch(.{ .method = .POST, .location = .{ .url = url }, .extra_headers = headers, .response_storage = .{ .dynamic = &response_body }, .payload = "{\"jsonrpc\":\"2.0\",\"method\":\"web3_clientVersion\",\"params\":[],\"id\":1}" });

    try writer.print("Status: {d}\n", .{response.status});

    return response_body;
}
