package types

// P2PHandler a general p2p msg handler interface for each chain
type P2PHandler interface {
	OnBlock(peer string, msg *BlockGeneralInfo) error
	OnGoAway(peer string, reason uint8, nodeID Checksum256) error
}
