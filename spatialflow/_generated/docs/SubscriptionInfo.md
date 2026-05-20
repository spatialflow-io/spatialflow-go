# SubscriptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanName** | Pointer to **string** |  | [optional] [default to "free"]
**Status** | Pointer to **string** |  | [optional] [default to "none"]
**BillingPeriod** | Pointer to **NullableString** |  | [optional] 
**CurrentPeriodEnd** | Pointer to **NullableString** |  | [optional] 
**LastPaymentAmount** | Pointer to **NullableFloat32** |  | [optional] 
**LastPaymentDate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubscriptionInfo

`func NewSubscriptionInfo() *SubscriptionInfo`

NewSubscriptionInfo instantiates a new SubscriptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionInfoWithDefaults

`func NewSubscriptionInfoWithDefaults() *SubscriptionInfo`

NewSubscriptionInfoWithDefaults instantiates a new SubscriptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanName

`func (o *SubscriptionInfo) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *SubscriptionInfo) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *SubscriptionInfo) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *SubscriptionInfo) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *SubscriptionInfo) GetBillingPeriod() string`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *SubscriptionInfo) GetBillingPeriodOk() (*string, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *SubscriptionInfo) SetBillingPeriod(v string)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *SubscriptionInfo) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### SetBillingPeriodNil

`func (o *SubscriptionInfo) SetBillingPeriodNil(b bool)`

 SetBillingPeriodNil sets the value for BillingPeriod to be an explicit nil

### UnsetBillingPeriod
`func (o *SubscriptionInfo) UnsetBillingPeriod()`

UnsetBillingPeriod ensures that no value is present for BillingPeriod, not even an explicit nil
### GetCurrentPeriodEnd

`func (o *SubscriptionInfo) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *SubscriptionInfo) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *SubscriptionInfo) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *SubscriptionInfo) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### SetCurrentPeriodEndNil

`func (o *SubscriptionInfo) SetCurrentPeriodEndNil(b bool)`

 SetCurrentPeriodEndNil sets the value for CurrentPeriodEnd to be an explicit nil

### UnsetCurrentPeriodEnd
`func (o *SubscriptionInfo) UnsetCurrentPeriodEnd()`

UnsetCurrentPeriodEnd ensures that no value is present for CurrentPeriodEnd, not even an explicit nil
### GetLastPaymentAmount

`func (o *SubscriptionInfo) GetLastPaymentAmount() float32`

GetLastPaymentAmount returns the LastPaymentAmount field if non-nil, zero value otherwise.

### GetLastPaymentAmountOk

`func (o *SubscriptionInfo) GetLastPaymentAmountOk() (*float32, bool)`

GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAmount

`func (o *SubscriptionInfo) SetLastPaymentAmount(v float32)`

SetLastPaymentAmount sets LastPaymentAmount field to given value.

### HasLastPaymentAmount

`func (o *SubscriptionInfo) HasLastPaymentAmount() bool`

HasLastPaymentAmount returns a boolean if a field has been set.

### SetLastPaymentAmountNil

`func (o *SubscriptionInfo) SetLastPaymentAmountNil(b bool)`

 SetLastPaymentAmountNil sets the value for LastPaymentAmount to be an explicit nil

### UnsetLastPaymentAmount
`func (o *SubscriptionInfo) UnsetLastPaymentAmount()`

UnsetLastPaymentAmount ensures that no value is present for LastPaymentAmount, not even an explicit nil
### GetLastPaymentDate

`func (o *SubscriptionInfo) GetLastPaymentDate() string`

GetLastPaymentDate returns the LastPaymentDate field if non-nil, zero value otherwise.

### GetLastPaymentDateOk

`func (o *SubscriptionInfo) GetLastPaymentDateOk() (*string, bool)`

GetLastPaymentDateOk returns a tuple with the LastPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentDate

`func (o *SubscriptionInfo) SetLastPaymentDate(v string)`

SetLastPaymentDate sets LastPaymentDate field to given value.

### HasLastPaymentDate

`func (o *SubscriptionInfo) HasLastPaymentDate() bool`

HasLastPaymentDate returns a boolean if a field has been set.

### SetLastPaymentDateNil

`func (o *SubscriptionInfo) SetLastPaymentDateNil(b bool)`

 SetLastPaymentDateNil sets the value for LastPaymentDate to be an explicit nil

### UnsetLastPaymentDate
`func (o *SubscriptionInfo) UnsetLastPaymentDate()`

UnsetLastPaymentDate ensures that no value is present for LastPaymentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


