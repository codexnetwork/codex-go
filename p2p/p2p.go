package p2p

import (
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/codexnetwork/codex-go/types"
)

// ClientInterface interface for common p2p client
type ClientInterface interface {
	Type() types.ClientType
	Start() error
	CloseConnection() error
	IsClosed() bool
	SetReadTimeout(readTimeout time.Duration)
	RegHandler(handler types.P2PHandler)
}

// P2PInitParams init params for P2P client
type P2PInitParams struct {
	Name          string   `json:"name"`
	ClientID      string   `json:"clientID"`
	Peers         []string `json:"peers"`
	StartBlockNum uint32   `json:"start"`
	Logger        *zap.Logger
}

func NewP2PClient(typ types.ClientType, params P2PInitParams) ClientInterface {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	switch typ {
	case types.EOSForce:
		return NewP2PClient4EOSForce(params.Name, params.ClientID, params.StartBlockNum, params.Peers, params.Logger)
	case types.FORCEIO:
		return NewP2PClient4Forceio(params.Name, params.ClientID, params.StartBlockNum, params.Peers, params.Logger)
	case types.EOSIO:
		return NewP2PClient4EOS(params.Name, params.ClientID, params.StartBlockNum, params.Peers, params.Logger)
	case types.ENU:
		return NewP2PClient4EOS(params.Name, params.ClientID, params.StartBlockNum, params.Peers, params.Logger)
	default:
		panic(errors.New("no support type for p2p"))
	}
}
