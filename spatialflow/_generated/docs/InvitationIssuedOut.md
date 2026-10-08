# InvitationIssuedOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**Role** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ExpiresAt** | **time.Time** |  | 
**InvitedByEmail** | Pointer to **NullableString** |  | [optional] 
**InviteLink** | **string** |  | 
**InviteTtlHours** | **int32** |  | 

## Methods

### NewInvitationIssuedOut

`func NewInvitationIssuedOut(id string, email string, role string, status string, createdAt time.Time, expiresAt time.Time, inviteLink string, inviteTtlHours int32, ) *InvitationIssuedOut`

NewInvitationIssuedOut instantiates a new InvitationIssuedOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationIssuedOutWithDefaults

`func NewInvitationIssuedOutWithDefaults() *InvitationIssuedOut`

NewInvitationIssuedOutWithDefaults instantiates a new InvitationIssuedOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvitationIssuedOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitationIssuedOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitationIssuedOut) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *InvitationIssuedOut) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitationIssuedOut) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitationIssuedOut) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InvitationIssuedOut) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InvitationIssuedOut) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InvitationIssuedOut) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *InvitationIssuedOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvitationIssuedOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvitationIssuedOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *InvitationIssuedOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvitationIssuedOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvitationIssuedOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *InvitationIssuedOut) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InvitationIssuedOut) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InvitationIssuedOut) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetInvitedByEmail

`func (o *InvitationIssuedOut) GetInvitedByEmail() string`

GetInvitedByEmail returns the InvitedByEmail field if non-nil, zero value otherwise.

### GetInvitedByEmailOk

`func (o *InvitationIssuedOut) GetInvitedByEmailOk() (*string, bool)`

GetInvitedByEmailOk returns a tuple with the InvitedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedByEmail

`func (o *InvitationIssuedOut) SetInvitedByEmail(v string)`

SetInvitedByEmail sets InvitedByEmail field to given value.

### HasInvitedByEmail

`func (o *InvitationIssuedOut) HasInvitedByEmail() bool`

HasInvitedByEmail returns a boolean if a field has been set.

### SetInvitedByEmailNil

`func (o *InvitationIssuedOut) SetInvitedByEmailNil(b bool)`

 SetInvitedByEmailNil sets the value for InvitedByEmail to be an explicit nil

### UnsetInvitedByEmail
`func (o *InvitationIssuedOut) UnsetInvitedByEmail()`

UnsetInvitedByEmail ensures that no value is present for InvitedByEmail, not even an explicit nil
### GetInviteLink

`func (o *InvitationIssuedOut) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *InvitationIssuedOut) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *InvitationIssuedOut) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.


### GetInviteTtlHours

`func (o *InvitationIssuedOut) GetInviteTtlHours() int32`

GetInviteTtlHours returns the InviteTtlHours field if non-nil, zero value otherwise.

### GetInviteTtlHoursOk

`func (o *InvitationIssuedOut) GetInviteTtlHoursOk() (*int32, bool)`

GetInviteTtlHoursOk returns a tuple with the InviteTtlHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteTtlHours

`func (o *InvitationIssuedOut) SetInviteTtlHours(v int32)`

SetInviteTtlHours sets InviteTtlHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


