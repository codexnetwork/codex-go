package types

import (
	chain "github.com/eoscanada/eos-go"
)

type switcher2EOSIO struct {
}

func (s switcher2EOSIO) Type() ClientType {
	return EOSIO
}

func (s switcher2EOSIO) NameFromCommon(n string) interface{} {
	return chain.Name(n)
}

func (s switcher2EOSIO) Checksum256FromCommon(c Checksum256) interface{} {
	return chain.Checksum256(c)
}

func (s switcher2EOSIO) PushTransactionFullRespToCommon(r interface{}) (*PushTransactionFullResp, error) {
	p := &PushTransactionFullResp{}

	rsp, ok := r.(*chain.PushTransactionFullResp)
	if !ok {
		return nil, ErrTypeErrToChain
	}

	p.StatusCode = rsp.StatusCode
	p.TransactionID = rsp.TransactionID
	p.BlockID = rsp.BlockID
	p.BlockNum = rsp.BlockNum

	return p, p.FillProcessedDatas(rsp.Processed)
}

func (s switcher2EOSIO) InfoRespToCommon(r interface{}) (*InfoResp, error) {
	i := &InfoResp{}

	info, ok := r.(*chain.InfoResp)
	if !ok {
		return nil, ErrTypeErrToChain
	}

	i.ServerVersion = info.ServerVersion
	i.ChainID = Checksum256(info.ChainID)
	i.HeadBlockNum = info.HeadBlockNum
	i.LastIrreversibleBlockNum = info.LastIrreversibleBlockNum
	i.LastIrreversibleBlockID = Checksum256(info.LastIrreversibleBlockID)
	i.HeadBlockID = Checksum256(info.HeadBlockID)
	i.HeadBlockTime = info.HeadBlockTime.Time
	i.HeadBlockProducer = string(info.HeadBlockProducer)
	i.VirtualBlockCPULimit = int64(info.VirtualBlockCPULimit)
	i.VirtualBlockNetLimit = int64(info.VirtualBlockNetLimit)
	i.BlockCPULimit = int64(info.BlockCPULimit)
	i.BlockNetLimit = int64(info.BlockNetLimit)
	i.ServerVersionString = info.ServerVersionString

	return i, nil
}

func (s switcher2EOSIO) ActionToCommon(d interface{}) (*Action, error) {
	res := &Action{}

	r, ok := d.(*chain.Action)
	if !ok {
		return nil, ErrTypeErrToChain
	}

	return res, res.FromEOSIO(r)
}

func (s switcher2EOSIO) ActionFromCommon(d *Action) (interface{}, error) {
	return d.ToEOSIO()
}

func (s switcher2EOSIO) TransactionToCommon(r interface{}) (*TransactionGeneralInfo, error) {
	t := &TransactionGeneralInfo{}

	trx, ok := r.(*chain.TransactionWithID)
	if !ok {
		return nil, ErrTypeErrToChain
	}

	t.ID = Checksum256(trx.ID)
	trxData, err := trx.Packed.Unpack()
	if err != nil {
		return nil, err
	}

	t.Expiration = trxData.Expiration.Time
	t.RefBlockNum = trxData.RefBlockNum
	t.RefBlockPrefix = trxData.RefBlockPrefix
	t.MaxNetUsageWords = uint32(trxData.MaxNetUsageWords)
	t.MaxCPUUsageMS = trxData.MaxCPUUsageMS
	t.DelaySec = uint32(trxData.DelaySec)

	t.ContextFreeActions = make([]*Action, 0, len(trxData.ContextFreeActions))
	for _, a := range trxData.ContextFreeActions {
		act, err := s.ActionToCommon(a)
		if err != nil {
			return nil, err
		}

		t.ContextFreeActions = append(t.ContextFreeActions, act)
	}

	t.Actions = make([]*Action, 0, len(trxData.Actions))
	for _, a := range trxData.Actions {
		act, err := s.ActionToCommon(a)
		if err != nil {
			return nil, err
		}

		t.Actions = append(t.Actions, act)
	}

	t.ContextFreeData = make([][]byte, 0, len(trxData.ContextFreeData))
	for _, cd := range trxData.ContextFreeData {
		t.ContextFreeData = append(t.ContextFreeData, []byte(cd))
	}

	return t, nil
}

func (s switcher2EOSIO) BlockToCommon(r interface{}) (*BlockGeneralInfo, error) {
	b := &BlockGeneralInfo{}

	block, ok := r.(*chain.SignedBlock)
	if !ok {
		return nil, ErrTypeErrToChain
	}

	id, _ := block.BlockID()

	b.ID = Checksum256(id)
	b.BlockNum = block.BlockNumber()
	b.Timestamp = block.Timestamp.Time
	b.Producer = string(block.Producer)
	b.Confirmed = block.Confirmed
	b.Previous = Checksum256(block.Previous)
	b.TransactionMRoot = Checksum256(block.TransactionMRoot)
	b.ActionMRoot = Checksum256(block.ActionMRoot)
	b.ScheduleVersion = block.ScheduleVersion

	b.Transactions = make([]TransactionReceipt, 0, len(block.Transactions))
	for _, trx := range block.Transactions {
		t := &TransactionGeneralInfo{}
		t, err := s.TransactionToCommon(&trx.Transaction)
		if err != nil {
			return nil, err
		}
		b.Transactions = append(b.Transactions, TransactionReceipt{
			Status:               TransactionStatus(trx.Status),
			CPUUsageMicroSeconds: trx.CPUUsageMicroSeconds,
			NetUsageWords:        uint32(trx.NetUsageWords),
			Transaction:          *t,
		})
	}

	return b, nil
}

func (s switcher2EOSIO) BlockRspToCommon(r interface{}) (*BlockResp, error) {
	b := &BlockResp{}

	block, ok := r.(*chain.BlockResp)
	if !ok {
		return nil, ErrTypeErrToChain
	}

	blockInRsp, err := s.BlockToCommon(&block.SignedBlock)
	if err != nil {
		return nil, err
	}

	b.BlockGeneralInfo = *blockInRsp

	b.ID = Checksum256(block.ID)
	b.BlockNum = block.BlockNum
	b.RefBlockPrefix = block.RefBlockPrefix

	return b, nil
}
