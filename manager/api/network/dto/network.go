package dto

type Network struct {
	ChainId    int    `json:"chain_id" validate:"required,gt=0" example:"1"`
	Name       string `json:"name" validate:"required,regexp=^[a-zA-Z0-9_-]+$" example:"eth-mainnet"`
	Blockchain string `json:"blockchain" validate:"required,regexp=^[a-zA-Z0-9_-]+$" example:"ethereum"`
}

type CreateNetworkRes struct {
	ChainId int `json:"chain_id"`
}

type DeleteNetworkRes struct {
	Name string `json:"name"`
}

type ListNetworkReq struct {
	ChainId    string `query:"chain_id" validate:"omitempty,numeric"`
	Name       string `query:"name" validate:"omitempty,regexp=^[a-zA-Z0-9_-]+$"`
	Blockchain string `query:"blockchain" validate:"omitempty,regexp=^[a-zA-Z0-9_-]+$"`
}

type ListNetworkRes struct {
	Networks []Network `json:"items"`
}
