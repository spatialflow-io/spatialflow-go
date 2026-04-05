# TestPointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Geometry** | Pointer to [**NullableGeoJSONPoint**](GeoJSONPoint.md) |  | [optional] 
**Lat** | Pointer to **NullableFloat32** |  | [optional] 
**Lng** | Pointer to **NullableFloat32** |  | [optional] 
**GeofenceIds** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**RequestMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTestPointRequest

`func NewTestPointRequest() *TestPointRequest`

NewTestPointRequest instantiates a new TestPointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointRequestWithDefaults

`func NewTestPointRequestWithDefaults() *TestPointRequest`

NewTestPointRequestWithDefaults instantiates a new TestPointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeometry

`func (o *TestPointRequest) GetGeometry() GeoJSONPoint`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *TestPointRequest) GetGeometryOk() (*GeoJSONPoint, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *TestPointRequest) SetGeometry(v GeoJSONPoint)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *TestPointRequest) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *TestPointRequest) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *TestPointRequest) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetLat

`func (o *TestPointRequest) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *TestPointRequest) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *TestPointRequest) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *TestPointRequest) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *TestPointRequest) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *TestPointRequest) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *TestPointRequest) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *TestPointRequest) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *TestPointRequest) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *TestPointRequest) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *TestPointRequest) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *TestPointRequest) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetGeofenceIds

`func (o *TestPointRequest) GetGeofenceIds() []string`

GetGeofenceIds returns the GeofenceIds field if non-nil, zero value otherwise.

### GetGeofenceIdsOk

`func (o *TestPointRequest) GetGeofenceIdsOk() (*[]string, bool)`

GetGeofenceIdsOk returns a tuple with the GeofenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceIds

`func (o *TestPointRequest) SetGeofenceIds(v []string)`

SetGeofenceIds sets GeofenceIds field to given value.

### HasGeofenceIds

`func (o *TestPointRequest) HasGeofenceIds() bool`

HasGeofenceIds returns a boolean if a field has been set.

### SetGeofenceIdsNil

`func (o *TestPointRequest) SetGeofenceIdsNil(b bool)`

 SetGeofenceIdsNil sets the value for GeofenceIds to be an explicit nil

### UnsetGeofenceIds
`func (o *TestPointRequest) UnsetGeofenceIds()`

UnsetGeofenceIds ensures that no value is present for GeofenceIds, not even an explicit nil
### GetMetadata

`func (o *TestPointRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TestPointRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TestPointRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TestPointRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TestPointRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TestPointRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRequestMetadata

`func (o *TestPointRequest) GetRequestMetadata() map[string]interface{}`

GetRequestMetadata returns the RequestMetadata field if non-nil, zero value otherwise.

### GetRequestMetadataOk

`func (o *TestPointRequest) GetRequestMetadataOk() (*map[string]interface{}, bool)`

GetRequestMetadataOk returns a tuple with the RequestMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMetadata

`func (o *TestPointRequest) SetRequestMetadata(v map[string]interface{})`

SetRequestMetadata sets RequestMetadata field to given value.

### HasRequestMetadata

`func (o *TestPointRequest) HasRequestMetadata() bool`

HasRequestMetadata returns a boolean if a field has been set.

### SetRequestMetadataNil

`func (o *TestPointRequest) SetRequestMetadataNil(b bool)`

 SetRequestMetadataNil sets the value for RequestMetadata to be an explicit nil

### UnsetRequestMetadata
`func (o *TestPointRequest) UnsetRequestMetadata()`

UnsetRequestMetadata ensures that no value is present for RequestMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


