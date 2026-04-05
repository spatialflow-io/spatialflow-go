# \E2ETestAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsTestApiCleanupE2eData**](E2ETestAPI.md#AppsTestApiCleanupE2eData) | **Post** /api/v1/test/cleanup-e2e-data | Cleanup E2E Data
[**AppsTestApiCleanupTestData**](E2ETestAPI.md#AppsTestApiCleanupTestData) | **Delete** /api/v1/test/cleanup | Cleanup Test Data
[**AppsTestApiCreateTestUser**](E2ETestAPI.md#AppsTestApiCreateTestUser) | **Post** /api/v1/test/create-user | Create Test User
[**AppsTestApiSeedE2eData**](E2ETestAPI.md#AppsTestApiSeedE2eData) | **Post** /api/v1/test/seed-e2e-data | Seed E2E Data



## AppsTestApiCleanupE2eData

> map[string]interface{} AppsTestApiCleanupE2eData(ctx).Execute()

Cleanup E2E Data



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
	resp, r, err := apiClient.E2ETestAPI.AppsTestApiCleanupE2eData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E2ETestAPI.AppsTestApiCleanupE2eData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsTestApiCleanupE2eData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `E2ETestAPI.AppsTestApiCleanupE2eData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTestApiCleanupE2eDataRequest struct via the builder pattern


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


## AppsTestApiCleanupTestData

> map[string]interface{} AppsTestApiCleanupTestData(ctx).Execute()

Cleanup Test Data



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
	resp, r, err := apiClient.E2ETestAPI.AppsTestApiCleanupTestData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E2ETestAPI.AppsTestApiCleanupTestData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsTestApiCleanupTestData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `E2ETestAPI.AppsTestApiCleanupTestData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTestApiCleanupTestDataRequest struct via the builder pattern


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


## AppsTestApiCreateTestUser

> map[string]interface{} AppsTestApiCreateTestUser(ctx).CreateUserSchema(createUserSchema).Execute()

Create Test User



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
	createUserSchema := *openapiclient.NewCreateUserSchema("Email_example", "Password_example") // CreateUserSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.E2ETestAPI.AppsTestApiCreateTestUser(context.Background()).CreateUserSchema(createUserSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E2ETestAPI.AppsTestApiCreateTestUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsTestApiCreateTestUser`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `E2ETestAPI.AppsTestApiCreateTestUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsTestApiCreateTestUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserSchema** | [**CreateUserSchema**](CreateUserSchema.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsTestApiSeedE2eData

> SeedDataResponseSchema AppsTestApiSeedE2eData(ctx).Execute()

Seed E2E Data



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
	resp, r, err := apiClient.E2ETestAPI.AppsTestApiSeedE2eData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E2ETestAPI.AppsTestApiSeedE2eData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsTestApiSeedE2eData`: SeedDataResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `E2ETestAPI.AppsTestApiSeedE2eData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTestApiSeedE2eDataRequest struct via the builder pattern


### Return type

[**SeedDataResponseSchema**](SeedDataResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

