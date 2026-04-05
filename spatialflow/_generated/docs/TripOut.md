# TripOut

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

## Methods

### NewTripOut

`func NewTripOut(id string, deviceId string, deviceName string, status string, name string, createdAt time.Time, updatedAt time.Time, ) *TripOut`

NewTripOut instantiates a new TripOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTripOutWithDefaults

`func NewTripOutWithDefaults() *TripOut`

NewTripOutWithDefaults instantiates a new TripOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TripOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TripOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TripOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *TripOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *TripOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *TripOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceName

`func (o *TripOut) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *TripOut) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *TripOut) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetStatus

`func (o *TripOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TripOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TripOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetName

`func (o *TripOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TripOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TripOut) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *TripOut) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TripOut) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TripOut) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TripOut) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TripOut) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TripOut) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *TripOut) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TripOut) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TripOut) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TripOut) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TripOut) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TripOut) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDurationSeconds

`func (o *TripOut) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *TripOut) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *TripOut) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *TripOut) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *TripOut) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *TripOut) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetDistanceMeters

`func (o *TripOut) GetDistanceMeters() float32`

GetDistanceMeters returns the DistanceMeters field if non-nil, zero value otherwise.

### GetDistanceMetersOk

`func (o *TripOut) GetDistanceMetersOk() (*float32, bool)`

GetDistanceMetersOk returns a tuple with the DistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMeters

`func (o *TripOut) SetDistanceMeters(v float32)`

SetDistanceMeters sets DistanceMeters field to given value.

### HasDistanceMeters

`func (o *TripOut) HasDistanceMeters() bool`

HasDistanceMeters returns a boolean if a field has been set.

### SetDistanceMetersNil

`func (o *TripOut) SetDistanceMetersNil(b bool)`

 SetDistanceMetersNil sets the value for DistanceMeters to be an explicit nil

### UnsetDistanceMeters
`func (o *TripOut) UnsetDistanceMeters()`

UnsetDistanceMeters ensures that no value is present for DistanceMeters, not even an explicit nil
### GetSessionId

`func (o *TripOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *TripOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *TripOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *TripOut) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *TripOut) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *TripOut) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetHasPlannedRoute

`func (o *TripOut) GetHasPlannedRoute() bool`

GetHasPlannedRoute returns the HasPlannedRoute field if non-nil, zero value otherwise.

### GetHasPlannedRouteOk

`func (o *TripOut) GetHasPlannedRouteOk() (*bool, bool)`

GetHasPlannedRouteOk returns a tuple with the HasPlannedRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPlannedRoute

`func (o *TripOut) SetHasPlannedRoute(v bool)`

SetHasPlannedRoute sets HasPlannedRoute field to given value.

### HasHasPlannedRoute

`func (o *TripOut) HasHasPlannedRoute() bool`

HasHasPlannedRoute returns a boolean if a field has been set.

### GetHasTrackGeometry

`func (o *TripOut) GetHasTrackGeometry() bool`

GetHasTrackGeometry returns the HasTrackGeometry field if non-nil, zero value otherwise.

### GetHasTrackGeometryOk

`func (o *TripOut) GetHasTrackGeometryOk() (*bool, bool)`

GetHasTrackGeometryOk returns a tuple with the HasTrackGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTrackGeometry

`func (o *TripOut) SetHasTrackGeometry(v bool)`

SetHasTrackGeometry sets HasTrackGeometry field to given value.

### HasHasTrackGeometry

`func (o *TripOut) HasHasTrackGeometry() bool`

HasHasTrackGeometry returns a boolean if a field has been set.

### GetMetadata

`func (o *TripOut) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TripOut) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TripOut) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TripOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TripOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TripOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TripOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TripOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TripOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TripOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


