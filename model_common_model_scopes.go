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
	OrganizationLevelScopes *CommonModelScopeData `json:"organization_level_scopes,omitempty"`
	ScopeOverrides []CommonModelScopeData `json:"scope_overrides"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCommonModelScopes instantiates a new CommonModelScopes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonModelScopes(scopeOverrides []CommonModelScopeData) *CommonModelScopes {
	this := CommonModelScopes{}
	this.ScopeOverrides = scopeOverrides
	return &this
}

// NewCommonModelScopesWithDefaults instantiates a new CommonModelScopes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonModelScopesWithDefaults() *CommonModelScopes {
	this := CommonModelScopes{}
	return &this
}

// GetOrganizationLevelScopes returns the OrganizationLevelScopes field value if set, zero value otherwise.
func (o *CommonModelScopes) GetOrganizationLevelScopes() CommonModelScopeData {
	if o == nil || o.OrganizationLevelScopes == nil {
		var ret CommonModelScopeData
		return ret
	}
	return *o.OrganizationLevelScopes
}

// GetOrganizationLevelScopesOk returns a tuple with the OrganizationLevelScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonModelScopes) GetOrganizationLevelScopesOk() (*CommonModelScopeData, bool) {
	if o == nil || o.OrganizationLevelScopes == nil {
		return nil, false
	}
	return o.OrganizationLevelScopes, true
}

// HasOrganizationLevelScopes returns a boolean if a field has been set.
func (o *CommonModelScopes) HasOrganizationLevelScopes() bool {
	if o != nil && o.OrganizationLevelScopes != nil {
		return true
	}

	return false
}

// SetOrganizationLevelScopes gets a reference to the given CommonModelScopeData and assigns it to the OrganizationLevelScopes field.
func (o *CommonModelScopes) SetOrganizationLevelScopes(v CommonModelScopeData) {
	o.OrganizationLevelScopes = &v
}

// GetScopeOverrides returns the ScopeOverrides field value
func (o *CommonModelScopes) GetScopeOverrides() []CommonModelScopeData {
	if o == nil {
		var ret []CommonModelScopeData
		return ret
	}

	return o.ScopeOverrides
}

// GetScopeOverridesOk returns a tuple with the ScopeOverrides field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopes) GetScopeOverridesOk() (*[]CommonModelScopeData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScopeOverrides, true
}

// SetScopeOverrides sets field value
func (o *CommonModelScopes) SetScopeOverrides(v []CommonModelScopeData) {
	o.ScopeOverrides = v
}

func (o CommonModelScopes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrganizationLevelScopes != nil {
		toSerialize["organization_level_scopes"] = o.OrganizationLevelScopes
	}
	if true {
		toSerialize["scope_overrides"] = o.ScopeOverrides
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


