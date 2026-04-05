# \SubscriptionsAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsSubscriptionsApiCancelSubscription**](SubscriptionsAPI.md#AppsSubscriptionsApiCancelSubscription) | **Put** /api/v1/subscriptions/cancel | Cancel Subscription
[**AppsSubscriptionsApiCreateCheckoutSession**](SubscriptionsAPI.md#AppsSubscriptionsApiCreateCheckoutSession) | **Post** /api/v1/subscriptions/create-checkout-session | Create Checkout Session
[**AppsSubscriptionsApiCreatePortalSession**](SubscriptionsAPI.md#AppsSubscriptionsApiCreatePortalSession) | **Post** /api/v1/subscriptions/create-portal-session | Create Portal Session
[**AppsSubscriptionsApiGetCurrentSubscription**](SubscriptionsAPI.md#AppsSubscriptionsApiGetCurrentSubscription) | **Get** /api/v1/subscriptions/current | Get Current Subscription
[**AppsSubscriptionsApiGetSubscriptionPlans**](SubscriptionsAPI.md#AppsSubscriptionsApiGetSubscriptionPlans) | **Get** /api/v1/subscriptions/plans | Get Subscription Plans
[**AppsSubscriptionsApiGetUsageMetrics**](SubscriptionsAPI.md#AppsSubscriptionsApiGetUsageMetrics) | **Get** /api/v1/subscriptions/usage | Get Usage Metrics
[**AppsSubscriptionsApiHandleStripeWebhook**](SubscriptionsAPI.md#AppsSubscriptionsApiHandleStripeWebhook) | **Post** /api/v1/subscriptions/stripe-webhook | Handle Stripe Webhook
[**AppsSubscriptionsApiHealthCheck**](SubscriptionsAPI.md#AppsSubscriptionsApiHealthCheck) | **Get** /api/v1/subscriptions/health | Health Check
[**AppsSubscriptionsApiReactivateSubscription**](SubscriptionsAPI.md#AppsSubscriptionsApiReactivateSubscription) | **Put** /api/v1/subscriptions/reactivate | Reactivate Subscription



## AppsSubscriptionsApiCancelSubscription

> SubscriptionActionResponse AppsSubscriptionsApiCancelSubscription(ctx).Execute()

Cancel Subscription



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
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiCancelSubscription(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiCancelSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiCancelSubscription`: SubscriptionActionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiCancelSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiCancelSubscriptionRequest struct via the builder pattern


### Return type

[**SubscriptionActionResponse**](SubscriptionActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsApiCreateCheckoutSession

> CheckoutSessionResponse AppsSubscriptionsApiCreateCheckoutSession(ctx).CheckoutSessionRequest(checkoutSessionRequest).Execute()

Create Checkout Session



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
	checkoutSessionRequest := *openapiclient.NewCheckoutSessionRequest("PriceId_example", "SuccessUrl_example", "CancelUrl_example") // CheckoutSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiCreateCheckoutSession(context.Background()).CheckoutSessionRequest(checkoutSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiCreateCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiCreateCheckoutSession`: CheckoutSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiCreateCheckoutSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiCreateCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutSessionRequest** | [**CheckoutSessionRequest**](CheckoutSessionRequest.md) |  | 

### Return type

[**CheckoutSessionResponse**](CheckoutSessionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsApiCreatePortalSession

> PortalSessionResponse AppsSubscriptionsApiCreatePortalSession(ctx).PortalSessionRequest(portalSessionRequest).Execute()

Create Portal Session



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
	portalSessionRequest := *openapiclient.NewPortalSessionRequest("ReturnUrl_example") // PortalSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiCreatePortalSession(context.Background()).PortalSessionRequest(portalSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiCreatePortalSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiCreatePortalSession`: PortalSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiCreatePortalSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiCreatePortalSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portalSessionRequest** | [**PortalSessionRequest**](PortalSessionRequest.md) |  | 

### Return type

[**PortalSessionResponse**](PortalSessionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsApiGetCurrentSubscription

> SubscriptionResponse AppsSubscriptionsApiGetCurrentSubscription(ctx).Execute()

Get Current Subscription



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
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiGetCurrentSubscription(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiGetCurrentSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiGetCurrentSubscription`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiGetCurrentSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiGetCurrentSubscriptionRequest struct via the builder pattern


### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsApiGetSubscriptionPlans

> []PlanResponse AppsSubscriptionsApiGetSubscriptionPlans(ctx).Execute()

Get Subscription Plans



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
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiGetSubscriptionPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiGetSubscriptionPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiGetSubscriptionPlans`: []PlanResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiGetSubscriptionPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiGetSubscriptionPlansRequest struct via the builder pattern


### Return type

[**[]PlanResponse**](PlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsApiGetUsageMetrics

> UsageResponse AppsSubscriptionsApiGetUsageMetrics(ctx).Execute()

Get Usage Metrics



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
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiGetUsageMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiGetUsageMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiGetUsageMetrics`: UsageResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiGetUsageMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiGetUsageMetricsRequest struct via the builder pattern


### Return type

[**UsageResponse**](UsageResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsApiHandleStripeWebhook

> map[string]interface{} AppsSubscriptionsApiHandleStripeWebhook(ctx).Execute()

Handle Stripe Webhook



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
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiHandleStripeWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiHandleStripeWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiHandleStripeWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiHandleStripeWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiHandleStripeWebhookRequest struct via the builder pattern


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


## AppsSubscriptionsApiHealthCheck

> HealthResponse AppsSubscriptionsApiHealthCheck(ctx).Execute()

Health Check



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
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiHealthCheck`: HealthResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiHealthCheckRequest struct via the builder pattern


### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsApiReactivateSubscription

> SubscriptionActionResponse AppsSubscriptionsApiReactivateSubscription(ctx).Execute()

Reactivate Subscription



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
	resp, r, err := apiClient.SubscriptionsAPI.AppsSubscriptionsApiReactivateSubscription(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.AppsSubscriptionsApiReactivateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsApiReactivateSubscription`: SubscriptionActionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.AppsSubscriptionsApiReactivateSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsApiReactivateSubscriptionRequest struct via the builder pattern


### Return type

[**SubscriptionActionResponse**](SubscriptionActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

