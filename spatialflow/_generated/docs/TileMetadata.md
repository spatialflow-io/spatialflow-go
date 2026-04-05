# TileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Format** | **string** |  | 
**Minzoom** | **int32** |  | 
**Maxzoom** | **int32** |  | 
**Bounds** | **[]float32** |  | 
**Center** | **[]float32** |  | 
**Layers** | **[]map[string]interface{}** |  | 

## Methods

### NewTileMetadata

`func NewTileMetadata(name string, format string, minzoom int32, maxzoom int32, bounds []float32, center []float32, layers []*map[string]interface{}, ) *TileMetadata`

NewTileMetadata instantiates a new TileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTileMetadataWithDefaults

`func NewTileMetadataWithDefaults() *TileMetadata`

NewTileMetadataWithDefaults instantiates a new TileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TileMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TileMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TileMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetFormat

`func (o *TileMetadata) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TileMetadata) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TileMetadata) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetMinzoom

`func (o *TileMetadata) GetMinzoom() int32`

GetMinzoom returns the Minzoom field if non-nil, zero value otherwise.

### GetMinzoomOk

`func (o *TileMetadata) GetMinzoomOk() (*int32, bool)`

GetMinzoomOk returns a tuple with the Minzoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinzoom

`func (o *TileMetadata) SetMinzoom(v int32)`

SetMinzoom sets Minzoom field to given value.


### GetMaxzoom

`func (o *TileMetadata) GetMaxzoom() int32`

GetMaxzoom returns the Maxzoom field if non-nil, zero value otherwise.

### GetMaxzoomOk

`func (o *TileMetadata) GetMaxzoomOk() (*int32, bool)`

GetMaxzoomOk returns a tuple with the Maxzoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxzoom

`func (o *TileMetadata) SetMaxzoom(v int32)`

SetMaxzoom sets Maxzoom field to given value.


### GetBounds

`func (o *TileMetadata) GetBounds() []float32`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *TileMetadata) GetBoundsOk() (*[]float32, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *TileMetadata) SetBounds(v []float32)`

SetBounds sets Bounds field to given value.


### GetCenter

`func (o *TileMetadata) GetCenter() []float32`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *TileMetadata) GetCenterOk() (*[]float32, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *TileMetadata) SetCenter(v []float32)`

SetCenter sets Center field to given value.


### GetLayers

`func (o *TileMetadata) GetLayers() []*map[string]interface{}`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *TileMetadata) GetLayersOk() (*[]*map[string]interface{}, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *TileMetadata) SetLayers(v []*map[string]interface{})`

SetLayers sets Layers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


