# CreateRouteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitBreaker** | Pointer to [**CircuitbreakerSettings**](CircuitbreakerSettings.md) |  | [optional] 
**HealthCheck** | Pointer to [**HealthcheckHealthCheckConfig**](HealthcheckHealthCheckConfig.md) |  | [optional] 
**Host** | **string** |  | 
**LoadBalancerStrategy** | [**BalancerLoadBalancerStrategy**](BalancerLoadBalancerStrategy.md) |  | 
**Path** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRouteReq

`func NewCreateRouteReq(host string, loadBalancerStrategy BalancerLoadBalancerStrategy, ) *CreateRouteReq`

NewCreateRouteReq instantiates a new CreateRouteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteReqWithDefaults

`func NewCreateRouteReqWithDefaults() *CreateRouteReq`

NewCreateRouteReqWithDefaults instantiates a new CreateRouteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitBreaker

`func (o *CreateRouteReq) GetCircuitBreaker() CircuitbreakerSettings`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *CreateRouteReq) GetCircuitBreakerOk() (*CircuitbreakerSettings, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *CreateRouteReq) SetCircuitBreaker(v CircuitbreakerSettings)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *CreateRouteReq) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### GetHealthCheck

`func (o *CreateRouteReq) GetHealthCheck() HealthcheckHealthCheckConfig`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *CreateRouteReq) GetHealthCheckOk() (*HealthcheckHealthCheckConfig, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *CreateRouteReq) SetHealthCheck(v HealthcheckHealthCheckConfig)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *CreateRouteReq) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetHost

`func (o *CreateRouteReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateRouteReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateRouteReq) SetHost(v string)`

SetHost sets Host field to given value.


### GetLoadBalancerStrategy

`func (o *CreateRouteReq) GetLoadBalancerStrategy() BalancerLoadBalancerStrategy`

GetLoadBalancerStrategy returns the LoadBalancerStrategy field if non-nil, zero value otherwise.

### GetLoadBalancerStrategyOk

`func (o *CreateRouteReq) GetLoadBalancerStrategyOk() (*BalancerLoadBalancerStrategy, bool)`

GetLoadBalancerStrategyOk returns a tuple with the LoadBalancerStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerStrategy

`func (o *CreateRouteReq) SetLoadBalancerStrategy(v BalancerLoadBalancerStrategy)`

SetLoadBalancerStrategy sets LoadBalancerStrategy field to given value.


### GetPath

`func (o *CreateRouteReq) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateRouteReq) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateRouteReq) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateRouteReq) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUuid

`func (o *CreateRouteReq) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateRouteReq) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateRouteReq) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CreateRouteReq) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


