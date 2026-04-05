# CheckoutSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | Stripe checkout session ID | 
**Url** | **string** | Checkout URL to redirect user to | 

## Methods

### NewCheckoutSessionResponse

`func NewCheckoutSessionResponse(sessionId string, url string, ) *CheckoutSessionResponse`

NewCheckoutSessionResponse instantiates a new CheckoutSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionResponseWithDefaults

`func NewCheckoutSessionResponseWithDefaults() *CheckoutSessionResponse`

NewCheckoutSessionResponseWithDefaults instantiates a new CheckoutSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *CheckoutSessionResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CheckoutSessionResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CheckoutSessionResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetUrl

`func (o *CheckoutSessionResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutSessionResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutSessionResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


