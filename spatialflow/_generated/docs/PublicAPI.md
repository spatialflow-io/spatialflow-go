# \PublicAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsPublicApiContactSales**](PublicAPI.md#AppsPublicApiContactSales) | **Post** /api/v1/public/contact | Contact Sales
[**AppsPublicApiGetApiDocs**](PublicAPI.md#AppsPublicApiGetApiDocs) | **Get** /api/v1/public/docs | Get Api Docs
[**AppsPublicApiGetWebsocketRoutes**](PublicAPI.md#AppsPublicApiGetWebsocketRoutes) | **Get** /api/v1/public/websocket-routes | Get Websocket Routes
[**AppsPublicApiHealthCheck**](PublicAPI.md#AppsPublicApiHealthCheck) | **Get** /api/v1/public/health | Health Check
[**AppsPublicApiRuntimeConfig**](PublicAPI.md#AppsPublicApiRuntimeConfig) | **Get** /api/v1/public/runtime-config | Runtime Config
[**AppsPublicApiSignup**](PublicAPI.md#AppsPublicApiSignup) | **Post** /api/v1/public/signup | Signup
[**AppsPublicApiStatus**](PublicAPI.md#AppsPublicApiStatus) | **Get** /api/v1/public/status | Status
[**AppsPublicApiSwaggerUi**](PublicAPI.md#AppsPublicApiSwaggerUi) | **Get** /api/v1/public/docs/ui | Swagger Ui



## AppsPublicApiContactSales

> map[string]interface{} AppsPublicApiContactSales(ctx).ContactRequest(contactRequest).Execute()

Contact Sales



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
	contactRequest := *openapiclient.NewContactRequest("Name_example", "Email_example", "UseCase_example", "Message_example") // ContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.AppsPublicApiContactSales(context.Background()).ContactRequest(contactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.AppsPublicApiContactSales``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicApiContactSales`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.AppsPublicApiContactSales`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicApiContactSalesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactRequest** | [**ContactRequest**](ContactRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPublicApiGetApiDocs

> map[string]interface{} AppsPublicApiGetApiDocs(ctx).Execute()

Get Api Docs



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
	resp, r, err := apiClient.PublicAPI.AppsPublicApiGetApiDocs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.AppsPublicApiGetApiDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicApiGetApiDocs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.AppsPublicApiGetApiDocs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicApiGetApiDocsRequest struct via the builder pattern


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


## AppsPublicApiGetWebsocketRoutes

> map[string]interface{} AppsPublicApiGetWebsocketRoutes(ctx).Execute()

Get Websocket Routes



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
	resp, r, err := apiClient.PublicAPI.AppsPublicApiGetWebsocketRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.AppsPublicApiGetWebsocketRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicApiGetWebsocketRoutes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.AppsPublicApiGetWebsocketRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicApiGetWebsocketRoutesRequest struct via the builder pattern


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


## AppsPublicApiHealthCheck

> map[string]interface{} AppsPublicApiHealthCheck(ctx).Execute()

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
	resp, r, err := apiClient.PublicAPI.AppsPublicApiHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.AppsPublicApiHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicApiHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.AppsPublicApiHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicApiHealthCheckRequest struct via the builder pattern


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


## AppsPublicApiRuntimeConfig

> map[string]interface{} AppsPublicApiRuntimeConfig(ctx).Execute()

Runtime Config



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
	resp, r, err := apiClient.PublicAPI.AppsPublicApiRuntimeConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.AppsPublicApiRuntimeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicApiRuntimeConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.AppsPublicApiRuntimeConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicApiRuntimeConfigRequest struct via the builder pattern


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


## AppsPublicApiSignup

> map[string]interface{} AppsPublicApiSignup(ctx).SignupRequest(signupRequest).Execute()

Signup



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
	signupRequest := *openapiclient.NewSignupRequest("Email_example", "Password_example") // SignupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.AppsPublicApiSignup(context.Background()).SignupRequest(signupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.AppsPublicApiSignup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicApiSignup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.AppsPublicApiSignup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicApiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupRequest** | [**SignupRequest**](SignupRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPublicApiStatus

> map[string]interface{} AppsPublicApiStatus(ctx).Execute()

Status



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
	resp, r, err := apiClient.PublicAPI.AppsPublicApiStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.AppsPublicApiStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicApiStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.AppsPublicApiStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicApiStatusRequest struct via the builder pattern


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


## AppsPublicApiSwaggerUi

> map[string]interface{} AppsPublicApiSwaggerUi(ctx).Execute()

Swagger Ui



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
	resp, r, err := apiClient.PublicAPI.AppsPublicApiSwaggerUi(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.AppsPublicApiSwaggerUi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicApiSwaggerUi`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.AppsPublicApiSwaggerUi`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicApiSwaggerUiRequest struct via the builder pattern


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

