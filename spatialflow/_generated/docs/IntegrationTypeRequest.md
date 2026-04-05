# IntegrationTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Unique identifier for this integration type | 
**Name** | **string** | Display name for this integration type | 
**Description** | **string** | User-friendly description | 
**Icon** | Pointer to **string** | Icon identifier for UI display | [optional] [default to "default"]
**Category** | Pointer to **string** | Category: general, communication, cloud, analytics, automation, custom | [optional] [default to "general"]
**HandlerClass** | **string** | Python path to the action handler class | 
**ValidatorClass** | Pointer to **NullableString** |  | [optional] 
**OauthEnabled** | Pointer to **bool** | Whether this integration supports OAuth authentication | [optional] [default to false]
**OauthConfig** | Pointer to **map[string]interface{}** | OAuth configuration | [optional] 
**DocumentationUrl** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **bool** | Whether this integration type is available for use | [optional] [default to true]

## Methods

### NewIntegrationTypeRequest

`func NewIntegrationTypeRequest(key string, name string, description string, handlerClass string, ) *IntegrationTypeRequest`

NewIntegrationTypeRequest instantiates a new IntegrationTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationTypeRequestWithDefaults

`func NewIntegrationTypeRequestWithDefaults() *IntegrationTypeRequest`

NewIntegrationTypeRequestWithDefaults instantiates a new IntegrationTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *IntegrationTypeRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IntegrationTypeRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IntegrationTypeRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *IntegrationTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IntegrationTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIcon

`func (o *IntegrationTypeRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IntegrationTypeRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IntegrationTypeRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IntegrationTypeRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetCategory

`func (o *IntegrationTypeRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IntegrationTypeRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IntegrationTypeRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IntegrationTypeRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetHandlerClass

`func (o *IntegrationTypeRequest) GetHandlerClass() string`

GetHandlerClass returns the HandlerClass field if non-nil, zero value otherwise.

### GetHandlerClassOk

`func (o *IntegrationTypeRequest) GetHandlerClassOk() (*string, bool)`

GetHandlerClassOk returns a tuple with the HandlerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerClass

`func (o *IntegrationTypeRequest) SetHandlerClass(v string)`

SetHandlerClass sets HandlerClass field to given value.


### GetValidatorClass

`func (o *IntegrationTypeRequest) GetValidatorClass() string`

GetValidatorClass returns the ValidatorClass field if non-nil, zero value otherwise.

### GetValidatorClassOk

`func (o *IntegrationTypeRequest) GetValidatorClassOk() (*string, bool)`

GetValidatorClassOk returns a tuple with the ValidatorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorClass

`func (o *IntegrationTypeRequest) SetValidatorClass(v string)`

SetValidatorClass sets ValidatorClass field to given value.

### HasValidatorClass

`func (o *IntegrationTypeRequest) HasValidatorClass() bool`

HasValidatorClass returns a boolean if a field has been set.

### SetValidatorClassNil

`func (o *IntegrationTypeRequest) SetValidatorClassNil(b bool)`

 SetValidatorClassNil sets the value for ValidatorClass to be an explicit nil

### UnsetValidatorClass
`func (o *IntegrationTypeRequest) UnsetValidatorClass()`

UnsetValidatorClass ensures that no value is present for ValidatorClass, not even an explicit nil
### GetOauthEnabled

`func (o *IntegrationTypeRequest) GetOauthEnabled() bool`

GetOauthEnabled returns the OauthEnabled field if non-nil, zero value otherwise.

### GetOauthEnabledOk

`func (o *IntegrationTypeRequest) GetOauthEnabledOk() (*bool, bool)`

GetOauthEnabledOk returns a tuple with the OauthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthEnabled

`func (o *IntegrationTypeRequest) SetOauthEnabled(v bool)`

SetOauthEnabled sets OauthEnabled field to given value.

### HasOauthEnabled

`func (o *IntegrationTypeRequest) HasOauthEnabled() bool`

HasOauthEnabled returns a boolean if a field has been set.

### GetOauthConfig

`func (o *IntegrationTypeRequest) GetOauthConfig() map[string]interface{}`

GetOauthConfig returns the OauthConfig field if non-nil, zero value otherwise.

### GetOauthConfigOk

`func (o *IntegrationTypeRequest) GetOauthConfigOk() (*map[string]interface{}, bool)`

GetOauthConfigOk returns a tuple with the OauthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthConfig

`func (o *IntegrationTypeRequest) SetOauthConfig(v map[string]interface{})`

SetOauthConfig sets OauthConfig field to given value.

### HasOauthConfig

`func (o *IntegrationTypeRequest) HasOauthConfig() bool`

HasOauthConfig returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *IntegrationTypeRequest) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *IntegrationTypeRequest) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *IntegrationTypeRequest) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *IntegrationTypeRequest) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### SetDocumentationUrlNil

`func (o *IntegrationTypeRequest) SetDocumentationUrlNil(b bool)`

 SetDocumentationUrlNil sets the value for DocumentationUrl to be an explicit nil

### UnsetDocumentationUrl
`func (o *IntegrationTypeRequest) UnsetDocumentationUrl()`

UnsetDocumentationUrl ensures that no value is present for DocumentationUrl, not even an explicit nil
### GetIsActive

`func (o *IntegrationTypeRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *IntegrationTypeRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *IntegrationTypeRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *IntegrationTypeRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


