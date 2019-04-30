package p2p

import (
	"time"

	"github.com/fanyang1988/force-go/types"
)

// ClientInterface interface for common p2p client
type ClientInterface interface {
	Type() types.ClientType
	Start() error
	CloseConnection() error
	SetReadTimeout(readTimeout time.Duration)
	RegHandler(handler types.P2PHandler)
}
