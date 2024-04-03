package model

type CollectorName string

const (
	CollectorEthBlockNumber CollectorName = "eth_block_number"
	CollectorNetPeerCount   CollectorName = "net_peer_count"
)

type RPCType string

const (
	RPCTypePrivate RPCType = "private"
	RPCTypePublic  RPCType = "public"
)

type RPC struct {
	UUID        string   `json:"uuid"`
	ChainID     int      `json:"chain_id"`
	NetworkName string   `json:"network_name"`
	URL         string   `json:"url"`
	Type        RPCType  `json:"type"`
	Provider    string   `json:"provider"`
	Collectors  []string `json:"collectors"`
}
