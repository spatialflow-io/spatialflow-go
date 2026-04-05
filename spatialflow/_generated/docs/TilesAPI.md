# \TilesAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsTilesApiGetRasterTile**](TilesAPI.md#AppsTilesApiGetRasterTile) | **Get** /api/v1/tiles/{z}/{x}/{y}.png | Get Raster Tile
[**AppsTilesApiGetTile**](TilesAPI.md#AppsTilesApiGetTile) | **Get** /api/v1/tiles/{z}/{x}/{y}.mvt | Get Tile
[**AppsTilesApiGetTileMetadata**](TilesAPI.md#AppsTilesApiGetTileMetadata) | **Get** /api/v1/tiles/metadata | Get Tile Metadata
[**AppsTilesApiGetTileStyle**](TilesAPI.md#AppsTilesApiGetTileStyle) | **Get** /api/v1/tiles/style/{style_id} | Get Tile Style
[**AppsTilesApiHealthCheck**](TilesAPI.md#AppsTilesApiHealthCheck) | **Get** /api/v1/tiles/health | Health Check
[**AppsTilesApiInvalidateTiles**](TilesAPI.md#AppsTilesApiInvalidateTiles) | **Post** /api/v1/tiles/invalidate | Invalidate Tiles



## AppsTilesApiGetRasterTile

> AppsTilesApiGetRasterTile(ctx, z, x, y).Execute()

Get Raster Tile



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
	z := int32(56) // int32 | 
	x := int32(56) // int32 | 
	y := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TilesAPI.AppsTilesApiGetRasterTile(context.Background(), z, x, y).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TilesAPI.AppsTilesApiGetRasterTile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**z** | **int32** |  | 
**x** | **int32** |  | 
**y** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTilesApiGetRasterTileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[TileJWTAuth](../README.md#TileJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsTilesApiGetTile

> *os.File AppsTilesApiGetTile(ctx, z, x, y).Execute()

Get Tile



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
	z := int32(56) // int32 | 
	x := int32(56) // int32 | 
	y := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TilesAPI.AppsTilesApiGetTile(context.Background(), z, x, y).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TilesAPI.AppsTilesApiGetTile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsTilesApiGetTile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TilesAPI.AppsTilesApiGetTile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**z** | **int32** |  | 
**x** | **int32** |  | 
**y** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTilesApiGetTileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[TileJWTAuth](../README.md#TileJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsTilesApiGetTileMetadata

> TileMetadata AppsTilesApiGetTileMetadata(ctx).Execute()

Get Tile Metadata



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
	resp, r, err := apiClient.TilesAPI.AppsTilesApiGetTileMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TilesAPI.AppsTilesApiGetTileMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsTilesApiGetTileMetadata`: TileMetadata
	fmt.Fprintf(os.Stdout, "Response from `TilesAPI.AppsTilesApiGetTileMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTilesApiGetTileMetadataRequest struct via the builder pattern


### Return type

[**TileMetadata**](TileMetadata.md)

### Authorization

[TileJWTAuth](../README.md#TileJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsTilesApiGetTileStyle

> AppsTilesApiGetTileStyle(ctx, styleId).Execute()

Get Tile Style



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
	styleId := "styleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TilesAPI.AppsTilesApiGetTileStyle(context.Background(), styleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TilesAPI.AppsTilesApiGetTileStyle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**styleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTilesApiGetTileStyleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TileJWTAuth](../README.md#TileJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsTilesApiHealthCheck

> map[string]interface{} AppsTilesApiHealthCheck(ctx).Execute()

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
	resp, r, err := apiClient.TilesAPI.AppsTilesApiHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TilesAPI.AppsTilesApiHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsTilesApiHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TilesAPI.AppsTilesApiHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTilesApiHealthCheckRequest struct via the builder pattern


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


## AppsTilesApiInvalidateTiles

> map[string]interface{} AppsTilesApiInvalidateTiles(ctx).Execute()

Invalidate Tiles



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
	resp, r, err := apiClient.TilesAPI.AppsTilesApiInvalidateTiles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TilesAPI.AppsTilesApiInvalidateTiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsTilesApiInvalidateTiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TilesAPI.AppsTilesApiInvalidateTiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsTilesApiInvalidateTilesRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[TileJWTAuth](../README.md#TileJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

