# GPXRouteOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**DeviceId** | **string** |  | 
**DeviceName** | **string** |  | 
**TotalPoints** | **int32** |  | 
**TotalDistanceKm** | **float32** |  | 
**TotalDurationSeconds** | **float32** |  | 
**StartTime** | **NullableString** |  | 
**EndTime** | **NullableString** |  | 
**IsActive** | **bool** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewGPXRouteOut

`func NewGPXRouteOut(id string, name string, description string, deviceId string, deviceName string, totalPoints int32, totalDistanceKm float32, totalDurationSeconds float32, startTime NullableString, endTime NullableString, isActive bool, createdAt string, ) *GPXRouteOut`

NewGPXRouteOut instantiates a new GPXRouteOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPXRouteOutWithDefaults

`func NewGPXRouteOutWithDefaults() *GPXRouteOut`

NewGPXRouteOutWithDefaults instantiates a new GPXRouteOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GPXRouteOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GPXRouteOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GPXRouteOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GPXRouteOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GPXRouteOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GPXRouteOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GPXRouteOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GPXRouteOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GPXRouteOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDeviceId

`func (o *GPXRouteOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GPXRouteOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GPXRouteOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceName

`func (o *GPXRouteOut) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *GPXRouteOut) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *GPXRouteOut) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetTotalPoints

`func (o *GPXRouteOut) GetTotalPoints() int32`

GetTotalPoints returns the TotalPoints field if non-nil, zero value otherwise.

### GetTotalPointsOk

`func (o *GPXRouteOut) GetTotalPointsOk() (*int32, bool)`

GetTotalPointsOk returns a tuple with the TotalPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPoints

`func (o *GPXRouteOut) SetTotalPoints(v int32)`

SetTotalPoints sets TotalPoints field to given value.


### GetTotalDistanceKm

`func (o *GPXRouteOut) GetTotalDistanceKm() float32`

GetTotalDistanceKm returns the TotalDistanceKm field if non-nil, zero value otherwise.

### GetTotalDistanceKmOk

`func (o *GPXRouteOut) GetTotalDistanceKmOk() (*float32, bool)`

GetTotalDistanceKmOk returns a tuple with the TotalDistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDistanceKm

`func (o *GPXRouteOut) SetTotalDistanceKm(v float32)`

SetTotalDistanceKm sets TotalDistanceKm field to given value.


### GetTotalDurationSeconds

`func (o *GPXRouteOut) GetTotalDurationSeconds() float32`

GetTotalDurationSeconds returns the TotalDurationSeconds field if non-nil, zero value otherwise.

### GetTotalDurationSecondsOk

`func (o *GPXRouteOut) GetTotalDurationSecondsOk() (*float32, bool)`

GetTotalDurationSecondsOk returns a tuple with the TotalDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationSeconds

`func (o *GPXRouteOut) SetTotalDurationSeconds(v float32)`

SetTotalDurationSeconds sets TotalDurationSeconds field to given value.


### GetStartTime

`func (o *GPXRouteOut) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GPXRouteOut) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GPXRouteOut) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *GPXRouteOut) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *GPXRouteOut) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *GPXRouteOut) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GPXRouteOut) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GPXRouteOut) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *GPXRouteOut) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *GPXRouteOut) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetIsActive

`func (o *GPXRouteOut) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GPXRouteOut) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GPXRouteOut) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCreatedAt

`func (o *GPXRouteOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GPXRouteOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GPXRouteOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


