# ResetPasswordSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Token** | **string** |  | 
**NewPassword** | **string** |  | 

## Methods

### NewResetPasswordSchema

`func NewResetPasswordSchema(uid string, token string, newPassword string, ) *ResetPasswordSchema`

NewResetPasswordSchema instantiates a new ResetPasswordSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordSchemaWithDefaults

`func NewResetPasswordSchemaWithDefaults() *ResetPasswordSchema`

NewResetPasswordSchemaWithDefaults instantiates a new ResetPasswordSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ResetPasswordSchema) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ResetPasswordSchema) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ResetPasswordSchema) SetUid(v string)`

SetUid sets Uid field to given value.


### GetToken

`func (o *ResetPasswordSchema) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResetPasswordSchema) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResetPasswordSchema) SetToken(v string)`

SetToken sets Token field to given value.


### GetNewPassword

`func (o *ResetPasswordSchema) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetPasswordSchema) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetPasswordSchema) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


