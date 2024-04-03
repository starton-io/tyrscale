package dto

import (
	"errors"
)

type StrategyName string

const (
	StrategyCustom           StrategyName = "STRATEGY_CUSTOM"
	StrategyHighestBlock     StrategyName = "STRATEGY_HIGHEST_BLOCK"
	StrategyAirUnderTheCurve StrategyName = "STRATEGY_AIR_UNDER_THE_CURVE"
)

var strategyValues = []StrategyName{
	StrategyCustom,
	StrategyHighestBlock,
	StrategyAirUnderTheCurve,
}

func (s StrategyName) String() string {
	return string(s)
}
func (s StrategyName) Validate() error {
	for _, value := range strategyValues {
		if s == value {
			return nil
		}
	}
	return errors.New("invalid strategy")
}

type CreateRecommendationReq struct {
	RouteUuid   string       `json:"route_uuid" validate:"required,uuid"`
	Schedule    string       `json:"schedule" validate:"required,cron"`
	NetworkName string       `json:"network_name" validate:"required,regexp=^[a-zA-Z0-9-]+$"`
	Strategy    StrategyName `json:"strategy" validate:"required"`
}

type CreateRecommendationRes struct {
	RouteUuid string `json:"route_uuid"`
}

type UpdateRecommendationReq struct {
	RouteUuid   string       `json:"route_uuid" validate:"required,uuid"`
	Schedule    string       `json:"schedule" validate:"required,cron"`
	NetworkName string       `json:"network_name" validate:"required,regexp=^[a-zA-Z0-9-]+$"`
	Strategy    StrategyName `json:"strategy" validate:"required"`
}

type ListRecommendationReq struct {
	NetworkName string `query:"network_name" validate:"omitempty,regexp=^[a-zA-Z0-9-]+$"`
	RouteUuid   string `query:"route_uuid" validate:"omitempty,uuid"`
	Strategy    string `query:"strategy" validate:"omitempty"`
}

type Recommendation struct {
	RouteUuid   string       `json:"route_uuid"`
	Schedule    string       `json:"schedule"`
	NetworkName string       `json:"network_name"`
	Strategy    StrategyName `json:"strategy"`
}

type ListRecommendationRes struct {
	Recommendations []Recommendation `json:"items"`
}

type DeleteRecommendationReq struct {
	RouteUuid string `query:"route_uuid" validate:"required,uuid"`
}
