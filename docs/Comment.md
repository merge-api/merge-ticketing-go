# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**User** | Pointer to **NullableString** | The author of the Comment, if the author is a User. | [optional] 
**Contact** | Pointer to **NullableString** | The author of the Comment, if the author is a Contact. | [optional] 
**Body** | Pointer to **NullableString** | The comment&#39;s text body. | [optional] 
**HtmlBody** | Pointer to **NullableString** | The comment&#39;s text body formatted as html. | [optional] 
**Ticket** | Pointer to **NullableString** | The ticket associated with the comment.  | [optional] 
**IsPrivate** | Pointer to **NullableBool** | Whether or not the comment is internal. | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the third party&#39;s comment was created. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

## Methods

### NewComment

`func NewComment() *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Comment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Comment) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Comment) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Comment) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Comment) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Comment) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Comment) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetUser

`func (o *Comment) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Comment) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Comment) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Comment) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Comment) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Comment) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetContact

`func (o *Comment) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Comment) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Comment) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Comment) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Comment) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Comment) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetBody

`func (o *Comment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Comment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Comment) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Comment) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Comment) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Comment) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetHtmlBody

`func (o *Comment) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *Comment) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *Comment) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *Comment) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### SetHtmlBodyNil

`func (o *Comment) SetHtmlBodyNil(b bool)`

 SetHtmlBodyNil sets the value for HtmlBody to be an explicit nil

### UnsetHtmlBody
`func (o *Comment) UnsetHtmlBody()`

UnsetHtmlBody ensures that no value is present for HtmlBody, not even an explicit nil
### GetTicket

`func (o *Comment) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *Comment) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *Comment) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *Comment) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### SetTicketNil

`func (o *Comment) SetTicketNil(b bool)`

 SetTicketNil sets the value for Ticket to be an explicit nil

### UnsetTicket
`func (o *Comment) UnsetTicket()`

UnsetTicket ensures that no value is present for Ticket, not even an explicit nil
### GetIsPrivate

`func (o *Comment) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Comment) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Comment) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *Comment) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### SetIsPrivateNil

`func (o *Comment) SetIsPrivateNil(b bool)`

 SetIsPrivateNil sets the value for IsPrivate to be an explicit nil

### UnsetIsPrivate
`func (o *Comment) UnsetIsPrivate()`

UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
### GetRemoteCreatedAt

`func (o *Comment) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *Comment) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *Comment) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *Comment) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *Comment) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *Comment) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Comment) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Comment) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Comment) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Comment) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Comment) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Comment) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Comment) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Comment) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Comment) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Comment) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *Comment) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Comment) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Comment) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Comment) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *Comment) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Comment) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Comment) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Comment) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Comment) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Comment) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


