# IntegrationTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Icon** | **string** |  | 
**Category** | **string** |  | 
**HandlerClass** | **string** |  | 
**ValidatorClass** | **NullableString** |  | 
**IsActive** | **bool** |  | 
**IsBuiltin** | **bool** |  | 
**OauthEnabled** | **bool** |  | 
**OauthConfig** | **map[string]interface{}** |  | 
**DocumentationUrl** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ConfigFieldsCount** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewIntegrationTypeResponse

`func NewIntegrationTypeResponse(id string, key string, name string, description string, icon string, category string, handlerClass string, validatorClass NullableString, isActive bool, isBuiltin bool, oauthEnabled bool, oauthConfig map[string]interface{}, documentationUrl NullableString, createdAt time.Time, updatedAt time.Time, ) *IntegrationTypeResponse`

NewIntegrationTypeResponse instantiates a new IntegrationTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationTypeResponseWithDefaults

`func NewIntegrationTypeResponseWithDefaults() *IntegrationTypeResponse`

NewIntegrationTypeResponseWithDefaults instantiates a new IntegrationTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationTypeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationTypeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationTypeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *IntegrationTypeResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IntegrationTypeResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IntegrationTypeResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *IntegrationTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationTypeResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IntegrationTypeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationTypeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationTypeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIcon

`func (o *IntegrationTypeResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IntegrationTypeResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IntegrationTypeResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetCategory

`func (o *IntegrationTypeResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IntegrationTypeResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IntegrationTypeResponse) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetHandlerClass

`func (o *IntegrationTypeResponse) GetHandlerClass() string`

GetHandlerClass returns the HandlerClass field if non-nil, zero value otherwise.

### GetHandlerClassOk

`func (o *IntegrationTypeResponse) GetHandlerClassOk() (*string, bool)`

GetHandlerClassOk returns a tuple with the HandlerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerClass

`func (o *IntegrationTypeResponse) SetHandlerClass(v string)`

SetHandlerClass sets HandlerClass field to given value.


### GetValidatorClass

`func (o *IntegrationTypeResponse) GetValidatorClass() string`

GetValidatorClass returns the ValidatorClass field if non-nil, zero value otherwise.

### GetValidatorClassOk

`func (o *IntegrationTypeResponse) GetValidatorClassOk() (*string, bool)`

GetValidatorClassOk returns a tuple with the ValidatorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorClass

`func (o *IntegrationTypeResponse) SetValidatorClass(v string)`

SetValidatorClass sets ValidatorClass field to given value.


### SetValidatorClassNil

`func (o *IntegrationTypeResponse) SetValidatorClassNil(b bool)`

 SetValidatorClassNil sets the value for ValidatorClass to be an explicit nil

### UnsetValidatorClass
`func (o *IntegrationTypeResponse) UnsetValidatorClass()`

UnsetValidatorClass ensures that no value is present for ValidatorClass, not even an explicit nil
### GetIsActive

`func (o *IntegrationTypeResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *IntegrationTypeResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *IntegrationTypeResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsBuiltin

`func (o *IntegrationTypeResponse) GetIsBuiltin() bool`

GetIsBuiltin returns the IsBuiltin field if non-nil, zero value otherwise.

### GetIsBuiltinOk

`func (o *IntegrationTypeResponse) GetIsBuiltinOk() (*bool, bool)`

GetIsBuiltinOk returns a tuple with the IsBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltin

`func (o *IntegrationTypeResponse) SetIsBuiltin(v bool)`

SetIsBuiltin sets IsBuiltin field to given value.


### GetOauthEnabled

`func (o *IntegrationTypeResponse) GetOauthEnabled() bool`

GetOauthEnabled returns the OauthEnabled field if non-nil, zero value otherwise.

### GetOauthEnabledOk

`func (o *IntegrationTypeResponse) GetOauthEnabledOk() (*bool, bool)`

GetOauthEnabledOk returns a tuple with the OauthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthEnabled

`func (o *IntegrationTypeResponse) SetOauthEnabled(v bool)`

SetOauthEnabled sets OauthEnabled field to given value.


### GetOauthConfig

`func (o *IntegrationTypeResponse) GetOauthConfig() map[string]interface{}`

GetOauthConfig returns the OauthConfig field if non-nil, zero value otherwise.

### GetOauthConfigOk

`func (o *IntegrationTypeResponse) GetOauthConfigOk() (*map[string]interface{}, bool)`

GetOauthConfigOk returns a tuple with the OauthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthConfig

`func (o *IntegrationTypeResponse) SetOauthConfig(v map[string]interface{})`

SetOauthConfig sets OauthConfig field to given value.


### GetDocumentationUrl

`func (o *IntegrationTypeResponse) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *IntegrationTypeResponse) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *IntegrationTypeResponse) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.


### SetDocumentationUrlNil

`func (o *IntegrationTypeResponse) SetDocumentationUrlNil(b bool)`

 SetDocumentationUrlNil sets the value for DocumentationUrl to be an explicit nil

### UnsetDocumentationUrl
`func (o *IntegrationTypeResponse) UnsetDocumentationUrl()`

UnsetDocumentationUrl ensures that no value is present for DocumentationUrl, not even an explicit nil
### GetCreatedAt

`func (o *IntegrationTypeResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationTypeResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationTypeResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IntegrationTypeResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IntegrationTypeResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IntegrationTypeResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetConfigFieldsCount

`func (o *IntegrationTypeResponse) GetConfigFieldsCount() int32`

GetConfigFieldsCount returns the ConfigFieldsCount field if non-nil, zero value otherwise.

### GetConfigFieldsCountOk

`func (o *IntegrationTypeResponse) GetConfigFieldsCountOk() (*int32, bool)`

GetConfigFieldsCountOk returns a tuple with the ConfigFieldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFieldsCount

`func (o *IntegrationTypeResponse) SetConfigFieldsCount(v int32)`

SetConfigFieldsCount sets ConfigFieldsCount field to given value.

### HasConfigFieldsCount

`func (o *IntegrationTypeResponse) HasConfigFieldsCount() bool`

HasConfigFieldsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


