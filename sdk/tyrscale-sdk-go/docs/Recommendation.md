# Recommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkName** | Pointer to **string** |  | [optional] 
**RouteUuid** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to [**StrategyName**](StrategyName.md) |  | [optional] 

## Methods

### NewRecommendation

`func NewRecommendation() *Recommendation`

NewRecommendation instantiates a new Recommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationWithDefaults

`func NewRecommendationWithDefaults() *Recommendation`

NewRecommendationWithDefaults instantiates a new Recommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkName

`func (o *Recommendation) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *Recommendation) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *Recommendation) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *Recommendation) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetRouteUuid

`func (o *Recommendation) GetRouteUuid() string`

GetRouteUuid returns the RouteUuid field if non-nil, zero value otherwise.

### GetRouteUuidOk

`func (o *Recommendation) GetRouteUuidOk() (*string, bool)`

GetRouteUuidOk returns a tuple with the RouteUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteUuid

`func (o *Recommendation) SetRouteUuid(v string)`

SetRouteUuid sets RouteUuid field to given value.

### HasRouteUuid

`func (o *Recommendation) HasRouteUuid() bool`

HasRouteUuid returns a boolean if a field has been set.

### GetSchedule

`func (o *Recommendation) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Recommendation) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Recommendation) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Recommendation) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStrategy

`func (o *Recommendation) GetStrategy() StrategyName`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Recommendation) GetStrategyOk() (*StrategyName, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *Recommendation) SetStrategy(v StrategyName)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *Recommendation) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


