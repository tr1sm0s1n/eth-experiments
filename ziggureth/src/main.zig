const std = @import("std");
const lib = @import("ziggureth_lib");
const writer = std.io.getStdOut().writer();

pub fn main() !void {
    const alloc = std.heap.page_allocator;
    var arena = std.heap.ArenaAllocator.init(alloc);
    const allocator = arena.allocator();

    defer arena.deinit();

    const args = try std.process.argsAlloc(allocator);
    if (args.len < 2) {
        return error.EmptyURL;
    }
    const rpc_url = args[1];

    var client = std.http.Client{
        .allocator = allocator,
    };

    var response = try lib.rpcCall(
        &client,
        alloc,
        rpc_url,
        "web3_clientVersion",
    );

    try writer.print("Client: \x1b[36m{s}\x1b[0m\n", .{response.value.result});

    response = try lib.rpcCall(
        &client,
        alloc,
        rpc_url,
        "net_version",
    );

    try writer.print("Network: \x1b[35m{s}\x1b[0m\n", .{response.value.result});

    response = try lib.rpcCall(
        &client,
        alloc,
        rpc_url,
        "eth_chainId",
    );

    try writer.print("Chain ID: \x1b[34m{s}\x1b[0m\n", .{response.value.result});

    response = try lib.rpcCall(
        &client,
        alloc,
        rpc_url,
        "eth_blockNumber",
    );

    try writer.print("Block Number: \x1b[34m{s}\x1b[0m\n", .{response.value.result});

    const header = try lib.latestHeader(
        &client,
        alloc,
        rpc_url,
    );

    var header_string = std.ArrayList(u8).init(alloc);
    try std.json.stringify(header, .{}, header_string.writer());

    try writer.print("Latest Header: \x1b[32m{s}\x1b[0m\n", .{header_string.items});
}
