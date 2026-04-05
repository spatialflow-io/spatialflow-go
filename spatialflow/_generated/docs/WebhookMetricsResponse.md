# WebhookMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | **map[string]interface{}** |  | 
**CloudwatchDashboard** | **map[string]string** |  | 

## Methods

### NewWebhookMetricsResponse

`func NewWebhookMetricsResponse(metrics map[string]interface{}, cloudwatchDashboard map[string]string, ) *WebhookMetricsResponse`

NewWebhookMetricsResponse instantiates a new WebhookMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookMetricsResponseWithDefaults

`func NewWebhookMetricsResponseWithDefaults() *WebhookMetricsResponse`

NewWebhookMetricsResponseWithDefaults instantiates a new WebhookMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *WebhookMetricsResponse) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *WebhookMetricsResponse) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *WebhookMetricsResponse) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.


### GetCloudwatchDashboard

`func (o *WebhookMetricsResponse) GetCloudwatchDashboard() map[string]string`

GetCloudwatchDashboard returns the CloudwatchDashboard field if non-nil, zero value otherwise.

### GetCloudwatchDashboardOk

`func (o *WebhookMetricsResponse) GetCloudwatchDashboardOk() (*map[string]string, bool)`

GetCloudwatchDashboardOk returns a tuple with the CloudwatchDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudwatchDashboard

`func (o *WebhookMetricsResponse) SetCloudwatchDashboard(v map[string]string)`

SetCloudwatchDashboard sets CloudwatchDashboard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


