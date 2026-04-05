# IntegrationResponseSchema

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

## Methods

### NewIntegrationResponseSchema

`func NewIntegrationResponseSchema(id string, name string, type_ string, description string, isActive bool, isVerified bool, healthStatus string, healthMessage string, usageCount int32, tags []string, createdAt string, updatedAt string, ) *IntegrationResponseSchema`

NewIntegrationResponseSchema instantiates a new IntegrationResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationResponseSchemaWithDefaults

`func NewIntegrationResponseSchemaWithDefaults() *IntegrationResponseSchema`

NewIntegrationResponseSchemaWithDefaults instantiates a new IntegrationResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationResponseSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationResponseSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationResponseSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IntegrationResponseSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationResponseSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationResponseSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntegrationResponseSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationResponseSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationResponseSchema) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *IntegrationResponseSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationResponseSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationResponseSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsActive

`func (o *IntegrationResponseSchema) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *IntegrationResponseSchema) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *IntegrationResponseSchema) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsVerified

`func (o *IntegrationResponseSchema) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *IntegrationResponseSchema) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *IntegrationResponseSchema) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.


### GetHealthStatus

`func (o *IntegrationResponseSchema) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *IntegrationResponseSchema) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *IntegrationResponseSchema) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.


### GetHealthMessage

`func (o *IntegrationResponseSchema) GetHealthMessage() string`

GetHealthMessage returns the HealthMessage field if non-nil, zero value otherwise.

### GetHealthMessageOk

`func (o *IntegrationResponseSchema) GetHealthMessageOk() (*string, bool)`

GetHealthMessageOk returns a tuple with the HealthMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthMessage

`func (o *IntegrationResponseSchema) SetHealthMessage(v string)`

SetHealthMessage sets HealthMessage field to given value.


### GetUsageCount

`func (o *IntegrationResponseSchema) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *IntegrationResponseSchema) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *IntegrationResponseSchema) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.


### GetLastUsedAt

`func (o *IntegrationResponseSchema) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *IntegrationResponseSchema) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *IntegrationResponseSchema) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *IntegrationResponseSchema) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *IntegrationResponseSchema) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *IntegrationResponseSchema) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetLastVerifiedAt

`func (o *IntegrationResponseSchema) GetLastVerifiedAt() string`

GetLastVerifiedAt returns the LastVerifiedAt field if non-nil, zero value otherwise.

### GetLastVerifiedAtOk

`func (o *IntegrationResponseSchema) GetLastVerifiedAtOk() (*string, bool)`

GetLastVerifiedAtOk returns a tuple with the LastVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedAt

`func (o *IntegrationResponseSchema) SetLastVerifiedAt(v string)`

SetLastVerifiedAt sets LastVerifiedAt field to given value.

### HasLastVerifiedAt

`func (o *IntegrationResponseSchema) HasLastVerifiedAt() bool`

HasLastVerifiedAt returns a boolean if a field has been set.

### SetLastVerifiedAtNil

`func (o *IntegrationResponseSchema) SetLastVerifiedAtNil(b bool)`

 SetLastVerifiedAtNil sets the value for LastVerifiedAt to be an explicit nil

### UnsetLastVerifiedAt
`func (o *IntegrationResponseSchema) UnsetLastVerifiedAt()`

UnsetLastVerifiedAt ensures that no value is present for LastVerifiedAt, not even an explicit nil
### GetTags

`func (o *IntegrationResponseSchema) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IntegrationResponseSchema) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IntegrationResponseSchema) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreatedAt

`func (o *IntegrationResponseSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationResponseSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationResponseSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IntegrationResponseSchema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IntegrationResponseSchema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IntegrationResponseSchema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


