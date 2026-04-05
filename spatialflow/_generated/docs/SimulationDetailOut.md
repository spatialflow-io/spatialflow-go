# SimulationDetailOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Simulation** | [**SimulationOut**](SimulationOut.md) |  | 
**Routes** | [**[]SimulationRouteOut**](SimulationRouteOut.md) |  | 
**RecentEvents** | [**[]SimulationEventOut**](SimulationEventOut.md) |  | 

## Methods

### NewSimulationDetailOut

`func NewSimulationDetailOut(simulation SimulationOut, routes []SimulationRouteOut, recentEvents []SimulationEventOut, ) *SimulationDetailOut`

NewSimulationDetailOut instantiates a new SimulationDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationDetailOutWithDefaults

`func NewSimulationDetailOutWithDefaults() *SimulationDetailOut`

NewSimulationDetailOutWithDefaults instantiates a new SimulationDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimulation

`func (o *SimulationDetailOut) GetSimulation() SimulationOut`

GetSimulation returns the Simulation field if non-nil, zero value otherwise.

### GetSimulationOk

`func (o *SimulationDetailOut) GetSimulationOk() (*SimulationOut, bool)`

GetSimulationOk returns a tuple with the Simulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulation

`func (o *SimulationDetailOut) SetSimulation(v SimulationOut)`

SetSimulation sets Simulation field to given value.


### GetRoutes

`func (o *SimulationDetailOut) GetRoutes() []SimulationRouteOut`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *SimulationDetailOut) GetRoutesOk() (*[]SimulationRouteOut, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *SimulationDetailOut) SetRoutes(v []SimulationRouteOut)`

SetRoutes sets Routes field to given value.


### GetRecentEvents

`func (o *SimulationDetailOut) GetRecentEvents() []SimulationEventOut`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *SimulationDetailOut) GetRecentEventsOk() (*[]SimulationEventOut, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *SimulationDetailOut) SetRecentEvents(v []SimulationEventOut)`

SetRecentEvents sets RecentEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


