# BatchResendIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvitationIds** | Pointer to **[]string** |  | [optional] 
**Before** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBatchResendIn

`func NewBatchResendIn() *BatchResendIn`

NewBatchResendIn instantiates a new BatchResendIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResendInWithDefaults

`func NewBatchResendInWithDefaults() *BatchResendIn`

NewBatchResendInWithDefaults instantiates a new BatchResendIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitationIds

`func (o *BatchResendIn) GetInvitationIds() []string`

GetInvitationIds returns the InvitationIds field if non-nil, zero value otherwise.

### GetInvitationIdsOk

`func (o *BatchResendIn) GetInvitationIdsOk() (*[]string, bool)`

GetInvitationIdsOk returns a tuple with the InvitationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationIds

`func (o *BatchResendIn) SetInvitationIds(v []string)`

SetInvitationIds sets InvitationIds field to given value.

### HasInvitationIds

`func (o *BatchResendIn) HasInvitationIds() bool`

HasInvitationIds returns a boolean if a field has been set.

### SetInvitationIdsNil

`func (o *BatchResendIn) SetInvitationIdsNil(b bool)`

 SetInvitationIdsNil sets the value for InvitationIds to be an explicit nil

### UnsetInvitationIds
`func (o *BatchResendIn) UnsetInvitationIds()`

UnsetInvitationIds ensures that no value is present for InvitationIds, not even an explicit nil
### GetBefore

`func (o *BatchResendIn) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *BatchResendIn) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *BatchResendIn) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *BatchResendIn) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *BatchResendIn) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *BatchResendIn) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


