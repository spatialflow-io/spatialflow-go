# WebhookDeliveryResponse

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

## Methods

### NewWebhookDeliveryResponse

`func NewWebhookDeliveryResponse(id string, webhookId NullableString, geofenceId NullableString, eventType string, eventId string, url string, method string, status DeliveryStatusEnum, responseStatusCode NullableInt32, responseTimeMs NullableFloat32, errorMessage NullableString, attemptCount int32, createdAt time.Time, deliveredAt NullableTime, nextRetryAt NullableTime, ) *WebhookDeliveryResponse`

NewWebhookDeliveryResponse instantiates a new WebhookDeliveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryResponseWithDefaults

`func NewWebhookDeliveryResponseWithDefaults() *WebhookDeliveryResponse`

NewWebhookDeliveryResponseWithDefaults instantiates a new WebhookDeliveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookDeliveryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDeliveryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDeliveryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetWebhookId

`func (o *WebhookDeliveryResponse) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookDeliveryResponse) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookDeliveryResponse) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### SetWebhookIdNil

`func (o *WebhookDeliveryResponse) SetWebhookIdNil(b bool)`

 SetWebhookIdNil sets the value for WebhookId to be an explicit nil

### UnsetWebhookId
`func (o *WebhookDeliveryResponse) UnsetWebhookId()`

UnsetWebhookId ensures that no value is present for WebhookId, not even an explicit nil
### GetGeofenceId

`func (o *WebhookDeliveryResponse) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *WebhookDeliveryResponse) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *WebhookDeliveryResponse) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.


### SetGeofenceIdNil

`func (o *WebhookDeliveryResponse) SetGeofenceIdNil(b bool)`

 SetGeofenceIdNil sets the value for GeofenceId to be an explicit nil

### UnsetGeofenceId
`func (o *WebhookDeliveryResponse) UnsetGeofenceId()`

UnsetGeofenceId ensures that no value is present for GeofenceId, not even an explicit nil
### GetEventType

`func (o *WebhookDeliveryResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookDeliveryResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookDeliveryResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventId

`func (o *WebhookDeliveryResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookDeliveryResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookDeliveryResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetUrl

`func (o *WebhookDeliveryResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookDeliveryResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookDeliveryResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *WebhookDeliveryResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebhookDeliveryResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebhookDeliveryResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetStatus

`func (o *WebhookDeliveryResponse) GetStatus() DeliveryStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookDeliveryResponse) GetStatusOk() (*DeliveryStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookDeliveryResponse) SetStatus(v DeliveryStatusEnum)`

SetStatus sets Status field to given value.


### GetResponseStatusCode

`func (o *WebhookDeliveryResponse) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *WebhookDeliveryResponse) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *WebhookDeliveryResponse) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.


### SetResponseStatusCodeNil

`func (o *WebhookDeliveryResponse) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *WebhookDeliveryResponse) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetResponseStatus

`func (o *WebhookDeliveryResponse) GetResponseStatus() int32`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *WebhookDeliveryResponse) GetResponseStatusOk() (*int32, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *WebhookDeliveryResponse) SetResponseStatus(v int32)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *WebhookDeliveryResponse) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### SetResponseStatusNil

`func (o *WebhookDeliveryResponse) SetResponseStatusNil(b bool)`

 SetResponseStatusNil sets the value for ResponseStatus to be an explicit nil

### UnsetResponseStatus
`func (o *WebhookDeliveryResponse) UnsetResponseStatus()`

UnsetResponseStatus ensures that no value is present for ResponseStatus, not even an explicit nil
### GetResponseTimeMs

`func (o *WebhookDeliveryResponse) GetResponseTimeMs() float32`

GetResponseTimeMs returns the ResponseTimeMs field if non-nil, zero value otherwise.

### GetResponseTimeMsOk

`func (o *WebhookDeliveryResponse) GetResponseTimeMsOk() (*float32, bool)`

GetResponseTimeMsOk returns a tuple with the ResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMs

`func (o *WebhookDeliveryResponse) SetResponseTimeMs(v float32)`

SetResponseTimeMs sets ResponseTimeMs field to given value.


### SetResponseTimeMsNil

`func (o *WebhookDeliveryResponse) SetResponseTimeMsNil(b bool)`

 SetResponseTimeMsNil sets the value for ResponseTimeMs to be an explicit nil

### UnsetResponseTimeMs
`func (o *WebhookDeliveryResponse) UnsetResponseTimeMs()`

UnsetResponseTimeMs ensures that no value is present for ResponseTimeMs, not even an explicit nil
### GetErrorMessage

`func (o *WebhookDeliveryResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WebhookDeliveryResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WebhookDeliveryResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *WebhookDeliveryResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *WebhookDeliveryResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetAttemptCount

`func (o *WebhookDeliveryResponse) GetAttemptCount() int32`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *WebhookDeliveryResponse) GetAttemptCountOk() (*int32, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *WebhookDeliveryResponse) SetAttemptCount(v int32)`

SetAttemptCount sets AttemptCount field to given value.


### GetRetryCount

`func (o *WebhookDeliveryResponse) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WebhookDeliveryResponse) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WebhookDeliveryResponse) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WebhookDeliveryResponse) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookDeliveryResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookDeliveryResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookDeliveryResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeliveredAt

`func (o *WebhookDeliveryResponse) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *WebhookDeliveryResponse) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *WebhookDeliveryResponse) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.


### SetDeliveredAtNil

`func (o *WebhookDeliveryResponse) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *WebhookDeliveryResponse) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetNextRetryAt

`func (o *WebhookDeliveryResponse) GetNextRetryAt() time.Time`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *WebhookDeliveryResponse) GetNextRetryAtOk() (*time.Time, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *WebhookDeliveryResponse) SetNextRetryAt(v time.Time)`

SetNextRetryAt sets NextRetryAt field to given value.


### SetNextRetryAtNil

`func (o *WebhookDeliveryResponse) SetNextRetryAtNil(b bool)`

 SetNextRetryAtNil sets the value for NextRetryAt to be an explicit nil

### UnsetNextRetryAt
`func (o *WebhookDeliveryResponse) UnsetNextRetryAt()`

UnsetNextRetryAt ensures that no value is present for NextRetryAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


