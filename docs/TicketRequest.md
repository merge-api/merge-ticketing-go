# TicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**CompletedAt** | Pointer to **NullableTime** | When the ticket was completed. | [optional] 
**TicketUrl** | Pointer to **NullableString** | The 3rd party url of the Ticket. | [optional] 
**Priority** | Pointer to [**NullablePriorityEnum**](PriorityEnum.md) | The priority or urgency of the Ticket.  * &#x60;URGENT&#x60; - URGENT * &#x60;HIGH&#x60; - HIGH * &#x60;NORMAL&#x60; - NORMAL * &#x60;LOW&#x60; - LOW | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 
**RemoteFields** | Pointer to [**[]RemoteFieldRequest**](RemoteFieldRequest.md) |  | [optional] 

## Methods

### NewTicketRequest

`func NewTicketRequest() *TicketRequest`

NewTicketRequest instantiates a new TicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketRequestWithDefaults

`func NewTicketRequestWithDefaults() *TicketRequest`

NewTicketRequestWithDefaults instantiates a new TicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TicketRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TicketRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TicketRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TicketRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TicketRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TicketRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssignees

`func (o *TicketRequest) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *TicketRequest) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *TicketRequest) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *TicketRequest) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetCreator

`func (o *TicketRequest) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *TicketRequest) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *TicketRequest) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *TicketRequest) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *TicketRequest) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *TicketRequest) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetDueDate

`func (o *TicketRequest) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TicketRequest) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TicketRequest) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TicketRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *TicketRequest) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *TicketRequest) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetStatus

`func (o *TicketRequest) GetStatus() TicketStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TicketRequest) GetStatusOk() (*TicketStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TicketRequest) SetStatus(v TicketStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TicketRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TicketRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TicketRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *TicketRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TicketRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TicketRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TicketRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TicketRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TicketRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProject

`func (o *TicketRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TicketRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TicketRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *TicketRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *TicketRequest) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *TicketRequest) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetCollections

`func (o *TicketRequest) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *TicketRequest) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *TicketRequest) SetCollections(v []string)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *TicketRequest) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetTicketType

`func (o *TicketRequest) GetTicketType() string`

GetTicketType returns the TicketType field if non-nil, zero value otherwise.

### GetTicketTypeOk

`func (o *TicketRequest) GetTicketTypeOk() (*string, bool)`

GetTicketTypeOk returns a tuple with the TicketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketType

`func (o *TicketRequest) SetTicketType(v string)`

SetTicketType sets TicketType field to given value.

### HasTicketType

`func (o *TicketRequest) HasTicketType() bool`

HasTicketType returns a boolean if a field has been set.

### SetTicketTypeNil

`func (o *TicketRequest) SetTicketTypeNil(b bool)`

 SetTicketTypeNil sets the value for TicketType to be an explicit nil

### UnsetTicketType
`func (o *TicketRequest) UnsetTicketType()`

UnsetTicketType ensures that no value is present for TicketType, not even an explicit nil
### GetAccount

`func (o *TicketRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TicketRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TicketRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TicketRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *TicketRequest) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *TicketRequest) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetContact

`func (o *TicketRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *TicketRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *TicketRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *TicketRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *TicketRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *TicketRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetParentTicket

`func (o *TicketRequest) GetParentTicket() string`

GetParentTicket returns the ParentTicket field if non-nil, zero value otherwise.

### GetParentTicketOk

`func (o *TicketRequest) GetParentTicketOk() (*string, bool)`

GetParentTicketOk returns a tuple with the ParentTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTicket

`func (o *TicketRequest) SetParentTicket(v string)`

SetParentTicket sets ParentTicket field to given value.

### HasParentTicket

`func (o *TicketRequest) HasParentTicket() bool`

HasParentTicket returns a boolean if a field has been set.

### SetParentTicketNil

`func (o *TicketRequest) SetParentTicketNil(b bool)`

 SetParentTicketNil sets the value for ParentTicket to be an explicit nil

### UnsetParentTicket
`func (o *TicketRequest) UnsetParentTicket()`

UnsetParentTicket ensures that no value is present for ParentTicket, not even an explicit nil
### GetAttachments

`func (o *TicketRequest) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TicketRequest) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TicketRequest) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TicketRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetTags

`func (o *TicketRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TicketRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TicketRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TicketRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCompletedAt

`func (o *TicketRequest) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TicketRequest) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TicketRequest) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *TicketRequest) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *TicketRequest) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *TicketRequest) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetTicketUrl

`func (o *TicketRequest) GetTicketUrl() string`

GetTicketUrl returns the TicketUrl field if non-nil, zero value otherwise.

### GetTicketUrlOk

`func (o *TicketRequest) GetTicketUrlOk() (*string, bool)`

GetTicketUrlOk returns a tuple with the TicketUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketUrl

`func (o *TicketRequest) SetTicketUrl(v string)`

SetTicketUrl sets TicketUrl field to given value.

### HasTicketUrl

`func (o *TicketRequest) HasTicketUrl() bool`

HasTicketUrl returns a boolean if a field has been set.

### SetTicketUrlNil

`func (o *TicketRequest) SetTicketUrlNil(b bool)`

 SetTicketUrlNil sets the value for TicketUrl to be an explicit nil

### UnsetTicketUrl
`func (o *TicketRequest) UnsetTicketUrl()`

UnsetTicketUrl ensures that no value is present for TicketUrl, not even an explicit nil
### GetPriority

`func (o *TicketRequest) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TicketRequest) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TicketRequest) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TicketRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *TicketRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *TicketRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetIntegrationParams

`func (o *TicketRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *TicketRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *TicketRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *TicketRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *TicketRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *TicketRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *TicketRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *TicketRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *TicketRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *TicketRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *TicketRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *TicketRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil
### GetRemoteFields

`func (o *TicketRequest) GetRemoteFields() []RemoteFieldRequest`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *TicketRequest) GetRemoteFieldsOk() (*[]RemoteFieldRequest, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *TicketRequest) SetRemoteFields(v []RemoteFieldRequest)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *TicketRequest) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


