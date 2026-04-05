# SubscriptionActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Subscription** | Pointer to [**NullableSubscriptionResponse**](SubscriptionResponse.md) |  | [optional] 

## Methods

### NewSubscriptionActionResponse

`func NewSubscriptionActionResponse(success bool, message string, ) *SubscriptionActionResponse`

NewSubscriptionActionResponse instantiates a new SubscriptionActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionActionResponseWithDefaults

`func NewSubscriptionActionResponseWithDefaults() *SubscriptionActionResponse`

NewSubscriptionActionResponseWithDefaults instantiates a new SubscriptionActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SubscriptionActionResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SubscriptionActionResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SubscriptionActionResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *SubscriptionActionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubscriptionActionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubscriptionActionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSubscription

`func (o *SubscriptionActionResponse) GetSubscription() SubscriptionResponse`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SubscriptionActionResponse) GetSubscriptionOk() (*SubscriptionResponse, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SubscriptionActionResponse) SetSubscription(v SubscriptionResponse)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *SubscriptionActionResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### SetSubscriptionNil

`func (o *SubscriptionActionResponse) SetSubscriptionNil(b bool)`

 SetSubscriptionNil sets the value for Subscription to be an explicit nil

### UnsetSubscription
`func (o *SubscriptionActionResponse) UnsetSubscription()`

UnsetSubscription ensures that no value is present for Subscription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


