# TripsListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trips** | [**[]TripOut**](TripOut.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewTripsListOut

`func NewTripsListOut(trips []TripOut, totalCount int32, ) *TripsListOut`

NewTripsListOut instantiates a new TripsListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTripsListOutWithDefaults

`func NewTripsListOutWithDefaults() *TripsListOut`

NewTripsListOutWithDefaults instantiates a new TripsListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrips

`func (o *TripsListOut) GetTrips() []TripOut`

GetTrips returns the Trips field if non-nil, zero value otherwise.

### GetTripsOk

`func (o *TripsListOut) GetTripsOk() (*[]TripOut, bool)`

GetTripsOk returns a tuple with the Trips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrips

`func (o *TripsListOut) SetTrips(v []TripOut)`

SetTrips sets Trips field to given value.


### GetTotalCount

`func (o *TripsListOut) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TripsListOut) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TripsListOut) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


