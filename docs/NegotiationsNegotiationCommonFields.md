# NegotiationsNegotiationCommonFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]VacancyNegotiationActions**](VacancyNegotiationActions.md) | Возможные [действия по отклику/приглашению](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#actions-info)  | [optional] 
**EmployerState** | Pointer to [**EmployersEmployersState**](EmployersEmployersState.md) |  | [optional] 
**FunnelStage** | Pointer to [**NullableEmployersFunnelStage**](EmployersFunnelStage.md) |  | [optional] 
**Templates** | Pointer to [**[]VacancyTemplates**](VacancyTemplates.md) | Шаблоны писем | [optional] 
**TestResult** | Pointer to [**NullableSkillVerificationsTestResultWithUrl**](SkillVerificationsTestResultWithUrl.md) |  | [optional] 

## Methods

### NewNegotiationsNegotiationCommonFields

`func NewNegotiationsNegotiationCommonFields() *NegotiationsNegotiationCommonFields`

NewNegotiationsNegotiationCommonFields instantiates a new NegotiationsNegotiationCommonFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsNegotiationCommonFieldsWithDefaults

`func NewNegotiationsNegotiationCommonFieldsWithDefaults() *NegotiationsNegotiationCommonFields`

NewNegotiationsNegotiationCommonFieldsWithDefaults instantiates a new NegotiationsNegotiationCommonFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *NegotiationsNegotiationCommonFields) GetActions() []VacancyNegotiationActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NegotiationsNegotiationCommonFields) GetActionsOk() (*[]VacancyNegotiationActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NegotiationsNegotiationCommonFields) SetActions(v []VacancyNegotiationActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *NegotiationsNegotiationCommonFields) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetEmployerState

`func (o *NegotiationsNegotiationCommonFields) GetEmployerState() EmployersEmployersState`

GetEmployerState returns the EmployerState field if non-nil, zero value otherwise.

### GetEmployerStateOk

`func (o *NegotiationsNegotiationCommonFields) GetEmployerStateOk() (*EmployersEmployersState, bool)`

GetEmployerStateOk returns a tuple with the EmployerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerState

`func (o *NegotiationsNegotiationCommonFields) SetEmployerState(v EmployersEmployersState)`

SetEmployerState sets EmployerState field to given value.

### HasEmployerState

`func (o *NegotiationsNegotiationCommonFields) HasEmployerState() bool`

HasEmployerState returns a boolean if a field has been set.

### GetFunnelStage

`func (o *NegotiationsNegotiationCommonFields) GetFunnelStage() EmployersFunnelStage`

GetFunnelStage returns the FunnelStage field if non-nil, zero value otherwise.

### GetFunnelStageOk

`func (o *NegotiationsNegotiationCommonFields) GetFunnelStageOk() (*EmployersFunnelStage, bool)`

GetFunnelStageOk returns a tuple with the FunnelStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStage

`func (o *NegotiationsNegotiationCommonFields) SetFunnelStage(v EmployersFunnelStage)`

SetFunnelStage sets FunnelStage field to given value.

### HasFunnelStage

`func (o *NegotiationsNegotiationCommonFields) HasFunnelStage() bool`

HasFunnelStage returns a boolean if a field has been set.

### SetFunnelStageNil

`func (o *NegotiationsNegotiationCommonFields) SetFunnelStageNil(b bool)`

 SetFunnelStageNil sets the value for FunnelStage to be an explicit nil

### UnsetFunnelStage
`func (o *NegotiationsNegotiationCommonFields) UnsetFunnelStage()`

UnsetFunnelStage ensures that no value is present for FunnelStage, not even an explicit nil
### GetTemplates

`func (o *NegotiationsNegotiationCommonFields) GetTemplates() []VacancyTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *NegotiationsNegotiationCommonFields) GetTemplatesOk() (*[]VacancyTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *NegotiationsNegotiationCommonFields) SetTemplates(v []VacancyTemplates)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *NegotiationsNegotiationCommonFields) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTestResult

`func (o *NegotiationsNegotiationCommonFields) GetTestResult() SkillVerificationsTestResultWithUrl`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *NegotiationsNegotiationCommonFields) GetTestResultOk() (*SkillVerificationsTestResultWithUrl, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResult

`func (o *NegotiationsNegotiationCommonFields) SetTestResult(v SkillVerificationsTestResultWithUrl)`

SetTestResult sets TestResult field to given value.

### HasTestResult

`func (o *NegotiationsNegotiationCommonFields) HasTestResult() bool`

HasTestResult returns a boolean if a field has been set.

### SetTestResultNil

`func (o *NegotiationsNegotiationCommonFields) SetTestResultNil(b bool)`

 SetTestResultNil sets the value for TestResult to be an explicit nil

### UnsetTestResult
`func (o *NegotiationsNegotiationCommonFields) UnsetTestResult()`

UnsetTestResult ensures that no value is present for TestResult, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


