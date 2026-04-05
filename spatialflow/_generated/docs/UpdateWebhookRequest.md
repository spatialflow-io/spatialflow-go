# UpdateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**AuthType** | Pointer to [**NullableAuthTypeEnum**](AuthTypeEnum.md) |  | [optional] 
**AuthConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**Method** | Pointer to [**NullableMethodEnum**](MethodEnum.md) |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**CustomPayloadTemplate** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**MaxRetries** | Pointer to **NullableInt32** |  | [optional] 
**TimeoutSeconds** | Pointer to **NullableInt32** |  | [optional] 
**RateLimitPerMinute** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUpdateWebhookRequest

`func NewUpdateWebhookRequest() *UpdateWebhookRequest`

NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookRequestWithDefaults

`func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest`

NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateWebhookRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateWebhookRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateWebhookRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateWebhookRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *UpdateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateWebhookRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *UpdateWebhookRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *UpdateWebhookRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetEvents

`func (o *UpdateWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UpdateWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UpdateWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UpdateWebhookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *UpdateWebhookRequest) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *UpdateWebhookRequest) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetHeaders

`func (o *UpdateWebhookRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *UpdateWebhookRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *UpdateWebhookRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *UpdateWebhookRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *UpdateWebhookRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *UpdateWebhookRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetAuthType

`func (o *UpdateWebhookRequest) GetAuthType() AuthTypeEnum`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UpdateWebhookRequest) GetAuthTypeOk() (*AuthTypeEnum, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UpdateWebhookRequest) SetAuthType(v AuthTypeEnum)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *UpdateWebhookRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *UpdateWebhookRequest) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *UpdateWebhookRequest) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetAuthConfig

`func (o *UpdateWebhookRequest) GetAuthConfig() map[string]interface{}`

GetAuthConfig returns the AuthConfig field if non-nil, zero value otherwise.

### GetAuthConfigOk

`func (o *UpdateWebhookRequest) GetAuthConfigOk() (*map[string]interface{}, bool)`

GetAuthConfigOk returns a tuple with the AuthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthConfig

`func (o *UpdateWebhookRequest) SetAuthConfig(v map[string]interface{})`

SetAuthConfig sets AuthConfig field to given value.

### HasAuthConfig

`func (o *UpdateWebhookRequest) HasAuthConfig() bool`

HasAuthConfig returns a boolean if a field has been set.

### SetAuthConfigNil

`func (o *UpdateWebhookRequest) SetAuthConfigNil(b bool)`

 SetAuthConfigNil sets the value for AuthConfig to be an explicit nil

### UnsetAuthConfig
`func (o *UpdateWebhookRequest) UnsetAuthConfig()`

UnsetAuthConfig ensures that no value is present for AuthConfig, not even an explicit nil
### GetMethod

`func (o *UpdateWebhookRequest) GetMethod() MethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UpdateWebhookRequest) GetMethodOk() (*MethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UpdateWebhookRequest) SetMethod(v MethodEnum)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UpdateWebhookRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *UpdateWebhookRequest) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *UpdateWebhookRequest) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetContentType

`func (o *UpdateWebhookRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UpdateWebhookRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UpdateWebhookRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *UpdateWebhookRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *UpdateWebhookRequest) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *UpdateWebhookRequest) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetCustomPayloadTemplate

`func (o *UpdateWebhookRequest) GetCustomPayloadTemplate() string`

GetCustomPayloadTemplate returns the CustomPayloadTemplate field if non-nil, zero value otherwise.

### GetCustomPayloadTemplateOk

`func (o *UpdateWebhookRequest) GetCustomPayloadTemplateOk() (*string, bool)`

GetCustomPayloadTemplateOk returns a tuple with the CustomPayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadTemplate

`func (o *UpdateWebhookRequest) SetCustomPayloadTemplate(v string)`

SetCustomPayloadTemplate sets CustomPayloadTemplate field to given value.

### HasCustomPayloadTemplate

`func (o *UpdateWebhookRequest) HasCustomPayloadTemplate() bool`

HasCustomPayloadTemplate returns a boolean if a field has been set.

### SetCustomPayloadTemplateNil

`func (o *UpdateWebhookRequest) SetCustomPayloadTemplateNil(b bool)`

 SetCustomPayloadTemplateNil sets the value for CustomPayloadTemplate to be an explicit nil

### UnsetCustomPayloadTemplate
`func (o *UpdateWebhookRequest) UnsetCustomPayloadTemplate()`

UnsetCustomPayloadTemplate ensures that no value is present for CustomPayloadTemplate, not even an explicit nil
### GetIsActive

`func (o *UpdateWebhookRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateWebhookRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateWebhookRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateWebhookRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateWebhookRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateWebhookRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetMaxRetries

`func (o *UpdateWebhookRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *UpdateWebhookRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *UpdateWebhookRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *UpdateWebhookRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *UpdateWebhookRequest) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *UpdateWebhookRequest) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetTimeoutSeconds

`func (o *UpdateWebhookRequest) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *UpdateWebhookRequest) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *UpdateWebhookRequest) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *UpdateWebhookRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### SetTimeoutSecondsNil

`func (o *UpdateWebhookRequest) SetTimeoutSecondsNil(b bool)`

 SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil

### UnsetTimeoutSeconds
`func (o *UpdateWebhookRequest) UnsetTimeoutSeconds()`

UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil
### GetRateLimitPerMinute

`func (o *UpdateWebhookRequest) GetRateLimitPerMinute() int32`

GetRateLimitPerMinute returns the RateLimitPerMinute field if non-nil, zero value otherwise.

### GetRateLimitPerMinuteOk

`func (o *UpdateWebhookRequest) GetRateLimitPerMinuteOk() (*int32, bool)`

GetRateLimitPerMinuteOk returns a tuple with the RateLimitPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerMinute

`func (o *UpdateWebhookRequest) SetRateLimitPerMinute(v int32)`

SetRateLimitPerMinute sets RateLimitPerMinute field to given value.

### HasRateLimitPerMinute

`func (o *UpdateWebhookRequest) HasRateLimitPerMinute() bool`

HasRateLimitPerMinute returns a boolean if a field has been set.

### SetRateLimitPerMinuteNil

`func (o *UpdateWebhookRequest) SetRateLimitPerMinuteNil(b bool)`

 SetRateLimitPerMinuteNil sets the value for RateLimitPerMinute to be an explicit nil

### UnsetRateLimitPerMinute
`func (o *UpdateWebhookRequest) UnsetRateLimitPerMinute()`

UnsetRateLimitPerMinute ensures that no value is present for RateLimitPerMinute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


