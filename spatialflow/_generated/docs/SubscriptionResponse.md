# SubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**Plan** | [**PlanResponse**](PlanResponse.md) |  | 
**Status** | **string** | Subscription status (active/canceled/past_due) | 
**CurrentPeriodStart** | Pointer to **NullableString** |  | [optional] 
**CurrentPeriodEnd** | Pointer to **NullableString** |  | [optional] 
**CancelAtPeriodEnd** | Pointer to **bool** |  | [optional] [default to false]
**CreatedAt** | **string** | ISO 8601 datetime | 
**UpdatedAt** | **string** | ISO 8601 datetime | 

## Methods

### NewSubscriptionResponse

`func NewSubscriptionResponse(userId string, plan PlanResponse, status string, createdAt string, updatedAt string, ) *SubscriptionResponse`

NewSubscriptionResponse instantiates a new SubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionResponseWithDefaults

`func NewSubscriptionResponseWithDefaults() *SubscriptionResponse`

NewSubscriptionResponseWithDefaults instantiates a new SubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SubscriptionResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SubscriptionResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SubscriptionResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPlan

`func (o *SubscriptionResponse) GetPlan() PlanResponse`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionResponse) GetPlanOk() (*PlanResponse, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionResponse) SetPlan(v PlanResponse)`

SetPlan sets Plan field to given value.


### GetStatus

`func (o *SubscriptionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrentPeriodStart

`func (o *SubscriptionResponse) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *SubscriptionResponse) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *SubscriptionResponse) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *SubscriptionResponse) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### SetCurrentPeriodStartNil

`func (o *SubscriptionResponse) SetCurrentPeriodStartNil(b bool)`

 SetCurrentPeriodStartNil sets the value for CurrentPeriodStart to be an explicit nil

### UnsetCurrentPeriodStart
`func (o *SubscriptionResponse) UnsetCurrentPeriodStart()`

UnsetCurrentPeriodStart ensures that no value is present for CurrentPeriodStart, not even an explicit nil
### GetCurrentPeriodEnd

`func (o *SubscriptionResponse) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *SubscriptionResponse) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *SubscriptionResponse) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *SubscriptionResponse) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### SetCurrentPeriodEndNil

`func (o *SubscriptionResponse) SetCurrentPeriodEndNil(b bool)`

 SetCurrentPeriodEndNil sets the value for CurrentPeriodEnd to be an explicit nil

### UnsetCurrentPeriodEnd
`func (o *SubscriptionResponse) UnsetCurrentPeriodEnd()`

UnsetCurrentPeriodEnd ensures that no value is present for CurrentPeriodEnd, not even an explicit nil
### GetCancelAtPeriodEnd

`func (o *SubscriptionResponse) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *SubscriptionResponse) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *SubscriptionResponse) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *SubscriptionResponse) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriptionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


