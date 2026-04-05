# \AccountAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsAccountsApiAccountHealthCheck**](AccountAPI.md#AppsAccountsApiAccountHealthCheck) | **Get** /api/v1/account/health | Account Health Check
[**AppsAccountsApiCreateApiKey**](AccountAPI.md#AppsAccountsApiCreateApiKey) | **Post** /api/v1/account/api-keys | Create Api Key
[**AppsAccountsApiCreateErasureJob**](AccountAPI.md#AppsAccountsApiCreateErasureJob) | **Post** /api/v1/account/privacy/erasure | Create Erasure Job
[**AppsAccountsApiDeleteApiKey**](AccountAPI.md#AppsAccountsApiDeleteApiKey) | **Delete** /api/v1/account/api-keys/{api_key_id} | Delete Api Key
[**AppsAccountsApiDeleteOwnAccount**](AccountAPI.md#AppsAccountsApiDeleteOwnAccount) | **Delete** /api/v1/account/me | Delete Own Account
[**AppsAccountsApiGetApiKey**](AccountAPI.md#AppsAccountsApiGetApiKey) | **Get** /api/v1/account/api-keys/{api_key_id} | Get Api Key
[**AppsAccountsApiGetApiKeys**](AccountAPI.md#AppsAccountsApiGetApiKeys) | **Get** /api/v1/account/api-keys | Get Api Keys
[**AppsAccountsApiGetDashboardMetrics**](AccountAPI.md#AppsAccountsApiGetDashboardMetrics) | **Get** /api/v1/account/dashboard/metrics | Get Dashboard Metrics
[**AppsAccountsApiGetErasureJob**](AccountAPI.md#AppsAccountsApiGetErasureJob) | **Get** /api/v1/account/privacy/erasure/{job_id} | Get Erasure Job
[**AppsAccountsApiGetNotifications**](AccountAPI.md#AppsAccountsApiGetNotifications) | **Get** /api/v1/account/notifications | Get Notifications
[**AppsAccountsApiGetUserProfile**](AccountAPI.md#AppsAccountsApiGetUserProfile) | **Get** /api/v1/account/me | Get User Profile
[**AppsAccountsApiListExpiringApiKeys**](AccountAPI.md#AppsAccountsApiListExpiringApiKeys) | **Get** /api/v1/account/api-keys/expiring | List Expiring Api Keys
[**AppsAccountsApiMarkAllNotificationsRead**](AccountAPI.md#AppsAccountsApiMarkAllNotificationsRead) | **Post** /api/v1/account/notifications/read-all | Mark All Notifications Read
[**AppsAccountsApiMarkNotificationRead**](AccountAPI.md#AppsAccountsApiMarkNotificationRead) | **Post** /api/v1/account/notifications/{notification_id}/read | Mark Notification Read
[**AppsAccountsApiRegenerateApiKey**](AccountAPI.md#AppsAccountsApiRegenerateApiKey) | **Post** /api/v1/account/api-keys/{api_key_id}/regenerate | Regenerate Api Key
[**AppsAccountsApiRotateApiKey**](AccountAPI.md#AppsAccountsApiRotateApiKey) | **Post** /api/v1/account/api-keys/{api_key_id}/rotate | Rotate Api Key
[**AppsAccountsApiSetApiKeyExpiration**](AccountAPI.md#AppsAccountsApiSetApiKeyExpiration) | **Patch** /api/v1/account/api-keys/{api_key_id}/expiration | Set Api Key Expiration
[**AppsAccountsApiUpdateApiKey**](AccountAPI.md#AppsAccountsApiUpdateApiKey) | **Put** /api/v1/account/api-keys/{api_key_id} | Update Api Key
[**AppsAccountsApiUpdateUserProfile**](AccountAPI.md#AppsAccountsApiUpdateUserProfile) | **Put** /api/v1/account/me | Update User Profile



## AppsAccountsApiAccountHealthCheck

> map[string]interface{} AppsAccountsApiAccountHealthCheck(ctx).Execute()

Account Health Check



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
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiAccountHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiAccountHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiAccountHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiAccountHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiAccountHealthCheckRequest struct via the builder pattern


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


## AppsAccountsApiCreateApiKey

> ApiKeyCreateResponse AppsAccountsApiCreateApiKey(ctx).ApiKeyCreateRequest(apiKeyCreateRequest).Execute()

Create Api Key



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
	apiKeyCreateRequest := *openapiclient.NewApiKeyCreateRequest() // ApiKeyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiCreateApiKey(context.Background()).ApiKeyCreateRequest(apiKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiCreateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiCreateApiKey`: ApiKeyCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiCreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKeyCreateRequest** | [**ApiKeyCreateRequest**](ApiKeyCreateRequest.md) |  | 

### Return type

[**ApiKeyCreateResponse**](ApiKeyCreateResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiCreateErasureJob

> PrivacyErasureResponse AppsAccountsApiCreateErasureJob(ctx).PrivacyErasureRequest(privacyErasureRequest).Execute()

Create Erasure Job



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
	privacyErasureRequest := *openapiclient.NewPrivacyErasureRequest("Scope_example") // PrivacyErasureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiCreateErasureJob(context.Background()).PrivacyErasureRequest(privacyErasureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiCreateErasureJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiCreateErasureJob`: PrivacyErasureResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiCreateErasureJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiCreateErasureJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privacyErasureRequest** | [**PrivacyErasureRequest**](PrivacyErasureRequest.md) |  | 

### Return type

[**PrivacyErasureResponse**](PrivacyErasureResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiDeleteApiKey

> map[string]interface{} AppsAccountsApiDeleteApiKey(ctx, apiKeyId).Execute()

Delete Api Key



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
	apiKeyId := "apiKeyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiDeleteApiKey(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiDeleteApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiDeleteApiKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiDeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiDeleteApiKeyRequest struct via the builder pattern


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


## AppsAccountsApiDeleteOwnAccount

> map[string]interface{} AppsAccountsApiDeleteOwnAccount(ctx).Execute()

Delete Own Account



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
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiDeleteOwnAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiDeleteOwnAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiDeleteOwnAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiDeleteOwnAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiDeleteOwnAccountRequest struct via the builder pattern


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


## AppsAccountsApiGetApiKey

> ApiKeyResponse AppsAccountsApiGetApiKey(ctx, apiKeyId).Execute()

Get Api Key



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
	apiKeyId := "apiKeyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiGetApiKey(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiGetApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiGetApiKey`: ApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiGetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiGetApiKeys

> []ApiKeyResponse AppsAccountsApiGetApiKeys(ctx).Execute()

Get Api Keys



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
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiGetApiKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiGetApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiGetApiKeys`: []ApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiGetApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiGetApiKeysRequest struct via the builder pattern


### Return type

[**[]ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiGetDashboardMetrics

> DashboardMetricsResponse AppsAccountsApiGetDashboardMetrics(ctx).TimeRange(timeRange).StartDate(startDate).EndDate(endDate).Execute()

Get Dashboard Metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	timeRange := "timeRange_example" // string | Time range: today, 7d, 30d, or custom (optional) (default to "today")
	startDate := time.Now() // time.Time | Custom start date (ISO 8601, required if time_range=custom) (optional)
	endDate := time.Now() // time.Time | Custom end date (ISO 8601, required if time_range=custom) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiGetDashboardMetrics(context.Background()).TimeRange(timeRange).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiGetDashboardMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiGetDashboardMetrics`: DashboardMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiGetDashboardMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiGetDashboardMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeRange** | **string** | Time range: today, 7d, 30d, or custom | [default to &quot;today&quot;]
 **startDate** | **time.Time** | Custom start date (ISO 8601, required if time_range&#x3D;custom) | 
 **endDate** | **time.Time** | Custom end date (ISO 8601, required if time_range&#x3D;custom) | 

### Return type

[**DashboardMetricsResponse**](DashboardMetricsResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiGetErasureJob

> PrivacyErasureResponse AppsAccountsApiGetErasureJob(ctx, jobId).Execute()

Get Erasure Job



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiGetErasureJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiGetErasureJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiGetErasureJob`: PrivacyErasureResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiGetErasureJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiGetErasureJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivacyErasureResponse**](PrivacyErasureResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiGetNotifications

> map[string]interface{} AppsAccountsApiGetNotifications(ctx).UnreadOnly(unreadOnly).Execute()

Get Notifications



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
	unreadOnly := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiGetNotifications(context.Background()).UnreadOnly(unreadOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiGetNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiGetNotifications`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiGetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unreadOnly** | **bool** |  | [default to false]

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


## AppsAccountsApiGetUserProfile

> UserProfileResponse AppsAccountsApiGetUserProfile(ctx).Execute()

Get User Profile



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
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiGetUserProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiGetUserProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiGetUserProfile`: UserProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiGetUserProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiGetUserProfileRequest struct via the builder pattern


### Return type

[**UserProfileResponse**](UserProfileResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiListExpiringApiKeys

> []map[string]interface{} AppsAccountsApiListExpiringApiKeys(ctx).Days(days).Execute()

List Expiring Api Keys



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
	days := int32(56) // int32 |  (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiListExpiringApiKeys(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiListExpiringApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiListExpiringApiKeys`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiListExpiringApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiListExpiringApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** |  | [default to 30]

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiMarkAllNotificationsRead

> map[string]interface{} AppsAccountsApiMarkAllNotificationsRead(ctx).Execute()

Mark All Notifications Read



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
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiMarkAllNotificationsRead(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiMarkAllNotificationsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiMarkAllNotificationsRead`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiMarkAllNotificationsRead`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiMarkAllNotificationsReadRequest struct via the builder pattern


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


## AppsAccountsApiMarkNotificationRead

> map[string]interface{} AppsAccountsApiMarkNotificationRead(ctx, notificationId).Execute()

Mark Notification Read



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
	notificationId := "notificationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiMarkNotificationRead(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiMarkNotificationRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiMarkNotificationRead`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiMarkNotificationRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiMarkNotificationReadRequest struct via the builder pattern


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


## AppsAccountsApiRegenerateApiKey

> ApiKeyCreateResponse AppsAccountsApiRegenerateApiKey(ctx, apiKeyId).Execute()

Regenerate Api Key



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
	apiKeyId := "apiKeyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiRegenerateApiKey(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiRegenerateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiRegenerateApiKey`: ApiKeyCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiRegenerateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiRegenerateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKeyCreateResponse**](ApiKeyCreateResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiRotateApiKey

> ApiKeyCreateResponse AppsAccountsApiRotateApiKey(ctx, apiKeyId).GracePeriodDays(gracePeriodDays).Execute()

Rotate Api Key



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
	apiKeyId := "apiKeyId_example" // string | 
	gracePeriodDays := int32(56) // int32 |  (optional) (default to 7)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiRotateApiKey(context.Background(), apiKeyId).GracePeriodDays(gracePeriodDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiRotateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiRotateApiKey`: ApiKeyCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiRotateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiRotateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gracePeriodDays** | **int32** |  | [default to 7]

### Return type

[**ApiKeyCreateResponse**](ApiKeyCreateResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiSetApiKeyExpiration

> map[string]string AppsAccountsApiSetApiKeyExpiration(ctx, apiKeyId).ExpiresAt(expiresAt).Execute()

Set Api Key Expiration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	apiKeyId := "apiKeyId_example" // string | 
	expiresAt := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiSetApiKeyExpiration(context.Background(), apiKeyId).ExpiresAt(expiresAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiSetApiKeyExpiration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiSetApiKeyExpiration`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiSetApiKeyExpiration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiSetApiKeyExpirationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expiresAt** | **time.Time** |  | 

### Return type

**map[string]string**

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiUpdateApiKey

> map[string]interface{} AppsAccountsApiUpdateApiKey(ctx, apiKeyId).ApiKeyUpdateRequest(apiKeyUpdateRequest).Execute()

Update Api Key



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
	apiKeyId := "apiKeyId_example" // string | 
	apiKeyUpdateRequest := *openapiclient.NewApiKeyUpdateRequest() // ApiKeyUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiUpdateApiKey(context.Background(), apiKeyId).ApiKeyUpdateRequest(apiKeyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiUpdateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiUpdateApiKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiUpdateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiUpdateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKeyUpdateRequest** | [**ApiKeyUpdateRequest**](ApiKeyUpdateRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAccountsApiUpdateUserProfile

> map[string]interface{} AppsAccountsApiUpdateUserProfile(ctx).UpdateProfileRequest(updateProfileRequest).Execute()

Update User Profile



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
	updateProfileRequest := *openapiclient.NewUpdateProfileRequest() // UpdateProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AppsAccountsApiUpdateUserProfile(context.Background()).UpdateProfileRequest(updateProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AppsAccountsApiUpdateUserProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAccountsApiUpdateUserProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AppsAccountsApiUpdateUserProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAccountsApiUpdateUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProfileRequest** | [**UpdateProfileRequest**](UpdateProfileRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

