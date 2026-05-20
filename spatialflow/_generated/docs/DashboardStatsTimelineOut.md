# DashboardStatsTimelineOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | **string** |  | 
**Buckets** | [**[]TimelineBucketOut**](TimelineBucketOut.md) |  | 
**Previous** | [**TimelinePreviousOut**](TimelinePreviousOut.md) |  | 

## Methods

### NewDashboardStatsTimelineOut

`func NewDashboardStatsTimelineOut(timeRange string, buckets []TimelineBucketOut, previous TimelinePreviousOut, ) *DashboardStatsTimelineOut`

NewDashboardStatsTimelineOut instantiates a new DashboardStatsTimelineOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardStatsTimelineOutWithDefaults

`func NewDashboardStatsTimelineOutWithDefaults() *DashboardStatsTimelineOut`

NewDashboardStatsTimelineOutWithDefaults instantiates a new DashboardStatsTimelineOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *DashboardStatsTimelineOut) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DashboardStatsTimelineOut) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DashboardStatsTimelineOut) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetBuckets

`func (o *DashboardStatsTimelineOut) GetBuckets() []TimelineBucketOut`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *DashboardStatsTimelineOut) GetBucketsOk() (*[]TimelineBucketOut, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *DashboardStatsTimelineOut) SetBuckets(v []TimelineBucketOut)`

SetBuckets sets Buckets field to given value.


### GetPrevious

`func (o *DashboardStatsTimelineOut) GetPrevious() TimelinePreviousOut`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *DashboardStatsTimelineOut) GetPreviousOk() (*TimelinePreviousOut, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *DashboardStatsTimelineOut) SetPrevious(v TimelinePreviousOut)`

SetPrevious sets Previous field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


