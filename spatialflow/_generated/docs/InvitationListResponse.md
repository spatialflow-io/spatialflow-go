# InvitationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | [**[]InvitationOut**](InvitationOut.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewInvitationListResponse

`func NewInvitationListResponse(invitations []InvitationOut, total int32, ) *InvitationListResponse`

NewInvitationListResponse instantiates a new InvitationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationListResponseWithDefaults

`func NewInvitationListResponseWithDefaults() *InvitationListResponse`

NewInvitationListResponseWithDefaults instantiates a new InvitationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *InvitationListResponse) GetInvitations() []InvitationOut`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *InvitationListResponse) GetInvitationsOk() (*[]InvitationOut, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *InvitationListResponse) SetInvitations(v []InvitationOut)`

SetInvitations sets Invitations field to given value.


### GetTotal

`func (o *InvitationListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvitationListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvitationListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


