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

// AccountDetailsAndActionsIntegration struct for AccountDetailsAndActionsIntegration
type AccountDetailsAndActionsIntegration struct {
	Name string `json:"name"`
	Categories []CategoriesEnum `json:"categories"`
	Image *string `json:"image,omitempty"`
	SquareImage *string `json:"square_image,omitempty"`
	Color string `json:"color"`
	Slug string `json:"slug"`
	PassthroughAvailable bool `json:"passthrough_available"`
	AvailableModelOperations *[]ModelOperation `json:"available_model_operations,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewAccountDetailsAndActionsIntegration instantiates a new AccountDetailsAndActionsIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDetailsAndActionsIntegration(name string, categories []CategoriesEnum, color string, slug string, passthroughAvailable bool) *AccountDetailsAndActionsIntegration {
	this := AccountDetailsAndActionsIntegration{}
	this.Name = name
	this.Categories = categories
	this.Color = color
	this.Slug = slug
	this.PassthroughAvailable = passthroughAvailable
	return &this
}

// NewAccountDetailsAndActionsIntegrationWithDefaults instantiates a new AccountDetailsAndActionsIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDetailsAndActionsIntegrationWithDefaults() *AccountDetailsAndActionsIntegration {
	this := AccountDetailsAndActionsIntegration{}
	return &this
}

// GetName returns the Name field value
func (o *AccountDetailsAndActionsIntegration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountDetailsAndActionsIntegration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountDetailsAndActionsIntegration) SetName(v string) {
	o.Name = v
}

// GetCategories returns the Categories field value
func (o *AccountDetailsAndActionsIntegration) GetCategories() []CategoriesEnum {
	if o == nil {
		var ret []CategoriesEnum
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *AccountDetailsAndActionsIntegration) GetCategoriesOk() (*[]CategoriesEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *AccountDetailsAndActionsIntegration) SetCategories(v []CategoriesEnum) {
	o.Categories = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AccountDetailsAndActionsIntegration) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetailsAndActionsIntegration) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AccountDetailsAndActionsIntegration) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *AccountDetailsAndActionsIntegration) SetImage(v string) {
	o.Image = &v
}

// GetSquareImage returns the SquareImage field value if set, zero value otherwise.
func (o *AccountDetailsAndActionsIntegration) GetSquareImage() string {
	if o == nil || o.SquareImage == nil {
		var ret string
		return ret
	}
	return *o.SquareImage
}

// GetSquareImageOk returns a tuple with the SquareImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetailsAndActionsIntegration) GetSquareImageOk() (*string, bool) {
	if o == nil || o.SquareImage == nil {
		return nil, false
	}
	return o.SquareImage, true
}

// HasSquareImage returns a boolean if a field has been set.
func (o *AccountDetailsAndActionsIntegration) HasSquareImage() bool {
	if o != nil && o.SquareImage != nil {
		return true
	}

	return false
}

// SetSquareImage gets a reference to the given string and assigns it to the SquareImage field.
func (o *AccountDetailsAndActionsIntegration) SetSquareImage(v string) {
	o.SquareImage = &v
}

// GetColor returns the Color field value
func (o *AccountDetailsAndActionsIntegration) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *AccountDetailsAndActionsIntegration) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *AccountDetailsAndActionsIntegration) SetColor(v string) {
	o.Color = v
}

// GetSlug returns the Slug field value
func (o *AccountDetailsAndActionsIntegration) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *AccountDetailsAndActionsIntegration) GetSlugOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *AccountDetailsAndActionsIntegration) SetSlug(v string) {
	o.Slug = v
}

// GetPassthroughAvailable returns the PassthroughAvailable field value
func (o *AccountDetailsAndActionsIntegration) GetPassthroughAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PassthroughAvailable
}

// GetPassthroughAvailableOk returns a tuple with the PassthroughAvailable field value
// and a boolean to check if the value has been set.
func (o *AccountDetailsAndActionsIntegration) GetPassthroughAvailableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PassthroughAvailable, true
}

// SetPassthroughAvailable sets field value
func (o *AccountDetailsAndActionsIntegration) SetPassthroughAvailable(v bool) {
	o.PassthroughAvailable = v
}

// GetAvailableModelOperations returns the AvailableModelOperations field value if set, zero value otherwise.
func (o *AccountDetailsAndActionsIntegration) GetAvailableModelOperations() []ModelOperation {
	if o == nil || o.AvailableModelOperations == nil {
		var ret []ModelOperation
		return ret
	}
	return *o.AvailableModelOperations
}

// GetAvailableModelOperationsOk returns a tuple with the AvailableModelOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetailsAndActionsIntegration) GetAvailableModelOperationsOk() (*[]ModelOperation, bool) {
	if o == nil || o.AvailableModelOperations == nil {
		return nil, false
	}
	return o.AvailableModelOperations, true
}

// HasAvailableModelOperations returns a boolean if a field has been set.
func (o *AccountDetailsAndActionsIntegration) HasAvailableModelOperations() bool {
	if o != nil && o.AvailableModelOperations != nil {
		return true
	}

	return false
}

// SetAvailableModelOperations gets a reference to the given []ModelOperation and assigns it to the AvailableModelOperations field.
func (o *AccountDetailsAndActionsIntegration) SetAvailableModelOperations(v []ModelOperation) {
	o.AvailableModelOperations = &v
}

func (o AccountDetailsAndActionsIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["categories"] = o.Categories
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.SquareImage != nil {
		toSerialize["square_image"] = o.SquareImage
	}
	if true {
		toSerialize["color"] = o.Color
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["passthrough_available"] = o.PassthroughAvailable
	}
	if o.AvailableModelOperations != nil {
		toSerialize["available_model_operations"] = o.AvailableModelOperations
	}
	return json.Marshal(toSerialize)
}

func (v *AccountDetailsAndActionsIntegration) UnmarshalJSON(src []byte) error {
    type AccountDetailsAndActionsIntegrationUnmarshalTarget AccountDetailsAndActionsIntegration

	var intermediateResult AccountDetailsAndActionsIntegrationUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = AccountDetailsAndActionsIntegration(intermediateResult)
	return nil
}
type NullableAccountDetailsAndActionsIntegration struct {
	value *AccountDetailsAndActionsIntegration
	isSet bool
}

func (v NullableAccountDetailsAndActionsIntegration) Get() *AccountDetailsAndActionsIntegration {
	return v.value
}

func (v *NullableAccountDetailsAndActionsIntegration) Set(val *AccountDetailsAndActionsIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDetailsAndActionsIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDetailsAndActionsIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDetailsAndActionsIntegration(val *AccountDetailsAndActionsIntegration) *NullableAccountDetailsAndActionsIntegration {
	return &NullableAccountDetailsAndActionsIntegration{value: val, isSet: true}
}

func (v NullableAccountDetailsAndActionsIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDetailsAndActionsIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


