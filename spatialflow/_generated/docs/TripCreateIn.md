# TripCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] [default to ""]
**PlannedRoute** | Pointer to **map[string]interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTripCreateIn

`func NewTripCreateIn(deviceId string, ) *TripCreateIn`

NewTripCreateIn instantiates a new TripCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTripCreateInWithDefaults

`func NewTripCreateInWithDefaults() *TripCreateIn`

NewTripCreateInWithDefaults instantiates a new TripCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *TripCreateIn) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *TripCreateIn) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *TripCreateIn) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetName

`func (o *TripCreateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TripCreateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TripCreateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TripCreateIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlannedRoute

`func (o *TripCreateIn) GetPlannedRoute() map[string]interface{}`

GetPlannedRoute returns the PlannedRoute field if non-nil, zero value otherwise.

### GetPlannedRouteOk

`func (o *TripCreateIn) GetPlannedRouteOk() (*map[string]interface{}, bool)`

GetPlannedRouteOk returns a tuple with the PlannedRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRoute

`func (o *TripCreateIn) SetPlannedRoute(v map[string]interface{})`

SetPlannedRoute sets PlannedRoute field to given value.

### HasPlannedRoute

`func (o *TripCreateIn) HasPlannedRoute() bool`

HasPlannedRoute returns a boolean if a field has been set.

### SetPlannedRouteNil

`func (o *TripCreateIn) SetPlannedRouteNil(b bool)`

 SetPlannedRouteNil sets the value for PlannedRoute to be an explicit nil

### UnsetPlannedRoute
`func (o *TripCreateIn) UnsetPlannedRoute()`

UnsetPlannedRoute ensures that no value is present for PlannedRoute, not even an explicit nil
### GetMetadata

`func (o *TripCreateIn) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TripCreateIn) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TripCreateIn) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TripCreateIn) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TripCreateIn) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TripCreateIn) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


