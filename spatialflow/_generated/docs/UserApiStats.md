# UserApiStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**Email** | **string** |  | 
**Name** | **NullableString** |  | 
**ApiCalls** | **int32** |  | 

## Methods

### NewUserApiStats

`func NewUserApiStats(userId string, email string, name NullableString, apiCalls int32, ) *UserApiStats`

NewUserApiStats instantiates a new UserApiStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserApiStatsWithDefaults

`func NewUserApiStatsWithDefaults() *UserApiStats`

NewUserApiStatsWithDefaults instantiates a new UserApiStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserApiStats) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserApiStats) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserApiStats) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *UserApiStats) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserApiStats) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserApiStats) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *UserApiStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserApiStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserApiStats) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UserApiStats) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserApiStats) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetApiCalls

`func (o *UserApiStats) GetApiCalls() int32`

GetApiCalls returns the ApiCalls field if non-nil, zero value otherwise.

### GetApiCallsOk

`func (o *UserApiStats) GetApiCallsOk() (*int32, bool)`

GetApiCallsOk returns a tuple with the ApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCalls

`func (o *UserApiStats) SetApiCalls(v int32)`

SetApiCalls sets ApiCalls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


