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

// CommonModelScopesPutInnerDeserializerRequest struct for CommonModelScopesPutInnerDeserializerRequest
type CommonModelScopesPutInnerDeserializerRequest struct {
	ModelId string `json:"model_id"`
	EnabledActions []CommonModelScopesPutInnerDeserializerEnabledActionsEnum `json:"enabled_actions"`
	DisabledFields []string `json:"disabled_fields"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCommonModelScopesPutInnerDeserializerRequest instantiates a new CommonModelScopesPutInnerDeserializerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonModelScopesPutInnerDeserializerRequest(modelId string, enabledActions []CommonModelScopesPutInnerDeserializerEnabledActionsEnum, disabledFields []string) *CommonModelScopesPutInnerDeserializerRequest {
	this := CommonModelScopesPutInnerDeserializerRequest{}
	this.ModelId = modelId
	this.EnabledActions = enabledActions
	this.DisabledFields = disabledFields
	return &this
}

// NewCommonModelScopesPutInnerDeserializerRequestWithDefaults instantiates a new CommonModelScopesPutInnerDeserializerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonModelScopesPutInnerDeserializerRequestWithDefaults() *CommonModelScopesPutInnerDeserializerRequest {
	this := CommonModelScopesPutInnerDeserializerRequest{}
	return &this
}

// GetModelId returns the ModelId field value
func (o *CommonModelScopesPutInnerDeserializerRequest) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopesPutInnerDeserializerRequest) GetModelIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *CommonModelScopesPutInnerDeserializerRequest) SetModelId(v string) {
	o.ModelId = v
}

// GetEnabledActions returns the EnabledActions field value
func (o *CommonModelScopesPutInnerDeserializerRequest) GetEnabledActions() []CommonModelScopesPutInnerDeserializerEnabledActionsEnum {
	if o == nil {
		var ret []CommonModelScopesPutInnerDeserializerEnabledActionsEnum
		return ret
	}

	return o.EnabledActions
}

// GetEnabledActionsOk returns a tuple with the EnabledActions field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopesPutInnerDeserializerRequest) GetEnabledActionsOk() (*[]CommonModelScopesPutInnerDeserializerEnabledActionsEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnabledActions, true
}

// SetEnabledActions sets field value
func (o *CommonModelScopesPutInnerDeserializerRequest) SetEnabledActions(v []CommonModelScopesPutInnerDeserializerEnabledActionsEnum) {
	o.EnabledActions = v
}

// GetDisabledFields returns the DisabledFields field value
func (o *CommonModelScopesPutInnerDeserializerRequest) GetDisabledFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DisabledFields
}

// GetDisabledFieldsOk returns a tuple with the DisabledFields field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopesPutInnerDeserializerRequest) GetDisabledFieldsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisabledFields, true
}

// SetDisabledFields sets field value
func (o *CommonModelScopesPutInnerDeserializerRequest) SetDisabledFields(v []string) {
	o.DisabledFields = v
}

func (o CommonModelScopesPutInnerDeserializerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model_id"] = o.ModelId
	}
	if true {
		toSerialize["enabled_actions"] = o.EnabledActions
	}
	if true {
		toSerialize["disabled_fields"] = o.DisabledFields
	}
	return json.Marshal(toSerialize)
}

func (v *CommonModelScopesPutInnerDeserializerRequest) UnmarshalJSON(src []byte) error {
    type CommonModelScopesPutInnerDeserializerRequestUnmarshalTarget CommonModelScopesPutInnerDeserializerRequest

	var intermediateResult CommonModelScopesPutInnerDeserializerRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CommonModelScopesPutInnerDeserializerRequest(intermediateResult)
	return nil
}
type NullableCommonModelScopesPutInnerDeserializerRequest struct {
	value *CommonModelScopesPutInnerDeserializerRequest
	isSet bool
}

func (v NullableCommonModelScopesPutInnerDeserializerRequest) Get() *CommonModelScopesPutInnerDeserializerRequest {
	return v.value
}

func (v *NullableCommonModelScopesPutInnerDeserializerRequest) Set(val *CommonModelScopesPutInnerDeserializerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonModelScopesPutInnerDeserializerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonModelScopesPutInnerDeserializerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonModelScopesPutInnerDeserializerRequest(val *CommonModelScopesPutInnerDeserializerRequest) *NullableCommonModelScopesPutInnerDeserializerRequest {
	return &NullableCommonModelScopesPutInnerDeserializerRequest{value: val, isSet: true}
}

func (v NullableCommonModelScopesPutInnerDeserializerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonModelScopesPutInnerDeserializerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}

