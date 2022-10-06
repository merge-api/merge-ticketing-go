# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**FileName** | Pointer to **NullableString** | The attachment&#39;s name. | [optional] 
**Ticket** | Pointer to **NullableString** |  | [optional] 
**FileUrl** | Pointer to **NullableString** | The attachment&#39;s url. | [optional] 
**ContentType** | Pointer to **NullableString** | The attachment&#39;s file format. | [optional] 
**UploadedBy** | Pointer to **NullableString** |  | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the third party&#39;s attachment was created. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewAttachment

`func NewAttachment() *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Attachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Attachment) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Attachment) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Attachment) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Attachment) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Attachment) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Attachment) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetFileName

`func (o *Attachment) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Attachment) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Attachment) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Attachment) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *Attachment) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *Attachment) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetTicket

`func (o *Attachment) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *Attachment) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *Attachment) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *Attachment) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### SetTicketNil

`func (o *Attachment) SetTicketNil(b bool)`

 SetTicketNil sets the value for Ticket to be an explicit nil

### UnsetTicket
`func (o *Attachment) UnsetTicket()`

UnsetTicket ensures that no value is present for Ticket, not even an explicit nil
### GetFileUrl

`func (o *Attachment) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *Attachment) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *Attachment) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *Attachment) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *Attachment) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *Attachment) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetContentType

`func (o *Attachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Attachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Attachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Attachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *Attachment) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *Attachment) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetUploadedBy

`func (o *Attachment) GetUploadedBy() string`

GetUploadedBy returns the UploadedBy field if non-nil, zero value otherwise.

### GetUploadedByOk

`func (o *Attachment) GetUploadedByOk() (*string, bool)`

GetUploadedByOk returns a tuple with the UploadedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedBy

`func (o *Attachment) SetUploadedBy(v string)`

SetUploadedBy sets UploadedBy field to given value.

### HasUploadedBy

`func (o *Attachment) HasUploadedBy() bool`

HasUploadedBy returns a boolean if a field has been set.

### SetUploadedByNil

`func (o *Attachment) SetUploadedByNil(b bool)`

 SetUploadedByNil sets the value for UploadedBy to be an explicit nil

### UnsetUploadedBy
`func (o *Attachment) UnsetUploadedBy()`

UnsetUploadedBy ensures that no value is present for UploadedBy, not even an explicit nil
### GetRemoteCreatedAt

`func (o *Attachment) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *Attachment) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *Attachment) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *Attachment) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *Attachment) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *Attachment) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteData

`func (o *Attachment) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Attachment) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Attachment) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Attachment) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Attachment) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Attachment) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Attachment) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Attachment) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Attachment) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Attachment) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


