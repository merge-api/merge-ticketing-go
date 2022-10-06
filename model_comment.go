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
	"time"
)

// Comment # The Comment Object ### Description The `Comment` object is used to represent a comment on a ticket.  ### Usage Example TODO
type Comment struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	User NullableString `json:"user,omitempty"`
	Contact NullableString `json:"contact,omitempty"`
	// The comment's text body.
	Body NullableString `json:"body,omitempty"`
	// The comment's text body formatted as html.
	HtmlBody NullableString `json:"html_body,omitempty"`
	Ticket NullableString `json:"ticket,omitempty"`
	// Whether or not the comment is internal.
	IsPrivate NullableBool `json:"is_private,omitempty"`
	// When the third party's comment was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewComment instantiates a new Comment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComment() *Comment {
	this := Comment{}
	return &this
}

// NewCommentWithDefaults instantiates a new Comment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentWithDefaults() *Comment {
	this := Comment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Comment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Comment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Comment) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Comment) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Comment) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Comment) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Comment) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetUser() string {
	if o == nil || o.User.Get() == nil {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *Comment) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *Comment) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *Comment) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *Comment) UnsetUser() {
	o.User.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetContact() string {
	if o == nil || o.Contact.Get() == nil {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetContactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *Comment) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *Comment) SetContact(v string) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *Comment) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *Comment) UnsetContact() {
	o.Contact.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *Comment) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *Comment) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *Comment) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *Comment) UnsetBody() {
	o.Body.Unset()
}

// GetHtmlBody returns the HtmlBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetHtmlBody() string {
	if o == nil || o.HtmlBody.Get() == nil {
		var ret string
		return ret
	}
	return *o.HtmlBody.Get()
}

// GetHtmlBodyOk returns a tuple with the HtmlBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetHtmlBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HtmlBody.Get(), o.HtmlBody.IsSet()
}

// HasHtmlBody returns a boolean if a field has been set.
func (o *Comment) HasHtmlBody() bool {
	if o != nil && o.HtmlBody.IsSet() {
		return true
	}

	return false
}

// SetHtmlBody gets a reference to the given NullableString and assigns it to the HtmlBody field.
func (o *Comment) SetHtmlBody(v string) {
	o.HtmlBody.Set(&v)
}
// SetHtmlBodyNil sets the value for HtmlBody to be an explicit nil
func (o *Comment) SetHtmlBodyNil() {
	o.HtmlBody.Set(nil)
}

// UnsetHtmlBody ensures that no value is present for HtmlBody, not even an explicit nil
func (o *Comment) UnsetHtmlBody() {
	o.HtmlBody.Unset()
}

// GetTicket returns the Ticket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetTicket() string {
	if o == nil || o.Ticket.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ticket.Get()
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetTicketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ticket.Get(), o.Ticket.IsSet()
}

// HasTicket returns a boolean if a field has been set.
func (o *Comment) HasTicket() bool {
	if o != nil && o.Ticket.IsSet() {
		return true
	}

	return false
}

// SetTicket gets a reference to the given NullableString and assigns it to the Ticket field.
func (o *Comment) SetTicket(v string) {
	o.Ticket.Set(&v)
}
// SetTicketNil sets the value for Ticket to be an explicit nil
func (o *Comment) SetTicketNil() {
	o.Ticket.Set(nil)
}

// UnsetTicket ensures that no value is present for Ticket, not even an explicit nil
func (o *Comment) UnsetTicket() {
	o.Ticket.Unset()
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetIsPrivate() bool {
	if o == nil || o.IsPrivate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate.Get()
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetIsPrivateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPrivate.Get(), o.IsPrivate.IsSet()
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *Comment) HasIsPrivate() bool {
	if o != nil && o.IsPrivate.IsSet() {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given NullableBool and assigns it to the IsPrivate field.
func (o *Comment) SetIsPrivate(v bool) {
	o.IsPrivate.Set(&v)
}
// SetIsPrivateNil sets the value for IsPrivate to be an explicit nil
func (o *Comment) SetIsPrivateNil() {
	o.IsPrivate.Set(nil)
}

// UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
func (o *Comment) UnsetIsPrivate() {
	o.IsPrivate.Unset()
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *Comment) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *Comment) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *Comment) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *Comment) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Comment) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Comment) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *Comment) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *Comment) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *Comment) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

func (o Comment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.HtmlBody.IsSet() {
		toSerialize["html_body"] = o.HtmlBody.Get()
	}
	if o.Ticket.IsSet() {
		toSerialize["ticket"] = o.Ticket.Get()
	}
	if o.IsPrivate.IsSet() {
		toSerialize["is_private"] = o.IsPrivate.Get()
	}
	if o.RemoteCreatedAt.IsSet() {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	return json.Marshal(toSerialize)
}

func (v *Comment) UnmarshalJSON(src []byte) error {
    type CommentUnmarshalTarget Comment

	var intermediateResult CommentUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Comment(intermediateResult)
	return nil
}
type NullableComment struct {
	value *Comment
	isSet bool
}

func (v NullableComment) Get() *Comment {
	return v.value
}

func (v *NullableComment) Set(val *Comment) {
	v.value = val
	v.isSet = true
}

func (v NullableComment) IsSet() bool {
	return v.isSet
}

func (v *NullableComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComment(val *Comment) *NullableComment {
	return &NullableComment{value: val, isSet: true}
}

func (v NullableComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


