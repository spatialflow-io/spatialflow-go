# \SystemAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsCoreHealthApiHealthCheck**](SystemAPI.md#AppsCoreHealthApiHealthCheck) | **Get** /api/v1/health | Api Health Check
[**AppsCoreHealthApiHealthCheckReady**](SystemAPI.md#AppsCoreHealthApiHealthCheckReady) | **Get** /api/v1/health/ready | Api Health Check Ready



## AppsCoreHealthApiHealthCheck

> map[string]interface{} AppsCoreHealthApiHealthCheck(ctx).Execute()

Api Health Check



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
	resp, r, err := apiClient.SystemAPI.AppsCoreHealthApiHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.AppsCoreHealthApiHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsCoreHealthApiHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.AppsCoreHealthApiHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsCoreHealthApiHealthCheckRequest struct via the builder pattern


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


## AppsCoreHealthApiHealthCheckReady

> map[string]interface{} AppsCoreHealthApiHealthCheckReady(ctx).Execute()

Api Health Check Ready



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
	resp, r, err := apiClient.SystemAPI.AppsCoreHealthApiHealthCheckReady(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.AppsCoreHealthApiHealthCheckReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsCoreHealthApiHealthCheckReady`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.AppsCoreHealthApiHealthCheckReady`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsCoreHealthApiHealthCheckReadyRequest struct via the builder pattern


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

