# SimulationEventOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EventType** | **string** |  | 
**GeofenceName** | **string** |  | 
**WorkflowName** | **string** |  | 
**ActionType** | **string** |  | 
**Latitude** | **float32** |  | 
**Longitude** | **float32** |  | 
**Details** | **map[string]interface{}** |  | 
**DwellDurationSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**Timestamp** | **string** |  | 

## Methods

### NewSimulationEventOut

`func NewSimulationEventOut(id string, eventType string, geofenceName string, workflowName string, actionType string, latitude float32, longitude float32, details map[string]interface{}, timestamp string, ) *SimulationEventOut`

NewSimulationEventOut instantiates a new SimulationEventOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationEventOutWithDefaults

`func NewSimulationEventOutWithDefaults() *SimulationEventOut`

NewSimulationEventOutWithDefaults instantiates a new SimulationEventOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimulationEventOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimulationEventOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimulationEventOut) SetId(v string)`

SetId sets Id field to given value.


### GetEventType

`func (o *SimulationEventOut) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SimulationEventOut) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SimulationEventOut) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetGeofenceName

`func (o *SimulationEventOut) GetGeofenceName() string`

GetGeofenceName returns the GeofenceName field if non-nil, zero value otherwise.

### GetGeofenceNameOk

`func (o *SimulationEventOut) GetGeofenceNameOk() (*string, bool)`

GetGeofenceNameOk returns a tuple with the GeofenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceName

`func (o *SimulationEventOut) SetGeofenceName(v string)`

SetGeofenceName sets GeofenceName field to given value.


### GetWorkflowName

`func (o *SimulationEventOut) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *SimulationEventOut) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *SimulationEventOut) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetActionType

`func (o *SimulationEventOut) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *SimulationEventOut) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *SimulationEventOut) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetLatitude

`func (o *SimulationEventOut) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SimulationEventOut) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SimulationEventOut) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *SimulationEventOut) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SimulationEventOut) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SimulationEventOut) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetDetails

`func (o *SimulationEventOut) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SimulationEventOut) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SimulationEventOut) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.


### GetDwellDurationSeconds

`func (o *SimulationEventOut) GetDwellDurationSeconds() float32`

GetDwellDurationSeconds returns the DwellDurationSeconds field if non-nil, zero value otherwise.

### GetDwellDurationSecondsOk

`func (o *SimulationEventOut) GetDwellDurationSecondsOk() (*float32, bool)`

GetDwellDurationSecondsOk returns a tuple with the DwellDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwellDurationSeconds

`func (o *SimulationEventOut) SetDwellDurationSeconds(v float32)`

SetDwellDurationSeconds sets DwellDurationSeconds field to given value.

### HasDwellDurationSeconds

`func (o *SimulationEventOut) HasDwellDurationSeconds() bool`

HasDwellDurationSeconds returns a boolean if a field has been set.

### SetDwellDurationSecondsNil

`func (o *SimulationEventOut) SetDwellDurationSecondsNil(b bool)`

 SetDwellDurationSecondsNil sets the value for DwellDurationSeconds to be an explicit nil

### UnsetDwellDurationSeconds
`func (o *SimulationEventOut) UnsetDwellDurationSeconds()`

UnsetDwellDurationSeconds ensures that no value is present for DwellDurationSeconds, not even an explicit nil
### GetTimestamp

`func (o *SimulationEventOut) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SimulationEventOut) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SimulationEventOut) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


