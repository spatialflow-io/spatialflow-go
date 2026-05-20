# EmailHistoryItemOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ToEmail** | **string** |  | 
**Subject** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | **string** |  | 
**DeliveredAt** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailHistoryItemOut

`func NewEmailHistoryItemOut(id string, toEmail string, subject string, status string, createdAt string, ) *EmailHistoryItemOut`

NewEmailHistoryItemOut instantiates a new EmailHistoryItemOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailHistoryItemOutWithDefaults

`func NewEmailHistoryItemOutWithDefaults() *EmailHistoryItemOut`

NewEmailHistoryItemOutWithDefaults instantiates a new EmailHistoryItemOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailHistoryItemOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailHistoryItemOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailHistoryItemOut) SetId(v string)`

SetId sets Id field to given value.


### GetToEmail

`func (o *EmailHistoryItemOut) GetToEmail() string`

GetToEmail returns the ToEmail field if non-nil, zero value otherwise.

### GetToEmailOk

`func (o *EmailHistoryItemOut) GetToEmailOk() (*string, bool)`

GetToEmailOk returns a tuple with the ToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmail

`func (o *EmailHistoryItemOut) SetToEmail(v string)`

SetToEmail sets ToEmail field to given value.


### GetSubject

`func (o *EmailHistoryItemOut) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailHistoryItemOut) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailHistoryItemOut) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetStatus

`func (o *EmailHistoryItemOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailHistoryItemOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailHistoryItemOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *EmailHistoryItemOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailHistoryItemOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailHistoryItemOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeliveredAt

`func (o *EmailHistoryItemOut) GetDeliveredAt() string`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *EmailHistoryItemOut) GetDeliveredAtOk() (*string, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *EmailHistoryItemOut) SetDeliveredAt(v string)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *EmailHistoryItemOut) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *EmailHistoryItemOut) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *EmailHistoryItemOut) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetTemplate

`func (o *EmailHistoryItemOut) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailHistoryItemOut) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailHistoryItemOut) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailHistoryItemOut) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *EmailHistoryItemOut) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *EmailHistoryItemOut) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


