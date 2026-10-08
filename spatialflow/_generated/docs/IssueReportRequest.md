# IssueReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** |  | [optional] [default to "Other"]
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** |  | 
**Steps** | Pointer to **NullableString** |  | [optional] 
**Expected** | Pointer to **NullableString** |  | [optional] 
**IncludeDiagnostics** | Pointer to **bool** |  | [optional] [default to true]
**IncludePage** | Pointer to **bool** |  | [optional] [default to true]
**IncludeBrowser** | Pointer to **bool** |  | [optional] [default to true]
**IncludeVersion** | Pointer to **bool** |  | [optional] [default to true]
**Diagnostics** | Pointer to **map[string]interface{}** |  | [optional] 
**PageUrl** | Pointer to **NullableString** |  | [optional] 
**Browser** | Pointer to **map[string]interface{}** |  | [optional] 
**AppVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIssueReportRequest

`func NewIssueReportRequest(description string, ) *IssueReportRequest`

NewIssueReportRequest instantiates a new IssueReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueReportRequestWithDefaults

`func NewIssueReportRequestWithDefaults() *IssueReportRequest`

NewIssueReportRequestWithDefaults instantiates a new IssueReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *IssueReportRequest) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *IssueReportRequest) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *IssueReportRequest) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *IssueReportRequest) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetTitle

`func (o *IssueReportRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssueReportRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssueReportRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IssueReportRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *IssueReportRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *IssueReportRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *IssueReportRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueReportRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueReportRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSteps

`func (o *IssueReportRequest) GetSteps() string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *IssueReportRequest) GetStepsOk() (*string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *IssueReportRequest) SetSteps(v string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *IssueReportRequest) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *IssueReportRequest) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *IssueReportRequest) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetExpected

`func (o *IssueReportRequest) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *IssueReportRequest) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *IssueReportRequest) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *IssueReportRequest) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### SetExpectedNil

`func (o *IssueReportRequest) SetExpectedNil(b bool)`

 SetExpectedNil sets the value for Expected to be an explicit nil

### UnsetExpected
`func (o *IssueReportRequest) UnsetExpected()`

UnsetExpected ensures that no value is present for Expected, not even an explicit nil
### GetIncludeDiagnostics

`func (o *IssueReportRequest) GetIncludeDiagnostics() bool`

GetIncludeDiagnostics returns the IncludeDiagnostics field if non-nil, zero value otherwise.

### GetIncludeDiagnosticsOk

`func (o *IssueReportRequest) GetIncludeDiagnosticsOk() (*bool, bool)`

GetIncludeDiagnosticsOk returns a tuple with the IncludeDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDiagnostics

`func (o *IssueReportRequest) SetIncludeDiagnostics(v bool)`

SetIncludeDiagnostics sets IncludeDiagnostics field to given value.

### HasIncludeDiagnostics

`func (o *IssueReportRequest) HasIncludeDiagnostics() bool`

HasIncludeDiagnostics returns a boolean if a field has been set.

### GetIncludePage

`func (o *IssueReportRequest) GetIncludePage() bool`

GetIncludePage returns the IncludePage field if non-nil, zero value otherwise.

### GetIncludePageOk

`func (o *IssueReportRequest) GetIncludePageOk() (*bool, bool)`

GetIncludePageOk returns a tuple with the IncludePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePage

`func (o *IssueReportRequest) SetIncludePage(v bool)`

SetIncludePage sets IncludePage field to given value.

### HasIncludePage

`func (o *IssueReportRequest) HasIncludePage() bool`

HasIncludePage returns a boolean if a field has been set.

### GetIncludeBrowser

`func (o *IssueReportRequest) GetIncludeBrowser() bool`

GetIncludeBrowser returns the IncludeBrowser field if non-nil, zero value otherwise.

### GetIncludeBrowserOk

`func (o *IssueReportRequest) GetIncludeBrowserOk() (*bool, bool)`

GetIncludeBrowserOk returns a tuple with the IncludeBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBrowser

`func (o *IssueReportRequest) SetIncludeBrowser(v bool)`

SetIncludeBrowser sets IncludeBrowser field to given value.

### HasIncludeBrowser

`func (o *IssueReportRequest) HasIncludeBrowser() bool`

HasIncludeBrowser returns a boolean if a field has been set.

### GetIncludeVersion

`func (o *IssueReportRequest) GetIncludeVersion() bool`

GetIncludeVersion returns the IncludeVersion field if non-nil, zero value otherwise.

### GetIncludeVersionOk

`func (o *IssueReportRequest) GetIncludeVersionOk() (*bool, bool)`

GetIncludeVersionOk returns a tuple with the IncludeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVersion

`func (o *IssueReportRequest) SetIncludeVersion(v bool)`

SetIncludeVersion sets IncludeVersion field to given value.

### HasIncludeVersion

`func (o *IssueReportRequest) HasIncludeVersion() bool`

HasIncludeVersion returns a boolean if a field has been set.

### GetDiagnostics

`func (o *IssueReportRequest) GetDiagnostics() map[string]interface{}`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *IssueReportRequest) GetDiagnosticsOk() (*map[string]interface{}, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *IssueReportRequest) SetDiagnostics(v map[string]interface{})`

SetDiagnostics sets Diagnostics field to given value.

### HasDiagnostics

`func (o *IssueReportRequest) HasDiagnostics() bool`

HasDiagnostics returns a boolean if a field has been set.

### GetPageUrl

`func (o *IssueReportRequest) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *IssueReportRequest) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *IssueReportRequest) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.

### HasPageUrl

`func (o *IssueReportRequest) HasPageUrl() bool`

HasPageUrl returns a boolean if a field has been set.

### SetPageUrlNil

`func (o *IssueReportRequest) SetPageUrlNil(b bool)`

 SetPageUrlNil sets the value for PageUrl to be an explicit nil

### UnsetPageUrl
`func (o *IssueReportRequest) UnsetPageUrl()`

UnsetPageUrl ensures that no value is present for PageUrl, not even an explicit nil
### GetBrowser

`func (o *IssueReportRequest) GetBrowser() map[string]interface{}`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *IssueReportRequest) GetBrowserOk() (*map[string]interface{}, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *IssueReportRequest) SetBrowser(v map[string]interface{})`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *IssueReportRequest) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetAppVersion

`func (o *IssueReportRequest) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *IssueReportRequest) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *IssueReportRequest) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *IssueReportRequest) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *IssueReportRequest) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *IssueReportRequest) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


