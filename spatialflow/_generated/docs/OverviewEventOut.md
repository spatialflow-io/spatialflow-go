# OverviewEventOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EventType** | **string** |  | 
**Device** | [**OverviewEventDeviceOut**](OverviewEventDeviceOut.md) |  | 
**Geofence** | [**OverviewEventGeofenceOut**](OverviewEventGeofenceOut.md) |  | 
**Timestamp** | **time.Time** |  | 
**Location** | [**OverviewEventLocationOut**](OverviewEventLocationOut.md) |  | 
**WorkflowsTriggered** | **[]string** |  | 
**WorkflowRuns** | [**[]OverviewEventWorkflowRunOut**](OverviewEventWorkflowRunOut.md) |  | 
**WebhooksTriggered** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewOverviewEventOut

`func NewOverviewEventOut(id string, eventType string, device OverviewEventDeviceOut, geofence OverviewEventGeofenceOut, timestamp time.Time, location OverviewEventLocationOut, workflowsTriggered []string, workflowRuns []OverviewEventWorkflowRunOut, webhooksTriggered []string, createdAt time.Time, ) *OverviewEventOut`

NewOverviewEventOut instantiates a new OverviewEventOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewEventOutWithDefaults

`func NewOverviewEventOutWithDefaults() *OverviewEventOut`

NewOverviewEventOutWithDefaults instantiates a new OverviewEventOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OverviewEventOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OverviewEventOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OverviewEventOut) SetId(v string)`

SetId sets Id field to given value.


### GetEventType

`func (o *OverviewEventOut) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OverviewEventOut) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OverviewEventOut) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetDevice

`func (o *OverviewEventOut) GetDevice() OverviewEventDeviceOut`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *OverviewEventOut) GetDeviceOk() (*OverviewEventDeviceOut, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *OverviewEventOut) SetDevice(v OverviewEventDeviceOut)`

SetDevice sets Device field to given value.


### GetGeofence

`func (o *OverviewEventOut) GetGeofence() OverviewEventGeofenceOut`

GetGeofence returns the Geofence field if non-nil, zero value otherwise.

### GetGeofenceOk

`func (o *OverviewEventOut) GetGeofenceOk() (*OverviewEventGeofenceOut, bool)`

GetGeofenceOk returns a tuple with the Geofence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofence

`func (o *OverviewEventOut) SetGeofence(v OverviewEventGeofenceOut)`

SetGeofence sets Geofence field to given value.


### GetTimestamp

`func (o *OverviewEventOut) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OverviewEventOut) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OverviewEventOut) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetLocation

`func (o *OverviewEventOut) GetLocation() OverviewEventLocationOut`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OverviewEventOut) GetLocationOk() (*OverviewEventLocationOut, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OverviewEventOut) SetLocation(v OverviewEventLocationOut)`

SetLocation sets Location field to given value.


### GetWorkflowsTriggered

`func (o *OverviewEventOut) GetWorkflowsTriggered() []string`

GetWorkflowsTriggered returns the WorkflowsTriggered field if non-nil, zero value otherwise.

### GetWorkflowsTriggeredOk

`func (o *OverviewEventOut) GetWorkflowsTriggeredOk() (*[]string, bool)`

GetWorkflowsTriggeredOk returns a tuple with the WorkflowsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowsTriggered

`func (o *OverviewEventOut) SetWorkflowsTriggered(v []string)`

SetWorkflowsTriggered sets WorkflowsTriggered field to given value.


### GetWorkflowRuns

`func (o *OverviewEventOut) GetWorkflowRuns() []OverviewEventWorkflowRunOut`

GetWorkflowRuns returns the WorkflowRuns field if non-nil, zero value otherwise.

### GetWorkflowRunsOk

`func (o *OverviewEventOut) GetWorkflowRunsOk() (*[]OverviewEventWorkflowRunOut, bool)`

GetWorkflowRunsOk returns a tuple with the WorkflowRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRuns

`func (o *OverviewEventOut) SetWorkflowRuns(v []OverviewEventWorkflowRunOut)`

SetWorkflowRuns sets WorkflowRuns field to given value.


### GetWebhooksTriggered

`func (o *OverviewEventOut) GetWebhooksTriggered() []string`

GetWebhooksTriggered returns the WebhooksTriggered field if non-nil, zero value otherwise.

### GetWebhooksTriggeredOk

`func (o *OverviewEventOut) GetWebhooksTriggeredOk() (*[]string, bool)`

GetWebhooksTriggeredOk returns a tuple with the WebhooksTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksTriggered

`func (o *OverviewEventOut) SetWebhooksTriggered(v []string)`

SetWebhooksTriggered sets WebhooksTriggered field to given value.


### GetCreatedAt

`func (o *OverviewEventOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OverviewEventOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OverviewEventOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


