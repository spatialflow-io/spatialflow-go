# CreateInvitationIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "field_worker"]

## Methods

### NewCreateInvitationIn

`func NewCreateInvitationIn(email string, ) *CreateInvitationIn`

NewCreateInvitationIn instantiates a new CreateInvitationIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvitationInWithDefaults

`func NewCreateInvitationInWithDefaults() *CreateInvitationIn`

NewCreateInvitationInWithDefaults instantiates a new CreateInvitationIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateInvitationIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateInvitationIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateInvitationIn) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *CreateInvitationIn) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateInvitationIn) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateInvitationIn) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateInvitationIn) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


