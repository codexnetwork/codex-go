package p2p

import (
	"errors"
	"time"

	"github.com/fanyang1988/force-go/types"
	"go.uber.org/zap"
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

type P2PSyncData struct {
	HeadBlockNum             uint32
	HeadBlockID              types.Checksum256
	HeadBlockTime            time.Time
	LastIrreversibleBlockNum uint32
	LastIrreversibleBlockID  types.Checksum256
}

type P2PInitParams struct {
	Name       string `json:"name"`
	ClientID   string `json:"clientID"`
	Peers      []string
	StartBlock *P2PSyncData
	Logger     *zap.Logger
}

func NewP2PClient(typ types.ClientType, params P2PInitParams) ClientInterface {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	switch typ {
	case types.EOSForce:
		return NewP2PClient4EOSForce(params.Name, params.ClientID, params.StartBlock, params.Peers, params.Logger)
	case types.FORCEIO:
		return NewP2PClient4Forceio(params.Name, params.ClientID, params.StartBlock, params.Peers, params.Logger)
	case types.EOSIO:
		return NewP2PClient4EOS(params.Name, params.ClientID, params.StartBlock, params.Peers, params.Logger)
	default:
		panic(errors.New("no support type for p2p"))
	}
}
