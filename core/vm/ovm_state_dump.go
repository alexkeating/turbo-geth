package vm

import "github.com/ledgerwatch/turbo-geth/common"

// AbiBytesTrue represents the ABI encoding of "true" as a byte slice
var AbiBytesTrue = common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000001")

// AbiBytesFalse represents the ABI encoding of "false" as a byte slice
var AbiBytesFalse = common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000")

var UsingOVM bool
