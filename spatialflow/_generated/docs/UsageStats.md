# UsageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationEvents** | **int32** |  | 
**ActionDeliveries** | **int32** |  | 
**EventUnits** | **float32** |  | 

## Methods

### NewUsageStats

`func NewUsageStats(locationEvents int32, actionDeliveries int32, eventUnits float32, ) *UsageStats`

NewUsageStats instantiates a new UsageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageStatsWithDefaults

`func NewUsageStatsWithDefaults() *UsageStats`

NewUsageStatsWithDefaults instantiates a new UsageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationEvents

`func (o *UsageStats) GetLocationEvents() int32`

GetLocationEvents returns the LocationEvents field if non-nil, zero value otherwise.

### GetLocationEventsOk

`func (o *UsageStats) GetLocationEventsOk() (*int32, bool)`

GetLocationEventsOk returns a tuple with the LocationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEvents

`func (o *UsageStats) SetLocationEvents(v int32)`

SetLocationEvents sets LocationEvents field to given value.


### GetActionDeliveries

`func (o *UsageStats) GetActionDeliveries() int32`

GetActionDeliveries returns the ActionDeliveries field if non-nil, zero value otherwise.

### GetActionDeliveriesOk

`func (o *UsageStats) GetActionDeliveriesOk() (*int32, bool)`

GetActionDeliveriesOk returns a tuple with the ActionDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDeliveries

`func (o *UsageStats) SetActionDeliveries(v int32)`

SetActionDeliveries sets ActionDeliveries field to given value.


### GetEventUnits

`func (o *UsageStats) GetEventUnits() float32`

GetEventUnits returns the EventUnits field if non-nil, zero value otherwise.

### GetEventUnitsOk

`func (o *UsageStats) GetEventUnitsOk() (*float32, bool)`

GetEventUnitsOk returns a tuple with the EventUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUnits

`func (o *UsageStats) SetEventUnits(v float32)`

SetEventUnits sets EventUnits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


