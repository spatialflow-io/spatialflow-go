# NotificationTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Provider** | **string** |  | 
**Message** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**ResponseTimeMs** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewNotificationTestResponse

`func NewNotificationTestResponse(success bool, provider string, ) *NotificationTestResponse`

NewNotificationTestResponse instantiates a new NotificationTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTestResponseWithDefaults

`func NewNotificationTestResponseWithDefaults() *NotificationTestResponse`

NewNotificationTestResponseWithDefaults instantiates a new NotificationTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *NotificationTestResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *NotificationTestResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *NotificationTestResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetProvider

`func (o *NotificationTestResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationTestResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationTestResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetMessage

`func (o *NotificationTestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationTestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationTestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotificationTestResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NotificationTestResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NotificationTestResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetError

`func (o *NotificationTestResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NotificationTestResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NotificationTestResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *NotificationTestResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *NotificationTestResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *NotificationTestResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetResponseTimeMs

`func (o *NotificationTestResponse) GetResponseTimeMs() float32`

GetResponseTimeMs returns the ResponseTimeMs field if non-nil, zero value otherwise.

### GetResponseTimeMsOk

`func (o *NotificationTestResponse) GetResponseTimeMsOk() (*float32, bool)`

GetResponseTimeMsOk returns a tuple with the ResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMs

`func (o *NotificationTestResponse) SetResponseTimeMs(v float32)`

SetResponseTimeMs sets ResponseTimeMs field to given value.

### HasResponseTimeMs

`func (o *NotificationTestResponse) HasResponseTimeMs() bool`

HasResponseTimeMs returns a boolean if a field has been set.

### SetResponseTimeMsNil

`func (o *NotificationTestResponse) SetResponseTimeMsNil(b bool)`

 SetResponseTimeMsNil sets the value for ResponseTimeMs to be an explicit nil

### UnsetResponseTimeMs
`func (o *NotificationTestResponse) UnsetResponseTimeMs()`

UnsetResponseTimeMs ensures that no value is present for ResponseTimeMs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


