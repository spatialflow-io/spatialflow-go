# AppleTokenExchangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**RefreshToken** | **string** |  | 
**TokenType** | **string** |  | 
**ExpiresIn** | **int32** |  | 
**User** | **map[string]interface{}** |  | 
**Created** | **bool** |  | 

## Methods

### NewAppleTokenExchangeResponse

`func NewAppleTokenExchangeResponse(accessToken string, refreshToken string, tokenType string, expiresIn int32, user map[string]interface{}, created bool, ) *AppleTokenExchangeResponse`

NewAppleTokenExchangeResponse instantiates a new AppleTokenExchangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleTokenExchangeResponseWithDefaults

`func NewAppleTokenExchangeResponseWithDefaults() *AppleTokenExchangeResponse`

NewAppleTokenExchangeResponseWithDefaults instantiates a new AppleTokenExchangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AppleTokenExchangeResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AppleTokenExchangeResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AppleTokenExchangeResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *AppleTokenExchangeResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AppleTokenExchangeResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AppleTokenExchangeResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetTokenType

`func (o *AppleTokenExchangeResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AppleTokenExchangeResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AppleTokenExchangeResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *AppleTokenExchangeResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AppleTokenExchangeResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AppleTokenExchangeResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetUser

`func (o *AppleTokenExchangeResponse) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AppleTokenExchangeResponse) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AppleTokenExchangeResponse) SetUser(v map[string]interface{})`

SetUser sets User field to given value.


### GetCreated

`func (o *AppleTokenExchangeResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppleTokenExchangeResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppleTokenExchangeResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


