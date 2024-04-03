package model

type Network struct {
	ChainID    int    `json:"chain_id" validate:"required,gt=0"`
	Name       string `json:"name" validate:"required"`
	Blockchain string `json:"blockchain" validate:"required"`
}
