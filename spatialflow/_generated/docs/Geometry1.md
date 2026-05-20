# Geometry1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Coordinates** | [**[][][][]float32**]([][][]float32.md) |  | 
**Center** | **[]float32** | [longitude, latitude] of circle center | 
**RadiusMeters** | **float32** | Circle radius in meters (max 100km) | 

## Methods

### NewGeometry1

`func NewGeometry1(type_ string, coordinates [][][][]float32, center []float32, radiusMeters float32, ) *Geometry1`

NewGeometry1 instantiates a new Geometry1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeometry1WithDefaults

`func NewGeometry1WithDefaults() *Geometry1`

NewGeometry1WithDefaults instantiates a new Geometry1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Geometry1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Geometry1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Geometry1) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *Geometry1) GetCoordinates() [][][][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Geometry1) GetCoordinatesOk() (*[][][][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Geometry1) SetCoordinates(v [][][][]float32)`

SetCoordinates sets Coordinates field to given value.


### GetCenter

`func (o *Geometry1) GetCenter() []float32`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *Geometry1) GetCenterOk() (*[]float32, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *Geometry1) SetCenter(v []float32)`

SetCenter sets Center field to given value.


### GetRadiusMeters

`func (o *Geometry1) GetRadiusMeters() float32`

GetRadiusMeters returns the RadiusMeters field if non-nil, zero value otherwise.

### GetRadiusMetersOk

`func (o *Geometry1) GetRadiusMetersOk() (*float32, bool)`

GetRadiusMetersOk returns a tuple with the RadiusMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusMeters

`func (o *Geometry1) SetRadiusMeters(v float32)`

SetRadiusMeters sets RadiusMeters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


