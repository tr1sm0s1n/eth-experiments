const std = @import("std");
const lib = @import("ziggureth_lib");
const writer = std.io.getStdOut().writer();

pub fn main() !void {
    const alloc = std.heap.page_allocator;
    var arena = std.heap.ArenaAllocator.init(alloc);
    const allocator = arena.allocator();

    defer arena.deinit();

    const rpc_url = "http://127.0.0.1:8545";

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
}
