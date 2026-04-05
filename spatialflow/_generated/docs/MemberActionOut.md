# MemberActionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**UserId** | **string** |  | 
**Role** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMemberActionOut

`func NewMemberActionOut(message string, userId string, ) *MemberActionOut`

NewMemberActionOut instantiates a new MemberActionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberActionOutWithDefaults

`func NewMemberActionOutWithDefaults() *MemberActionOut`

NewMemberActionOutWithDefaults instantiates a new MemberActionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MemberActionOut) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MemberActionOut) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MemberActionOut) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUserId

`func (o *MemberActionOut) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberActionOut) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberActionOut) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRole

`func (o *MemberActionOut) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberActionOut) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberActionOut) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberActionOut) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *MemberActionOut) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *MemberActionOut) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


