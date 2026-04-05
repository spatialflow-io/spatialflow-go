# SimulationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Status** | **string** |  | 
**SpeedMultiplier** | **float32** |  | 
**LoopEnabled** | **bool** |  | 
**MuteExternalActions** | **bool** |  | 
**ProgressPercent** | **float32** |  | 
**TotalEventsTriggered** | **int32** |  | 
**RouteCount** | **int32** |  | 
**CreatedAt** | **string** |  | 
**StartedAt** | **NullableString** |  | 
**PausedAt** | **NullableString** |  | 
**CompletedAt** | **NullableString** |  | 
**ErrorMessage** | **string** |  | 

## Methods

### NewSimulationOut

`func NewSimulationOut(id string, name string, description string, status string, speedMultiplier float32, loopEnabled bool, muteExternalActions bool, progressPercent float32, totalEventsTriggered int32, routeCount int32, createdAt string, startedAt NullableString, pausedAt NullableString, completedAt NullableString, errorMessage string, ) *SimulationOut`

NewSimulationOut instantiates a new SimulationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationOutWithDefaults

`func NewSimulationOutWithDefaults() *SimulationOut`

NewSimulationOutWithDefaults instantiates a new SimulationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimulationOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimulationOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimulationOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SimulationOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimulationOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimulationOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SimulationOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimulationOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimulationOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *SimulationOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimulationOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimulationOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSpeedMultiplier

`func (o *SimulationOut) GetSpeedMultiplier() float32`

GetSpeedMultiplier returns the SpeedMultiplier field if non-nil, zero value otherwise.

### GetSpeedMultiplierOk

`func (o *SimulationOut) GetSpeedMultiplierOk() (*float32, bool)`

GetSpeedMultiplierOk returns a tuple with the SpeedMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMultiplier

`func (o *SimulationOut) SetSpeedMultiplier(v float32)`

SetSpeedMultiplier sets SpeedMultiplier field to given value.


### GetLoopEnabled

`func (o *SimulationOut) GetLoopEnabled() bool`

GetLoopEnabled returns the LoopEnabled field if non-nil, zero value otherwise.

### GetLoopEnabledOk

`func (o *SimulationOut) GetLoopEnabledOk() (*bool, bool)`

GetLoopEnabledOk returns a tuple with the LoopEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopEnabled

`func (o *SimulationOut) SetLoopEnabled(v bool)`

SetLoopEnabled sets LoopEnabled field to given value.


### GetMuteExternalActions

`func (o *SimulationOut) GetMuteExternalActions() bool`

GetMuteExternalActions returns the MuteExternalActions field if non-nil, zero value otherwise.

### GetMuteExternalActionsOk

`func (o *SimulationOut) GetMuteExternalActionsOk() (*bool, bool)`

GetMuteExternalActionsOk returns a tuple with the MuteExternalActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteExternalActions

`func (o *SimulationOut) SetMuteExternalActions(v bool)`

SetMuteExternalActions sets MuteExternalActions field to given value.


### GetProgressPercent

`func (o *SimulationOut) GetProgressPercent() float32`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *SimulationOut) GetProgressPercentOk() (*float32, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *SimulationOut) SetProgressPercent(v float32)`

SetProgressPercent sets ProgressPercent field to given value.


### GetTotalEventsTriggered

`func (o *SimulationOut) GetTotalEventsTriggered() int32`

GetTotalEventsTriggered returns the TotalEventsTriggered field if non-nil, zero value otherwise.

### GetTotalEventsTriggeredOk

`func (o *SimulationOut) GetTotalEventsTriggeredOk() (*int32, bool)`

GetTotalEventsTriggeredOk returns a tuple with the TotalEventsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEventsTriggered

`func (o *SimulationOut) SetTotalEventsTriggered(v int32)`

SetTotalEventsTriggered sets TotalEventsTriggered field to given value.


### GetRouteCount

`func (o *SimulationOut) GetRouteCount() int32`

GetRouteCount returns the RouteCount field if non-nil, zero value otherwise.

### GetRouteCountOk

`func (o *SimulationOut) GetRouteCountOk() (*int32, bool)`

GetRouteCountOk returns a tuple with the RouteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteCount

`func (o *SimulationOut) SetRouteCount(v int32)`

SetRouteCount sets RouteCount field to given value.


### GetCreatedAt

`func (o *SimulationOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimulationOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimulationOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *SimulationOut) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SimulationOut) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SimulationOut) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *SimulationOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *SimulationOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetPausedAt

`func (o *SimulationOut) GetPausedAt() string`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *SimulationOut) GetPausedAtOk() (*string, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *SimulationOut) SetPausedAt(v string)`

SetPausedAt sets PausedAt field to given value.


### SetPausedAtNil

`func (o *SimulationOut) SetPausedAtNil(b bool)`

 SetPausedAtNil sets the value for PausedAt to be an explicit nil

### UnsetPausedAt
`func (o *SimulationOut) UnsetPausedAt()`

UnsetPausedAt ensures that no value is present for PausedAt, not even an explicit nil
### GetCompletedAt

`func (o *SimulationOut) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *SimulationOut) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *SimulationOut) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *SimulationOut) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *SimulationOut) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetErrorMessage

`func (o *SimulationOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SimulationOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SimulationOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


