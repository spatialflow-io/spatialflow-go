# ActiveGeofenceSummaryItemOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Geometry** | **map[string]interface{}** |  | 
**Center** | **[]float32** |  | 
**EventCountToday** | **int32** |  | 

## Methods

### NewActiveGeofenceSummaryItemOut

`func NewActiveGeofenceSummaryItemOut(id string, name string, geometry map[string]interface{}, center []float32, eventCountToday int32, ) *ActiveGeofenceSummaryItemOut`

NewActiveGeofenceSummaryItemOut instantiates a new ActiveGeofenceSummaryItemOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGeofenceSummaryItemOutWithDefaults

`func NewActiveGeofenceSummaryItemOutWithDefaults() *ActiveGeofenceSummaryItemOut`

NewActiveGeofenceSummaryItemOutWithDefaults instantiates a new ActiveGeofenceSummaryItemOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActiveGeofenceSummaryItemOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveGeofenceSummaryItemOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveGeofenceSummaryItemOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ActiveGeofenceSummaryItemOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActiveGeofenceSummaryItemOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActiveGeofenceSummaryItemOut) SetName(v string)`

SetName sets Name field to given value.


### GetGeometry

`func (o *ActiveGeofenceSummaryItemOut) GetGeometry() map[string]interface{}`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *ActiveGeofenceSummaryItemOut) GetGeometryOk() (*map[string]interface{}, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *ActiveGeofenceSummaryItemOut) SetGeometry(v map[string]interface{})`

SetGeometry sets Geometry field to given value.


### GetCenter

`func (o *ActiveGeofenceSummaryItemOut) GetCenter() []float32`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *ActiveGeofenceSummaryItemOut) GetCenterOk() (*[]float32, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *ActiveGeofenceSummaryItemOut) SetCenter(v []float32)`

SetCenter sets Center field to given value.


### GetEventCountToday

`func (o *ActiveGeofenceSummaryItemOut) GetEventCountToday() int32`

GetEventCountToday returns the EventCountToday field if non-nil, zero value otherwise.

### GetEventCountTodayOk

`func (o *ActiveGeofenceSummaryItemOut) GetEventCountTodayOk() (*int32, bool)`

GetEventCountTodayOk returns a tuple with the EventCountToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCountToday

`func (o *ActiveGeofenceSummaryItemOut) SetEventCountToday(v int32)`

SetEventCountToday sets EventCountToday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


