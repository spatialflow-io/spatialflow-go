# \SimulationCenterAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsDevicesApiSimulationCreateSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationCreateSimulation) | **Post** /api/v1/simulations | Create Simulation
[**AppsDevicesApiSimulationDeleteSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationDeleteSimulation) | **Delete** /api/v1/simulations/{simulation_id} | Delete Simulation
[**AppsDevicesApiSimulationGetRouteTrackPoints**](SimulationCenterAPI.md#AppsDevicesApiSimulationGetRouteTrackPoints) | **Get** /api/v1/simulations/{simulation_id}/routes/{route_id}/track-points | Get Route Track Points
[**AppsDevicesApiSimulationGetSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationGetSimulation) | **Get** /api/v1/simulations/{simulation_id} | Get Simulation
[**AppsDevicesApiSimulationListSimulationEvents**](SimulationCenterAPI.md#AppsDevicesApiSimulationListSimulationEvents) | **Get** /api/v1/simulations/{simulation_id}/events | List Simulation Events
[**AppsDevicesApiSimulationListSimulations**](SimulationCenterAPI.md#AppsDevicesApiSimulationListSimulations) | **Get** /api/v1/simulations | List Simulations
[**AppsDevicesApiSimulationPauseSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationPauseSimulation) | **Post** /api/v1/simulations/{simulation_id}/pause | Pause Simulation
[**AppsDevicesApiSimulationRemoveRoute**](SimulationCenterAPI.md#AppsDevicesApiSimulationRemoveRoute) | **Delete** /api/v1/simulations/{simulation_id}/routes/{route_id} | Remove Route
[**AppsDevicesApiSimulationResetSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationResetSimulation) | **Post** /api/v1/simulations/{simulation_id}/reset | Reset Simulation
[**AppsDevicesApiSimulationResumeSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationResumeSimulation) | **Post** /api/v1/simulations/{simulation_id}/resume | Resume Simulation
[**AppsDevicesApiSimulationStartSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationStartSimulation) | **Post** /api/v1/simulations/{simulation_id}/start | Start Simulation
[**AppsDevicesApiSimulationStopSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationStopSimulation) | **Post** /api/v1/simulations/{simulation_id}/stop | Stop Simulation
[**AppsDevicesApiSimulationUpdateSimulation**](SimulationCenterAPI.md#AppsDevicesApiSimulationUpdateSimulation) | **Patch** /api/v1/simulations/{simulation_id} | Update Simulation
[**AppsDevicesApiSimulationUploadCsvRoute**](SimulationCenterAPI.md#AppsDevicesApiSimulationUploadCsvRoute) | **Post** /api/v1/simulations/{simulation_id}/routes/upload-csv | Upload Csv Route
[**AppsDevicesApiSimulationUploadGpxRoute**](SimulationCenterAPI.md#AppsDevicesApiSimulationUploadGpxRoute) | **Post** /api/v1/simulations/{simulation_id}/routes/upload-gpx | Upload Gpx Route



## AppsDevicesApiSimulationCreateSimulation

> SimulationOut AppsDevicesApiSimulationCreateSimulation(ctx).CreateSimulationRequest(createSimulationRequest).Execute()

Create Simulation



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
	createSimulationRequest := *openapiclient.NewCreateSimulationRequest("Name_example") // CreateSimulationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationCreateSimulation(context.Background()).CreateSimulationRequest(createSimulationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationCreateSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationCreateSimulation`: SimulationOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationCreateSimulation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationCreateSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSimulationRequest** | [**CreateSimulationRequest**](CreateSimulationRequest.md) |  | 

### Return type

[**SimulationOut**](SimulationOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationDeleteSimulation

> map[string]interface{} AppsDevicesApiSimulationDeleteSimulation(ctx, simulationId).Execute()

Delete Simulation



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationDeleteSimulation(context.Background(), simulationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationDeleteSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationDeleteSimulation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationDeleteSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationDeleteSimulationRequest struct via the builder pattern


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


## AppsDevicesApiSimulationGetRouteTrackPoints

> RouteTrackPointsOut AppsDevicesApiSimulationGetRouteTrackPoints(ctx, simulationId, routeId).Execute()

Get Route Track Points



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	routeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationGetRouteTrackPoints(context.Background(), simulationId, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationGetRouteTrackPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationGetRouteTrackPoints`: RouteTrackPointsOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationGetRouteTrackPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationGetRouteTrackPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RouteTrackPointsOut**](RouteTrackPointsOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationGetSimulation

> SimulationDetailOut AppsDevicesApiSimulationGetSimulation(ctx, simulationId).Execute()

Get Simulation



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationGetSimulation(context.Background(), simulationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationGetSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationGetSimulation`: SimulationDetailOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationGetSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationGetSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimulationDetailOut**](SimulationDetailOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationListSimulationEvents

> []SimulationEventOut AppsDevicesApiSimulationListSimulationEvents(ctx, simulationId).EventType(eventType).Limit(limit).Offset(offset).Execute()

List Simulation Events



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	eventType := "eventType_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationListSimulationEvents(context.Background(), simulationId).EventType(eventType).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationListSimulationEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationListSimulationEvents`: []SimulationEventOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationListSimulationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationListSimulationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventType** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]SimulationEventOut**](SimulationEventOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationListSimulations

> []SimulationOut AppsDevicesApiSimulationListSimulations(ctx).Status(status).Limit(limit).Offset(offset).Execute()

List Simulations



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
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationListSimulations(context.Background()).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationListSimulations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationListSimulations`: []SimulationOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationListSimulations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationListSimulationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]SimulationOut**](SimulationOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationPauseSimulation

> SimulationOut AppsDevicesApiSimulationPauseSimulation(ctx, simulationId).Execute()

Pause Simulation



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationPauseSimulation(context.Background(), simulationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationPauseSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationPauseSimulation`: SimulationOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationPauseSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationPauseSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimulationOut**](SimulationOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationRemoveRoute

> map[string]interface{} AppsDevicesApiSimulationRemoveRoute(ctx, simulationId, routeId).Execute()

Remove Route



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	routeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationRemoveRoute(context.Background(), simulationId, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationRemoveRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationRemoveRoute`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationRemoveRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationRemoveRouteRequest struct via the builder pattern


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


## AppsDevicesApiSimulationResetSimulation

> SimulationOut AppsDevicesApiSimulationResetSimulation(ctx, simulationId).Execute()

Reset Simulation



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationResetSimulation(context.Background(), simulationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationResetSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationResetSimulation`: SimulationOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationResetSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationResetSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimulationOut**](SimulationOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationResumeSimulation

> SimulationOut AppsDevicesApiSimulationResumeSimulation(ctx, simulationId).Execute()

Resume Simulation



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationResumeSimulation(context.Background(), simulationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationResumeSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationResumeSimulation`: SimulationOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationResumeSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationResumeSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimulationOut**](SimulationOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationStartSimulation

> SimulationOut AppsDevicesApiSimulationStartSimulation(ctx, simulationId).Execute()

Start Simulation



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationStartSimulation(context.Background(), simulationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationStartSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationStartSimulation`: SimulationOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationStartSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationStartSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimulationOut**](SimulationOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationStopSimulation

> SimulationOut AppsDevicesApiSimulationStopSimulation(ctx, simulationId).Execute()

Stop Simulation



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationStopSimulation(context.Background(), simulationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationStopSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationStopSimulation`: SimulationOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationStopSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationStopSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimulationOut**](SimulationOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationUpdateSimulation

> SimulationOut AppsDevicesApiSimulationUpdateSimulation(ctx, simulationId).UpdateSimulationRequest(updateSimulationRequest).Execute()

Update Simulation



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateSimulationRequest := *openapiclient.NewUpdateSimulationRequest() // UpdateSimulationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationUpdateSimulation(context.Background(), simulationId).UpdateSimulationRequest(updateSimulationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationUpdateSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationUpdateSimulation`: SimulationOut
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationUpdateSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationUpdateSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSimulationRequest** | [**UpdateSimulationRequest**](UpdateSimulationRequest.md) |  | 

### Return type

[**SimulationOut**](SimulationOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationUploadCsvRoute

> RouteUploadResponse AppsDevicesApiSimulationUploadCsvRoute(ctx, simulationId).File(file).SimulatedDeviceName(simulatedDeviceName).SimulatedDeviceType(simulatedDeviceType).Execute()

Upload Csv Route



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 
	simulatedDeviceName := "simulatedDeviceName_example" // string |  (optional) (default to "")
	simulatedDeviceType := "simulatedDeviceType_example" // string |  (optional) (default to "simulation")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationUploadCsvRoute(context.Background(), simulationId).File(file).SimulatedDeviceName(simulatedDeviceName).SimulatedDeviceType(simulatedDeviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationUploadCsvRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationUploadCsvRoute`: RouteUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationUploadCsvRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationUploadCsvRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **simulatedDeviceName** | **string** |  | [default to &quot;&quot;]
 **simulatedDeviceType** | **string** |  | [default to &quot;simulation&quot;]

### Return type

[**RouteUploadResponse**](RouteUploadResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiSimulationUploadGpxRoute

> RouteUploadResponse AppsDevicesApiSimulationUploadGpxRoute(ctx, simulationId).File(file).SimulatedDeviceName(simulatedDeviceName).SimulatedDeviceType(simulatedDeviceType).Execute()

Upload Gpx Route



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
	simulationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 
	simulatedDeviceName := "simulatedDeviceName_example" // string |  (optional) (default to "")
	simulatedDeviceType := "simulatedDeviceType_example" // string |  (optional) (default to "simulation")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulationCenterAPI.AppsDevicesApiSimulationUploadGpxRoute(context.Background(), simulationId).File(file).SimulatedDeviceName(simulatedDeviceName).SimulatedDeviceType(simulatedDeviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulationCenterAPI.AppsDevicesApiSimulationUploadGpxRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiSimulationUploadGpxRoute`: RouteUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `SimulationCenterAPI.AppsDevicesApiSimulationUploadGpxRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simulationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiSimulationUploadGpxRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **simulatedDeviceName** | **string** |  | [default to &quot;&quot;]
 **simulatedDeviceType** | **string** |  | [default to &quot;simulation&quot;]

### Return type

[**RouteUploadResponse**](RouteUploadResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

