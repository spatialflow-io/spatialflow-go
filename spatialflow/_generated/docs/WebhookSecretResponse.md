# WebhookSecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Url** | **string** |  | 
**Events** | **[]string** |  | 
**Headers** | **map[string]string** |  | 
**SensitiveHeadersConfigured** | Pointer to **[]string** |  | [optional] 
**AuthType** | **string** |  | 
**Method** | **string** |  | 
**ContentType** | **string** |  | 
**CustomPayloadTemplate** | **NullableString** |  | 
**IsActive** | **bool** |  | 
**MaxRetries** | **int32** |  | 
**TimeoutSeconds** | **int32** |  | 
**RateLimitPerMinute** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**LastTriggeredAt** | **NullableTime** |  | 
**TotalDeliveries** | **int32** |  | 
**SuccessfulDeliveries** | **int32** |  | 
**FailedDeliveries** | **int32** |  | 
**SuccessRate** | **NullableFloat32** |  | 
**AttachedGeofenceCount** | Pointer to **int32** |  | [optional] [default to 0]
**Secret** | **string** | HMAC-SHA256 key that signs every delivery&#39;s X-SF-Signature header. Shown only in this response; store it now, or rotate the secret to get a new one. | 

## Methods

### NewWebhookSecretResponse

`func NewWebhookSecretResponse(id string, name string, description NullableString, url string, events []string, headers map[string]string, authType string, method string, contentType string, customPayloadTemplate NullableString, isActive bool, maxRetries int32, timeoutSeconds int32, rateLimitPerMinute int32, createdAt time.Time, updatedAt time.Time, lastTriggeredAt NullableTime, totalDeliveries int32, successfulDeliveries int32, failedDeliveries int32, successRate NullableFloat32, secret string, ) *WebhookSecretResponse`

NewWebhookSecretResponse instantiates a new WebhookSecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSecretResponseWithDefaults

`func NewWebhookSecretResponseWithDefaults() *WebhookSecretResponse`

NewWebhookSecretResponseWithDefaults instantiates a new WebhookSecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookSecretResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookSecretResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookSecretResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebhookSecretResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookSecretResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookSecretResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WebhookSecretResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookSecretResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookSecretResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WebhookSecretResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookSecretResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *WebhookSecretResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookSecretResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookSecretResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *WebhookSecretResponse) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookSecretResponse) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookSecretResponse) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetHeaders

`func (o *WebhookSecretResponse) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookSecretResponse) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookSecretResponse) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetSensitiveHeadersConfigured

`func (o *WebhookSecretResponse) GetSensitiveHeadersConfigured() []string`

GetSensitiveHeadersConfigured returns the SensitiveHeadersConfigured field if non-nil, zero value otherwise.

### GetSensitiveHeadersConfiguredOk

`func (o *WebhookSecretResponse) GetSensitiveHeadersConfiguredOk() (*[]string, bool)`

GetSensitiveHeadersConfiguredOk returns a tuple with the SensitiveHeadersConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveHeadersConfigured

`func (o *WebhookSecretResponse) SetSensitiveHeadersConfigured(v []string)`

SetSensitiveHeadersConfigured sets SensitiveHeadersConfigured field to given value.

### HasSensitiveHeadersConfigured

`func (o *WebhookSecretResponse) HasSensitiveHeadersConfigured() bool`

HasSensitiveHeadersConfigured returns a boolean if a field has been set.

### GetAuthType

`func (o *WebhookSecretResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WebhookSecretResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WebhookSecretResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetMethod

`func (o *WebhookSecretResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebhookSecretResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebhookSecretResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetContentType

`func (o *WebhookSecretResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WebhookSecretResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WebhookSecretResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetCustomPayloadTemplate

`func (o *WebhookSecretResponse) GetCustomPayloadTemplate() string`

GetCustomPayloadTemplate returns the CustomPayloadTemplate field if non-nil, zero value otherwise.

### GetCustomPayloadTemplateOk

`func (o *WebhookSecretResponse) GetCustomPayloadTemplateOk() (*string, bool)`

GetCustomPayloadTemplateOk returns a tuple with the CustomPayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadTemplate

`func (o *WebhookSecretResponse) SetCustomPayloadTemplate(v string)`

SetCustomPayloadTemplate sets CustomPayloadTemplate field to given value.


### SetCustomPayloadTemplateNil

`func (o *WebhookSecretResponse) SetCustomPayloadTemplateNil(b bool)`

 SetCustomPayloadTemplateNil sets the value for CustomPayloadTemplate to be an explicit nil

### UnsetCustomPayloadTemplate
`func (o *WebhookSecretResponse) UnsetCustomPayloadTemplate()`

UnsetCustomPayloadTemplate ensures that no value is present for CustomPayloadTemplate, not even an explicit nil
### GetIsActive

`func (o *WebhookSecretResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebhookSecretResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebhookSecretResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetMaxRetries

`func (o *WebhookSecretResponse) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *WebhookSecretResponse) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *WebhookSecretResponse) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.


### GetTimeoutSeconds

`func (o *WebhookSecretResponse) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *WebhookSecretResponse) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *WebhookSecretResponse) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetRateLimitPerMinute

`func (o *WebhookSecretResponse) GetRateLimitPerMinute() int32`

GetRateLimitPerMinute returns the RateLimitPerMinute field if non-nil, zero value otherwise.

### GetRateLimitPerMinuteOk

`func (o *WebhookSecretResponse) GetRateLimitPerMinuteOk() (*int32, bool)`

GetRateLimitPerMinuteOk returns a tuple with the RateLimitPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerMinute

`func (o *WebhookSecretResponse) SetRateLimitPerMinute(v int32)`

SetRateLimitPerMinute sets RateLimitPerMinute field to given value.


### GetCreatedAt

`func (o *WebhookSecretResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookSecretResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookSecretResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WebhookSecretResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookSecretResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookSecretResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastTriggeredAt

`func (o *WebhookSecretResponse) GetLastTriggeredAt() time.Time`

GetLastTriggeredAt returns the LastTriggeredAt field if non-nil, zero value otherwise.

### GetLastTriggeredAtOk

`func (o *WebhookSecretResponse) GetLastTriggeredAtOk() (*time.Time, bool)`

GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredAt

`func (o *WebhookSecretResponse) SetLastTriggeredAt(v time.Time)`

SetLastTriggeredAt sets LastTriggeredAt field to given value.


### SetLastTriggeredAtNil

`func (o *WebhookSecretResponse) SetLastTriggeredAtNil(b bool)`

 SetLastTriggeredAtNil sets the value for LastTriggeredAt to be an explicit nil

### UnsetLastTriggeredAt
`func (o *WebhookSecretResponse) UnsetLastTriggeredAt()`

UnsetLastTriggeredAt ensures that no value is present for LastTriggeredAt, not even an explicit nil
### GetTotalDeliveries

`func (o *WebhookSecretResponse) GetTotalDeliveries() int32`

GetTotalDeliveries returns the TotalDeliveries field if non-nil, zero value otherwise.

### GetTotalDeliveriesOk

`func (o *WebhookSecretResponse) GetTotalDeliveriesOk() (*int32, bool)`

GetTotalDeliveriesOk returns a tuple with the TotalDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeliveries

`func (o *WebhookSecretResponse) SetTotalDeliveries(v int32)`

SetTotalDeliveries sets TotalDeliveries field to given value.


### GetSuccessfulDeliveries

`func (o *WebhookSecretResponse) GetSuccessfulDeliveries() int32`

GetSuccessfulDeliveries returns the SuccessfulDeliveries field if non-nil, zero value otherwise.

### GetSuccessfulDeliveriesOk

`func (o *WebhookSecretResponse) GetSuccessfulDeliveriesOk() (*int32, bool)`

GetSuccessfulDeliveriesOk returns a tuple with the SuccessfulDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulDeliveries

`func (o *WebhookSecretResponse) SetSuccessfulDeliveries(v int32)`

SetSuccessfulDeliveries sets SuccessfulDeliveries field to given value.


### GetFailedDeliveries

`func (o *WebhookSecretResponse) GetFailedDeliveries() int32`

GetFailedDeliveries returns the FailedDeliveries field if non-nil, zero value otherwise.

### GetFailedDeliveriesOk

`func (o *WebhookSecretResponse) GetFailedDeliveriesOk() (*int32, bool)`

GetFailedDeliveriesOk returns a tuple with the FailedDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeliveries

`func (o *WebhookSecretResponse) SetFailedDeliveries(v int32)`

SetFailedDeliveries sets FailedDeliveries field to given value.


### GetSuccessRate

`func (o *WebhookSecretResponse) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *WebhookSecretResponse) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *WebhookSecretResponse) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### SetSuccessRateNil

`func (o *WebhookSecretResponse) SetSuccessRateNil(b bool)`

 SetSuccessRateNil sets the value for SuccessRate to be an explicit nil

### UnsetSuccessRate
`func (o *WebhookSecretResponse) UnsetSuccessRate()`

UnsetSuccessRate ensures that no value is present for SuccessRate, not even an explicit nil
### GetAttachedGeofenceCount

`func (o *WebhookSecretResponse) GetAttachedGeofenceCount() int32`

GetAttachedGeofenceCount returns the AttachedGeofenceCount field if non-nil, zero value otherwise.

### GetAttachedGeofenceCountOk

`func (o *WebhookSecretResponse) GetAttachedGeofenceCountOk() (*int32, bool)`

GetAttachedGeofenceCountOk returns a tuple with the AttachedGeofenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedGeofenceCount

`func (o *WebhookSecretResponse) SetAttachedGeofenceCount(v int32)`

SetAttachedGeofenceCount sets AttachedGeofenceCount field to given value.

### HasAttachedGeofenceCount

`func (o *WebhookSecretResponse) HasAttachedGeofenceCount() bool`

HasAttachedGeofenceCount returns a boolean if a field has been set.

### GetSecret

`func (o *WebhookSecretResponse) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookSecretResponse) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookSecretResponse) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


