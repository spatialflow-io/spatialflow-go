# GeofenceTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeofenceId** | **string** |  | 
**GeofenceName** | **string** |  | 
**IsInside** | **bool** |  | 
**DistanceMeters** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewGeofenceTestResult

`func NewGeofenceTestResult(geofenceId string, geofenceName string, isInside bool, ) *GeofenceTestResult`

NewGeofenceTestResult instantiates a new GeofenceTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeofenceTestResultWithDefaults

`func NewGeofenceTestResultWithDefaults() *GeofenceTestResult`

NewGeofenceTestResultWithDefaults instantiates a new GeofenceTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeofenceId

`func (o *GeofenceTestResult) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *GeofenceTestResult) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *GeofenceTestResult) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.


### GetGeofenceName

`func (o *GeofenceTestResult) GetGeofenceName() string`

GetGeofenceName returns the GeofenceName field if non-nil, zero value otherwise.

### GetGeofenceNameOk

`func (o *GeofenceTestResult) GetGeofenceNameOk() (*string, bool)`

GetGeofenceNameOk returns a tuple with the GeofenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceName

`func (o *GeofenceTestResult) SetGeofenceName(v string)`

SetGeofenceName sets GeofenceName field to given value.


### GetIsInside

`func (o *GeofenceTestResult) GetIsInside() bool`

GetIsInside returns the IsInside field if non-nil, zero value otherwise.

### GetIsInsideOk

`func (o *GeofenceTestResult) GetIsInsideOk() (*bool, bool)`

GetIsInsideOk returns a tuple with the IsInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInside

`func (o *GeofenceTestResult) SetIsInside(v bool)`

SetIsInside sets IsInside field to given value.


### GetDistanceMeters

`func (o *GeofenceTestResult) GetDistanceMeters() float32`

GetDistanceMeters returns the DistanceMeters field if non-nil, zero value otherwise.

### GetDistanceMetersOk

`func (o *GeofenceTestResult) GetDistanceMetersOk() (*float32, bool)`

GetDistanceMetersOk returns a tuple with the DistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMeters

`func (o *GeofenceTestResult) SetDistanceMeters(v float32)`

SetDistanceMeters sets DistanceMeters field to given value.

### HasDistanceMeters

`func (o *GeofenceTestResult) HasDistanceMeters() bool`

HasDistanceMeters returns a boolean if a field has been set.

### SetDistanceMetersNil

`func (o *GeofenceTestResult) SetDistanceMetersNil(b bool)`

 SetDistanceMetersNil sets the value for DistanceMeters to be an explicit nil

### UnsetDistanceMeters
`func (o *GeofenceTestResult) UnsetDistanceMeters()`

UnsetDistanceMeters ensures that no value is present for DistanceMeters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


