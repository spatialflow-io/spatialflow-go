# UnsubscribeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewUnsubscribeResponse

`func NewUnsubscribeResponse(success bool, message string, ) *UnsubscribeResponse`

NewUnsubscribeResponse instantiates a new UnsubscribeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsubscribeResponseWithDefaults

`func NewUnsubscribeResponseWithDefaults() *UnsubscribeResponse`

NewUnsubscribeResponseWithDefaults instantiates a new UnsubscribeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UnsubscribeResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UnsubscribeResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UnsubscribeResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *UnsubscribeResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnsubscribeResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnsubscribeResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetEmail

`func (o *UnsubscribeResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnsubscribeResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnsubscribeResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnsubscribeResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


