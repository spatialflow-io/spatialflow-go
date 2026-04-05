# EmailHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Checks** | **map[string]interface{}** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewEmailHealthResponse

`func NewEmailHealthResponse(status string, checks map[string]interface{}, timestamp time.Time, ) *EmailHealthResponse`

NewEmailHealthResponse instantiates a new EmailHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailHealthResponseWithDefaults

`func NewEmailHealthResponseWithDefaults() *EmailHealthResponse`

NewEmailHealthResponseWithDefaults instantiates a new EmailHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EmailHealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailHealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailHealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetChecks

`func (o *EmailHealthResponse) GetChecks() map[string]interface{}`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *EmailHealthResponse) GetChecksOk() (*map[string]interface{}, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *EmailHealthResponse) SetChecks(v map[string]interface{})`

SetChecks sets Checks field to given value.


### GetTimestamp

`func (o *EmailHealthResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmailHealthResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmailHealthResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


