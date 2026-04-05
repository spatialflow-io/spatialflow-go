# PlanChangePreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPlan** | [**PlanResponse**](PlanResponse.md) |  | 
**NewPlan** | [**PlanResponse**](PlanResponse.md) |  | 
**ProratedAmount** | **int32** |  | 
**ImmediateCharge** | **int32** |  | 
**NextInvoiceChange** | **int32** |  | 
**EffectiveDate** | **string** |  | 

## Methods

### NewPlanChangePreviewResponse

`func NewPlanChangePreviewResponse(currentPlan PlanResponse, newPlan PlanResponse, proratedAmount int32, immediateCharge int32, nextInvoiceChange int32, effectiveDate string, ) *PlanChangePreviewResponse`

NewPlanChangePreviewResponse instantiates a new PlanChangePreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanChangePreviewResponseWithDefaults

`func NewPlanChangePreviewResponseWithDefaults() *PlanChangePreviewResponse`

NewPlanChangePreviewResponseWithDefaults instantiates a new PlanChangePreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPlan

`func (o *PlanChangePreviewResponse) GetCurrentPlan() PlanResponse`

GetCurrentPlan returns the CurrentPlan field if non-nil, zero value otherwise.

### GetCurrentPlanOk

`func (o *PlanChangePreviewResponse) GetCurrentPlanOk() (*PlanResponse, bool)`

GetCurrentPlanOk returns a tuple with the CurrentPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlan

`func (o *PlanChangePreviewResponse) SetCurrentPlan(v PlanResponse)`

SetCurrentPlan sets CurrentPlan field to given value.


### GetNewPlan

`func (o *PlanChangePreviewResponse) GetNewPlan() PlanResponse`

GetNewPlan returns the NewPlan field if non-nil, zero value otherwise.

### GetNewPlanOk

`func (o *PlanChangePreviewResponse) GetNewPlanOk() (*PlanResponse, bool)`

GetNewPlanOk returns a tuple with the NewPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlan

`func (o *PlanChangePreviewResponse) SetNewPlan(v PlanResponse)`

SetNewPlan sets NewPlan field to given value.


### GetProratedAmount

`func (o *PlanChangePreviewResponse) GetProratedAmount() int32`

GetProratedAmount returns the ProratedAmount field if non-nil, zero value otherwise.

### GetProratedAmountOk

`func (o *PlanChangePreviewResponse) GetProratedAmountOk() (*int32, bool)`

GetProratedAmountOk returns a tuple with the ProratedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProratedAmount

`func (o *PlanChangePreviewResponse) SetProratedAmount(v int32)`

SetProratedAmount sets ProratedAmount field to given value.


### GetImmediateCharge

`func (o *PlanChangePreviewResponse) GetImmediateCharge() int32`

GetImmediateCharge returns the ImmediateCharge field if non-nil, zero value otherwise.

### GetImmediateChargeOk

`func (o *PlanChangePreviewResponse) GetImmediateChargeOk() (*int32, bool)`

GetImmediateChargeOk returns a tuple with the ImmediateCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateCharge

`func (o *PlanChangePreviewResponse) SetImmediateCharge(v int32)`

SetImmediateCharge sets ImmediateCharge field to given value.


### GetNextInvoiceChange

`func (o *PlanChangePreviewResponse) GetNextInvoiceChange() int32`

GetNextInvoiceChange returns the NextInvoiceChange field if non-nil, zero value otherwise.

### GetNextInvoiceChangeOk

`func (o *PlanChangePreviewResponse) GetNextInvoiceChangeOk() (*int32, bool)`

GetNextInvoiceChangeOk returns a tuple with the NextInvoiceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceChange

`func (o *PlanChangePreviewResponse) SetNextInvoiceChange(v int32)`

SetNextInvoiceChange sets NextInvoiceChange field to given value.


### GetEffectiveDate

`func (o *PlanChangePreviewResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *PlanChangePreviewResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *PlanChangePreviewResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


