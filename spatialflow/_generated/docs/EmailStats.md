# EmailStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalEmails** | **int32** |  | 
**Delivered** | **int32** |  | 
**Failed** | **int32** |  | 

## Methods

### NewEmailStats

`func NewEmailStats(totalEmails int32, delivered int32, failed int32, ) *EmailStats`

NewEmailStats instantiates a new EmailStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStatsWithDefaults

`func NewEmailStatsWithDefaults() *EmailStats`

NewEmailStatsWithDefaults instantiates a new EmailStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalEmails

`func (o *EmailStats) GetTotalEmails() int32`

GetTotalEmails returns the TotalEmails field if non-nil, zero value otherwise.

### GetTotalEmailsOk

`func (o *EmailStats) GetTotalEmailsOk() (*int32, bool)`

GetTotalEmailsOk returns a tuple with the TotalEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEmails

`func (o *EmailStats) SetTotalEmails(v int32)`

SetTotalEmails sets TotalEmails field to given value.


### GetDelivered

`func (o *EmailStats) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *EmailStats) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *EmailStats) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.


### GetFailed

`func (o *EmailStats) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *EmailStats) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *EmailStats) SetFailed(v int32)`

SetFailed sets Failed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


