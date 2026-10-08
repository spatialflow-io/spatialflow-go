# Geometry2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Coordinates** | [**[][][]float32**]([][]float32.md) | Polygon rings with at most 1000 total positions, 64 rings, and 32 holes. | 

## Methods

### NewGeometry2

`func NewGeometry2(type_ string, coordinates [][][]float32, ) *Geometry2`

NewGeometry2 instantiates a new Geometry2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeometry2WithDefaults

`func NewGeometry2WithDefaults() *Geometry2`

NewGeometry2WithDefaults instantiates a new Geometry2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Geometry2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Geometry2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Geometry2) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *Geometry2) GetCoordinates() [][][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Geometry2) GetCoordinatesOk() (*[][][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Geometry2) SetCoordinates(v [][][]float32)`

SetCoordinates sets Coordinates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


