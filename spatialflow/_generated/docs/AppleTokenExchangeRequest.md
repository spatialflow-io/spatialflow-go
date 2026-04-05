# AppleTokenExchangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityToken** | **string** |  | 
**AuthorizationCode** | **string** |  | 
**FullName** | Pointer to [**NullableAppleFullName**](AppleFullName.md) |  | [optional] 
**Nonce** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppleTokenExchangeRequest

`func NewAppleTokenExchangeRequest(identityToken string, authorizationCode string, ) *AppleTokenExchangeRequest`

NewAppleTokenExchangeRequest instantiates a new AppleTokenExchangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleTokenExchangeRequestWithDefaults

`func NewAppleTokenExchangeRequestWithDefaults() *AppleTokenExchangeRequest`

NewAppleTokenExchangeRequestWithDefaults instantiates a new AppleTokenExchangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityToken

`func (o *AppleTokenExchangeRequest) GetIdentityToken() string`

GetIdentityToken returns the IdentityToken field if non-nil, zero value otherwise.

### GetIdentityTokenOk

`func (o *AppleTokenExchangeRequest) GetIdentityTokenOk() (*string, bool)`

GetIdentityTokenOk returns a tuple with the IdentityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityToken

`func (o *AppleTokenExchangeRequest) SetIdentityToken(v string)`

SetIdentityToken sets IdentityToken field to given value.


### GetAuthorizationCode

`func (o *AppleTokenExchangeRequest) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *AppleTokenExchangeRequest) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *AppleTokenExchangeRequest) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.


### GetFullName

`func (o *AppleTokenExchangeRequest) GetFullName() AppleFullName`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AppleTokenExchangeRequest) GetFullNameOk() (*AppleFullName, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AppleTokenExchangeRequest) SetFullName(v AppleFullName)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AppleTokenExchangeRequest) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *AppleTokenExchangeRequest) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *AppleTokenExchangeRequest) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetNonce

`func (o *AppleTokenExchangeRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *AppleTokenExchangeRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *AppleTokenExchangeRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *AppleTokenExchangeRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *AppleTokenExchangeRequest) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *AppleTokenExchangeRequest) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


