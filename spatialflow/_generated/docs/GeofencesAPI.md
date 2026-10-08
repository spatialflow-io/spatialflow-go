# \GeofencesAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsGeofencesApiArchiveGeofence**](GeofencesAPI.md#AppsGeofencesApiArchiveGeofence) | **Post** /api/v1/geofences/{geofence_id}/archive | Archive Geofence
[**AppsGeofencesApiBulkCreateGeofences**](GeofencesAPI.md#AppsGeofencesApiBulkCreateGeofences) | **Post** /api/v1/geofences/bulk | Bulk Create Geofences
[**AppsGeofencesApiBulkPreviewGeofences**](GeofencesAPI.md#AppsGeofencesApiBulkPreviewGeofences) | **Post** /api/v1/geofences/preview | Bulk Preview Geofences
[**AppsGeofencesApiCreateGeofence**](GeofencesAPI.md#AppsGeofencesApiCreateGeofence) | **Post** /api/v1/geofences/ | Create Geofence
[**AppsGeofencesApiDeleteGeofence**](GeofencesAPI.md#AppsGeofencesApiDeleteGeofence) | **Delete** /api/v1/geofences/{geofence_id} | Delete Geofence
[**AppsGeofencesApiGeocodeAutocomplete**](GeofencesAPI.md#AppsGeofencesApiGeocodeAutocomplete) | **Get** /api/v1/geofences/geocode/autocomplete | Geocode Autocomplete
[**AppsGeofencesApiGeocodePlace**](GeofencesAPI.md#AppsGeofencesApiGeocodePlace) | **Get** /api/v1/geofences/geocode/place | Geocode Place
[**AppsGeofencesApiGeofenceHealthCheck**](GeofencesAPI.md#AppsGeofencesApiGeofenceHealthCheck) | **Get** /api/v1/geofences/health | Geofence Health Check
[**AppsGeofencesApiGetActiveGeofencesSummary**](GeofencesAPI.md#AppsGeofencesApiGetActiveGeofencesSummary) | **Get** /api/v1/geofences/active-summary | Get Active Geofences Summary
[**AppsGeofencesApiGetGeofence**](GeofencesAPI.md#AppsGeofencesApiGetGeofence) | **Get** /api/v1/geofences/{geofence_id} | Get Geofence
[**AppsGeofencesApiGetTestEventHistory**](GeofencesAPI.md#AppsGeofencesApiGetTestEventHistory) | **Get** /api/v1/geofences/{geofence_id}/test-events | Get Test Event History
[**AppsGeofencesApiGetUploadJobStatus**](GeofencesAPI.md#AppsGeofencesApiGetUploadJobStatus) | **Get** /api/v1/geofences/upload/{job_id}/status | Get Upload Job Status
[**AppsGeofencesApiListGeofenceGroups**](GeofencesAPI.md#AppsGeofencesApiListGeofenceGroups) | **Get** /api/v1/geofences/groups | List Geofence Groups
[**AppsGeofencesApiListGeofenceWorkflows**](GeofencesAPI.md#AppsGeofencesApiListGeofenceWorkflows) | **Get** /api/v1/geofences/{geofence_id}/workflows | List Geofence Workflows
[**AppsGeofencesApiListGeofences**](GeofencesAPI.md#AppsGeofencesApiListGeofences) | **Get** /api/v1/geofences/ | List Geofences
[**AppsGeofencesApiListGroupGeofences**](GeofencesAPI.md#AppsGeofencesApiListGroupGeofences) | **Get** /api/v1/geofences/groups/{group_id}/geofences | List Group Geofences
[**AppsGeofencesApiListWorkspaceTags**](GeofencesAPI.md#AppsGeofencesApiListWorkspaceTags) | **Get** /api/v1/geofences/tags | List Workspace Tags
[**AppsGeofencesApiTestGroupPoint**](GeofencesAPI.md#AppsGeofencesApiTestGroupPoint) | **Post** /api/v1/geofences/groups/{group_id}/test-point | Test Group Point
[**AppsGeofencesApiTestPoint**](GeofencesAPI.md#AppsGeofencesApiTestPoint) | **Post** /api/v1/geofences/test-point | Test Point
[**AppsGeofencesApiTriggerTestEvent**](GeofencesAPI.md#AppsGeofencesApiTriggerTestEvent) | **Post** /api/v1/geofences/{geofence_id}/test-event | Trigger Test Event
[**AppsGeofencesApiUnarchiveGeofence**](GeofencesAPI.md#AppsGeofencesApiUnarchiveGeofence) | **Post** /api/v1/geofences/{geofence_id}/unarchive | Unarchive Geofence
[**AppsGeofencesApiUpdateGeofence**](GeofencesAPI.md#AppsGeofencesApiUpdateGeofence) | **Put** /api/v1/geofences/{geofence_id} | Update Geofence
[**AppsGeofencesApiUpdateGeofenceGroup**](GeofencesAPI.md#AppsGeofencesApiUpdateGeofenceGroup) | **Put** /api/v1/geofences/{geofence_id}/group | Update Geofence Group
[**AppsGeofencesApiUploadGeofencesAsync**](GeofencesAPI.md#AppsGeofencesApiUploadGeofencesAsync) | **Post** /api/v1/geofences/upload | Upload Geofences Async



## AppsGeofencesApiArchiveGeofence

> GeofenceResponse AppsGeofencesApiArchiveGeofence(ctx, geofenceId).Execute()

Archive Geofence



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
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiArchiveGeofence(context.Background(), geofenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiArchiveGeofence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiArchiveGeofence`: GeofenceResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiArchiveGeofence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiArchiveGeofenceRequest struct via the builder pattern


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


## AppsGeofencesApiBulkCreateGeofences

> BulkCreateResponse AppsGeofencesApiBulkCreateGeofences(ctx).BulkCreateRequest(bulkCreateRequest).Execute()

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
	bulkCreateRequest := *openapiclient.NewBulkCreateRequest([]openapiclient.BulkItemCommit{*openapiclient.NewBulkItemCommit(int32(123), "Address_example", "Status_example")}) // BulkCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiBulkCreateGeofences(context.Background()).BulkCreateRequest(bulkCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiBulkCreateGeofences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiBulkCreateGeofences`: BulkCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiBulkCreateGeofences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiBulkCreateGeofencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkCreateRequest** | [**BulkCreateRequest**](BulkCreateRequest.md) |  | 

### Return type

[**BulkCreateResponse**](BulkCreateResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiBulkPreviewGeofences

> BulkPreviewResponse AppsGeofencesApiBulkPreviewGeofences(ctx).BulkPreviewRequest(bulkPreviewRequest).Execute()

Bulk Preview Geofences



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
	bulkPreviewRequest := *openapiclient.NewBulkPreviewRequest([]openapiclient.BulkItemInput{*openapiclient.NewBulkItemInput("Address_example")}) // BulkPreviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiBulkPreviewGeofences(context.Background()).BulkPreviewRequest(bulkPreviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiBulkPreviewGeofences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiBulkPreviewGeofences`: BulkPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiBulkPreviewGeofences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiBulkPreviewGeofencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkPreviewRequest** | [**BulkPreviewRequest**](BulkPreviewRequest.md) |  | 

### Return type

[**BulkPreviewResponse**](BulkPreviewResponse.md)

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
	createGeofenceRequest := *openapiclient.NewCreateGeofenceRequest("Name_example") // CreateGeofenceRequest | 

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

> AppsGeofencesApiDeleteGeofence(ctx, geofenceId).Execute()

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
	r, err := apiClient.GeofencesAPI.AppsGeofencesApiDeleteGeofence(context.Background(), geofenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiDeleteGeofence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiGeocodeAutocomplete

> AutocompleteResponse AppsGeofencesApiGeocodeAutocomplete(ctx).Query(query).MaxResults(maxResults).Execute()

Geocode Autocomplete



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
	query := "query_example" // string | 
	maxResults := int32(56) // int32 |  (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiGeocodeAutocomplete(context.Background()).Query(query).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiGeocodeAutocomplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiGeocodeAutocomplete`: AutocompleteResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiGeocodeAutocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiGeocodeAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **maxResults** | **int32** |  | [default to 5]

### Return type

[**AutocompleteResponse**](AutocompleteResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiGeocodePlace

> GeocodePlaceResponse AppsGeofencesApiGeocodePlace(ctx).PlaceId(placeId).Execute()

Geocode Place



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
	placeId := "placeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiGeocodePlace(context.Background()).PlaceId(placeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiGeocodePlace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiGeocodePlace`: GeocodePlaceResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiGeocodePlace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiGeocodePlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placeId** | **string** |  | 

### Return type

[**GeocodePlaceResponse**](GeocodePlaceResponse.md)

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

> ActiveGeofenceSummaryOut AppsGeofencesApiGetActiveGeofencesSummary(ctx).Execute()

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
	// response from `AppsGeofencesApiGetActiveGeofencesSummary`: ActiveGeofenceSummaryOut
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiGetActiveGeofencesSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiGetActiveGeofencesSummaryRequest struct via the builder pattern


### Return type

[**ActiveGeofenceSummaryOut**](ActiveGeofenceSummaryOut.md)

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

> TestEventHistoryOut AppsGeofencesApiGetTestEventHistory(ctx, geofenceId).Limit(limit).Offset(offset).Execute()

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
	// response from `AppsGeofencesApiGetTestEventHistory`: TestEventHistoryOut
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

[**TestEventHistoryOut**](TestEventHistoryOut.md)

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

> GeofenceGroupsOut AppsGeofencesApiListGeofenceGroups(ctx).Execute()

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
	// response from `AppsGeofencesApiListGeofenceGroups`: GeofenceGroupsOut
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiListGeofenceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiListGeofenceGroupsRequest struct via the builder pattern


### Return type

[**GeofenceGroupsOut**](GeofenceGroupsOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiListGeofenceWorkflows

> GeofenceWorkflowReferencesOut AppsGeofencesApiListGeofenceWorkflows(ctx, geofenceId).Execute()

List Geofence Workflows



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
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiListGeofenceWorkflows(context.Background(), geofenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiListGeofenceWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiListGeofenceWorkflows`: GeofenceWorkflowReferencesOut
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiListGeofenceWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiListGeofenceWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GeofenceWorkflowReferencesOut**](GeofenceWorkflowReferencesOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiListGeofences

> GeofenceListResponse AppsGeofencesApiListGeofences(ctx).Limit(limit).Offset(offset).ActiveOnly(activeOnly).Tags(tags).IncludeArchived(includeArchived).Execute()

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
	limit := int32(56) // int32 | Maximum 500; larger values are capped (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)
	activeOnly := true // bool |  (optional) (default to true)
	tags := []string{"Inner_example"} // []string | Filter by tag names. Repeat key for AND-semantics: ?tags=foo&tags=bar. (optional)
	includeArchived := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiListGeofences(context.Background()).Limit(limit).Offset(offset).ActiveOnly(activeOnly).Tags(tags).IncludeArchived(includeArchived).Execute()
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
 **limit** | **int32** | Maximum 500; larger values are capped | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **activeOnly** | **bool** |  | [default to true]
 **tags** | **[]string** | Filter by tag names. Repeat key for AND-semantics: ?tags&#x3D;foo&amp;tags&#x3D;bar. | 
 **includeArchived** | **bool** |  | [default to false]

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

> GroupGeofencesOut AppsGeofencesApiListGroupGeofences(ctx, groupId).Execute()

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
	// response from `AppsGeofencesApiListGroupGeofences`: GroupGeofencesOut
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

[**GroupGeofencesOut**](GroupGeofencesOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiListWorkspaceTags

> TagListResponse AppsGeofencesApiListWorkspaceTags(ctx).Execute()

List Workspace Tags



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
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiListWorkspaceTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiListWorkspaceTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiListWorkspaceTags`: TagListResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiListWorkspaceTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiListWorkspaceTagsRequest struct via the builder pattern


### Return type

[**TagListResponse**](TagListResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGeofencesApiTestGroupPoint

> GroupTestPointOut AppsGeofencesApiTestGroupPoint(ctx, groupId).TestPointRequest(testPointRequest).Execute()

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
	// response from `AppsGeofencesApiTestGroupPoint`: GroupTestPointOut
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

[**GroupTestPointOut**](GroupTestPointOut.md)

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

> map[string]interface{} AppsGeofencesApiTriggerTestEvent(ctx, geofenceId).TestEventRequest(testEventRequest).Execute()

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
	testEventRequest := *openapiclient.NewTestEventRequest("EventType_example") // TestEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiTriggerTestEvent(context.Background(), geofenceId).TestEventRequest(testEventRequest).Execute()
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


## AppsGeofencesApiUnarchiveGeofence

> GeofenceResponse AppsGeofencesApiUnarchiveGeofence(ctx, geofenceId).Execute()

Unarchive Geofence



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
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiUnarchiveGeofence(context.Background(), geofenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiUnarchiveGeofence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiUnarchiveGeofence`: GeofenceResponse
	fmt.Fprintf(os.Stdout, "Response from `GeofencesAPI.AppsGeofencesApiUnarchiveGeofence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geofenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGeofencesApiUnarchiveGeofenceRequest struct via the builder pattern


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

> GeofenceGroupUpdateOut AppsGeofencesApiUpdateGeofenceGroup(ctx, geofenceId).Body(body).Execute()

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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeofencesAPI.AppsGeofencesApiUpdateGeofenceGroup(context.Background(), geofenceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeofencesAPI.AppsGeofencesApiUpdateGeofenceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGeofencesApiUpdateGeofenceGroup`: GeofenceGroupUpdateOut
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

 **body** | **string** |  | 

### Return type

[**GeofenceGroupUpdateOut**](GeofenceGroupUpdateOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
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

