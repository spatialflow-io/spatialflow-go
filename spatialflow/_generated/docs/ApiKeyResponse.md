# ApiKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**KeyPrefix** | **string** |  | 
**Permissions** | **[]string** |  | 
**IsActive** | **bool** |  | 
**UsageCount** | **int32** |  | 
**RateLimitPerHour** | **int32** |  | 
**LastUsedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ReadOnly** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **NullableString** |  | [optional] 
**WorkspaceId** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewApiKeyResponse

`func NewApiKeyResponse(id string, name string, keyPrefix string, permissions []string, isActive bool, usageCount int32, rateLimitPerHour int32, createdAt time.Time, ) *ApiKeyResponse`

NewApiKeyResponse instantiates a new ApiKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyResponseWithDefaults

`func NewApiKeyResponseWithDefaults() *ApiKeyResponse`

NewApiKeyResponseWithDefaults instantiates a new ApiKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiKeyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetKeyPrefix

`func (o *ApiKeyResponse) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *ApiKeyResponse) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *ApiKeyResponse) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.


### GetPermissions

`func (o *ApiKeyResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiKeyResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiKeyResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetIsActive

`func (o *ApiKeyResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ApiKeyResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ApiKeyResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetUsageCount

`func (o *ApiKeyResponse) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *ApiKeyResponse) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *ApiKeyResponse) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.


### GetRateLimitPerHour

`func (o *ApiKeyResponse) GetRateLimitPerHour() int32`

GetRateLimitPerHour returns the RateLimitPerHour field if non-nil, zero value otherwise.

### GetRateLimitPerHourOk

`func (o *ApiKeyResponse) GetRateLimitPerHourOk() (*int32, bool)`

GetRateLimitPerHourOk returns a tuple with the RateLimitPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerHour

`func (o *ApiKeyResponse) SetRateLimitPerHour(v int32)`

SetRateLimitPerHour sets RateLimitPerHour field to given value.


### GetLastUsedAt

`func (o *ApiKeyResponse) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiKeyResponse) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiKeyResponse) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiKeyResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *ApiKeyResponse) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *ApiKeyResponse) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetCreatedAt

`func (o *ApiKeyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKeyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKeyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReadOnly

`func (o *ApiKeyResponse) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiKeyResponse) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiKeyResponse) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiKeyResponse) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiKeyResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiKeyResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWorkspaceId

`func (o *ApiKeyResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ApiKeyResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ApiKeyResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *ApiKeyResponse) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *ApiKeyResponse) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *ApiKeyResponse) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetExpiresAt

`func (o *ApiKeyResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiKeyResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiKeyResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiKeyResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ApiKeyResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ApiKeyResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


