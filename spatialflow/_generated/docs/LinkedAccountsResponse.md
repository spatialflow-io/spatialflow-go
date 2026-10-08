# LinkedAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedAccounts** | [**[]LinkedAccountOut**](LinkedAccountOut.md) |  | 
**HasPassword** | **bool** |  | 

## Methods

### NewLinkedAccountsResponse

`func NewLinkedAccountsResponse(linkedAccounts []LinkedAccountOut, hasPassword bool, ) *LinkedAccountsResponse`

NewLinkedAccountsResponse instantiates a new LinkedAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedAccountsResponseWithDefaults

`func NewLinkedAccountsResponseWithDefaults() *LinkedAccountsResponse`

NewLinkedAccountsResponseWithDefaults instantiates a new LinkedAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedAccounts

`func (o *LinkedAccountsResponse) GetLinkedAccounts() []LinkedAccountOut`

GetLinkedAccounts returns the LinkedAccounts field if non-nil, zero value otherwise.

### GetLinkedAccountsOk

`func (o *LinkedAccountsResponse) GetLinkedAccountsOk() (*[]LinkedAccountOut, bool)`

GetLinkedAccountsOk returns a tuple with the LinkedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccounts

`func (o *LinkedAccountsResponse) SetLinkedAccounts(v []LinkedAccountOut)`

SetLinkedAccounts sets LinkedAccounts field to given value.


### GetHasPassword

`func (o *LinkedAccountsResponse) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *LinkedAccountsResponse) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *LinkedAccountsResponse) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


