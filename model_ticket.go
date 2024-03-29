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

// Ticket # The Ticket Object ### Description The `Ticket` object is used to represent a ticket or a task within a system.  ### Usage Example TODO
type Ticket struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The ticket's name.
	Name NullableString `json:"name,omitempty"`
	Assignees *[]string `json:"assignees,omitempty"`
	// The user who created this ticket.
	Creator NullableString `json:"creator,omitempty"`
	// The ticket's due date.
	DueDate NullableTime `json:"due_date,omitempty"`
	// The current status of the ticket.  * `OPEN` - OPEN * `CLOSED` - CLOSED * `IN_PROGRESS` - IN_PROGRESS * `ON_HOLD` - ON_HOLD
	Status NullableTicketStatusEnum `json:"status,omitempty"`
	// The ticket’s description. HTML version of description is mapped if supported by the third-party platform.
	Description NullableString `json:"description,omitempty"`
	// The project the ticket belongs to.
	Project NullableString `json:"project,omitempty"`
	Collections *[]string `json:"collections,omitempty"`
	// The ticket's type.
	TicketType NullableString `json:"ticket_type,omitempty"`
	// The account associated with the ticket.
	Account NullableString `json:"account,omitempty"`
	// The contact associated with the ticket.
	Contact NullableString `json:"contact,omitempty"`
	// The ticket's parent ticket.
	ParentTicket NullableString `json:"parent_ticket,omitempty"`
	Attachments *[]string `json:"attachments,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	// When the third party's ticket was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
	// When the third party's ticket was updated.
	RemoteUpdatedAt NullableTime `json:"remote_updated_at,omitempty"`
	// When the ticket was completed.
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
	// The 3rd party url of the Ticket.
	TicketUrl NullableString `json:"ticket_url,omitempty"`
	// The priority or urgency of the Ticket.  * `URGENT` - URGENT * `HIGH` - HIGH * `NORMAL` - NORMAL * `LOW` - LOW
	Priority NullablePriorityEnum `json:"priority,omitempty"`
	FieldMappings map[string]interface{} `json:"field_mappings,omitempty"`
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	RemoteFields *[]RemoteField `json:"remote_fields,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewTicket instantiates a new Ticket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicket() *Ticket {
	this := Ticket{}
	return &this
}

// NewTicketWithDefaults instantiates a new Ticket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketWithDefaults() *Ticket {
	this := Ticket{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Ticket) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ticket) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ticket) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Ticket) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Ticket) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Ticket) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Ticket) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Ticket) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Ticket) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Ticket) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Ticket) UnsetName() {
	o.Name.Unset()
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *Ticket) GetAssignees() []string {
	if o == nil || o.Assignees == nil {
		var ret []string
		return ret
	}
	return *o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetAssigneesOk() (*[]string, bool) {
	if o == nil || o.Assignees == nil {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *Ticket) HasAssignees() bool {
	if o != nil && o.Assignees != nil {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []string and assigns it to the Assignees field.
func (o *Ticket) SetAssignees(v []string) {
	o.Assignees = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetCreator() string {
	if o == nil || o.Creator.Get() == nil {
		var ret string
		return ret
	}
	return *o.Creator.Get()
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetCreatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Creator.Get(), o.Creator.IsSet()
}

// HasCreator returns a boolean if a field has been set.
func (o *Ticket) HasCreator() bool {
	if o != nil && o.Creator.IsSet() {
		return true
	}

	return false
}

// SetCreator gets a reference to the given NullableString and assigns it to the Creator field.
func (o *Ticket) SetCreator(v string) {
	o.Creator.Set(&v)
}
// SetCreatorNil sets the value for Creator to be an explicit nil
func (o *Ticket) SetCreatorNil() {
	o.Creator.Set(nil)
}

// UnsetCreator ensures that no value is present for Creator, not even an explicit nil
func (o *Ticket) UnsetCreator() {
	o.Creator.Unset()
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetDueDate() time.Time {
	if o == nil || o.DueDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetDueDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *Ticket) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *Ticket) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *Ticket) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *Ticket) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetStatus() TicketStatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret TicketStatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetStatusOk() (*TicketStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Ticket) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableTicketStatusEnum and assigns it to the Status field.
func (o *Ticket) SetStatus(v TicketStatusEnum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Ticket) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Ticket) UnsetStatus() {
	o.Status.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Ticket) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Ticket) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Ticket) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Ticket) UnsetDescription() {
	o.Description.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetProject() string {
	if o == nil || o.Project.Get() == nil {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *Ticket) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *Ticket) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *Ticket) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *Ticket) UnsetProject() {
	o.Project.Unset()
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *Ticket) GetCollections() []string {
	if o == nil || o.Collections == nil {
		var ret []string
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetCollectionsOk() (*[]string, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *Ticket) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []string and assigns it to the Collections field.
func (o *Ticket) SetCollections(v []string) {
	o.Collections = &v
}

// GetTicketType returns the TicketType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetTicketType() string {
	if o == nil || o.TicketType.Get() == nil {
		var ret string
		return ret
	}
	return *o.TicketType.Get()
}

// GetTicketTypeOk returns a tuple with the TicketType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetTicketTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TicketType.Get(), o.TicketType.IsSet()
}

// HasTicketType returns a boolean if a field has been set.
func (o *Ticket) HasTicketType() bool {
	if o != nil && o.TicketType.IsSet() {
		return true
	}

	return false
}

// SetTicketType gets a reference to the given NullableString and assigns it to the TicketType field.
func (o *Ticket) SetTicketType(v string) {
	o.TicketType.Set(&v)
}
// SetTicketTypeNil sets the value for TicketType to be an explicit nil
func (o *Ticket) SetTicketTypeNil() {
	o.TicketType.Set(nil)
}

// UnsetTicketType ensures that no value is present for TicketType, not even an explicit nil
func (o *Ticket) UnsetTicketType() {
	o.TicketType.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *Ticket) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *Ticket) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *Ticket) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *Ticket) UnsetAccount() {
	o.Account.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetContact() string {
	if o == nil || o.Contact.Get() == nil {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetContactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *Ticket) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *Ticket) SetContact(v string) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *Ticket) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *Ticket) UnsetContact() {
	o.Contact.Unset()
}

// GetParentTicket returns the ParentTicket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetParentTicket() string {
	if o == nil || o.ParentTicket.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentTicket.Get()
}

// GetParentTicketOk returns a tuple with the ParentTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetParentTicketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentTicket.Get(), o.ParentTicket.IsSet()
}

// HasParentTicket returns a boolean if a field has been set.
func (o *Ticket) HasParentTicket() bool {
	if o != nil && o.ParentTicket.IsSet() {
		return true
	}

	return false
}

// SetParentTicket gets a reference to the given NullableString and assigns it to the ParentTicket field.
func (o *Ticket) SetParentTicket(v string) {
	o.ParentTicket.Set(&v)
}
// SetParentTicketNil sets the value for ParentTicket to be an explicit nil
func (o *Ticket) SetParentTicketNil() {
	o.ParentTicket.Set(nil)
}

// UnsetParentTicket ensures that no value is present for ParentTicket, not even an explicit nil
func (o *Ticket) UnsetParentTicket() {
	o.ParentTicket.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *Ticket) GetAttachments() []string {
	if o == nil || o.Attachments == nil {
		var ret []string
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetAttachmentsOk() (*[]string, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Ticket) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *Ticket) SetAttachments(v []string) {
	o.Attachments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Ticket) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Ticket) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Ticket) SetTags(v []string) {
	o.Tags = &v
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *Ticket) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *Ticket) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *Ticket) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *Ticket) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetRemoteUpdatedAt returns the RemoteUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetRemoteUpdatedAt() time.Time {
	if o == nil || o.RemoteUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteUpdatedAt.Get()
}

// GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetRemoteUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteUpdatedAt.Get(), o.RemoteUpdatedAt.IsSet()
}

// HasRemoteUpdatedAt returns a boolean if a field has been set.
func (o *Ticket) HasRemoteUpdatedAt() bool {
	if o != nil && o.RemoteUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteUpdatedAt gets a reference to the given NullableTime and assigns it to the RemoteUpdatedAt field.
func (o *Ticket) SetRemoteUpdatedAt(v time.Time) {
	o.RemoteUpdatedAt.Set(&v)
}
// SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil
func (o *Ticket) SetRemoteUpdatedAtNil() {
	o.RemoteUpdatedAt.Set(nil)
}

// UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
func (o *Ticket) UnsetRemoteUpdatedAt() {
	o.RemoteUpdatedAt.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *Ticket) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *Ticket) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *Ticket) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *Ticket) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *Ticket) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *Ticket) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *Ticket) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

// GetTicketUrl returns the TicketUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetTicketUrl() string {
	if o == nil || o.TicketUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.TicketUrl.Get()
}

// GetTicketUrlOk returns a tuple with the TicketUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetTicketUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TicketUrl.Get(), o.TicketUrl.IsSet()
}

// HasTicketUrl returns a boolean if a field has been set.
func (o *Ticket) HasTicketUrl() bool {
	if o != nil && o.TicketUrl.IsSet() {
		return true
	}

	return false
}

// SetTicketUrl gets a reference to the given NullableString and assigns it to the TicketUrl field.
func (o *Ticket) SetTicketUrl(v string) {
	o.TicketUrl.Set(&v)
}
// SetTicketUrlNil sets the value for TicketUrl to be an explicit nil
func (o *Ticket) SetTicketUrlNil() {
	o.TicketUrl.Set(nil)
}

// UnsetTicketUrl ensures that no value is present for TicketUrl, not even an explicit nil
func (o *Ticket) UnsetTicketUrl() {
	o.TicketUrl.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetPriority() PriorityEnum {
	if o == nil || o.Priority.Get() == nil {
		var ret PriorityEnum
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetPriorityOk() (*PriorityEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *Ticket) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullablePriorityEnum and assigns it to the Priority field.
func (o *Ticket) SetPriority(v PriorityEnum) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *Ticket) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *Ticket) UnsetPriority() {
	o.Priority.Unset()
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetFieldMappings() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *Ticket) HasFieldMappings() bool {
	if o != nil && o.FieldMappings != nil {
		return true
	}

	return false
}

// SetFieldMappings gets a reference to the given map[string]interface{} and assigns it to the FieldMappings field.
func (o *Ticket) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Ticket) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Ticket) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Ticket) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ticket) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ticket) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Ticket) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Ticket) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetRemoteFields returns the RemoteFields field value if set, zero value otherwise.
func (o *Ticket) GetRemoteFields() []RemoteField {
	if o == nil || o.RemoteFields == nil {
		var ret []RemoteField
		return ret
	}
	return *o.RemoteFields
}

// GetRemoteFieldsOk returns a tuple with the RemoteFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetRemoteFieldsOk() (*[]RemoteField, bool) {
	if o == nil || o.RemoteFields == nil {
		return nil, false
	}
	return o.RemoteFields, true
}

// HasRemoteFields returns a boolean if a field has been set.
func (o *Ticket) HasRemoteFields() bool {
	if o != nil && o.RemoteFields != nil {
		return true
	}

	return false
}

// SetRemoteFields gets a reference to the given []RemoteField and assigns it to the RemoteFields field.
func (o *Ticket) SetRemoteFields(v []RemoteField) {
	o.RemoteFields = &v
}

func (o Ticket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Assignees != nil {
		toSerialize["assignees"] = o.Assignees
	}
	if o.Creator.IsSet() {
		toSerialize["creator"] = o.Creator.Get()
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
	if o.Collections != nil {
		toSerialize["collections"] = o.Collections
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
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	if o.TicketUrl.IsSet() {
		toSerialize["ticket_url"] = o.TicketUrl.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.FieldMappings != nil {
		toSerialize["field_mappings"] = o.FieldMappings
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	if o.RemoteFields != nil {
		toSerialize["remote_fields"] = o.RemoteFields
	}
	return json.Marshal(toSerialize)
}

func (v *Ticket) UnmarshalJSON(src []byte) error {
    type TicketUnmarshalTarget Ticket

	var intermediateResult TicketUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Ticket(intermediateResult)
	return nil
}
type NullableTicket struct {
	value *Ticket
	isSet bool
}

func (v NullableTicket) Get() *Ticket {
	return v.value
}

func (v *NullableTicket) Set(val *Ticket) {
	v.value = val
	v.isSet = true
}

func (v NullableTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicket(val *Ticket) *NullableTicket {
	return &NullableTicket{value: val, isSet: true}
}

func (v NullableTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


