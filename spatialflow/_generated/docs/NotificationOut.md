# NotificationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Message** | **string** |  | 
**Type** | **string** |  | 
**IsRead** | **bool** |  | 
**ReadAt** | Pointer to **NullableTime** |  | [optional] 
**ActionUrl** | Pointer to **NullableString** |  | [optional] 
**ActionLabel** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewNotificationOut

`func NewNotificationOut(id string, title string, message string, type_ string, isRead bool, createdAt time.Time, ) *NotificationOut`

NewNotificationOut instantiates a new NotificationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationOutWithDefaults

`func NewNotificationOutWithDefaults() *NotificationOut`

NewNotificationOutWithDefaults instantiates a new NotificationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationOut) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *NotificationOut) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationOut) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationOut) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *NotificationOut) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationOut) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationOut) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *NotificationOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationOut) SetType(v string)`

SetType sets Type field to given value.


### GetIsRead

`func (o *NotificationOut) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *NotificationOut) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *NotificationOut) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.


### GetReadAt

`func (o *NotificationOut) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *NotificationOut) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *NotificationOut) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *NotificationOut) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.

### SetReadAtNil

`func (o *NotificationOut) SetReadAtNil(b bool)`

 SetReadAtNil sets the value for ReadAt to be an explicit nil

### UnsetReadAt
`func (o *NotificationOut) UnsetReadAt()`

UnsetReadAt ensures that no value is present for ReadAt, not even an explicit nil
### GetActionUrl

`func (o *NotificationOut) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *NotificationOut) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *NotificationOut) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *NotificationOut) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### SetActionUrlNil

`func (o *NotificationOut) SetActionUrlNil(b bool)`

 SetActionUrlNil sets the value for ActionUrl to be an explicit nil

### UnsetActionUrl
`func (o *NotificationOut) UnsetActionUrl()`

UnsetActionUrl ensures that no value is present for ActionUrl, not even an explicit nil
### GetActionLabel

`func (o *NotificationOut) GetActionLabel() string`

GetActionLabel returns the ActionLabel field if non-nil, zero value otherwise.

### GetActionLabelOk

`func (o *NotificationOut) GetActionLabelOk() (*string, bool)`

GetActionLabelOk returns a tuple with the ActionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionLabel

`func (o *NotificationOut) SetActionLabel(v string)`

SetActionLabel sets ActionLabel field to given value.

### HasActionLabel

`func (o *NotificationOut) HasActionLabel() bool`

HasActionLabel returns a boolean if a field has been set.

### SetActionLabelNil

`func (o *NotificationOut) SetActionLabelNil(b bool)`

 SetActionLabelNil sets the value for ActionLabel to be an explicit nil

### UnsetActionLabel
`func (o *NotificationOut) UnsetActionLabel()`

UnsetActionLabel ensures that no value is present for ActionLabel, not even an explicit nil
### GetCreatedAt

`func (o *NotificationOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


