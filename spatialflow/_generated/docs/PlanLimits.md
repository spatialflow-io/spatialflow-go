# PlanLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiCalls** | **int32** | Monthly API call limit (-1 for unlimited) | 
**Geofences** | **int32** | Maximum number of geofences (-1 for unlimited) | 
**WebhooksDelivered** | **int32** | Monthly webhook delivery limit (-1 for unlimited) | 
**TestPoints** | **int32** | Monthly test point limit (-1 for unlimited) | 
**RateLimitPerHour** | **int32** | Hourly rate limit | 
**LogRetentionDays** | Pointer to **int32** | Log retention in days (-1 for unlimited) | [optional] [default to 7]
**LocationRetentionDays** | Pointer to **int32** | Location data retention in days | [optional] [default to 90]

## Methods

### NewPlanLimits

`func NewPlanLimits(apiCalls int32, geofences int32, webhooksDelivered int32, testPoints int32, rateLimitPerHour int32, ) *PlanLimits`

NewPlanLimits instantiates a new PlanLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanLimitsWithDefaults

`func NewPlanLimitsWithDefaults() *PlanLimits`

NewPlanLimitsWithDefaults instantiates a new PlanLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiCalls

`func (o *PlanLimits) GetApiCalls() int32`

GetApiCalls returns the ApiCalls field if non-nil, zero value otherwise.

### GetApiCallsOk

`func (o *PlanLimits) GetApiCallsOk() (*int32, bool)`

GetApiCallsOk returns a tuple with the ApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCalls

`func (o *PlanLimits) SetApiCalls(v int32)`

SetApiCalls sets ApiCalls field to given value.


### GetGeofences

`func (o *PlanLimits) GetGeofences() int32`

GetGeofences returns the Geofences field if non-nil, zero value otherwise.

### GetGeofencesOk

`func (o *PlanLimits) GetGeofencesOk() (*int32, bool)`

GetGeofencesOk returns a tuple with the Geofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofences

`func (o *PlanLimits) SetGeofences(v int32)`

SetGeofences sets Geofences field to given value.


### GetWebhooksDelivered

`func (o *PlanLimits) GetWebhooksDelivered() int32`

GetWebhooksDelivered returns the WebhooksDelivered field if non-nil, zero value otherwise.

### GetWebhooksDeliveredOk

`func (o *PlanLimits) GetWebhooksDeliveredOk() (*int32, bool)`

GetWebhooksDeliveredOk returns a tuple with the WebhooksDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksDelivered

`func (o *PlanLimits) SetWebhooksDelivered(v int32)`

SetWebhooksDelivered sets WebhooksDelivered field to given value.


### GetTestPoints

`func (o *PlanLimits) GetTestPoints() int32`

GetTestPoints returns the TestPoints field if non-nil, zero value otherwise.

### GetTestPointsOk

`func (o *PlanLimits) GetTestPointsOk() (*int32, bool)`

GetTestPointsOk returns a tuple with the TestPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoints

`func (o *PlanLimits) SetTestPoints(v int32)`

SetTestPoints sets TestPoints field to given value.


### GetRateLimitPerHour

`func (o *PlanLimits) GetRateLimitPerHour() int32`

GetRateLimitPerHour returns the RateLimitPerHour field if non-nil, zero value otherwise.

### GetRateLimitPerHourOk

`func (o *PlanLimits) GetRateLimitPerHourOk() (*int32, bool)`

GetRateLimitPerHourOk returns a tuple with the RateLimitPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerHour

`func (o *PlanLimits) SetRateLimitPerHour(v int32)`

SetRateLimitPerHour sets RateLimitPerHour field to given value.


### GetLogRetentionDays

`func (o *PlanLimits) GetLogRetentionDays() int32`

GetLogRetentionDays returns the LogRetentionDays field if non-nil, zero value otherwise.

### GetLogRetentionDaysOk

`func (o *PlanLimits) GetLogRetentionDaysOk() (*int32, bool)`

GetLogRetentionDaysOk returns a tuple with the LogRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetentionDays

`func (o *PlanLimits) SetLogRetentionDays(v int32)`

SetLogRetentionDays sets LogRetentionDays field to given value.

### HasLogRetentionDays

`func (o *PlanLimits) HasLogRetentionDays() bool`

HasLogRetentionDays returns a boolean if a field has been set.

### GetLocationRetentionDays

`func (o *PlanLimits) GetLocationRetentionDays() int32`

GetLocationRetentionDays returns the LocationRetentionDays field if non-nil, zero value otherwise.

### GetLocationRetentionDaysOk

`func (o *PlanLimits) GetLocationRetentionDaysOk() (*int32, bool)`

GetLocationRetentionDaysOk returns a tuple with the LocationRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationRetentionDays

`func (o *PlanLimits) SetLocationRetentionDays(v int32)`

SetLocationRetentionDays sets LocationRetentionDays field to given value.

### HasLocationRetentionDays

`func (o *PlanLimits) HasLocationRetentionDays() bool`

HasLocationRetentionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


