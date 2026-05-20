# GeofenceEventOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EventType** | **string** |  | 
**Device** | [**EventDeviceOut**](EventDeviceOut.md) |  | 
**Geofence** | [**EventGeofenceOut**](EventGeofenceOut.md) |  | 
**Timestamp** | **string** |  | 
**Location** | [**EventLocationOut**](EventLocationOut.md) |  | 
**WorkflowsTriggered** | Pointer to **[]string** |  | [optional] [default to []]
**WebhooksTriggered** | Pointer to **[]string** |  | [optional] [default to []]
**CreatedAt** | **string** |  | 

## Methods

### NewGeofenceEventOut

`func NewGeofenceEventOut(id string, eventType string, device EventDeviceOut, geofence EventGeofenceOut, timestamp string, location EventLocationOut, createdAt string, ) *GeofenceEventOut`

NewGeofenceEventOut instantiates a new GeofenceEventOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeofenceEventOutWithDefaults

`func NewGeofenceEventOutWithDefaults() *GeofenceEventOut`

NewGeofenceEventOutWithDefaults instantiates a new GeofenceEventOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GeofenceEventOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeofenceEventOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeofenceEventOut) SetId(v string)`

SetId sets Id field to given value.


### GetEventType

`func (o *GeofenceEventOut) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GeofenceEventOut) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GeofenceEventOut) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetDevice

`func (o *GeofenceEventOut) GetDevice() EventDeviceOut`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GeofenceEventOut) GetDeviceOk() (*EventDeviceOut, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GeofenceEventOut) SetDevice(v EventDeviceOut)`

SetDevice sets Device field to given value.


### GetGeofence

`func (o *GeofenceEventOut) GetGeofence() EventGeofenceOut`

GetGeofence returns the Geofence field if non-nil, zero value otherwise.

### GetGeofenceOk

`func (o *GeofenceEventOut) GetGeofenceOk() (*EventGeofenceOut, bool)`

GetGeofenceOk returns a tuple with the Geofence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofence

`func (o *GeofenceEventOut) SetGeofence(v EventGeofenceOut)`

SetGeofence sets Geofence field to given value.


### GetTimestamp

`func (o *GeofenceEventOut) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GeofenceEventOut) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GeofenceEventOut) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetLocation

`func (o *GeofenceEventOut) GetLocation() EventLocationOut`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GeofenceEventOut) GetLocationOk() (*EventLocationOut, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GeofenceEventOut) SetLocation(v EventLocationOut)`

SetLocation sets Location field to given value.


### GetWorkflowsTriggered

`func (o *GeofenceEventOut) GetWorkflowsTriggered() []string`

GetWorkflowsTriggered returns the WorkflowsTriggered field if non-nil, zero value otherwise.

### GetWorkflowsTriggeredOk

`func (o *GeofenceEventOut) GetWorkflowsTriggeredOk() (*[]string, bool)`

GetWorkflowsTriggeredOk returns a tuple with the WorkflowsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowsTriggered

`func (o *GeofenceEventOut) SetWorkflowsTriggered(v []string)`

SetWorkflowsTriggered sets WorkflowsTriggered field to given value.

### HasWorkflowsTriggered

`func (o *GeofenceEventOut) HasWorkflowsTriggered() bool`

HasWorkflowsTriggered returns a boolean if a field has been set.

### GetWebhooksTriggered

`func (o *GeofenceEventOut) GetWebhooksTriggered() []string`

GetWebhooksTriggered returns the WebhooksTriggered field if non-nil, zero value otherwise.

### GetWebhooksTriggeredOk

`func (o *GeofenceEventOut) GetWebhooksTriggeredOk() (*[]string, bool)`

GetWebhooksTriggeredOk returns a tuple with the WebhooksTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksTriggered

`func (o *GeofenceEventOut) SetWebhooksTriggered(v []string)`

SetWebhooksTriggered sets WebhooksTriggered field to given value.

### HasWebhooksTriggered

`func (o *GeofenceEventOut) HasWebhooksTriggered() bool`

HasWebhooksTriggered returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GeofenceEventOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GeofenceEventOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GeofenceEventOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


