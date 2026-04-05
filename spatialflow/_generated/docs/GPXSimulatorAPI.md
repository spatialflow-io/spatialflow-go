# \GPXSimulatorAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsDevicesApiGpxCancelGpxPlayback**](GPXSimulatorAPI.md#AppsDevicesApiGpxCancelGpxPlayback) | **Post** /api/v1/gpx/playbacks/{playback_id}/cancel | Cancel Gpx Playback
[**AppsDevicesApiGpxDeleteGpxRoute**](GPXSimulatorAPI.md#AppsDevicesApiGpxDeleteGpxRoute) | **Delete** /api/v1/gpx/routes/{route_id} | Delete Gpx Route
[**AppsDevicesApiGpxGetGpxPlayback**](GPXSimulatorAPI.md#AppsDevicesApiGpxGetGpxPlayback) | **Get** /api/v1/gpx/playbacks/{playback_id} | Get Gpx Playback
[**AppsDevicesApiGpxGetGpxRoute**](GPXSimulatorAPI.md#AppsDevicesApiGpxGetGpxRoute) | **Get** /api/v1/gpx/routes/{route_id} | Get Gpx Route
[**AppsDevicesApiGpxListGpxPlaybacks**](GPXSimulatorAPI.md#AppsDevicesApiGpxListGpxPlaybacks) | **Get** /api/v1/gpx/playbacks | List Gpx Playbacks
[**AppsDevicesApiGpxListGpxRoutes**](GPXSimulatorAPI.md#AppsDevicesApiGpxListGpxRoutes) | **Get** /api/v1/gpx/routes | List Gpx Routes
[**AppsDevicesApiGpxPauseGpxPlayback**](GPXSimulatorAPI.md#AppsDevicesApiGpxPauseGpxPlayback) | **Post** /api/v1/gpx/playbacks/{playback_id}/pause | Pause Gpx Playback
[**AppsDevicesApiGpxResumeGpxPlayback**](GPXSimulatorAPI.md#AppsDevicesApiGpxResumeGpxPlayback) | **Post** /api/v1/gpx/playbacks/{playback_id}/resume | Resume Gpx Playback
[**AppsDevicesApiGpxStartGpxPlayback**](GPXSimulatorAPI.md#AppsDevicesApiGpxStartGpxPlayback) | **Post** /api/v1/gpx/routes/{route_id}/playback | Start Gpx Playback
[**AppsDevicesApiGpxUploadGpxRoute**](GPXSimulatorAPI.md#AppsDevicesApiGpxUploadGpxRoute) | **Post** /api/v1/gpx/routes/upload | Upload Gpx Route



## AppsDevicesApiGpxCancelGpxPlayback

> map[string]interface{} AppsDevicesApiGpxCancelGpxPlayback(ctx, playbackId).Execute()

Cancel Gpx Playback



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
	playbackId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxCancelGpxPlayback(context.Background(), playbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxCancelGpxPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxCancelGpxPlayback`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxCancelGpxPlayback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playbackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxCancelGpxPlaybackRequest struct via the builder pattern


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


## AppsDevicesApiGpxDeleteGpxRoute

> map[string]interface{} AppsDevicesApiGpxDeleteGpxRoute(ctx, routeId).Execute()

Delete Gpx Route



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
	routeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxDeleteGpxRoute(context.Background(), routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxDeleteGpxRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxDeleteGpxRoute`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxDeleteGpxRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxDeleteGpxRouteRequest struct via the builder pattern


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


## AppsDevicesApiGpxGetGpxPlayback

> GPXPlaybackOut AppsDevicesApiGpxGetGpxPlayback(ctx, playbackId).Execute()

Get Gpx Playback



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
	playbackId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxGetGpxPlayback(context.Background(), playbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxGetGpxPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxGetGpxPlayback`: GPXPlaybackOut
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxGetGpxPlayback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playbackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxGetGpxPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GPXPlaybackOut**](GPXPlaybackOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGpxGetGpxRoute

> GPXRouteOut AppsDevicesApiGpxGetGpxRoute(ctx, routeId).Execute()

Get Gpx Route



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
	routeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxGetGpxRoute(context.Background(), routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxGetGpxRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxGetGpxRoute`: GPXRouteOut
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxGetGpxRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxGetGpxRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GPXRouteOut**](GPXRouteOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGpxListGpxPlaybacks

> []GPXPlaybackOut AppsDevicesApiGpxListGpxPlaybacks(ctx).Status(status).Limit(limit).Offset(offset).Execute()

List Gpx Playbacks



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
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxListGpxPlaybacks(context.Background()).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxListGpxPlaybacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxListGpxPlaybacks`: []GPXPlaybackOut
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxListGpxPlaybacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxListGpxPlaybacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]GPXPlaybackOut**](GPXPlaybackOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGpxListGpxRoutes

> []GPXRouteOut AppsDevicesApiGpxListGpxRoutes(ctx).DeviceId(deviceId).Limit(limit).Offset(offset).Execute()

List Gpx Routes



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
	deviceId := "deviceId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxListGpxRoutes(context.Background()).DeviceId(deviceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxListGpxRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxListGpxRoutes`: []GPXRouteOut
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxListGpxRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxListGpxRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]GPXRouteOut**](GPXRouteOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGpxPauseGpxPlayback

> GPXPlaybackOut AppsDevicesApiGpxPauseGpxPlayback(ctx, playbackId).Execute()

Pause Gpx Playback



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
	playbackId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxPauseGpxPlayback(context.Background(), playbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxPauseGpxPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxPauseGpxPlayback`: GPXPlaybackOut
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxPauseGpxPlayback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playbackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxPauseGpxPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GPXPlaybackOut**](GPXPlaybackOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGpxResumeGpxPlayback

> GPXPlaybackOut AppsDevicesApiGpxResumeGpxPlayback(ctx, playbackId).Execute()

Resume Gpx Playback



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
	playbackId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxResumeGpxPlayback(context.Background(), playbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxResumeGpxPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxResumeGpxPlayback`: GPXPlaybackOut
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxResumeGpxPlayback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playbackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxResumeGpxPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GPXPlaybackOut**](GPXPlaybackOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGpxStartGpxPlayback

> GPXPlaybackOut AppsDevicesApiGpxStartGpxPlayback(ctx, routeId).StartPlaybackRequest(startPlaybackRequest).Execute()

Start Gpx Playback



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
	routeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startPlaybackRequest := *openapiclient.NewStartPlaybackRequest() // StartPlaybackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxStartGpxPlayback(context.Background(), routeId).StartPlaybackRequest(startPlaybackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxStartGpxPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxStartGpxPlayback`: GPXPlaybackOut
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxStartGpxPlayback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxStartGpxPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startPlaybackRequest** | [**StartPlaybackRequest**](StartPlaybackRequest.md) |  | 

### Return type

[**GPXPlaybackOut**](GPXPlaybackOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGpxUploadGpxRoute

> GPXRouteOut AppsDevicesApiGpxUploadGpxRoute(ctx).DeviceId(deviceId).Name(name).File(file).Description(description).Execute()

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
	deviceId := "deviceId_example" // string | 
	name := "name_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 
	description := "description_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GPXSimulatorAPI.AppsDevicesApiGpxUploadGpxRoute(context.Background()).DeviceId(deviceId).Name(name).File(file).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GPXSimulatorAPI.AppsDevicesApiGpxUploadGpxRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGpxUploadGpxRoute`: GPXRouteOut
	fmt.Fprintf(os.Stdout, "Response from `GPXSimulatorAPI.AppsDevicesApiGpxUploadGpxRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGpxUploadGpxRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** |  | 
 **name** | **string** |  | 
 **file** | ***os.File** |  | 
 **description** | **string** |  | [default to &quot;&quot;]

### Return type

[**GPXRouteOut**](GPXRouteOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

