# UserActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**UserId** | **string** |  | 
**Email** | **string** |  | 
**EmailVerified** | Pointer to **NullableBool** |  | [optional] 
**AdminApproved** | Pointer to **NullableBool** |  | [optional] 
**AdminApprovedAt** | Pointer to **NullableString** |  | [optional] 
**PasswordSet** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUserActionResponse

`func NewUserActionResponse(message string, userId string, email string, ) *UserActionResponse`

NewUserActionResponse instantiates a new UserActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionResponseWithDefaults

`func NewUserActionResponseWithDefaults() *UserActionResponse`

NewUserActionResponseWithDefaults instantiates a new UserActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UserActionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserActionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserActionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUserId

`func (o *UserActionResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserActionResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserActionResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *UserActionResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserActionResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserActionResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailVerified

`func (o *UserActionResponse) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserActionResponse) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserActionResponse) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *UserActionResponse) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *UserActionResponse) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *UserActionResponse) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetAdminApproved

`func (o *UserActionResponse) GetAdminApproved() bool`

GetAdminApproved returns the AdminApproved field if non-nil, zero value otherwise.

### GetAdminApprovedOk

`func (o *UserActionResponse) GetAdminApprovedOk() (*bool, bool)`

GetAdminApprovedOk returns a tuple with the AdminApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApproved

`func (o *UserActionResponse) SetAdminApproved(v bool)`

SetAdminApproved sets AdminApproved field to given value.

### HasAdminApproved

`func (o *UserActionResponse) HasAdminApproved() bool`

HasAdminApproved returns a boolean if a field has been set.

### SetAdminApprovedNil

`func (o *UserActionResponse) SetAdminApprovedNil(b bool)`

 SetAdminApprovedNil sets the value for AdminApproved to be an explicit nil

### UnsetAdminApproved
`func (o *UserActionResponse) UnsetAdminApproved()`

UnsetAdminApproved ensures that no value is present for AdminApproved, not even an explicit nil
### GetAdminApprovedAt

`func (o *UserActionResponse) GetAdminApprovedAt() string`

GetAdminApprovedAt returns the AdminApprovedAt field if non-nil, zero value otherwise.

### GetAdminApprovedAtOk

`func (o *UserActionResponse) GetAdminApprovedAtOk() (*string, bool)`

GetAdminApprovedAtOk returns a tuple with the AdminApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApprovedAt

`func (o *UserActionResponse) SetAdminApprovedAt(v string)`

SetAdminApprovedAt sets AdminApprovedAt field to given value.

### HasAdminApprovedAt

`func (o *UserActionResponse) HasAdminApprovedAt() bool`

HasAdminApprovedAt returns a boolean if a field has been set.

### SetAdminApprovedAtNil

`func (o *UserActionResponse) SetAdminApprovedAtNil(b bool)`

 SetAdminApprovedAtNil sets the value for AdminApprovedAt to be an explicit nil

### UnsetAdminApprovedAt
`func (o *UserActionResponse) UnsetAdminApprovedAt()`

UnsetAdminApprovedAt ensures that no value is present for AdminApprovedAt, not even an explicit nil
### GetPasswordSet

`func (o *UserActionResponse) GetPasswordSet() bool`

GetPasswordSet returns the PasswordSet field if non-nil, zero value otherwise.

### GetPasswordSetOk

`func (o *UserActionResponse) GetPasswordSetOk() (*bool, bool)`

GetPasswordSetOk returns a tuple with the PasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSet

`func (o *UserActionResponse) SetPasswordSet(v bool)`

SetPasswordSet sets PasswordSet field to given value.

### HasPasswordSet

`func (o *UserActionResponse) HasPasswordSet() bool`

HasPasswordSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


