# BatchResendOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]BatchResendItemOut**](BatchResendItemOut.md) |  | 
**ResentCount** | **int32** |  | 
**SkippedCount** | **int32** |  | 
**NextCursor** | Pointer to **NullableString** |  | [optional] 
**RemainingCount** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewBatchResendOut

`func NewBatchResendOut(results []BatchResendItemOut, resentCount int32, skippedCount int32, ) *BatchResendOut`

NewBatchResendOut instantiates a new BatchResendOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResendOutWithDefaults

`func NewBatchResendOutWithDefaults() *BatchResendOut`

NewBatchResendOutWithDefaults instantiates a new BatchResendOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BatchResendOut) GetResults() []BatchResendItemOut`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResendOut) GetResultsOk() (*[]BatchResendItemOut, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResendOut) SetResults(v []BatchResendItemOut)`

SetResults sets Results field to given value.


### GetResentCount

`func (o *BatchResendOut) GetResentCount() int32`

GetResentCount returns the ResentCount field if non-nil, zero value otherwise.

### GetResentCountOk

`func (o *BatchResendOut) GetResentCountOk() (*int32, bool)`

GetResentCountOk returns a tuple with the ResentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResentCount

`func (o *BatchResendOut) SetResentCount(v int32)`

SetResentCount sets ResentCount field to given value.


### GetSkippedCount

`func (o *BatchResendOut) GetSkippedCount() int32`

GetSkippedCount returns the SkippedCount field if non-nil, zero value otherwise.

### GetSkippedCountOk

`func (o *BatchResendOut) GetSkippedCountOk() (*int32, bool)`

GetSkippedCountOk returns a tuple with the SkippedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedCount

`func (o *BatchResendOut) SetSkippedCount(v int32)`

SetSkippedCount sets SkippedCount field to given value.


### GetNextCursor

`func (o *BatchResendOut) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *BatchResendOut) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *BatchResendOut) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *BatchResendOut) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *BatchResendOut) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *BatchResendOut) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetRemainingCount

`func (o *BatchResendOut) GetRemainingCount() int32`

GetRemainingCount returns the RemainingCount field if non-nil, zero value otherwise.

### GetRemainingCountOk

`func (o *BatchResendOut) GetRemainingCountOk() (*int32, bool)`

GetRemainingCountOk returns a tuple with the RemainingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCount

`func (o *BatchResendOut) SetRemainingCount(v int32)`

SetRemainingCount sets RemainingCount field to given value.

### HasRemainingCount

`func (o *BatchResendOut) HasRemainingCount() bool`

HasRemainingCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


