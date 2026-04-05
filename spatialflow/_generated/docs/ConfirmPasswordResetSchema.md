# ConfirmPasswordResetSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Password** | **string** |  | 

## Methods

### NewConfirmPasswordResetSchema

`func NewConfirmPasswordResetSchema(token string, password string, ) *ConfirmPasswordResetSchema`

NewConfirmPasswordResetSchema instantiates a new ConfirmPasswordResetSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmPasswordResetSchemaWithDefaults

`func NewConfirmPasswordResetSchemaWithDefaults() *ConfirmPasswordResetSchema`

NewConfirmPasswordResetSchemaWithDefaults instantiates a new ConfirmPasswordResetSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ConfirmPasswordResetSchema) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConfirmPasswordResetSchema) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConfirmPasswordResetSchema) SetToken(v string)`

SetToken sets Token field to given value.


### GetPassword

`func (o *ConfirmPasswordResetSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConfirmPasswordResetSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConfirmPasswordResetSchema) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


