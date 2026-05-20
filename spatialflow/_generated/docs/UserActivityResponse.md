# UserActivityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**Email** | **string** |  | 
**Activities** | [**[]UserActivityItem**](UserActivityItem.md) |  | 
**Total** | **int32** |  | 
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 

## Methods

### NewUserActivityResponse

`func NewUserActivityResponse(userId string, email string, activities []UserActivityItem, total int32, limit int32, offset int32, ) *UserActivityResponse`

NewUserActivityResponse instantiates a new UserActivityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityResponseWithDefaults

`func NewUserActivityResponseWithDefaults() *UserActivityResponse`

NewUserActivityResponseWithDefaults instantiates a new UserActivityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserActivityResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserActivityResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserActivityResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *UserActivityResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserActivityResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserActivityResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetActivities

`func (o *UserActivityResponse) GetActivities() []UserActivityItem`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *UserActivityResponse) GetActivitiesOk() (*[]UserActivityItem, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *UserActivityResponse) SetActivities(v []UserActivityItem)`

SetActivities sets Activities field to given value.


### GetTotal

`func (o *UserActivityResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserActivityResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserActivityResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *UserActivityResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UserActivityResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UserActivityResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *UserActivityResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UserActivityResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UserActivityResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


