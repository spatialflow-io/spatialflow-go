# ApiKeyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**RateLimitPerHour** | Pointer to **NullableInt32** |  | [optional] 
**ReadOnly** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiKeyUpdateRequest

`func NewApiKeyUpdateRequest() *ApiKeyUpdateRequest`

NewApiKeyUpdateRequest instantiates a new ApiKeyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyUpdateRequestWithDefaults

`func NewApiKeyUpdateRequestWithDefaults() *ApiKeyUpdateRequest`

NewApiKeyUpdateRequestWithDefaults instantiates a new ApiKeyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiKeyUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeyUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiKeyUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiKeyUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermissions

`func (o *ApiKeyUpdateRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiKeyUpdateRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiKeyUpdateRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApiKeyUpdateRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ApiKeyUpdateRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ApiKeyUpdateRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetIsActive

`func (o *ApiKeyUpdateRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ApiKeyUpdateRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ApiKeyUpdateRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ApiKeyUpdateRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ApiKeyUpdateRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ApiKeyUpdateRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetRateLimitPerHour

`func (o *ApiKeyUpdateRequest) GetRateLimitPerHour() int32`

GetRateLimitPerHour returns the RateLimitPerHour field if non-nil, zero value otherwise.

### GetRateLimitPerHourOk

`func (o *ApiKeyUpdateRequest) GetRateLimitPerHourOk() (*int32, bool)`

GetRateLimitPerHourOk returns a tuple with the RateLimitPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerHour

`func (o *ApiKeyUpdateRequest) SetRateLimitPerHour(v int32)`

SetRateLimitPerHour sets RateLimitPerHour field to given value.

### HasRateLimitPerHour

`func (o *ApiKeyUpdateRequest) HasRateLimitPerHour() bool`

HasRateLimitPerHour returns a boolean if a field has been set.

### SetRateLimitPerHourNil

`func (o *ApiKeyUpdateRequest) SetRateLimitPerHourNil(b bool)`

 SetRateLimitPerHourNil sets the value for RateLimitPerHour to be an explicit nil

### UnsetRateLimitPerHour
`func (o *ApiKeyUpdateRequest) UnsetRateLimitPerHour()`

UnsetRateLimitPerHour ensures that no value is present for RateLimitPerHour, not even an explicit nil
### GetReadOnly

`func (o *ApiKeyUpdateRequest) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApiKeyUpdateRequest) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApiKeyUpdateRequest) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApiKeyUpdateRequest) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *ApiKeyUpdateRequest) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *ApiKeyUpdateRequest) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetDescription

`func (o *ApiKeyUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiKeyUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiKeyUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


