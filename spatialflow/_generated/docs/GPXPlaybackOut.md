# GPXPlaybackOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RouteId** | **string** |  | 
**RouteName** | **string** |  | 
**DeviceId** | **string** |  | 
**Status** | **string** |  | 
**SpeedMultiplier** | **float32** |  | 
**LoopEnabled** | **bool** |  | 
**CurrentPointIndex** | **int32** |  | 
**ProgressPercent** | **float32** |  | 
**EventsTriggered** | **int32** |  | 
**StartedAt** | **NullableString** |  | 
**PausedAt** | **NullableString** |  | 
**CompletedAt** | **NullableString** |  | 
**ErrorMessage** | **string** |  | 

## Methods

### NewGPXPlaybackOut

`func NewGPXPlaybackOut(id string, routeId string, routeName string, deviceId string, status string, speedMultiplier float32, loopEnabled bool, currentPointIndex int32, progressPercent float32, eventsTriggered int32, startedAt NullableString, pausedAt NullableString, completedAt NullableString, errorMessage string, ) *GPXPlaybackOut`

NewGPXPlaybackOut instantiates a new GPXPlaybackOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPXPlaybackOutWithDefaults

`func NewGPXPlaybackOutWithDefaults() *GPXPlaybackOut`

NewGPXPlaybackOutWithDefaults instantiates a new GPXPlaybackOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GPXPlaybackOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GPXPlaybackOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GPXPlaybackOut) SetId(v string)`

SetId sets Id field to given value.


### GetRouteId

`func (o *GPXPlaybackOut) GetRouteId() string`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *GPXPlaybackOut) GetRouteIdOk() (*string, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *GPXPlaybackOut) SetRouteId(v string)`

SetRouteId sets RouteId field to given value.


### GetRouteName

`func (o *GPXPlaybackOut) GetRouteName() string`

GetRouteName returns the RouteName field if non-nil, zero value otherwise.

### GetRouteNameOk

`func (o *GPXPlaybackOut) GetRouteNameOk() (*string, bool)`

GetRouteNameOk returns a tuple with the RouteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteName

`func (o *GPXPlaybackOut) SetRouteName(v string)`

SetRouteName sets RouteName field to given value.


### GetDeviceId

`func (o *GPXPlaybackOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GPXPlaybackOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GPXPlaybackOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetStatus

`func (o *GPXPlaybackOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GPXPlaybackOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GPXPlaybackOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSpeedMultiplier

`func (o *GPXPlaybackOut) GetSpeedMultiplier() float32`

GetSpeedMultiplier returns the SpeedMultiplier field if non-nil, zero value otherwise.

### GetSpeedMultiplierOk

`func (o *GPXPlaybackOut) GetSpeedMultiplierOk() (*float32, bool)`

GetSpeedMultiplierOk returns a tuple with the SpeedMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMultiplier

`func (o *GPXPlaybackOut) SetSpeedMultiplier(v float32)`

SetSpeedMultiplier sets SpeedMultiplier field to given value.


### GetLoopEnabled

`func (o *GPXPlaybackOut) GetLoopEnabled() bool`

GetLoopEnabled returns the LoopEnabled field if non-nil, zero value otherwise.

### GetLoopEnabledOk

`func (o *GPXPlaybackOut) GetLoopEnabledOk() (*bool, bool)`

GetLoopEnabledOk returns a tuple with the LoopEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopEnabled

`func (o *GPXPlaybackOut) SetLoopEnabled(v bool)`

SetLoopEnabled sets LoopEnabled field to given value.


### GetCurrentPointIndex

`func (o *GPXPlaybackOut) GetCurrentPointIndex() int32`

GetCurrentPointIndex returns the CurrentPointIndex field if non-nil, zero value otherwise.

### GetCurrentPointIndexOk

`func (o *GPXPlaybackOut) GetCurrentPointIndexOk() (*int32, bool)`

GetCurrentPointIndexOk returns a tuple with the CurrentPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPointIndex

`func (o *GPXPlaybackOut) SetCurrentPointIndex(v int32)`

SetCurrentPointIndex sets CurrentPointIndex field to given value.


### GetProgressPercent

`func (o *GPXPlaybackOut) GetProgressPercent() float32`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *GPXPlaybackOut) GetProgressPercentOk() (*float32, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *GPXPlaybackOut) SetProgressPercent(v float32)`

SetProgressPercent sets ProgressPercent field to given value.


### GetEventsTriggered

`func (o *GPXPlaybackOut) GetEventsTriggered() int32`

GetEventsTriggered returns the EventsTriggered field if non-nil, zero value otherwise.

### GetEventsTriggeredOk

`func (o *GPXPlaybackOut) GetEventsTriggeredOk() (*int32, bool)`

GetEventsTriggeredOk returns a tuple with the EventsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsTriggered

`func (o *GPXPlaybackOut) SetEventsTriggered(v int32)`

SetEventsTriggered sets EventsTriggered field to given value.


### GetStartedAt

`func (o *GPXPlaybackOut) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GPXPlaybackOut) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GPXPlaybackOut) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *GPXPlaybackOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GPXPlaybackOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetPausedAt

`func (o *GPXPlaybackOut) GetPausedAt() string`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *GPXPlaybackOut) GetPausedAtOk() (*string, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *GPXPlaybackOut) SetPausedAt(v string)`

SetPausedAt sets PausedAt field to given value.


### SetPausedAtNil

`func (o *GPXPlaybackOut) SetPausedAtNil(b bool)`

 SetPausedAtNil sets the value for PausedAt to be an explicit nil

### UnsetPausedAt
`func (o *GPXPlaybackOut) UnsetPausedAt()`

UnsetPausedAt ensures that no value is present for PausedAt, not even an explicit nil
### GetCompletedAt

`func (o *GPXPlaybackOut) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GPXPlaybackOut) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GPXPlaybackOut) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *GPXPlaybackOut) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GPXPlaybackOut) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetErrorMessage

`func (o *GPXPlaybackOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GPXPlaybackOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GPXPlaybackOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


