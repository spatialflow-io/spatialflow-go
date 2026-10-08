# BadgeCountsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentsOpen** | **int32** |  | 
**DlqPending** | Pointer to **NullableInt32** |  | [optional] 
**IntegrationsDegraded** | **int32** |  | 

## Methods

### NewBadgeCountsOut

`func NewBadgeCountsOut(incidentsOpen int32, integrationsDegraded int32, ) *BadgeCountsOut`

NewBadgeCountsOut instantiates a new BadgeCountsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadgeCountsOutWithDefaults

`func NewBadgeCountsOutWithDefaults() *BadgeCountsOut`

NewBadgeCountsOutWithDefaults instantiates a new BadgeCountsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentsOpen

`func (o *BadgeCountsOut) GetIncidentsOpen() int32`

GetIncidentsOpen returns the IncidentsOpen field if non-nil, zero value otherwise.

### GetIncidentsOpenOk

`func (o *BadgeCountsOut) GetIncidentsOpenOk() (*int32, bool)`

GetIncidentsOpenOk returns a tuple with the IncidentsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentsOpen

`func (o *BadgeCountsOut) SetIncidentsOpen(v int32)`

SetIncidentsOpen sets IncidentsOpen field to given value.


### GetDlqPending

`func (o *BadgeCountsOut) GetDlqPending() int32`

GetDlqPending returns the DlqPending field if non-nil, zero value otherwise.

### GetDlqPendingOk

`func (o *BadgeCountsOut) GetDlqPendingOk() (*int32, bool)`

GetDlqPendingOk returns a tuple with the DlqPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlqPending

`func (o *BadgeCountsOut) SetDlqPending(v int32)`

SetDlqPending sets DlqPending field to given value.

### HasDlqPending

`func (o *BadgeCountsOut) HasDlqPending() bool`

HasDlqPending returns a boolean if a field has been set.

### SetDlqPendingNil

`func (o *BadgeCountsOut) SetDlqPendingNil(b bool)`

 SetDlqPendingNil sets the value for DlqPending to be an explicit nil

### UnsetDlqPending
`func (o *BadgeCountsOut) UnsetDlqPending()`

UnsetDlqPending ensures that no value is present for DlqPending, not even an explicit nil
### GetIntegrationsDegraded

`func (o *BadgeCountsOut) GetIntegrationsDegraded() int32`

GetIntegrationsDegraded returns the IntegrationsDegraded field if non-nil, zero value otherwise.

### GetIntegrationsDegradedOk

`func (o *BadgeCountsOut) GetIntegrationsDegradedOk() (*int32, bool)`

GetIntegrationsDegradedOk returns a tuple with the IntegrationsDegraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsDegraded

`func (o *BadgeCountsOut) SetIntegrationsDegraded(v int32)`

SetIntegrationsDegraded sets IntegrationsDegraded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


