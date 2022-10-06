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

// CategoriesEnum the model 'CategoriesEnum'
type CategoriesEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of CategoriesEnum
const (
    CATEGORIESENUM_MERGE_NONSTANDARD_VALUE CategoriesEnum = "MERGE_NONSTANDARD_VALUE"
    
	CATEGORIESENUM_HRIS CategoriesEnum = "hris"
	CATEGORIESENUM_ATS CategoriesEnum = "ats"
	CATEGORIESENUM_ACCOUNTING CategoriesEnum = "accounting"
	CATEGORIESENUM_TICKETING CategoriesEnum = "ticketing"
	CATEGORIESENUM_CRM CategoriesEnum = "crm"
)

var allowedCategoriesEnumEnumValues = []CategoriesEnum{
	"hris",
	"ats",
	"accounting",
	"ticketing",
	"crm",
}

func (v *CategoriesEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CategoriesEnum(value)
	for _, existing := range allowedCategoriesEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = CATEGORIESENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewCategoriesEnumFromValue returns a pointer to a valid CategoriesEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCategoriesEnumFromValue(v string) (*CategoriesEnum, error) {
	ev := CategoriesEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := CATEGORIESENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CategoriesEnum) IsValid() bool {
	for _, existing := range allowedCategoriesEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CategoriesEnum value
func (v CategoriesEnum) Ptr() *CategoriesEnum {
	return &v
}

type NullableCategoriesEnum struct {
	value *CategoriesEnum
	isSet bool
}

func (v NullableCategoriesEnum) Get() *CategoriesEnum {
	return v.value
}

func (v *NullableCategoriesEnum) Set(val *CategoriesEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoriesEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoriesEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoriesEnum(val *CategoriesEnum) *NullableCategoriesEnum {
	return &NullableCategoriesEnum{value: val, isSet: true}
}

func (v NullableCategoriesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoriesEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
