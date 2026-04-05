# UserListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserSummary**](UserSummary.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**Limit** | **int32** |  | 
**Pages** | **int32** |  | 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserListResponse

`func NewUserListResponse(users []UserSummary, total int32, page int32, limit int32, pages int32, ) *UserListResponse`

NewUserListResponse instantiates a new UserListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListResponseWithDefaults

`func NewUserListResponseWithDefaults() *UserListResponse`

NewUserListResponseWithDefaults instantiates a new UserListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UserListResponse) GetUsers() []UserSummary`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserListResponse) GetUsersOk() (*[]UserSummary, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserListResponse) SetUsers(v []UserSummary)`

SetUsers sets Users field to given value.


### GetTotal

`func (o *UserListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *UserListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UserListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UserListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *UserListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UserListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UserListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPages

`func (o *UserListResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *UserListResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *UserListResponse) SetPages(v int32)`

SetPages sets Pages field to given value.


### GetNextCursor

`func (o *UserListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *UserListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *UserListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *UserListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *UserListResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *UserListResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


