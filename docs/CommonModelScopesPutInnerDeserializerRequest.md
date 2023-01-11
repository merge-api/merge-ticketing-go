# CommonModelScopesPutInnerDeserializerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | **string** |  | 
**EnabledActions** | [**[]CommonModelScopesPutInnerDeserializerEnabledActionsEnum**](CommonModelScopesPutInnerDeserializerEnabledActionsEnum.md) |  | 
**DisabledFields** | **[]string** |  | 

## Methods

### NewCommonModelScopesPutInnerDeserializerRequest

`func NewCommonModelScopesPutInnerDeserializerRequest(modelId string, enabledActions []CommonModelScopesPutInnerDeserializerEnabledActionsEnum, disabledFields []string, ) *CommonModelScopesPutInnerDeserializerRequest`

NewCommonModelScopesPutInnerDeserializerRequest instantiates a new CommonModelScopesPutInnerDeserializerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonModelScopesPutInnerDeserializerRequestWithDefaults

`func NewCommonModelScopesPutInnerDeserializerRequestWithDefaults() *CommonModelScopesPutInnerDeserializerRequest`

NewCommonModelScopesPutInnerDeserializerRequestWithDefaults instantiates a new CommonModelScopesPutInnerDeserializerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *CommonModelScopesPutInnerDeserializerRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CommonModelScopesPutInnerDeserializerRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CommonModelScopesPutInnerDeserializerRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetEnabledActions

`func (o *CommonModelScopesPutInnerDeserializerRequest) GetEnabledActions() []CommonModelScopesPutInnerDeserializerEnabledActionsEnum`

GetEnabledActions returns the EnabledActions field if non-nil, zero value otherwise.

### GetEnabledActionsOk

`func (o *CommonModelScopesPutInnerDeserializerRequest) GetEnabledActionsOk() (*[]CommonModelScopesPutInnerDeserializerEnabledActionsEnum, bool)`

GetEnabledActionsOk returns a tuple with the EnabledActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledActions

`func (o *CommonModelScopesPutInnerDeserializerRequest) SetEnabledActions(v []CommonModelScopesPutInnerDeserializerEnabledActionsEnum)`

SetEnabledActions sets EnabledActions field to given value.


### GetDisabledFields

`func (o *CommonModelScopesPutInnerDeserializerRequest) GetDisabledFields() []string`

GetDisabledFields returns the DisabledFields field if non-nil, zero value otherwise.

### GetDisabledFieldsOk

`func (o *CommonModelScopesPutInnerDeserializerRequest) GetDisabledFieldsOk() (*[]string, bool)`

GetDisabledFieldsOk returns a tuple with the DisabledFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledFields

`func (o *CommonModelScopesPutInnerDeserializerRequest) SetDisabledFields(v []string)`

SetDisabledFields sets DisabledFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


