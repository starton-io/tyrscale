package dto

import (
	"errors"
)

type RPCType string

const (
	RPCTypePrivate RPCType = "private"
	RPCTypePublic  RPCType = "public"
)

var rpcTypeValues = []RPCType{
	RPCTypePrivate,
	RPCTypePublic,
}

func (r RPCType) String() string {
	return string(r)
}

func (r RPCType) Validate() error {
	for _, value := range rpcTypeValues {
		if r == value {
			return nil
		}
	}
	return errors.New("invalid type")
}

// Create a rpc request to create a rpc endpoints
type CreateRpcReq struct {
	UUID        string   `json:"uuid" validate:"omitempty,uuid"`
	URL         string   `json:"url" validate:"required,http_url"`
	Type        RPCType  `json:"type" validate:"required"`
	Provider    string   `json:"provider" validate:"required,regexp=^[a-zA-Z0-9_-]+$"`
	NetworkName string   `json:"network_name" validate:"required"`
	Collectors  []string `json:"collectors" validate:"gte=1,dive,oneof=eth_block_number net_peers"`
}

type CreateRpcRes struct {
	UUID string `json:"uuid"`
}

type CreateRpcCtx struct {
	UUID string `json:"uuid"`
}

type ListReq struct {
	ListFilterReq
	*ListSortReq
}

type ListFilterReq struct {
	UUID        string `query:"uuid" validate:"omitempty,uuid"`
	ChainID     string `query:"chain_id" validate:"omitempty,numeric"`
	Type        string `query:"type" validate:"omitempty"`
	Provider    string `query:"provider" validate:"omitempty,regexp=^[a-zA-Z0-9_-]+$"`
	NetworkName string `query:"network_name" validate:"omitempty,regexp=^[a-zA-Z0-9_-]+$"`
	URL         string `query:"url" validate:"omitempty"`
}

type ListSortReq struct {
	SortBy         string `query:"sort_by" validate:"omitempty,oneof=uuid type provider network_name"`
	SortDescending bool   `query:"sort_descending" validate:"omitempty"`
}

type Rpc struct {
	UUID        string   `json:"uuid" validate:"omitempty,uuid"`
	ChainId     int64    `json:"chain_id"`
	URL         string   `json:"url" validate:"required,http_url"`
	Type        RPCType  `json:"type" validate:"required"`
	Provider    string   `json:"provider" validate:"required,regexp=^[a-zA-Z0-9_-]+$"`
	NetworkName string   `json:"network_name" validate:"required"`
	Collectors  []string `json:"collectors" validate:"gte=1,dive,oneof=eth_block_number net_peers"`
}

type ListRpcRes struct {
	RPCs []Rpc `json:"items"`
}

type DeleteRpcOptReq struct {
	CascadeDeleteUpstream *bool `json:"cascade_delete_upstream,omitempty"`
}

type DeleteRpcReq struct {
	UUID                  string
	CascadeDeleteUpstream bool
}

type UpdateRpcReq struct {
	UUID       string   `json:"uuid" validate:"required"`
	URL        string   `json:"url" validate:"omitempty,http_url"`
	Type       *RPCType `json:"type" validate:"omitempty"`
	Provider   string   `json:"provider" validate:"omitempty"`
	Collectors []string `json:"collectors" validate:"omitempty,gte=1,dive,oneof=eth_block_number net_peers"`
}
