# EmailTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | Pointer to **NullableString** |  | [optional] 
**Recipient** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to [**NullableEmailConfigInfo**](EmailConfigInfo.md) |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailTestResponse

`func NewEmailTestResponse(success bool, ) *EmailTestResponse`

NewEmailTestResponse instantiates a new EmailTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTestResponseWithDefaults

`func NewEmailTestResponseWithDefaults() *EmailTestResponse`

NewEmailTestResponseWithDefaults instantiates a new EmailTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *EmailTestResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *EmailTestResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *EmailTestResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *EmailTestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EmailTestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EmailTestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EmailTestResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *EmailTestResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EmailTestResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRecipient

`func (o *EmailTestResponse) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *EmailTestResponse) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *EmailTestResponse) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *EmailTestResponse) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *EmailTestResponse) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *EmailTestResponse) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetConfig

`func (o *EmailTestResponse) GetConfig() EmailConfigInfo`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EmailTestResponse) GetConfigOk() (*EmailConfigInfo, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EmailTestResponse) SetConfig(v EmailConfigInfo)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *EmailTestResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *EmailTestResponse) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *EmailTestResponse) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetError

`func (o *EmailTestResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EmailTestResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EmailTestResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EmailTestResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *EmailTestResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *EmailTestResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


