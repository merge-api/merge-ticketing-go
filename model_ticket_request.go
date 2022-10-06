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

// TicketRequest # The Ticket Object ### Description The `Ticket` object is used to represent a ticket or a task within a system.  ### Usage Example TODO
type TicketRequest struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The ticket's name.
	Name NullableString `json:"name,omitempty"`
	Assignees *[]string `json:"assignees,omitempty"`
	// The ticket's due date.
	DueDate NullableTime `json:"due_date,omitempty"`
	// The current status of the ticket.
	Status NullableTicketStatusEnum `json:"status,omitempty"`
	// The ticket's description.
	Description NullableString `json:"description,omitempty"`
	Project NullableString `json:"project,omitempty"`
	// The ticket's type.
	TicketType NullableString `json:"ticket_type,omitempty"`
	Account NullableString `json:"account,omitempty"`
	Contact NullableString `json:"contact,omitempty"`
	ParentTicket NullableString `json:"parent_ticket,omitempty"`
	Attachments *[]string `json:"attachments,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	// When the third party's ticket was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
	// When the third party's ticket was updated.
	RemoteUpdatedAt NullableTime `json:"remote_updated_at,omitempty"`
	// When the ticket was completed.
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	// The 3rd party url of the Ticket.
	TicketUrl NullableString `json:"ticket_url,omitempty"`
	// The priority or urgency of the Ticket. Possible values include: URGENT, HIGH, NORMAL, LOW - in cases where there is no clear mapping - the original value passed through.
	Priority NullablePriorityEnum `json:"priority,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewTicketRequest instantiates a new TicketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketRequest() *TicketRequest {
	this := TicketRequest{}
	return &this
}

// NewTicketRequestWithDefaults instantiates a new TicketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketRequestWithDefaults() *TicketRequest {
	this := TicketRequest{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *TicketRequest) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *TicketRequest) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *TicketRequest) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *TicketRequest) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TicketRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TicketRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TicketRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TicketRequest) UnsetName() {
	o.Name.Unset()
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *TicketRequest) GetAssignees() []string {
	if o == nil || o.Assignees == nil {
		var ret []string
		return ret
	}
	return *o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketRequest) GetAssigneesOk() (*[]string, bool) {
	if o == nil || o.Assignees == nil {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *TicketRequest) HasAssignees() bool {
	if o != nil && o.Assignees != nil {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []string and assigns it to the Assignees field.
func (o *TicketRequest) SetAssignees(v []string) {
	o.Assignees = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetDueDate() time.Time {
	if o == nil || o.DueDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetDueDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *TicketRequest) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *TicketRequest) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *TicketRequest) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *TicketRequest) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetStatus() TicketStatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret TicketStatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetStatusOk() (*TicketStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TicketRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableTicketStatusEnum and assigns it to the Status field.
func (o *TicketRequest) SetStatus(v TicketStatusEnum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TicketRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TicketRequest) UnsetStatus() {
	o.Status.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TicketRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TicketRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TicketRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TicketRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetProject() string {
	if o == nil || o.Project.Get() == nil {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *TicketRequest) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *TicketRequest) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *TicketRequest) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *TicketRequest) UnsetProject() {
	o.Project.Unset()
}

// GetTicketType returns the TicketType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetTicketType() string {
	if o == nil || o.TicketType.Get() == nil {
		var ret string
		return ret
	}
	return *o.TicketType.Get()
}

// GetTicketTypeOk returns a tuple with the TicketType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetTicketTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TicketType.Get(), o.TicketType.IsSet()
}

// HasTicketType returns a boolean if a field has been set.
func (o *TicketRequest) HasTicketType() bool {
	if o != nil && o.TicketType.IsSet() {
		return true
	}

	return false
}

// SetTicketType gets a reference to the given NullableString and assigns it to the TicketType field.
func (o *TicketRequest) SetTicketType(v string) {
	o.TicketType.Set(&v)
}
// SetTicketTypeNil sets the value for TicketType to be an explicit nil
func (o *TicketRequest) SetTicketTypeNil() {
	o.TicketType.Set(nil)
}

// UnsetTicketType ensures that no value is present for TicketType, not even an explicit nil
func (o *TicketRequest) UnsetTicketType() {
	o.TicketType.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *TicketRequest) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *TicketRequest) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *TicketRequest) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *TicketRequest) UnsetAccount() {
	o.Account.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetContact() string {
	if o == nil || o.Contact.Get() == nil {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetContactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *TicketRequest) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *TicketRequest) SetContact(v string) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *TicketRequest) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *TicketRequest) UnsetContact() {
	o.Contact.Unset()
}

// GetParentTicket returns the ParentTicket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetParentTicket() string {
	if o == nil || o.ParentTicket.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentTicket.Get()
}

// GetParentTicketOk returns a tuple with the ParentTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetParentTicketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentTicket.Get(), o.ParentTicket.IsSet()
}

// HasParentTicket returns a boolean if a field has been set.
func (o *TicketRequest) HasParentTicket() bool {
	if o != nil && o.ParentTicket.IsSet() {
		return true
	}

	return false
}

// SetParentTicket gets a reference to the given NullableString and assigns it to the ParentTicket field.
func (o *TicketRequest) SetParentTicket(v string) {
	o.ParentTicket.Set(&v)
}
// SetParentTicketNil sets the value for ParentTicket to be an explicit nil
func (o *TicketRequest) SetParentTicketNil() {
	o.ParentTicket.Set(nil)
}

// UnsetParentTicket ensures that no value is present for ParentTicket, not even an explicit nil
func (o *TicketRequest) UnsetParentTicket() {
	o.ParentTicket.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TicketRequest) GetAttachments() []string {
	if o == nil || o.Attachments == nil {
		var ret []string
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketRequest) GetAttachmentsOk() (*[]string, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TicketRequest) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *TicketRequest) SetAttachments(v []string) {
	o.Attachments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TicketRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketRequest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TicketRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TicketRequest) SetTags(v []string) {
	o.Tags = &v
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *TicketRequest) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *TicketRequest) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *TicketRequest) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *TicketRequest) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetRemoteUpdatedAt returns the RemoteUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetRemoteUpdatedAt() time.Time {
	if o == nil || o.RemoteUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteUpdatedAt.Get()
}

// GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetRemoteUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteUpdatedAt.Get(), o.RemoteUpdatedAt.IsSet()
}

// HasRemoteUpdatedAt returns a boolean if a field has been set.
func (o *TicketRequest) HasRemoteUpdatedAt() bool {
	if o != nil && o.RemoteUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteUpdatedAt gets a reference to the given NullableTime and assigns it to the RemoteUpdatedAt field.
func (o *TicketRequest) SetRemoteUpdatedAt(v time.Time) {
	o.RemoteUpdatedAt.Set(&v)
}
// SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil
func (o *TicketRequest) SetRemoteUpdatedAtNil() {
	o.RemoteUpdatedAt.Set(nil)
}

// UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
func (o *TicketRequest) UnsetRemoteUpdatedAt() {
	o.RemoteUpdatedAt.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *TicketRequest) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *TicketRequest) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *TicketRequest) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *TicketRequest) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetTicketUrl returns the TicketUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetTicketUrl() string {
	if o == nil || o.TicketUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.TicketUrl.Get()
}

// GetTicketUrlOk returns a tuple with the TicketUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetTicketUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TicketUrl.Get(), o.TicketUrl.IsSet()
}

// HasTicketUrl returns a boolean if a field has been set.
func (o *TicketRequest) HasTicketUrl() bool {
	if o != nil && o.TicketUrl.IsSet() {
		return true
	}

	return false
}

// SetTicketUrl gets a reference to the given NullableString and assigns it to the TicketUrl field.
func (o *TicketRequest) SetTicketUrl(v string) {
	o.TicketUrl.Set(&v)
}
// SetTicketUrlNil sets the value for TicketUrl to be an explicit nil
func (o *TicketRequest) SetTicketUrlNil() {
	o.TicketUrl.Set(nil)
}

// UnsetTicketUrl ensures that no value is present for TicketUrl, not even an explicit nil
func (o *TicketRequest) UnsetTicketUrl() {
	o.TicketUrl.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TicketRequest) GetPriority() PriorityEnum {
	if o == nil || o.Priority.Get() == nil {
		var ret PriorityEnum
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TicketRequest) GetPriorityOk() (*PriorityEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *TicketRequest) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullablePriorityEnum and assigns it to the Priority field.
func (o *TicketRequest) SetPriority(v PriorityEnum) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *TicketRequest) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *TicketRequest) UnsetPriority() {
	o.Priority.Unset()
}

func (o TicketRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Assignees != nil {
		toSerialize["assignees"] = o.Assignees
	}
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.TicketType.IsSet() {
		toSerialize["ticket_type"] = o.TicketType.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.ParentTicket.IsSet() {
		toSerialize["parent_ticket"] = o.ParentTicket.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.RemoteCreatedAt.IsSet() {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
	}
	if o.RemoteUpdatedAt.IsSet() {
		toSerialize["remote_updated_at"] = o.RemoteUpdatedAt.Get()
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.TicketUrl.IsSet() {
		toSerialize["ticket_url"] = o.TicketUrl.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	return json.Marshal(toSerialize)
}

func (v *TicketRequest) UnmarshalJSON(src []byte) error {
    type TicketRequestUnmarshalTarget TicketRequest

	var intermediateResult TicketRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = TicketRequest(intermediateResult)
	return nil
}
type NullableTicketRequest struct {
	value *TicketRequest
	isSet bool
}

func (v NullableTicketRequest) Get() *TicketRequest {
	return v.value
}

func (v *NullableTicketRequest) Set(val *TicketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketRequest(val *TicketRequest) *NullableTicketRequest {
	return &NullableTicketRequest{value: val, isSet: true}
}

func (v NullableTicketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


