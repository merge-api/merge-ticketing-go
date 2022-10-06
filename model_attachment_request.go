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

// AttachmentRequest # The Attachment Object ### Description The `Attachment` object is used to represent an attachment for a ticket.  ### Usage Example TODO
type AttachmentRequest struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The attachment's name.
	FileName NullableString `json:"file_name,omitempty"`
	Ticket NullableString `json:"ticket,omitempty"`
	// The attachment's url.
	FileUrl NullableString `json:"file_url,omitempty"`
	// The attachment's file format.
	ContentType NullableString `json:"content_type,omitempty"`
	UploadedBy NullableString `json:"uploaded_by,omitempty"`
	// When the third party's attachment was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewAttachmentRequest instantiates a new AttachmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentRequest() *AttachmentRequest {
	this := AttachmentRequest{}
	return &this
}

// NewAttachmentRequestWithDefaults instantiates a new AttachmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentRequestWithDefaults() *AttachmentRequest {
	this := AttachmentRequest{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentRequest) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentRequest) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *AttachmentRequest) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *AttachmentRequest) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *AttachmentRequest) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *AttachmentRequest) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentRequest) GetFileName() string {
	if o == nil || o.FileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentRequest) GetFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *AttachmentRequest) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *AttachmentRequest) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *AttachmentRequest) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *AttachmentRequest) UnsetFileName() {
	o.FileName.Unset()
}

// GetTicket returns the Ticket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentRequest) GetTicket() string {
	if o == nil || o.Ticket.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ticket.Get()
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentRequest) GetTicketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ticket.Get(), o.Ticket.IsSet()
}

// HasTicket returns a boolean if a field has been set.
func (o *AttachmentRequest) HasTicket() bool {
	if o != nil && o.Ticket.IsSet() {
		return true
	}

	return false
}

// SetTicket gets a reference to the given NullableString and assigns it to the Ticket field.
func (o *AttachmentRequest) SetTicket(v string) {
	o.Ticket.Set(&v)
}
// SetTicketNil sets the value for Ticket to be an explicit nil
func (o *AttachmentRequest) SetTicketNil() {
	o.Ticket.Set(nil)
}

// UnsetTicket ensures that no value is present for Ticket, not even an explicit nil
func (o *AttachmentRequest) UnsetTicket() {
	o.Ticket.Unset()
}

// GetFileUrl returns the FileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentRequest) GetFileUrl() string {
	if o == nil || o.FileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileUrl.Get()
}

// GetFileUrlOk returns a tuple with the FileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentRequest) GetFileUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileUrl.Get(), o.FileUrl.IsSet()
}

// HasFileUrl returns a boolean if a field has been set.
func (o *AttachmentRequest) HasFileUrl() bool {
	if o != nil && o.FileUrl.IsSet() {
		return true
	}

	return false
}

// SetFileUrl gets a reference to the given NullableString and assigns it to the FileUrl field.
func (o *AttachmentRequest) SetFileUrl(v string) {
	o.FileUrl.Set(&v)
}
// SetFileUrlNil sets the value for FileUrl to be an explicit nil
func (o *AttachmentRequest) SetFileUrlNil() {
	o.FileUrl.Set(nil)
}

// UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
func (o *AttachmentRequest) UnsetFileUrl() {
	o.FileUrl.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentRequest) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentRequest) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *AttachmentRequest) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *AttachmentRequest) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *AttachmentRequest) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *AttachmentRequest) UnsetContentType() {
	o.ContentType.Unset()
}

// GetUploadedBy returns the UploadedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentRequest) GetUploadedBy() string {
	if o == nil || o.UploadedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UploadedBy.Get()
}

// GetUploadedByOk returns a tuple with the UploadedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentRequest) GetUploadedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UploadedBy.Get(), o.UploadedBy.IsSet()
}

// HasUploadedBy returns a boolean if a field has been set.
func (o *AttachmentRequest) HasUploadedBy() bool {
	if o != nil && o.UploadedBy.IsSet() {
		return true
	}

	return false
}

// SetUploadedBy gets a reference to the given NullableString and assigns it to the UploadedBy field.
func (o *AttachmentRequest) SetUploadedBy(v string) {
	o.UploadedBy.Set(&v)
}
// SetUploadedByNil sets the value for UploadedBy to be an explicit nil
func (o *AttachmentRequest) SetUploadedByNil() {
	o.UploadedBy.Set(nil)
}

// UnsetUploadedBy ensures that no value is present for UploadedBy, not even an explicit nil
func (o *AttachmentRequest) UnsetUploadedBy() {
	o.UploadedBy.Unset()
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentRequest) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentRequest) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *AttachmentRequest) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *AttachmentRequest) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *AttachmentRequest) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *AttachmentRequest) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

func (o AttachmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.FileName.IsSet() {
		toSerialize["file_name"] = o.FileName.Get()
	}
	if o.Ticket.IsSet() {
		toSerialize["ticket"] = o.Ticket.Get()
	}
	if o.FileUrl.IsSet() {
		toSerialize["file_url"] = o.FileUrl.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.UploadedBy.IsSet() {
		toSerialize["uploaded_by"] = o.UploadedBy.Get()
	}
	if o.RemoteCreatedAt.IsSet() {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
	}
	return json.Marshal(toSerialize)
}

func (v *AttachmentRequest) UnmarshalJSON(src []byte) error {
    type AttachmentRequestUnmarshalTarget AttachmentRequest

	var intermediateResult AttachmentRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = AttachmentRequest(intermediateResult)
	return nil
}
type NullableAttachmentRequest struct {
	value *AttachmentRequest
	isSet bool
}

func (v NullableAttachmentRequest) Get() *AttachmentRequest {
	return v.value
}

func (v *NullableAttachmentRequest) Set(val *AttachmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentRequest(val *AttachmentRequest) *NullableAttachmentRequest {
	return &NullableAttachmentRequest{value: val, isSet: true}
}

func (v NullableAttachmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


