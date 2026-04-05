# PolicyListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | [**[]PolicyOut**](PolicyOut.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewPolicyListOut

`func NewPolicyListOut(policies []PolicyOut, totalCount int32, ) *PolicyListOut`

NewPolicyListOut instantiates a new PolicyListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyListOutWithDefaults

`func NewPolicyListOutWithDefaults() *PolicyListOut`

NewPolicyListOutWithDefaults instantiates a new PolicyListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *PolicyListOut) GetPolicies() []PolicyOut`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *PolicyListOut) GetPoliciesOk() (*[]PolicyOut, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *PolicyListOut) SetPolicies(v []PolicyOut)`

SetPolicies sets Policies field to given value.


### GetTotalCount

`func (o *PolicyListOut) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PolicyListOut) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PolicyListOut) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


