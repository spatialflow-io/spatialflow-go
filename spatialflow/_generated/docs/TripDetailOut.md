# TripDetailOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**DeviceName** | **string** |  | 
**Status** | **string** |  | 
**Name** | **string** |  | 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**DurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**DistanceMeters** | Pointer to **NullableFloat32** |  | [optional] 
**SessionId** | Pointer to **NullableString** |  | [optional] 
**HasPlannedRoute** | Pointer to **bool** |  | [optional] [default to false]
**HasTrackGeometry** | Pointer to **bool** |  | [optional] [default to false]
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**PlannedRouteGeojson** | Pointer to **map[string]interface{}** |  | [optional] 
**TrackGeometryGeojson** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTripDetailOut

`func NewTripDetailOut(id string, deviceId string, deviceName string, status string, name string, createdAt time.Time, updatedAt time.Time, ) *TripDetailOut`

NewTripDetailOut instantiates a new TripDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTripDetailOutWithDefaults

`func NewTripDetailOutWithDefaults() *TripDetailOut`

NewTripDetailOutWithDefaults instantiates a new TripDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TripDetailOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TripDetailOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TripDetailOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *TripDetailOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *TripDetailOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *TripDetailOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceName

`func (o *TripDetailOut) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *TripDetailOut) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *TripDetailOut) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetStatus

`func (o *TripDetailOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TripDetailOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TripDetailOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetName

`func (o *TripDetailOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TripDetailOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TripDetailOut) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *TripDetailOut) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TripDetailOut) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TripDetailOut) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TripDetailOut) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TripDetailOut) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TripDetailOut) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *TripDetailOut) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TripDetailOut) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TripDetailOut) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TripDetailOut) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TripDetailOut) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TripDetailOut) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDurationSeconds

`func (o *TripDetailOut) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *TripDetailOut) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *TripDetailOut) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *TripDetailOut) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *TripDetailOut) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *TripDetailOut) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetDistanceMeters

`func (o *TripDetailOut) GetDistanceMeters() float32`

GetDistanceMeters returns the DistanceMeters field if non-nil, zero value otherwise.

### GetDistanceMetersOk

`func (o *TripDetailOut) GetDistanceMetersOk() (*float32, bool)`

GetDistanceMetersOk returns a tuple with the DistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMeters

`func (o *TripDetailOut) SetDistanceMeters(v float32)`

SetDistanceMeters sets DistanceMeters field to given value.

### HasDistanceMeters

`func (o *TripDetailOut) HasDistanceMeters() bool`

HasDistanceMeters returns a boolean if a field has been set.

### SetDistanceMetersNil

`func (o *TripDetailOut) SetDistanceMetersNil(b bool)`

 SetDistanceMetersNil sets the value for DistanceMeters to be an explicit nil

### UnsetDistanceMeters
`func (o *TripDetailOut) UnsetDistanceMeters()`

UnsetDistanceMeters ensures that no value is present for DistanceMeters, not even an explicit nil
### GetSessionId

`func (o *TripDetailOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *TripDetailOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *TripDetailOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *TripDetailOut) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *TripDetailOut) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *TripDetailOut) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetHasPlannedRoute

`func (o *TripDetailOut) GetHasPlannedRoute() bool`

GetHasPlannedRoute returns the HasPlannedRoute field if non-nil, zero value otherwise.

### GetHasPlannedRouteOk

`func (o *TripDetailOut) GetHasPlannedRouteOk() (*bool, bool)`

GetHasPlannedRouteOk returns a tuple with the HasPlannedRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPlannedRoute

`func (o *TripDetailOut) SetHasPlannedRoute(v bool)`

SetHasPlannedRoute sets HasPlannedRoute field to given value.

### HasHasPlannedRoute

`func (o *TripDetailOut) HasHasPlannedRoute() bool`

HasHasPlannedRoute returns a boolean if a field has been set.

### GetHasTrackGeometry

`func (o *TripDetailOut) GetHasTrackGeometry() bool`

GetHasTrackGeometry returns the HasTrackGeometry field if non-nil, zero value otherwise.

### GetHasTrackGeometryOk

`func (o *TripDetailOut) GetHasTrackGeometryOk() (*bool, bool)`

GetHasTrackGeometryOk returns a tuple with the HasTrackGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTrackGeometry

`func (o *TripDetailOut) SetHasTrackGeometry(v bool)`

SetHasTrackGeometry sets HasTrackGeometry field to given value.

### HasHasTrackGeometry

`func (o *TripDetailOut) HasHasTrackGeometry() bool`

HasHasTrackGeometry returns a boolean if a field has been set.

### GetMetadata

`func (o *TripDetailOut) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TripDetailOut) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TripDetailOut) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TripDetailOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TripDetailOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TripDetailOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TripDetailOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TripDetailOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TripDetailOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TripDetailOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPlannedRouteGeojson

`func (o *TripDetailOut) GetPlannedRouteGeojson() map[string]interface{}`

GetPlannedRouteGeojson returns the PlannedRouteGeojson field if non-nil, zero value otherwise.

### GetPlannedRouteGeojsonOk

`func (o *TripDetailOut) GetPlannedRouteGeojsonOk() (*map[string]interface{}, bool)`

GetPlannedRouteGeojsonOk returns a tuple with the PlannedRouteGeojson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRouteGeojson

`func (o *TripDetailOut) SetPlannedRouteGeojson(v map[string]interface{})`

SetPlannedRouteGeojson sets PlannedRouteGeojson field to given value.

### HasPlannedRouteGeojson

`func (o *TripDetailOut) HasPlannedRouteGeojson() bool`

HasPlannedRouteGeojson returns a boolean if a field has been set.

### SetPlannedRouteGeojsonNil

`func (o *TripDetailOut) SetPlannedRouteGeojsonNil(b bool)`

 SetPlannedRouteGeojsonNil sets the value for PlannedRouteGeojson to be an explicit nil

### UnsetPlannedRouteGeojson
`func (o *TripDetailOut) UnsetPlannedRouteGeojson()`

UnsetPlannedRouteGeojson ensures that no value is present for PlannedRouteGeojson, not even an explicit nil
### GetTrackGeometryGeojson

`func (o *TripDetailOut) GetTrackGeometryGeojson() map[string]interface{}`

GetTrackGeometryGeojson returns the TrackGeometryGeojson field if non-nil, zero value otherwise.

### GetTrackGeometryGeojsonOk

`func (o *TripDetailOut) GetTrackGeometryGeojsonOk() (*map[string]interface{}, bool)`

GetTrackGeometryGeojsonOk returns a tuple with the TrackGeometryGeojson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackGeometryGeojson

`func (o *TripDetailOut) SetTrackGeometryGeojson(v map[string]interface{})`

SetTrackGeometryGeojson sets TrackGeometryGeojson field to given value.

### HasTrackGeometryGeojson

`func (o *TripDetailOut) HasTrackGeometryGeojson() bool`

HasTrackGeometryGeojson returns a boolean if a field has been set.

### SetTrackGeometryGeojsonNil

`func (o *TripDetailOut) SetTrackGeometryGeojsonNil(b bool)`

 SetTrackGeometryGeojsonNil sets the value for TrackGeometryGeojson to be an explicit nil

### UnsetTrackGeometryGeojson
`func (o *TripDetailOut) UnsetTrackGeometryGeojson()`

UnsetTrackGeometryGeojson ensures that no value is present for TrackGeometryGeojson, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


