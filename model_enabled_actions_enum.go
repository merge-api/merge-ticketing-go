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

// EnabledActionsEnum * `READ` - READ * `WRITE` - WRITE
type EnabledActionsEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of EnabledActionsEnum
const (
    ENABLEDACTIONSENUM_MERGE_NONSTANDARD_VALUE EnabledActionsEnum = "MERGE_NONSTANDARD_VALUE"
    
	ENABLEDACTIONSENUM_READ EnabledActionsEnum = "READ"
	ENABLEDACTIONSENUM_WRITE EnabledActionsEnum = "WRITE"
)

var allowedEnabledActionsEnumEnumValues = []EnabledActionsEnum{
	"READ",
	"WRITE",
}

func (v *EnabledActionsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnabledActionsEnum(value)
	for _, existing := range allowedEnabledActionsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = ENABLEDACTIONSENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewEnabledActionsEnumFromValue returns a pointer to a valid EnabledActionsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnabledActionsEnumFromValue(v string) (*EnabledActionsEnum, error) {
	ev := EnabledActionsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := ENABLEDACTIONSENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnabledActionsEnum) IsValid() bool {
	for _, existing := range allowedEnabledActionsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnabledActionsEnum value
func (v EnabledActionsEnum) Ptr() *EnabledActionsEnum {
	return &v
}

type NullableEnabledActionsEnum struct {
	value *EnabledActionsEnum
	isSet bool
}

func (v NullableEnabledActionsEnum) Get() *EnabledActionsEnum {
	return v.value
}

func (v *NullableEnabledActionsEnum) Set(val *EnabledActionsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEnabledActionsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEnabledActionsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnabledActionsEnum(val *EnabledActionsEnum) *NullableEnabledActionsEnum {
	return &NullableEnabledActionsEnum{value: val, isSet: true}
}

func (v NullableEnabledActionsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnabledActionsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

