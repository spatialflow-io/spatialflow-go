# EmailQueueStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | **int32** |  | 
**Queued** | **int32** |  | 
**Delivered** | **int32** |  | 
**Skipped** | **int32** |  | 
**Failed** | **int32** |  | 
**TotalToday** | **int32** |  | 
**TotalThisHour** | **int32** |  | 

## Methods

### NewEmailQueueStats

`func NewEmailQueueStats(pending int32, queued int32, delivered int32, skipped int32, failed int32, totalToday int32, totalThisHour int32, ) *EmailQueueStats`

NewEmailQueueStats instantiates a new EmailQueueStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailQueueStatsWithDefaults

`func NewEmailQueueStatsWithDefaults() *EmailQueueStats`

NewEmailQueueStatsWithDefaults instantiates a new EmailQueueStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *EmailQueueStats) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *EmailQueueStats) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *EmailQueueStats) SetPending(v int32)`

SetPending sets Pending field to given value.


### GetQueued

`func (o *EmailQueueStats) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *EmailQueueStats) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *EmailQueueStats) SetQueued(v int32)`

SetQueued sets Queued field to given value.


### GetDelivered

`func (o *EmailQueueStats) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *EmailQueueStats) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *EmailQueueStats) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.


### GetSkipped

`func (o *EmailQueueStats) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *EmailQueueStats) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *EmailQueueStats) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.


### GetFailed

`func (o *EmailQueueStats) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *EmailQueueStats) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *EmailQueueStats) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetTotalToday

`func (o *EmailQueueStats) GetTotalToday() int32`

GetTotalToday returns the TotalToday field if non-nil, zero value otherwise.

### GetTotalTodayOk

`func (o *EmailQueueStats) GetTotalTodayOk() (*int32, bool)`

GetTotalTodayOk returns a tuple with the TotalToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalToday

`func (o *EmailQueueStats) SetTotalToday(v int32)`

SetTotalToday sets TotalToday field to given value.


### GetTotalThisHour

`func (o *EmailQueueStats) GetTotalThisHour() int32`

GetTotalThisHour returns the TotalThisHour field if non-nil, zero value otherwise.

### GetTotalThisHourOk

`func (o *EmailQueueStats) GetTotalThisHourOk() (*int32, bool)`

GetTotalThisHourOk returns a tuple with the TotalThisHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalThisHour

`func (o *EmailQueueStats) SetTotalThisHour(v int32)`

SetTotalThisHour sets TotalThisHour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


