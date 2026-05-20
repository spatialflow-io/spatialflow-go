# StatusOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Service** | **string** |  | 
**AdminApprovalRequired** | Pointer to **NullableBool** |  | [optional] 
**ApprovedUsers** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewStatusOut

`func NewStatusOut(status string, service string, ) *StatusOut`

NewStatusOut instantiates a new StatusOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusOutWithDefaults

`func NewStatusOutWithDefaults() *StatusOut`

NewStatusOutWithDefaults instantiates a new StatusOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StatusOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetService

`func (o *StatusOut) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *StatusOut) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *StatusOut) SetService(v string)`

SetService sets Service field to given value.


### GetAdminApprovalRequired

`func (o *StatusOut) GetAdminApprovalRequired() bool`

GetAdminApprovalRequired returns the AdminApprovalRequired field if non-nil, zero value otherwise.

### GetAdminApprovalRequiredOk

`func (o *StatusOut) GetAdminApprovalRequiredOk() (*bool, bool)`

GetAdminApprovalRequiredOk returns a tuple with the AdminApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApprovalRequired

`func (o *StatusOut) SetAdminApprovalRequired(v bool)`

SetAdminApprovalRequired sets AdminApprovalRequired field to given value.

### HasAdminApprovalRequired

`func (o *StatusOut) HasAdminApprovalRequired() bool`

HasAdminApprovalRequired returns a boolean if a field has been set.

### SetAdminApprovalRequiredNil

`func (o *StatusOut) SetAdminApprovalRequiredNil(b bool)`

 SetAdminApprovalRequiredNil sets the value for AdminApprovalRequired to be an explicit nil

### UnsetAdminApprovalRequired
`func (o *StatusOut) UnsetAdminApprovalRequired()`

UnsetAdminApprovalRequired ensures that no value is present for AdminApprovalRequired, not even an explicit nil
### GetApprovedUsers

`func (o *StatusOut) GetApprovedUsers() int32`

GetApprovedUsers returns the ApprovedUsers field if non-nil, zero value otherwise.

### GetApprovedUsersOk

`func (o *StatusOut) GetApprovedUsersOk() (*int32, bool)`

GetApprovedUsersOk returns a tuple with the ApprovedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedUsers

`func (o *StatusOut) SetApprovedUsers(v int32)`

SetApprovedUsers sets ApprovedUsers field to given value.

### HasApprovedUsers

`func (o *StatusOut) HasApprovedUsers() bool`

HasApprovedUsers returns a boolean if a field has been set.

### SetApprovedUsersNil

`func (o *StatusOut) SetApprovedUsersNil(b bool)`

 SetApprovedUsersNil sets the value for ApprovedUsers to be an explicit nil

### UnsetApprovedUsers
`func (o *StatusOut) UnsetApprovedUsers()`

UnsetApprovedUsers ensures that no value is present for ApprovedUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


