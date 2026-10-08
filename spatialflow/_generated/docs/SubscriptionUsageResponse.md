# SubscriptionUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**PeriodStart** | **string** | ISO 8601 datetime | 
**PeriodEnd** | **string** | ISO 8601 datetime | 
**Usage** | [**UsageMetrics**](UsageMetrics.md) |  | 
**Limits** | [**PlanLimits**](PlanLimits.md) |  | 
**PlanName** | **string** |  | 
**ThrottleStatus** | Pointer to **string** | Current throttle status: ok, warning, alert, throttled, blocked | [optional] [default to "ok"]
**IsFirstBillingMonth** | Pointer to **bool** | Whether workspace is in its first billing month | [optional] [default to false]

## Methods

### NewSubscriptionUsageResponse

`func NewSubscriptionUsageResponse(userId string, periodStart string, periodEnd string, usage UsageMetrics, limits PlanLimits, planName string, ) *SubscriptionUsageResponse`

NewSubscriptionUsageResponse instantiates a new SubscriptionUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUsageResponseWithDefaults

`func NewSubscriptionUsageResponseWithDefaults() *SubscriptionUsageResponse`

NewSubscriptionUsageResponseWithDefaults instantiates a new SubscriptionUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SubscriptionUsageResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SubscriptionUsageResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SubscriptionUsageResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPeriodStart

`func (o *SubscriptionUsageResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *SubscriptionUsageResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *SubscriptionUsageResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *SubscriptionUsageResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *SubscriptionUsageResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *SubscriptionUsageResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetUsage

`func (o *SubscriptionUsageResponse) GetUsage() UsageMetrics`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SubscriptionUsageResponse) GetUsageOk() (*UsageMetrics, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SubscriptionUsageResponse) SetUsage(v UsageMetrics)`

SetUsage sets Usage field to given value.


### GetLimits

`func (o *SubscriptionUsageResponse) GetLimits() PlanLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SubscriptionUsageResponse) GetLimitsOk() (*PlanLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SubscriptionUsageResponse) SetLimits(v PlanLimits)`

SetLimits sets Limits field to given value.


### GetPlanName

`func (o *SubscriptionUsageResponse) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *SubscriptionUsageResponse) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *SubscriptionUsageResponse) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetThrottleStatus

`func (o *SubscriptionUsageResponse) GetThrottleStatus() string`

GetThrottleStatus returns the ThrottleStatus field if non-nil, zero value otherwise.

### GetThrottleStatusOk

`func (o *SubscriptionUsageResponse) GetThrottleStatusOk() (*string, bool)`

GetThrottleStatusOk returns a tuple with the ThrottleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleStatus

`func (o *SubscriptionUsageResponse) SetThrottleStatus(v string)`

SetThrottleStatus sets ThrottleStatus field to given value.

### HasThrottleStatus

`func (o *SubscriptionUsageResponse) HasThrottleStatus() bool`

HasThrottleStatus returns a boolean if a field has been set.

### GetIsFirstBillingMonth

`func (o *SubscriptionUsageResponse) GetIsFirstBillingMonth() bool`

GetIsFirstBillingMonth returns the IsFirstBillingMonth field if non-nil, zero value otherwise.

### GetIsFirstBillingMonthOk

`func (o *SubscriptionUsageResponse) GetIsFirstBillingMonthOk() (*bool, bool)`

GetIsFirstBillingMonthOk returns a tuple with the IsFirstBillingMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstBillingMonth

`func (o *SubscriptionUsageResponse) SetIsFirstBillingMonth(v bool)`

SetIsFirstBillingMonth sets IsFirstBillingMonth field to given value.

### HasIsFirstBillingMonth

`func (o *SubscriptionUsageResponse) HasIsFirstBillingMonth() bool`

HasIsFirstBillingMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


