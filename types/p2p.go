package types

type P2PHandler interface {
	OnBlock(peer string, msg *BlockGeneralInfo) error
	OnGoAway(peer string, reason uint8, nodeID Checksum256) error
}
