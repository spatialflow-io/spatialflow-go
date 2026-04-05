# LocationPointIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**Lat** | **float32** |  | 
**Lon** | **float32** |  | 
**Ts** | **time.Time** |  | 
**Accuracy** | Pointer to **NullableFloat32** |  | [optional] 
**Speed** | Pointer to **NullableFloat32** |  | [optional] 
**Heading** | Pointer to **NullableFloat32** |  | [optional] 
**Altitude** | Pointer to **NullableFloat32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLocationPointIn

`func NewLocationPointIn(deviceId string, lat float32, lon float32, ts time.Time, ) *LocationPointIn`

NewLocationPointIn instantiates a new LocationPointIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationPointInWithDefaults

`func NewLocationPointInWithDefaults() *LocationPointIn`

NewLocationPointInWithDefaults instantiates a new LocationPointIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *LocationPointIn) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *LocationPointIn) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *LocationPointIn) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetLat

`func (o *LocationPointIn) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *LocationPointIn) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *LocationPointIn) SetLat(v float32)`

SetLat sets Lat field to given value.


### GetLon

`func (o *LocationPointIn) GetLon() float32`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *LocationPointIn) GetLonOk() (*float32, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *LocationPointIn) SetLon(v float32)`

SetLon sets Lon field to given value.


### GetTs

`func (o *LocationPointIn) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *LocationPointIn) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *LocationPointIn) SetTs(v time.Time)`

SetTs sets Ts field to given value.


### GetAccuracy

`func (o *LocationPointIn) GetAccuracy() float32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *LocationPointIn) GetAccuracyOk() (*float32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *LocationPointIn) SetAccuracy(v float32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *LocationPointIn) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### SetAccuracyNil

`func (o *LocationPointIn) SetAccuracyNil(b bool)`

 SetAccuracyNil sets the value for Accuracy to be an explicit nil

### UnsetAccuracy
`func (o *LocationPointIn) UnsetAccuracy()`

UnsetAccuracy ensures that no value is present for Accuracy, not even an explicit nil
### GetSpeed

`func (o *LocationPointIn) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *LocationPointIn) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *LocationPointIn) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *LocationPointIn) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *LocationPointIn) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *LocationPointIn) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetHeading

`func (o *LocationPointIn) GetHeading() float32`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *LocationPointIn) GetHeadingOk() (*float32, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *LocationPointIn) SetHeading(v float32)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *LocationPointIn) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### SetHeadingNil

`func (o *LocationPointIn) SetHeadingNil(b bool)`

 SetHeadingNil sets the value for Heading to be an explicit nil

### UnsetHeading
`func (o *LocationPointIn) UnsetHeading()`

UnsetHeading ensures that no value is present for Heading, not even an explicit nil
### GetAltitude

`func (o *LocationPointIn) GetAltitude() float32`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *LocationPointIn) GetAltitudeOk() (*float32, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *LocationPointIn) SetAltitude(v float32)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *LocationPointIn) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### SetAltitudeNil

`func (o *LocationPointIn) SetAltitudeNil(b bool)`

 SetAltitudeNil sets the value for Altitude to be an explicit nil

### UnsetAltitude
`func (o *LocationPointIn) UnsetAltitude()`

UnsetAltitude ensures that no value is present for Altitude, not even an explicit nil
### GetMetadata

`func (o *LocationPointIn) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LocationPointIn) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LocationPointIn) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LocationPointIn) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *LocationPointIn) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *LocationPointIn) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


