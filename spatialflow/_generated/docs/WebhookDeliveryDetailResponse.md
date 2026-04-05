# WebhookDeliveryDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WebhookId** | **NullableString** |  | 
**GeofenceId** | **NullableString** |  | 
**EventType** | **string** |  | 
**EventId** | **string** |  | 
**Url** | **string** |  | 
**Method** | **string** |  | 
**Status** | [**DeliveryStatusEnum**](DeliveryStatusEnum.md) |  | 
**ResponseStatusCode** | **NullableInt32** |  | 
**ResponseStatus** | Pointer to **NullableInt32** |  | [optional] 
**ResponseTimeMs** | **NullableFloat32** |  | 
**ErrorMessage** | **NullableString** |  | 
**AttemptCount** | **int32** |  | 
**RetryCount** | Pointer to **int32** |  | [optional] [default to 0]
**CreatedAt** | **time.Time** |  | 
**DeliveredAt** | **NullableTime** |  | 
**NextRetryAt** | **NullableTime** |  | 
**Payload** | **map[string]interface{}** |  | 
**ResponseBody** | **map[string]interface{}** |  | 
**ResponseHeaders** | **map[string]interface{}** |  | 

## Methods

### NewWebhookDeliveryDetailResponse

`func NewWebhookDeliveryDetailResponse(id string, webhookId NullableString, geofenceId NullableString, eventType string, eventId string, url string, method string, status DeliveryStatusEnum, responseStatusCode NullableInt32, responseTimeMs NullableFloat32, errorMessage NullableString, attemptCount int32, createdAt time.Time, deliveredAt NullableTime, nextRetryAt NullableTime, payload map[string]interface{}, responseBody map[string]interface{}, responseHeaders map[string]interface{}, ) *WebhookDeliveryDetailResponse`

NewWebhookDeliveryDetailResponse instantiates a new WebhookDeliveryDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryDetailResponseWithDefaults

`func NewWebhookDeliveryDetailResponseWithDefaults() *WebhookDeliveryDetailResponse`

NewWebhookDeliveryDetailResponseWithDefaults instantiates a new WebhookDeliveryDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookDeliveryDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDeliveryDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDeliveryDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetWebhookId

`func (o *WebhookDeliveryDetailResponse) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookDeliveryDetailResponse) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookDeliveryDetailResponse) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### SetWebhookIdNil

`func (o *WebhookDeliveryDetailResponse) SetWebhookIdNil(b bool)`

 SetWebhookIdNil sets the value for WebhookId to be an explicit nil

### UnsetWebhookId
`func (o *WebhookDeliveryDetailResponse) UnsetWebhookId()`

UnsetWebhookId ensures that no value is present for WebhookId, not even an explicit nil
### GetGeofenceId

`func (o *WebhookDeliveryDetailResponse) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *WebhookDeliveryDetailResponse) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *WebhookDeliveryDetailResponse) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.


### SetGeofenceIdNil

`func (o *WebhookDeliveryDetailResponse) SetGeofenceIdNil(b bool)`

 SetGeofenceIdNil sets the value for GeofenceId to be an explicit nil

### UnsetGeofenceId
`func (o *WebhookDeliveryDetailResponse) UnsetGeofenceId()`

UnsetGeofenceId ensures that no value is present for GeofenceId, not even an explicit nil
### GetEventType

`func (o *WebhookDeliveryDetailResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookDeliveryDetailResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookDeliveryDetailResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventId

`func (o *WebhookDeliveryDetailResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookDeliveryDetailResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookDeliveryDetailResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetUrl

`func (o *WebhookDeliveryDetailResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookDeliveryDetailResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookDeliveryDetailResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *WebhookDeliveryDetailResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebhookDeliveryDetailResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebhookDeliveryDetailResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetStatus

`func (o *WebhookDeliveryDetailResponse) GetStatus() DeliveryStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookDeliveryDetailResponse) GetStatusOk() (*DeliveryStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookDeliveryDetailResponse) SetStatus(v DeliveryStatusEnum)`

SetStatus sets Status field to given value.


### GetResponseStatusCode

`func (o *WebhookDeliveryDetailResponse) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *WebhookDeliveryDetailResponse) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *WebhookDeliveryDetailResponse) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.


### SetResponseStatusCodeNil

`func (o *WebhookDeliveryDetailResponse) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *WebhookDeliveryDetailResponse) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetResponseStatus

`func (o *WebhookDeliveryDetailResponse) GetResponseStatus() int32`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *WebhookDeliveryDetailResponse) GetResponseStatusOk() (*int32, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *WebhookDeliveryDetailResponse) SetResponseStatus(v int32)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *WebhookDeliveryDetailResponse) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### SetResponseStatusNil

`func (o *WebhookDeliveryDetailResponse) SetResponseStatusNil(b bool)`

 SetResponseStatusNil sets the value for ResponseStatus to be an explicit nil

### UnsetResponseStatus
`func (o *WebhookDeliveryDetailResponse) UnsetResponseStatus()`

UnsetResponseStatus ensures that no value is present for ResponseStatus, not even an explicit nil
### GetResponseTimeMs

`func (o *WebhookDeliveryDetailResponse) GetResponseTimeMs() float32`

GetResponseTimeMs returns the ResponseTimeMs field if non-nil, zero value otherwise.

### GetResponseTimeMsOk

`func (o *WebhookDeliveryDetailResponse) GetResponseTimeMsOk() (*float32, bool)`

GetResponseTimeMsOk returns a tuple with the ResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMs

`func (o *WebhookDeliveryDetailResponse) SetResponseTimeMs(v float32)`

SetResponseTimeMs sets ResponseTimeMs field to given value.


### SetResponseTimeMsNil

`func (o *WebhookDeliveryDetailResponse) SetResponseTimeMsNil(b bool)`

 SetResponseTimeMsNil sets the value for ResponseTimeMs to be an explicit nil

### UnsetResponseTimeMs
`func (o *WebhookDeliveryDetailResponse) UnsetResponseTimeMs()`

UnsetResponseTimeMs ensures that no value is present for ResponseTimeMs, not even an explicit nil
### GetErrorMessage

`func (o *WebhookDeliveryDetailResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WebhookDeliveryDetailResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WebhookDeliveryDetailResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *WebhookDeliveryDetailResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *WebhookDeliveryDetailResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetAttemptCount

`func (o *WebhookDeliveryDetailResponse) GetAttemptCount() int32`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *WebhookDeliveryDetailResponse) GetAttemptCountOk() (*int32, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *WebhookDeliveryDetailResponse) SetAttemptCount(v int32)`

SetAttemptCount sets AttemptCount field to given value.


### GetRetryCount

`func (o *WebhookDeliveryDetailResponse) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WebhookDeliveryDetailResponse) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WebhookDeliveryDetailResponse) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WebhookDeliveryDetailResponse) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookDeliveryDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookDeliveryDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookDeliveryDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeliveredAt

`func (o *WebhookDeliveryDetailResponse) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *WebhookDeliveryDetailResponse) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *WebhookDeliveryDetailResponse) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.


### SetDeliveredAtNil

`func (o *WebhookDeliveryDetailResponse) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *WebhookDeliveryDetailResponse) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetNextRetryAt

`func (o *WebhookDeliveryDetailResponse) GetNextRetryAt() time.Time`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *WebhookDeliveryDetailResponse) GetNextRetryAtOk() (*time.Time, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *WebhookDeliveryDetailResponse) SetNextRetryAt(v time.Time)`

SetNextRetryAt sets NextRetryAt field to given value.


### SetNextRetryAtNil

`func (o *WebhookDeliveryDetailResponse) SetNextRetryAtNil(b bool)`

 SetNextRetryAtNil sets the value for NextRetryAt to be an explicit nil

### UnsetNextRetryAt
`func (o *WebhookDeliveryDetailResponse) UnsetNextRetryAt()`

UnsetNextRetryAt ensures that no value is present for NextRetryAt, not even an explicit nil
### GetPayload

`func (o *WebhookDeliveryDetailResponse) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookDeliveryDetailResponse) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookDeliveryDetailResponse) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### SetPayloadNil

`func (o *WebhookDeliveryDetailResponse) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *WebhookDeliveryDetailResponse) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetResponseBody

`func (o *WebhookDeliveryDetailResponse) GetResponseBody() map[string]interface{}`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *WebhookDeliveryDetailResponse) GetResponseBodyOk() (*map[string]interface{}, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *WebhookDeliveryDetailResponse) SetResponseBody(v map[string]interface{})`

SetResponseBody sets ResponseBody field to given value.


### SetResponseBodyNil

`func (o *WebhookDeliveryDetailResponse) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *WebhookDeliveryDetailResponse) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseHeaders

`func (o *WebhookDeliveryDetailResponse) GetResponseHeaders() map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *WebhookDeliveryDetailResponse) GetResponseHeadersOk() (*map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *WebhookDeliveryDetailResponse) SetResponseHeaders(v map[string]interface{})`

SetResponseHeaders sets ResponseHeaders field to given value.


### SetResponseHeadersNil

`func (o *WebhookDeliveryDetailResponse) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *WebhookDeliveryDetailResponse) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


