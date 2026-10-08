# LocalShiftStopIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShiftStartedAt** | **time.Time** |  | 
**StoppedAt** | **time.Time** |  | 

## Methods

### NewLocalShiftStopIn

`func NewLocalShiftStopIn(shiftStartedAt time.Time, stoppedAt time.Time, ) *LocalShiftStopIn`

NewLocalShiftStopIn instantiates a new LocalShiftStopIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalShiftStopInWithDefaults

`func NewLocalShiftStopInWithDefaults() *LocalShiftStopIn`

NewLocalShiftStopInWithDefaults instantiates a new LocalShiftStopIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShiftStartedAt

`func (o *LocalShiftStopIn) GetShiftStartedAt() time.Time`

GetShiftStartedAt returns the ShiftStartedAt field if non-nil, zero value otherwise.

### GetShiftStartedAtOk

`func (o *LocalShiftStopIn) GetShiftStartedAtOk() (*time.Time, bool)`

GetShiftStartedAtOk returns a tuple with the ShiftStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftStartedAt

`func (o *LocalShiftStopIn) SetShiftStartedAt(v time.Time)`

SetShiftStartedAt sets ShiftStartedAt field to given value.


### GetStoppedAt

`func (o *LocalShiftStopIn) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *LocalShiftStopIn) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *LocalShiftStopIn) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


