# DeviceSessionsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | [**[]DeviceSessionOut**](DeviceSessionOut.md) |  | 
**TotalCount** | **int32** |  | 
**OpenSession** | Pointer to [**NullableDeviceSessionOut**](DeviceSessionOut.md) |  | [optional] 

## Methods

### NewDeviceSessionsOut

`func NewDeviceSessionsOut(sessions []DeviceSessionOut, totalCount int32, ) *DeviceSessionsOut`

NewDeviceSessionsOut instantiates a new DeviceSessionsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSessionsOutWithDefaults

`func NewDeviceSessionsOutWithDefaults() *DeviceSessionsOut`

NewDeviceSessionsOutWithDefaults instantiates a new DeviceSessionsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *DeviceSessionsOut) GetSessions() []DeviceSessionOut`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *DeviceSessionsOut) GetSessionsOk() (*[]DeviceSessionOut, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *DeviceSessionsOut) SetSessions(v []DeviceSessionOut)`

SetSessions sets Sessions field to given value.


### GetTotalCount

`func (o *DeviceSessionsOut) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DeviceSessionsOut) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DeviceSessionsOut) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetOpenSession

`func (o *DeviceSessionsOut) GetOpenSession() DeviceSessionOut`

GetOpenSession returns the OpenSession field if non-nil, zero value otherwise.

### GetOpenSessionOk

`func (o *DeviceSessionsOut) GetOpenSessionOk() (*DeviceSessionOut, bool)`

GetOpenSessionOk returns a tuple with the OpenSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSession

`func (o *DeviceSessionsOut) SetOpenSession(v DeviceSessionOut)`

SetOpenSession sets OpenSession field to given value.

### HasOpenSession

`func (o *DeviceSessionsOut) HasOpenSession() bool`

HasOpenSession returns a boolean if a field has been set.

### SetOpenSessionNil

`func (o *DeviceSessionsOut) SetOpenSessionNil(b bool)`

 SetOpenSessionNil sets the value for OpenSession to be an explicit nil

### UnsetOpenSession
`func (o *DeviceSessionsOut) UnsetOpenSession()`

UnsetOpenSession ensures that no value is present for OpenSession, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


