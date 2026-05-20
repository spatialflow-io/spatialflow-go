# \DevicesAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsDevicesApiActivateDevice**](DevicesAPI.md#AppsDevicesApiActivateDevice) | **Post** /api/v1/devices/{device_id}/activate | Activate Device
[**AppsDevicesApiBatchUpdateLocations**](DevicesAPI.md#AppsDevicesApiBatchUpdateLocations) | **Post** /api/v1/devices/batch-update | Batch Update Locations
[**AppsDevicesApiCreateDevice**](DevicesAPI.md#AppsDevicesApiCreateDevice) | **Post** /api/v1/devices/ | Create Device
[**AppsDevicesApiCreateManagerSessionNote**](DevicesAPI.md#AppsDevicesApiCreateManagerSessionNote) | **Post** /api/v1/devices/{device_uuid}/sessions/{session_id}/notes | Create Manager Session Note
[**AppsDevicesApiDeactivateDevice**](DevicesAPI.md#AppsDevicesApiDeactivateDevice) | **Post** /api/v1/devices/{device_id}/deactivate | Deactivate Device
[**AppsDevicesApiDeleteDevice**](DevicesAPI.md#AppsDevicesApiDeleteDevice) | **Delete** /api/v1/devices/{device_id} | Delete Device
[**AppsDevicesApiEndShift**](DevicesAPI.md#AppsDevicesApiEndShift) | **Post** /api/v1/devices/{device_id}/end-shift | End Shift
[**AppsDevicesApiExportEventsEndpoint**](DevicesAPI.md#AppsDevicesApiExportEventsEndpoint) | **Get** /api/v1/devices/events/export | Export Events Endpoint
[**AppsDevicesApiGetActiveSession**](DevicesAPI.md#AppsDevicesApiGetActiveSession) | **Get** /api/v1/devices/{device_uuid}/active-session | Get Active Session
[**AppsDevicesApiGetDashboardStats**](DevicesAPI.md#AppsDevicesApiGetDashboardStats) | **Get** /api/v1/devices/dashboard-stats | Get Dashboard Stats
[**AppsDevicesApiGetDashboardStatsTimeline**](DevicesAPI.md#AppsDevicesApiGetDashboardStatsTimeline) | **Get** /api/v1/devices/dashboard-stats/timeline | Get Dashboard Stats Timeline
[**AppsDevicesApiGetDevice**](DevicesAPI.md#AppsDevicesApiGetDevice) | **Get** /api/v1/devices/{device_id} | Get Device
[**AppsDevicesApiGetDeviceEvents**](DevicesAPI.md#AppsDevicesApiGetDeviceEvents) | **Get** /api/v1/devices/{device_id}/events | Get Device Events
[**AppsDevicesApiGetDeviceSessions**](DevicesAPI.md#AppsDevicesApiGetDeviceSessions) | **Get** /api/v1/devices/{device_id}/sessions | Get Device Sessions
[**AppsDevicesApiGetEventDetail**](DevicesAPI.md#AppsDevicesApiGetEventDetail) | **Get** /api/v1/devices/events/{event_id} | Get Event Detail
[**AppsDevicesApiGetImportJob**](DevicesAPI.md#AppsDevicesApiGetImportJob) | **Get** /api/v1/devices/locations/import/{job_id} | Get Import Job
[**AppsDevicesApiGetLocationStats**](DevicesAPI.md#AppsDevicesApiGetLocationStats) | **Get** /api/v1/devices/stats | Get Location Stats
[**AppsDevicesApiGetRecentEvents**](DevicesAPI.md#AppsDevicesApiGetRecentEvents) | **Get** /api/v1/devices/events/recent | Get Recent Events
[**AppsDevicesApiGetRecentLocations**](DevicesAPI.md#AppsDevicesApiGetRecentLocations) | **Get** /api/v1/devices/{device_id}/locations/recent | Get Recent Locations
[**AppsDevicesApiGetSessionDetail**](DevicesAPI.md#AppsDevicesApiGetSessionDetail) | **Get** /api/v1/devices/{device_id}/sessions/{session_id} | Get Session Detail
[**AppsDevicesApiGetSessionLocations**](DevicesAPI.md#AppsDevicesApiGetSessionLocations) | **Get** /api/v1/devices/{device_id}/sessions/{session_id}/locations | Get Session Locations
[**AppsDevicesApiListDevices**](DevicesAPI.md#AppsDevicesApiListDevices) | **Get** /api/v1/devices/ | List Devices
[**AppsDevicesApiListSessionAttachments**](DevicesAPI.md#AppsDevicesApiListSessionAttachments) | **Get** /api/v1/devices/{device_uuid}/sessions/{session_id}/attachments | List Session Attachments
[**AppsDevicesApiListSessionNotes**](DevicesAPI.md#AppsDevicesApiListSessionNotes) | **Get** /api/v1/devices/{device_uuid}/sessions/{session_id}/notes | List Session Notes
[**AppsDevicesApiListSessionPhotos**](DevicesAPI.md#AppsDevicesApiListSessionPhotos) | **Get** /api/v1/devices/{device_uuid}/sessions/{session_id}/photos | List Session Photos
[**AppsDevicesApiListWorkspacePhotos**](DevicesAPI.md#AppsDevicesApiListWorkspacePhotos) | **Get** /api/v1/devices/photos | List Workspace Photos
[**AppsDevicesApiPauseShift**](DevicesAPI.md#AppsDevicesApiPauseShift) | **Post** /api/v1/devices/{device_id}/pause-shift | Pause Shift
[**AppsDevicesApiResumeShift**](DevicesAPI.md#AppsDevicesApiResumeShift) | **Post** /api/v1/devices/{device_id}/resume-shift | Resume Shift
[**AppsDevicesApiStartShift**](DevicesAPI.md#AppsDevicesApiStartShift) | **Post** /api/v1/devices/{device_id}/start-shift | Start Shift
[**AppsDevicesApiUpdateDevice**](DevicesAPI.md#AppsDevicesApiUpdateDevice) | **Put** /api/v1/devices/{device_id} | Update Device
[**AppsDevicesApiUpdateDeviceLocation**](DevicesAPI.md#AppsDevicesApiUpdateDeviceLocation) | **Post** /api/v1/devices/{device_id}/location | Update Device Location
[**AppsDevicesApiUpdateSessionNotes**](DevicesAPI.md#AppsDevicesApiUpdateSessionNotes) | **Post** /api/v1/devices/{device_id}/notes | Update Session Notes
[**AppsDevicesApiUploadCsvImport**](DevicesAPI.md#AppsDevicesApiUploadCsvImport) | **Post** /api/v1/devices/locations/import | Upload Csv Import



## AppsDevicesApiActivateDevice

> DeviceOut AppsDevicesApiActivateDevice(ctx, deviceId).Execute()

Activate Device



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiActivateDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiActivateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiActivateDevice`: DeviceOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiActivateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiActivateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceOut**](DeviceOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiBatchUpdateLocations

> []LocationUpdateOut AppsDevicesApiBatchUpdateLocations(ctx).BatchLocationUpdateIn(batchLocationUpdateIn).Execute()

Batch Update Locations



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
	batchLocationUpdateIn := []openapiclient.BatchLocationUpdateIn{*openapiclient.NewBatchLocationUpdateIn("DeviceId_example", []openapiclient.LocationUpdateIn{*openapiclient.NewLocationUpdateIn(float32(123), float32(123))})} // []BatchLocationUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiBatchUpdateLocations(context.Background()).BatchLocationUpdateIn(batchLocationUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiBatchUpdateLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiBatchUpdateLocations`: []LocationUpdateOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiBatchUpdateLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiBatchUpdateLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchLocationUpdateIn** | [**[]BatchLocationUpdateIn**](BatchLocationUpdateIn.md) |  | 

### Return type

[**[]LocationUpdateOut**](LocationUpdateOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiCreateDevice

> DeviceOut AppsDevicesApiCreateDevice(ctx).DeviceIn(deviceIn).Execute()

Create Device



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
	deviceIn := *openapiclient.NewDeviceIn("DeviceId_example", "Name_example") // DeviceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiCreateDevice(context.Background()).DeviceIn(deviceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiCreateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiCreateDevice`: DeviceOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiCreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceIn** | [**DeviceIn**](DeviceIn.md) |  | 

### Return type

[**DeviceOut**](DeviceOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiCreateManagerSessionNote

> SessionNoteOut AppsDevicesApiCreateManagerSessionNote(ctx, deviceUuid, sessionId).SessionNoteIn(sessionNoteIn).Execute()

Create Manager Session Note



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
	deviceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionNoteIn := *openapiclient.NewSessionNoteIn("Body_example") // SessionNoteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiCreateManagerSessionNote(context.Background(), deviceUuid, sessionId).SessionNoteIn(sessionNoteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiCreateManagerSessionNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiCreateManagerSessionNote`: SessionNoteOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiCreateManagerSessionNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiCreateManagerSessionNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sessionNoteIn** | [**SessionNoteIn**](SessionNoteIn.md) |  | 

### Return type

[**SessionNoteOut**](SessionNoteOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiDeactivateDevice

> DeviceOut AppsDevicesApiDeactivateDevice(ctx, deviceId).Execute()

Deactivate Device



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiDeactivateDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiDeactivateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiDeactivateDevice`: DeviceOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiDeactivateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiDeactivateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceOut**](DeviceOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiDeleteDevice

> AppsDevicesApiDeleteDevice(ctx, deviceId).Execute()

Delete Device



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.AppsDevicesApiDeleteDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiDeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiDeleteDeviceRequest struct via the builder pattern


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


## AppsDevicesApiEndShift

> ShiftActionOut AppsDevicesApiEndShift(ctx, deviceId).Execute()

End Shift



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiEndShift(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiEndShift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiEndShift`: ShiftActionOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiEndShift`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiEndShiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShiftActionOut**](ShiftActionOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiExportEventsEndpoint

> map[string]interface{} AppsDevicesApiExportEventsEndpoint(ctx).Format(format).FromDate(fromDate).ToDate(toDate).DeviceIds(deviceIds).GeofenceIds(geofenceIds).EventType(eventType).Execute()

Export Events Endpoint



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
	format := "format_example" // string | 
	fromDate := "fromDate_example" // string |  (optional)
	toDate := "toDate_example" // string |  (optional)
	deviceIds := "deviceIds_example" // string |  (optional)
	geofenceIds := "geofenceIds_example" // string |  (optional)
	eventType := "eventType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiExportEventsEndpoint(context.Background()).Format(format).FromDate(fromDate).ToDate(toDate).DeviceIds(deviceIds).GeofenceIds(geofenceIds).EventType(eventType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiExportEventsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiExportEventsEndpoint`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiExportEventsEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiExportEventsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 
 **deviceIds** | **string** |  | 
 **geofenceIds** | **string** |  | 
 **eventType** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetActiveSession

> DeviceSessionOut AppsDevicesApiGetActiveSession(ctx, deviceUuid).Execute()

Get Active Session



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
	deviceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetActiveSession(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetActiveSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetActiveSession`: DeviceSessionOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetActiveSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetActiveSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceSessionOut**](DeviceSessionOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetDashboardStats

> DashboardStatsOut AppsDevicesApiGetDashboardStats(ctx).Execute()

Get Dashboard Stats



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
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetDashboardStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetDashboardStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetDashboardStats`: DashboardStatsOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetDashboardStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetDashboardStatsRequest struct via the builder pattern


### Return type

[**DashboardStatsOut**](DashboardStatsOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetDashboardStatsTimeline

> DashboardStatsTimelineOut AppsDevicesApiGetDashboardStatsTimeline(ctx).TimeRange(timeRange).StartDate(startDate).EndDate(endDate).Execute()

Get Dashboard Stats Timeline



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
	timeRange := "timeRange_example" // string | Time range: today, 7d, 30d, or custom (optional) (default to "today")
	startDate := "startDate_example" // string | Custom start date (ISO 8601) (optional)
	endDate := "endDate_example" // string | Custom end date (ISO 8601) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetDashboardStatsTimeline(context.Background()).TimeRange(timeRange).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetDashboardStatsTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetDashboardStatsTimeline`: DashboardStatsTimelineOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetDashboardStatsTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetDashboardStatsTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeRange** | **string** | Time range: today, 7d, 30d, or custom | [default to &quot;today&quot;]
 **startDate** | **string** | Custom start date (ISO 8601) | 
 **endDate** | **string** | Custom end date (ISO 8601) | 

### Return type

[**DashboardStatsTimelineOut**](DashboardStatsTimelineOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetDevice

> DeviceOut AppsDevicesApiGetDevice(ctx, deviceId).Execute()

Get Device



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetDevice`: DeviceOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceOut**](DeviceOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetDeviceEvents

> []GeofenceEventOut AppsDevicesApiGetDeviceEvents(ctx, deviceId).Limit(limit).Offset(offset).Execute()

Get Device Events



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetDeviceEvents(context.Background(), deviceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetDeviceEvents`: []GeofenceEventOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]GeofenceEventOut**](GeofenceEventOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetDeviceSessions

> DeviceSessionsOut AppsDevicesApiGetDeviceSessions(ctx, deviceId).Limit(limit).Offset(offset).Execute()

Get Device Sessions



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 |  (optional) (default to 20)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetDeviceSessions(context.Background(), deviceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetDeviceSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetDeviceSessions`: DeviceSessionsOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetDeviceSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetDeviceSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**DeviceSessionsOut**](DeviceSessionsOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetEventDetail

> map[string]interface{} AppsDevicesApiGetEventDetail(ctx, eventId).Execute()

Get Event Detail



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetEventDetail(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetEventDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetEventDetail`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetEventDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetEventDetailRequest struct via the builder pattern


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


## AppsDevicesApiGetImportJob

> LocationImportResponse AppsDevicesApiGetImportJob(ctx, jobId).Execute()

Get Import Job



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetImportJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetImportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetImportJob`: LocationImportResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetImportJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetImportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocationImportResponse**](LocationImportResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetLocationStats

> DeviceStatsOut AppsDevicesApiGetLocationStats(ctx).Execute()

Get Location Stats



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
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetLocationStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetLocationStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetLocationStats`: DeviceStatsOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetLocationStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetLocationStatsRequest struct via the builder pattern


### Return type

[**DeviceStatsOut**](DeviceStatsOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetRecentEvents

> RecentEventsOut AppsDevicesApiGetRecentEvents(ctx).Limit(limit).Offset(offset).DeviceId(deviceId).GeofenceId(geofenceId).EventType(eventType).TimeRange(timeRange).StartDate(startDate).EndDate(endDate).Sort(sort).Execute()

Get Recent Events



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
	deviceId := "deviceId_example" // string |  (optional)
	geofenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	eventType := "eventType_example" // string |  (optional)
	timeRange := "timeRange_example" // string | Time range: today, 7d, 30d, or custom (optional)
	startDate := "startDate_example" // string | Custom start date (YYYY-MM-DD) (optional)
	endDate := "endDate_example" // string | Custom end date (YYYY-MM-DD) (optional)
	sort := "sort_example" // string | Sort order: -timestamp (default) or timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetRecentEvents(context.Background()).Limit(limit).Offset(offset).DeviceId(deviceId).GeofenceId(geofenceId).EventType(eventType).TimeRange(timeRange).StartDate(startDate).EndDate(endDate).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetRecentEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetRecentEvents`: RecentEventsOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetRecentEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetRecentEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **deviceId** | **string** |  | 
 **geofenceId** | **string** |  | 
 **eventType** | **string** |  | 
 **timeRange** | **string** | Time range: today, 7d, 30d, or custom | 
 **startDate** | **string** | Custom start date (YYYY-MM-DD) | 
 **endDate** | **string** | Custom end date (YYYY-MM-DD) | 
 **sort** | **string** | Sort order: -timestamp (default) or timestamp | 

### Return type

[**RecentEventsOut**](RecentEventsOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetRecentLocations

> RecentLocationsOut AppsDevicesApiGetRecentLocations(ctx, deviceId).Limit(limit).Execute()

Get Recent Locations



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 |  (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetRecentLocations(context.Background(), deviceId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetRecentLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetRecentLocations`: RecentLocationsOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetRecentLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetRecentLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 500]

### Return type

[**RecentLocationsOut**](RecentLocationsOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetSessionDetail

> DeviceSessionOut AppsDevicesApiGetSessionDetail(ctx, deviceId, sessionId).Execute()

Get Session Detail



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetSessionDetail(context.Background(), deviceId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetSessionDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetSessionDetail`: DeviceSessionOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetSessionDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetSessionDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceSessionOut**](DeviceSessionOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiGetSessionLocations

> SessionLocationsOut AppsDevicesApiGetSessionLocations(ctx, deviceId, sessionId).Limit(limit).Offset(offset).MaxPoints(maxPoints).Execute()

Get Session Locations



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 |  (optional) (default to 1000)
	offset := int32(56) // int32 |  (optional) (default to 0)
	maxPoints := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiGetSessionLocations(context.Background(), deviceId, sessionId).Limit(limit).Offset(offset).MaxPoints(maxPoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiGetSessionLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiGetSessionLocations`: SessionLocationsOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiGetSessionLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiGetSessionLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | [default to 1000]
 **offset** | **int32** |  | [default to 0]
 **maxPoints** | **int32** |  | 

### Return type

[**SessionLocationsOut**](SessionLocationsOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiListDevices

> []DeviceOut AppsDevicesApiListDevices(ctx).IsActive(isActive).IncludeGeofences(includeGeofences).Execute()

List Devices



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
	isActive := true // bool |  (optional)
	includeGeofences := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiListDevices(context.Background()).IsActive(isActive).IncludeGeofences(includeGeofences).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiListDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiListDevices`: []DeviceOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiListDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiListDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isActive** | **bool** |  | 
 **includeGeofences** | **bool** |  | [default to false]

### Return type

[**[]DeviceOut**](DeviceOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiListSessionAttachments

> []StoredFileAttachmentOut AppsDevicesApiListSessionAttachments(ctx, deviceUuid, sessionId).Execute()

List Session Attachments



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
	deviceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiListSessionAttachments(context.Background(), deviceUuid, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiListSessionAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiListSessionAttachments`: []StoredFileAttachmentOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiListSessionAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiListSessionAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]StoredFileAttachmentOut**](StoredFileAttachmentOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiListSessionNotes

> []SessionNoteOut AppsDevicesApiListSessionNotes(ctx, deviceUuid, sessionId).Execute()

List Session Notes



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
	deviceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiListSessionNotes(context.Background(), deviceUuid, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiListSessionNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiListSessionNotes`: []SessionNoteOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiListSessionNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiListSessionNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SessionNoteOut**](SessionNoteOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiListSessionPhotos

> []PhotoOut AppsDevicesApiListSessionPhotos(ctx, deviceUuid, sessionId).Execute()

List Session Photos



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
	deviceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiListSessionPhotos(context.Background(), deviceUuid, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiListSessionPhotos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiListSessionPhotos`: []PhotoOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiListSessionPhotos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiListSessionPhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PhotoOut**](PhotoOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiListWorkspacePhotos

> WorkspacePhotosOut AppsDevicesApiListWorkspacePhotos(ctx).From(from).To(to).DeviceIds(deviceIds).Limit(limit).Execute()

List Workspace Photos



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
	from := time.Now() // time.Time | Earliest captured_at (default: now - 24h) (optional)
	to := time.Now() // time.Time | Latest captured_at (default: now) (optional)
	deviceIds := []*string{"Inner_example"} // []*string | Restrict to these device IDs (max 50) (optional)
	limit := int32(56) // int32 | Soft cap; max 500 newest (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiListWorkspacePhotos(context.Background()).From(from).To(to).DeviceIds(deviceIds).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiListWorkspacePhotos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiListWorkspacePhotos`: WorkspacePhotosOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiListWorkspacePhotos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiListWorkspacePhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | Earliest captured_at (default: now - 24h) | 
 **to** | **time.Time** | Latest captured_at (default: now) | 
 **deviceIds** | **[]string** | Restrict to these device IDs (max 50) | 
 **limit** | **int32** | Soft cap; max 500 newest | [default to 500]

### Return type

[**WorkspacePhotosOut**](WorkspacePhotosOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiPauseShift

> ShiftActionOut AppsDevicesApiPauseShift(ctx, deviceId).Execute()

Pause Shift



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiPauseShift(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiPauseShift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiPauseShift`: ShiftActionOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiPauseShift`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiPauseShiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShiftActionOut**](ShiftActionOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiResumeShift

> ShiftActionOut AppsDevicesApiResumeShift(ctx, deviceId).Execute()

Resume Shift



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiResumeShift(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiResumeShift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiResumeShift`: ShiftActionOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiResumeShift`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiResumeShiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShiftActionOut**](ShiftActionOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiStartShift

> ShiftActionOut AppsDevicesApiStartShift(ctx, deviceId).Execute()

Start Shift



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiStartShift(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiStartShift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiStartShift`: ShiftActionOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiStartShift`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiStartShiftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShiftActionOut**](ShiftActionOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiUpdateDevice

> DeviceOut AppsDevicesApiUpdateDevice(ctx, deviceId).UpdateDeviceIn(updateDeviceIn).Execute()

Update Device



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateDeviceIn := *openapiclient.NewUpdateDeviceIn() // UpdateDeviceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiUpdateDevice(context.Background(), deviceId).UpdateDeviceIn(updateDeviceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiUpdateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiUpdateDevice`: DeviceOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiUpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceIn** | [**UpdateDeviceIn**](UpdateDeviceIn.md) |  | 

### Return type

[**DeviceOut**](DeviceOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiUpdateDeviceLocation

> LocationUpdateOut AppsDevicesApiUpdateDeviceLocation(ctx, deviceId).LocationUpdateIn(locationUpdateIn).Execute()

Update Device Location



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	locationUpdateIn := *openapiclient.NewLocationUpdateIn(float32(123), float32(123)) // LocationUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiUpdateDeviceLocation(context.Background(), deviceId).LocationUpdateIn(locationUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiUpdateDeviceLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiUpdateDeviceLocation`: LocationUpdateOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiUpdateDeviceLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiUpdateDeviceLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationUpdateIn** | [**LocationUpdateIn**](LocationUpdateIn.md) |  | 

### Return type

[**LocationUpdateOut**](LocationUpdateOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiUpdateSessionNotes

> NotesUpdateOut AppsDevicesApiUpdateSessionNotes(ctx, deviceId).NotesUpdateIn(notesUpdateIn).Execute()

Update Session Notes



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	notesUpdateIn := *openapiclient.NewNotesUpdateIn("Notes_example") // NotesUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiUpdateSessionNotes(context.Background(), deviceId).NotesUpdateIn(notesUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiUpdateSessionNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiUpdateSessionNotes`: NotesUpdateOut
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiUpdateSessionNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiUpdateSessionNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notesUpdateIn** | [**NotesUpdateIn**](NotesUpdateIn.md) |  | 

### Return type

[**NotesUpdateOut**](NotesUpdateOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiUploadCsvImport

> LocationImportResponse AppsDevicesApiUploadCsvImport(ctx).File(file).Execute()

Upload Csv Import



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
	resp, r, err := apiClient.DevicesAPI.AppsDevicesApiUploadCsvImport(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AppsDevicesApiUploadCsvImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiUploadCsvImport`: LocationImportResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.AppsDevicesApiUploadCsvImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiUploadCsvImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**LocationImportResponse**](LocationImportResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

