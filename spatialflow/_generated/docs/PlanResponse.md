# PlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Price** | **float32** |  | 
**Interval** | **string** | Billing interval (month/year) | 
**Features** | [**PlanFeatures**](PlanFeatures.md) |  | 
**Limits** | [**PlanLimits**](PlanLimits.md) |  | 
**StripePriceId** | Pointer to **NullableString** |  | [optional] 
**Tier** | Pointer to **string** | Lowercase plan name (e.g., &#39;free&#39;, &#39;pro&#39;) | [optional] [default to ""]
**DisplayName** | Pointer to **string** | Human-readable plan name | [optional] [default to ""]
**PriceMonthly** | Pointer to **float32** | Monthly price in dollars | [optional] [default to 0]
**PriceYearly** | Pointer to **float32** | Yearly price in dollars | [optional] [default to 0]
**StripePriceMonthlyId** | Pointer to **NullableString** |  | [optional] 
**StripePriceYearlyId** | Pointer to **NullableString** |  | [optional] 
**EventOverageRate** | Pointer to **float32** | Price per extra 100k events | [optional] [default to 0]
**GeofenceOverageRate** | Pointer to **float32** | Price per extra 100 geofences | [optional] [default to 0]
**IsFeatured** | Pointer to **bool** | Whether this plan is featured/recommended | [optional] [default to false]

## Methods

### NewPlanResponse

`func NewPlanResponse(id string, name string, description string, price float32, interval string, features PlanFeatures, limits PlanLimits, ) *PlanResponse`

NewPlanResponse instantiates a new PlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanResponseWithDefaults

`func NewPlanResponseWithDefaults() *PlanResponse`

NewPlanResponseWithDefaults instantiates a new PlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PlanResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PlanResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPrice

`func (o *PlanResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlanResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlanResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetInterval

`func (o *PlanResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PlanResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PlanResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetFeatures

`func (o *PlanResponse) GetFeatures() PlanFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PlanResponse) GetFeaturesOk() (*PlanFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PlanResponse) SetFeatures(v PlanFeatures)`

SetFeatures sets Features field to given value.


### GetLimits

`func (o *PlanResponse) GetLimits() PlanLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *PlanResponse) GetLimitsOk() (*PlanLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *PlanResponse) SetLimits(v PlanLimits)`

SetLimits sets Limits field to given value.


### GetStripePriceId

`func (o *PlanResponse) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *PlanResponse) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *PlanResponse) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *PlanResponse) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.

### SetStripePriceIdNil

`func (o *PlanResponse) SetStripePriceIdNil(b bool)`

 SetStripePriceIdNil sets the value for StripePriceId to be an explicit nil

### UnsetStripePriceId
`func (o *PlanResponse) UnsetStripePriceId()`

UnsetStripePriceId ensures that no value is present for StripePriceId, not even an explicit nil
### GetTier

`func (o *PlanResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PlanResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PlanResponse) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *PlanResponse) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetDisplayName

`func (o *PlanResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PlanResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PlanResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PlanResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPriceMonthly

`func (o *PlanResponse) GetPriceMonthly() float32`

GetPriceMonthly returns the PriceMonthly field if non-nil, zero value otherwise.

### GetPriceMonthlyOk

`func (o *PlanResponse) GetPriceMonthlyOk() (*float32, bool)`

GetPriceMonthlyOk returns a tuple with the PriceMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonthly

`func (o *PlanResponse) SetPriceMonthly(v float32)`

SetPriceMonthly sets PriceMonthly field to given value.

### HasPriceMonthly

`func (o *PlanResponse) HasPriceMonthly() bool`

HasPriceMonthly returns a boolean if a field has been set.

### GetPriceYearly

`func (o *PlanResponse) GetPriceYearly() float32`

GetPriceYearly returns the PriceYearly field if non-nil, zero value otherwise.

### GetPriceYearlyOk

`func (o *PlanResponse) GetPriceYearlyOk() (*float32, bool)`

GetPriceYearlyOk returns a tuple with the PriceYearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceYearly

`func (o *PlanResponse) SetPriceYearly(v float32)`

SetPriceYearly sets PriceYearly field to given value.

### HasPriceYearly

`func (o *PlanResponse) HasPriceYearly() bool`

HasPriceYearly returns a boolean if a field has been set.

### GetStripePriceMonthlyId

`func (o *PlanResponse) GetStripePriceMonthlyId() string`

GetStripePriceMonthlyId returns the StripePriceMonthlyId field if non-nil, zero value otherwise.

### GetStripePriceMonthlyIdOk

`func (o *PlanResponse) GetStripePriceMonthlyIdOk() (*string, bool)`

GetStripePriceMonthlyIdOk returns a tuple with the StripePriceMonthlyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceMonthlyId

`func (o *PlanResponse) SetStripePriceMonthlyId(v string)`

SetStripePriceMonthlyId sets StripePriceMonthlyId field to given value.

### HasStripePriceMonthlyId

`func (o *PlanResponse) HasStripePriceMonthlyId() bool`

HasStripePriceMonthlyId returns a boolean if a field has been set.

### SetStripePriceMonthlyIdNil

`func (o *PlanResponse) SetStripePriceMonthlyIdNil(b bool)`

 SetStripePriceMonthlyIdNil sets the value for StripePriceMonthlyId to be an explicit nil

### UnsetStripePriceMonthlyId
`func (o *PlanResponse) UnsetStripePriceMonthlyId()`

UnsetStripePriceMonthlyId ensures that no value is present for StripePriceMonthlyId, not even an explicit nil
### GetStripePriceYearlyId

`func (o *PlanResponse) GetStripePriceYearlyId() string`

GetStripePriceYearlyId returns the StripePriceYearlyId field if non-nil, zero value otherwise.

### GetStripePriceYearlyIdOk

`func (o *PlanResponse) GetStripePriceYearlyIdOk() (*string, bool)`

GetStripePriceYearlyIdOk returns a tuple with the StripePriceYearlyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceYearlyId

`func (o *PlanResponse) SetStripePriceYearlyId(v string)`

SetStripePriceYearlyId sets StripePriceYearlyId field to given value.

### HasStripePriceYearlyId

`func (o *PlanResponse) HasStripePriceYearlyId() bool`

HasStripePriceYearlyId returns a boolean if a field has been set.

### SetStripePriceYearlyIdNil

`func (o *PlanResponse) SetStripePriceYearlyIdNil(b bool)`

 SetStripePriceYearlyIdNil sets the value for StripePriceYearlyId to be an explicit nil

### UnsetStripePriceYearlyId
`func (o *PlanResponse) UnsetStripePriceYearlyId()`

UnsetStripePriceYearlyId ensures that no value is present for StripePriceYearlyId, not even an explicit nil
### GetEventOverageRate

`func (o *PlanResponse) GetEventOverageRate() float32`

GetEventOverageRate returns the EventOverageRate field if non-nil, zero value otherwise.

### GetEventOverageRateOk

`func (o *PlanResponse) GetEventOverageRateOk() (*float32, bool)`

GetEventOverageRateOk returns a tuple with the EventOverageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOverageRate

`func (o *PlanResponse) SetEventOverageRate(v float32)`

SetEventOverageRate sets EventOverageRate field to given value.

### HasEventOverageRate

`func (o *PlanResponse) HasEventOverageRate() bool`

HasEventOverageRate returns a boolean if a field has been set.

### GetGeofenceOverageRate

`func (o *PlanResponse) GetGeofenceOverageRate() float32`

GetGeofenceOverageRate returns the GeofenceOverageRate field if non-nil, zero value otherwise.

### GetGeofenceOverageRateOk

`func (o *PlanResponse) GetGeofenceOverageRateOk() (*float32, bool)`

GetGeofenceOverageRateOk returns a tuple with the GeofenceOverageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceOverageRate

`func (o *PlanResponse) SetGeofenceOverageRate(v float32)`

SetGeofenceOverageRate sets GeofenceOverageRate field to given value.

### HasGeofenceOverageRate

`func (o *PlanResponse) HasGeofenceOverageRate() bool`

HasGeofenceOverageRate returns a boolean if a field has been set.

### GetIsFeatured

`func (o *PlanResponse) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *PlanResponse) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *PlanResponse) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.

### HasIsFeatured

`func (o *PlanResponse) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


