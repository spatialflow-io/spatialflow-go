# \IntegrationsAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsIntegrationsApiBulkExportIntegrations**](IntegrationsAPI.md#AppsIntegrationsApiBulkExportIntegrations) | **Post** /api/v1/integrations/bulk-export | Bulk Export Integrations
[**AppsIntegrationsApiBulkImportIntegrations**](IntegrationsAPI.md#AppsIntegrationsApiBulkImportIntegrations) | **Post** /api/v1/integrations/bulk-import | Bulk Import Integrations
[**AppsIntegrationsApiCreateConfigField**](IntegrationsAPI.md#AppsIntegrationsApiCreateConfigField) | **Post** /api/v1/integrations/admin/integration-types/{integration_type_id}/fields | Create Config Field
[**AppsIntegrationsApiCreateIntegration**](IntegrationsAPI.md#AppsIntegrationsApiCreateIntegration) | **Post** /api/v1/integrations/ | Create Integration
[**AppsIntegrationsApiCreateIntegrationType**](IntegrationsAPI.md#AppsIntegrationsApiCreateIntegrationType) | **Post** /api/v1/integrations/admin/integration-types | Create Integration Type
[**AppsIntegrationsApiDeleteConfigField**](IntegrationsAPI.md#AppsIntegrationsApiDeleteConfigField) | **Delete** /api/v1/integrations/admin/integration-types/{integration_type_id}/fields/{field_id} | Delete Config Field
[**AppsIntegrationsApiDeleteIntegration**](IntegrationsAPI.md#AppsIntegrationsApiDeleteIntegration) | **Delete** /api/v1/integrations/{integration_id} | Delete Integration
[**AppsIntegrationsApiDeleteIntegrationType**](IntegrationsAPI.md#AppsIntegrationsApiDeleteIntegrationType) | **Delete** /api/v1/integrations/admin/integration-types/{integration_type_id} | Delete Integration Type
[**AppsIntegrationsApiExportIntegration**](IntegrationsAPI.md#AppsIntegrationsApiExportIntegration) | **Get** /api/v1/integrations/{integration_id}/export | Export Integration
[**AppsIntegrationsApiGetAvailableIntegrationTypes**](IntegrationsAPI.md#AppsIntegrationsApiGetAvailableIntegrationTypes) | **Get** /api/v1/integrations/types/available | Get Available Integration Types
[**AppsIntegrationsApiGetIntegration**](IntegrationsAPI.md#AppsIntegrationsApiGetIntegration) | **Get** /api/v1/integrations/{integration_id} | Get Integration
[**AppsIntegrationsApiGetIntegrationErrorStats**](IntegrationsAPI.md#AppsIntegrationsApiGetIntegrationErrorStats) | **Get** /api/v1/integrations/error-stats | Get Integration Error Stats
[**AppsIntegrationsApiGetIntegrationStats**](IntegrationsAPI.md#AppsIntegrationsApiGetIntegrationStats) | **Get** /api/v1/integrations/{integration_id}/stats | Get Integration Stats
[**AppsIntegrationsApiGetIntegrationType**](IntegrationsAPI.md#AppsIntegrationsApiGetIntegrationType) | **Get** /api/v1/integrations/admin/integration-types/{integration_type_id} | Get Integration Type
[**AppsIntegrationsApiImportIntegration**](IntegrationsAPI.md#AppsIntegrationsApiImportIntegration) | **Post** /api/v1/integrations/import | Import Integration
[**AppsIntegrationsApiListConfigFields**](IntegrationsAPI.md#AppsIntegrationsApiListConfigFields) | **Get** /api/v1/integrations/admin/integration-types/{integration_type_id}/fields | List Config Fields
[**AppsIntegrationsApiListIntegrationTypes**](IntegrationsAPI.md#AppsIntegrationsApiListIntegrationTypes) | **Get** /api/v1/integrations/admin/integration-types | List Integration Types
[**AppsIntegrationsApiListIntegrations**](IntegrationsAPI.md#AppsIntegrationsApiListIntegrations) | **Get** /api/v1/integrations/ | List Integrations
[**AppsIntegrationsApiOauthAuthorize**](IntegrationsAPI.md#AppsIntegrationsApiOauthAuthorize) | **Get** /api/v1/integrations/{integration_type}/oauth/authorize | Oauth Authorize
[**AppsIntegrationsApiOauthCallback**](IntegrationsAPI.md#AppsIntegrationsApiOauthCallback) | **Post** /api/v1/integrations/{integration_type}/oauth/callback | Oauth Callback
[**AppsIntegrationsApiTestAllIntegrations**](IntegrationsAPI.md#AppsIntegrationsApiTestAllIntegrations) | **Post** /api/v1/integrations/test-bulk | Test All Integrations
[**AppsIntegrationsApiTestIntegration**](IntegrationsAPI.md#AppsIntegrationsApiTestIntegration) | **Post** /api/v1/integrations/{integration_id}/test | Test Integration
[**AppsIntegrationsApiUpdateConfigField**](IntegrationsAPI.md#AppsIntegrationsApiUpdateConfigField) | **Put** /api/v1/integrations/admin/integration-types/{integration_type_id}/fields/{field_id} | Update Config Field
[**AppsIntegrationsApiUpdateIntegration**](IntegrationsAPI.md#AppsIntegrationsApiUpdateIntegration) | **Put** /api/v1/integrations/{integration_id} | Update Integration
[**AppsIntegrationsApiUpdateIntegrationType**](IntegrationsAPI.md#AppsIntegrationsApiUpdateIntegrationType) | **Put** /api/v1/integrations/admin/integration-types/{integration_type_id} | Update Integration Type



## AppsIntegrationsApiBulkExportIntegrations

> []ExportIntegrationSchema AppsIntegrationsApiBulkExportIntegrations(ctx).IncludeSecrets(includeSecrets).RequestBody(requestBody).Execute()

Bulk Export Integrations



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
	includeSecrets := true // bool |  (optional) (default to false)
	requestBody := []*string{"Property_example"} // []*string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiBulkExportIntegrations(context.Background()).IncludeSecrets(includeSecrets).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiBulkExportIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiBulkExportIntegrations`: []ExportIntegrationSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiBulkExportIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiBulkExportIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSecrets** | **bool** |  | [default to false]
 **requestBody** | **[]string** |  | 

### Return type

[**[]ExportIntegrationSchema**](ExportIntegrationSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiBulkImportIntegrations

> []ImportResultSchema AppsIntegrationsApiBulkImportIntegrations(ctx).ExportIntegrationSchema(exportIntegrationSchema).Overwrite(overwrite).DecryptKey(decryptKey).Execute()

Bulk Import Integrations



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
	exportIntegrationSchema := []openapiclient.ExportIntegrationSchema{*openapiclient.NewExportIntegrationSchema("Name_example", "Type_example", "Description_example", map[string]interface{}{"key": interface{}(123)}, []string{"Tags_example"})} // []ExportIntegrationSchema | 
	overwrite := true // bool |  (optional) (default to false)
	decryptKey := "decryptKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiBulkImportIntegrations(context.Background()).ExportIntegrationSchema(exportIntegrationSchema).Overwrite(overwrite).DecryptKey(decryptKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiBulkImportIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiBulkImportIntegrations`: []ImportResultSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiBulkImportIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiBulkImportIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportIntegrationSchema** | [**[]ExportIntegrationSchema**](ExportIntegrationSchema.md) |  | 
 **overwrite** | **bool** |  | [default to false]
 **decryptKey** | **string** |  | 

### Return type

[**[]ImportResultSchema**](ImportResultSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiCreateConfigField

> ConfigFieldDefinitionResponse AppsIntegrationsApiCreateConfigField(ctx, integrationTypeId).ConfigFieldDefinitionRequest(configFieldDefinitionRequest).Execute()

Create Config Field



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
	integrationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configFieldDefinitionRequest := *openapiclient.NewConfigFieldDefinitionRequest("Name_example", "Label_example") // ConfigFieldDefinitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiCreateConfigField(context.Background(), integrationTypeId).ConfigFieldDefinitionRequest(configFieldDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiCreateConfigField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiCreateConfigField`: ConfigFieldDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiCreateConfigField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiCreateConfigFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configFieldDefinitionRequest** | [**ConfigFieldDefinitionRequest**](ConfigFieldDefinitionRequest.md) |  | 

### Return type

[**ConfigFieldDefinitionResponse**](ConfigFieldDefinitionResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiCreateIntegration

> IntegrationDetailSchema AppsIntegrationsApiCreateIntegration(ctx).CreateIntegrationSchema(createIntegrationSchema).Execute()

Create Integration



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
	createIntegrationSchema := *openapiclient.NewCreateIntegrationSchema("Name_example", "Type_example", map[string]interface{}{"key": interface{}(123)}) // CreateIntegrationSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiCreateIntegration(context.Background()).CreateIntegrationSchema(createIntegrationSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiCreateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiCreateIntegration`: IntegrationDetailSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiCreateIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiCreateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIntegrationSchema** | [**CreateIntegrationSchema**](CreateIntegrationSchema.md) |  | 

### Return type

[**IntegrationDetailSchema**](IntegrationDetailSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiCreateIntegrationType

> IntegrationTypeResponse AppsIntegrationsApiCreateIntegrationType(ctx).IntegrationTypeRequest(integrationTypeRequest).Execute()

Create Integration Type



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
	integrationTypeRequest := *openapiclient.NewIntegrationTypeRequest("Key_example", "Name_example", "Description_example", "HandlerClass_example") // IntegrationTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiCreateIntegrationType(context.Background()).IntegrationTypeRequest(integrationTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiCreateIntegrationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiCreateIntegrationType`: IntegrationTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiCreateIntegrationType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiCreateIntegrationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationTypeRequest** | [**IntegrationTypeRequest**](IntegrationTypeRequest.md) |  | 

### Return type

[**IntegrationTypeResponse**](IntegrationTypeResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiDeleteConfigField

> ActionResponse AppsIntegrationsApiDeleteConfigField(ctx, integrationTypeId, fieldId).Execute()

Delete Config Field



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
	integrationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fieldId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiDeleteConfigField(context.Background(), integrationTypeId, fieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiDeleteConfigField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiDeleteConfigField`: ActionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiDeleteConfigField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationTypeId** | **string** |  | 
**fieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiDeleteConfigFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiDeleteIntegration

> map[string]interface{} AppsIntegrationsApiDeleteIntegration(ctx, integrationId).Execute()

Delete Integration



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
	integrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiDeleteIntegration(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiDeleteIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiDeleteIntegration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiDeleteIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiDeleteIntegrationRequest struct via the builder pattern


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


## AppsIntegrationsApiDeleteIntegrationType

> ActionResponse AppsIntegrationsApiDeleteIntegrationType(ctx, integrationTypeId).Execute()

Delete Integration Type



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
	integrationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiDeleteIntegrationType(context.Background(), integrationTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiDeleteIntegrationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiDeleteIntegrationType`: ActionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiDeleteIntegrationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiDeleteIntegrationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiExportIntegration

> ExportIntegrationSchema AppsIntegrationsApiExportIntegration(ctx, integrationId).IncludeSecrets(includeSecrets).Execute()

Export Integration



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
	integrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeSecrets := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiExportIntegration(context.Background(), integrationId).IncludeSecrets(includeSecrets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiExportIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiExportIntegration`: ExportIntegrationSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiExportIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiExportIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSecrets** | **bool** |  | [default to false]

### Return type

[**ExportIntegrationSchema**](ExportIntegrationSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiGetAvailableIntegrationTypes

> []map[string]interface{} AppsIntegrationsApiGetAvailableIntegrationTypes(ctx).Execute()

Get Available Integration Types



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
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiGetAvailableIntegrationTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiGetAvailableIntegrationTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiGetAvailableIntegrationTypes`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiGetAvailableIntegrationTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiGetAvailableIntegrationTypesRequest struct via the builder pattern


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


## AppsIntegrationsApiGetIntegration

> IntegrationDetailSchema AppsIntegrationsApiGetIntegration(ctx, integrationId).Execute()

Get Integration



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
	integrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiGetIntegration(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiGetIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiGetIntegration`: IntegrationDetailSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiGetIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiGetIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationDetailSchema**](IntegrationDetailSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiGetIntegrationErrorStats

> map[string]interface{} AppsIntegrationsApiGetIntegrationErrorStats(ctx).Execute()

Get Integration Error Stats



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
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiGetIntegrationErrorStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiGetIntegrationErrorStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiGetIntegrationErrorStats`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiGetIntegrationErrorStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiGetIntegrationErrorStatsRequest struct via the builder pattern


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


## AppsIntegrationsApiGetIntegrationStats

> IntegrationStatsSchema AppsIntegrationsApiGetIntegrationStats(ctx, integrationId).Execute()

Get Integration Stats



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
	integrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiGetIntegrationStats(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiGetIntegrationStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiGetIntegrationStats`: IntegrationStatsSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiGetIntegrationStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiGetIntegrationStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationStatsSchema**](IntegrationStatsSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiGetIntegrationType

> IntegrationTypeResponse AppsIntegrationsApiGetIntegrationType(ctx, integrationTypeId).Execute()

Get Integration Type



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
	integrationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiGetIntegrationType(context.Background(), integrationTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiGetIntegrationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiGetIntegrationType`: IntegrationTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiGetIntegrationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiGetIntegrationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationTypeResponse**](IntegrationTypeResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiImportIntegration

> ImportResultSchema AppsIntegrationsApiImportIntegration(ctx).ImportIntegrationSchema(importIntegrationSchema).Execute()

Import Integration



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
	importIntegrationSchema := *openapiclient.NewImportIntegrationSchema(*openapiclient.NewExportIntegrationSchema("Name_example", "Type_example", "Description_example", map[string]interface{}{"key": interface{}(123)}, []string{"Tags_example"})) // ImportIntegrationSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiImportIntegration(context.Background()).ImportIntegrationSchema(importIntegrationSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiImportIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiImportIntegration`: ImportResultSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiImportIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiImportIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importIntegrationSchema** | [**ImportIntegrationSchema**](ImportIntegrationSchema.md) |  | 

### Return type

[**ImportResultSchema**](ImportResultSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiListConfigFields

> []ConfigFieldDefinitionResponse AppsIntegrationsApiListConfigFields(ctx, integrationTypeId).Execute()

List Config Fields



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
	integrationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiListConfigFields(context.Background(), integrationTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiListConfigFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiListConfigFields`: []ConfigFieldDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiListConfigFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiListConfigFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigFieldDefinitionResponse**](ConfigFieldDefinitionResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiListIntegrationTypes

> IntegrationTypeListResponse AppsIntegrationsApiListIntegrationTypes(ctx).Page(page).PageSize(pageSize).Category(category).IsActive(isActive).Search(search).Execute()

List Integration Types



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	category := "category_example" // string |  (optional)
	isActive := true // bool |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiListIntegrationTypes(context.Background()).Page(page).PageSize(pageSize).Category(category).IsActive(isActive).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiListIntegrationTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiListIntegrationTypes`: IntegrationTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiListIntegrationTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiListIntegrationTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **category** | **string** |  | 
 **isActive** | **bool** |  | 
 **search** | **string** |  | 

### Return type

[**IntegrationTypeListResponse**](IntegrationTypeListResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiListIntegrations

> []IntegrationResponseSchema AppsIntegrationsApiListIntegrations(ctx).Type_(type_).IsActive(isActive).IsVerified(isVerified).Search(search).Execute()

List Integrations



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
	type_ := "type__example" // string |  (optional)
	isActive := true // bool |  (optional)
	isVerified := true // bool |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiListIntegrations(context.Background()).Type_(type_).IsActive(isActive).IsVerified(isVerified).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiListIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiListIntegrations`: []IntegrationResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **isActive** | **bool** |  | 
 **isVerified** | **bool** |  | 
 **search** | **string** |  | 

### Return type

[**[]IntegrationResponseSchema**](IntegrationResponseSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiOauthAuthorize

> map[string]interface{} AppsIntegrationsApiOauthAuthorize(ctx, integrationType).Execute()

Oauth Authorize



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
	integrationType := "integrationType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiOauthAuthorize(context.Background(), integrationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiOauthAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiOauthAuthorize`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiOauthAuthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiOauthAuthorizeRequest struct via the builder pattern


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


## AppsIntegrationsApiOauthCallback

> IntegrationDetailSchema AppsIntegrationsApiOauthCallback(ctx, integrationType).Code(code).State(state).Execute()

Oauth Callback



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
	integrationType := "integrationType_example" // string | 
	code := "code_example" // string | 
	state := "state_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiOauthCallback(context.Background(), integrationType).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiOauthCallback`: IntegrationDetailSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiOauthCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** |  | 
 **state** | **string** |  | 

### Return type

[**IntegrationDetailSchema**](IntegrationDetailSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiTestAllIntegrations

> []map[string]interface{} AppsIntegrationsApiTestAllIntegrations(ctx).Execute()

Test All Integrations



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
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiTestAllIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiTestAllIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiTestAllIntegrations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiTestAllIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiTestAllIntegrationsRequest struct via the builder pattern


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


## AppsIntegrationsApiTestIntegration

> TestIntegrationResponseSchema AppsIntegrationsApiTestIntegration(ctx, integrationId).Execute()

Test Integration



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
	integrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiTestIntegration(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiTestIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiTestIntegration`: TestIntegrationResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiTestIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiTestIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestIntegrationResponseSchema**](TestIntegrationResponseSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiUpdateConfigField

> ConfigFieldDefinitionResponse AppsIntegrationsApiUpdateConfigField(ctx, integrationTypeId, fieldId).ConfigFieldDefinitionRequest(configFieldDefinitionRequest).Execute()

Update Config Field



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
	integrationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fieldId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configFieldDefinitionRequest := *openapiclient.NewConfigFieldDefinitionRequest("Name_example", "Label_example") // ConfigFieldDefinitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiUpdateConfigField(context.Background(), integrationTypeId, fieldId).ConfigFieldDefinitionRequest(configFieldDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiUpdateConfigField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiUpdateConfigField`: ConfigFieldDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiUpdateConfigField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationTypeId** | **string** |  | 
**fieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiUpdateConfigFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configFieldDefinitionRequest** | [**ConfigFieldDefinitionRequest**](ConfigFieldDefinitionRequest.md) |  | 

### Return type

[**ConfigFieldDefinitionResponse**](ConfigFieldDefinitionResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiUpdateIntegration

> IntegrationDetailSchema AppsIntegrationsApiUpdateIntegration(ctx, integrationId).UpdateIntegrationSchema(updateIntegrationSchema).Execute()

Update Integration



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
	integrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateIntegrationSchema := *openapiclient.NewUpdateIntegrationSchema() // UpdateIntegrationSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiUpdateIntegration(context.Background(), integrationId).UpdateIntegrationSchema(updateIntegrationSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiUpdateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiUpdateIntegration`: IntegrationDetailSchema
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiUpdateIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiUpdateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIntegrationSchema** | [**UpdateIntegrationSchema**](UpdateIntegrationSchema.md) |  | 

### Return type

[**IntegrationDetailSchema**](IntegrationDetailSchema.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsIntegrationsApiUpdateIntegrationType

> IntegrationTypeResponse AppsIntegrationsApiUpdateIntegrationType(ctx, integrationTypeId).IntegrationTypeRequest(integrationTypeRequest).Execute()

Update Integration Type



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
	integrationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	integrationTypeRequest := *openapiclient.NewIntegrationTypeRequest("Key_example", "Name_example", "Description_example", "HandlerClass_example") // IntegrationTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.AppsIntegrationsApiUpdateIntegrationType(context.Background(), integrationTypeId).IntegrationTypeRequest(integrationTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AppsIntegrationsApiUpdateIntegrationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsIntegrationsApiUpdateIntegrationType`: IntegrationTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.AppsIntegrationsApiUpdateIntegrationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsIntegrationsApiUpdateIntegrationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationTypeRequest** | [**IntegrationTypeRequest**](IntegrationTypeRequest.md) |  | 

### Return type

[**IntegrationTypeResponse**](IntegrationTypeResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

