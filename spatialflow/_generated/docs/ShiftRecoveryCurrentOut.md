# ShiftRecoveryCurrentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | **bool** |  | 
**SessionId** | Pointer to **NullableString** |  | [optional] 
**ShiftStatus** | **string** |  | 
**ShiftStartedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftPausedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftResumedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftEndedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewShiftRecoveryCurrentOut

`func NewShiftRecoveryCurrentOut(isActive bool, shiftStatus string, ) *ShiftRecoveryCurrentOut`

NewShiftRecoveryCurrentOut instantiates a new ShiftRecoveryCurrentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftRecoveryCurrentOutWithDefaults

`func NewShiftRecoveryCurrentOutWithDefaults() *ShiftRecoveryCurrentOut`

NewShiftRecoveryCurrentOutWithDefaults instantiates a new ShiftRecoveryCurrentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *ShiftRecoveryCurrentOut) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ShiftRecoveryCurrentOut) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ShiftRecoveryCurrentOut) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetSessionId

`func (o *ShiftRecoveryCurrentOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ShiftRecoveryCurrentOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ShiftRecoveryCurrentOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ShiftRecoveryCurrentOut) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *ShiftRecoveryCurrentOut) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *ShiftRecoveryCurrentOut) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetShiftStatus

`func (o *ShiftRecoveryCurrentOut) GetShiftStatus() string`

GetShiftStatus returns the ShiftStatus field if non-nil, zero value otherwise.

### GetShiftStatusOk

`func (o *ShiftRecoveryCurrentOut) GetShiftStatusOk() (*string, bool)`

GetShiftStatusOk returns a tuple with the ShiftStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftStatus

`func (o *ShiftRecoveryCurrentOut) SetShiftStatus(v string)`

SetShiftStatus sets ShiftStatus field to given value.


### GetShiftStartedAt

`func (o *ShiftRecoveryCurrentOut) GetShiftStartedAt() time.Time`

GetShiftStartedAt returns the ShiftStartedAt field if non-nil, zero value otherwise.

### GetShiftStartedAtOk

`func (o *ShiftRecoveryCurrentOut) GetShiftStartedAtOk() (*time.Time, bool)`

GetShiftStartedAtOk returns a tuple with the ShiftStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftStartedAt

`func (o *ShiftRecoveryCurrentOut) SetShiftStartedAt(v time.Time)`

SetShiftStartedAt sets ShiftStartedAt field to given value.

### HasShiftStartedAt

`func (o *ShiftRecoveryCurrentOut) HasShiftStartedAt() bool`

HasShiftStartedAt returns a boolean if a field has been set.

### SetShiftStartedAtNil

`func (o *ShiftRecoveryCurrentOut) SetShiftStartedAtNil(b bool)`

 SetShiftStartedAtNil sets the value for ShiftStartedAt to be an explicit nil

### UnsetShiftStartedAt
`func (o *ShiftRecoveryCurrentOut) UnsetShiftStartedAt()`

UnsetShiftStartedAt ensures that no value is present for ShiftStartedAt, not even an explicit nil
### GetShiftPausedAt

`func (o *ShiftRecoveryCurrentOut) GetShiftPausedAt() time.Time`

GetShiftPausedAt returns the ShiftPausedAt field if non-nil, zero value otherwise.

### GetShiftPausedAtOk

`func (o *ShiftRecoveryCurrentOut) GetShiftPausedAtOk() (*time.Time, bool)`

GetShiftPausedAtOk returns a tuple with the ShiftPausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftPausedAt

`func (o *ShiftRecoveryCurrentOut) SetShiftPausedAt(v time.Time)`

SetShiftPausedAt sets ShiftPausedAt field to given value.

### HasShiftPausedAt

`func (o *ShiftRecoveryCurrentOut) HasShiftPausedAt() bool`

HasShiftPausedAt returns a boolean if a field has been set.

### SetShiftPausedAtNil

`func (o *ShiftRecoveryCurrentOut) SetShiftPausedAtNil(b bool)`

 SetShiftPausedAtNil sets the value for ShiftPausedAt to be an explicit nil

### UnsetShiftPausedAt
`func (o *ShiftRecoveryCurrentOut) UnsetShiftPausedAt()`

UnsetShiftPausedAt ensures that no value is present for ShiftPausedAt, not even an explicit nil
### GetShiftResumedAt

`func (o *ShiftRecoveryCurrentOut) GetShiftResumedAt() time.Time`

GetShiftResumedAt returns the ShiftResumedAt field if non-nil, zero value otherwise.

### GetShiftResumedAtOk

`func (o *ShiftRecoveryCurrentOut) GetShiftResumedAtOk() (*time.Time, bool)`

GetShiftResumedAtOk returns a tuple with the ShiftResumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftResumedAt

`func (o *ShiftRecoveryCurrentOut) SetShiftResumedAt(v time.Time)`

SetShiftResumedAt sets ShiftResumedAt field to given value.

### HasShiftResumedAt

`func (o *ShiftRecoveryCurrentOut) HasShiftResumedAt() bool`

HasShiftResumedAt returns a boolean if a field has been set.

### SetShiftResumedAtNil

`func (o *ShiftRecoveryCurrentOut) SetShiftResumedAtNil(b bool)`

 SetShiftResumedAtNil sets the value for ShiftResumedAt to be an explicit nil

### UnsetShiftResumedAt
`func (o *ShiftRecoveryCurrentOut) UnsetShiftResumedAt()`

UnsetShiftResumedAt ensures that no value is present for ShiftResumedAt, not even an explicit nil
### GetShiftEndedAt

`func (o *ShiftRecoveryCurrentOut) GetShiftEndedAt() time.Time`

GetShiftEndedAt returns the ShiftEndedAt field if non-nil, zero value otherwise.

### GetShiftEndedAtOk

`func (o *ShiftRecoveryCurrentOut) GetShiftEndedAtOk() (*time.Time, bool)`

GetShiftEndedAtOk returns a tuple with the ShiftEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftEndedAt

`func (o *ShiftRecoveryCurrentOut) SetShiftEndedAt(v time.Time)`

SetShiftEndedAt sets ShiftEndedAt field to given value.

### HasShiftEndedAt

`func (o *ShiftRecoveryCurrentOut) HasShiftEndedAt() bool`

HasShiftEndedAt returns a boolean if a field has been set.

### SetShiftEndedAtNil

`func (o *ShiftRecoveryCurrentOut) SetShiftEndedAtNil(b bool)`

 SetShiftEndedAtNil sets the value for ShiftEndedAt to be an explicit nil

### UnsetShiftEndedAt
`func (o *ShiftRecoveryCurrentOut) UnsetShiftEndedAt()`

UnsetShiftEndedAt ensures that no value is present for ShiftEndedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


