# CreateUserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] [default to "Test"]
**LastName** | Pointer to **string** |  | [optional] [default to "User"]
**EmailVerified** | Pointer to **bool** |  | [optional] [default to true]
**IsStaff** | Pointer to **bool** |  | [optional] [default to false]
**IsSuperuser** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCreateUserSchema

`func NewCreateUserSchema(email string, password string, ) *CreateUserSchema`

NewCreateUserSchema instantiates a new CreateUserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserSchemaWithDefaults

`func NewCreateUserSchemaWithDefaults() *CreateUserSchema`

NewCreateUserSchemaWithDefaults instantiates a new CreateUserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateUserSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserSchema) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CreateUserSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserSchema) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetFirstName

`func (o *CreateUserSchema) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateUserSchema) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateUserSchema) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CreateUserSchema) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CreateUserSchema) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateUserSchema) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateUserSchema) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CreateUserSchema) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmailVerified

`func (o *CreateUserSchema) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *CreateUserSchema) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *CreateUserSchema) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *CreateUserSchema) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetIsStaff

`func (o *CreateUserSchema) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *CreateUserSchema) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *CreateUserSchema) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *CreateUserSchema) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsSuperuser

`func (o *CreateUserSchema) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *CreateUserSchema) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *CreateUserSchema) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *CreateUserSchema) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


