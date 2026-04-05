# BatchLocationUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**Locations** | [**[]LocationUpdateIn**](LocationUpdateIn.md) |  | 

## Methods

### NewBatchLocationUpdateIn

`func NewBatchLocationUpdateIn(deviceId string, locations []LocationUpdateIn, ) *BatchLocationUpdateIn`

NewBatchLocationUpdateIn instantiates a new BatchLocationUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchLocationUpdateInWithDefaults

`func NewBatchLocationUpdateInWithDefaults() *BatchLocationUpdateIn`

NewBatchLocationUpdateInWithDefaults instantiates a new BatchLocationUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *BatchLocationUpdateIn) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BatchLocationUpdateIn) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BatchLocationUpdateIn) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetLocations

`func (o *BatchLocationUpdateIn) GetLocations() []LocationUpdateIn`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *BatchLocationUpdateIn) GetLocationsOk() (*[]LocationUpdateIn, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *BatchLocationUpdateIn) SetLocations(v []LocationUpdateIn)`

SetLocations sets Locations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


