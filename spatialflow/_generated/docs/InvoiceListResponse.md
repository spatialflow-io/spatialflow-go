# InvoiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | [**[]InvoiceResponse**](InvoiceResponse.md) |  | 
**HasMore** | **bool** |  | 

## Methods

### NewInvoiceListResponse

`func NewInvoiceListResponse(invoices []InvoiceResponse, hasMore bool, ) *InvoiceListResponse`

NewInvoiceListResponse instantiates a new InvoiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceListResponseWithDefaults

`func NewInvoiceListResponseWithDefaults() *InvoiceListResponse`

NewInvoiceListResponseWithDefaults instantiates a new InvoiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoices

`func (o *InvoiceListResponse) GetInvoices() []InvoiceResponse`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *InvoiceListResponse) GetInvoicesOk() (*[]InvoiceResponse, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *InvoiceListResponse) SetInvoices(v []InvoiceResponse)`

SetInvoices sets Invoices field to given value.


### GetHasMore

`func (o *InvoiceListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InvoiceListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InvoiceListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


