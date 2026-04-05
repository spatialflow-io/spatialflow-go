# \GeofencesAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsGeofencesApiBulkCreateGeofences**](GeofencesAPI.md#AppsGeofencesApiBulkCreateGeofences) | **Post** /api/v1/geofences/bulk | Bulk Create Geofences
[**AppsGeofencesApiCreateGeofence**](GeofencesAPI.md#AppsGeofencesApiCreateGeofence) | **Post** /api/v1/geofences/ | Create Geofence
[**AppsGeofencesApiDeleteGeofence**](GeofencesAPI.md#AppsGeofencesApiDeleteGeofence) | **Delete** /api/v1/geofences/{geofence_id} | Delete Geofence
[**AppsGeofencesApiGeofenceHealthCheck**](GeofencesAPI.md#AppsGeofencesApiGeofenceHealthCheck) | **Get** /api/v1/geofences/health | Geofence Health Check
[**AppsGeofencesApiGetActiveGeofencesSummary**](GeofencesAPI.md#AppsGeofencesApiGetActiveGeofencesSummary) | **Get** /api/v1/geofences/active-summary | Get Active Geofences Summary
[**AppsGeofencesApiGetGeofence**](GeofencesAPI.md#AppsGeofencesApiGetGeofence) | **Get** /api/v1/geofences/{geofence_id} | Get Geofence
[**AppsGeofencesApiGetTestEventHistory**](GeofencesAPI.md#AppsGeofencesApiGetTestEventHistory) | **Get** /api/v1/geofences/{geofence_id}/test-events | Get Test Event History
[**AppsGeofencesApiGetUploadJobStatus**](GeofencesAPI.md#AppsGeofencesApiGetUploadJobStatus) | **Get** /api/v1/geofences/upload/{job_id}/status | Get Upload Job Status
[**AppsGeofencesApiListGeofenceGroups**](GeofencesAPI.md#AppsGeofencesApiListGeofenceGroups) | **Get** /api/v1/geofences/groups | List Geofence Groups
[**AppsGeofencesApiListGeofences**](GeofencesAPI.md#AppsGeofencesApiListGeofences) | **Get** /api/v1/geofences/ | List Geofences
[**AppsGeofencesApiListGroupGeofences**](GeofencesAPI.md#AppsGeofencesApiListGroupGeofences) | **Get** /api/v1/geofences/groups/{group_id}/geofences | List Group Geofences
[**AppsGeofencesApiTestGroupPoint**](GeofencesAPI.md#AppsGeofencesApiTestGroupPoint) | **Post** /api/v1/geofences/groups/{group_id}/test-point | Test Group Point
[**AppsGeofencesApiTestPoint**](GeofencesAPI.md#AppsGeofencesApiTestPoint) | **Post** /api/v1/geofences/test-point | Test Point
[**AppsGeofencesApiTriggerTestEvent**](GeofencesAPI.md#AppsGeofencesApiTriggerTestEvent) | **Post** /api/v1/geofences/{geofence_id}/test-event | Trigger Test Event
[**AppsGeofencesApiUpdateGeofence**](GeofencesAPI.md#AppsGeofencesApiUpdateGeofence) | **Put** /api/v1/geofences/{geofence_id} | Update Geofence
[**AppsGeofencesApiUpdateGeofenceGroup**](GeofencesAPI.md#AppsGeofencesApiUpdateGeofenceGroup) | **Put** /api/v1/geofences/{geofence_id}/group | Update Geofence Group
[**AppsGeofencesApiUploadGeofencesAsync**](GeofencesAPI.md#AppsGeofencesApiUploadGeofencesAsync) | **Post** /api/v1/geofences/upload | Upload Geofences Async



## AppsGeofencesApiBulkCreateGeofences

> map[string]interface{} AppsGeofencesApiBulkCreateGeofences(ctx).BulkGeofenceRequest(bulkGeofenceRequest).Execute()

Bulk Create Geofences



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
	bulkGeofenceRequest := *openapiclient.NewBulkGeofenceRequest([]openapiclient.CreateGeofenceRequest{*openapiclient.NewCreateGeofenceRequest("Name_example", map[string]interface{}{"key": interface{}(123)})}) // BulkGeofenceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiBulkCreateGeofences(context.Background()).BulkGeofenceRequest(bulkGeofenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiBulkCreateGeofences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiBulkCreateGeofences`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiBulkCreateGeofences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiBulkCreateGeofencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkGeofenceRequest** | [**BulkGeofenceRequest**](BulkGeofenceRequest.md) |  | 

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


## AppsGeofencesApiCreateGeofence

> GeofenceResponse AppsGeofencesApiCreateGeofence(ctx).CreateGeofenceRequest(createGeofenceRequest).Execute()

Create Geofence



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
	createGeofenceRequest := *openapiclient.NewCreateGeofenceRequest("Name_example", map[string]interface{}{"key": interface{}(123)}) // CreateGeofenceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiCreateGeofence(context.Background()).CreateGeofenceRequest(createGeofenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiCreateGeofence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiCreateGeofence`: GeofenceResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiCreateGeofence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiCreateGeofenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGeofenceRequest** | [**CreateGeofenceRequest**](CreateGeofenceRequest.md) |  | 

### Return type

[**GeofenceResponse**](GeofenceResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiDeleteGeofence

> map[string]interface{} AppsGeofencesApiDeleteGeofence(ctx, geofenceId).Execute()

Delete Geofence



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
	geofenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiDeleteGeofence(context.Background(), geofenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiDeleteGeofence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiDeleteGeofence`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiDeleteGeofence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiDeleteGeofenceRequest struct via the builder pattern


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


## AppsGeofencesApiGeofenceHealthCheck

> map[string]interface{} AppsGeofencesApiGeofenceHealthCheck(ctx).Execute()

Geofence Health Check



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
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiGeofenceHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiGeofenceHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiGeofenceHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiGeofenceHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiGeofenceHealthCheckRequest struct via the builder pattern


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


## AppsGeofencesApiGetActiveGeofencesSummary

> map[string]interface{} AppsGeofencesApiGetActiveGeofencesSummary(ctx).Execute()

Get Active Geofences Summary



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
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiGetActiveGeofencesSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiGetActiveGeofencesSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiGetActiveGeofencesSummary`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiGetActiveGeofencesSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiGetActiveGeofencesSummaryRequest struct via the builder pattern


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


## AppsGeofencesApiGetGeofence

> GeofenceResponse AppsGeofencesApiGetGeofence(ctx, geofenceId).Execute()

Get Geofence



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
	geofenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiGetGeofence(context.Background(), geofenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiGetGeofence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiGetGeofence`: GeofenceResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiGetGeofence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiGetGeofenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GeofenceResponse**](GeofenceResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiGetTestEventHistory

> map[string]interface{} AppsGeofencesApiGetTestEventHistory(ctx, geofenceId).Limit(limit).Offset(offset).Execute()

Get Test Event History



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
	geofenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiGetTestEventHistory(context.Background(), geofenceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiGetTestEventHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiGetTestEventHistory`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiGetTestEventHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiGetTestEventHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

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


## AppsGeofencesApiGetUploadJobStatus

> UploadJobStatus AppsGeofencesApiGetUploadJobStatus(ctx, jobId).Execute()

Get Upload Job Status



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiGetUploadJobStatus(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiGetUploadJobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiGetUploadJobStatus`: UploadJobStatus
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiGetUploadJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiGetUploadJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UploadJobStatus**](UploadJobStatus.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiListGeofenceGroups

> map[string]interface{} AppsGeofencesApiListGeofenceGroups(ctx).Execute()

List Geofence Groups



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
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiListGeofenceGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiListGeofenceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiListGeofenceGroups`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiListGeofenceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiListGeofenceGroupsRequest struct via the builder pattern


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


## AppsGeofencesApiListGeofences

> GeofenceListResponse AppsGeofencesApiListGeofences(ctx).Limit(limit).Offset(offset).ActiveOnly(activeOnly).Execute()

List Geofences



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
	activeOnly := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiListGeofences(context.Background()).Limit(limit).Offset(offset).ActiveOnly(activeOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiListGeofences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiListGeofences`: GeofenceListResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiListGeofences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiListGeofencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **activeOnly** | **bool** |  | [default to true]

### Return type

[**GeofenceListResponse**](GeofenceListResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiListGroupGeofences

> map[string]interface{} AppsGeofencesApiListGroupGeofences(ctx, groupId).Execute()

List Group Geofences



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiListGroupGeofences(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiListGroupGeofences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiListGroupGeofences`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiListGroupGeofences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiListGroupGeofencesRequest struct via the builder pattern


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


## AppsGeofencesApiTestGroupPoint

> map[string]interface{} AppsGeofencesApiTestGroupPoint(ctx, groupId).TestPointRequest(testPointRequest).Execute()

Test Group Point



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testPointRequest := *openapiclient.NewTestPointRequest() // TestPointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiTestGroupPoint(context.Background(), groupId).TestPointRequest(testPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiTestGroupPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiTestGroupPoint`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiTestGroupPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiTestGroupPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testPointRequest** | [**TestPointRequest**](TestPointRequest.md) |  | 

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


## AppsGeofencesApiTestPoint

> TestPointResponse AppsGeofencesApiTestPoint(ctx).TestPointRequest(testPointRequest).Execute()

Test Point



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
	testPointRequest := *openapiclient.NewTestPointRequest() // TestPointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiTestPoint(context.Background()).TestPointRequest(testPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiTestPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiTestPoint`: TestPointResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiTestPoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiTestPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testPointRequest** | [**TestPointRequest**](TestPointRequest.md) |  | 

### Return type

[**TestPointResponse**](TestPointResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiTriggerTestEvent

> map[string]interface{} AppsGeofencesApiTriggerTestEvent(ctx, geofenceId).EventType(eventType).TestEventRequest(testEventRequest).Execute()

Trigger Test Event



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
	geofenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	eventType := "eventType_example" // string | 
	testEventRequest := *openapiclient.NewTestEventRequest() // TestEventRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiTriggerTestEvent(context.Background(), geofenceId).EventType(eventType).TestEventRequest(testEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiTriggerTestEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiTriggerTestEvent`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiTriggerTestEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiTriggerTestEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventType** | **string** |  | 
 **testEventRequest** | [**TestEventRequest**](TestEventRequest.md) |  | 

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


## AppsGeofencesApiUpdateGeofence

> GeofenceResponse AppsGeofencesApiUpdateGeofence(ctx, geofenceId).UpdateGeofenceRequest(updateGeofenceRequest).Execute()

Update Geofence



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
	geofenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateGeofenceRequest := *openapiclient.NewUpdateGeofenceRequest() // UpdateGeofenceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiUpdateGeofence(context.Background(), geofenceId).UpdateGeofenceRequest(updateGeofenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiUpdateGeofence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiUpdateGeofence`: GeofenceResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiUpdateGeofence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiUpdateGeofenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGeofenceRequest** | [**UpdateGeofenceRequest**](UpdateGeofenceRequest.md) |  | 

### Return type

[**GeofenceResponse**](GeofenceResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiUpdateGeofenceGroup

> map[string]interface{} AppsGeofencesApiUpdateGeofenceGroup(ctx, geofenceId).GroupName(groupName).Execute()

Update Geofence Group



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
	geofenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	groupName := "groupName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiUpdateGeofenceGroup(context.Background(), geofenceId).GroupName(groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiUpdateGeofenceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiUpdateGeofenceGroup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiUpdateGeofenceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiUpdateGeofenceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupName** | **string** |  | 

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


## AppsGeofencesApiUploadGeofencesAsync

> AsyncUploadGeofencesResponse AppsGeofencesApiUploadGeofencesAsync(ctx).UploadGeofencesRequest(uploadGeofencesRequest).Execute()

Upload Geofences Async



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
	uploadGeofencesRequest := *openapiclient.NewUploadGeofencesRequest("FileId_example") // UploadGeofencesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiUploadGeofencesAsync(context.Background()).UploadGeofencesRequest(uploadGeofencesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiUploadGeofencesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiUploadGeofencesAsync`: AsyncUploadGeofencesResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiUploadGeofencesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiUploadGeofencesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadGeofencesRequest** | [**UploadGeofencesRequest**](UploadGeofencesRequest.md) |  | 

### Return type

[**AsyncUploadGeofencesResponse**](AsyncUploadGeofencesResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

