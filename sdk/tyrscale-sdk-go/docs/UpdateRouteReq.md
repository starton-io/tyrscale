# UpdateRouteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitBreaker** | Pointer to [**CircuitbreakerSettings**](CircuitbreakerSettings.md) |  | [optional] 
**HealthCheck** | Pointer to [**HealthcheckHealthCheckConfig**](HealthcheckHealthCheckConfig.md) |  | [optional] 
**Host** | Pointer to **string** | Uuid                 string                         &#x60;json:\&quot;uuid\&quot; validate:\&quot;required,uuid\&quot;&#x60; | [optional] 
**LoadBalancerStrategy** | Pointer to [**BalancerLoadBalancerStrategy**](BalancerLoadBalancerStrategy.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateRouteReq

`func NewUpdateRouteReq() *UpdateRouteReq`

NewUpdateRouteReq instantiates a new UpdateRouteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRouteReqWithDefaults

`func NewUpdateRouteReqWithDefaults() *UpdateRouteReq`

NewUpdateRouteReqWithDefaults instantiates a new UpdateRouteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitBreaker

`func (o *UpdateRouteReq) GetCircuitBreaker() CircuitbreakerSettings`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *UpdateRouteReq) GetCircuitBreakerOk() (*CircuitbreakerSettings, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *UpdateRouteReq) SetCircuitBreaker(v CircuitbreakerSettings)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *UpdateRouteReq) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### GetHealthCheck

`func (o *UpdateRouteReq) GetHealthCheck() HealthcheckHealthCheckConfig`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *UpdateRouteReq) GetHealthCheckOk() (*HealthcheckHealthCheckConfig, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *UpdateRouteReq) SetHealthCheck(v HealthcheckHealthCheckConfig)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *UpdateRouteReq) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetHost

`func (o *UpdateRouteReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateRouteReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateRouteReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateRouteReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetLoadBalancerStrategy

`func (o *UpdateRouteReq) GetLoadBalancerStrategy() BalancerLoadBalancerStrategy`

GetLoadBalancerStrategy returns the LoadBalancerStrategy field if non-nil, zero value otherwise.

### GetLoadBalancerStrategyOk

`func (o *UpdateRouteReq) GetLoadBalancerStrategyOk() (*BalancerLoadBalancerStrategy, bool)`

GetLoadBalancerStrategyOk returns a tuple with the LoadBalancerStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerStrategy

`func (o *UpdateRouteReq) SetLoadBalancerStrategy(v BalancerLoadBalancerStrategy)`

SetLoadBalancerStrategy sets LoadBalancerStrategy field to given value.

### HasLoadBalancerStrategy

`func (o *UpdateRouteReq) HasLoadBalancerStrategy() bool`

HasLoadBalancerStrategy returns a boolean if a field has been set.

### GetPath

`func (o *UpdateRouteReq) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateRouteReq) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateRouteReq) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateRouteReq) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


