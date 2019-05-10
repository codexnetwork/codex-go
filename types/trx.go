package types

import (
	"time"
)

// TransactionGeneralInfo Transaction data for each chain
type TransactionGeneralInfo struct {
	ID                 Checksum256
	Expiration         time.Time `json:"expiration"`
	RefBlockNum        uint16    `json:"ref_block_num"`
	RefBlockPrefix     uint32    `json:"ref_block_prefix"`
	MaxNetUsageWords   uint32    `json:"max_net_usage_words"`
	MaxCPUUsageMS      uint8     `json:"max_cpu_usage_ms"`
	DelaySec           uint32    `json:"delay_sec"`
	ContextFreeActions []*Action `json:"context_free_actions"`
	Actions            []*Action `json:"actions"`
	ContextFreeData    [][]byte  `json:"context_free_data"`
}
