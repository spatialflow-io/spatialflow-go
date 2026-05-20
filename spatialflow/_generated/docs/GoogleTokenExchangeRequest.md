# GoogleTokenExchangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdToken** | **string** |  | 
**InviteId** | Pointer to **NullableString** |  | [optional] 
**InviteToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGoogleTokenExchangeRequest

`func NewGoogleTokenExchangeRequest(idToken string, ) *GoogleTokenExchangeRequest`

NewGoogleTokenExchangeRequest instantiates a new GoogleTokenExchangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleTokenExchangeRequestWithDefaults

`func NewGoogleTokenExchangeRequestWithDefaults() *GoogleTokenExchangeRequest`

NewGoogleTokenExchangeRequestWithDefaults instantiates a new GoogleTokenExchangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdToken

`func (o *GoogleTokenExchangeRequest) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *GoogleTokenExchangeRequest) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *GoogleTokenExchangeRequest) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.


### GetInviteId

`func (o *GoogleTokenExchangeRequest) GetInviteId() string`

GetInviteId returns the InviteId field if non-nil, zero value otherwise.

### GetInviteIdOk

`func (o *GoogleTokenExchangeRequest) GetInviteIdOk() (*string, bool)`

GetInviteIdOk returns a tuple with the InviteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteId

`func (o *GoogleTokenExchangeRequest) SetInviteId(v string)`

SetInviteId sets InviteId field to given value.

### HasInviteId

`func (o *GoogleTokenExchangeRequest) HasInviteId() bool`

HasInviteId returns a boolean if a field has been set.

### SetInviteIdNil

`func (o *GoogleTokenExchangeRequest) SetInviteIdNil(b bool)`

 SetInviteIdNil sets the value for InviteId to be an explicit nil

### UnsetInviteId
`func (o *GoogleTokenExchangeRequest) UnsetInviteId()`

UnsetInviteId ensures that no value is present for InviteId, not even an explicit nil
### GetInviteToken

`func (o *GoogleTokenExchangeRequest) GetInviteToken() string`

GetInviteToken returns the InviteToken field if non-nil, zero value otherwise.

### GetInviteTokenOk

`func (o *GoogleTokenExchangeRequest) GetInviteTokenOk() (*string, bool)`

GetInviteTokenOk returns a tuple with the InviteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteToken

`func (o *GoogleTokenExchangeRequest) SetInviteToken(v string)`

SetInviteToken sets InviteToken field to given value.

### HasInviteToken

`func (o *GoogleTokenExchangeRequest) HasInviteToken() bool`

HasInviteToken returns a boolean if a field has been set.

### SetInviteTokenNil

`func (o *GoogleTokenExchangeRequest) SetInviteTokenNil(b bool)`

 SetInviteTokenNil sets the value for InviteToken to be an explicit nil

### UnsetInviteToken
`func (o *GoogleTokenExchangeRequest) UnsetInviteToken()`

UnsetInviteToken ensures that no value is present for InviteToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


