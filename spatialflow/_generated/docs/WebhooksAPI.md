# \WebhooksAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsWebhooksApiCreateWebhook**](WebhooksAPI.md#AppsWebhooksApiCreateWebhook) | **Post** /api/v1/webhooks/ | Create Webhook
[**AppsWebhooksApiDeleteWebhook**](WebhooksAPI.md#AppsWebhooksApiDeleteWebhook) | **Delete** /api/v1/webhooks/{webhook_id} | Delete Webhook
[**AppsWebhooksApiGetDlqStats**](WebhooksAPI.md#AppsWebhooksApiGetDlqStats) | **Get** /api/v1/webhooks/dlq/stats | Get Dlq Stats
[**AppsWebhooksApiGetWebhook**](WebhooksAPI.md#AppsWebhooksApiGetWebhook) | **Get** /api/v1/webhooks/{webhook_id} | Get Webhook
[**AppsWebhooksApiGetWebhookDeliveries**](WebhooksAPI.md#AppsWebhooksApiGetWebhookDeliveries) | **Get** /api/v1/webhooks/{webhook_id}/deliveries | Get Webhook Deliveries
[**AppsWebhooksApiGetWebhookDeliveryDetail**](WebhooksAPI.md#AppsWebhooksApiGetWebhookDeliveryDetail) | **Get** /api/v1/webhooks/{webhook_id}/deliveries/{delivery_id} | Get Webhook Delivery Detail
[**AppsWebhooksApiGetWebhookMetrics**](WebhooksAPI.md#AppsWebhooksApiGetWebhookMetrics) | **Get** /api/v1/webhooks/metrics | Get Webhook Metrics
[**AppsWebhooksApiGetWebhookSuccessTimeline**](WebhooksAPI.md#AppsWebhooksApiGetWebhookSuccessTimeline) | **Get** /api/v1/webhooks/success-timeline | Get Webhook Success Timeline
[**AppsWebhooksApiListDlqEntries**](WebhooksAPI.md#AppsWebhooksApiListDlqEntries) | **Get** /api/v1/webhooks/dlq | List Dlq Entries
[**AppsWebhooksApiListWebhooks**](WebhooksAPI.md#AppsWebhooksApiListWebhooks) | **Get** /api/v1/webhooks/ | List Webhooks
[**AppsWebhooksApiReceiveWebhook**](WebhooksAPI.md#AppsWebhooksApiReceiveWebhook) | **Post** /api/v1/webhooks/receive/{webhook_id} | Receive Webhook
[**AppsWebhooksApiRetryFromDlq**](WebhooksAPI.md#AppsWebhooksApiRetryFromDlq) | **Post** /api/v1/webhooks/dlq/{dlq_id}/retry | Retry From Dlq
[**AppsWebhooksApiRetryWebhookDelivery**](WebhooksAPI.md#AppsWebhooksApiRetryWebhookDelivery) | **Post** /api/v1/webhooks/{webhook_id}/deliveries/{delivery_id}/retry | Retry Webhook Delivery
[**AppsWebhooksApiTestWebhook**](WebhooksAPI.md#AppsWebhooksApiTestWebhook) | **Post** /api/v1/webhooks/{webhook_id}/test | Test Webhook
[**AppsWebhooksApiUpdateWebhook**](WebhooksAPI.md#AppsWebhooksApiUpdateWebhook) | **Put** /api/v1/webhooks/{webhook_id} | Update Webhook
[**AppsWebhooksApiWebhookHealthCheck**](WebhooksAPI.md#AppsWebhooksApiWebhookHealthCheck) | **Get** /api/v1/webhooks/health | Webhook Health Check



## AppsWebhooksApiCreateWebhook

> WebhookResponse AppsWebhooksApiCreateWebhook(ctx).CreateWebhookRequest(createWebhookRequest).Execute()

Create Webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	createWebhookRequest := *openapiclient.NewCreateWebhookRequest("Name_example", "Url_example") // CreateWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiCreateWebhook(context.Background()).CreateWebhookRequest(createWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiCreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiCreateWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiCreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookRequest** | [**CreateWebhookRequest**](CreateWebhookRequest.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiDeleteWebhook

> SuccessResponse AppsWebhooksApiDeleteWebhook(ctx, webhookId).Execute()

Delete Webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiDeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiDeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiDeleteWebhook`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiDeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiGetDlqStats

> map[string]interface{} AppsWebhooksApiGetDlqStats(ctx).Execute()

Get Dlq Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiGetDlqStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiGetDlqStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiGetDlqStats`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiGetDlqStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiGetDlqStatsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiGetWebhook

> WebhookResponse AppsWebhooksApiGetWebhook(ctx, webhookId).Execute()

Get Webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiGetWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiGetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiGetWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiGetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiGetWebhookDeliveries

> WebhookDeliveryListResponse AppsWebhooksApiGetWebhookDeliveries(ctx, webhookId).Limit(limit).Offset(offset).Status(status).EventType(eventType).Execute()

Get Webhook Deliveries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)
	status := "status_example" // string |  (optional)
	eventType := "eventType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiGetWebhookDeliveries(context.Background(), webhookId).Limit(limit).Offset(offset).Status(status).EventType(eventType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiGetWebhookDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiGetWebhookDeliveries`: WebhookDeliveryListResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiGetWebhookDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiGetWebhookDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **status** | **string** |  | 
 **eventType** | **string** |  | 

### Return type

[**WebhookDeliveryListResponse**](WebhookDeliveryListResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiGetWebhookDeliveryDetail

> WebhookDeliveryDetailResponse AppsWebhooksApiGetWebhookDeliveryDetail(ctx, webhookId, deliveryId).Execute()

Get Webhook Delivery Detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deliveryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiGetWebhookDeliveryDetail(context.Background(), webhookId, deliveryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiGetWebhookDeliveryDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiGetWebhookDeliveryDetail`: WebhookDeliveryDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiGetWebhookDeliveryDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiGetWebhookDeliveryDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebhookDeliveryDetailResponse**](WebhookDeliveryDetailResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiGetWebhookMetrics

> WebhookMetricsResponse AppsWebhooksApiGetWebhookMetrics(ctx).Execute()

Get Webhook Metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiGetWebhookMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiGetWebhookMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiGetWebhookMetrics`: WebhookMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiGetWebhookMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiGetWebhookMetricsRequest struct via the builder pattern


### Return type

[**WebhookMetricsResponse**](WebhookMetricsResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiGetWebhookSuccessTimeline

> map[string]interface{} AppsWebhooksApiGetWebhookSuccessTimeline(ctx).TimeRange(timeRange).StartDate(startDate).EndDate(endDate).Execute()

Get Webhook Success Timeline



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	timeRange := "timeRange_example" // string |  (optional) (default to "today")
	startDate := "startDate_example" // string |  (optional)
	endDate := "endDate_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiGetWebhookSuccessTimeline(context.Background()).TimeRange(timeRange).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiGetWebhookSuccessTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiGetWebhookSuccessTimeline`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiGetWebhookSuccessTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiGetWebhookSuccessTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeRange** | **string** |  | [default to &quot;today&quot;]
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiListDlqEntries

> map[string]interface{} AppsWebhooksApiListDlqEntries(ctx).Limit(limit).Offset(offset).Requeued(requeued).Execute()

List Dlq Entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)
	requeued := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiListDlqEntries(context.Background()).Limit(limit).Offset(offset).Requeued(requeued).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiListDlqEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiListDlqEntries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiListDlqEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiListDlqEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **requeued** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiListWebhooks

> WebhookListResponse AppsWebhooksApiListWebhooks(ctx).Limit(limit).Offset(offset).IsActive(isActive).Execute()

List Webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)
	isActive := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiListWebhooks(context.Background()).Limit(limit).Offset(offset).IsActive(isActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiListWebhooks`: WebhookListResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiListWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **isActive** | **bool** |  | 

### Return type

[**WebhookListResponse**](WebhookListResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiReceiveWebhook

> SuccessResponse AppsWebhooksApiReceiveWebhook(ctx, webhookId).Execute()

Receive Webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiReceiveWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiReceiveWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiReceiveWebhook`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiReceiveWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiReceiveWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiRetryFromDlq

> map[string]interface{} AppsWebhooksApiRetryFromDlq(ctx, dlqId).Execute()

Retry From Dlq



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	dlqId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiRetryFromDlq(context.Background(), dlqId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiRetryFromDlq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiRetryFromDlq`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiRetryFromDlq`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dlqId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiRetryFromDlqRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiRetryWebhookDelivery

> map[string]interface{} AppsWebhooksApiRetryWebhookDelivery(ctx, webhookId, deliveryId).Execute()

Retry Webhook Delivery



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deliveryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiRetryWebhookDelivery(context.Background(), webhookId, deliveryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiRetryWebhookDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiRetryWebhookDelivery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiRetryWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiRetryWebhookDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiTestWebhook

> WebhookTestResponse AppsWebhooksApiTestWebhook(ctx, webhookId).TestWebhookRequest(testWebhookRequest).Execute()

Test Webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testWebhookRequest := *openapiclient.NewTestWebhookRequest() // TestWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiTestWebhook(context.Background(), webhookId).TestWebhookRequest(testWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiTestWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiTestWebhook`: WebhookTestResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiTestWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiTestWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testWebhookRequest** | [**TestWebhookRequest**](TestWebhookRequest.md) |  | 

### Return type

[**WebhookTestResponse**](WebhookTestResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiUpdateWebhook

> WebhookResponse AppsWebhooksApiUpdateWebhook(ctx, webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()

Update Webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateWebhookRequest := *openapiclient.NewUpdateWebhookRequest() // UpdateWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiUpdateWebhook(context.Background(), webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiUpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiUpdateWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWebhookRequest** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWebhooksApiWebhookHealthCheck

> map[string]interface{} AppsWebhooksApiWebhookHealthCheck(ctx).Execute()

Webhook Health Check



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.AppsWebhooksApiWebhookHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.AppsWebhooksApiWebhookHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWebhooksApiWebhookHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.AppsWebhooksApiWebhookHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWebhooksApiWebhookHealthCheckRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

