# ItemSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemType** | Pointer to **bool** |  | [optional] 
**ItemFormat** | Pointer to **bool** |  | [optional] 
**ItemChoices** | Pointer to **[]string** |  | [optional] 

## Methods

### NewItemSchema

`func NewItemSchema() *ItemSchema`

NewItemSchema instantiates a new ItemSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSchemaWithDefaults

`func NewItemSchemaWithDefaults() *ItemSchema`

NewItemSchemaWithDefaults instantiates a new ItemSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemType

`func (o *ItemSchema) GetItemType() bool`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ItemSchema) GetItemTypeOk() (*bool, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ItemSchema) SetItemType(v bool)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ItemSchema) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetItemFormat

`func (o *ItemSchema) GetItemFormat() bool`

GetItemFormat returns the ItemFormat field if non-nil, zero value otherwise.

### GetItemFormatOk

`func (o *ItemSchema) GetItemFormatOk() (*bool, bool)`

GetItemFormatOk returns a tuple with the ItemFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemFormat

`func (o *ItemSchema) SetItemFormat(v bool)`

SetItemFormat sets ItemFormat field to given value.

### HasItemFormat

`func (o *ItemSchema) HasItemFormat() bool`

HasItemFormat returns a boolean if a field has been set.

### GetItemChoices

`func (o *ItemSchema) GetItemChoices() []string`

GetItemChoices returns the ItemChoices field if non-nil, zero value otherwise.

### GetItemChoicesOk

`func (o *ItemSchema) GetItemChoicesOk() (*[]string, bool)`

GetItemChoicesOk returns a tuple with the ItemChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemChoices

`func (o *ItemSchema) SetItemChoices(v []string)`

SetItemChoices sets ItemChoices field to given value.

### HasItemChoices

`func (o *ItemSchema) HasItemChoices() bool`

HasItemChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


