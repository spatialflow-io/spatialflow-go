# \ReportsAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsDevicesApiReportsReportAttendance**](ReportsAPI.md#AppsDevicesApiReportsReportAttendance) | **Get** /api/v1/reports/attendance | Shift attendance report
[**AppsDevicesApiReportsReportDetention**](ReportsAPI.md#AppsDevicesApiReportsReportDetention) | **Get** /api/v1/reports/detention | Detention report — one row per facility visit with billable overage
[**AppsDevicesApiReportsReportIdleStops**](ReportsAPI.md#AppsDevicesApiReportsReportIdleStops) | **Get** /api/v1/reports/idle-stops | Idle/stop summary report
[**AppsDevicesApiReportsReportMileage**](ReportsAPI.md#AppsDevicesApiReportsReportMileage) | **Get** /api/v1/reports/mileage | Mileage summary report
[**AppsDevicesApiReportsReportTimeOnSite**](ReportsAPI.md#AppsDevicesApiReportsReportTimeOnSite) | **Get** /api/v1/reports/time-on-site | Time-on-site / job-site report
[**AppsDevicesApiReportsReportTrips**](ReportsAPI.md#AppsDevicesApiReportsReportTrips) | **Get** /api/v1/reports/trips | Trip history report



## AppsDevicesApiReportsReportAttendance

> ReportOut AppsDevicesApiReportsReportAttendance(ctx).Start(start).End(end).DeviceId(deviceId).Group(group).Format(format).Execute()

Shift attendance report



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
	start := time.Now() // string |  (optional)
	end := time.Now() // string |  (optional)
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	group := "group_example" // string |  (optional)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.AppsDevicesApiReportsReportAttendance(context.Background()).Start(start).End(end).DeviceId(deviceId).Group(group).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AppsDevicesApiReportsReportAttendance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiReportsReportAttendance`: ReportOut
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AppsDevicesApiReportsReportAttendance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiReportsReportAttendanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 
 **deviceId** | **string** |  | 
 **group** | **string** |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**ReportOut**](ReportOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiReportsReportDetention

> ReportOut AppsDevicesApiReportsReportDetention(ctx).Start(start).End(end).DeviceId(deviceId).Group(group).FreeTimeMinutes(freeTimeMinutes).RatePerHourCents(ratePerHourCents).Format(format).Execute()

Detention report — one row per facility visit with billable overage



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
	start := time.Now() // string |  (optional)
	end := time.Now() // string |  (optional)
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	group := "group_example" // string |  (optional)
	freeTimeMinutes := int32(56) // int32 |  (optional) (default to 120)
	ratePerHourCents := int32(56) // int32 |  (optional) (default to 0)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.AppsDevicesApiReportsReportDetention(context.Background()).Start(start).End(end).DeviceId(deviceId).Group(group).FreeTimeMinutes(freeTimeMinutes).RatePerHourCents(ratePerHourCents).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AppsDevicesApiReportsReportDetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiReportsReportDetention`: ReportOut
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AppsDevicesApiReportsReportDetention`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiReportsReportDetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 
 **deviceId** | **string** |  | 
 **group** | **string** |  | 
 **freeTimeMinutes** | **int32** |  | [default to 120]
 **ratePerHourCents** | **int32** |  | [default to 0]
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**ReportOut**](ReportOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiReportsReportIdleStops

> ReportOut AppsDevicesApiReportsReportIdleStops(ctx).Start(start).End(end).DeviceId(deviceId).Group(group).ThresholdMinutes(thresholdMinutes).Format(format).Execute()

Idle/stop summary report



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
	start := time.Now() // string |  (optional)
	end := time.Now() // string |  (optional)
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	group := "group_example" // string |  (optional)
	thresholdMinutes := int32(56) // int32 |  (optional) (default to 5)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.AppsDevicesApiReportsReportIdleStops(context.Background()).Start(start).End(end).DeviceId(deviceId).Group(group).ThresholdMinutes(thresholdMinutes).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AppsDevicesApiReportsReportIdleStops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiReportsReportIdleStops`: ReportOut
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AppsDevicesApiReportsReportIdleStops`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiReportsReportIdleStopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 
 **deviceId** | **string** |  | 
 **group** | **string** |  | 
 **thresholdMinutes** | **int32** |  | [default to 5]
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**ReportOut**](ReportOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiReportsReportMileage

> ReportOut AppsDevicesApiReportsReportMileage(ctx).Start(start).End(end).DeviceId(deviceId).Group(group).Format(format).Execute()

Mileage summary report



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
	start := time.Now() // string |  (optional)
	end := time.Now() // string |  (optional)
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	group := "group_example" // string |  (optional)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.AppsDevicesApiReportsReportMileage(context.Background()).Start(start).End(end).DeviceId(deviceId).Group(group).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AppsDevicesApiReportsReportMileage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiReportsReportMileage`: ReportOut
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AppsDevicesApiReportsReportMileage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiReportsReportMileageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 
 **deviceId** | **string** |  | 
 **group** | **string** |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**ReportOut**](ReportOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiReportsReportTimeOnSite

> ReportOut AppsDevicesApiReportsReportTimeOnSite(ctx).Start(start).End(end).DeviceId(deviceId).Group(group).Format(format).Execute()

Time-on-site / job-site report



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
	start := time.Now() // string |  (optional)
	end := time.Now() // string |  (optional)
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	group := "group_example" // string |  (optional)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.AppsDevicesApiReportsReportTimeOnSite(context.Background()).Start(start).End(end).DeviceId(deviceId).Group(group).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AppsDevicesApiReportsReportTimeOnSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiReportsReportTimeOnSite`: ReportOut
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AppsDevicesApiReportsReportTimeOnSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiReportsReportTimeOnSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 
 **deviceId** | **string** |  | 
 **group** | **string** |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**ReportOut**](ReportOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiReportsReportTrips

> ReportOut AppsDevicesApiReportsReportTrips(ctx).Start(start).End(end).DeviceId(deviceId).Group(group).Format(format).Execute()

Trip history report



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
	start := time.Now() // string |  (optional)
	end := time.Now() // string |  (optional)
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	group := "group_example" // string |  (optional)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.AppsDevicesApiReportsReportTrips(context.Background()).Start(start).End(end).DeviceId(deviceId).Group(group).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AppsDevicesApiReportsReportTrips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiReportsReportTrips`: ReportOut
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AppsDevicesApiReportsReportTrips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiReportsReportTripsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 
 **deviceId** | **string** |  | 
 **group** | **string** |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**ReportOut**](ReportOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

