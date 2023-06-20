# Ticket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Name** | Pointer to **NullableString** | The ticket&#39;s name. | [optional] 
**Assignees** | Pointer to **[]string** |  | [optional] 
**Creator** | Pointer to **NullableString** | The user who created this ticket. | [optional] 
**DueDate** | Pointer to **NullableTime** | The ticket&#39;s due date. | [optional] 
**Status** | Pointer to [**NullableTicketStatusEnum**](TicketStatusEnum.md) | The current status of the ticket.  * &#x60;OPEN&#x60; - OPEN * &#x60;CLOSED&#x60; - CLOSED * &#x60;IN_PROGRESS&#x60; - IN_PROGRESS * &#x60;ON_HOLD&#x60; - ON_HOLD | [optional] 
**Description** | Pointer to **NullableString** | The ticket’s description. HTML version of description is mapped if supported by the third-party platform. | [optional] 
**Project** | Pointer to **NullableString** | The project the ticket belongs to. | [optional] 
**Collections** | Pointer to **[]string** |  | [optional] 
**TicketType** | Pointer to **NullableString** | The ticket&#39;s type. | [optional] 
**Account** | Pointer to **NullableString** | The account associated with the ticket. | [optional] 
**Contact** | Pointer to **NullableString** | The contact associated with the ticket. | [optional] 
**ParentTicket** | Pointer to **NullableString** | The ticket&#39;s parent ticket. | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the third party&#39;s ticket was created. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s ticket was updated. | [optional] 
**CompletedAt** | Pointer to **NullableTime** | When the ticket was completed. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 
**TicketUrl** | Pointer to **NullableString** | The 3rd party url of the Ticket. | [optional] 
**Priority** | Pointer to [**NullablePriorityEnum**](PriorityEnum.md) | The priority or urgency of the Ticket.  * &#x60;URGENT&#x60; - URGENT * &#x60;HIGH&#x60; - HIGH * &#x60;NORMAL&#x60; - NORMAL * &#x60;LOW&#x60; - LOW | [optional] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**RemoteFields** | Pointer to [**[]RemoteField**](RemoteField.md) |  | [optional] [readonly] 

## Methods

### NewTicket

`func NewTicket() *Ticket`

NewTicket instantiates a new Ticket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketWithDefaults

`func NewTicketWithDefaults() *Ticket`

NewTicketWithDefaults instantiates a new Ticket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ticket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ticket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ticket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ticket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Ticket) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Ticket) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Ticket) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Ticket) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Ticket) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Ticket) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetName

`func (o *Ticket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ticket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ticket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ticket) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Ticket) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Ticket) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssignees

`func (o *Ticket) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *Ticket) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *Ticket) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *Ticket) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetCreator

`func (o *Ticket) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Ticket) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Ticket) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Ticket) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *Ticket) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *Ticket) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetDueDate

`func (o *Ticket) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Ticket) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Ticket) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Ticket) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *Ticket) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *Ticket) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetStatus

`func (o *Ticket) GetStatus() TicketStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ticket) GetStatusOk() (*TicketStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ticket) SetStatus(v TicketStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ticket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Ticket) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Ticket) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *Ticket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ticket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Ticket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Ticket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Ticket) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Ticket) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProject

`func (o *Ticket) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Ticket) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Ticket) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Ticket) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *Ticket) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *Ticket) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetCollections

`func (o *Ticket) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *Ticket) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *Ticket) SetCollections(v []string)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *Ticket) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetTicketType

`func (o *Ticket) GetTicketType() string`

GetTicketType returns the TicketType field if non-nil, zero value otherwise.

### GetTicketTypeOk

`func (o *Ticket) GetTicketTypeOk() (*string, bool)`

GetTicketTypeOk returns a tuple with the TicketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketType

`func (o *Ticket) SetTicketType(v string)`

SetTicketType sets TicketType field to given value.

### HasTicketType

`func (o *Ticket) HasTicketType() bool`

HasTicketType returns a boolean if a field has been set.

### SetTicketTypeNil

`func (o *Ticket) SetTicketTypeNil(b bool)`

 SetTicketTypeNil sets the value for TicketType to be an explicit nil

### UnsetTicketType
`func (o *Ticket) UnsetTicketType()`

UnsetTicketType ensures that no value is present for TicketType, not even an explicit nil
### GetAccount

`func (o *Ticket) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Ticket) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Ticket) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Ticket) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *Ticket) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Ticket) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetContact

`func (o *Ticket) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Ticket) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Ticket) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Ticket) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Ticket) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Ticket) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetParentTicket

`func (o *Ticket) GetParentTicket() string`

GetParentTicket returns the ParentTicket field if non-nil, zero value otherwise.

### GetParentTicketOk

`func (o *Ticket) GetParentTicketOk() (*string, bool)`

GetParentTicketOk returns a tuple with the ParentTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTicket

`func (o *Ticket) SetParentTicket(v string)`

SetParentTicket sets ParentTicket field to given value.

### HasParentTicket

`func (o *Ticket) HasParentTicket() bool`

HasParentTicket returns a boolean if a field has been set.

### SetParentTicketNil

`func (o *Ticket) SetParentTicketNil(b bool)`

 SetParentTicketNil sets the value for ParentTicket to be an explicit nil

### UnsetParentTicket
`func (o *Ticket) UnsetParentTicket()`

UnsetParentTicket ensures that no value is present for ParentTicket, not even an explicit nil
### GetAttachments

`func (o *Ticket) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Ticket) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Ticket) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Ticket) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetTags

`func (o *Ticket) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Ticket) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Ticket) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Ticket) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRemoteCreatedAt

`func (o *Ticket) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *Ticket) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *Ticket) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *Ticket) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *Ticket) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *Ticket) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *Ticket) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *Ticket) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *Ticket) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *Ticket) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *Ticket) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *Ticket) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetCompletedAt

`func (o *Ticket) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Ticket) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Ticket) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Ticket) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *Ticket) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Ticket) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Ticket) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Ticket) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Ticket) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Ticket) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetTicketUrl

`func (o *Ticket) GetTicketUrl() string`

GetTicketUrl returns the TicketUrl field if non-nil, zero value otherwise.

### GetTicketUrlOk

`func (o *Ticket) GetTicketUrlOk() (*string, bool)`

GetTicketUrlOk returns a tuple with the TicketUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketUrl

`func (o *Ticket) SetTicketUrl(v string)`

SetTicketUrl sets TicketUrl field to given value.

### HasTicketUrl

`func (o *Ticket) HasTicketUrl() bool`

HasTicketUrl returns a boolean if a field has been set.

### SetTicketUrlNil

`func (o *Ticket) SetTicketUrlNil(b bool)`

 SetTicketUrlNil sets the value for TicketUrl to be an explicit nil

### UnsetTicketUrl
`func (o *Ticket) UnsetTicketUrl()`

UnsetTicketUrl ensures that no value is present for TicketUrl, not even an explicit nil
### GetPriority

`func (o *Ticket) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Ticket) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Ticket) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Ticket) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *Ticket) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *Ticket) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetFieldMappings

`func (o *Ticket) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Ticket) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Ticket) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Ticket) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Ticket) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Ticket) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *Ticket) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Ticket) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Ticket) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Ticket) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *Ticket) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Ticket) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Ticket) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Ticket) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Ticket) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Ticket) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetRemoteFields

`func (o *Ticket) GetRemoteFields() []RemoteField`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *Ticket) GetRemoteFieldsOk() (*[]RemoteField, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *Ticket) SetRemoteFields(v []RemoteField)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *Ticket) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


