# BatchResendItemOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvitationId** | **string** |  | 
**Status** | **string** |  | 
**Invitation** | Pointer to [**NullableInvitationIssuedOut**](InvitationIssuedOut.md) |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBatchResendItemOut

`func NewBatchResendItemOut(invitationId string, status string, ) *BatchResendItemOut`

NewBatchResendItemOut instantiates a new BatchResendItemOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResendItemOutWithDefaults

`func NewBatchResendItemOutWithDefaults() *BatchResendItemOut`

NewBatchResendItemOutWithDefaults instantiates a new BatchResendItemOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitationId

`func (o *BatchResendItemOut) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *BatchResendItemOut) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *BatchResendItemOut) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.


### GetStatus

`func (o *BatchResendItemOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResendItemOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResendItemOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInvitation

`func (o *BatchResendItemOut) GetInvitation() InvitationIssuedOut`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *BatchResendItemOut) GetInvitationOk() (*InvitationIssuedOut, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitation

`func (o *BatchResendItemOut) SetInvitation(v InvitationIssuedOut)`

SetInvitation sets Invitation field to given value.

### HasInvitation

`func (o *BatchResendItemOut) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### SetInvitationNil

`func (o *BatchResendItemOut) SetInvitationNil(b bool)`

 SetInvitationNil sets the value for Invitation to be an explicit nil

### UnsetInvitation
`func (o *BatchResendItemOut) UnsetInvitation()`

UnsetInvitation ensures that no value is present for Invitation, not even an explicit nil
### GetDetail

`func (o *BatchResendItemOut) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BatchResendItemOut) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BatchResendItemOut) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BatchResendItemOut) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *BatchResendItemOut) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *BatchResendItemOut) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetErrorCode

`func (o *BatchResendItemOut) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BatchResendItemOut) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BatchResendItemOut) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BatchResendItemOut) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BatchResendItemOut) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BatchResendItemOut) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


