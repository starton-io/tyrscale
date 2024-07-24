# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitBreaker** | Pointer to [**CircuitbreakerSettings**](CircuitbreakerSettings.md) |  | [optional] 
**HealthCheck** | Pointer to [**HealthcheckHealthCheckConfig**](HealthcheckHealthCheckConfig.md) |  | [optional] 
**Host** | **string** |  | 
**LoadBalancerStrategy** | [**BalancerLoadBalancerStrategy**](BalancerLoadBalancerStrategy.md) |  | 
**Path** | Pointer to **string** |  | [optional] 
**Plugins** | Pointer to [**Plugins**](Plugins.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewRoute

`func NewRoute(host string, loadBalancerStrategy BalancerLoadBalancerStrategy, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitBreaker

`func (o *Route) GetCircuitBreaker() CircuitbreakerSettings`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *Route) GetCircuitBreakerOk() (*CircuitbreakerSettings, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *Route) SetCircuitBreaker(v CircuitbreakerSettings)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *Route) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### GetHealthCheck

`func (o *Route) GetHealthCheck() HealthcheckHealthCheckConfig`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *Route) GetHealthCheckOk() (*HealthcheckHealthCheckConfig, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *Route) SetHealthCheck(v HealthcheckHealthCheckConfig)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *Route) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetHost

`func (o *Route) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Route) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Route) SetHost(v string)`

SetHost sets Host field to given value.


### GetLoadBalancerStrategy

`func (o *Route) GetLoadBalancerStrategy() BalancerLoadBalancerStrategy`

GetLoadBalancerStrategy returns the LoadBalancerStrategy field if non-nil, zero value otherwise.

### GetLoadBalancerStrategyOk

`func (o *Route) GetLoadBalancerStrategyOk() (*BalancerLoadBalancerStrategy, bool)`

GetLoadBalancerStrategyOk returns a tuple with the LoadBalancerStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerStrategy

`func (o *Route) SetLoadBalancerStrategy(v BalancerLoadBalancerStrategy)`

SetLoadBalancerStrategy sets LoadBalancerStrategy field to given value.


### GetPath

`func (o *Route) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Route) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Route) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Route) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPlugins

`func (o *Route) GetPlugins() Plugins`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *Route) GetPluginsOk() (*Plugins, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *Route) SetPlugins(v Plugins)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *Route) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetUuid

`func (o *Route) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Route) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Route) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Route) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


