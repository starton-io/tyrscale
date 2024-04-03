# UpdateRecommendationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkName** | **string** |  | 
**RouteUuid** | **string** |  | 
**Schedule** | **string** |  | 
**Strategy** | [**StrategyName**](StrategyName.md) |  | 

## Methods

### NewUpdateRecommendationReq

`func NewUpdateRecommendationReq(networkName string, routeUuid string, schedule string, strategy StrategyName, ) *UpdateRecommendationReq`

NewUpdateRecommendationReq instantiates a new UpdateRecommendationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRecommendationReqWithDefaults

`func NewUpdateRecommendationReqWithDefaults() *UpdateRecommendationReq`

NewUpdateRecommendationReqWithDefaults instantiates a new UpdateRecommendationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkName

`func (o *UpdateRecommendationReq) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *UpdateRecommendationReq) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *UpdateRecommendationReq) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetRouteUuid

`func (o *UpdateRecommendationReq) GetRouteUuid() string`

GetRouteUuid returns the RouteUuid field if non-nil, zero value otherwise.

### GetRouteUuidOk

`func (o *UpdateRecommendationReq) GetRouteUuidOk() (*string, bool)`

GetRouteUuidOk returns a tuple with the RouteUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteUuid

`func (o *UpdateRecommendationReq) SetRouteUuid(v string)`

SetRouteUuid sets RouteUuid field to given value.


### GetSchedule

`func (o *UpdateRecommendationReq) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateRecommendationReq) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateRecommendationReq) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetStrategy

`func (o *UpdateRecommendationReq) GetStrategy() StrategyName`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *UpdateRecommendationReq) GetStrategyOk() (*StrategyName, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *UpdateRecommendationReq) SetStrategy(v StrategyName)`

SetStrategy sets Strategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


