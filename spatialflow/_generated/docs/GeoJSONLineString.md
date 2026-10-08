# GeoJSONLineString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Coordinates** | **[][]float32** | Two-dimensional [longitude, latitude] positions; maximum 2000 per planned route. | 

## Methods

### NewGeoJSONLineString

`func NewGeoJSONLineString(type_ string, coordinates [][]float32, ) *GeoJSONLineString`

NewGeoJSONLineString instantiates a new GeoJSONLineString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoJSONLineStringWithDefaults

`func NewGeoJSONLineStringWithDefaults() *GeoJSONLineString`

NewGeoJSONLineStringWithDefaults instantiates a new GeoJSONLineString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoJSONLineString) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoJSONLineString) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoJSONLineString) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *GeoJSONLineString) GetCoordinates() [][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *GeoJSONLineString) GetCoordinatesOk() (*[][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *GeoJSONLineString) SetCoordinates(v [][]float32)`

SetCoordinates sets Coordinates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


