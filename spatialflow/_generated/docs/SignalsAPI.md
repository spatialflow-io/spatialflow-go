# \SignalsAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsDevicesApiSignalsGetSignal**](SignalsAPI.md#AppsDevicesApiSignalsGetSignal) | **Get** /api/v1/signals/{signal_id}/ | Get signal event detail
[**AppsDevicesApiSignalsListSignals**](SignalsAPI.md#AppsDevicesApiSignalsListSignals) | **Get** /api/v1/signals/ | List signal events



## AppsDevicesApiSignalsGetSignal

> SignalEventDetailOut AppsDevicesApiSignalsGetSignal(ctx, signalId).Execute()

Get signal event detail

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
	signalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalsAPI.AppsDevicesApiSignalsGetSignal(context.Background(), signalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalsAPI.AppsDevicesApiSignalsGetSignal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSignalsGetSignal`: SignalEventDetailOut
	fmt.Fprintf(os.Stdout, "Response from `SignalsAPI.AppsDevicesApiSignalsGetSignal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSignalsGetSignalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignalEventDetailOut**](SignalEventDetailOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSignalsListSignals

> SignalEventsListOut AppsDevicesApiSignalsListSignals(ctx).SignalType(signalType).State(state).DeviceId(deviceId).GeofenceId(geofenceId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

List signal events

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
	signalType := "signalType_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	geofenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalsAPI.AppsDevicesApiSignalsListSignals(context.Background()).SignalType(signalType).State(state).DeviceId(deviceId).GeofenceId(geofenceId).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalsAPI.AppsDevicesApiSignalsListSignals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSignalsListSignals`: SignalEventsListOut
	fmt.Fprintf(os.Stdout, "Response from `SignalsAPI.AppsDevicesApiSignalsListSignals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSignalsListSignalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalType** | **string** |  | 
 **state** | **string** |  | 
 **deviceId** | **string** |  | 
 **geofenceId** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**SignalEventsListOut**](SignalEventsListOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

