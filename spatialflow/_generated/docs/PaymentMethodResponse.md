# PaymentMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Last4** | **string** |  | 
**ExpMonth** | Pointer to **NullableInt32** |  | [optional] 
**ExpYear** | Pointer to **NullableInt32** |  | [optional] 
**IsDefault** | **bool** |  | 
**CreatedAt** | **int32** |  | 

## Methods

### NewPaymentMethodResponse

`func NewPaymentMethodResponse(id string, type_ string, last4 string, isDefault bool, createdAt int32, ) *PaymentMethodResponse`

NewPaymentMethodResponse instantiates a new PaymentMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodResponseWithDefaults

`func NewPaymentMethodResponseWithDefaults() *PaymentMethodResponse`

NewPaymentMethodResponseWithDefaults instantiates a new PaymentMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethodResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PaymentMethodResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodResponse) SetType(v string)`

SetType sets Type field to given value.


### GetBrand

`func (o *PaymentMethodResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PaymentMethodResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PaymentMethodResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *PaymentMethodResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *PaymentMethodResponse) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *PaymentMethodResponse) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetLast4

`func (o *PaymentMethodResponse) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentMethodResponse) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentMethodResponse) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetExpMonth

`func (o *PaymentMethodResponse) GetExpMonth() int32`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *PaymentMethodResponse) GetExpMonthOk() (*int32, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *PaymentMethodResponse) SetExpMonth(v int32)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *PaymentMethodResponse) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### SetExpMonthNil

`func (o *PaymentMethodResponse) SetExpMonthNil(b bool)`

 SetExpMonthNil sets the value for ExpMonth to be an explicit nil

### UnsetExpMonth
`func (o *PaymentMethodResponse) UnsetExpMonth()`

UnsetExpMonth ensures that no value is present for ExpMonth, not even an explicit nil
### GetExpYear

`func (o *PaymentMethodResponse) GetExpYear() int32`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *PaymentMethodResponse) GetExpYearOk() (*int32, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *PaymentMethodResponse) SetExpYear(v int32)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *PaymentMethodResponse) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### SetExpYearNil

`func (o *PaymentMethodResponse) SetExpYearNil(b bool)`

 SetExpYearNil sets the value for ExpYear to be an explicit nil

### UnsetExpYear
`func (o *PaymentMethodResponse) UnsetExpYear()`

UnsetExpYear ensures that no value is present for ExpYear, not even an explicit nil
### GetIsDefault

`func (o *PaymentMethodResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PaymentMethodResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PaymentMethodResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCreatedAt

`func (o *PaymentMethodResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


