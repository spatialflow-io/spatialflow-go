# \AdminAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsAdminPortalApiActivateUser**](AdminAPI.md#AppsAdminPortalApiActivateUser) | **Put** /api/v1/admin/users/{user_id}/activate | Activate User
[**AppsAdminPortalApiAdminHealthCheck**](AdminAPI.md#AppsAdminPortalApiAdminHealthCheck) | **Get** /api/v1/admin/health | Admin Health Check
[**AppsAdminPortalApiAdminPing**](AdminAPI.md#AppsAdminPortalApiAdminPing) | **Get** /api/v1/admin/ping | Admin Ping
[**AppsAdminPortalApiAdminResetPassword**](AdminAPI.md#AppsAdminPortalApiAdminResetPassword) | **Post** /api/v1/admin/users/{user_id}/reset-password | Admin Reset Password
[**AppsAdminPortalApiApproveUser**](AdminAPI.md#AppsAdminPortalApiApproveUser) | **Post** /api/v1/admin/users/{user_id}/approve | Approve User
[**AppsAdminPortalApiDeactivateUser**](AdminAPI.md#AppsAdminPortalApiDeactivateUser) | **Put** /api/v1/admin/users/{user_id}/deactivate | Deactivate User
[**AppsAdminPortalApiDeleteUser**](AdminAPI.md#AppsAdminPortalApiDeleteUser) | **Delete** /api/v1/admin/users/{user_id} | Delete User
[**AppsAdminPortalApiDeleteWorkspace**](AdminAPI.md#AppsAdminPortalApiDeleteWorkspace) | **Delete** /api/v1/admin/workspaces/{workspace_id} | Delete Workspace
[**AppsAdminPortalApiExportMarketingSubscribers**](AdminAPI.md#AppsAdminPortalApiExportMarketingSubscribers) | **Get** /api/v1/admin/users/marketing-subscribers | Export Marketing Subscribers
[**AppsAdminPortalApiGetConfiguration**](AdminAPI.md#AppsAdminPortalApiGetConfiguration) | **Get** /api/v1/admin/configurations/{key} | Get Configuration
[**AppsAdminPortalApiGetDashboardStats**](AdminAPI.md#AppsAdminPortalApiGetDashboardStats) | **Get** /api/v1/admin/dashboard-stats | Get Dashboard Stats
[**AppsAdminPortalApiGetNotificationConfig**](AdminAPI.md#AppsAdminPortalApiGetNotificationConfig) | **Get** /api/v1/admin/slack-config | Get Notification Config
[**AppsAdminPortalApiGetUserDetail**](AdminAPI.md#AppsAdminPortalApiGetUserDetail) | **Get** /api/v1/admin/users/{user_id} | Get User Detail
[**AppsAdminPortalApiGetUserUsage**](AdminAPI.md#AppsAdminPortalApiGetUserUsage) | **Get** /api/v1/admin/users/{user_id}/usage | Get User Usage
[**AppsAdminPortalApiGetUsersWithStats**](AdminAPI.md#AppsAdminPortalApiGetUsersWithStats) | **Get** /api/v1/admin/users/stats | Get Users With Stats
[**AppsAdminPortalApiGetWorkspace**](AdminAPI.md#AppsAdminPortalApiGetWorkspace) | **Get** /api/v1/admin/workspaces/{workspace_id} | Get Workspace
[**AppsAdminPortalApiGetWorkspaceMembers**](AdminAPI.md#AppsAdminPortalApiGetWorkspaceMembers) | **Get** /api/v1/admin/workspaces/{workspace_id}/members | Get Workspace Members
[**AppsAdminPortalApiInviteUser**](AdminAPI.md#AppsAdminPortalApiInviteUser) | **Post** /api/v1/admin/users/invite | Invite User
[**AppsAdminPortalApiListConfigurations**](AdminAPI.md#AppsAdminPortalApiListConfigurations) | **Get** /api/v1/admin/configurations | List Configurations
[**AppsAdminPortalApiListPendingUsers**](AdminAPI.md#AppsAdminPortalApiListPendingUsers) | **Get** /api/v1/admin/users/pending | List Pending Users
[**AppsAdminPortalApiListUsers**](AdminAPI.md#AppsAdminPortalApiListUsers) | **Get** /api/v1/admin/users | List Users
[**AppsAdminPortalApiListWorkspaces**](AdminAPI.md#AppsAdminPortalApiListWorkspaces) | **Get** /api/v1/admin/workspaces | List Workspaces
[**AppsAdminPortalApiRejectUser**](AdminAPI.md#AppsAdminPortalApiRejectUser) | **Post** /api/v1/admin/users/{user_id}/reject | Reject User
[**AppsAdminPortalApiRemoveMember**](AdminAPI.md#AppsAdminPortalApiRemoveMember) | **Delete** /api/v1/admin/workspaces/{workspace_id}/members/{user_id} | Remove Member
[**AppsAdminPortalApiRemoveUserWorkspace**](AdminAPI.md#AppsAdminPortalApiRemoveUserWorkspace) | **Delete** /api/v1/admin/users/{user_id}/workspace | Remove User Workspace
[**AppsAdminPortalApiResendPasswordReset**](AdminAPI.md#AppsAdminPortalApiResendPasswordReset) | **Post** /api/v1/admin/users/{user_id}/resend-reset | Resend Password Reset
[**AppsAdminPortalApiResendVerificationEmail**](AdminAPI.md#AppsAdminPortalApiResendVerificationEmail) | **Post** /api/v1/admin/users/{user_id}/resend-verification | Resend Verification Email
[**AppsAdminPortalApiResetConfiguration**](AdminAPI.md#AppsAdminPortalApiResetConfiguration) | **Delete** /api/v1/admin/configurations/{key} | Reset Configuration
[**AppsAdminPortalApiRevokeInvitation**](AdminAPI.md#AppsAdminPortalApiRevokeInvitation) | **Delete** /api/v1/admin/invitations/{invite_id} | Revoke Invitation
[**AppsAdminPortalApiSendTestEmail**](AdminAPI.md#AppsAdminPortalApiSendTestEmail) | **Post** /api/v1/admin/email/test | Send Test Email
[**AppsAdminPortalApiTestNotification**](AdminAPI.md#AppsAdminPortalApiTestNotification) | **Post** /api/v1/admin/slack-config/test | Test Notification
[**AppsAdminPortalApiUpdateConfiguration**](AdminAPI.md#AppsAdminPortalApiUpdateConfiguration) | **Put** /api/v1/admin/configurations/{key} | Update Configuration
[**AppsAdminPortalApiUpdateMemberRole**](AdminAPI.md#AppsAdminPortalApiUpdateMemberRole) | **Patch** /api/v1/admin/workspaces/{workspace_id}/members/{user_id} | Update Member Role
[**AppsAdminPortalApiUpdateNotificationConfig**](AdminAPI.md#AppsAdminPortalApiUpdateNotificationConfig) | **Put** /api/v1/admin/slack-config | Update Notification Config
[**AppsAdminPortalApiUpdateUserWorkspace**](AdminAPI.md#AppsAdminPortalApiUpdateUserWorkspace) | **Patch** /api/v1/admin/users/{user_id}/workspace | Update User Workspace
[**AppsAdminPortalApiUpdateWorkspace**](AdminAPI.md#AppsAdminPortalApiUpdateWorkspace) | **Put** /api/v1/admin/workspaces/{workspace_id} | Update Workspace



## AppsAdminPortalApiActivateUser

> UserActionResponse AppsAdminPortalApiActivateUser(ctx, userId).Execute()

Activate User



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiActivateUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiActivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiActivateUser`: UserActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiActivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiActivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserActionResponse**](UserActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiAdminHealthCheck

> HealthCheckResponse AppsAdminPortalApiAdminHealthCheck(ctx).Execute()

Admin Health Check



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
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiAdminHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiAdminHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiAdminHealthCheck`: HealthCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiAdminHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiAdminHealthCheckRequest struct via the builder pattern


### Return type

[**HealthCheckResponse**](HealthCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiAdminPing

> PingResponse AppsAdminPortalApiAdminPing(ctx).Execute()

Admin Ping



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
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiAdminPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiAdminPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiAdminPing`: PingResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiAdminPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiAdminPingRequest struct via the builder pattern


### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiAdminResetPassword

> UserActionResponse AppsAdminPortalApiAdminResetPassword(ctx, userId).Execute()

Admin Reset Password



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiAdminResetPassword(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiAdminResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiAdminResetPassword`: UserActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiAdminResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiAdminResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserActionResponse**](UserActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiApproveUser

> UserActionResponse AppsAdminPortalApiApproveUser(ctx, userId).UserApprovalRequest(userApprovalRequest).Execute()

Approve User



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
	userId := "userId_example" // string | 
	userApprovalRequest := *openapiclient.NewUserApprovalRequest() // UserApprovalRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiApproveUser(context.Background(), userId).UserApprovalRequest(userApprovalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiApproveUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiApproveUser`: UserActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiApproveUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiApproveUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userApprovalRequest** | [**UserApprovalRequest**](UserApprovalRequest.md) |  | 

### Return type

[**UserActionResponse**](UserActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiDeactivateUser

> UserActionResponse AppsAdminPortalApiDeactivateUser(ctx, userId).Execute()

Deactivate User



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiDeactivateUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiDeactivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiDeactivateUser`: UserActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiDeactivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiDeactivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserActionResponse**](UserActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiDeleteUser

> map[string]interface{} AppsAdminPortalApiDeleteUser(ctx, userId).ConfirmWorkspaceDeletion(confirmWorkspaceDeletion).Execute()

Delete User



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
	userId := "userId_example" // string | 
	confirmWorkspaceDeletion := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiDeleteUser(context.Background(), userId).ConfirmWorkspaceDeletion(confirmWorkspaceDeletion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiDeleteUser`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiDeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmWorkspaceDeletion** | **bool** |  | [default to false]

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


## AppsAdminPortalApiDeleteWorkspace

> WorkspaceDeleteResponse AppsAdminPortalApiDeleteWorkspace(ctx, workspaceId).Execute()

Delete Workspace



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
	workspaceId := "workspaceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiDeleteWorkspace(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiDeleteWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiDeleteWorkspace`: WorkspaceDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiDeleteWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiDeleteWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceDeleteResponse**](WorkspaceDeleteResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiExportMarketingSubscribers

> MarketingSubscriberExportResponse AppsAdminPortalApiExportMarketingSubscribers(ctx).Format(format).Execute()

Export Marketing Subscribers



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
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiExportMarketingSubscribers(context.Background()).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiExportMarketingSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiExportMarketingSubscribers`: MarketingSubscriberExportResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiExportMarketingSubscribers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiExportMarketingSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**MarketingSubscriberExportResponse**](MarketingSubscriberExportResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiGetConfiguration

> ConfigurationItem AppsAdminPortalApiGetConfiguration(ctx, key).Execute()

Get Configuration



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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiGetConfiguration(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiGetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiGetConfiguration`: ConfigurationItem
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiGetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationItem**](ConfigurationItem.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiGetDashboardStats

> DashboardStatsResponse AppsAdminPortalApiGetDashboardStats(ctx).Execute()

Get Dashboard Stats



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
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiGetDashboardStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiGetDashboardStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiGetDashboardStats`: DashboardStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiGetDashboardStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiGetDashboardStatsRequest struct via the builder pattern


### Return type

[**DashboardStatsResponse**](DashboardStatsResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiGetNotificationConfig

> NotificationConfigResponse AppsAdminPortalApiGetNotificationConfig(ctx).Execute()

Get Notification Config



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
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiGetNotificationConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiGetNotificationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiGetNotificationConfig`: NotificationConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiGetNotificationConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiGetNotificationConfigRequest struct via the builder pattern


### Return type

[**NotificationConfigResponse**](NotificationConfigResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiGetUserDetail

> UserSummary AppsAdminPortalApiGetUserDetail(ctx, userId).Execute()

Get User Detail



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiGetUserDetail(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiGetUserDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiGetUserDetail`: UserSummary
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiGetUserDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiGetUserDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSummary**](UserSummary.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiGetUserUsage

> UserUsageResponse AppsAdminPortalApiGetUserUsage(ctx, userId).Execute()

Get User Usage



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiGetUserUsage(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiGetUserUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiGetUserUsage`: UserUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiGetUserUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiGetUserUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserUsageResponse**](UserUsageResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiGetUsersWithStats

> AdminUserStatsResponse AppsAdminPortalApiGetUsersWithStats(ctx).Days(days).Limit(limit).Offset(offset).Sort(sort).UserIds(userIds).Execute()

Get Users With Stats



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
	days := int32(56) // int32 |  (optional) (default to 30)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)
	sort := "sort_example" // string |  (optional) (default to "api_calls")
	userIds := "userIds_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiGetUsersWithStats(context.Background()).Days(days).Limit(limit).Offset(offset).Sort(sort).UserIds(userIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiGetUsersWithStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiGetUsersWithStats`: AdminUserStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiGetUsersWithStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiGetUsersWithStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** |  | [default to 30]
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **sort** | **string** |  | [default to &quot;api_calls&quot;]
 **userIds** | **string** |  | 

### Return type

[**AdminUserStatsResponse**](AdminUserStatsResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiGetWorkspace

> WorkspaceDetailResponse AppsAdminPortalApiGetWorkspace(ctx, workspaceId).Execute()

Get Workspace



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
	workspaceId := "workspaceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiGetWorkspace(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiGetWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiGetWorkspace`: WorkspaceDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiGetWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiGetWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceDetailResponse**](WorkspaceDetailResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiGetWorkspaceMembers

> WorkspaceMembersResponse AppsAdminPortalApiGetWorkspaceMembers(ctx, workspaceId).Page(page).Limit(limit).Execute()

Get Workspace Members



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
	workspaceId := "workspaceId_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiGetWorkspaceMembers(context.Background(), workspaceId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiGetWorkspaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiGetWorkspaceMembers`: WorkspaceMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiGetWorkspaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiGetWorkspaceMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 20]

### Return type

[**WorkspaceMembersResponse**](WorkspaceMembersResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiInviteUser

> UserInviteResponse AppsAdminPortalApiInviteUser(ctx).UserInviteRequest(userInviteRequest).Execute()

Invite User



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
	userInviteRequest := *openapiclient.NewUserInviteRequest("Email_example") // UserInviteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiInviteUser(context.Background()).UserInviteRequest(userInviteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiInviteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiInviteUser`: UserInviteResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiInviteUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userInviteRequest** | [**UserInviteRequest**](UserInviteRequest.md) |  | 

### Return type

[**UserInviteResponse**](UserInviteResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiListConfigurations

> ConfigurationListResponse AppsAdminPortalApiListConfigurations(ctx).Execute()

List Configurations



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
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiListConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiListConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiListConfigurations`: ConfigurationListResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiListConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiListConfigurationsRequest struct via the builder pattern


### Return type

[**ConfigurationListResponse**](ConfigurationListResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiListPendingUsers

> UserListResponse AppsAdminPortalApiListPendingUsers(ctx).Limit(limit).Cursor(cursor).Execute()

List Pending Users



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
	limit := int32(56) // int32 |  (optional) (default to 20)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiListPendingUsers(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiListPendingUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiListPendingUsers`: UserListResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiListPendingUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiListPendingUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **cursor** | **string** |  | 

### Return type

[**UserListResponse**](UserListResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiListUsers

> UserListResponse AppsAdminPortalApiListUsers(ctx).Page(page).Limit(limit).Search(search).Role(role).Status(status).Execute()

List Users



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
	limit := int32(56) // int32 |  (optional) (default to 20)
	search := "search_example" // string |  (optional)
	role := "role_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiListUsers(context.Background()).Page(page).Limit(limit).Search(search).Role(role).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiListUsers`: UserListResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 20]
 **search** | **string** |  | 
 **role** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**UserListResponse**](UserListResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiListWorkspaces

> WorkspaceListResponse AppsAdminPortalApiListWorkspaces(ctx).Page(page).Limit(limit).Search(search).Execute()

List Workspaces



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
	limit := int32(56) // int32 |  (optional) (default to 20)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiListWorkspaces(context.Background()).Page(page).Limit(limit).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiListWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiListWorkspaces`: WorkspaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiListWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 20]
 **search** | **string** |  | 

### Return type

[**WorkspaceListResponse**](WorkspaceListResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiRejectUser

> UserActionResponse AppsAdminPortalApiRejectUser(ctx, userId).UserRejectionRequest(userRejectionRequest).Execute()

Reject User



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
	userId := "userId_example" // string | 
	userRejectionRequest := *openapiclient.NewUserRejectionRequest("Reason_example") // UserRejectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiRejectUser(context.Background(), userId).UserRejectionRequest(userRejectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiRejectUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiRejectUser`: UserActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiRejectUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiRejectUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRejectionRequest** | [**UserRejectionRequest**](UserRejectionRequest.md) |  | 

### Return type

[**UserActionResponse**](UserActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiRemoveMember

> MemberActionResponse AppsAdminPortalApiRemoveMember(ctx, workspaceId, userId).Execute()

Remove Member



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
	workspaceId := "workspaceId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiRemoveMember(context.Background(), workspaceId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiRemoveMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiRemoveMember`: MemberActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiRemoveMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiRemoveMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MemberActionResponse**](MemberActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiRemoveUserWorkspace

> UserWorkspaceResponse AppsAdminPortalApiRemoveUserWorkspace(ctx, userId).Execute()

Remove User Workspace



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiRemoveUserWorkspace(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiRemoveUserWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiRemoveUserWorkspace`: UserWorkspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiRemoveUserWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiRemoveUserWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserWorkspaceResponse**](UserWorkspaceResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiResendPasswordReset

> UserActionResponse AppsAdminPortalApiResendPasswordReset(ctx, userId).Execute()

Resend Password Reset



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiResendPasswordReset(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiResendPasswordReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiResendPasswordReset`: UserActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiResendPasswordReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiResendPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserActionResponse**](UserActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiResendVerificationEmail

> UserActionResponse AppsAdminPortalApiResendVerificationEmail(ctx, userId).Execute()

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiResendVerificationEmail(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiResendVerificationEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiResendVerificationEmail`: UserActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiResendVerificationEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiResendVerificationEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserActionResponse**](UserActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiResetConfiguration

> ConfigurationResetResponse AppsAdminPortalApiResetConfiguration(ctx, key).Execute()

Reset Configuration



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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiResetConfiguration(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiResetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiResetConfiguration`: ConfigurationResetResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiResetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiResetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationResetResponse**](ConfigurationResetResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiRevokeInvitation

> map[string]interface{} AppsAdminPortalApiRevokeInvitation(ctx, inviteId).Execute()

Revoke Invitation



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
	inviteId := "inviteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiRevokeInvitation(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiRevokeInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiRevokeInvitation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiRevokeInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiRevokeInvitationRequest struct via the builder pattern


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


## AppsAdminPortalApiSendTestEmail

> EmailTestResponse AppsAdminPortalApiSendTestEmail(ctx).Execute()

Send Test Email



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
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiSendTestEmail(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiSendTestEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiSendTestEmail`: EmailTestResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiSendTestEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiSendTestEmailRequest struct via the builder pattern


### Return type

[**EmailTestResponse**](EmailTestResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiTestNotification

> NotificationTestResponse AppsAdminPortalApiTestNotification(ctx).Execute()

Test Notification



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
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiTestNotification(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiTestNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiTestNotification`: NotificationTestResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiTestNotification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiTestNotificationRequest struct via the builder pattern


### Return type

[**NotificationTestResponse**](NotificationTestResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiUpdateConfiguration

> ConfigurationUpdateResponse AppsAdminPortalApiUpdateConfiguration(ctx, key).ConfigurationUpdateRequest(configurationUpdateRequest).Execute()

Update Configuration



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
	key := "key_example" // string | 
	configurationUpdateRequest := *openapiclient.NewConfigurationUpdateRequest(interface{}(123)) // ConfigurationUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiUpdateConfiguration(context.Background(), key).ConfigurationUpdateRequest(configurationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiUpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiUpdateConfiguration`: ConfigurationUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiUpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationUpdateRequest** | [**ConfigurationUpdateRequest**](ConfigurationUpdateRequest.md) |  | 

### Return type

[**ConfigurationUpdateResponse**](ConfigurationUpdateResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiUpdateMemberRole

> MemberActionResponse AppsAdminPortalApiUpdateMemberRole(ctx, workspaceId, userId).UpdateMemberRoleRequest(updateMemberRoleRequest).Execute()

Update Member Role



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
	workspaceId := "workspaceId_example" // string | 
	userId := "userId_example" // string | 
	updateMemberRoleRequest := *openapiclient.NewUpdateMemberRoleRequest("Role_example") // UpdateMemberRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiUpdateMemberRole(context.Background(), workspaceId, userId).UpdateMemberRoleRequest(updateMemberRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiUpdateMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiUpdateMemberRole`: MemberActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiUpdateMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiUpdateMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMemberRoleRequest** | [**UpdateMemberRoleRequest**](UpdateMemberRoleRequest.md) |  | 

### Return type

[**MemberActionResponse**](MemberActionResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiUpdateNotificationConfig

> NotificationConfigResponse AppsAdminPortalApiUpdateNotificationConfig(ctx).NotificationConfigUpdateRequest(notificationConfigUpdateRequest).Execute()

Update Notification Config



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
	notificationConfigUpdateRequest := *openapiclient.NewNotificationConfigUpdateRequest() // NotificationConfigUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiUpdateNotificationConfig(context.Background()).NotificationConfigUpdateRequest(notificationConfigUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiUpdateNotificationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiUpdateNotificationConfig`: NotificationConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiUpdateNotificationConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiUpdateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationConfigUpdateRequest** | [**NotificationConfigUpdateRequest**](NotificationConfigUpdateRequest.md) |  | 

### Return type

[**NotificationConfigResponse**](NotificationConfigResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiUpdateUserWorkspace

> UserWorkspaceResponse AppsAdminPortalApiUpdateUserWorkspace(ctx, userId).UpdateUserWorkspaceRequest(updateUserWorkspaceRequest).Execute()

Update User Workspace



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
	userId := "userId_example" // string | 
	updateUserWorkspaceRequest := *openapiclient.NewUpdateUserWorkspaceRequest("WorkspaceId_example") // UpdateUserWorkspaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiUpdateUserWorkspace(context.Background(), userId).UpdateUserWorkspaceRequest(updateUserWorkspaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiUpdateUserWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiUpdateUserWorkspace`: UserWorkspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiUpdateUserWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiUpdateUserWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserWorkspaceRequest** | [**UpdateUserWorkspaceRequest**](UpdateUserWorkspaceRequest.md) |  | 

### Return type

[**UserWorkspaceResponse**](UserWorkspaceResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAdminPortalApiUpdateWorkspace

> WorkspaceUpdateResponse AppsAdminPortalApiUpdateWorkspace(ctx, workspaceId).WorkspaceUpdateRequest(workspaceUpdateRequest).Execute()

Update Workspace



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
	workspaceId := "workspaceId_example" // string | 
	workspaceUpdateRequest := *openapiclient.NewWorkspaceUpdateRequest() // WorkspaceUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AppsAdminPortalApiUpdateWorkspace(context.Background(), workspaceId).WorkspaceUpdateRequest(workspaceUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AppsAdminPortalApiUpdateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAdminPortalApiUpdateWorkspace`: WorkspaceUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AppsAdminPortalApiUpdateWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAdminPortalApiUpdateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceUpdateRequest** | [**WorkspaceUpdateRequest**](WorkspaceUpdateRequest.md) |  | 

### Return type

[**WorkspaceUpdateResponse**](WorkspaceUpdateResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

