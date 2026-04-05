# SimulationRouteOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SimulationId** | **string** |  | 
**GpxRouteId** | **string** |  | 
**GpxRouteName** | **string** |  | 
**SimulatedDeviceName** | **string** |  | 
**SimulatedDeviceType** | **string** |  | 
**Status** | **string** |  | 
**CurrentPointIndex** | **int32** |  | 
**TotalPoints** | **int32** |  | 
**ProgressPercent** | **float32** |  | 
**EventsTriggered** | **int32** |  | 
**CreatedAt** | **string** |  | 
**StartedAt** | **NullableString** |  | 
**PausedAt** | **NullableString** |  | 
**CompletedAt** | **NullableString** |  | 
**ErrorMessage** | **string** |  | 

## Methods

### NewSimulationRouteOut

`func NewSimulationRouteOut(id string, simulationId string, gpxRouteId string, gpxRouteName string, simulatedDeviceName string, simulatedDeviceType string, status string, currentPointIndex int32, totalPoints int32, progressPercent float32, eventsTriggered int32, createdAt string, startedAt NullableString, pausedAt NullableString, completedAt NullableString, errorMessage string, ) *SimulationRouteOut`

NewSimulationRouteOut instantiates a new SimulationRouteOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationRouteOutWithDefaults

`func NewSimulationRouteOutWithDefaults() *SimulationRouteOut`

NewSimulationRouteOutWithDefaults instantiates a new SimulationRouteOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimulationRouteOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimulationRouteOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimulationRouteOut) SetId(v string)`

SetId sets Id field to given value.


### GetSimulationId

`func (o *SimulationRouteOut) GetSimulationId() string`

GetSimulationId returns the SimulationId field if non-nil, zero value otherwise.

### GetSimulationIdOk

`func (o *SimulationRouteOut) GetSimulationIdOk() (*string, bool)`

GetSimulationIdOk returns a tuple with the SimulationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationId

`func (o *SimulationRouteOut) SetSimulationId(v string)`

SetSimulationId sets SimulationId field to given value.


### GetGpxRouteId

`func (o *SimulationRouteOut) GetGpxRouteId() string`

GetGpxRouteId returns the GpxRouteId field if non-nil, zero value otherwise.

### GetGpxRouteIdOk

`func (o *SimulationRouteOut) GetGpxRouteIdOk() (*string, bool)`

GetGpxRouteIdOk returns a tuple with the GpxRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpxRouteId

`func (o *SimulationRouteOut) SetGpxRouteId(v string)`

SetGpxRouteId sets GpxRouteId field to given value.


### GetGpxRouteName

`func (o *SimulationRouteOut) GetGpxRouteName() string`

GetGpxRouteName returns the GpxRouteName field if non-nil, zero value otherwise.

### GetGpxRouteNameOk

`func (o *SimulationRouteOut) GetGpxRouteNameOk() (*string, bool)`

GetGpxRouteNameOk returns a tuple with the GpxRouteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpxRouteName

`func (o *SimulationRouteOut) SetGpxRouteName(v string)`

SetGpxRouteName sets GpxRouteName field to given value.


### GetSimulatedDeviceName

`func (o *SimulationRouteOut) GetSimulatedDeviceName() string`

GetSimulatedDeviceName returns the SimulatedDeviceName field if non-nil, zero value otherwise.

### GetSimulatedDeviceNameOk

`func (o *SimulationRouteOut) GetSimulatedDeviceNameOk() (*string, bool)`

GetSimulatedDeviceNameOk returns a tuple with the SimulatedDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatedDeviceName

`func (o *SimulationRouteOut) SetSimulatedDeviceName(v string)`

SetSimulatedDeviceName sets SimulatedDeviceName field to given value.


### GetSimulatedDeviceType

`func (o *SimulationRouteOut) GetSimulatedDeviceType() string`

GetSimulatedDeviceType returns the SimulatedDeviceType field if non-nil, zero value otherwise.

### GetSimulatedDeviceTypeOk

`func (o *SimulationRouteOut) GetSimulatedDeviceTypeOk() (*string, bool)`

GetSimulatedDeviceTypeOk returns a tuple with the SimulatedDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatedDeviceType

`func (o *SimulationRouteOut) SetSimulatedDeviceType(v string)`

SetSimulatedDeviceType sets SimulatedDeviceType field to given value.


### GetStatus

`func (o *SimulationRouteOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimulationRouteOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimulationRouteOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrentPointIndex

`func (o *SimulationRouteOut) GetCurrentPointIndex() int32`

GetCurrentPointIndex returns the CurrentPointIndex field if non-nil, zero value otherwise.

### GetCurrentPointIndexOk

`func (o *SimulationRouteOut) GetCurrentPointIndexOk() (*int32, bool)`

GetCurrentPointIndexOk returns a tuple with the CurrentPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPointIndex

`func (o *SimulationRouteOut) SetCurrentPointIndex(v int32)`

SetCurrentPointIndex sets CurrentPointIndex field to given value.


### GetTotalPoints

`func (o *SimulationRouteOut) GetTotalPoints() int32`

GetTotalPoints returns the TotalPoints field if non-nil, zero value otherwise.

### GetTotalPointsOk

`func (o *SimulationRouteOut) GetTotalPointsOk() (*int32, bool)`

GetTotalPointsOk returns a tuple with the TotalPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPoints

`func (o *SimulationRouteOut) SetTotalPoints(v int32)`

SetTotalPoints sets TotalPoints field to given value.


### GetProgressPercent

`func (o *SimulationRouteOut) GetProgressPercent() float32`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *SimulationRouteOut) GetProgressPercentOk() (*float32, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *SimulationRouteOut) SetProgressPercent(v float32)`

SetProgressPercent sets ProgressPercent field to given value.


### GetEventsTriggered

`func (o *SimulationRouteOut) GetEventsTriggered() int32`

GetEventsTriggered returns the EventsTriggered field if non-nil, zero value otherwise.

### GetEventsTriggeredOk

`func (o *SimulationRouteOut) GetEventsTriggeredOk() (*int32, bool)`

GetEventsTriggeredOk returns a tuple with the EventsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsTriggered

`func (o *SimulationRouteOut) SetEventsTriggered(v int32)`

SetEventsTriggered sets EventsTriggered field to given value.


### GetCreatedAt

`func (o *SimulationRouteOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimulationRouteOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimulationRouteOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *SimulationRouteOut) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SimulationRouteOut) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SimulationRouteOut) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *SimulationRouteOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *SimulationRouteOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetPausedAt

`func (o *SimulationRouteOut) GetPausedAt() string`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *SimulationRouteOut) GetPausedAtOk() (*string, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *SimulationRouteOut) SetPausedAt(v string)`

SetPausedAt sets PausedAt field to given value.


### SetPausedAtNil

`func (o *SimulationRouteOut) SetPausedAtNil(b bool)`

 SetPausedAtNil sets the value for PausedAt to be an explicit nil

### UnsetPausedAt
`func (o *SimulationRouteOut) UnsetPausedAt()`

UnsetPausedAt ensures that no value is present for PausedAt, not even an explicit nil
### GetCompletedAt

`func (o *SimulationRouteOut) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *SimulationRouteOut) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *SimulationRouteOut) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *SimulationRouteOut) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *SimulationRouteOut) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetErrorMessage

`func (o *SimulationRouteOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SimulationRouteOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SimulationRouteOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


