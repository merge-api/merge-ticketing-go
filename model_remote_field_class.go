/*
 * Merge Ticketing API
 *
 * The unified API for building rich integrations with multiple Ticketing platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_ticketing_client

import (
	"encoding/json"
)

// RemoteFieldClass struct for RemoteFieldClass
type RemoteFieldClass struct {
	Id *string `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	RemoteKeyName *string `json:"remote_key_name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsCustom *bool `json:"is_custom,omitempty"`
	IsRequired *bool `json:"is_required,omitempty"`
	FieldType *FieldTypeEnum `json:"field_type,omitempty"`
	FieldFormat *FieldFormatEnum `json:"field_format,omitempty"`
	FieldChoices *[]string `json:"field_choices,omitempty"`
	ItemSchema *ItemSchema `json:"item_schema,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewRemoteFieldClass instantiates a new RemoteFieldClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteFieldClass() *RemoteFieldClass {
	this := RemoteFieldClass{}
	return &this
}

// NewRemoteFieldClassWithDefaults instantiates a new RemoteFieldClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteFieldClassWithDefaults() *RemoteFieldClass {
	this := RemoteFieldClass{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteFieldClass) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RemoteFieldClass) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetRemoteKeyName returns the RemoteKeyName field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetRemoteKeyName() string {
	if o == nil || o.RemoteKeyName == nil {
		var ret string
		return ret
	}
	return *o.RemoteKeyName
}

// GetRemoteKeyNameOk returns a tuple with the RemoteKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetRemoteKeyNameOk() (*string, bool) {
	if o == nil || o.RemoteKeyName == nil {
		return nil, false
	}
	return o.RemoteKeyName, true
}

// HasRemoteKeyName returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasRemoteKeyName() bool {
	if o != nil && o.RemoteKeyName != nil {
		return true
	}

	return false
}

// SetRemoteKeyName gets a reference to the given string and assigns it to the RemoteKeyName field.
func (o *RemoteFieldClass) SetRemoteKeyName(v string) {
	o.RemoteKeyName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RemoteFieldClass) SetDescription(v string) {
	o.Description = &v
}

// GetIsCustom returns the IsCustom field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetIsCustom() bool {
	if o == nil || o.IsCustom == nil {
		var ret bool
		return ret
	}
	return *o.IsCustom
}

// GetIsCustomOk returns a tuple with the IsCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetIsCustomOk() (*bool, bool) {
	if o == nil || o.IsCustom == nil {
		return nil, false
	}
	return o.IsCustom, true
}

// HasIsCustom returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasIsCustom() bool {
	if o != nil && o.IsCustom != nil {
		return true
	}

	return false
}

// SetIsCustom gets a reference to the given bool and assigns it to the IsCustom field.
func (o *RemoteFieldClass) SetIsCustom(v bool) {
	o.IsCustom = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetIsRequired() bool {
	if o == nil || o.IsRequired == nil {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetIsRequiredOk() (*bool, bool) {
	if o == nil || o.IsRequired == nil {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasIsRequired() bool {
	if o != nil && o.IsRequired != nil {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *RemoteFieldClass) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetFieldType returns the FieldType field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetFieldType() FieldTypeEnum {
	if o == nil || o.FieldType == nil {
		var ret FieldTypeEnum
		return ret
	}
	return *o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetFieldTypeOk() (*FieldTypeEnum, bool) {
	if o == nil || o.FieldType == nil {
		return nil, false
	}
	return o.FieldType, true
}

// HasFieldType returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasFieldType() bool {
	if o != nil && o.FieldType != nil {
		return true
	}

	return false
}

// SetFieldType gets a reference to the given FieldTypeEnum and assigns it to the FieldType field.
func (o *RemoteFieldClass) SetFieldType(v FieldTypeEnum) {
	o.FieldType = &v
}

// GetFieldFormat returns the FieldFormat field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetFieldFormat() FieldFormatEnum {
	if o == nil || o.FieldFormat == nil {
		var ret FieldFormatEnum
		return ret
	}
	return *o.FieldFormat
}

// GetFieldFormatOk returns a tuple with the FieldFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetFieldFormatOk() (*FieldFormatEnum, bool) {
	if o == nil || o.FieldFormat == nil {
		return nil, false
	}
	return o.FieldFormat, true
}

// HasFieldFormat returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasFieldFormat() bool {
	if o != nil && o.FieldFormat != nil {
		return true
	}

	return false
}

// SetFieldFormat gets a reference to the given FieldFormatEnum and assigns it to the FieldFormat field.
func (o *RemoteFieldClass) SetFieldFormat(v FieldFormatEnum) {
	o.FieldFormat = &v
}

// GetFieldChoices returns the FieldChoices field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetFieldChoices() []string {
	if o == nil || o.FieldChoices == nil {
		var ret []string
		return ret
	}
	return *o.FieldChoices
}

// GetFieldChoicesOk returns a tuple with the FieldChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetFieldChoicesOk() (*[]string, bool) {
	if o == nil || o.FieldChoices == nil {
		return nil, false
	}
	return o.FieldChoices, true
}

// HasFieldChoices returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasFieldChoices() bool {
	if o != nil && o.FieldChoices != nil {
		return true
	}

	return false
}

// SetFieldChoices gets a reference to the given []string and assigns it to the FieldChoices field.
func (o *RemoteFieldClass) SetFieldChoices(v []string) {
	o.FieldChoices = &v
}

// GetItemSchema returns the ItemSchema field value if set, zero value otherwise.
func (o *RemoteFieldClass) GetItemSchema() ItemSchema {
	if o == nil || o.ItemSchema == nil {
		var ret ItemSchema
		return ret
	}
	return *o.ItemSchema
}

// GetItemSchemaOk returns a tuple with the ItemSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteFieldClass) GetItemSchemaOk() (*ItemSchema, bool) {
	if o == nil || o.ItemSchema == nil {
		return nil, false
	}
	return o.ItemSchema, true
}

// HasItemSchema returns a boolean if a field has been set.
func (o *RemoteFieldClass) HasItemSchema() bool {
	if o != nil && o.ItemSchema != nil {
		return true
	}

	return false
}

// SetItemSchema gets a reference to the given ItemSchema and assigns it to the ItemSchema field.
func (o *RemoteFieldClass) SetItemSchema(v ItemSchema) {
	o.ItemSchema = &v
}

func (o RemoteFieldClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.RemoteKeyName != nil {
		toSerialize["remote_key_name"] = o.RemoteKeyName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IsCustom != nil {
		toSerialize["is_custom"] = o.IsCustom
	}
	if o.IsRequired != nil {
		toSerialize["is_required"] = o.IsRequired
	}
	if o.FieldType != nil {
		toSerialize["field_type"] = o.FieldType
	}
	if o.FieldFormat != nil {
		toSerialize["field_format"] = o.FieldFormat
	}
	if o.FieldChoices != nil {
		toSerialize["field_choices"] = o.FieldChoices
	}
	if o.ItemSchema != nil {
		toSerialize["item_schema"] = o.ItemSchema
	}
	return json.Marshal(toSerialize)
}

func (v *RemoteFieldClass) UnmarshalJSON(src []byte) error {
    type RemoteFieldClassUnmarshalTarget RemoteFieldClass

	var intermediateResult RemoteFieldClassUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = RemoteFieldClass(intermediateResult)
	return nil
}
type NullableRemoteFieldClass struct {
	value *RemoteFieldClass
	isSet bool
}

func (v NullableRemoteFieldClass) Get() *RemoteFieldClass {
	return v.value
}

func (v *NullableRemoteFieldClass) Set(val *RemoteFieldClass) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteFieldClass) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteFieldClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteFieldClass(val *RemoteFieldClass) *NullableRemoteFieldClass {
	return &NullableRemoteFieldClass{value: val, isSet: true}
}

func (v NullableRemoteFieldClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteFieldClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


