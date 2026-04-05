# UsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationEvents** | **int32** |  | 
**ActionDeliveries** | **int32** |  | 
**EventUnits** | **float32** |  | 
**Tier** | **string** |  | 
**TierLimit** | **NullableInt32** |  | 
**PeriodStart** | **time.Time** |  | 
**PeriodEnd** | **time.Time** |  | 

## Methods

### NewUsageResponse

`func NewUsageResponse(locationEvents int32, actionDeliveries int32, eventUnits float32, tier string, tierLimit NullableInt32, periodStart time.Time, periodEnd time.Time, ) *UsageResponse`

NewUsageResponse instantiates a new UsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageResponseWithDefaults

`func NewUsageResponseWithDefaults() *UsageResponse`

NewUsageResponseWithDefaults instantiates a new UsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationEvents

`func (o *UsageResponse) GetLocationEvents() int32`

GetLocationEvents returns the LocationEvents field if non-nil, zero value otherwise.

### GetLocationEventsOk

`func (o *UsageResponse) GetLocationEventsOk() (*int32, bool)`

GetLocationEventsOk returns a tuple with the LocationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEvents

`func (o *UsageResponse) SetLocationEvents(v int32)`

SetLocationEvents sets LocationEvents field to given value.


### GetActionDeliveries

`func (o *UsageResponse) GetActionDeliveries() int32`

GetActionDeliveries returns the ActionDeliveries field if non-nil, zero value otherwise.

### GetActionDeliveriesOk

`func (o *UsageResponse) GetActionDeliveriesOk() (*int32, bool)`

GetActionDeliveriesOk returns a tuple with the ActionDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDeliveries

`func (o *UsageResponse) SetActionDeliveries(v int32)`

SetActionDeliveries sets ActionDeliveries field to given value.


### GetEventUnits

`func (o *UsageResponse) GetEventUnits() float32`

GetEventUnits returns the EventUnits field if non-nil, zero value otherwise.

### GetEventUnitsOk

`func (o *UsageResponse) GetEventUnitsOk() (*float32, bool)`

GetEventUnitsOk returns a tuple with the EventUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUnits

`func (o *UsageResponse) SetEventUnits(v float32)`

SetEventUnits sets EventUnits field to given value.


### GetTier

`func (o *UsageResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *UsageResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *UsageResponse) SetTier(v string)`

SetTier sets Tier field to given value.


### GetTierLimit

`func (o *UsageResponse) GetTierLimit() int32`

GetTierLimit returns the TierLimit field if non-nil, zero value otherwise.

### GetTierLimitOk

`func (o *UsageResponse) GetTierLimitOk() (*int32, bool)`

GetTierLimitOk returns a tuple with the TierLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierLimit

`func (o *UsageResponse) SetTierLimit(v int32)`

SetTierLimit sets TierLimit field to given value.


### SetTierLimitNil

`func (o *UsageResponse) SetTierLimitNil(b bool)`

 SetTierLimitNil sets the value for TierLimit to be an explicit nil

### UnsetTierLimit
`func (o *UsageResponse) UnsetTierLimit()`

UnsetTierLimit ensures that no value is present for TierLimit, not even an explicit nil
### GetPeriodStart

`func (o *UsageResponse) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UsageResponse) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UsageResponse) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *UsageResponse) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UsageResponse) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UsageResponse) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


