# NotificationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | [**[]NotificationOut**](NotificationOut.md) |  | 
**Count** | **int32** |  | 

## Methods

### NewNotificationListResponse

`func NewNotificationListResponse(notifications []NotificationOut, count int32, ) *NotificationListResponse`

NewNotificationListResponse instantiates a new NotificationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationListResponseWithDefaults

`func NewNotificationListResponseWithDefaults() *NotificationListResponse`

NewNotificationListResponseWithDefaults instantiates a new NotificationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *NotificationListResponse) GetNotifications() []NotificationOut`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationListResponse) GetNotificationsOk() (*[]NotificationOut, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationListResponse) SetNotifications(v []NotificationOut)`

SetNotifications sets Notifications field to given value.


### GetCount

`func (o *NotificationListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NotificationListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NotificationListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


