# BulkInviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invites** | [**[]BulkInviteItem**](BulkInviteItem.md) |  | 

## Methods

### NewBulkInviteRequest

`func NewBulkInviteRequest(invites []BulkInviteItem, ) *BulkInviteRequest`

NewBulkInviteRequest instantiates a new BulkInviteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInviteRequestWithDefaults

`func NewBulkInviteRequestWithDefaults() *BulkInviteRequest`

NewBulkInviteRequestWithDefaults instantiates a new BulkInviteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvites

`func (o *BulkInviteRequest) GetInvites() []BulkInviteItem`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *BulkInviteRequest) GetInvitesOk() (*[]BulkInviteItem, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *BulkInviteRequest) SetInvites(v []BulkInviteItem)`

SetInvites sets Invites field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


