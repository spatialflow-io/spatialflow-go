# RecentLocationsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | [**[]LocationPointOut**](LocationPointOut.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewRecentLocationsOut

`func NewRecentLocationsOut(locations []LocationPointOut, totalCount int32, ) *RecentLocationsOut`

NewRecentLocationsOut instantiates a new RecentLocationsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecentLocationsOutWithDefaults

`func NewRecentLocationsOutWithDefaults() *RecentLocationsOut`

NewRecentLocationsOutWithDefaults instantiates a new RecentLocationsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *RecentLocationsOut) GetLocations() []LocationPointOut`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *RecentLocationsOut) GetLocationsOk() (*[]LocationPointOut, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *RecentLocationsOut) SetLocations(v []LocationPointOut)`

SetLocations sets Locations field to given value.


### GetTotalCount

`func (o *RecentLocationsOut) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *RecentLocationsOut) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *RecentLocationsOut) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


