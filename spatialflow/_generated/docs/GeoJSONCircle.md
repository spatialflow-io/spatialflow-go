# GeoJSONCircle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Center** | **[]float32** | [longitude, latitude] of circle center | 
**RadiusMeters** | **float32** | Circle radius in meters (max 100km) | 

## Methods

### NewGeoJSONCircle

`func NewGeoJSONCircle(type_ string, center []float32, radiusMeters float32, ) *GeoJSONCircle`

NewGeoJSONCircle instantiates a new GeoJSONCircle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoJSONCircleWithDefaults

`func NewGeoJSONCircleWithDefaults() *GeoJSONCircle`

NewGeoJSONCircleWithDefaults instantiates a new GeoJSONCircle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoJSONCircle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoJSONCircle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoJSONCircle) SetType(v string)`

SetType sets Type field to given value.


### GetCenter

`func (o *GeoJSONCircle) GetCenter() []float32`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *GeoJSONCircle) GetCenterOk() (*[]float32, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *GeoJSONCircle) SetCenter(v []float32)`

SetCenter sets Center field to given value.


### GetRadiusMeters

`func (o *GeoJSONCircle) GetRadiusMeters() float32`

GetRadiusMeters returns the RadiusMeters field if non-nil, zero value otherwise.

### GetRadiusMetersOk

`func (o *GeoJSONCircle) GetRadiusMetersOk() (*float32, bool)`

GetRadiusMetersOk returns a tuple with the RadiusMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusMeters

`func (o *GeoJSONCircle) SetRadiusMeters(v float32)`

SetRadiusMeters sets RadiusMeters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


