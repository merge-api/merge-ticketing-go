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

// CommentRequest # The Comment Object ### Description The `Comment` object is used to represent a comment on a ticket.  ### Usage Example TODO
type CommentRequest struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The author of the Comment, if the author is a User.
	User NullableString `json:"user,omitempty"`
	// The author of the Comment, if the author is a Contact.
	Contact NullableString `json:"contact,omitempty"`
	// The comment's text body.
	Body NullableString `json:"body,omitempty"`
	// The comment's text body formatted as html.
	HtmlBody NullableString `json:"html_body,omitempty"`
	// The ticket associated with the comment. 
	Ticket NullableString `json:"ticket,omitempty"`
	// Whether or not the comment is internal.
	IsPrivate NullableBool `json:"is_private,omitempty"`
	// When the third party's comment was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCommentRequest instantiates a new CommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentRequest() *CommentRequest {
	this := CommentRequest{}
	return &this
}

// NewCommentRequestWithDefaults instantiates a new CommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentRequestWithDefaults() *CommentRequest {
	this := CommentRequest{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *CommentRequest) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *CommentRequest) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *CommentRequest) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *CommentRequest) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetUser() string {
	if o == nil || o.User.Get() == nil {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *CommentRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *CommentRequest) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *CommentRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *CommentRequest) UnsetUser() {
	o.User.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetContact() string {
	if o == nil || o.Contact.Get() == nil {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetContactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *CommentRequest) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *CommentRequest) SetContact(v string) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *CommentRequest) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *CommentRequest) UnsetContact() {
	o.Contact.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *CommentRequest) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *CommentRequest) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *CommentRequest) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *CommentRequest) UnsetBody() {
	o.Body.Unset()
}

// GetHtmlBody returns the HtmlBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetHtmlBody() string {
	if o == nil || o.HtmlBody.Get() == nil {
		var ret string
		return ret
	}
	return *o.HtmlBody.Get()
}

// GetHtmlBodyOk returns a tuple with the HtmlBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetHtmlBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HtmlBody.Get(), o.HtmlBody.IsSet()
}

// HasHtmlBody returns a boolean if a field has been set.
func (o *CommentRequest) HasHtmlBody() bool {
	if o != nil && o.HtmlBody.IsSet() {
		return true
	}

	return false
}

// SetHtmlBody gets a reference to the given NullableString and assigns it to the HtmlBody field.
func (o *CommentRequest) SetHtmlBody(v string) {
	o.HtmlBody.Set(&v)
}
// SetHtmlBodyNil sets the value for HtmlBody to be an explicit nil
func (o *CommentRequest) SetHtmlBodyNil() {
	o.HtmlBody.Set(nil)
}

// UnsetHtmlBody ensures that no value is present for HtmlBody, not even an explicit nil
func (o *CommentRequest) UnsetHtmlBody() {
	o.HtmlBody.Unset()
}

// GetTicket returns the Ticket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetTicket() string {
	if o == nil || o.Ticket.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ticket.Get()
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetTicketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ticket.Get(), o.Ticket.IsSet()
}

// HasTicket returns a boolean if a field has been set.
func (o *CommentRequest) HasTicket() bool {
	if o != nil && o.Ticket.IsSet() {
		return true
	}

	return false
}

// SetTicket gets a reference to the given NullableString and assigns it to the Ticket field.
func (o *CommentRequest) SetTicket(v string) {
	o.Ticket.Set(&v)
}
// SetTicketNil sets the value for Ticket to be an explicit nil
func (o *CommentRequest) SetTicketNil() {
	o.Ticket.Set(nil)
}

// UnsetTicket ensures that no value is present for Ticket, not even an explicit nil
func (o *CommentRequest) UnsetTicket() {
	o.Ticket.Unset()
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetIsPrivate() bool {
	if o == nil || o.IsPrivate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate.Get()
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetIsPrivateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPrivate.Get(), o.IsPrivate.IsSet()
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *CommentRequest) HasIsPrivate() bool {
	if o != nil && o.IsPrivate.IsSet() {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given NullableBool and assigns it to the IsPrivate field.
func (o *CommentRequest) SetIsPrivate(v bool) {
	o.IsPrivate.Set(&v)
}
// SetIsPrivateNil sets the value for IsPrivate to be an explicit nil
func (o *CommentRequest) SetIsPrivateNil() {
	o.IsPrivate.Set(nil)
}

// UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
func (o *CommentRequest) UnsetIsPrivate() {
	o.IsPrivate.Unset()
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *CommentRequest) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *CommentRequest) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *CommentRequest) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *CommentRequest) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *CommentRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *CommentRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *CommentRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *CommentRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

func (o CommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.IntegrationParams != nil {
		toSerialize["integration_params"] = o.IntegrationParams
	}
	if o.LinkedAccountParams != nil {
		toSerialize["linked_account_params"] = o.LinkedAccountParams
	}
	return json.Marshal(toSerialize)
}

func (v *CommentRequest) UnmarshalJSON(src []byte) error {
    type CommentRequestUnmarshalTarget CommentRequest

	var intermediateResult CommentRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CommentRequest(intermediateResult)
	return nil
}
type NullableCommentRequest struct {
	value *CommentRequest
	isSet bool
}

func (v NullableCommentRequest) Get() *CommentRequest {
	return v.value
}

func (v *NullableCommentRequest) Set(val *CommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentRequest(val *CommentRequest) *NullableCommentRequest {
	return &NullableCommentRequest{value: val, isSet: true}
}

func (v NullableCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


