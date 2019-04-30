package types

import (
	"time"
)

type BlockGeneralInfo struct {
	ID               Checksum256          `json:"id"`
	BlockNum         uint32               `json:"block_num"`
	Timestamp        time.Time            `json:"timestamp"`
	Producer         string               `json:"producer"`
	Confirmed        uint16               `json:"confirmed"`
	Previous         Checksum256          `json:"previous"`
	TransactionMRoot Checksum256          `json:"transaction_mroot"`
	ActionMRoot      Checksum256          `json:"action_mroot"`
	ScheduleVersion  uint32               `json:"schedule_version"`
	Transactions     []TransactionReceipt `json:"transactions"`
}

type TransactionStatus uint8

const (
	TransactionStatusExecuted TransactionStatus = iota ///< succeed, no error handler executed
	TransactionStatusSoftFail                          ///< objectively failed (not executed), error handler executed
	TransactionStatusHardFail                          ///< objectively failed and error handler objectively failed thus no state change
	TransactionStatusDelayed                           ///< transaction delayed
	TransactionStatusExpired                           ///< transaction expired
	TransactionStatusUnknown  = TransactionStatus(255)
)

type TransactionReceipt struct {
	Status               TransactionStatus      `json:"status"`
	CPUUsageMicroSeconds uint32                 `json:"cpu_usage_us"`
	NetUsageWords        uint32                 `json:"net_usage_words"`
	Transaction          TransactionGeneralInfo `json:"trx"`
}
