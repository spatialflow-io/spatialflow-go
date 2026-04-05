# LocationUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float32** |  | 
**Longitude** | **float32** |  | 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Accuracy** | Pointer to **NullableFloat32** |  | [optional] 
**Speed** | Pointer to **NullableFloat32** |  | [optional] 
**Heading** | Pointer to **NullableFloat32** |  | [optional] 
**Altitude** | Pointer to **NullableFloat32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLocationUpdateIn

`func NewLocationUpdateIn(latitude float32, longitude float32, ) *LocationUpdateIn`

NewLocationUpdateIn instantiates a new LocationUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationUpdateInWithDefaults

`func NewLocationUpdateInWithDefaults() *LocationUpdateIn`

NewLocationUpdateInWithDefaults instantiates a new LocationUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *LocationUpdateIn) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationUpdateIn) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationUpdateIn) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *LocationUpdateIn) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationUpdateIn) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationUpdateIn) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetTimestamp

`func (o *LocationUpdateIn) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LocationUpdateIn) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LocationUpdateIn) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LocationUpdateIn) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LocationUpdateIn) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LocationUpdateIn) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetAccuracy

`func (o *LocationUpdateIn) GetAccuracy() float32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *LocationUpdateIn) GetAccuracyOk() (*float32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *LocationUpdateIn) SetAccuracy(v float32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *LocationUpdateIn) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### SetAccuracyNil

`func (o *LocationUpdateIn) SetAccuracyNil(b bool)`

 SetAccuracyNil sets the value for Accuracy to be an explicit nil

### UnsetAccuracy
`func (o *LocationUpdateIn) UnsetAccuracy()`

UnsetAccuracy ensures that no value is present for Accuracy, not even an explicit nil
### GetSpeed

`func (o *LocationUpdateIn) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *LocationUpdateIn) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *LocationUpdateIn) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *LocationUpdateIn) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *LocationUpdateIn) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *LocationUpdateIn) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetHeading

`func (o *LocationUpdateIn) GetHeading() float32`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *LocationUpdateIn) GetHeadingOk() (*float32, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *LocationUpdateIn) SetHeading(v float32)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *LocationUpdateIn) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### SetHeadingNil

`func (o *LocationUpdateIn) SetHeadingNil(b bool)`

 SetHeadingNil sets the value for Heading to be an explicit nil

### UnsetHeading
`func (o *LocationUpdateIn) UnsetHeading()`

UnsetHeading ensures that no value is present for Heading, not even an explicit nil
### GetAltitude

`func (o *LocationUpdateIn) GetAltitude() float32`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *LocationUpdateIn) GetAltitudeOk() (*float32, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *LocationUpdateIn) SetAltitude(v float32)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *LocationUpdateIn) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### SetAltitudeNil

`func (o *LocationUpdateIn) SetAltitudeNil(b bool)`

 SetAltitudeNil sets the value for Altitude to be an explicit nil

### UnsetAltitude
`func (o *LocationUpdateIn) UnsetAltitude()`

UnsetAltitude ensures that no value is present for Altitude, not even an explicit nil
### GetMetadata

`func (o *LocationUpdateIn) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LocationUpdateIn) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LocationUpdateIn) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LocationUpdateIn) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *LocationUpdateIn) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *LocationUpdateIn) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


