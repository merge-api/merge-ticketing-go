# CommonModelScopesPostInnerDeserializerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | **string** |  | 
**EnabledActions** | [**[]EnabledActionsEnum**](EnabledActionsEnum.md) |  | 
**DisabledFields** | **[]string** |  | 

## Methods

### NewCommonModelScopesPostInnerDeserializerRequest

`func NewCommonModelScopesPostInnerDeserializerRequest(modelId string, enabledActions []EnabledActionsEnum, disabledFields []string, ) *CommonModelScopesPostInnerDeserializerRequest`

NewCommonModelScopesPostInnerDeserializerRequest instantiates a new CommonModelScopesPostInnerDeserializerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonModelScopesPostInnerDeserializerRequestWithDefaults

`func NewCommonModelScopesPostInnerDeserializerRequestWithDefaults() *CommonModelScopesPostInnerDeserializerRequest`

NewCommonModelScopesPostInnerDeserializerRequestWithDefaults instantiates a new CommonModelScopesPostInnerDeserializerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *CommonModelScopesPostInnerDeserializerRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CommonModelScopesPostInnerDeserializerRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CommonModelScopesPostInnerDeserializerRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetEnabledActions

`func (o *CommonModelScopesPostInnerDeserializerRequest) GetEnabledActions() []EnabledActionsEnum`

GetEnabledActions returns the EnabledActions field if non-nil, zero value otherwise.

### GetEnabledActionsOk

`func (o *CommonModelScopesPostInnerDeserializerRequest) GetEnabledActionsOk() (*[]EnabledActionsEnum, bool)`

GetEnabledActionsOk returns a tuple with the EnabledActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledActions

`func (o *CommonModelScopesPostInnerDeserializerRequest) SetEnabledActions(v []EnabledActionsEnum)`

SetEnabledActions sets EnabledActions field to given value.


### GetDisabledFields

`func (o *CommonModelScopesPostInnerDeserializerRequest) GetDisabledFields() []string`

GetDisabledFields returns the DisabledFields field if non-nil, zero value otherwise.

### GetDisabledFieldsOk

`func (o *CommonModelScopesPostInnerDeserializerRequest) GetDisabledFieldsOk() (*[]string, bool)`

GetDisabledFieldsOk returns a tuple with the DisabledFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledFields

`func (o *CommonModelScopesPostInnerDeserializerRequest) SetDisabledFields(v []string)`

SetDisabledFields sets DisabledFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


