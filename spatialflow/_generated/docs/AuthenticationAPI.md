# \AuthenticationAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsAuthenticationApiAcceptInvitation**](AuthenticationAPI.md#AppsAuthenticationApiAcceptInvitation) | **Post** /api/v1/auth/accept-invite | Accept Invitation
[**AppsAuthenticationApiChangePassword**](AuthenticationAPI.md#AppsAuthenticationApiChangePassword) | **Post** /api/v1/auth/change-password | Change Password
[**AppsAuthenticationApiConfirmPasswordReset**](AuthenticationAPI.md#AppsAuthenticationApiConfirmPasswordReset) | **Post** /api/v1/auth/password-reset/confirm | Confirm Password Reset
[**AppsAuthenticationApiForgotPassword**](AuthenticationAPI.md#AppsAuthenticationApiForgotPassword) | **Post** /api/v1/auth/forgot-password | Forgot Password
[**AppsAuthenticationApiGetCurrentUser**](AuthenticationAPI.md#AppsAuthenticationApiGetCurrentUser) | **Get** /api/v1/auth/me | Get Current User
[**AppsAuthenticationApiGetFeatureFlags**](AuthenticationAPI.md#AppsAuthenticationApiGetFeatureFlags) | **Get** /api/v1/auth/feature-flags | Get Feature Flags
[**AppsAuthenticationApiHealthCheck**](AuthenticationAPI.md#AppsAuthenticationApiHealthCheck) | **Get** /api/v1/auth/health | Health Check
[**AppsAuthenticationApiLogin**](AuthenticationAPI.md#AppsAuthenticationApiLogin) | **Post** /api/v1/auth/login | Login
[**AppsAuthenticationApiLogout**](AuthenticationAPI.md#AppsAuthenticationApiLogout) | **Post** /api/v1/auth/logout | Logout
[**AppsAuthenticationApiPasswordResetAlias**](AuthenticationAPI.md#AppsAuthenticationApiPasswordResetAlias) | **Post** /api/v1/auth/password-reset | Password Reset Alias
[**AppsAuthenticationApiRefreshToken**](AuthenticationAPI.md#AppsAuthenticationApiRefreshToken) | **Post** /api/v1/auth/refresh | Refresh Token
[**AppsAuthenticationApiRegister**](AuthenticationAPI.md#AppsAuthenticationApiRegister) | **Post** /api/v1/auth/register | Register
[**AppsAuthenticationApiResendVerification**](AuthenticationAPI.md#AppsAuthenticationApiResendVerification) | **Post** /api/v1/auth/resend-verification-authenticated | Resend Verification
[**AppsAuthenticationApiResendVerificationAlias**](AuthenticationAPI.md#AppsAuthenticationApiResendVerificationAlias) | **Post** /api/v1/auth/resend-verification | Resend Verification Alias
[**AppsAuthenticationApiResendVerificationEmail**](AuthenticationAPI.md#AppsAuthenticationApiResendVerificationEmail) | **Post** /api/v1/auth/resend-verification-email | Resend Verification Email
[**AppsAuthenticationApiResetPassword**](AuthenticationAPI.md#AppsAuthenticationApiResetPassword) | **Post** /api/v1/auth/reset-password | Reset Password
[**AppsAuthenticationApiSsoStart**](AuthenticationAPI.md#AppsAuthenticationApiSsoStart) | **Get** /api/v1/auth/sso/start | Sso Start
[**AppsAuthenticationApiVerifyEmail**](AuthenticationAPI.md#AppsAuthenticationApiVerifyEmail) | **Get** /api/v1/auth/verify-email | Verify Email
[**AppsAuthenticationApiVerifyEmailPath**](AuthenticationAPI.md#AppsAuthenticationApiVerifyEmailPath) | **Get** /api/v1/auth/verify-email/{token} | Verify Email Path
[**AppsAuthenticationApiVerifyEmailPost**](AuthenticationAPI.md#AppsAuthenticationApiVerifyEmailPost) | **Post** /api/v1/auth/verify-email | Verify Email Post
[**AppsAuthenticationAppleMobileApiAppleNonce**](AuthenticationAPI.md#AppsAuthenticationAppleMobileApiAppleNonce) | **Post** /api/v1/auth/apple/nonce | Apple Nonce
[**AppsAuthenticationAppleMobileApiAppleTokenExchange**](AuthenticationAPI.md#AppsAuthenticationAppleMobileApiAppleTokenExchange) | **Post** /api/v1/auth/apple/token-exchange | Apple Token Exchange
[**AppsAuthenticationGoogleMobileApiGoogleTokenExchange**](AuthenticationAPI.md#AppsAuthenticationGoogleMobileApiGoogleTokenExchange) | **Post** /api/v1/auth/google/token-exchange | Google Token Exchange
[**AppsAuthenticationOauthApiDisconnectOauthAccount**](AuthenticationAPI.md#AppsAuthenticationOauthApiDisconnectOauthAccount) | **Delete** /api/v1/auth/oauth/{provider}/disconnect | Disconnect Oauth Account
[**AppsAuthenticationOauthApiGetLinkedAccounts**](AuthenticationAPI.md#AppsAuthenticationOauthApiGetLinkedAccounts) | **Get** /api/v1/auth/oauth/user/linked-accounts | Get Linked Accounts
[**AppsAuthenticationOauthApiGetOauthProviders**](AuthenticationAPI.md#AppsAuthenticationOauthApiGetOauthProviders) | **Get** /api/v1/auth/oauth/providers | Get Oauth Providers
[**AppsAuthenticationOauthApiLinkOauthAccount**](AuthenticationAPI.md#AppsAuthenticationOauthApiLinkOauthAccount) | **Post** /api/v1/auth/oauth/{provider}/link | Link Oauth Account
[**AppsAuthenticationOauthApiOauthAuthorize**](AuthenticationAPI.md#AppsAuthenticationOauthApiOauthAuthorize) | **Get** /api/v1/auth/oauth/{provider}/authorize | Oauth Authorize
[**AppsAuthenticationOauthApiOauthCallback**](AuthenticationAPI.md#AppsAuthenticationOauthApiOauthCallback) | **Get** /api/v1/auth/oauth/{provider}/callback | Oauth Callback
[**AppsAuthenticationSamlApiDetectMethod**](AuthenticationAPI.md#AppsAuthenticationSamlApiDetectMethod) | **Post** /api/v1/auth/saml/detect-method | Detect Method
[**AppsAuthenticationSamlApiInitiate**](AuthenticationAPI.md#AppsAuthenticationSamlApiInitiate) | **Get** /api/v1/auth/saml/{slug}/initiate | Initiate
[**AppsAuthenticationSamlApiMetadata**](AuthenticationAPI.md#AppsAuthenticationSamlApiMetadata) | **Get** /api/v1/auth/saml/{slug}/metadata | Metadata
[**AppsAuthenticationSamlApiSamlAcs**](AuthenticationAPI.md#AppsAuthenticationSamlApiSamlAcs) | **Post** /api/v1/auth/saml/{slug}/acs | Saml Acs



## AppsAuthenticationApiAcceptInvitation

> map[string]interface{} AppsAuthenticationApiAcceptInvitation(ctx).AcceptInviteSchema(acceptInviteSchema).Execute()

Accept Invitation



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
	acceptInviteSchema := *openapiclient.NewAcceptInviteSchema("Token_example", "InviteId_example") // AcceptInviteSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiAcceptInvitation(context.Background()).AcceptInviteSchema(acceptInviteSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiAcceptInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiAcceptInvitation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiAcceptInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiAcceptInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptInviteSchema** | [**AcceptInviteSchema**](AcceptInviteSchema.md) |  | 

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


## AppsAuthenticationApiChangePassword

> map[string]interface{} AppsAuthenticationApiChangePassword(ctx).ChangePasswordSchema(changePasswordSchema).Execute()

Change Password



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
	changePasswordSchema := *openapiclient.NewChangePasswordSchema("CurrentPassword_example", "NewPassword_example") // ChangePasswordSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiChangePassword(context.Background()).ChangePasswordSchema(changePasswordSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiChangePassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiChangePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePasswordSchema** | [**ChangePasswordSchema**](ChangePasswordSchema.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationApiConfirmPasswordReset

> map[string]interface{} AppsAuthenticationApiConfirmPasswordReset(ctx).ConfirmPasswordResetSchema(confirmPasswordResetSchema).Execute()

Confirm Password Reset



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
	confirmPasswordResetSchema := *openapiclient.NewConfirmPasswordResetSchema("Token_example", "Password_example") // ConfirmPasswordResetSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiConfirmPasswordReset(context.Background()).ConfirmPasswordResetSchema(confirmPasswordResetSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiConfirmPasswordReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiConfirmPasswordReset`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiConfirmPasswordReset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiConfirmPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmPasswordResetSchema** | [**ConfirmPasswordResetSchema**](ConfirmPasswordResetSchema.md) |  | 

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


## AppsAuthenticationApiForgotPassword

> map[string]interface{} AppsAuthenticationApiForgotPassword(ctx).ForgotPasswordSchema(forgotPasswordSchema).Execute()

Forgot Password



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
	forgotPasswordSchema := *openapiclient.NewForgotPasswordSchema("Email_example") // ForgotPasswordSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiForgotPassword(context.Background()).ForgotPasswordSchema(forgotPasswordSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiForgotPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiForgotPassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiForgotPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordSchema** | [**ForgotPasswordSchema**](ForgotPasswordSchema.md) |  | 

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


## AppsAuthenticationApiGetCurrentUser

> UserResponse AppsAuthenticationApiGetCurrentUser(ctx).Execute()

Get Current User



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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiGetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiGetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiGetCurrentUser`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiGetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiGetCurrentUserRequest struct via the builder pattern


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationApiGetFeatureFlags

> map[string]interface{} AppsAuthenticationApiGetFeatureFlags(ctx).Execute()

Get Feature Flags



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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiGetFeatureFlags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiGetFeatureFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiGetFeatureFlags`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiGetFeatureFlags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiGetFeatureFlagsRequest struct via the builder pattern


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


## AppsAuthenticationApiHealthCheck

> map[string]interface{} AppsAuthenticationApiHealthCheck(ctx).Execute()

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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiHealthCheckRequest struct via the builder pattern


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


## AppsAuthenticationApiLogin

> LoginResponse AppsAuthenticationApiLogin(ctx).LoginSchema(loginSchema).Execute()

Login



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
	loginSchema := *openapiclient.NewLoginSchema("Email_example", "Password_example") // LoginSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiLogin(context.Background()).LoginSchema(loginSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiLogin`: LoginResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginSchema** | [**LoginSchema**](LoginSchema.md) |  | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationApiLogout

> AppsAuthenticationApiLogout(ctx).Execute()

Logout



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
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiLogoutRequest struct via the builder pattern


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


## AppsAuthenticationApiPasswordResetAlias

> map[string]interface{} AppsAuthenticationApiPasswordResetAlias(ctx).ForgotPasswordSchema(forgotPasswordSchema).Execute()

Password Reset Alias



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
	forgotPasswordSchema := *openapiclient.NewForgotPasswordSchema("Email_example") // ForgotPasswordSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiPasswordResetAlias(context.Background()).ForgotPasswordSchema(forgotPasswordSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiPasswordResetAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiPasswordResetAlias`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiPasswordResetAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiPasswordResetAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordSchema** | [**ForgotPasswordSchema**](ForgotPasswordSchema.md) |  | 

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


## AppsAuthenticationApiRefreshToken

> LoginResponse AppsAuthenticationApiRefreshToken(ctx).RefreshTokenSchema(refreshTokenSchema).Execute()

Refresh Token



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
	refreshTokenSchema := *openapiclient.NewRefreshTokenSchema() // RefreshTokenSchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiRefreshToken(context.Background()).RefreshTokenSchema(refreshTokenSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiRefreshToken`: LoginResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenSchema** | [**RefreshTokenSchema**](RefreshTokenSchema.md) |  | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationApiRegister

> AppsAuthenticationApiRegister(ctx).RegisterSchema(registerSchema).Execute()

Register



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
	registerSchema := *openapiclient.NewRegisterSchema("Email_example", "Password_example", "Name_example") // RegisterSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiRegister(context.Background()).RegisterSchema(registerSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerSchema** | [**RegisterSchema**](RegisterSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationApiResendVerification

> map[string]interface{} AppsAuthenticationApiResendVerification(ctx).Execute()

Resend Verification



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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiResendVerification(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiResendVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiResendVerification`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiResendVerification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiResendVerificationRequest struct via the builder pattern


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


## AppsAuthenticationApiResendVerificationAlias

> map[string]interface{} AppsAuthenticationApiResendVerificationAlias(ctx).ResendVerificationSchema(resendVerificationSchema).Execute()

Resend Verification Alias



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
	resendVerificationSchema := *openapiclient.NewResendVerificationSchema("Email_example") // ResendVerificationSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiResendVerificationAlias(context.Background()).ResendVerificationSchema(resendVerificationSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiResendVerificationAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiResendVerificationAlias`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiResendVerificationAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiResendVerificationAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resendVerificationSchema** | [**ResendVerificationSchema**](ResendVerificationSchema.md) |  | 

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


## AppsAuthenticationApiResendVerificationEmail

> map[string]interface{} AppsAuthenticationApiResendVerificationEmail(ctx).ResendVerificationSchema(resendVerificationSchema).Execute()

Resend Verification Email



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
	resendVerificationSchema := *openapiclient.NewResendVerificationSchema("Email_example") // ResendVerificationSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiResendVerificationEmail(context.Background()).ResendVerificationSchema(resendVerificationSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiResendVerificationEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiResendVerificationEmail`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiResendVerificationEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiResendVerificationEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resendVerificationSchema** | [**ResendVerificationSchema**](ResendVerificationSchema.md) |  | 

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


## AppsAuthenticationApiResetPassword

> AppsAuthenticationApiResetPassword(ctx).ResetPasswordSchema(resetPasswordSchema).Execute()

Reset Password



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
	resetPasswordSchema := *openapiclient.NewResetPasswordSchema("Uid_example", "Token_example", "NewPassword_example") // ResetPasswordSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiResetPassword(context.Background()).ResetPasswordSchema(resetPasswordSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordSchema** | [**ResetPasswordSchema**](ResetPasswordSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAuthenticationApiSsoStart

> AppsAuthenticationApiSsoStart(ctx).Return_(return_).WorkspaceSlug(workspaceSlug).Execute()

Sso Start



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
	return_ := "return__example" // string | 
	workspaceSlug := "workspaceSlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiSsoStart(context.Background()).Return_(return_).WorkspaceSlug(workspaceSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiSsoStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiSsoStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **return_** | **string** |  | 
 **workspaceSlug** | **string** |  | 

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


## AppsAuthenticationApiVerifyEmail

> map[string]interface{} AppsAuthenticationApiVerifyEmail(ctx).Token(token).Execute()

Verify Email



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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiVerifyEmail(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiVerifyEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiVerifyEmail`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiVerifyEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiVerifyEmailRequest struct via the builder pattern


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


## AppsAuthenticationApiVerifyEmailPath

> map[string]interface{} AppsAuthenticationApiVerifyEmailPath(ctx, token).Execute()

Verify Email Path



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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiVerifyEmailPath(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiVerifyEmailPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiVerifyEmailPath`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiVerifyEmailPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiVerifyEmailPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AppsAuthenticationApiVerifyEmailPost

> map[string]interface{} AppsAuthenticationApiVerifyEmailPost(ctx).VerifyEmailSchema(verifyEmailSchema).Execute()

Verify Email Post



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
	verifyEmailSchema := *openapiclient.NewVerifyEmailSchema("Token_example") // VerifyEmailSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationApiVerifyEmailPost(context.Background()).VerifyEmailSchema(verifyEmailSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationApiVerifyEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationApiVerifyEmailPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationApiVerifyEmailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationApiVerifyEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyEmailSchema** | [**VerifyEmailSchema**](VerifyEmailSchema.md) |  | 

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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationAppleMobileApiAppleNonce(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationAppleMobileApiAppleNonce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationAppleMobileApiAppleNonce`: AppleNonceResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationAppleMobileApiAppleNonce`: %v\n", resp)
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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationAppleMobileApiAppleTokenExchange(context.Background()).AppleTokenExchangeRequest(appleTokenExchangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationAppleMobileApiAppleTokenExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationAppleMobileApiAppleTokenExchange`: AppleTokenExchangeResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationAppleMobileApiAppleTokenExchange`: %v\n", resp)
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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationGoogleMobileApiGoogleTokenExchange(context.Background()).GoogleTokenExchangeRequest(googleTokenExchangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationGoogleMobileApiGoogleTokenExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationGoogleMobileApiGoogleTokenExchange`: GoogleTokenExchangeResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationGoogleMobileApiGoogleTokenExchange`: %v\n", resp)
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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationOauthApiDisconnectOauthAccount(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationOauthApiDisconnectOauthAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiDisconnectOauthAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationOauthApiDisconnectOauthAccount`: %v\n", resp)
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

> AppsAuthenticationOauthApiGetLinkedAccounts(ctx).Execute()

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
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationOauthApiGetLinkedAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationOauthApiGetLinkedAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAuthenticationOauthApiGetLinkedAccountsRequest struct via the builder pattern


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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationOauthApiGetOauthProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationOauthApiGetOauthProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiGetOauthProviders`: OAuthProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationOauthApiGetOauthProviders`: %v\n", resp)
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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationOauthApiLinkOauthAccount(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationOauthApiLinkOauthAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiLinkOauthAccount`: OAuthLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationOauthApiLinkOauthAccount`: %v\n", resp)
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

> OAuthAuthorizeResponse AppsAuthenticationOauthApiOauthAuthorize(ctx, provider).Next(next).InviteId(inviteId).InviteToken(inviteToken).Execute()

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
	inviteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	inviteToken := "inviteToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationOauthApiOauthAuthorize(context.Background(), provider).Next(next).InviteId(inviteId).InviteToken(inviteToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationOauthApiOauthAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationOauthApiOauthAuthorize`: OAuthAuthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationOauthApiOauthAuthorize`: %v\n", resp)
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
 **inviteId** | **string** |  | 
 **inviteToken** | **string** |  | 

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

> AppsAuthenticationOauthApiOauthCallback(ctx, provider).Code(code).State(state).Error_(error_).ErrorDescription(errorDescription).Execute()

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
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationOauthApiOauthCallback(context.Background(), provider).Code(code).State(state).Error_(error_).ErrorDescription(errorDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationOauthApiOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

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
	resp, r, err := apiClient.AuthenticationAPI.AppsAuthenticationSamlApiDetectMethod(context.Background()).DetectMethodRequest(detectMethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationSamlApiDetectMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAuthenticationSamlApiDetectMethod`: DetectMethodResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AppsAuthenticationSamlApiDetectMethod`: %v\n", resp)
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
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationSamlApiInitiate(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationSamlApiInitiate``: %v\n", err)
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
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationSamlApiMetadata(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationSamlApiMetadata``: %v\n", err)
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
	r, err := apiClient.AuthenticationAPI.AppsAuthenticationSamlApiSamlAcs(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AppsAuthenticationSamlApiSamlAcs``: %v\n", err)
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

