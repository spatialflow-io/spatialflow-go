# InvitationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**Role** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | **string** |  | 
**ExpiresAt** | **string** |  | 
**InvitedByEmail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvitationOut

`func NewInvitationOut(id string, email string, role string, status string, createdAt string, expiresAt string, ) *InvitationOut`

NewInvitationOut instantiates a new InvitationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationOutWithDefaults

`func NewInvitationOutWithDefaults() *InvitationOut`

NewInvitationOutWithDefaults instantiates a new InvitationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvitationOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitationOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitationOut) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *InvitationOut) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitationOut) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitationOut) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InvitationOut) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InvitationOut) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InvitationOut) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *InvitationOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvitationOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvitationOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *InvitationOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvitationOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvitationOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *InvitationOut) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InvitationOut) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InvitationOut) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetInvitedByEmail

`func (o *InvitationOut) GetInvitedByEmail() string`

GetInvitedByEmail returns the InvitedByEmail field if non-nil, zero value otherwise.

### GetInvitedByEmailOk

`func (o *InvitationOut) GetInvitedByEmailOk() (*string, bool)`

GetInvitedByEmailOk returns a tuple with the InvitedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedByEmail

`func (o *InvitationOut) SetInvitedByEmail(v string)`

SetInvitedByEmail sets InvitedByEmail field to given value.

### HasInvitedByEmail

`func (o *InvitationOut) HasInvitedByEmail() bool`

HasInvitedByEmail returns a boolean if a field has been set.

### SetInvitedByEmailNil

`func (o *InvitationOut) SetInvitedByEmailNil(b bool)`

 SetInvitedByEmailNil sets the value for InvitedByEmail to be an explicit nil

### UnsetInvitedByEmail
`func (o *InvitationOut) UnsetInvitedByEmail()`

UnsetInvitedByEmail ensures that no value is present for InvitedByEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


