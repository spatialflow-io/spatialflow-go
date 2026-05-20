# EventLocationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float32** |  | 
**Longitude** | **float32** |  | 
**Accuracy** | Pointer to **NullableFloat32** |  | [optional] 
**Speed** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewEventLocationOut

`func NewEventLocationOut(latitude float32, longitude float32, ) *EventLocationOut`

NewEventLocationOut instantiates a new EventLocationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLocationOutWithDefaults

`func NewEventLocationOutWithDefaults() *EventLocationOut`

NewEventLocationOutWithDefaults instantiates a new EventLocationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *EventLocationOut) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *EventLocationOut) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *EventLocationOut) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *EventLocationOut) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *EventLocationOut) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *EventLocationOut) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetAccuracy

`func (o *EventLocationOut) GetAccuracy() float32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *EventLocationOut) GetAccuracyOk() (*float32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *EventLocationOut) SetAccuracy(v float32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *EventLocationOut) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### SetAccuracyNil

`func (o *EventLocationOut) SetAccuracyNil(b bool)`

 SetAccuracyNil sets the value for Accuracy to be an explicit nil

### UnsetAccuracy
`func (o *EventLocationOut) UnsetAccuracy()`

UnsetAccuracy ensures that no value is present for Accuracy, not even an explicit nil
### GetSpeed

`func (o *EventLocationOut) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *EventLocationOut) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *EventLocationOut) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *EventLocationOut) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *EventLocationOut) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *EventLocationOut) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


