# TileHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Service** | **string** |  | 
**MvtEnabled** | **bool** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewTileHealthResponse

`func NewTileHealthResponse(status string, service string, mvtEnabled bool, timestamp time.Time, ) *TileHealthResponse`

NewTileHealthResponse instantiates a new TileHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTileHealthResponseWithDefaults

`func NewTileHealthResponseWithDefaults() *TileHealthResponse`

NewTileHealthResponseWithDefaults instantiates a new TileHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TileHealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TileHealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TileHealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetService

`func (o *TileHealthResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *TileHealthResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *TileHealthResponse) SetService(v string)`

SetService sets Service field to given value.


### GetMvtEnabled

`func (o *TileHealthResponse) GetMvtEnabled() bool`

GetMvtEnabled returns the MvtEnabled field if non-nil, zero value otherwise.

### GetMvtEnabledOk

`func (o *TileHealthResponse) GetMvtEnabledOk() (*bool, bool)`

GetMvtEnabledOk returns a tuple with the MvtEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvtEnabled

`func (o *TileHealthResponse) SetMvtEnabled(v bool)`

SetMvtEnabled sets MvtEnabled field to given value.


### GetTimestamp

`func (o *TileHealthResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TileHealthResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TileHealthResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


