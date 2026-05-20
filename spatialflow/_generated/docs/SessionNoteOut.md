# SessionNoteOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionId** | **string** |  | 
**AuthorId** | Pointer to **NullableString** |  | [optional] 
**AuthorName** | **string** |  | 
**AuthorRole** | **string** |  | 
**Body** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewSessionNoteOut

`func NewSessionNoteOut(id string, sessionId string, authorName string, authorRole string, body string, createdAt time.Time, ) *SessionNoteOut`

NewSessionNoteOut instantiates a new SessionNoteOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionNoteOutWithDefaults

`func NewSessionNoteOutWithDefaults() *SessionNoteOut`

NewSessionNoteOutWithDefaults instantiates a new SessionNoteOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionNoteOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionNoteOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionNoteOut) SetId(v string)`

SetId sets Id field to given value.


### GetSessionId

`func (o *SessionNoteOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionNoteOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionNoteOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetAuthorId

`func (o *SessionNoteOut) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *SessionNoteOut) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *SessionNoteOut) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *SessionNoteOut) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### SetAuthorIdNil

`func (o *SessionNoteOut) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *SessionNoteOut) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
### GetAuthorName

`func (o *SessionNoteOut) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *SessionNoteOut) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *SessionNoteOut) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetAuthorRole

`func (o *SessionNoteOut) GetAuthorRole() string`

GetAuthorRole returns the AuthorRole field if non-nil, zero value otherwise.

### GetAuthorRoleOk

`func (o *SessionNoteOut) GetAuthorRoleOk() (*string, bool)`

GetAuthorRoleOk returns a tuple with the AuthorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorRole

`func (o *SessionNoteOut) SetAuthorRole(v string)`

SetAuthorRole sets AuthorRole field to given value.


### GetBody

`func (o *SessionNoteOut) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SessionNoteOut) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SessionNoteOut) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *SessionNoteOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionNoteOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionNoteOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


