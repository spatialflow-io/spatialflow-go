# \RouteTesterAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsDevicesApiRouteTesterCreateRouteTest**](RouteTesterAPI.md#AppsDevicesApiRouteTesterCreateRouteTest) | **Post** /api/v1/route-tester/test | Run route test
[**AppsDevicesApiRouteTesterGetRouteTest**](RouteTesterAPI.md#AppsDevicesApiRouteTesterGetRouteTest) | **Get** /api/v1/route-tester/test/{test_id} | Get route test status/results



## AppsDevicesApiRouteTesterCreateRouteTest

> RouteTestResultsOut AppsDevicesApiRouteTesterCreateRouteTest(ctx).File(file).Execute()

Run route test



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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTesterAPI.AppsDevicesApiRouteTesterCreateRouteTest(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTesterAPI.AppsDevicesApiRouteTesterCreateRouteTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiRouteTesterCreateRouteTest`: RouteTestResultsOut
	fmt.Fprintf(os.Stdout, "Response from `RouteTesterAPI.AppsDevicesApiRouteTesterCreateRouteTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiRouteTesterCreateRouteTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**RouteTestResultsOut**](RouteTestResultsOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiRouteTesterGetRouteTest

> RouteTestStatusOut AppsDevicesApiRouteTesterGetRouteTest(ctx, testId).Execute()

Get route test status/results



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
	testId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteTesterAPI.AppsDevicesApiRouteTesterGetRouteTest(context.Background(), testId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteTesterAPI.AppsDevicesApiRouteTesterGetRouteTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiRouteTesterGetRouteTest`: RouteTestStatusOut
	fmt.Fprintf(os.Stdout, "Response from `RouteTesterAPI.AppsDevicesApiRouteTesterGetRouteTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiRouteTesterGetRouteTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteTestStatusOut**](RouteTestStatusOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

