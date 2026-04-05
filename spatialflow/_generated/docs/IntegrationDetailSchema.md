# IntegrationDetailSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Description** | **string** |  | 
**IsActive** | **bool** |  | 
**IsVerified** | **bool** |  | 
**HealthStatus** | **string** |  | 
**HealthMessage** | **string** |  | 
**UsageCount** | **int32** |  | 
**LastUsedAt** | Pointer to **NullableString** |  | [optional] 
**LastVerifiedAt** | Pointer to **NullableString** |  | [optional] 
**Tags** | **[]string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Config** | **map[string]interface{}** |  | 

## Methods

### NewIntegrationDetailSchema

`func NewIntegrationDetailSchema(id string, name string, type_ string, description string, isActive bool, isVerified bool, healthStatus string, healthMessage string, usageCount int32, tags []string, createdAt string, updatedAt string, config map[string]interface{}, ) *IntegrationDetailSchema`

NewIntegrationDetailSchema instantiates a new IntegrationDetailSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationDetailSchemaWithDefaults

`func NewIntegrationDetailSchemaWithDefaults() *IntegrationDetailSchema`

NewIntegrationDetailSchemaWithDefaults instantiates a new IntegrationDetailSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationDetailSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationDetailSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationDetailSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IntegrationDetailSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationDetailSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationDetailSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntegrationDetailSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationDetailSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationDetailSchema) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *IntegrationDetailSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationDetailSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationDetailSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsActive

`func (o *IntegrationDetailSchema) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *IntegrationDetailSchema) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *IntegrationDetailSchema) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsVerified

`func (o *IntegrationDetailSchema) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *IntegrationDetailSchema) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *IntegrationDetailSchema) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.


### GetHealthStatus

`func (o *IntegrationDetailSchema) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *IntegrationDetailSchema) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *IntegrationDetailSchema) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.


### GetHealthMessage

`func (o *IntegrationDetailSchema) GetHealthMessage() string`

GetHealthMessage returns the HealthMessage field if non-nil, zero value otherwise.

### GetHealthMessageOk

`func (o *IntegrationDetailSchema) GetHealthMessageOk() (*string, bool)`

GetHealthMessageOk returns a tuple with the HealthMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthMessage

`func (o *IntegrationDetailSchema) SetHealthMessage(v string)`

SetHealthMessage sets HealthMessage field to given value.


### GetUsageCount

`func (o *IntegrationDetailSchema) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *IntegrationDetailSchema) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *IntegrationDetailSchema) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.


### GetLastUsedAt

`func (o *IntegrationDetailSchema) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *IntegrationDetailSchema) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *IntegrationDetailSchema) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *IntegrationDetailSchema) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *IntegrationDetailSchema) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *IntegrationDetailSchema) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetLastVerifiedAt

`func (o *IntegrationDetailSchema) GetLastVerifiedAt() string`

GetLastVerifiedAt returns the LastVerifiedAt field if non-nil, zero value otherwise.

### GetLastVerifiedAtOk

`func (o *IntegrationDetailSchema) GetLastVerifiedAtOk() (*string, bool)`

GetLastVerifiedAtOk returns a tuple with the LastVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedAt

`func (o *IntegrationDetailSchema) SetLastVerifiedAt(v string)`

SetLastVerifiedAt sets LastVerifiedAt field to given value.

### HasLastVerifiedAt

`func (o *IntegrationDetailSchema) HasLastVerifiedAt() bool`

HasLastVerifiedAt returns a boolean if a field has been set.

### SetLastVerifiedAtNil

`func (o *IntegrationDetailSchema) SetLastVerifiedAtNil(b bool)`

 SetLastVerifiedAtNil sets the value for LastVerifiedAt to be an explicit nil

### UnsetLastVerifiedAt
`func (o *IntegrationDetailSchema) UnsetLastVerifiedAt()`

UnsetLastVerifiedAt ensures that no value is present for LastVerifiedAt, not even an explicit nil
### GetTags

`func (o *IntegrationDetailSchema) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IntegrationDetailSchema) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IntegrationDetailSchema) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreatedAt

`func (o *IntegrationDetailSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationDetailSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationDetailSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IntegrationDetailSchema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IntegrationDetailSchema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IntegrationDetailSchema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetConfig

`func (o *IntegrationDetailSchema) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationDetailSchema) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationDetailSchema) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


