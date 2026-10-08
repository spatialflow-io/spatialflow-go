# RunHistoryEntryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EvaluatedAt** | **time.Time** |  | 
**TriggerType** | **string** |  | 
**DeviceId** | Pointer to **NullableString** |  | [optional] 
**GeofenceId** | Pointer to **NullableString** |  | [optional] 
**SignalEventId** | Pointer to **NullableString** |  | [optional] 
**EventTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Matched** | **bool** |  | 
**SkipReason** | Pointer to **NullableString** |  | [optional] 
**SkipReasonLabel** | Pointer to **NullableString** |  | [optional] 
**Execution** | Pointer to [**NullableExecutionOut**](ExecutionOut.md) |  | [optional] 
**Actions** | Pointer to [**[]ActionDeliveryOut**](ActionDeliveryOut.md) |  | [optional] [default to []]

## Methods

### NewRunHistoryEntryOut

`func NewRunHistoryEntryOut(id string, evaluatedAt time.Time, triggerType string, matched bool, ) *RunHistoryEntryOut`

NewRunHistoryEntryOut instantiates a new RunHistoryEntryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunHistoryEntryOutWithDefaults

`func NewRunHistoryEntryOutWithDefaults() *RunHistoryEntryOut`

NewRunHistoryEntryOutWithDefaults instantiates a new RunHistoryEntryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunHistoryEntryOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunHistoryEntryOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunHistoryEntryOut) SetId(v string)`

SetId sets Id field to given value.


### GetEvaluatedAt

`func (o *RunHistoryEntryOut) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *RunHistoryEntryOut) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *RunHistoryEntryOut) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.


### GetTriggerType

`func (o *RunHistoryEntryOut) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *RunHistoryEntryOut) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *RunHistoryEntryOut) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetDeviceId

`func (o *RunHistoryEntryOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RunHistoryEntryOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RunHistoryEntryOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *RunHistoryEntryOut) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *RunHistoryEntryOut) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *RunHistoryEntryOut) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetGeofenceId

`func (o *RunHistoryEntryOut) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *RunHistoryEntryOut) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *RunHistoryEntryOut) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.

### HasGeofenceId

`func (o *RunHistoryEntryOut) HasGeofenceId() bool`

HasGeofenceId returns a boolean if a field has been set.

### SetGeofenceIdNil

`func (o *RunHistoryEntryOut) SetGeofenceIdNil(b bool)`

 SetGeofenceIdNil sets the value for GeofenceId to be an explicit nil

### UnsetGeofenceId
`func (o *RunHistoryEntryOut) UnsetGeofenceId()`

UnsetGeofenceId ensures that no value is present for GeofenceId, not even an explicit nil
### GetSignalEventId

`func (o *RunHistoryEntryOut) GetSignalEventId() string`

GetSignalEventId returns the SignalEventId field if non-nil, zero value otherwise.

### GetSignalEventIdOk

`func (o *RunHistoryEntryOut) GetSignalEventIdOk() (*string, bool)`

GetSignalEventIdOk returns a tuple with the SignalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalEventId

`func (o *RunHistoryEntryOut) SetSignalEventId(v string)`

SetSignalEventId sets SignalEventId field to given value.

### HasSignalEventId

`func (o *RunHistoryEntryOut) HasSignalEventId() bool`

HasSignalEventId returns a boolean if a field has been set.

### SetSignalEventIdNil

`func (o *RunHistoryEntryOut) SetSignalEventIdNil(b bool)`

 SetSignalEventIdNil sets the value for SignalEventId to be an explicit nil

### UnsetSignalEventId
`func (o *RunHistoryEntryOut) UnsetSignalEventId()`

UnsetSignalEventId ensures that no value is present for SignalEventId, not even an explicit nil
### GetEventTimestamp

`func (o *RunHistoryEntryOut) GetEventTimestamp() time.Time`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *RunHistoryEntryOut) GetEventTimestampOk() (*time.Time, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *RunHistoryEntryOut) SetEventTimestamp(v time.Time)`

SetEventTimestamp sets EventTimestamp field to given value.

### HasEventTimestamp

`func (o *RunHistoryEntryOut) HasEventTimestamp() bool`

HasEventTimestamp returns a boolean if a field has been set.

### SetEventTimestampNil

`func (o *RunHistoryEntryOut) SetEventTimestampNil(b bool)`

 SetEventTimestampNil sets the value for EventTimestamp to be an explicit nil

### UnsetEventTimestamp
`func (o *RunHistoryEntryOut) UnsetEventTimestamp()`

UnsetEventTimestamp ensures that no value is present for EventTimestamp, not even an explicit nil
### GetMatched

`func (o *RunHistoryEntryOut) GetMatched() bool`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *RunHistoryEntryOut) GetMatchedOk() (*bool, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *RunHistoryEntryOut) SetMatched(v bool)`

SetMatched sets Matched field to given value.


### GetSkipReason

`func (o *RunHistoryEntryOut) GetSkipReason() string`

GetSkipReason returns the SkipReason field if non-nil, zero value otherwise.

### GetSkipReasonOk

`func (o *RunHistoryEntryOut) GetSkipReasonOk() (*string, bool)`

GetSkipReasonOk returns a tuple with the SkipReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReason

`func (o *RunHistoryEntryOut) SetSkipReason(v string)`

SetSkipReason sets SkipReason field to given value.

### HasSkipReason

`func (o *RunHistoryEntryOut) HasSkipReason() bool`

HasSkipReason returns a boolean if a field has been set.

### SetSkipReasonNil

`func (o *RunHistoryEntryOut) SetSkipReasonNil(b bool)`

 SetSkipReasonNil sets the value for SkipReason to be an explicit nil

### UnsetSkipReason
`func (o *RunHistoryEntryOut) UnsetSkipReason()`

UnsetSkipReason ensures that no value is present for SkipReason, not even an explicit nil
### GetSkipReasonLabel

`func (o *RunHistoryEntryOut) GetSkipReasonLabel() string`

GetSkipReasonLabel returns the SkipReasonLabel field if non-nil, zero value otherwise.

### GetSkipReasonLabelOk

`func (o *RunHistoryEntryOut) GetSkipReasonLabelOk() (*string, bool)`

GetSkipReasonLabelOk returns a tuple with the SkipReasonLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReasonLabel

`func (o *RunHistoryEntryOut) SetSkipReasonLabel(v string)`

SetSkipReasonLabel sets SkipReasonLabel field to given value.

### HasSkipReasonLabel

`func (o *RunHistoryEntryOut) HasSkipReasonLabel() bool`

HasSkipReasonLabel returns a boolean if a field has been set.

### SetSkipReasonLabelNil

`func (o *RunHistoryEntryOut) SetSkipReasonLabelNil(b bool)`

 SetSkipReasonLabelNil sets the value for SkipReasonLabel to be an explicit nil

### UnsetSkipReasonLabel
`func (o *RunHistoryEntryOut) UnsetSkipReasonLabel()`

UnsetSkipReasonLabel ensures that no value is present for SkipReasonLabel, not even an explicit nil
### GetExecution

`func (o *RunHistoryEntryOut) GetExecution() ExecutionOut`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *RunHistoryEntryOut) GetExecutionOk() (*ExecutionOut, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *RunHistoryEntryOut) SetExecution(v ExecutionOut)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *RunHistoryEntryOut) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### SetExecutionNil

`func (o *RunHistoryEntryOut) SetExecutionNil(b bool)`

 SetExecutionNil sets the value for Execution to be an explicit nil

### UnsetExecution
`func (o *RunHistoryEntryOut) UnsetExecution()`

UnsetExecution ensures that no value is present for Execution, not even an explicit nil
### GetActions

`func (o *RunHistoryEntryOut) GetActions() []ActionDeliveryOut`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RunHistoryEntryOut) GetActionsOk() (*[]ActionDeliveryOut, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RunHistoryEntryOut) SetActions(v []ActionDeliveryOut)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RunHistoryEntryOut) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


