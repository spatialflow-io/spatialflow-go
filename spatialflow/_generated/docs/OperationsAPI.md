# \OperationsAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsOperationsApiAcknowledgeIncident**](OperationsAPI.md#AppsOperationsApiAcknowledgeIncident) | **Post** /api/v1/incidents/{incident_id}/ack | Acknowledge Incident
[**AppsOperationsApiBulkUpdateIncidents**](OperationsAPI.md#AppsOperationsApiBulkUpdateIncidents) | **Post** /api/v1/incidents/bulk | Bulk Update Incidents
[**AppsOperationsApiCreateSavedView**](OperationsAPI.md#AppsOperationsApiCreateSavedView) | **Post** /api/v1/saved-views/ | Create Saved View
[**AppsOperationsApiDeleteSavedView**](OperationsAPI.md#AppsOperationsApiDeleteSavedView) | **Delete** /api/v1/saved-views/{view_id} | Delete Saved View
[**AppsOperationsApiGetIncident**](OperationsAPI.md#AppsOperationsApiGetIncident) | **Get** /api/v1/incidents/{incident_id} | Get Incident
[**AppsOperationsApiGetOverview**](OperationsAPI.md#AppsOperationsApiGetOverview) | **Get** /api/v1/overview | Get Overview
[**AppsOperationsApiListIncidentGroups**](OperationsAPI.md#AppsOperationsApiListIncidentGroups) | **Get** /api/v1/incidents/groups | List Incident Groups
[**AppsOperationsApiListIncidents**](OperationsAPI.md#AppsOperationsApiListIncidents) | **Get** /api/v1/incidents/ | List Incidents
[**AppsOperationsApiListSavedViews**](OperationsAPI.md#AppsOperationsApiListSavedViews) | **Get** /api/v1/saved-views/ | List Saved Views
[**AppsOperationsApiMuteIncident**](OperationsAPI.md#AppsOperationsApiMuteIncident) | **Post** /api/v1/incidents/{incident_id}/mute | Mute Incident
[**AppsOperationsApiResolveIncident**](OperationsAPI.md#AppsOperationsApiResolveIncident) | **Post** /api/v1/incidents/{incident_id}/resolve | Resolve Incident
[**AppsOperationsApiStreamBadges**](OperationsAPI.md#AppsOperationsApiStreamBadges) | **Get** /api/v1/stream/badges | Stream Badges
[**AppsOperationsApiUpdateSavedView**](OperationsAPI.md#AppsOperationsApiUpdateSavedView) | **Patch** /api/v1/saved-views/{view_id} | Update Saved View



## AppsOperationsApiAcknowledgeIncident

> IncidentOut AppsOperationsApiAcknowledgeIncident(ctx, incidentId).IncidentAckIn(incidentAckIn).Execute()

Acknowledge Incident

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
	incidentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	incidentAckIn := *openapiclient.NewIncidentAckIn() // IncidentAckIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiAcknowledgeIncident(context.Background(), incidentId).IncidentAckIn(incidentAckIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiAcknowledgeIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiAcknowledgeIncident`: IncidentOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiAcknowledgeIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiAcknowledgeIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **incidentAckIn** | [**IncidentAckIn**](IncidentAckIn.md) |  | 

### Return type

[**IncidentOut**](IncidentOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiBulkUpdateIncidents

> IncidentBulkOut AppsOperationsApiBulkUpdateIncidents(ctx).IncidentBulkIn(incidentBulkIn).Execute()

Bulk Update Incidents



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
	incidentBulkIn := *openapiclient.NewIncidentBulkIn("Action_example", time.Now()) // IncidentBulkIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiBulkUpdateIncidents(context.Background()).IncidentBulkIn(incidentBulkIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiBulkUpdateIncidents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiBulkUpdateIncidents`: IncidentBulkOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiBulkUpdateIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiBulkUpdateIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentBulkIn** | [**IncidentBulkIn**](IncidentBulkIn.md) |  | 

### Return type

[**IncidentBulkOut**](IncidentBulkOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiCreateSavedView

> SavedViewOut AppsOperationsApiCreateSavedView(ctx).SavedViewIn(savedViewIn).Execute()

Create Saved View

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
	savedViewIn := *openapiclient.NewSavedViewIn("Surface_example", "Name_example", "Url_example") // SavedViewIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiCreateSavedView(context.Background()).SavedViewIn(savedViewIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiCreateSavedView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiCreateSavedView`: SavedViewOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiCreateSavedView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiCreateSavedViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **savedViewIn** | [**SavedViewIn**](SavedViewIn.md) |  | 

### Return type

[**SavedViewOut**](SavedViewOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiDeleteSavedView

> SavedViewDeleteOut AppsOperationsApiDeleteSavedView(ctx, viewId).Execute()

Delete Saved View

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
	viewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiDeleteSavedView(context.Background(), viewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiDeleteSavedView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiDeleteSavedView`: SavedViewDeleteOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiDeleteSavedView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiDeleteSavedViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SavedViewDeleteOut**](SavedViewDeleteOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiGetIncident

> IncidentOut AppsOperationsApiGetIncident(ctx, incidentId).Execute()

Get Incident

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
	incidentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiGetIncident(context.Background(), incidentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiGetIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiGetIncident`: IncidentOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiGetIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiGetIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IncidentOut**](IncidentOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiGetOverview

> OverviewOut AppsOperationsApiGetOverview(ctx).Execute()

Get Overview

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
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiGetOverview(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiGetOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiGetOverview`: OverviewOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiGetOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiGetOverviewRequest struct via the builder pattern


### Return type

[**OverviewOut**](OverviewOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiListIncidentGroups

> IncidentGroupListOut AppsOperationsApiListIncidentGroups(ctx).Status(status).Severity(severity).Execute()

List Incident Groups



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
	status := "status_example" // string |  (optional) (default to "open")
	severity := "severity_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiListIncidentGroups(context.Background()).Status(status).Severity(severity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiListIncidentGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiListIncidentGroups`: IncidentGroupListOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiListIncidentGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiListIncidentGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | [default to &quot;open&quot;]
 **severity** | **string** |  | 

### Return type

[**IncidentGroupListOut**](IncidentGroupListOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiListIncidents

> IncidentListOut AppsOperationsApiListIncidents(ctx).Status(status).Severity(severity).Owner(owner).SignalType(signalType).Cursor(cursor).Limit(limit).Execute()

List Incidents

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
	status := "status_example" // string |  (optional) (default to "open")
	severity := "severity_example" // string |  (optional)
	owner := "owner_example" // string |  (optional)
	signalType := "signalType_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiListIncidents(context.Background()).Status(status).Severity(severity).Owner(owner).SignalType(signalType).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiListIncidents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiListIncidents`: IncidentListOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiListIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiListIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | [default to &quot;open&quot;]
 **severity** | **string** |  | 
 **owner** | **string** |  | 
 **signalType** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**IncidentListOut**](IncidentListOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiListSavedViews

> []SavedViewOut AppsOperationsApiListSavedViews(ctx).Surface(surface).Execute()

List Saved Views

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
	surface := "surface_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiListSavedViews(context.Background()).Surface(surface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiListSavedViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiListSavedViews`: []SavedViewOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiListSavedViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiListSavedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **surface** | **string** |  | 

### Return type

[**[]SavedViewOut**](SavedViewOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiMuteIncident

> IncidentOut AppsOperationsApiMuteIncident(ctx, incidentId).IncidentMuteIn(incidentMuteIn).Execute()

Mute Incident

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
	incidentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	incidentMuteIn := *openapiclient.NewIncidentMuteIn(int32(123)) // IncidentMuteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiMuteIncident(context.Background(), incidentId).IncidentMuteIn(incidentMuteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiMuteIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiMuteIncident`: IncidentOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiMuteIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiMuteIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **incidentMuteIn** | [**IncidentMuteIn**](IncidentMuteIn.md) |  | 

### Return type

[**IncidentOut**](IncidentOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiResolveIncident

> IncidentOut AppsOperationsApiResolveIncident(ctx, incidentId).Execute()

Resolve Incident

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
	incidentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiResolveIncident(context.Background(), incidentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiResolveIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiResolveIncident`: IncidentOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiResolveIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiResolveIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IncidentOut**](IncidentOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiStreamBadges

> string AppsOperationsApiStreamBadges(ctx).Execute()

Stream Badges

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
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiStreamBadges(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiStreamBadges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiStreamBadges`: string
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiStreamBadges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiStreamBadgesRequest struct via the builder pattern


### Return type

**string**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsOperationsApiUpdateSavedView

> SavedViewOut AppsOperationsApiUpdateSavedView(ctx, viewId).SavedViewPatchIn(savedViewPatchIn).Execute()

Update Saved View

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
	viewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	savedViewPatchIn := *openapiclient.NewSavedViewPatchIn() // SavedViewPatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.AppsOperationsApiUpdateSavedView(context.Background(), viewId).SavedViewPatchIn(savedViewPatchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.AppsOperationsApiUpdateSavedView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsOperationsApiUpdateSavedView`: SavedViewOut
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.AppsOperationsApiUpdateSavedView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsOperationsApiUpdateSavedViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **savedViewPatchIn** | [**SavedViewPatchIn**](SavedViewPatchIn.md) |  | 

### Return type

[**SavedViewOut**](SavedViewOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

