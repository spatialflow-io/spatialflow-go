# \StorageAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsStorageApiCreatePresignedUrl**](StorageAPI.md#AppsStorageApiCreatePresignedUrl) | **Post** /api/v1/storage/presigned-url | Create Presigned Url
[**AppsStorageApiDeleteFile**](StorageAPI.md#AppsStorageApiDeleteFile) | **Delete** /api/v1/storage/{file_type}/{filename} | Delete File
[**AppsStorageApiGetDownloadUrl**](StorageAPI.md#AppsStorageApiGetDownloadUrl) | **Get** /api/v1/storage/download/{file_id} | Get Download Url
[**AppsStorageApiGetFileTypes**](StorageAPI.md#AppsStorageApiGetFileTypes) | **Get** /api/v1/storage/types | Get File Types
[**AppsStorageApiHealthCheck**](StorageAPI.md#AppsStorageApiHealthCheck) | **Get** /api/v1/storage/health | Health Check
[**AppsStorageApiListFiles**](StorageAPI.md#AppsStorageApiListFiles) | **Get** /api/v1/storage/list/{file_type} | List Files



## AppsStorageApiCreatePresignedUrl

> PresignedUrlResponse AppsStorageApiCreatePresignedUrl(ctx).PresignedUrlRequest(presignedUrlRequest).Execute()

Create Presigned Url



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
	presignedUrlRequest := *openapiclient.NewPresignedUrlRequest("FileType_example", "Filename_example", int32(123)) // PresignedUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.AppsStorageApiCreatePresignedUrl(context.Background()).PresignedUrlRequest(presignedUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.AppsStorageApiCreatePresignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsStorageApiCreatePresignedUrl`: PresignedUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.AppsStorageApiCreatePresignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsStorageApiCreatePresignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presignedUrlRequest** | [**PresignedUrlRequest**](PresignedUrlRequest.md) |  | 

### Return type

[**PresignedUrlResponse**](PresignedUrlResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsStorageApiDeleteFile

> DeleteFileResponse AppsStorageApiDeleteFile(ctx, fileType, filename).Execute()

Delete File



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
	fileType := "fileType_example" // string | 
	filename := "filename_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.AppsStorageApiDeleteFile(context.Background(), fileType, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.AppsStorageApiDeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsStorageApiDeleteFile`: DeleteFileResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.AppsStorageApiDeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileType** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsStorageApiDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteFileResponse**](DeleteFileResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsStorageApiGetDownloadUrl

> map[string]interface{} AppsStorageApiGetDownloadUrl(ctx, fileId).Execute()

Get Download Url



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
	fileId := "fileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.AppsStorageApiGetDownloadUrl(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.AppsStorageApiGetDownloadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsStorageApiGetDownloadUrl`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.AppsStorageApiGetDownloadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsStorageApiGetDownloadUrlRequest struct via the builder pattern


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


## AppsStorageApiGetFileTypes

> map[string]interface{} AppsStorageApiGetFileTypes(ctx).Execute()

Get File Types



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
	resp, r, err := apiClient.StorageAPI.AppsStorageApiGetFileTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.AppsStorageApiGetFileTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsStorageApiGetFileTypes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.AppsStorageApiGetFileTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsStorageApiGetFileTypesRequest struct via the builder pattern


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


## AppsStorageApiHealthCheck

> map[string]interface{} AppsStorageApiHealthCheck(ctx).Execute()

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
	resp, r, err := apiClient.StorageAPI.AppsStorageApiHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.AppsStorageApiHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsStorageApiHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.AppsStorageApiHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsStorageApiHealthCheckRequest struct via the builder pattern


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


## AppsStorageApiListFiles

> FileListResponse AppsStorageApiListFiles(ctx, fileType).Execute()

List Files



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
	fileType := "fileType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.AppsStorageApiListFiles(context.Background(), fileType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.AppsStorageApiListFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsStorageApiListFiles`: FileListResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.AppsStorageApiListFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsStorageApiListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileListResponse**](FileListResponse.md)

### Authorization

[APIKeyBearer](../README.md#APIKeyBearer), [JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

