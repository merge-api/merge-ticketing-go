# RemoteFieldClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**RemoteKeyName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**FieldType** | Pointer to [**FieldTypeEnum**](FieldTypeEnum.md) |  | [optional] [readonly] 
**FieldFormat** | Pointer to [**FieldFormatEnum**](FieldFormatEnum.md) |  | [optional] [readonly] 
**FieldChoices** | Pointer to **[]string** |  | [optional] [readonly] 
**ItemSchema** | Pointer to [**NullableRemoteFieldClassItemSchema**](RemoteFieldClassItemSchema.md) |  | [optional] 
**IsCustom** | Pointer to **NullableBool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteFields** | Pointer to [**[]RemoteField**](RemoteField.md) |  | [optional] [readonly] 

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

### SetDisplayNameNil

`func (o *RemoteFieldClass) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RemoteFieldClass) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
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

### SetRemoteKeyNameNil

`func (o *RemoteFieldClass) SetRemoteKeyNameNil(b bool)`

 SetRemoteKeyNameNil sets the value for RemoteKeyName to be an explicit nil

### UnsetRemoteKeyName
`func (o *RemoteFieldClass) UnsetRemoteKeyName()`

UnsetRemoteKeyName ensures that no value is present for RemoteKeyName, not even an explicit nil
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

### SetDescriptionNil

`func (o *RemoteFieldClass) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RemoteFieldClass) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetFieldChoicesNil

`func (o *RemoteFieldClass) SetFieldChoicesNil(b bool)`

 SetFieldChoicesNil sets the value for FieldChoices to be an explicit nil

### UnsetFieldChoices
`func (o *RemoteFieldClass) UnsetFieldChoices()`

UnsetFieldChoices ensures that no value is present for FieldChoices, not even an explicit nil
### GetItemSchema

`func (o *RemoteFieldClass) GetItemSchema() RemoteFieldClassItemSchema`

GetItemSchema returns the ItemSchema field if non-nil, zero value otherwise.

### GetItemSchemaOk

`func (o *RemoteFieldClass) GetItemSchemaOk() (*RemoteFieldClassItemSchema, bool)`

GetItemSchemaOk returns a tuple with the ItemSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSchema

`func (o *RemoteFieldClass) SetItemSchema(v RemoteFieldClassItemSchema)`

SetItemSchema sets ItemSchema field to given value.

### HasItemSchema

`func (o *RemoteFieldClass) HasItemSchema() bool`

HasItemSchema returns a boolean if a field has been set.

### SetItemSchemaNil

`func (o *RemoteFieldClass) SetItemSchemaNil(b bool)`

 SetItemSchemaNil sets the value for ItemSchema to be an explicit nil

### UnsetItemSchema
`func (o *RemoteFieldClass) UnsetItemSchema()`

UnsetItemSchema ensures that no value is present for ItemSchema, not even an explicit nil
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

### SetIsCustomNil

`func (o *RemoteFieldClass) SetIsCustomNil(b bool)`

 SetIsCustomNil sets the value for IsCustom to be an explicit nil

### UnsetIsCustom
`func (o *RemoteFieldClass) UnsetIsCustom()`

UnsetIsCustom ensures that no value is present for IsCustom, not even an explicit nil
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

### GetRemoteFields

`func (o *RemoteFieldClass) GetRemoteFields() []RemoteField`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *RemoteFieldClass) GetRemoteFieldsOk() (*[]RemoteField, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *RemoteFieldClass) SetRemoteFields(v []RemoteField)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *RemoteFieldClass) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


