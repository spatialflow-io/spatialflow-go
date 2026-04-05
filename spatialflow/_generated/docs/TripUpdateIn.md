# TripUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**PlannedRoute** | Pointer to **map[string]interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTripUpdateIn

`func NewTripUpdateIn() *TripUpdateIn`

NewTripUpdateIn instantiates a new TripUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTripUpdateInWithDefaults

`func NewTripUpdateInWithDefaults() *TripUpdateIn`

NewTripUpdateInWithDefaults instantiates a new TripUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TripUpdateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TripUpdateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TripUpdateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TripUpdateIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TripUpdateIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TripUpdateIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlannedRoute

`func (o *TripUpdateIn) GetPlannedRoute() map[string]interface{}`

GetPlannedRoute returns the PlannedRoute field if non-nil, zero value otherwise.

### GetPlannedRouteOk

`func (o *TripUpdateIn) GetPlannedRouteOk() (*map[string]interface{}, bool)`

GetPlannedRouteOk returns a tuple with the PlannedRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRoute

`func (o *TripUpdateIn) SetPlannedRoute(v map[string]interface{})`

SetPlannedRoute sets PlannedRoute field to given value.

### HasPlannedRoute

`func (o *TripUpdateIn) HasPlannedRoute() bool`

HasPlannedRoute returns a boolean if a field has been set.

### SetPlannedRouteNil

`func (o *TripUpdateIn) SetPlannedRouteNil(b bool)`

 SetPlannedRouteNil sets the value for PlannedRoute to be an explicit nil

### UnsetPlannedRoute
`func (o *TripUpdateIn) UnsetPlannedRoute()`

UnsetPlannedRoute ensures that no value is present for PlannedRoute, not even an explicit nil
### GetMetadata

`func (o *TripUpdateIn) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TripUpdateIn) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TripUpdateIn) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TripUpdateIn) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TripUpdateIn) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TripUpdateIn) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


