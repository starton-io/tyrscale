# CreateRecommendationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkName** | **string** |  | 
**RouteUuid** | **string** |  | 
**Schedule** | **string** |  | 
**Strategy** | [**StrategyName**](StrategyName.md) |  | 

## Methods

### NewCreateRecommendationReq

`func NewCreateRecommendationReq(networkName string, routeUuid string, schedule string, strategy StrategyName, ) *CreateRecommendationReq`

NewCreateRecommendationReq instantiates a new CreateRecommendationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecommendationReqWithDefaults

`func NewCreateRecommendationReqWithDefaults() *CreateRecommendationReq`

NewCreateRecommendationReqWithDefaults instantiates a new CreateRecommendationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkName

`func (o *CreateRecommendationReq) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *CreateRecommendationReq) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *CreateRecommendationReq) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetRouteUuid

`func (o *CreateRecommendationReq) GetRouteUuid() string`

GetRouteUuid returns the RouteUuid field if non-nil, zero value otherwise.

### GetRouteUuidOk

`func (o *CreateRecommendationReq) GetRouteUuidOk() (*string, bool)`

GetRouteUuidOk returns a tuple with the RouteUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteUuid

`func (o *CreateRecommendationReq) SetRouteUuid(v string)`

SetRouteUuid sets RouteUuid field to given value.


### GetSchedule

`func (o *CreateRecommendationReq) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateRecommendationReq) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateRecommendationReq) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetStrategy

`func (o *CreateRecommendationReq) GetStrategy() StrategyName`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *CreateRecommendationReq) GetStrategyOk() (*StrategyName, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *CreateRecommendationReq) SetStrategy(v StrategyName)`

SetStrategy sets Strategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


