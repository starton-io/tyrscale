# ListRecommendationRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Recommendation**](Recommendation.md) |  | [optional] 

## Methods

### NewListRecommendationRes

`func NewListRecommendationRes() *ListRecommendationRes`

NewListRecommendationRes instantiates a new ListRecommendationRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRecommendationResWithDefaults

`func NewListRecommendationResWithDefaults() *ListRecommendationRes`

NewListRecommendationResWithDefaults instantiates a new ListRecommendationRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListRecommendationRes) GetItems() []Recommendation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListRecommendationRes) GetItemsOk() (*[]Recommendation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListRecommendationRes) SetItems(v []Recommendation)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListRecommendationRes) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


