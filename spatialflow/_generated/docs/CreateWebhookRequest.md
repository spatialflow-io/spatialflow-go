# CreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**Events** | Pointer to **[]string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**AuthType** | Pointer to [**AuthTypeEnum**](AuthTypeEnum.md) |  | [optional] [default to NONE]
**AuthConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**Method** | Pointer to [**MethodEnum**](MethodEnum.md) |  | [optional] [default to POST]
**ContentType** | Pointer to **string** |  | [optional] [default to "application/json"]
**CustomPayloadTemplate** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] [default to true]
**MaxRetries** | Pointer to **int32** |  | [optional] [default to 3]
**TimeoutSeconds** | Pointer to **int32** |  | [optional] [default to 30]
**RateLimitPerMinute** | Pointer to **int32** |  | [optional] [default to 60]

## Methods

### NewCreateWebhookRequest

`func NewCreateWebhookRequest(name string, url string, ) *CreateWebhookRequest`

NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestWithDefaults

`func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest`

NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateWebhookRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateWebhookRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *CreateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *CreateWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CreateWebhookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHeaders

`func (o *CreateWebhookRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CreateWebhookRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CreateWebhookRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CreateWebhookRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *CreateWebhookRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *CreateWebhookRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetAuthType

`func (o *CreateWebhookRequest) GetAuthType() AuthTypeEnum`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateWebhookRequest) GetAuthTypeOk() (*AuthTypeEnum, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateWebhookRequest) SetAuthType(v AuthTypeEnum)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *CreateWebhookRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthConfig

`func (o *CreateWebhookRequest) GetAuthConfig() map[string]interface{}`

GetAuthConfig returns the AuthConfig field if non-nil, zero value otherwise.

### GetAuthConfigOk

`func (o *CreateWebhookRequest) GetAuthConfigOk() (*map[string]interface{}, bool)`

GetAuthConfigOk returns a tuple with the AuthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthConfig

`func (o *CreateWebhookRequest) SetAuthConfig(v map[string]interface{})`

SetAuthConfig sets AuthConfig field to given value.

### HasAuthConfig

`func (o *CreateWebhookRequest) HasAuthConfig() bool`

HasAuthConfig returns a boolean if a field has been set.

### SetAuthConfigNil

`func (o *CreateWebhookRequest) SetAuthConfigNil(b bool)`

 SetAuthConfigNil sets the value for AuthConfig to be an explicit nil

### UnsetAuthConfig
`func (o *CreateWebhookRequest) UnsetAuthConfig()`

UnsetAuthConfig ensures that no value is present for AuthConfig, not even an explicit nil
### GetMethod

`func (o *CreateWebhookRequest) GetMethod() MethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CreateWebhookRequest) GetMethodOk() (*MethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CreateWebhookRequest) SetMethod(v MethodEnum)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CreateWebhookRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetContentType

`func (o *CreateWebhookRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateWebhookRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateWebhookRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CreateWebhookRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCustomPayloadTemplate

`func (o *CreateWebhookRequest) GetCustomPayloadTemplate() string`

GetCustomPayloadTemplate returns the CustomPayloadTemplate field if non-nil, zero value otherwise.

### GetCustomPayloadTemplateOk

`func (o *CreateWebhookRequest) GetCustomPayloadTemplateOk() (*string, bool)`

GetCustomPayloadTemplateOk returns a tuple with the CustomPayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadTemplate

`func (o *CreateWebhookRequest) SetCustomPayloadTemplate(v string)`

SetCustomPayloadTemplate sets CustomPayloadTemplate field to given value.

### HasCustomPayloadTemplate

`func (o *CreateWebhookRequest) HasCustomPayloadTemplate() bool`

HasCustomPayloadTemplate returns a boolean if a field has been set.

### SetCustomPayloadTemplateNil

`func (o *CreateWebhookRequest) SetCustomPayloadTemplateNil(b bool)`

 SetCustomPayloadTemplateNil sets the value for CustomPayloadTemplate to be an explicit nil

### UnsetCustomPayloadTemplate
`func (o *CreateWebhookRequest) UnsetCustomPayloadTemplate()`

UnsetCustomPayloadTemplate ensures that no value is present for CustomPayloadTemplate, not even an explicit nil
### GetIsActive

`func (o *CreateWebhookRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateWebhookRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateWebhookRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateWebhookRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMaxRetries

`func (o *CreateWebhookRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *CreateWebhookRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *CreateWebhookRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *CreateWebhookRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *CreateWebhookRequest) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *CreateWebhookRequest) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *CreateWebhookRequest) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *CreateWebhookRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetRateLimitPerMinute

`func (o *CreateWebhookRequest) GetRateLimitPerMinute() int32`

GetRateLimitPerMinute returns the RateLimitPerMinute field if non-nil, zero value otherwise.

### GetRateLimitPerMinuteOk

`func (o *CreateWebhookRequest) GetRateLimitPerMinuteOk() (*int32, bool)`

GetRateLimitPerMinuteOk returns a tuple with the RateLimitPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerMinute

`func (o *CreateWebhookRequest) SetRateLimitPerMinute(v int32)`

SetRateLimitPerMinute sets RateLimitPerMinute field to given value.

### HasRateLimitPerMinute

`func (o *CreateWebhookRequest) HasRateLimitPerMinute() bool`

HasRateLimitPerMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


