# SsoExchangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**CodeVerifier** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewSsoExchangeRequest

`func NewSsoExchangeRequest(code string, codeVerifier string, state string, ) *SsoExchangeRequest`

NewSsoExchangeRequest instantiates a new SsoExchangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoExchangeRequestWithDefaults

`func NewSsoExchangeRequestWithDefaults() *SsoExchangeRequest`

NewSsoExchangeRequestWithDefaults instantiates a new SsoExchangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SsoExchangeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SsoExchangeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SsoExchangeRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeVerifier

`func (o *SsoExchangeRequest) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *SsoExchangeRequest) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *SsoExchangeRequest) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.


### GetState

`func (o *SsoExchangeRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SsoExchangeRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SsoExchangeRequest) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


