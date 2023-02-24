# CommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **NullableString** | The author of the Comment, if the author is a User. | [optional] 
**Contact** | Pointer to **NullableString** | The author of the Comment, if the author is a Contact. | [optional] 
**Body** | Pointer to **NullableString** | The comment&#39;s text body. | [optional] 
**HtmlBody** | Pointer to **NullableString** | The comment&#39;s text body formatted as html. | [optional] 
**Ticket** | Pointer to **NullableString** | The ticket associated with the comment.  | [optional] 
**IsPrivate** | Pointer to **NullableBool** | Whether or not the comment is internal. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCommentRequest

`func NewCommentRequest() *CommentRequest`

NewCommentRequest instantiates a new CommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentRequestWithDefaults

`func NewCommentRequestWithDefaults() *CommentRequest`

NewCommentRequestWithDefaults instantiates a new CommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *CommentRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CommentRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CommentRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CommentRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *CommentRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CommentRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetContact

`func (o *CommentRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CommentRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CommentRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *CommentRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *CommentRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *CommentRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetBody

`func (o *CommentRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CommentRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CommentRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CommentRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *CommentRequest) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *CommentRequest) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetHtmlBody

`func (o *CommentRequest) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *CommentRequest) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *CommentRequest) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *CommentRequest) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### SetHtmlBodyNil

`func (o *CommentRequest) SetHtmlBodyNil(b bool)`

 SetHtmlBodyNil sets the value for HtmlBody to be an explicit nil

### UnsetHtmlBody
`func (o *CommentRequest) UnsetHtmlBody()`

UnsetHtmlBody ensures that no value is present for HtmlBody, not even an explicit nil
### GetTicket

`func (o *CommentRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *CommentRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *CommentRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *CommentRequest) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### SetTicketNil

`func (o *CommentRequest) SetTicketNil(b bool)`

 SetTicketNil sets the value for Ticket to be an explicit nil

### UnsetTicket
`func (o *CommentRequest) UnsetTicket()`

UnsetTicket ensures that no value is present for Ticket, not even an explicit nil
### GetIsPrivate

`func (o *CommentRequest) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *CommentRequest) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *CommentRequest) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *CommentRequest) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### SetIsPrivateNil

`func (o *CommentRequest) SetIsPrivateNil(b bool)`

 SetIsPrivateNil sets the value for IsPrivate to be an explicit nil

### UnsetIsPrivate
`func (o *CommentRequest) UnsetIsPrivate()`

UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
### GetIntegrationParams

`func (o *CommentRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *CommentRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *CommentRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *CommentRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *CommentRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *CommentRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *CommentRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *CommentRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *CommentRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *CommentRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *CommentRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *CommentRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


