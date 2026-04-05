# \EmailAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsEmailApiGetEmailHistory**](EmailAPI.md#AppsEmailApiGetEmailHistory) | **Get** /api/v1/email/history | Get Email History
[**AppsEmailApiGetEmailStatus**](EmailAPI.md#AppsEmailApiGetEmailStatus) | **Get** /api/v1/email/status/{email_id} | Get Email Status
[**AppsEmailApiGetQueueStats**](EmailAPI.md#AppsEmailApiGetQueueStats) | **Get** /api/v1/email/queue-stats | Get Queue Stats
[**AppsEmailApiHealthCheck**](EmailAPI.md#AppsEmailApiHealthCheck) | **Get** /api/v1/email/health | Health Check
[**AppsEmailApiPreviewEmailTemplate**](EmailAPI.md#AppsEmailApiPreviewEmailTemplate) | **Get** /api/v1/email/preview/{template_name} | Preview Email Template
[**AppsEmailApiSendEmail**](EmailAPI.md#AppsEmailApiSendEmail) | **Post** /api/v1/email/send | Send Email
[**AppsEmailApiSendTestEmail**](EmailAPI.md#AppsEmailApiSendTestEmail) | **Post** /api/v1/email/test | Send Test Email



## AppsEmailApiGetEmailHistory

> map[string]interface{} AppsEmailApiGetEmailHistory(ctx).Limit(limit).Offset(offset).Execute()

Get Email History



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.AppsEmailApiGetEmailHistory(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.AppsEmailApiGetEmailHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsEmailApiGetEmailHistory`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.AppsEmailApiGetEmailHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailApiGetEmailHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

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


## AppsEmailApiGetEmailStatus

> EmailStatusResponse AppsEmailApiGetEmailStatus(ctx, emailId).Execute()

Get Email Status



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
	emailId := "emailId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.AppsEmailApiGetEmailStatus(context.Background(), emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.AppsEmailApiGetEmailStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsEmailApiGetEmailStatus`: EmailStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.AppsEmailApiGetEmailStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailApiGetEmailStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailStatusResponse**](EmailStatusResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsEmailApiGetQueueStats

> EmailQueueStats AppsEmailApiGetQueueStats(ctx).Execute()

Get Queue Stats



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
	resp, r, err := apiClient.EmailAPI.AppsEmailApiGetQueueStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.AppsEmailApiGetQueueStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsEmailApiGetQueueStats`: EmailQueueStats
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.AppsEmailApiGetQueueStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailApiGetQueueStatsRequest struct via the builder pattern


### Return type

[**EmailQueueStats**](EmailQueueStats.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsEmailApiHealthCheck

> EmailHealthResponse AppsEmailApiHealthCheck(ctx).Execute()

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
	resp, r, err := apiClient.EmailAPI.AppsEmailApiHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.AppsEmailApiHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsEmailApiHealthCheck`: EmailHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.AppsEmailApiHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailApiHealthCheckRequest struct via the builder pattern


### Return type

[**EmailHealthResponse**](EmailHealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsEmailApiPreviewEmailTemplate

> AppsEmailApiPreviewEmailTemplate(ctx, templateName).Format(format).Execute()

Preview Email Template



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
	templateName := "templateName_example" // string | 
	format := "format_example" // string |  (optional) (default to "html")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.AppsEmailApiPreviewEmailTemplate(context.Background(), templateName).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.AppsEmailApiPreviewEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailApiPreviewEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | [default to &quot;html&quot;]

### Return type

 (empty response body)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsEmailApiSendEmail

> map[string]interface{} AppsEmailApiSendEmail(ctx).SendEmailRequest(sendEmailRequest).Execute()

Send Email



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
	sendEmailRequest := *openapiclient.NewSendEmailRequest("ToEmail_example", "Template_example", map[string]interface{}{"key": interface{}(123)}) // SendEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.AppsEmailApiSendEmail(context.Background()).SendEmailRequest(sendEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.AppsEmailApiSendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsEmailApiSendEmail`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.AppsEmailApiSendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailApiSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendEmailRequest** | [**SendEmailRequest**](SendEmailRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsEmailApiSendTestEmail

> map[string]interface{} AppsEmailApiSendTestEmail(ctx).Execute()

Send Test Email



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
	resp, r, err := apiClient.EmailAPI.AppsEmailApiSendTestEmail(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.AppsEmailApiSendTestEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsEmailApiSendTestEmail`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.AppsEmailApiSendTestEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailApiSendTestEmailRequest struct via the builder pattern


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

