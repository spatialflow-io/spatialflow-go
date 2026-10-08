# CurrentSessionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**PhotoCount** | **NullableInt32** |  | 
**NoteCount** | **NullableInt32** |  | 

## Methods

### NewCurrentSessionOut

`func NewCurrentSessionOut(id string, startedAt time.Time, photoCount NullableInt32, noteCount NullableInt32, ) *CurrentSessionOut`

NewCurrentSessionOut instantiates a new CurrentSessionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentSessionOutWithDefaults

`func NewCurrentSessionOutWithDefaults() *CurrentSessionOut`

NewCurrentSessionOutWithDefaults instantiates a new CurrentSessionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrentSessionOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentSessionOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentSessionOut) SetId(v string)`

SetId sets Id field to given value.


### GetStartedAt

`func (o *CurrentSessionOut) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CurrentSessionOut) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CurrentSessionOut) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetPhotoCount

`func (o *CurrentSessionOut) GetPhotoCount() int32`

GetPhotoCount returns the PhotoCount field if non-nil, zero value otherwise.

### GetPhotoCountOk

`func (o *CurrentSessionOut) GetPhotoCountOk() (*int32, bool)`

GetPhotoCountOk returns a tuple with the PhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCount

`func (o *CurrentSessionOut) SetPhotoCount(v int32)`

SetPhotoCount sets PhotoCount field to given value.


### SetPhotoCountNil

`func (o *CurrentSessionOut) SetPhotoCountNil(b bool)`

 SetPhotoCountNil sets the value for PhotoCount to be an explicit nil

### UnsetPhotoCount
`func (o *CurrentSessionOut) UnsetPhotoCount()`

UnsetPhotoCount ensures that no value is present for PhotoCount, not even an explicit nil
### GetNoteCount

`func (o *CurrentSessionOut) GetNoteCount() int32`

GetNoteCount returns the NoteCount field if non-nil, zero value otherwise.

### GetNoteCountOk

`func (o *CurrentSessionOut) GetNoteCountOk() (*int32, bool)`

GetNoteCountOk returns a tuple with the NoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteCount

`func (o *CurrentSessionOut) SetNoteCount(v int32)`

SetNoteCount sets NoteCount field to given value.


### SetNoteCountNil

`func (o *CurrentSessionOut) SetNoteCountNil(b bool)`

 SetNoteCountNil sets the value for NoteCount to be an explicit nil

### UnsetNoteCount
`func (o *CurrentSessionOut) UnsetNoteCount()`

UnsetNoteCount ensures that no value is present for NoteCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


