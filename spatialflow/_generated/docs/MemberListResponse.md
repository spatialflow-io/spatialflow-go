# MemberListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]MemberOut**](MemberOut.md) |  | 
**Total** | **int32** |  | 
**Truncated** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMemberListResponse

`func NewMemberListResponse(members []MemberOut, total int32, ) *MemberListResponse`

NewMemberListResponse instantiates a new MemberListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberListResponseWithDefaults

`func NewMemberListResponseWithDefaults() *MemberListResponse`

NewMemberListResponseWithDefaults instantiates a new MemberListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *MemberListResponse) GetMembers() []MemberOut`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MemberListResponse) GetMembersOk() (*[]MemberOut, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MemberListResponse) SetMembers(v []MemberOut)`

SetMembers sets Members field to given value.


### GetTotal

`func (o *MemberListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MemberListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MemberListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetTruncated

`func (o *MemberListResponse) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *MemberListResponse) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *MemberListResponse) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *MemberListResponse) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


