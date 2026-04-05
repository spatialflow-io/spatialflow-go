# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Url** | **string** |  | 
**Events** | **[]string** |  | 
**Headers** | **map[string]string** |  | 
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

## Methods

### NewWebhookResponse

`func NewWebhookResponse(id string, name string, description NullableString, url string, events []string, headers map[string]string, authType string, method string, contentType string, customPayloadTemplate NullableString, isActive bool, maxRetries int32, timeoutSeconds int32, rateLimitPerMinute int32, createdAt time.Time, updatedAt time.Time, lastTriggeredAt NullableTime, totalDeliveries int32, successfulDeliveries int32, failedDeliveries int32, successRate NullableFloat32, ) *WebhookResponse`

NewWebhookResponse instantiates a new WebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseWithDefaults

`func NewWebhookResponseWithDefaults() *WebhookResponse`

NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebhookResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WebhookResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WebhookResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *WebhookResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *WebhookResponse) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookResponse) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookResponse) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetHeaders

`func (o *WebhookResponse) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookResponse) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookResponse) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetAuthType

`func (o *WebhookResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WebhookResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WebhookResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetMethod

`func (o *WebhookResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebhookResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebhookResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetContentType

`func (o *WebhookResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WebhookResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WebhookResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetCustomPayloadTemplate

`func (o *WebhookResponse) GetCustomPayloadTemplate() string`

GetCustomPayloadTemplate returns the CustomPayloadTemplate field if non-nil, zero value otherwise.

### GetCustomPayloadTemplateOk

`func (o *WebhookResponse) GetCustomPayloadTemplateOk() (*string, bool)`

GetCustomPayloadTemplateOk returns a tuple with the CustomPayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadTemplate

`func (o *WebhookResponse) SetCustomPayloadTemplate(v string)`

SetCustomPayloadTemplate sets CustomPayloadTemplate field to given value.


### SetCustomPayloadTemplateNil

`func (o *WebhookResponse) SetCustomPayloadTemplateNil(b bool)`

 SetCustomPayloadTemplateNil sets the value for CustomPayloadTemplate to be an explicit nil

### UnsetCustomPayloadTemplate
`func (o *WebhookResponse) UnsetCustomPayloadTemplate()`

UnsetCustomPayloadTemplate ensures that no value is present for CustomPayloadTemplate, not even an explicit nil
### GetIsActive

`func (o *WebhookResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebhookResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebhookResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetMaxRetries

`func (o *WebhookResponse) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *WebhookResponse) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *WebhookResponse) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.


### GetTimeoutSeconds

`func (o *WebhookResponse) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *WebhookResponse) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *WebhookResponse) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetRateLimitPerMinute

`func (o *WebhookResponse) GetRateLimitPerMinute() int32`

GetRateLimitPerMinute returns the RateLimitPerMinute field if non-nil, zero value otherwise.

### GetRateLimitPerMinuteOk

`func (o *WebhookResponse) GetRateLimitPerMinuteOk() (*int32, bool)`

GetRateLimitPerMinuteOk returns a tuple with the RateLimitPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerMinute

`func (o *WebhookResponse) SetRateLimitPerMinute(v int32)`

SetRateLimitPerMinute sets RateLimitPerMinute field to given value.


### GetCreatedAt

`func (o *WebhookResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WebhookResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastTriggeredAt

`func (o *WebhookResponse) GetLastTriggeredAt() time.Time`

GetLastTriggeredAt returns the LastTriggeredAt field if non-nil, zero value otherwise.

### GetLastTriggeredAtOk

`func (o *WebhookResponse) GetLastTriggeredAtOk() (*time.Time, bool)`

GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredAt

`func (o *WebhookResponse) SetLastTriggeredAt(v time.Time)`

SetLastTriggeredAt sets LastTriggeredAt field to given value.


### SetLastTriggeredAtNil

`func (o *WebhookResponse) SetLastTriggeredAtNil(b bool)`

 SetLastTriggeredAtNil sets the value for LastTriggeredAt to be an explicit nil

### UnsetLastTriggeredAt
`func (o *WebhookResponse) UnsetLastTriggeredAt()`

UnsetLastTriggeredAt ensures that no value is present for LastTriggeredAt, not even an explicit nil
### GetTotalDeliveries

`func (o *WebhookResponse) GetTotalDeliveries() int32`

GetTotalDeliveries returns the TotalDeliveries field if non-nil, zero value otherwise.

### GetTotalDeliveriesOk

`func (o *WebhookResponse) GetTotalDeliveriesOk() (*int32, bool)`

GetTotalDeliveriesOk returns a tuple with the TotalDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeliveries

`func (o *WebhookResponse) SetTotalDeliveries(v int32)`

SetTotalDeliveries sets TotalDeliveries field to given value.


### GetSuccessfulDeliveries

`func (o *WebhookResponse) GetSuccessfulDeliveries() int32`

GetSuccessfulDeliveries returns the SuccessfulDeliveries field if non-nil, zero value otherwise.

### GetSuccessfulDeliveriesOk

`func (o *WebhookResponse) GetSuccessfulDeliveriesOk() (*int32, bool)`

GetSuccessfulDeliveriesOk returns a tuple with the SuccessfulDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulDeliveries

`func (o *WebhookResponse) SetSuccessfulDeliveries(v int32)`

SetSuccessfulDeliveries sets SuccessfulDeliveries field to given value.


### GetFailedDeliveries

`func (o *WebhookResponse) GetFailedDeliveries() int32`

GetFailedDeliveries returns the FailedDeliveries field if non-nil, zero value otherwise.

### GetFailedDeliveriesOk

`func (o *WebhookResponse) GetFailedDeliveriesOk() (*int32, bool)`

GetFailedDeliveriesOk returns a tuple with the FailedDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeliveries

`func (o *WebhookResponse) SetFailedDeliveries(v int32)`

SetFailedDeliveries sets FailedDeliveries field to given value.


### GetSuccessRate

`func (o *WebhookResponse) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *WebhookResponse) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *WebhookResponse) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### SetSuccessRateNil

`func (o *WebhookResponse) SetSuccessRateNil(b bool)`

 SetSuccessRateNil sets the value for SuccessRate to be an explicit nil

### UnsetSuccessRate
`func (o *WebhookResponse) UnsetSuccessRate()`

UnsetSuccessRate ensures that no value is present for SuccessRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


