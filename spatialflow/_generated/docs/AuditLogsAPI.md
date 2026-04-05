# \AuditLogsAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsWorkspacesApiAuditExportAuditLogs**](AuditLogsAPI.md#AppsWorkspacesApiAuditExportAuditLogs) | **Get** /api/v1/audit-logs/export/ | Export Audit Logs
[**AppsWorkspacesApiAuditListAuditLogs**](AuditLogsAPI.md#AppsWorkspacesApiAuditListAuditLogs) | **Get** /api/v1/audit-logs/ | List Audit Logs



## AppsWorkspacesApiAuditExportAuditLogs

> AppsWorkspacesApiAuditExportAuditLogs(ctx).ExportFormat(exportFormat).Action(action).UserId(userId).ResourceType(resourceType).DateFrom(dateFrom).DateTo(dateTo).Search(search).Execute()

Export Audit Logs



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
	exportFormat := "exportFormat_example" // string |  (optional) (default to "csv")
	action := "action_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	resourceType := "resourceType_example" // string |  (optional)
	dateFrom := "dateFrom_example" // string |  (optional)
	dateTo := "dateTo_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuditLogsAPI.AppsWorkspacesApiAuditExportAuditLogs(context.Background()).ExportFormat(exportFormat).Action(action).UserId(userId).ResourceType(resourceType).DateFrom(dateFrom).DateTo(dateTo).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.AppsWorkspacesApiAuditExportAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiAuditExportAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportFormat** | **string** |  | [default to &quot;csv&quot;]
 **action** | **string** |  | 
 **userId** | **string** |  | 
 **resourceType** | **string** |  | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 
 **search** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiAuditListAuditLogs

> AuditLogListResponse AppsWorkspacesApiAuditListAuditLogs(ctx).Action(action).UserId(userId).ResourceType(resourceType).DateFrom(dateFrom).DateTo(dateTo).Search(search).Page(page).PageSize(pageSize).Execute()

List Audit Logs



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
	action := "action_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	resourceType := "resourceType_example" // string |  (optional)
	dateFrom := "dateFrom_example" // string |  (optional)
	dateTo := "dateTo_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.AppsWorkspacesApiAuditListAuditLogs(context.Background()).Action(action).UserId(userId).ResourceType(resourceType).DateFrom(dateFrom).DateTo(dateTo).Search(search).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.AppsWorkspacesApiAuditListAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiAuditListAuditLogs`: AuditLogListResponse
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.AppsWorkspacesApiAuditListAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiAuditListAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** |  | 
 **userId** | **string** |  | 
 **resourceType** | **string** |  | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 
 **search** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]

### Return type

[**AuditLogListResponse**](AuditLogListResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

