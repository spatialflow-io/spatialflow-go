# WebhookTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**TestId** | **string** |  | 
**PayloadSent** | **map[string]interface{}** |  | 
**Message** | **string** |  | 
**ResponseStatusCode** | Pointer to **NullableInt32** |  | [optional] 
**ResponseTimeMs** | Pointer to **NullableFloat32** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebhookTestResponse

`func NewWebhookTestResponse(success bool, testId string, payloadSent map[string]interface{}, message string, ) *WebhookTestResponse`

NewWebhookTestResponse instantiates a new WebhookTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookTestResponseWithDefaults

`func NewWebhookTestResponseWithDefaults() *WebhookTestResponse`

NewWebhookTestResponseWithDefaults instantiates a new WebhookTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *WebhookTestResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WebhookTestResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WebhookTestResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTestId

`func (o *WebhookTestResponse) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *WebhookTestResponse) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *WebhookTestResponse) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetPayloadSent

`func (o *WebhookTestResponse) GetPayloadSent() map[string]interface{}`

GetPayloadSent returns the PayloadSent field if non-nil, zero value otherwise.

### GetPayloadSentOk

`func (o *WebhookTestResponse) GetPayloadSentOk() (*map[string]interface{}, bool)`

GetPayloadSentOk returns a tuple with the PayloadSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadSent

`func (o *WebhookTestResponse) SetPayloadSent(v map[string]interface{})`

SetPayloadSent sets PayloadSent field to given value.


### GetMessage

`func (o *WebhookTestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookTestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookTestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResponseStatusCode

`func (o *WebhookTestResponse) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *WebhookTestResponse) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *WebhookTestResponse) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *WebhookTestResponse) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCodeNil

`func (o *WebhookTestResponse) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *WebhookTestResponse) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetResponseTimeMs

`func (o *WebhookTestResponse) GetResponseTimeMs() float32`

GetResponseTimeMs returns the ResponseTimeMs field if non-nil, zero value otherwise.

### GetResponseTimeMsOk

`func (o *WebhookTestResponse) GetResponseTimeMsOk() (*float32, bool)`

GetResponseTimeMsOk returns a tuple with the ResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMs

`func (o *WebhookTestResponse) SetResponseTimeMs(v float32)`

SetResponseTimeMs sets ResponseTimeMs field to given value.

### HasResponseTimeMs

`func (o *WebhookTestResponse) HasResponseTimeMs() bool`

HasResponseTimeMs returns a boolean if a field has been set.

### SetResponseTimeMsNil

`func (o *WebhookTestResponse) SetResponseTimeMsNil(b bool)`

 SetResponseTimeMsNil sets the value for ResponseTimeMs to be an explicit nil

### UnsetResponseTimeMs
`func (o *WebhookTestResponse) UnsetResponseTimeMs()`

UnsetResponseTimeMs ensures that no value is present for ResponseTimeMs, not even an explicit nil
### GetError

`func (o *WebhookTestResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookTestResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookTestResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookTestResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *WebhookTestResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *WebhookTestResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


