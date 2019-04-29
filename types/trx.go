package types

import (
	"time"

	eosio "github.com/eoscanada/eos-go"
	forceio "github.com/eosforce/goforceio"
)

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

func (t *TransactionGeneralInfo) FromForceio(trx *forceio.TransactionWithID) error {
	t.ID = Checksum256(trx.ID)
	trxData, err := trx.Packed.Unpack()
	if err != nil {
		return err
	}

	t.Expiration = trxData.Expiration.Time
	t.RefBlockNum = trxData.RefBlockNum
	t.RefBlockPrefix = trxData.RefBlockPrefix
	t.MaxNetUsageWords = uint32(trxData.MaxNetUsageWords)
	t.MaxCPUUsageMS = trxData.MaxCPUUsageMS
	t.DelaySec = uint32(trxData.DelaySec)

	t.ContextFreeActions = make([]*Action, 0, len(trxData.ContextFreeActions))
	for _, a := range trxData.ContextFreeActions {
		act := &Action{}
		err := act.FromForceio(a)
		if err != nil {
			return err
		}

		t.ContextFreeActions = append(t.ContextFreeActions, act)
	}

	t.Actions = make([]*Action, 0, len(trxData.Actions))
	for _, a := range trxData.Actions {
		act := &Action{}
		err := act.FromForceio(a)
		if err != nil {
			return err
		}

		t.Actions = append(t.Actions, act)
	}

	t.ContextFreeData = make([][]byte, 0, len(trxData.ContextFreeData))
	for _, cd := range trxData.ContextFreeData {
		t.ContextFreeData = append(t.ContextFreeData, []byte(cd))
	}

	return nil
}

func (t *TransactionGeneralInfo) FromEOSIO(trx *eosio.TransactionWithID) error {
	t.ID = Checksum256(trx.ID)
	trxData, err := trx.Packed.Unpack()
	if err != nil {
		return err
	}

	t.Expiration = trxData.Expiration.Time
	t.RefBlockNum = trxData.RefBlockNum
	t.RefBlockPrefix = trxData.RefBlockPrefix
	t.MaxNetUsageWords = uint32(trxData.MaxNetUsageWords)
	t.MaxCPUUsageMS = trxData.MaxCPUUsageMS
	t.DelaySec = uint32(trxData.DelaySec)

	t.ContextFreeActions = make([]*Action, 0, len(trxData.ContextFreeActions))
	for _, a := range trxData.ContextFreeActions {
		act := &Action{}
		err := act.FromEOSIO(a)
		if err != nil {
			return err
		}

		t.ContextFreeActions = append(t.ContextFreeActions, act)
	}

	t.Actions = make([]*Action, 0, len(trxData.Actions))
	for _, a := range trxData.Actions {
		act := &Action{}
		err := act.FromEOSIO(a)
		if err != nil {
			return err
		}

		t.Actions = append(t.Actions, act)
	}

	t.ContextFreeData = make([][]byte, 0, len(trxData.ContextFreeData))
	for _, cd := range trxData.ContextFreeData {
		t.ContextFreeData = append(t.ContextFreeData, []byte(cd))
	}

	return nil
}
