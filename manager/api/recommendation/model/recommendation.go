package model

type StrategyName string

const (
	StrategyCustom           StrategyName = "strategy-custom"
	StrategyHighestBlock     StrategyName = "strategy-highest-block"
	StrategyAirUnderTheCurve StrategyName = "strategy-air-under-the-curve"
)

type Recommendation struct {
	UUID        string       `json:"uuid"`
	Schedule    string       `json:"schedule"`
	NetworkName string       `json:"network_name"`
	Strategy    StrategyName `json:"strategy"`
}
