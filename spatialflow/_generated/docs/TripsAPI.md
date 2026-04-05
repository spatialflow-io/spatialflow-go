# \TripsAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsDevicesApiTripsCancelTrip**](TripsAPI.md#AppsDevicesApiTripsCancelTrip) | **Post** /api/v1/trips/{trip_id}/cancel | Cancel a planned trip
[**AppsDevicesApiTripsCreateTrip**](TripsAPI.md#AppsDevicesApiTripsCreateTrip) | **Post** /api/v1/trips/ | Create a planned trip
[**AppsDevicesApiTripsGetTrip**](TripsAPI.md#AppsDevicesApiTripsGetTrip) | **Get** /api/v1/trips/{trip_id}/ | Get trip detail
[**AppsDevicesApiTripsListTrips**](TripsAPI.md#AppsDevicesApiTripsListTrips) | **Get** /api/v1/trips/ | List trips
[**AppsDevicesApiTripsUpdateTrip**](TripsAPI.md#AppsDevicesApiTripsUpdateTrip) | **Put** /api/v1/trips/{trip_id}/ | Update a trip



## AppsDevicesApiTripsCancelTrip

> TripOut AppsDevicesApiTripsCancelTrip(ctx, tripId).Execute()

Cancel a planned trip

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
	tripId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TripsAPI.AppsDevicesApiTripsCancelTrip(context.Background(), tripId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TripsAPI.AppsDevicesApiTripsCancelTrip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiTripsCancelTrip`: TripOut
	fmt.Fprintf(os.Stdout, "Response from `TripsAPI.AppsDevicesApiTripsCancelTrip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tripId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiTripsCancelTripRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TripOut**](TripOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiTripsCreateTrip

> TripOut AppsDevicesApiTripsCreateTrip(ctx).TripCreateIn(tripCreateIn).Execute()

Create a planned trip

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
	tripCreateIn := *openapiclient.NewTripCreateIn("DeviceId_example") // TripCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TripsAPI.AppsDevicesApiTripsCreateTrip(context.Background()).TripCreateIn(tripCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TripsAPI.AppsDevicesApiTripsCreateTrip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiTripsCreateTrip`: TripOut
	fmt.Fprintf(os.Stdout, "Response from `TripsAPI.AppsDevicesApiTripsCreateTrip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiTripsCreateTripRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tripCreateIn** | [**TripCreateIn**](TripCreateIn.md) |  | 

### Return type

[**TripOut**](TripOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiTripsGetTrip

> TripDetailOut AppsDevicesApiTripsGetTrip(ctx, tripId).Execute()

Get trip detail

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
	tripId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TripsAPI.AppsDevicesApiTripsGetTrip(context.Background(), tripId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TripsAPI.AppsDevicesApiTripsGetTrip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiTripsGetTrip`: TripDetailOut
	fmt.Fprintf(os.Stdout, "Response from `TripsAPI.AppsDevicesApiTripsGetTrip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tripId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiTripsGetTripRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TripDetailOut**](TripDetailOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiTripsListTrips

> TripsListOut AppsDevicesApiTripsListTrips(ctx).DeviceId(deviceId).Status(status).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

List trips

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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	status := "status_example" // string |  (optional)
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TripsAPI.AppsDevicesApiTripsListTrips(context.Background()).DeviceId(deviceId).Status(status).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TripsAPI.AppsDevicesApiTripsListTrips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiTripsListTrips`: TripsListOut
	fmt.Fprintf(os.Stdout, "Response from `TripsAPI.AppsDevicesApiTripsListTrips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiTripsListTripsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** |  | 
 **status** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**TripsListOut**](TripsListOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiTripsUpdateTrip

> TripOut AppsDevicesApiTripsUpdateTrip(ctx, tripId).TripUpdateIn(tripUpdateIn).Execute()

Update a trip

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
	tripId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tripUpdateIn := *openapiclient.NewTripUpdateIn() // TripUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TripsAPI.AppsDevicesApiTripsUpdateTrip(context.Background(), tripId).TripUpdateIn(tripUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TripsAPI.AppsDevicesApiTripsUpdateTrip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiTripsUpdateTrip`: TripOut
	fmt.Fprintf(os.Stdout, "Response from `TripsAPI.AppsDevicesApiTripsUpdateTrip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tripId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiTripsUpdateTripRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tripUpdateIn** | [**TripUpdateIn**](TripUpdateIn.md) |  | 

### Return type

[**TripOut**](TripOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

