# IngestStatsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalIngestedToday** | **int32** |  | 
**TotalIngestedWeek** | **int32** |  | 
**DevicesActive** | **int32** |  | 
**LastIngest** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIngestStatsOut

`func NewIngestStatsOut(totalIngestedToday int32, totalIngestedWeek int32, devicesActive int32, ) *IngestStatsOut`

NewIngestStatsOut instantiates a new IngestStatsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestStatsOutWithDefaults

`func NewIngestStatsOutWithDefaults() *IngestStatsOut`

NewIngestStatsOutWithDefaults instantiates a new IngestStatsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalIngestedToday

`func (o *IngestStatsOut) GetTotalIngestedToday() int32`

GetTotalIngestedToday returns the TotalIngestedToday field if non-nil, zero value otherwise.

### GetTotalIngestedTodayOk

`func (o *IngestStatsOut) GetTotalIngestedTodayOk() (*int32, bool)`

GetTotalIngestedTodayOk returns a tuple with the TotalIngestedToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIngestedToday

`func (o *IngestStatsOut) SetTotalIngestedToday(v int32)`

SetTotalIngestedToday sets TotalIngestedToday field to given value.


### GetTotalIngestedWeek

`func (o *IngestStatsOut) GetTotalIngestedWeek() int32`

GetTotalIngestedWeek returns the TotalIngestedWeek field if non-nil, zero value otherwise.

### GetTotalIngestedWeekOk

`func (o *IngestStatsOut) GetTotalIngestedWeekOk() (*int32, bool)`

GetTotalIngestedWeekOk returns a tuple with the TotalIngestedWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIngestedWeek

`func (o *IngestStatsOut) SetTotalIngestedWeek(v int32)`

SetTotalIngestedWeek sets TotalIngestedWeek field to given value.


### GetDevicesActive

`func (o *IngestStatsOut) GetDevicesActive() int32`

GetDevicesActive returns the DevicesActive field if non-nil, zero value otherwise.

### GetDevicesActiveOk

`func (o *IngestStatsOut) GetDevicesActiveOk() (*int32, bool)`

GetDevicesActiveOk returns a tuple with the DevicesActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesActive

`func (o *IngestStatsOut) SetDevicesActive(v int32)`

SetDevicesActive sets DevicesActive field to given value.


### GetLastIngest

`func (o *IngestStatsOut) GetLastIngest() string`

GetLastIngest returns the LastIngest field if non-nil, zero value otherwise.

### GetLastIngestOk

`func (o *IngestStatsOut) GetLastIngestOk() (*string, bool)`

GetLastIngestOk returns a tuple with the LastIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIngest

`func (o *IngestStatsOut) SetLastIngest(v string)`

SetLastIngest sets LastIngest field to given value.

### HasLastIngest

`func (o *IngestStatsOut) HasLastIngest() bool`

HasLastIngest returns a boolean if a field has been set.

### SetLastIngestNil

`func (o *IngestStatsOut) SetLastIngestNil(b bool)`

 SetLastIngestNil sets the value for LastIngest to be an explicit nil

### UnsetLastIngest
`func (o *IngestStatsOut) UnsetLastIngest()`

UnsetLastIngest ensures that no value is present for LastIngest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


