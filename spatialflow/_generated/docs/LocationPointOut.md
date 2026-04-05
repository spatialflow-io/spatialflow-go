# LocationPointOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float32** |  | 
**Longitude** | **float32** |  | 
**Timestamp** | **time.Time** |  | 
**Accuracy** | Pointer to **NullableFloat32** |  | [optional] 
**Speed** | Pointer to **NullableFloat32** |  | [optional] 
**Heading** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewLocationPointOut

`func NewLocationPointOut(latitude float32, longitude float32, timestamp time.Time, ) *LocationPointOut`

NewLocationPointOut instantiates a new LocationPointOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationPointOutWithDefaults

`func NewLocationPointOutWithDefaults() *LocationPointOut`

NewLocationPointOutWithDefaults instantiates a new LocationPointOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *LocationPointOut) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationPointOut) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationPointOut) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *LocationPointOut) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationPointOut) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationPointOut) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetTimestamp

`func (o *LocationPointOut) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LocationPointOut) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LocationPointOut) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetAccuracy

`func (o *LocationPointOut) GetAccuracy() float32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *LocationPointOut) GetAccuracyOk() (*float32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *LocationPointOut) SetAccuracy(v float32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *LocationPointOut) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### SetAccuracyNil

`func (o *LocationPointOut) SetAccuracyNil(b bool)`

 SetAccuracyNil sets the value for Accuracy to be an explicit nil

### UnsetAccuracy
`func (o *LocationPointOut) UnsetAccuracy()`

UnsetAccuracy ensures that no value is present for Accuracy, not even an explicit nil
### GetSpeed

`func (o *LocationPointOut) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *LocationPointOut) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *LocationPointOut) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *LocationPointOut) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *LocationPointOut) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *LocationPointOut) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetHeading

`func (o *LocationPointOut) GetHeading() float32`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *LocationPointOut) GetHeadingOk() (*float32, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *LocationPointOut) SetHeading(v float32)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *LocationPointOut) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### SetHeadingNil

`func (o *LocationPointOut) SetHeadingNil(b bool)`

 SetHeadingNil sets the value for Heading to be an explicit nil

### UnsetHeading
`func (o *LocationPointOut) UnsetHeading()`

UnsetHeading ensures that no value is present for Heading, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


