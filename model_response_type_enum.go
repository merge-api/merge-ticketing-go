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

// ResponseTypeEnum the model 'ResponseTypeEnum'
type ResponseTypeEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of ResponseTypeEnum
const (
    RESPONSETYPEENUM_MERGE_NONSTANDARD_VALUE ResponseTypeEnum = "MERGE_NONSTANDARD_VALUE"
    
	RESPONSETYPEENUM_JSON ResponseTypeEnum = "JSON"
	RESPONSETYPEENUM_BASE64_GZIP ResponseTypeEnum = "BASE64_GZIP"
)

var allowedResponseTypeEnumEnumValues = []ResponseTypeEnum{
	"JSON",
	"BASE64_GZIP",
}

func (v *ResponseTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResponseTypeEnum(value)
	for _, existing := range allowedResponseTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = RESPONSETYPEENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewResponseTypeEnumFromValue returns a pointer to a valid ResponseTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResponseTypeEnumFromValue(v string) (*ResponseTypeEnum, error) {
	ev := ResponseTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := RESPONSETYPEENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResponseTypeEnum) IsValid() bool {
	for _, existing := range allowedResponseTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResponseTypeEnum value
func (v ResponseTypeEnum) Ptr() *ResponseTypeEnum {
	return &v
}

type NullableResponseTypeEnum struct {
	value *ResponseTypeEnum
	isSet bool
}

func (v NullableResponseTypeEnum) Get() *ResponseTypeEnum {
	return v.value
}

func (v *NullableResponseTypeEnum) Set(val *ResponseTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTypeEnum(val *ResponseTypeEnum) *NullableResponseTypeEnum {
	return &NullableResponseTypeEnum{value: val, isSet: true}
}

func (v NullableResponseTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

