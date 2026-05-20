# BulkApproveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]string** |  | 
**Notes** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewBulkApproveRequest

`func NewBulkApproveRequest(userIds []string, ) *BulkApproveRequest`

NewBulkApproveRequest instantiates a new BulkApproveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkApproveRequestWithDefaults

`func NewBulkApproveRequestWithDefaults() *BulkApproveRequest`

NewBulkApproveRequestWithDefaults instantiates a new BulkApproveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *BulkApproveRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *BulkApproveRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *BulkApproveRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### GetNotes

`func (o *BulkApproveRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BulkApproveRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BulkApproveRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BulkApproveRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


