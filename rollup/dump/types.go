package dump

import (
	"github.com/ledgerwatch/turbo-geth/accounts/abi"
	"github.com/ledgerwatch/turbo-geth/common"
)

type OvmDumpAccount struct {
	Address  common.Address         `json:"address"`
	Code     string                 `json:"code"`
	CodeHash string                 `json:"codeHash"`
	Storage  map[common.Hash]string `json:"storage"`
	ABI      abi.ABI                `json:"abi"`
	Nonce    uint64                 `json:"nonce"`
}
