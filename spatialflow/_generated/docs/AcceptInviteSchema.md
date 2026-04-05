# AcceptInviteSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**InviteId** | **string** |  | 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAcceptInviteSchema

`func NewAcceptInviteSchema(token string, inviteId string, ) *AcceptInviteSchema`

NewAcceptInviteSchema instantiates a new AcceptInviteSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptInviteSchemaWithDefaults

`func NewAcceptInviteSchemaWithDefaults() *AcceptInviteSchema`

NewAcceptInviteSchemaWithDefaults instantiates a new AcceptInviteSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AcceptInviteSchema) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AcceptInviteSchema) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AcceptInviteSchema) SetToken(v string)`

SetToken sets Token field to given value.


### GetInviteId

`func (o *AcceptInviteSchema) GetInviteId() string`

GetInviteId returns the InviteId field if non-nil, zero value otherwise.

### GetInviteIdOk

`func (o *AcceptInviteSchema) GetInviteIdOk() (*string, bool)`

GetInviteIdOk returns a tuple with the InviteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteId

`func (o *AcceptInviteSchema) SetInviteId(v string)`

SetInviteId sets InviteId field to given value.


### GetPassword

`func (o *AcceptInviteSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AcceptInviteSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AcceptInviteSchema) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AcceptInviteSchema) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AcceptInviteSchema) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AcceptInviteSchema) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


