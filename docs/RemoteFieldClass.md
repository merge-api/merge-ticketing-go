# RemoteFieldClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**RemoteKeyName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsCustom** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**FieldType** | Pointer to [**FieldTypeEnum**](FieldTypeEnum.md) |  | [optional] 
**FieldFormat** | Pointer to [**FieldFormatEnum**](FieldFormatEnum.md) |  | [optional] 
**FieldChoices** | Pointer to **[]string** |  | [optional] 
**ItemSchema** | Pointer to [**ItemSchema**](ItemSchema.md) |  | [optional] 

## Methods

### NewRemoteFieldClass

`func NewRemoteFieldClass() *RemoteFieldClass`

NewRemoteFieldClass instantiates a new RemoteFieldClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteFieldClassWithDefaults

`func NewRemoteFieldClassWithDefaults() *RemoteFieldClass`

NewRemoteFieldClassWithDefaults instantiates a new RemoteFieldClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteFieldClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteFieldClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteFieldClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteFieldClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *RemoteFieldClass) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RemoteFieldClass) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RemoteFieldClass) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RemoteFieldClass) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRemoteKeyName

`func (o *RemoteFieldClass) GetRemoteKeyName() string`

GetRemoteKeyName returns the RemoteKeyName field if non-nil, zero value otherwise.

### GetRemoteKeyNameOk

`func (o *RemoteFieldClass) GetRemoteKeyNameOk() (*string, bool)`

GetRemoteKeyNameOk returns a tuple with the RemoteKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteKeyName

`func (o *RemoteFieldClass) SetRemoteKeyName(v string)`

SetRemoteKeyName sets RemoteKeyName field to given value.

### HasRemoteKeyName

`func (o *RemoteFieldClass) HasRemoteKeyName() bool`

HasRemoteKeyName returns a boolean if a field has been set.

### GetDescription

`func (o *RemoteFieldClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemoteFieldClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemoteFieldClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RemoteFieldClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsCustom

`func (o *RemoteFieldClass) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *RemoteFieldClass) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *RemoteFieldClass) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.

### HasIsCustom

`func (o *RemoteFieldClass) HasIsCustom() bool`

HasIsCustom returns a boolean if a field has been set.

### GetIsRequired

`func (o *RemoteFieldClass) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *RemoteFieldClass) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *RemoteFieldClass) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *RemoteFieldClass) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetFieldType

`func (o *RemoteFieldClass) GetFieldType() FieldTypeEnum`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *RemoteFieldClass) GetFieldTypeOk() (*FieldTypeEnum, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *RemoteFieldClass) SetFieldType(v FieldTypeEnum)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *RemoteFieldClass) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetFieldFormat

`func (o *RemoteFieldClass) GetFieldFormat() FieldFormatEnum`

GetFieldFormat returns the FieldFormat field if non-nil, zero value otherwise.

### GetFieldFormatOk

`func (o *RemoteFieldClass) GetFieldFormatOk() (*FieldFormatEnum, bool)`

GetFieldFormatOk returns a tuple with the FieldFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldFormat

`func (o *RemoteFieldClass) SetFieldFormat(v FieldFormatEnum)`

SetFieldFormat sets FieldFormat field to given value.

### HasFieldFormat

`func (o *RemoteFieldClass) HasFieldFormat() bool`

HasFieldFormat returns a boolean if a field has been set.

### GetFieldChoices

`func (o *RemoteFieldClass) GetFieldChoices() []string`

GetFieldChoices returns the FieldChoices field if non-nil, zero value otherwise.

### GetFieldChoicesOk

`func (o *RemoteFieldClass) GetFieldChoicesOk() (*[]string, bool)`

GetFieldChoicesOk returns a tuple with the FieldChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldChoices

`func (o *RemoteFieldClass) SetFieldChoices(v []string)`

SetFieldChoices sets FieldChoices field to given value.

### HasFieldChoices

`func (o *RemoteFieldClass) HasFieldChoices() bool`

HasFieldChoices returns a boolean if a field has been set.

### GetItemSchema

`func (o *RemoteFieldClass) GetItemSchema() ItemSchema`

GetItemSchema returns the ItemSchema field if non-nil, zero value otherwise.

### GetItemSchemaOk

`func (o *RemoteFieldClass) GetItemSchemaOk() (*ItemSchema, bool)`

GetItemSchemaOk returns a tuple with the ItemSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSchema

`func (o *RemoteFieldClass) SetItemSchema(v ItemSchema)`

SetItemSchema sets ItemSchema field to given value.

### HasItemSchema

`func (o *RemoteFieldClass) HasItemSchema() bool`

HasItemSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


