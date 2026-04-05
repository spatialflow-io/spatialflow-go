# ShiftActionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**DeviceId** | **string** |  | 
**ShiftStatus** | **string** |  | 
**ShiftStartedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftPausedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftEndedAt** | Pointer to **NullableTime** |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewShiftActionOut

`func NewShiftActionOut(success bool, deviceId string, shiftStatus string, message string, ) *ShiftActionOut`

NewShiftActionOut instantiates a new ShiftActionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftActionOutWithDefaults

`func NewShiftActionOutWithDefaults() *ShiftActionOut`

NewShiftActionOutWithDefaults instantiates a new ShiftActionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ShiftActionOut) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ShiftActionOut) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ShiftActionOut) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetDeviceId

`func (o *ShiftActionOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ShiftActionOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ShiftActionOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetShiftStatus

`func (o *ShiftActionOut) GetShiftStatus() string`

GetShiftStatus returns the ShiftStatus field if non-nil, zero value otherwise.

### GetShiftStatusOk

`func (o *ShiftActionOut) GetShiftStatusOk() (*string, bool)`

GetShiftStatusOk returns a tuple with the ShiftStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftStatus

`func (o *ShiftActionOut) SetShiftStatus(v string)`

SetShiftStatus sets ShiftStatus field to given value.


### GetShiftStartedAt

`func (o *ShiftActionOut) GetShiftStartedAt() time.Time`

GetShiftStartedAt returns the ShiftStartedAt field if non-nil, zero value otherwise.

### GetShiftStartedAtOk

`func (o *ShiftActionOut) GetShiftStartedAtOk() (*time.Time, bool)`

GetShiftStartedAtOk returns a tuple with the ShiftStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftStartedAt

`func (o *ShiftActionOut) SetShiftStartedAt(v time.Time)`

SetShiftStartedAt sets ShiftStartedAt field to given value.

### HasShiftStartedAt

`func (o *ShiftActionOut) HasShiftStartedAt() bool`

HasShiftStartedAt returns a boolean if a field has been set.

### SetShiftStartedAtNil

`func (o *ShiftActionOut) SetShiftStartedAtNil(b bool)`

 SetShiftStartedAtNil sets the value for ShiftStartedAt to be an explicit nil

### UnsetShiftStartedAt
`func (o *ShiftActionOut) UnsetShiftStartedAt()`

UnsetShiftStartedAt ensures that no value is present for ShiftStartedAt, not even an explicit nil
### GetShiftPausedAt

`func (o *ShiftActionOut) GetShiftPausedAt() time.Time`

GetShiftPausedAt returns the ShiftPausedAt field if non-nil, zero value otherwise.

### GetShiftPausedAtOk

`func (o *ShiftActionOut) GetShiftPausedAtOk() (*time.Time, bool)`

GetShiftPausedAtOk returns a tuple with the ShiftPausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftPausedAt

`func (o *ShiftActionOut) SetShiftPausedAt(v time.Time)`

SetShiftPausedAt sets ShiftPausedAt field to given value.

### HasShiftPausedAt

`func (o *ShiftActionOut) HasShiftPausedAt() bool`

HasShiftPausedAt returns a boolean if a field has been set.

### SetShiftPausedAtNil

`func (o *ShiftActionOut) SetShiftPausedAtNil(b bool)`

 SetShiftPausedAtNil sets the value for ShiftPausedAt to be an explicit nil

### UnsetShiftPausedAt
`func (o *ShiftActionOut) UnsetShiftPausedAt()`

UnsetShiftPausedAt ensures that no value is present for ShiftPausedAt, not even an explicit nil
### GetShiftEndedAt

`func (o *ShiftActionOut) GetShiftEndedAt() time.Time`

GetShiftEndedAt returns the ShiftEndedAt field if non-nil, zero value otherwise.

### GetShiftEndedAtOk

`func (o *ShiftActionOut) GetShiftEndedAtOk() (*time.Time, bool)`

GetShiftEndedAtOk returns a tuple with the ShiftEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftEndedAt

`func (o *ShiftActionOut) SetShiftEndedAt(v time.Time)`

SetShiftEndedAt sets ShiftEndedAt field to given value.

### HasShiftEndedAt

`func (o *ShiftActionOut) HasShiftEndedAt() bool`

HasShiftEndedAt returns a boolean if a field has been set.

### SetShiftEndedAtNil

`func (o *ShiftActionOut) SetShiftEndedAtNil(b bool)`

 SetShiftEndedAtNil sets the value for ShiftEndedAt to be an explicit nil

### UnsetShiftEndedAt
`func (o *ShiftActionOut) UnsetShiftEndedAt()`

UnsetShiftEndedAt ensures that no value is present for ShiftEndedAt, not even an explicit nil
### GetMessage

`func (o *ShiftActionOut) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ShiftActionOut) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ShiftActionOut) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


