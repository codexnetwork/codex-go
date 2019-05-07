package eosio

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/fanyang1988/force-go/config"
	"github.com/fanyang1988/force-go/types"
	"github.com/pkg/errors"
)

// API client api to forceio chain
type API struct {
	*eos.API
	Cfg      Config
	switcher types.SwitcherInterface
	typ      types.ClientType
}

func (api *API) Init(cfg *config.ConfigData) error {
	api.API = eos.New(cfg.URL)
	api.typ = types.EOSIO
	api.switcher = types.NewSwitcherInterface(api.typ)

	err := api.Cfg.Parse(cfg)
	if err != nil {
		return err
	}
	api.API.Debug = api.Cfg.IsDebug

	err = api.checkChainID()
	if err != nil {
		return err
	}
	api.SetSigner(NewConfigSigner(api))
	return nil
}

func (api *API) checkChainID() error {
	res, err := api.GetInfo()
	if err != nil {
		return errors.Wrapf(err, "check chain id to get err")
	}

	if api.Cfg.ChainID != nil && api.Cfg.ChainID.String() != res.ChainID.String() {
		return errors.Errorf("chain id diff from cfg by %s <-> %s",
			api.Cfg.ChainID.String(), res.ChainID.String())
	}

	api.Cfg.ChainID = res.ChainID
	return nil
}

func (api *API) PushActions(actions ...*types.Action) (*types.PushTransactionFullResp, error) {
	acts := make([]*eos.Action, 0, len(actions))
	for _, act := range actions {
		a, err := act.ToEOSIO()
		if err != nil {
			return nil, err
		}
		acts = append(acts, a)
	}
	rsp, err := PushActions(api, acts...)
	if err != nil {
		return nil, err
	}

	res, err := api.switcher.PushTransactionFullRespToCommon(rsp)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (api *API) GetInfoData() (*types.InfoResp, error) {
	rsp, err := api.GetInfo()
	if err != nil {
		return nil, err
	}

	res, err := api.switcher.InfoRespToCommon(rsp)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (api *API) GetBlockDataByID(id string) (*types.BlockGeneralInfo, error) {
	rsp, err := api.GetBlockByID(id)
	if err != nil {
		return nil, err
	}

	res, err := api.switcher.BlockRspToCommon(rsp)
	if err != nil {
		return nil, err
	}
	// fix id err
	res.BlockGeneralInfo.ID = types.Checksum256(rsp.ID)
	res.BlockGeneralInfo.BlockNum = rsp.BlockNum
	return &res.BlockGeneralInfo, err
}
func (api *API) GetBlockDataByNum(num uint32) (*types.BlockGeneralInfo, error) {
	rsp, err := api.GetBlockByNum(num)
	if err != nil {
		return nil, err
	}

	res, err := api.switcher.BlockRspToCommon(rsp)
	if err != nil {
		return nil, err
	}
	// fix id err
	res.BlockGeneralInfo.ID = types.Checksum256(rsp.ID)
	res.BlockGeneralInfo.BlockNum = rsp.BlockNum
	return &res.BlockGeneralInfo, err
}

func (api *API) Name(n string) interface{} {
	return eos.Name(n)
}

func (api *API) Asset(a *types.Asset) interface{} {
	return eos.Asset{
		Amount: eos.Int64(a.Amount),
		Symbol: eos.Symbol{
			Precision: a.Symbol.Precision,
			Symbol:    a.Symbol.Symbol,
		},
	}
}
