# \WorkspacesAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsWorkspacesApiCancelInvitation**](WorkspacesAPI.md#AppsWorkspacesApiCancelInvitation) | **Delete** /api/v1/workspaces/invitations/{invite_id} | Cancel Invitation
[**AppsWorkspacesApiCreateInvitation**](WorkspacesAPI.md#AppsWorkspacesApiCreateInvitation) | **Post** /api/v1/workspaces/invitations | Create Invitation
[**AppsWorkspacesApiDeleteSamlConfig**](WorkspacesAPI.md#AppsWorkspacesApiDeleteSamlConfig) | **Delete** /api/v1/workspaces/saml-config | Delete Saml Config
[**AppsWorkspacesApiGetSamlConfig**](WorkspacesAPI.md#AppsWorkspacesApiGetSamlConfig) | **Get** /api/v1/workspaces/saml-config | Get Saml Config
[**AppsWorkspacesApiGetWorkspace**](WorkspacesAPI.md#AppsWorkspacesApiGetWorkspace) | **Get** /api/v1/workspaces/ | Get Workspace
[**AppsWorkspacesApiGetWorkspaceUsage**](WorkspacesAPI.md#AppsWorkspacesApiGetWorkspaceUsage) | **Get** /api/v1/workspaces/usage | Get Workspace Usage
[**AppsWorkspacesApiListInvitations**](WorkspacesAPI.md#AppsWorkspacesApiListInvitations) | **Get** /api/v1/workspaces/invitations | List Invitations
[**AppsWorkspacesApiListWorkspaceMembers**](WorkspacesAPI.md#AppsWorkspacesApiListWorkspaceMembers) | **Get** /api/v1/workspaces/members | List Workspace Members
[**AppsWorkspacesApiRemoveMember**](WorkspacesAPI.md#AppsWorkspacesApiRemoveMember) | **Delete** /api/v1/workspaces/members/{user_id} | Remove Member
[**AppsWorkspacesApiResendInvitation**](WorkspacesAPI.md#AppsWorkspacesApiResendInvitation) | **Post** /api/v1/workspaces/invitations/{invite_id}/resend | Resend Invitation
[**AppsWorkspacesApiRevokeAllWorkspaceSessions**](WorkspacesAPI.md#AppsWorkspacesApiRevokeAllWorkspaceSessions) | **Post** /api/v1/workspaces/revoke-all-sessions | Revoke All Workspace Sessions
[**AppsWorkspacesApiUpdateMemberRole**](WorkspacesAPI.md#AppsWorkspacesApiUpdateMemberRole) | **Patch** /api/v1/workspaces/members/{user_id} | Update Member Role
[**AppsWorkspacesApiUpdateWorkspace**](WorkspacesAPI.md#AppsWorkspacesApiUpdateWorkspace) | **Put** /api/v1/workspaces/ | Update Workspace
[**AppsWorkspacesApiUpsertSamlConfig**](WorkspacesAPI.md#AppsWorkspacesApiUpsertSamlConfig) | **Put** /api/v1/workspaces/saml-config | Upsert Saml Config



## AppsWorkspacesApiCancelInvitation

> map[string]interface{} AppsWorkspacesApiCancelInvitation(ctx, inviteId).Execute()

Cancel Invitation



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
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiCancelInvitation(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiCancelInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiCancelInvitation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiCancelInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiCancelInvitationRequest struct via the builder pattern


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


## AppsWorkspacesApiCreateInvitation

> InvitationOut AppsWorkspacesApiCreateInvitation(ctx).CreateInvitationIn(createInvitationIn).Execute()

Create Invitation



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
	createInvitationIn := *openapiclient.NewCreateInvitationIn("Email_example") // CreateInvitationIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiCreateInvitation(context.Background()).CreateInvitationIn(createInvitationIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiCreateInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiCreateInvitation`: InvitationOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiCreateInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiCreateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvitationIn** | [**CreateInvitationIn**](CreateInvitationIn.md) |  | 

### Return type

[**InvitationOut**](InvitationOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiDeleteSamlConfig

> map[string]interface{} AppsWorkspacesApiDeleteSamlConfig(ctx).Execute()

Delete Saml Config



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
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiDeleteSamlConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiDeleteSamlConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiDeleteSamlConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiDeleteSamlConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiDeleteSamlConfigRequest struct via the builder pattern


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


## AppsWorkspacesApiGetSamlConfig

> SAMLConfigOut AppsWorkspacesApiGetSamlConfig(ctx).Execute()

Get Saml Config



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
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiGetSamlConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiGetSamlConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiGetSamlConfig`: SAMLConfigOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiGetSamlConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiGetSamlConfigRequest struct via the builder pattern


### Return type

[**SAMLConfigOut**](SAMLConfigOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiGetWorkspace

> WorkspaceOut AppsWorkspacesApiGetWorkspace(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiGetWorkspace(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiGetWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiGetWorkspace`: WorkspaceOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiGetWorkspace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiGetWorkspaceRequest struct via the builder pattern


### Return type

[**WorkspaceOut**](WorkspaceOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiGetWorkspaceUsage

> UsageResponse AppsWorkspacesApiGetWorkspaceUsage(ctx).Execute()

Get Workspace Usage



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
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiGetWorkspaceUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiGetWorkspaceUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiGetWorkspaceUsage`: UsageResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiGetWorkspaceUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiGetWorkspaceUsageRequest struct via the builder pattern


### Return type

[**UsageResponse**](UsageResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiListInvitations

> InvitationListResponse AppsWorkspacesApiListInvitations(ctx).Execute()

List Invitations



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
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiListInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiListInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiListInvitations`: InvitationListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiListInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiListInvitationsRequest struct via the builder pattern


### Return type

[**InvitationListResponse**](InvitationListResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiListWorkspaceMembers

> MemberListResponse AppsWorkspacesApiListWorkspaceMembers(ctx).Execute()

List Workspace Members



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
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiListWorkspaceMembers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiListWorkspaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiListWorkspaceMembers`: MemberListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiListWorkspaceMembers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiListWorkspaceMembersRequest struct via the builder pattern


### Return type

[**MemberListResponse**](MemberListResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiRemoveMember

> MemberActionOut AppsWorkspacesApiRemoveMember(ctx, userId).Execute()

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiRemoveMember(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiRemoveMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiRemoveMember`: MemberActionOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiRemoveMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiRemoveMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemberActionOut**](MemberActionOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiResendInvitation

> InvitationOut AppsWorkspacesApiResendInvitation(ctx, inviteId).Execute()

Resend Invitation



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
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiResendInvitation(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiResendInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiResendInvitation`: InvitationOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiResendInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiResendInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvitationOut**](InvitationOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiRevokeAllWorkspaceSessions

> RevokeAllSessionsOut AppsWorkspacesApiRevokeAllWorkspaceSessions(ctx).Execute()

Revoke All Workspace Sessions



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
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiRevokeAllWorkspaceSessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiRevokeAllWorkspaceSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiRevokeAllWorkspaceSessions`: RevokeAllSessionsOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiRevokeAllWorkspaceSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiRevokeAllWorkspaceSessionsRequest struct via the builder pattern


### Return type

[**RevokeAllSessionsOut**](RevokeAllSessionsOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiUpdateMemberRole

> MemberActionOut AppsWorkspacesApiUpdateMemberRole(ctx, userId).UpdateMemberRoleIn(updateMemberRoleIn).Execute()

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
	userId := "userId_example" // string | 
	updateMemberRoleIn := *openapiclient.NewUpdateMemberRoleIn("Role_example") // UpdateMemberRoleIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiUpdateMemberRole(context.Background(), userId).UpdateMemberRoleIn(updateMemberRoleIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiUpdateMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiUpdateMemberRole`: MemberActionOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiUpdateMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiUpdateMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMemberRoleIn** | [**UpdateMemberRoleIn**](UpdateMemberRoleIn.md) |  | 

### Return type

[**MemberActionOut**](MemberActionOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiUpdateWorkspace

> WorkspaceOut AppsWorkspacesApiUpdateWorkspace(ctx).WorkspaceIn(workspaceIn).Execute()

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
	workspaceIn := *openapiclient.NewWorkspaceIn("Name_example") // WorkspaceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiUpdateWorkspace(context.Background()).WorkspaceIn(workspaceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiUpdateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiUpdateWorkspace`: WorkspaceOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiUpdateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiUpdateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceIn** | [**WorkspaceIn**](WorkspaceIn.md) |  | 

### Return type

[**WorkspaceOut**](WorkspaceOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsWorkspacesApiUpsertSamlConfig

> SAMLConfigOut AppsWorkspacesApiUpsertSamlConfig(ctx).SAMLConfigIn(sAMLConfigIn).Execute()

Upsert Saml Config



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
	sAMLConfigIn := *openapiclient.NewSAMLConfigIn("EntityId_example", "SsoUrl_example", "Certificate_example", "CoveredDomain_example") // SAMLConfigIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.AppsWorkspacesApiUpsertSamlConfig(context.Background()).SAMLConfigIn(sAMLConfigIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.AppsWorkspacesApiUpsertSamlConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsWorkspacesApiUpsertSamlConfig`: SAMLConfigOut
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.AppsWorkspacesApiUpsertSamlConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsWorkspacesApiUpsertSamlConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sAMLConfigIn** | [**SAMLConfigIn**](SAMLConfigIn.md) |  | 

### Return type

[**SAMLConfigOut**](SAMLConfigOut.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

