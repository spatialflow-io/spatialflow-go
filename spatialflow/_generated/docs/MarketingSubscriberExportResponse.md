# MarketingSubscriberExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscribers** | [**[]MarketingSubscriber**](MarketingSubscriber.md) |  | 
**Count** | **int32** |  | 
**ExportedAt** | **string** |  | 

## Methods

### NewMarketingSubscriberExportResponse

`func NewMarketingSubscriberExportResponse(subscribers []MarketingSubscriber, count int32, exportedAt string, ) *MarketingSubscriberExportResponse`

NewMarketingSubscriberExportResponse instantiates a new MarketingSubscriberExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSubscriberExportResponseWithDefaults

`func NewMarketingSubscriberExportResponseWithDefaults() *MarketingSubscriberExportResponse`

NewMarketingSubscriberExportResponseWithDefaults instantiates a new MarketingSubscriberExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribers

`func (o *MarketingSubscriberExportResponse) GetSubscribers() []MarketingSubscriber`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *MarketingSubscriberExportResponse) GetSubscribersOk() (*[]MarketingSubscriber, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *MarketingSubscriberExportResponse) SetSubscribers(v []MarketingSubscriber)`

SetSubscribers sets Subscribers field to given value.


### GetCount

`func (o *MarketingSubscriberExportResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MarketingSubscriberExportResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MarketingSubscriberExportResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetExportedAt

`func (o *MarketingSubscriberExportResponse) GetExportedAt() string`

GetExportedAt returns the ExportedAt field if non-nil, zero value otherwise.

### GetExportedAtOk

`func (o *MarketingSubscriberExportResponse) GetExportedAtOk() (*string, bool)`

GetExportedAtOk returns a tuple with the ExportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedAt

`func (o *MarketingSubscriberExportResponse) SetExportedAt(v string)`

SetExportedAt sets ExportedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


