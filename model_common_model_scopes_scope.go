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

// CommonModelScopesScope struct for CommonModelScopesScope
type CommonModelScopesScope struct {
	Type TypeEnum `json:"type"`
	Value string `json:"value"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCommonModelScopesScope instantiates a new CommonModelScopesScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonModelScopesScope(type_ TypeEnum, value string) *CommonModelScopesScope {
	this := CommonModelScopesScope{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewCommonModelScopesScopeWithDefaults instantiates a new CommonModelScopesScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonModelScopesScopeWithDefaults() *CommonModelScopesScope {
	this := CommonModelScopesScope{}
	return &this
}

// GetType returns the Type field value
func (o *CommonModelScopesScope) GetType() TypeEnum {
	if o == nil {
		var ret TypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopesScope) GetTypeOk() (*TypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CommonModelScopesScope) SetType(v TypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *CommonModelScopesScope) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopesScope) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CommonModelScopesScope) SetValue(v string) {
	o.Value = v
}

func (o CommonModelScopesScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

func (v *CommonModelScopesScope) UnmarshalJSON(src []byte) error {
    type CommonModelScopesScopeUnmarshalTarget CommonModelScopesScope

	var intermediateResult CommonModelScopesScopeUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CommonModelScopesScope(intermediateResult)
	return nil
}
type NullableCommonModelScopesScope struct {
	value *CommonModelScopesScope
	isSet bool
}

func (v NullableCommonModelScopesScope) Get() *CommonModelScopesScope {
	return v.value
}

func (v *NullableCommonModelScopesScope) Set(val *CommonModelScopesScope) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonModelScopesScope) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonModelScopesScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonModelScopesScope(val *CommonModelScopesScope) *NullableCommonModelScopesScope {
	return &NullableCommonModelScopesScope{value: val, isSet: true}
}

func (v NullableCommonModelScopesScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonModelScopesScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


