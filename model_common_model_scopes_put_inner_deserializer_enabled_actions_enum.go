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
	"fmt"
)

// CommonModelScopesPutInnerDeserializerEnabledActionsEnum the model 'CommonModelScopesPutInnerDeserializerEnabledActionsEnum'
type CommonModelScopesPutInnerDeserializerEnabledActionsEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of CommonModelScopesPutInnerDeserializerEnabledActionsEnum
const (
    COMMONMODELSCOPESPUTINNERDESERIALIZERENABLEDACTIONSENUM_MERGE_NONSTANDARD_VALUE CommonModelScopesPutInnerDeserializerEnabledActionsEnum = "MERGE_NONSTANDARD_VALUE"
    
	COMMONMODELSCOPESPUTINNERDESERIALIZERENABLEDACTIONSENUM_READ CommonModelScopesPutInnerDeserializerEnabledActionsEnum = "READ"
	COMMONMODELSCOPESPUTINNERDESERIALIZERENABLEDACTIONSENUM_WRITE CommonModelScopesPutInnerDeserializerEnabledActionsEnum = "WRITE"
)

var allowedCommonModelScopesPutInnerDeserializerEnabledActionsEnumEnumValues = []CommonModelScopesPutInnerDeserializerEnabledActionsEnum{
	"READ",
	"WRITE",
}

func (v *CommonModelScopesPutInnerDeserializerEnabledActionsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommonModelScopesPutInnerDeserializerEnabledActionsEnum(value)
	for _, existing := range allowedCommonModelScopesPutInnerDeserializerEnabledActionsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = COMMONMODELSCOPESPUTINNERDESERIALIZERENABLEDACTIONSENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewCommonModelScopesPutInnerDeserializerEnabledActionsEnumFromValue returns a pointer to a valid CommonModelScopesPutInnerDeserializerEnabledActionsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommonModelScopesPutInnerDeserializerEnabledActionsEnumFromValue(v string) (*CommonModelScopesPutInnerDeserializerEnabledActionsEnum, error) {
	ev := CommonModelScopesPutInnerDeserializerEnabledActionsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := COMMONMODELSCOPESPUTINNERDESERIALIZERENABLEDACTIONSENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommonModelScopesPutInnerDeserializerEnabledActionsEnum) IsValid() bool {
	for _, existing := range allowedCommonModelScopesPutInnerDeserializerEnabledActionsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommonModelScopesPutInnerDeserializerEnabledActionsEnum value
func (v CommonModelScopesPutInnerDeserializerEnabledActionsEnum) Ptr() *CommonModelScopesPutInnerDeserializerEnabledActionsEnum {
	return &v
}

type NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum struct {
	value *CommonModelScopesPutInnerDeserializerEnabledActionsEnum
	isSet bool
}

func (v NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum) Get() *CommonModelScopesPutInnerDeserializerEnabledActionsEnum {
	return v.value
}

func (v *NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum) Set(val *CommonModelScopesPutInnerDeserializerEnabledActionsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum(val *CommonModelScopesPutInnerDeserializerEnabledActionsEnum) *NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum {
	return &NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum{value: val, isSet: true}
}

func (v NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonModelScopesPutInnerDeserializerEnabledActionsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
