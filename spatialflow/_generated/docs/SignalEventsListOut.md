# SignalEventsListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signals** | [**[]SignalEventOut**](SignalEventOut.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewSignalEventsListOut

`func NewSignalEventsListOut(signals []SignalEventOut, totalCount int32, ) *SignalEventsListOut`

NewSignalEventsListOut instantiates a new SignalEventsListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalEventsListOutWithDefaults

`func NewSignalEventsListOutWithDefaults() *SignalEventsListOut`

NewSignalEventsListOutWithDefaults instantiates a new SignalEventsListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignals

`func (o *SignalEventsListOut) GetSignals() []SignalEventOut`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *SignalEventsListOut) GetSignalsOk() (*[]SignalEventOut, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *SignalEventsListOut) SetSignals(v []SignalEventOut)`

SetSignals sets Signals field to given value.


### GetTotalCount

`func (o *SignalEventsListOut) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SignalEventsListOut) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SignalEventsListOut) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


