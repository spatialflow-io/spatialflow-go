# InvoiceLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Quantity** | **int32** |  | 
**UnitAmount** | **int32** |  | 
**Amount** | **int32** |  | 

## Methods

### NewInvoiceLineItem

`func NewInvoiceLineItem(description string, quantity int32, unitAmount int32, amount int32, ) *InvoiceLineItem`

NewInvoiceLineItem instantiates a new InvoiceLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineItemWithDefaults

`func NewInvoiceLineItemWithDefaults() *InvoiceLineItem`

NewInvoiceLineItemWithDefaults instantiates a new InvoiceLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InvoiceLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *InvoiceLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitAmount

`func (o *InvoiceLineItem) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *InvoiceLineItem) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *InvoiceLineItem) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.


### GetAmount

`func (o *InvoiceLineItem) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceLineItem) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceLineItem) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


