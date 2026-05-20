# DashboardTrendsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** |  | 
**Series** | [**[]TrendSeries**](TrendSeries.md) |  | 

## Methods

### NewDashboardTrendsResponse

`func NewDashboardTrendsResponse(days int32, series []TrendSeries, ) *DashboardTrendsResponse`

NewDashboardTrendsResponse instantiates a new DashboardTrendsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardTrendsResponseWithDefaults

`func NewDashboardTrendsResponseWithDefaults() *DashboardTrendsResponse`

NewDashboardTrendsResponseWithDefaults instantiates a new DashboardTrendsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *DashboardTrendsResponse) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *DashboardTrendsResponse) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *DashboardTrendsResponse) SetDays(v int32)`

SetDays sets Days field to given value.


### GetSeries

`func (o *DashboardTrendsResponse) GetSeries() []TrendSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *DashboardTrendsResponse) GetSeriesOk() (*[]TrendSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *DashboardTrendsResponse) SetSeries(v []TrendSeries)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


