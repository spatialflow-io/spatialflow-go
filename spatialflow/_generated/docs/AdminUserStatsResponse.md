# AdminUserStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserApiStats**](UserApiStats.md) |  | 
**Total** | **int32** |  | 
**Days** | **int32** |  | 
**Offset** | **int32** |  | 
**Limit** | **int32** |  | 
**HasMore** | **bool** |  | 

## Methods

### NewAdminUserStatsResponse

`func NewAdminUserStatsResponse(users []UserApiStats, total int32, days int32, offset int32, limit int32, hasMore bool, ) *AdminUserStatsResponse`

NewAdminUserStatsResponse instantiates a new AdminUserStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserStatsResponseWithDefaults

`func NewAdminUserStatsResponseWithDefaults() *AdminUserStatsResponse`

NewAdminUserStatsResponseWithDefaults instantiates a new AdminUserStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *AdminUserStatsResponse) GetUsers() []UserApiStats`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AdminUserStatsResponse) GetUsersOk() (*[]UserApiStats, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AdminUserStatsResponse) SetUsers(v []UserApiStats)`

SetUsers sets Users field to given value.


### GetTotal

`func (o *AdminUserStatsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AdminUserStatsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AdminUserStatsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetDays

`func (o *AdminUserStatsResponse) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *AdminUserStatsResponse) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *AdminUserStatsResponse) SetDays(v int32)`

SetDays sets Days field to given value.


### GetOffset

`func (o *AdminUserStatsResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AdminUserStatsResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AdminUserStatsResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *AdminUserStatsResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AdminUserStatsResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AdminUserStatsResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetHasMore

`func (o *AdminUserStatsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AdminUserStatsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AdminUserStatsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


