# ValidationIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Loc** | [**[]ValidationIssueLocInner**](ValidationIssueLocInner.md) |  | 
**Msg** | **string** |  | 
**Ctx** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewValidationIssue

`func NewValidationIssue(type_ string, loc []ValidationIssueLocInner, msg string, ) *ValidationIssue`

NewValidationIssue instantiates a new ValidationIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationIssueWithDefaults

`func NewValidationIssueWithDefaults() *ValidationIssue`

NewValidationIssueWithDefaults instantiates a new ValidationIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ValidationIssue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidationIssue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidationIssue) SetType(v string)`

SetType sets Type field to given value.


### GetLoc

`func (o *ValidationIssue) GetLoc() []ValidationIssueLocInner`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *ValidationIssue) GetLocOk() (*[]ValidationIssueLocInner, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *ValidationIssue) SetLoc(v []ValidationIssueLocInner)`

SetLoc sets Loc field to given value.


### GetMsg

`func (o *ValidationIssue) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ValidationIssue) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ValidationIssue) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetCtx

`func (o *ValidationIssue) GetCtx() map[string]interface{}`

GetCtx returns the Ctx field if non-nil, zero value otherwise.

### GetCtxOk

`func (o *ValidationIssue) GetCtxOk() (*map[string]interface{}, bool)`

GetCtxOk returns a tuple with the Ctx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtx

`func (o *ValidationIssue) SetCtx(v map[string]interface{})`

SetCtx sets Ctx field to given value.

### HasCtx

`func (o *ValidationIssue) HasCtx() bool`

HasCtx returns a boolean if a field has been set.

### SetCtxNil

`func (o *ValidationIssue) SetCtxNil(b bool)`

 SetCtxNil sets the value for Ctx to be an explicit nil

### UnsetCtx
`func (o *ValidationIssue) UnsetCtx()`

UnsetCtx ensures that no value is present for Ctx, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


