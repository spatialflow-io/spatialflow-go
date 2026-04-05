# \PoliciesAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsDevicesApiPoliciesCreatePolicy**](PoliciesAPI.md#AppsDevicesApiPoliciesCreatePolicy) | **Post** /api/v1/policies/ | Create policy
[**AppsDevicesApiPoliciesDeletePolicy**](PoliciesAPI.md#AppsDevicesApiPoliciesDeletePolicy) | **Delete** /api/v1/policies/{policy_id}/ | Delete policy
[**AppsDevicesApiPoliciesGetPolicy**](PoliciesAPI.md#AppsDevicesApiPoliciesGetPolicy) | **Get** /api/v1/policies/{policy_id}/ | Get policy detail
[**AppsDevicesApiPoliciesListPolicies**](PoliciesAPI.md#AppsDevicesApiPoliciesListPolicies) | **Get** /api/v1/policies/ | List policies
[**AppsDevicesApiPoliciesListTemplates**](PoliciesAPI.md#AppsDevicesApiPoliciesListTemplates) | **Get** /api/v1/policies/templates | List policy templates
[**AppsDevicesApiPoliciesUpdatePolicy**](PoliciesAPI.md#AppsDevicesApiPoliciesUpdatePolicy) | **Put** /api/v1/policies/{policy_id}/ | Update policy



## AppsDevicesApiPoliciesCreatePolicy

> PolicyOut AppsDevicesApiPoliciesCreatePolicy(ctx).PolicyCreateIn(policyCreateIn).Execute()

Create policy



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
	policyCreateIn := *openapiclient.NewPolicyCreateIn("Name_example", "GeofenceId_example") // PolicyCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.AppsDevicesApiPoliciesCreatePolicy(context.Background()).PolicyCreateIn(policyCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AppsDevicesApiPoliciesCreatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiPoliciesCreatePolicy`: PolicyOut
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AppsDevicesApiPoliciesCreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiPoliciesCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyCreateIn** | [**PolicyCreateIn**](PolicyCreateIn.md) |  | 

### Return type

[**PolicyOut**](PolicyOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiPoliciesDeletePolicy

> AppsDevicesApiPoliciesDeletePolicy(ctx, policyId).Execute()

Delete policy



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
	policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.AppsDevicesApiPoliciesDeletePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AppsDevicesApiPoliciesDeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiPoliciesDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AppsDevicesApiPoliciesGetPolicy

> PolicyOut AppsDevicesApiPoliciesGetPolicy(ctx, policyId).Execute()

Get policy detail



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
	policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.AppsDevicesApiPoliciesGetPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AppsDevicesApiPoliciesGetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiPoliciesGetPolicy`: PolicyOut
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AppsDevicesApiPoliciesGetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiPoliciesGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyOut**](PolicyOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiPoliciesListPolicies

> PolicyListOut AppsDevicesApiPoliciesListPolicies(ctx).EnabledOnly(enabledOnly).Limit(limit).Offset(offset).Execute()

List policies



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
	enabledOnly := true // bool |  (optional) (default to false)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.AppsDevicesApiPoliciesListPolicies(context.Background()).EnabledOnly(enabledOnly).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AppsDevicesApiPoliciesListPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiPoliciesListPolicies`: PolicyListOut
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AppsDevicesApiPoliciesListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiPoliciesListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabledOnly** | **bool** |  | [default to false]
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PolicyListOut**](PolicyListOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiPoliciesListTemplates

> []PolicyTemplateOut AppsDevicesApiPoliciesListTemplates(ctx).Execute()

List policy templates



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
	resp, r, err := apiClient.PoliciesAPI.AppsDevicesApiPoliciesListTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AppsDevicesApiPoliciesListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiPoliciesListTemplates`: []PolicyTemplateOut
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AppsDevicesApiPoliciesListTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiPoliciesListTemplatesRequest struct via the builder pattern


### Return type

[**[]PolicyTemplateOut**](PolicyTemplateOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDevicesApiPoliciesUpdatePolicy

> PolicyOut AppsDevicesApiPoliciesUpdatePolicy(ctx, policyId).PolicyUpdateIn(policyUpdateIn).Execute()

Update policy



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
	policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	policyUpdateIn := *openapiclient.NewPolicyUpdateIn() // PolicyUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.AppsDevicesApiPoliciesUpdatePolicy(context.Background(), policyId).PolicyUpdateIn(policyUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AppsDevicesApiPoliciesUpdatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsDevicesApiPoliciesUpdatePolicy`: PolicyOut
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AppsDevicesApiPoliciesUpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDevicesApiPoliciesUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyUpdateIn** | [**PolicyUpdateIn**](PolicyUpdateIn.md) |  | 

### Return type

[**PolicyOut**](PolicyOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

