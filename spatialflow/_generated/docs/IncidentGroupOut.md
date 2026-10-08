# IncidentGroupOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalType** | **string** |  | 
**Title** | **string** |  | 
**Severity** | **string** |  | 
**Count** | **int32** |  | 
**UnacknowledgedCount** | **int32** |  | 
**LatestAt** | **time.Time** |  | 

## Methods

### NewIncidentGroupOut

`func NewIncidentGroupOut(signalType string, title string, severity string, count int32, unacknowledgedCount int32, latestAt time.Time, ) *IncidentGroupOut`

NewIncidentGroupOut instantiates a new IncidentGroupOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentGroupOutWithDefaults

`func NewIncidentGroupOutWithDefaults() *IncidentGroupOut`

NewIncidentGroupOutWithDefaults instantiates a new IncidentGroupOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalType

`func (o *IncidentGroupOut) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *IncidentGroupOut) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *IncidentGroupOut) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetTitle

`func (o *IncidentGroupOut) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IncidentGroupOut) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IncidentGroupOut) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSeverity

`func (o *IncidentGroupOut) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IncidentGroupOut) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IncidentGroupOut) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetCount

`func (o *IncidentGroupOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IncidentGroupOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IncidentGroupOut) SetCount(v int32)`

SetCount sets Count field to given value.


### GetUnacknowledgedCount

`func (o *IncidentGroupOut) GetUnacknowledgedCount() int32`

GetUnacknowledgedCount returns the UnacknowledgedCount field if non-nil, zero value otherwise.

### GetUnacknowledgedCountOk

`func (o *IncidentGroupOut) GetUnacknowledgedCountOk() (*int32, bool)`

GetUnacknowledgedCountOk returns a tuple with the UnacknowledgedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnacknowledgedCount

`func (o *IncidentGroupOut) SetUnacknowledgedCount(v int32)`

SetUnacknowledgedCount sets UnacknowledgedCount field to given value.


### GetLatestAt

`func (o *IncidentGroupOut) GetLatestAt() time.Time`

GetLatestAt returns the LatestAt field if non-nil, zero value otherwise.

### GetLatestAtOk

`func (o *IncidentGroupOut) GetLatestAtOk() (*time.Time, bool)`

GetLatestAtOk returns a tuple with the LatestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAt

`func (o *IncidentGroupOut) SetLatestAt(v time.Time)`

SetLatestAt sets LatestAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


