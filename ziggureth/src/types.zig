pub const SimpleResponse = struct {
    result: []u8,
};

pub const Header = struct {
    hash: []u8,
    parentHash: []u8,
    sha3Uncles: []u8,
    miner: []u8,
    stateRoot: []u8,
    transactionsRoot: []u8,
    receiptsRoot: []u8,
    logsBloom: []u8,
    difficulty: []u8,
    number: []u8,
    gasLimit: []u8,
    gasUsed: []u8,
    timestamp: []u8,
    extraData: []u8,
    mixHash: []u8,
    nonce: []u8,
    baseFeePerGas: []u8,
    withdrawalsRoot: []u8,
    blobGasUsed: []u8,
    excessBlobGas: []u8,
    parentBeaconBlockRoot: []u8,
};
