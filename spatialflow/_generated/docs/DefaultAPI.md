# \DefaultAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsAuthenticationAppleMobileApiAppleNonce**](DefaultAPI.md#AppsAuthenticationAppleMobileApiAppleNonce) | **Post** /api/v1/auth/apple/nonce | Apple Nonce
[**AppsAuthenticationAppleMobileApiAppleTokenExchange**](DefaultAPI.md#AppsAuthenticationAppleMobileApiAppleTokenExchange) | **Post** /api/v1/auth/apple/token-exchange | Apple Token Exchange
[**AppsAuthenticationGoogleMobileApiGoogleTokenExchange**](DefaultAPI.md#AppsAuthenticationGoogleMobileApiGoogleTokenExchange) | **Post** /api/v1/auth/google/token-exchange | Google Token Exchange
[**AppsAuthenticationOauthApiDisconnectOauthAccount**](DefaultAPI.md#AppsAuthenticationOauthApiDisconnectOauthAccount) | **Delete** /api/v1/auth/oauth/{provider}/disconnect | Disconnect Oauth Account
[**AppsAuthenticationOauthApiGetLinkedAccounts**](DefaultAPI.md#AppsAuthenticationOauthApiGetLinkedAccounts) | **Get** /api/v1/auth/oauth/user/linked-accounts | Get Linked Accounts
[**AppsAuthenticationOauthApiGetOauthProviders**](DefaultAPI.md#AppsAuthenticationOauthApiGetOauthProviders) | **Get** /api/v1/auth/oauth/providers | Get Oauth Providers
[**AppsAuthenticationOauthApiLinkOauthAccount**](DefaultAPI.md#AppsAuthenticationOauthApiLinkOauthAccount) | **Post** /api/v1/auth/oauth/{provider}/link | Link Oauth Account
[**AppsAuthenticationOauthApiOauthAuthorize**](DefaultAPI.md#AppsAuthenticationOauthApiOauthAuthorize) | **Get** /api/v1/auth/oauth/{provider}/authorize | Oauth Authorize
[**AppsAuthenticationOauthApiOauthCallback**](DefaultAPI.md#AppsAuthenticationOauthApiOauthCallback) | **Get** /api/v1/auth/oauth/{provider}/callback | Oauth Callback
[**AppsAuthenticationSamlApiDetectMethod**](DefaultAPI.md#AppsAuthenticationSamlApiDetectMethod) | **Post** /api/v1/auth/saml/detect-method | Detect Method
[**AppsAuthenticationSamlApiInitiate**](DefaultAPI.md#AppsAuthenticationSamlApiInitiate) | **Get** /api/v1/auth/saml/{slug}/initiate | Initiate
[**AppsAuthenticationSamlApiMetadata**](DefaultAPI.md#AppsAuthenticationSamlApiMetadata) | **Get** /api/v1/auth/saml/{slug}/metadata | Metadata
[**AppsAuthenticationSamlApiSamlAcs**](DefaultAPI.md#AppsAuthenticationSamlApiSamlAcs) | **Post** /api/v1/auth/saml/{slug}/acs | Saml Acs
[**AppsEmailUnsubscribeUnsubscribe**](DefaultAPI.md#AppsEmailUnsubscribeUnsubscribe) | **Post** /api/v1/email/unsubscribe | Unsubscribe
[**AppsEmailUnsubscribeVerifyUnsubscribeToken**](DefaultAPI.md#AppsEmailUnsubscribeVerifyUnsubscribeToken) | **Get** /api/v1/email/unsubscribe/verify | Verify Unsubscribe Token



## AppsAuthenticationAppleMobileApiAppleNonce

> AppleNonceResponse AppsAuthenticationAppleMobileApiAppleNonce(ctx).Execute()

Apple Nonce



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
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationAppleMobileApiAppleNonce(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationAppleMobileApiAppleNonce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationAppleMobileApiAppleNonce`: AppleNonceResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationAppleMobileApiAppleNonce`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationAppleMobileApiAppleNonceRequest struct via the builder pattern


### Return type

[**AppleNonceResponse**](AppleNonceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationAppleMobileApiAppleTokenExchange

> AppleTokenExchangeResponse AppsAuthenticationAppleMobileApiAppleTokenExchange(ctx).AppleTokenExchangeRequest(appleTokenExchangeRequest).Execute()

Apple Token Exchange



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
	appleTokenExchangeRequest := *openapiclient.NewAppleTokenExchangeRequest("IdentityToken_example", "AuthorizationCode_example") // AppleTokenExchangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationAppleMobileApiAppleTokenExchange(context.Background()).AppleTokenExchangeRequest(appleTokenExchangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationAppleMobileApiAppleTokenExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationAppleMobileApiAppleTokenExchange`: AppleTokenExchangeResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationAppleMobileApiAppleTokenExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationAppleMobileApiAppleTokenExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appleTokenExchangeRequest** | [**AppleTokenExchangeRequest**](AppleTokenExchangeRequest.md) |  | 

### Return type

[**AppleTokenExchangeResponse**](AppleTokenExchangeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationGoogleMobileApiGoogleTokenExchange

> GoogleTokenExchangeResponse AppsAuthenticationGoogleMobileApiGoogleTokenExchange(ctx).GoogleTokenExchangeRequest(googleTokenExchangeRequest).Execute()

Google Token Exchange



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
	googleTokenExchangeRequest := *openapiclient.NewGoogleTokenExchangeRequest("IdToken_example") // GoogleTokenExchangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationGoogleMobileApiGoogleTokenExchange(context.Background()).GoogleTokenExchangeRequest(googleTokenExchangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationGoogleMobileApiGoogleTokenExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationGoogleMobileApiGoogleTokenExchange`: GoogleTokenExchangeResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationGoogleMobileApiGoogleTokenExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationGoogleMobileApiGoogleTokenExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleTokenExchangeRequest** | [**GoogleTokenExchangeRequest**](GoogleTokenExchangeRequest.md) |  | 

### Return type

[**GoogleTokenExchangeResponse**](GoogleTokenExchangeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationOauthApiDisconnectOauthAccount

> map[string]interface{} AppsAuthenticationOauthApiDisconnectOauthAccount(ctx, provider).Execute()

Disconnect Oauth Account



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationOauthApiDisconnectOauthAccount(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationOauthApiDisconnectOauthAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiDisconnectOauthAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationOauthApiDisconnectOauthAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationOauthApiDisconnectOauthAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AppsAuthenticationOauthApiGetLinkedAccounts

> map[string]interface{} AppsAuthenticationOauthApiGetLinkedAccounts(ctx).Execute()

Get Linked Accounts



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
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationOauthApiGetLinkedAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationOauthApiGetLinkedAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiGetLinkedAccounts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationOauthApiGetLinkedAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationOauthApiGetLinkedAccountsRequest struct via the builder pattern


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


## AppsAuthenticationOauthApiGetOauthProviders

> OAuthProvidersResponse AppsAuthenticationOauthApiGetOauthProviders(ctx).Execute()

Get Oauth Providers



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
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationOauthApiGetOauthProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationOauthApiGetOauthProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiGetOauthProviders`: OAuthProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationOauthApiGetOauthProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationOauthApiGetOauthProvidersRequest struct via the builder pattern


### Return type

[**OAuthProvidersResponse**](OAuthProvidersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationOauthApiLinkOauthAccount

> OAuthLinkResponse AppsAuthenticationOauthApiLinkOauthAccount(ctx, provider).Execute()

Link Oauth Account



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationOauthApiLinkOauthAccount(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationOauthApiLinkOauthAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiLinkOauthAccount`: OAuthLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationOauthApiLinkOauthAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationOauthApiLinkOauthAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuthLinkResponse**](OAuthLinkResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationOauthApiOauthAuthorize

> OAuthAuthorizeResponse AppsAuthenticationOauthApiOauthAuthorize(ctx, provider).Next(next).Execute()

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
	provider := "provider_example" // string | 
	next := "next_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationOauthApiOauthAuthorize(context.Background(), provider).Next(next).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationOauthApiOauthAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiOauthAuthorize`: OAuthAuthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationOauthApiOauthAuthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationOauthApiOauthAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **next** | **string** |  | 

### Return type

[**OAuthAuthorizeResponse**](OAuthAuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationOauthApiOauthCallback

> map[string]interface{} AppsAuthenticationOauthApiOauthCallback(ctx, provider).Code(code).State(state).Error_(error_).ErrorDescription(errorDescription).Execute()

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
	provider := "provider_example" // string | 
	code := "code_example" // string | 
	state := "state_example" // string | 
	error_ := "error__example" // string |  (optional)
	errorDescription := "errorDescription_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationOauthApiOauthCallback(context.Background(), provider).Code(code).State(state).Error_(error_).ErrorDescription(errorDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationOauthApiOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiOauthCallback`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationOauthApiOauthCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationOauthApiOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** |  | 
 **state** | **string** |  | 
 **error_** | **string** |  | 
 **errorDescription** | **string** |  | 

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


## AppsAuthenticationSamlApiDetectMethod

> DetectMethodResponse AppsAuthenticationSamlApiDetectMethod(ctx).DetectMethodRequest(detectMethodRequest).Execute()

Detect Method



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
	detectMethodRequest := *openapiclient.NewDetectMethodRequest("Email_example") // DetectMethodRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsAuthenticationSamlApiDetectMethod(context.Background()).DetectMethodRequest(detectMethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationSamlApiDetectMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationSamlApiDetectMethod`: DetectMethodResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsAuthenticationSamlApiDetectMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationSamlApiDetectMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detectMethodRequest** | [**DetectMethodRequest**](DetectMethodRequest.md) |  | 

### Return type

[**DetectMethodResponse**](DetectMethodResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationSamlApiInitiate

> AppsAuthenticationSamlApiInitiate(ctx, slug).Execute()

Initiate



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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AppsAuthenticationSamlApiInitiate(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationSamlApiInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationSamlApiInitiateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationSamlApiMetadata

> AppsAuthenticationSamlApiMetadata(ctx, slug).Execute()

Metadata



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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AppsAuthenticationSamlApiMetadata(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationSamlApiMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationSamlApiMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationSamlApiSamlAcs

> AppsAuthenticationSamlApiSamlAcs(ctx, slug).Execute()

Saml Acs



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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AppsAuthenticationSamlApiSamlAcs(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsAuthenticationSamlApiSamlAcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationSamlApiSamlAcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsEmailUnsubscribeUnsubscribe

> UnsubscribeResponse AppsEmailUnsubscribeUnsubscribe(ctx).UnsubscribeRequest(unsubscribeRequest).Execute()

Unsubscribe



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
	unsubscribeRequest := *openapiclient.NewUnsubscribeRequest("Token_example") // UnsubscribeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsEmailUnsubscribeUnsubscribe(context.Background()).UnsubscribeRequest(unsubscribeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsEmailUnsubscribeUnsubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsEmailUnsubscribeUnsubscribe`: UnsubscribeResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsEmailUnsubscribeUnsubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailUnsubscribeUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unsubscribeRequest** | [**UnsubscribeRequest**](UnsubscribeRequest.md) |  | 

### Return type

[**UnsubscribeResponse**](UnsubscribeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsEmailUnsubscribeVerifyUnsubscribeToken

> map[string]interface{} AppsEmailUnsubscribeVerifyUnsubscribeToken(ctx).Token(token).Execute()

Verify Unsubscribe Token



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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsEmailUnsubscribeVerifyUnsubscribeToken(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsEmailUnsubscribeVerifyUnsubscribeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsEmailUnsubscribeVerifyUnsubscribeToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsEmailUnsubscribeVerifyUnsubscribeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsEmailUnsubscribeVerifyUnsubscribeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** |  | 

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

