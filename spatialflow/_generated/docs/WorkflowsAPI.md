# \WorkflowsAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsWorkflowsApiActivateWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiActivateWorkflow) | **Post** /api/v1/workflows/{workflow_id}/activate | Activate Workflow
[**AppsWorkflowsApiCreateWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiCreateWorkflow) | **Post** /api/v1/workflows | Create Workflow
[**AppsWorkflowsApiCreateWorkflowFromTemplate**](WorkflowsAPI.md#AppsWorkflowsApiCreateWorkflowFromTemplate) | **Post** /api/v1/workflows/templates/{template_id}/use | Create Workflow From Template
[**AppsWorkflowsApiDeleteWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiDeleteWorkflow) | **Delete** /api/v1/workflows/{workflow_id} | Delete Workflow
[**AppsWorkflowsApiDuplicateWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiDuplicateWorkflow) | **Post** /api/v1/workflows/{workflow_id}/duplicate | Duplicate Workflow
[**AppsWorkflowsApiExecuteWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiExecuteWorkflow) | **Post** /api/v1/workflows/{workflow_id}/execute | Execute Workflow
[**AppsWorkflowsApiExportWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiExportWorkflow) | **Get** /api/v1/workflows/{workflow_id}/export | Export Workflow
[**AppsWorkflowsApiGetAvailableRetryPolicies**](WorkflowsAPI.md#AppsWorkflowsApiGetAvailableRetryPolicies) | **Get** /api/v1/workflows/retry-policies | Get Available Retry Policies
[**AppsWorkflowsApiGetExecutionDetails**](WorkflowsAPI.md#AppsWorkflowsApiGetExecutionDetails) | **Get** /api/v1/workflows/executions/{execution_id} | Get Execution Details
[**AppsWorkflowsApiGetSystemPerformanceSummary**](WorkflowsAPI.md#AppsWorkflowsApiGetSystemPerformanceSummary) | **Get** /api/v1/workflows/performance/summary | Get System Performance Summary
[**AppsWorkflowsApiGetVersionDiff**](WorkflowsAPI.md#AppsWorkflowsApiGetVersionDiff) | **Get** /api/v1/workflows/{workflow_id}/versions/{version_number}/diff | Get Version Diff
[**AppsWorkflowsApiGetWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflow) | **Get** /api/v1/workflows/{workflow_id} | Get Workflow
[**AppsWorkflowsApiGetWorkflowBottlenecks**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflowBottlenecks) | **Get** /api/v1/workflows/{workflow_id}/performance/bottlenecks | Get Workflow Bottlenecks
[**AppsWorkflowsApiGetWorkflowExecutionDetail**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflowExecutionDetail) | **Get** /api/v1/workflows/{workflow_id}/executions/{execution_id} | Get Workflow Execution Detail
[**AppsWorkflowsApiGetWorkflowExecutions**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflowExecutions) | **Get** /api/v1/workflows/{workflow_id}/executions | Get Workflow Executions
[**AppsWorkflowsApiGetWorkflowPerformance**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflowPerformance) | **Get** /api/v1/workflows/{workflow_id}/performance | Get Workflow Performance
[**AppsWorkflowsApiGetWorkflowRetryPolicy**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflowRetryPolicy) | **Get** /api/v1/workflows/{workflow_id}/retry-policy | Get Workflow Retry Policy
[**AppsWorkflowsApiGetWorkflowStatistics**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflowStatistics) | **Get** /api/v1/workflows/{workflow_id}/statistics | Get Workflow Statistics
[**AppsWorkflowsApiGetWorkflowStepPerformance**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflowStepPerformance) | **Get** /api/v1/workflows/{workflow_id}/performance/steps | Get Workflow Step Performance
[**AppsWorkflowsApiGetWorkflowTemplate**](WorkflowsAPI.md#AppsWorkflowsApiGetWorkflowTemplate) | **Get** /api/v1/workflows/templates/{template_id} | Get Workflow Template
[**AppsWorkflowsApiHealthCheck**](WorkflowsAPI.md#AppsWorkflowsApiHealthCheck) | **Get** /api/v1/workflows/health | Health Check
[**AppsWorkflowsApiImportWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiImportWorkflow) | **Post** /api/v1/workflows/import | Import Workflow
[**AppsWorkflowsApiListWorkflowExecutions**](WorkflowsAPI.md#AppsWorkflowsApiListWorkflowExecutions) | **Get** /api/v1/workflows/executions | List Workflow Executions
[**AppsWorkflowsApiListWorkflowTemplates**](WorkflowsAPI.md#AppsWorkflowsApiListWorkflowTemplates) | **Get** /api/v1/workflows/templates | List Workflow Templates
[**AppsWorkflowsApiListWorkflowVersions**](WorkflowsAPI.md#AppsWorkflowsApiListWorkflowVersions) | **Get** /api/v1/workflows/{workflow_id}/versions | List Workflow Versions
[**AppsWorkflowsApiListWorkflows**](WorkflowsAPI.md#AppsWorkflowsApiListWorkflows) | **Get** /api/v1/workflows | List Workflows
[**AppsWorkflowsApiRestoreWorkflowVersion**](WorkflowsAPI.md#AppsWorkflowsApiRestoreWorkflowVersion) | **Post** /api/v1/workflows/{workflow_id}/versions/{version_number}/restore | Restore Workflow Version
[**AppsWorkflowsApiTestWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiTestWorkflow) | **Post** /api/v1/workflows/{workflow_id}/test | Test Workflow
[**AppsWorkflowsApiToggleWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiToggleWorkflow) | **Post** /api/v1/workflows/{workflow_id}/toggle | Toggle Workflow
[**AppsWorkflowsApiUpdateWorkflow**](WorkflowsAPI.md#AppsWorkflowsApiUpdateWorkflow) | **Put** /api/v1/workflows/{workflow_id} | Update Workflow
[**AppsWorkflowsApiUpdateWorkflowRetryPolicy**](WorkflowsAPI.md#AppsWorkflowsApiUpdateWorkflowRetryPolicy) | **Put** /api/v1/workflows/{workflow_id}/retry-policy | Update Workflow Retry Policy



## AppsWorkflowsApiActivateWorkflow

> WorkflowOut AppsWorkflowsApiActivateWorkflow(ctx, workflowId).Execute()

Activate Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiActivateWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiActivateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiActivateWorkflow`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiActivateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiActivateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiCreateWorkflow

> WorkflowOut AppsWorkflowsApiCreateWorkflow(ctx).WorkflowIn(workflowIn).Execute()

Create Workflow



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
	workflowIn := *openapiclient.NewWorkflowIn("Name_example", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // WorkflowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiCreateWorkflow(context.Background()).WorkflowIn(workflowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiCreateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiCreateWorkflow`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiCreateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowIn** | [**WorkflowIn**](WorkflowIn.md) |  | 

### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiCreateWorkflowFromTemplate

> WorkflowOut AppsWorkflowsApiCreateWorkflowFromTemplate(ctx, templateId).CreateFromTemplateIn(createFromTemplateIn).Execute()

Create Workflow From Template



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
	templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createFromTemplateIn := *openapiclient.NewCreateFromTemplateIn("Name_example") // CreateFromTemplateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiCreateWorkflowFromTemplate(context.Background(), templateId).CreateFromTemplateIn(createFromTemplateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiCreateWorkflowFromTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiCreateWorkflowFromTemplate`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiCreateWorkflowFromTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiCreateWorkflowFromTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFromTemplateIn** | [**CreateFromTemplateIn**](CreateFromTemplateIn.md) |  | 

### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiDeleteWorkflow

> map[string]interface{} AppsWorkflowsApiDeleteWorkflow(ctx, workflowId).Execute()

Delete Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiDeleteWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiDeleteWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiDeleteWorkflow`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiDeleteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiDeleteWorkflowRequest struct via the builder pattern


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


## AppsWorkflowsApiDuplicateWorkflow

> WorkflowOut AppsWorkflowsApiDuplicateWorkflow(ctx, workflowId).Name(name).Execute()

Duplicate Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiDuplicateWorkflow(context.Background(), workflowId).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiDuplicateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiDuplicateWorkflow`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiDuplicateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiDuplicateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 

### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiExecuteWorkflow

> map[string]interface{} AppsWorkflowsApiExecuteWorkflow(ctx, workflowId).TestData(testData).Execute()

Execute Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testData := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiExecuteWorkflow(context.Background(), workflowId).TestData(testData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiExecuteWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiExecuteWorkflow`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiExecuteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiExecuteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testData** | **map[string]interface{}** |  | 

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


## AppsWorkflowsApiExportWorkflow

> map[string]interface{} AppsWorkflowsApiExportWorkflow(ctx, workflowId).Execute()

Export Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiExportWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiExportWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiExportWorkflow`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiExportWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiExportWorkflowRequest struct via the builder pattern


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


## AppsWorkflowsApiGetAvailableRetryPolicies

> RetryPolicyResponseSchema AppsWorkflowsApiGetAvailableRetryPolicies(ctx).Execute()

Get Available Retry Policies



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
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetAvailableRetryPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetAvailableRetryPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetAvailableRetryPolicies`: RetryPolicyResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetAvailableRetryPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetAvailableRetryPoliciesRequest struct via the builder pattern


### Return type

[**RetryPolicyResponseSchema**](RetryPolicyResponseSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiGetExecutionDetails

> map[string]interface{} AppsWorkflowsApiGetExecutionDetails(ctx, executionId).Execute()

Get Execution Details



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
	executionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetExecutionDetails(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetExecutionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetExecutionDetails`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetExecutionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetExecutionDetailsRequest struct via the builder pattern


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


## AppsWorkflowsApiGetSystemPerformanceSummary

> map[string]interface{} AppsWorkflowsApiGetSystemPerformanceSummary(ctx).Execute()

Get System Performance Summary



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
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetSystemPerformanceSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetSystemPerformanceSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetSystemPerformanceSummary`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetSystemPerformanceSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetSystemPerformanceSummaryRequest struct via the builder pattern


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


## AppsWorkflowsApiGetVersionDiff

> map[string]interface{} AppsWorkflowsApiGetVersionDiff(ctx, workflowId, versionNumber).Execute()

Get Version Diff



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	versionNumber := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetVersionDiff(context.Background(), workflowId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetVersionDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetVersionDiff`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetVersionDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**versionNumber** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetVersionDiffRequest struct via the builder pattern


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


## AppsWorkflowsApiGetWorkflow

> WorkflowOut AppsWorkflowsApiGetWorkflow(ctx, workflowId).Execute()

Get Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflow`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiGetWorkflowBottlenecks

> []map[string]interface{} AppsWorkflowsApiGetWorkflowBottlenecks(ctx, workflowId).Execute()

Get Workflow Bottlenecks



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflowBottlenecks(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflowBottlenecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflowBottlenecks`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflowBottlenecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowBottlenecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiGetWorkflowExecutionDetail

> map[string]interface{} AppsWorkflowsApiGetWorkflowExecutionDetail(ctx, workflowId, executionId).Execute()

Get Workflow Execution Detail



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	executionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflowExecutionDetail(context.Background(), workflowId, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflowExecutionDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflowExecutionDetail`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflowExecutionDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowExecutionDetailRequest struct via the builder pattern


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


## AppsWorkflowsApiGetWorkflowExecutions

> []ExecutionOut AppsWorkflowsApiGetWorkflowExecutions(ctx, workflowId).Status(status).Limit(limit).Offset(offset).Execute()

Get Workflow Executions



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflowExecutions(context.Background(), workflowId).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflowExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflowExecutions`: []ExecutionOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflowExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]ExecutionOut**](ExecutionOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiGetWorkflowPerformance

> map[string]interface{} AppsWorkflowsApiGetWorkflowPerformance(ctx, workflowId).StartDate(startDate).EndDate(endDate).Execute()

Get Workflow Performance



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflowPerformance(context.Background(), workflowId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflowPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflowPerformance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflowPerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

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


## AppsWorkflowsApiGetWorkflowRetryPolicy

> map[string]interface{} AppsWorkflowsApiGetWorkflowRetryPolicy(ctx, workflowId).Execute()

Get Workflow Retry Policy



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflowRetryPolicy(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflowRetryPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflowRetryPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflowRetryPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowRetryPolicyRequest struct via the builder pattern


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


## AppsWorkflowsApiGetWorkflowStatistics

> map[string]interface{} AppsWorkflowsApiGetWorkflowStatistics(ctx, workflowId).Execute()

Get Workflow Statistics



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflowStatistics(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflowStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflowStatistics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflowStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowStatisticsRequest struct via the builder pattern


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


## AppsWorkflowsApiGetWorkflowStepPerformance

> []map[string]interface{} AppsWorkflowsApiGetWorkflowStepPerformance(ctx, workflowId).Execute()

Get Workflow Step Performance



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflowStepPerformance(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflowStepPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflowStepPerformance`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflowStepPerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowStepPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiGetWorkflowTemplate

> map[string]interface{} AppsWorkflowsApiGetWorkflowTemplate(ctx, templateId).Execute()

Get Workflow Template



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
	templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiGetWorkflowTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiGetWorkflowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiGetWorkflowTemplate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiGetWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiGetWorkflowTemplateRequest struct via the builder pattern


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


## AppsWorkflowsApiHealthCheck

> map[string]interface{} AppsWorkflowsApiHealthCheck(ctx).Execute()

Health Check



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
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiHealthCheckRequest struct via the builder pattern


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


## AppsWorkflowsApiImportWorkflow

> WorkflowOut AppsWorkflowsApiImportWorkflow(ctx).WorkflowImportSchema(workflowImportSchema).Execute()

Import Workflow



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
	workflowImportSchema := *openapiclient.NewWorkflowImportSchema("Version_example", *openapiclient.NewWorkflowImportDataSchema("Name_example")) // WorkflowImportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiImportWorkflow(context.Background()).WorkflowImportSchema(workflowImportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiImportWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiImportWorkflow`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiImportWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiImportWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowImportSchema** | [**WorkflowImportSchema**](WorkflowImportSchema.md) |  | 

### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiListWorkflowExecutions

> []ExecutionOut AppsWorkflowsApiListWorkflowExecutions(ctx).WorkflowId(workflowId).Status(status).Limit(limit).Offset(offset).Execute()

List Workflow Executions



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiListWorkflowExecutions(context.Background()).WorkflowId(workflowId).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiListWorkflowExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiListWorkflowExecutions`: []ExecutionOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiListWorkflowExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiListWorkflowExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowId** | **string** |  | 
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]ExecutionOut**](ExecutionOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiListWorkflowTemplates

> []TemplateOut AppsWorkflowsApiListWorkflowTemplates(ctx).Category(category).Lane(lane).Execute()

List Workflow Templates



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
	category := "category_example" // string |  (optional)
	lane := "lane_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiListWorkflowTemplates(context.Background()).Category(category).Lane(lane).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiListWorkflowTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiListWorkflowTemplates`: []TemplateOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiListWorkflowTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiListWorkflowTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** |  | 
 **lane** | **string** |  | 

### Return type

[**[]TemplateOut**](TemplateOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiListWorkflowVersions

> []map[string]interface{} AppsWorkflowsApiListWorkflowVersions(ctx, workflowId).Execute()

List Workflow Versions



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiListWorkflowVersions(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiListWorkflowVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiListWorkflowVersions`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiListWorkflowVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiListWorkflowVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiListWorkflows

> WorkflowListResponse AppsWorkflowsApiListWorkflows(ctx).Limit(limit).Offset(offset).IsActive(isActive).Category(category).Search(search).Execute()

List Workflows



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
	isActive := true // bool |  (optional)
	category := "category_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiListWorkflows(context.Background()).Limit(limit).Offset(offset).IsActive(isActive).Category(category).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiListWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiListWorkflows`: WorkflowListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiListWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiListWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **isActive** | **bool** |  | 
 **category** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**WorkflowListResponse**](WorkflowListResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiRestoreWorkflowVersion

> WorkflowOut AppsWorkflowsApiRestoreWorkflowVersion(ctx, workflowId, versionNumber).Execute()

Restore Workflow Version



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	versionNumber := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiRestoreWorkflowVersion(context.Background(), workflowId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiRestoreWorkflowVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiRestoreWorkflowVersion`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiRestoreWorkflowVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**versionNumber** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiRestoreWorkflowVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiTestWorkflow

> map[string]interface{} AppsWorkflowsApiTestWorkflow(ctx, workflowId).TestWorkflowIn(testWorkflowIn).Execute()

Test Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testWorkflowIn := *openapiclient.NewTestWorkflowIn(map[string]interface{}{"key": interface{}(123)}) // TestWorkflowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiTestWorkflow(context.Background(), workflowId).TestWorkflowIn(testWorkflowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiTestWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiTestWorkflow`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiTestWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiTestWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testWorkflowIn** | [**TestWorkflowIn**](TestWorkflowIn.md) |  | 

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


## AppsWorkflowsApiToggleWorkflow

> WorkflowOut AppsWorkflowsApiToggleWorkflow(ctx, workflowId).Execute()

Toggle Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiToggleWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiToggleWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiToggleWorkflow`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiToggleWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiToggleWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiUpdateWorkflow

> WorkflowOut AppsWorkflowsApiUpdateWorkflow(ctx, workflowId).WorkflowUpdate(workflowUpdate).Execute()

Update Workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workflowUpdate := *openapiclient.NewWorkflowUpdate() // WorkflowUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiUpdateWorkflow(context.Background(), workflowId).WorkflowUpdate(workflowUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiUpdateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiUpdateWorkflow`: WorkflowOut
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiUpdateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowUpdate** | [**WorkflowUpdate**](WorkflowUpdate.md) |  | 

### Return type

[**WorkflowOut**](WorkflowOut.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkflowsApiUpdateWorkflowRetryPolicy

> map[string]interface{} AppsWorkflowsApiUpdateWorkflowRetryPolicy(ctx, workflowId).WorkflowRetryPolicyUpdateSchema(workflowRetryPolicyUpdateSchema).Execute()

Update Workflow Retry Policy



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workflowRetryPolicyUpdateSchema := *openapiclient.NewWorkflowRetryPolicyUpdateSchema("WorkflowId_example") // WorkflowRetryPolicyUpdateSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AppsWorkflowsApiUpdateWorkflowRetryPolicy(context.Background(), workflowId).WorkflowRetryPolicyUpdateSchema(workflowRetryPolicyUpdateSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AppsWorkflowsApiUpdateWorkflowRetryPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkflowsApiUpdateWorkflowRetryPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AppsWorkflowsApiUpdateWorkflowRetryPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkflowsApiUpdateWorkflowRetryPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowRetryPolicyUpdateSchema** | [**WorkflowRetryPolicyUpdateSchema**](WorkflowRetryPolicyUpdateSchema.md) |  | 

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

