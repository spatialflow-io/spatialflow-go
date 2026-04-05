# CheckoutSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceId** | **string** | Stripe price ID | 
**SuccessUrl** | **string** | URL to redirect after successful payment | 
**CancelUrl** | **string** | URL to redirect if payment is cancelled | 

## Methods

### NewCheckoutSessionRequest

`func NewCheckoutSessionRequest(priceId string, successUrl string, cancelUrl string, ) *CheckoutSessionRequest`

NewCheckoutSessionRequest instantiates a new CheckoutSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionRequestWithDefaults

`func NewCheckoutSessionRequestWithDefaults() *CheckoutSessionRequest`

NewCheckoutSessionRequestWithDefaults instantiates a new CheckoutSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceId

`func (o *CheckoutSessionRequest) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *CheckoutSessionRequest) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *CheckoutSessionRequest) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetSuccessUrl

`func (o *CheckoutSessionRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CheckoutSessionRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CheckoutSessionRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### GetCancelUrl

`func (o *CheckoutSessionRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CheckoutSessionRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CheckoutSessionRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


