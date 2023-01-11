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

// CommonModelScopes struct for CommonModelScopes
type CommonModelScopes struct {
	Scope CommonModelScopesScope `json:"scope"`
	CommonModels []CommonModelScopesDisabledModels `json:"common_models"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCommonModelScopes instantiates a new CommonModelScopes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonModelScopes(scope CommonModelScopesScope, commonModels []CommonModelScopesDisabledModels) *CommonModelScopes {
	this := CommonModelScopes{}
	this.Scope = scope
	this.CommonModels = commonModels
	return &this
}

// NewCommonModelScopesWithDefaults instantiates a new CommonModelScopes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonModelScopesWithDefaults() *CommonModelScopes {
	this := CommonModelScopes{}
	return &this
}

// GetScope returns the Scope field value
func (o *CommonModelScopes) GetScope() CommonModelScopesScope {
	if o == nil {
		var ret CommonModelScopesScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopes) GetScopeOk() (*CommonModelScopesScope, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CommonModelScopes) SetScope(v CommonModelScopesScope) {
	o.Scope = v
}

// GetCommonModels returns the CommonModels field value
func (o *CommonModelScopes) GetCommonModels() []CommonModelScopesDisabledModels {
	if o == nil {
		var ret []CommonModelScopesDisabledModels
		return ret
	}

	return o.CommonModels
}

// GetCommonModelsOk returns a tuple with the CommonModels field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopes) GetCommonModelsOk() (*[]CommonModelScopesDisabledModels, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommonModels, true
}

// SetCommonModels sets field value
func (o *CommonModelScopes) SetCommonModels(v []CommonModelScopesDisabledModels) {
	o.CommonModels = v
}

func (o CommonModelScopes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["common_models"] = o.CommonModels
	}
	return json.Marshal(toSerialize)
}

func (v *CommonModelScopes) UnmarshalJSON(src []byte) error {
    type CommonModelScopesUnmarshalTarget CommonModelScopes

	var intermediateResult CommonModelScopesUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CommonModelScopes(intermediateResult)
	return nil
}
type NullableCommonModelScopes struct {
	value *CommonModelScopes
	isSet bool
}

func (v NullableCommonModelScopes) Get() *CommonModelScopes {
	return v.value
}

func (v *NullableCommonModelScopes) Set(val *CommonModelScopes) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonModelScopes) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonModelScopes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonModelScopes(val *CommonModelScopes) *NullableCommonModelScopes {
	return &NullableCommonModelScopes{value: val, isSet: true}
}

func (v NullableCommonModelScopes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonModelScopes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}

