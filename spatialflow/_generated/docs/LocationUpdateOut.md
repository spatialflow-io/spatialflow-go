# LocationUpdateOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**DeviceId** | **string** |  | 
**EventsTriggered** | **int32** |  | 
**Message** | **string** |  | 

## Methods

### NewLocationUpdateOut

`func NewLocationUpdateOut(success bool, deviceId string, eventsTriggered int32, message string, ) *LocationUpdateOut`

NewLocationUpdateOut instantiates a new LocationUpdateOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationUpdateOutWithDefaults

`func NewLocationUpdateOutWithDefaults() *LocationUpdateOut`

NewLocationUpdateOutWithDefaults instantiates a new LocationUpdateOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LocationUpdateOut) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LocationUpdateOut) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LocationUpdateOut) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetDeviceId

`func (o *LocationUpdateOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *LocationUpdateOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *LocationUpdateOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetEventsTriggered

`func (o *LocationUpdateOut) GetEventsTriggered() int32`

GetEventsTriggered returns the EventsTriggered field if non-nil, zero value otherwise.

### GetEventsTriggeredOk

`func (o *LocationUpdateOut) GetEventsTriggeredOk() (*int32, bool)`

GetEventsTriggeredOk returns a tuple with the EventsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsTriggered

`func (o *LocationUpdateOut) SetEventsTriggered(v int32)`

SetEventsTriggered sets EventsTriggered field to given value.


### GetMessage

`func (o *LocationUpdateOut) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LocationUpdateOut) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LocationUpdateOut) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


