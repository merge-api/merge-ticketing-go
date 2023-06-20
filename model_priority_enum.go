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

// PriorityEnum * `URGENT` - URGENT * `HIGH` - HIGH * `NORMAL` - NORMAL * `LOW` - LOW
type PriorityEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of PriorityEnum
const (
    PRIORITYENUM_MERGE_NONSTANDARD_VALUE PriorityEnum = "MERGE_NONSTANDARD_VALUE"
    
	PRIORITYENUM_URGENT PriorityEnum = "URGENT"
	PRIORITYENUM_HIGH PriorityEnum = "HIGH"
	PRIORITYENUM_NORMAL PriorityEnum = "NORMAL"
	PRIORITYENUM_LOW PriorityEnum = "LOW"
)

var allowedPriorityEnumEnumValues = []PriorityEnum{
	"URGENT",
	"HIGH",
	"NORMAL",
	"LOW",
}

func (v *PriorityEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriorityEnum(value)
	for _, existing := range allowedPriorityEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = PRIORITYENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewPriorityEnumFromValue returns a pointer to a valid PriorityEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriorityEnumFromValue(v string) (*PriorityEnum, error) {
	ev := PriorityEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := PRIORITYENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriorityEnum) IsValid() bool {
	for _, existing := range allowedPriorityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PriorityEnum value
func (v PriorityEnum) Ptr() *PriorityEnum {
	return &v
}

type NullablePriorityEnum struct {
	value *PriorityEnum
	isSet bool
}

func (v NullablePriorityEnum) Get() *PriorityEnum {
	return v.value
}

func (v *NullablePriorityEnum) Set(val *PriorityEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePriorityEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePriorityEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriorityEnum(val *PriorityEnum) *NullablePriorityEnum {
	return &NullablePriorityEnum{value: val, isSet: true}
}

func (v NullablePriorityEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriorityEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

