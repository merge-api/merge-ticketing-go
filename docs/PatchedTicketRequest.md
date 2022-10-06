# PatchedTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The ticket&#39;s name. | [optional] 
**Assignees** | Pointer to **[]string** |  | [optional] 
**DueDate** | Pointer to **NullableTime** | The ticket&#39;s due date. | [optional] 
**Status** | Pointer to [**NullableTicketStatusEnum**](TicketStatusEnum.md) | The current status of the ticket. | [optional] 
**Description** | Pointer to **NullableString** | The ticket&#39;s description. | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**TicketType** | Pointer to **NullableString** | The ticket&#39;s type. | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Contact** | Pointer to **NullableString** |  | [optional] 
**ParentTicket** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the third party&#39;s ticket was created. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s ticket was updated. | [optional] 
**CompletedAt** | Pointer to **NullableTime** | When the ticket was completed. | [optional] 
**TicketUrl** | Pointer to **NullableString** | The 3rd party url of the Ticket. | [optional] 
**Priority** | Pointer to [**NullablePriorityEnum**](PriorityEnum.md) | The priority or urgency of the Ticket. Possible values include: URGENT, HIGH, NORMAL, LOW - in cases where there is no clear mapping - the original value passed through. | [optional] 

## Methods

### NewPatchedTicketRequest

`func NewPatchedTicketRequest() *PatchedTicketRequest`

NewPatchedTicketRequest instantiates a new PatchedTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTicketRequestWithDefaults

`func NewPatchedTicketRequestWithDefaults() *PatchedTicketRequest`

NewPatchedTicketRequestWithDefaults instantiates a new PatchedTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedTicketRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTicketRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTicketRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTicketRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedTicketRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedTicketRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssignees

`func (o *PatchedTicketRequest) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *PatchedTicketRequest) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *PatchedTicketRequest) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *PatchedTicketRequest) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetDueDate

`func (o *PatchedTicketRequest) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *PatchedTicketRequest) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *PatchedTicketRequest) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *PatchedTicketRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *PatchedTicketRequest) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *PatchedTicketRequest) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetStatus

`func (o *PatchedTicketRequest) GetStatus() TicketStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedTicketRequest) GetStatusOk() (*TicketStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedTicketRequest) SetStatus(v TicketStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedTicketRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchedTicketRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchedTicketRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *PatchedTicketRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTicketRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTicketRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTicketRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedTicketRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedTicketRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProject

`func (o *PatchedTicketRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedTicketRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedTicketRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedTicketRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *PatchedTicketRequest) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *PatchedTicketRequest) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetTicketType

`func (o *PatchedTicketRequest) GetTicketType() string`

GetTicketType returns the TicketType field if non-nil, zero value otherwise.

### GetTicketTypeOk

`func (o *PatchedTicketRequest) GetTicketTypeOk() (*string, bool)`

GetTicketTypeOk returns a tuple with the TicketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketType

`func (o *PatchedTicketRequest) SetTicketType(v string)`

SetTicketType sets TicketType field to given value.

### HasTicketType

`func (o *PatchedTicketRequest) HasTicketType() bool`

HasTicketType returns a boolean if a field has been set.

### SetTicketTypeNil

`func (o *PatchedTicketRequest) SetTicketTypeNil(b bool)`

 SetTicketTypeNil sets the value for TicketType to be an explicit nil

### UnsetTicketType
`func (o *PatchedTicketRequest) UnsetTicketType()`

UnsetTicketType ensures that no value is present for TicketType, not even an explicit nil
### GetAccount

`func (o *PatchedTicketRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PatchedTicketRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PatchedTicketRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PatchedTicketRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PatchedTicketRequest) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PatchedTicketRequest) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetContact

`func (o *PatchedTicketRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatchedTicketRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatchedTicketRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatchedTicketRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PatchedTicketRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PatchedTicketRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetParentTicket

`func (o *PatchedTicketRequest) GetParentTicket() string`

GetParentTicket returns the ParentTicket field if non-nil, zero value otherwise.

### GetParentTicketOk

`func (o *PatchedTicketRequest) GetParentTicketOk() (*string, bool)`

GetParentTicketOk returns a tuple with the ParentTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTicket

`func (o *PatchedTicketRequest) SetParentTicket(v string)`

SetParentTicket sets ParentTicket field to given value.

### HasParentTicket

`func (o *PatchedTicketRequest) HasParentTicket() bool`

HasParentTicket returns a boolean if a field has been set.

### SetParentTicketNil

`func (o *PatchedTicketRequest) SetParentTicketNil(b bool)`

 SetParentTicketNil sets the value for ParentTicket to be an explicit nil

### UnsetParentTicket
`func (o *PatchedTicketRequest) UnsetParentTicket()`

UnsetParentTicket ensures that no value is present for ParentTicket, not even an explicit nil
### GetTags

`func (o *PatchedTicketRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedTicketRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedTicketRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedTicketRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRemoteCreatedAt

`func (o *PatchedTicketRequest) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *PatchedTicketRequest) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *PatchedTicketRequest) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *PatchedTicketRequest) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *PatchedTicketRequest) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *PatchedTicketRequest) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *PatchedTicketRequest) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *PatchedTicketRequest) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *PatchedTicketRequest) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *PatchedTicketRequest) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *PatchedTicketRequest) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *PatchedTicketRequest) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetCompletedAt

`func (o *PatchedTicketRequest) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *PatchedTicketRequest) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *PatchedTicketRequest) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *PatchedTicketRequest) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *PatchedTicketRequest) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *PatchedTicketRequest) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetTicketUrl

`func (o *PatchedTicketRequest) GetTicketUrl() string`

GetTicketUrl returns the TicketUrl field if non-nil, zero value otherwise.

### GetTicketUrlOk

`func (o *PatchedTicketRequest) GetTicketUrlOk() (*string, bool)`

GetTicketUrlOk returns a tuple with the TicketUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketUrl

`func (o *PatchedTicketRequest) SetTicketUrl(v string)`

SetTicketUrl sets TicketUrl field to given value.

### HasTicketUrl

`func (o *PatchedTicketRequest) HasTicketUrl() bool`

HasTicketUrl returns a boolean if a field has been set.

### SetTicketUrlNil

`func (o *PatchedTicketRequest) SetTicketUrlNil(b bool)`

 SetTicketUrlNil sets the value for TicketUrl to be an explicit nil

### UnsetTicketUrl
`func (o *PatchedTicketRequest) UnsetTicketUrl()`

UnsetTicketUrl ensures that no value is present for TicketUrl, not even an explicit nil
### GetPriority

`func (o *PatchedTicketRequest) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedTicketRequest) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedTicketRequest) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedTicketRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *PatchedTicketRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *PatchedTicketRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

