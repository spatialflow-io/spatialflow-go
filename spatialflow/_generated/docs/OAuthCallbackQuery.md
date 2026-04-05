# OAuthCallbackQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**State** | **string** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**ErrorDescription** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOAuthCallbackQuery

`func NewOAuthCallbackQuery(code string, state string, ) *OAuthCallbackQuery`

NewOAuthCallbackQuery instantiates a new OAuthCallbackQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthCallbackQueryWithDefaults

`func NewOAuthCallbackQueryWithDefaults() *OAuthCallbackQuery`

NewOAuthCallbackQueryWithDefaults instantiates a new OAuthCallbackQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OAuthCallbackQuery) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OAuthCallbackQuery) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OAuthCallbackQuery) SetCode(v string)`

SetCode sets Code field to given value.


### GetState

`func (o *OAuthCallbackQuery) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OAuthCallbackQuery) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OAuthCallbackQuery) SetState(v string)`

SetState sets State field to given value.


### GetError

`func (o *OAuthCallbackQuery) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OAuthCallbackQuery) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OAuthCallbackQuery) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OAuthCallbackQuery) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *OAuthCallbackQuery) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *OAuthCallbackQuery) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetErrorDescription

`func (o *OAuthCallbackQuery) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *OAuthCallbackQuery) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *OAuthCallbackQuery) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *OAuthCallbackQuery) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *OAuthCallbackQuery) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *OAuthCallbackQuery) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


