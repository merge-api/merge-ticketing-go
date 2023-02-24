# PaginatedRemoteFieldClassList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]RemoteFieldClass**](RemoteFieldClass.md) |  | [optional] 

## Methods

### NewPaginatedRemoteFieldClassList

`func NewPaginatedRemoteFieldClassList() *PaginatedRemoteFieldClassList`

NewPaginatedRemoteFieldClassList instantiates a new PaginatedRemoteFieldClassList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRemoteFieldClassListWithDefaults

`func NewPaginatedRemoteFieldClassListWithDefaults() *PaginatedRemoteFieldClassList`

NewPaginatedRemoteFieldClassListWithDefaults instantiates a new PaginatedRemoteFieldClassList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedRemoteFieldClassList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedRemoteFieldClassList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedRemoteFieldClassList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedRemoteFieldClassList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedRemoteFieldClassList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedRemoteFieldClassList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedRemoteFieldClassList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedRemoteFieldClassList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedRemoteFieldClassList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedRemoteFieldClassList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedRemoteFieldClassList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedRemoteFieldClassList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedRemoteFieldClassList) GetResults() []RemoteFieldClass`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRemoteFieldClassList) GetResultsOk() (*[]RemoteFieldClass, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRemoteFieldClassList) SetResults(v []RemoteFieldClass)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedRemoteFieldClassList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


