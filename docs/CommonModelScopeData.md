# CommonModelScopeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonModels** | [**[]CommonModelScopesDisabledModels**](CommonModelScopesDisabledModels.md) |  | 
**LinkedAccountId** | Pointer to **string** |  | [optional] 

## Methods

### NewCommonModelScopeData

`func NewCommonModelScopeData(commonModels []CommonModelScopesDisabledModels, ) *CommonModelScopeData`

NewCommonModelScopeData instantiates a new CommonModelScopeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonModelScopeDataWithDefaults

`func NewCommonModelScopeDataWithDefaults() *CommonModelScopeData`

NewCommonModelScopeDataWithDefaults instantiates a new CommonModelScopeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonModels

`func (o *CommonModelScopeData) GetCommonModels() []CommonModelScopesDisabledModels`

GetCommonModels returns the CommonModels field if non-nil, zero value otherwise.

### GetCommonModelsOk

`func (o *CommonModelScopeData) GetCommonModelsOk() (*[]CommonModelScopesDisabledModels, bool)`

GetCommonModelsOk returns a tuple with the CommonModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonModels

`func (o *CommonModelScopeData) SetCommonModels(v []CommonModelScopesDisabledModels)`

SetCommonModels sets CommonModels field to given value.


### GetLinkedAccountId

`func (o *CommonModelScopeData) GetLinkedAccountId() string`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *CommonModelScopeData) GetLinkedAccountIdOk() (*string, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *CommonModelScopeData) SetLinkedAccountId(v string)`

SetLinkedAccountId sets LinkedAccountId field to given value.

### HasLinkedAccountId

`func (o *CommonModelScopeData) HasLinkedAccountId() bool`

HasLinkedAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


