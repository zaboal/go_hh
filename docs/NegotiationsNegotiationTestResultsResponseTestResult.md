# NegotiationsNegotiationTestResultsResponseTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **float32** | Время, затраченное на выполнение теста, в секундах | 
**Mark** | **string** | Дифференцированная оценка за тест: * &#x60;UNFAIR&#x60; — от 0 до 14 баллов; * &#x60;FAIR&#x60; — от 15 до 44 баллов; * &#x60;GOOD&#x60; — от 45 до 79 баллов; * &#x60;EXCELLENT&#x60; — от 80 до 100 баллов  | 
**Name** | **string** | Наименование теста | 
**Score** | **float32** | Результат прохождения теста в баллах (от 0 до 100) | 
**Tasks** | [**[]SkillVerificationsTestResultTasks**](SkillVerificationsTestResultTasks.md) |  | 

## Methods

### NewNegotiationsNegotiationTestResultsResponseTestResult

`func NewNegotiationsNegotiationTestResultsResponseTestResult(duration float32, mark string, name string, score float32, tasks []SkillVerificationsTestResultTasks, ) *NegotiationsNegotiationTestResultsResponseTestResult`

NewNegotiationsNegotiationTestResultsResponseTestResult instantiates a new NegotiationsNegotiationTestResultsResponseTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsNegotiationTestResultsResponseTestResultWithDefaults

`func NewNegotiationsNegotiationTestResultsResponseTestResultWithDefaults() *NegotiationsNegotiationTestResultsResponseTestResult`

NewNegotiationsNegotiationTestResultsResponseTestResultWithDefaults instantiates a new NegotiationsNegotiationTestResultsResponseTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetMark

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetMark() string`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetMarkOk() (*string, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) SetMark(v string)`

SetMark sets Mark field to given value.


### GetName

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) SetName(v string)`

SetName sets Name field to given value.


### GetScore

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) SetScore(v float32)`

SetScore sets Score field to given value.


### GetTasks

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetTasks() []SkillVerificationsTestResultTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) GetTasksOk() (*[]SkillVerificationsTestResultTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *NegotiationsNegotiationTestResultsResponseTestResult) SetTasks(v []SkillVerificationsTestResultTasks)`

SetTasks sets Tasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


