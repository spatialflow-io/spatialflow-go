# InvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Number** | **string** |  | 
**Status** | **string** |  | 
**AmountDue** | **int32** |  | 
**AmountPaid** | **int32** |  | 
**Currency** | **string** |  | 
**PeriodStart** | **int32** |  | 
**PeriodEnd** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**InvoicePdf** | Pointer to **NullableString** |  | [optional] 
**LineItems** | [**[]InvoiceLineItem**](InvoiceLineItem.md) |  | 

## Methods

### NewInvoiceResponse

`func NewInvoiceResponse(id string, number string, status string, amountDue int32, amountPaid int32, currency string, periodStart int32, periodEnd int32, createdAt int32, lineItems []InvoiceLineItem, ) *InvoiceResponse`

NewInvoiceResponse instantiates a new InvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceResponseWithDefaults

`func NewInvoiceResponseWithDefaults() *InvoiceResponse`

NewInvoiceResponseWithDefaults instantiates a new InvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *InvoiceResponse) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceResponse) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceResponse) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetStatus

`func (o *InvoiceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAmountDue

`func (o *InvoiceResponse) GetAmountDue() int32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *InvoiceResponse) GetAmountDueOk() (*int32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *InvoiceResponse) SetAmountDue(v int32)`

SetAmountDue sets AmountDue field to given value.


### GetAmountPaid

`func (o *InvoiceResponse) GetAmountPaid() int32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *InvoiceResponse) GetAmountPaidOk() (*int32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *InvoiceResponse) SetAmountPaid(v int32)`

SetAmountPaid sets AmountPaid field to given value.


### GetCurrency

`func (o *InvoiceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPeriodStart

`func (o *InvoiceResponse) GetPeriodStart() int32`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *InvoiceResponse) GetPeriodStartOk() (*int32, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *InvoiceResponse) SetPeriodStart(v int32)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *InvoiceResponse) GetPeriodEnd() int32`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *InvoiceResponse) GetPeriodEndOk() (*int32, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *InvoiceResponse) SetPeriodEnd(v int32)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetCreatedAt

`func (o *InvoiceResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetInvoicePdf

`func (o *InvoiceResponse) GetInvoicePdf() string`

GetInvoicePdf returns the InvoicePdf field if non-nil, zero value otherwise.

### GetInvoicePdfOk

`func (o *InvoiceResponse) GetInvoicePdfOk() (*string, bool)`

GetInvoicePdfOk returns a tuple with the InvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdf

`func (o *InvoiceResponse) SetInvoicePdf(v string)`

SetInvoicePdf sets InvoicePdf field to given value.

### HasInvoicePdf

`func (o *InvoiceResponse) HasInvoicePdf() bool`

HasInvoicePdf returns a boolean if a field has been set.

### SetInvoicePdfNil

`func (o *InvoiceResponse) SetInvoicePdfNil(b bool)`

 SetInvoicePdfNil sets the value for InvoicePdf to be an explicit nil

### UnsetInvoicePdf
`func (o *InvoiceResponse) UnsetInvoicePdf()`

UnsetInvoicePdf ensures that no value is present for InvoicePdf, not even an explicit nil
### GetLineItems

`func (o *InvoiceResponse) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *InvoiceResponse) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *InvoiceResponse) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


