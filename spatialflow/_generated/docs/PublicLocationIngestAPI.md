# \PublicLocationIngestAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsPublicLocationsApiGetIngestStats**](PublicLocationIngestAPI.md#AppsPublicLocationsApiGetIngestStats) | **Get** /api/v1/locations/stats | Get Ingest Stats
[**AppsPublicLocationsApiIngestLocation**](PublicLocationIngestAPI.md#AppsPublicLocationsApiIngestLocation) | **Post** /api/v1/locations | Ingest Location
[**AppsPublicLocationsApiIngestLocationBatch**](PublicLocationIngestAPI.md#AppsPublicLocationsApiIngestLocationBatch) | **Post** /api/v1/locations/batch | Ingest Location Batch



## AppsPublicLocationsApiGetIngestStats

> map[string]interface{} AppsPublicLocationsApiGetIngestStats(ctx).Execute()

Get Ingest Stats



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
	resp, r, err := apiClient.PublicLocationIngestAPI.AppsPublicLocationsApiGetIngestStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicLocationIngestAPI.AppsPublicLocationsApiGetIngestStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicLocationsApiGetIngestStats`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PublicLocationIngestAPI.AppsPublicLocationsApiGetIngestStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicLocationsApiGetIngestStatsRequest struct via the builder pattern


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


## AppsPublicLocationsApiIngestLocation

> LocationIngestResponse AppsPublicLocationsApiIngestLocation(ctx).LocationPointIn(locationPointIn).Execute()

Ingest Location



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
	locationPointIn := *openapiclient.NewLocationPointIn("DeviceId_example", float32(123), float32(123), time.Now()) // LocationPointIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicLocationIngestAPI.AppsPublicLocationsApiIngestLocation(context.Background()).LocationPointIn(locationPointIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicLocationIngestAPI.AppsPublicLocationsApiIngestLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicLocationsApiIngestLocation`: LocationIngestResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicLocationIngestAPI.AppsPublicLocationsApiIngestLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicLocationsApiIngestLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationPointIn** | [**LocationPointIn**](LocationPointIn.md) |  | 

### Return type

[**LocationIngestResponse**](LocationIngestResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPublicLocationsApiIngestLocationBatch

> LocationIngestResponse AppsPublicLocationsApiIngestLocationBatch(ctx).LocationBatchIn(locationBatchIn).Execute()

Ingest Location Batch



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
	locationBatchIn := *openapiclient.NewLocationBatchIn([]openapiclient.LocationPointIn{*openapiclient.NewLocationPointIn("DeviceId_example", float32(123), float32(123), time.Now())}) // LocationBatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicLocationIngestAPI.AppsPublicLocationsApiIngestLocationBatch(context.Background()).LocationBatchIn(locationBatchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicLocationIngestAPI.AppsPublicLocationsApiIngestLocationBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPublicLocationsApiIngestLocationBatch`: LocationIngestResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicLocationIngestAPI.AppsPublicLocationsApiIngestLocationBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsPublicLocationsApiIngestLocationBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationBatchIn** | [**LocationBatchIn**](LocationBatchIn.md) |  | 

### Return type

[**LocationIngestResponse**](LocationIngestResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

