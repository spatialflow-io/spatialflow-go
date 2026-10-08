# WorkspaceWebhookMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**WorkspaceWebhookMetricsSummary**](WorkspaceWebhookMetricsSummary.md) |  | 
**Performance** | [**WorkspaceWebhookMetricsPerformance**](WorkspaceWebhookMetricsPerformance.md) |  | 
**Webhooks** | [**WorkspaceWebhookCounts**](WorkspaceWebhookCounts.md) |  | 
**WindowHours** | **int32** |  | 
**GeneratedAt** | **time.Time** |  | 

## Methods

### NewWorkspaceWebhookMetricsResponse

`func NewWorkspaceWebhookMetricsResponse(summary WorkspaceWebhookMetricsSummary, performance WorkspaceWebhookMetricsPerformance, webhooks WorkspaceWebhookCounts, windowHours int32, generatedAt time.Time, ) *WorkspaceWebhookMetricsResponse`

NewWorkspaceWebhookMetricsResponse instantiates a new WorkspaceWebhookMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWebhookMetricsResponseWithDefaults

`func NewWorkspaceWebhookMetricsResponseWithDefaults() *WorkspaceWebhookMetricsResponse`

NewWorkspaceWebhookMetricsResponseWithDefaults instantiates a new WorkspaceWebhookMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *WorkspaceWebhookMetricsResponse) GetSummary() WorkspaceWebhookMetricsSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *WorkspaceWebhookMetricsResponse) GetSummaryOk() (*WorkspaceWebhookMetricsSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *WorkspaceWebhookMetricsResponse) SetSummary(v WorkspaceWebhookMetricsSummary)`

SetSummary sets Summary field to given value.


### GetPerformance

`func (o *WorkspaceWebhookMetricsResponse) GetPerformance() WorkspaceWebhookMetricsPerformance`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *WorkspaceWebhookMetricsResponse) GetPerformanceOk() (*WorkspaceWebhookMetricsPerformance, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *WorkspaceWebhookMetricsResponse) SetPerformance(v WorkspaceWebhookMetricsPerformance)`

SetPerformance sets Performance field to given value.


### GetWebhooks

`func (o *WorkspaceWebhookMetricsResponse) GetWebhooks() WorkspaceWebhookCounts`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WorkspaceWebhookMetricsResponse) GetWebhooksOk() (*WorkspaceWebhookCounts, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WorkspaceWebhookMetricsResponse) SetWebhooks(v WorkspaceWebhookCounts)`

SetWebhooks sets Webhooks field to given value.


### GetWindowHours

`func (o *WorkspaceWebhookMetricsResponse) GetWindowHours() int32`

GetWindowHours returns the WindowHours field if non-nil, zero value otherwise.

### GetWindowHoursOk

`func (o *WorkspaceWebhookMetricsResponse) GetWindowHoursOk() (*int32, bool)`

GetWindowHoursOk returns a tuple with the WindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowHours

`func (o *WorkspaceWebhookMetricsResponse) SetWindowHours(v int32)`

SetWindowHours sets WindowHours field to given value.


### GetGeneratedAt

`func (o *WorkspaceWebhookMetricsResponse) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *WorkspaceWebhookMetricsResponse) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *WorkspaceWebhookMetricsResponse) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


