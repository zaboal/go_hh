# SkillVerificationsTestResultTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClosedAnswers** | [**[]SkillVerificationsTestResultTasksClosedAnswersInner**](SkillVerificationsTestResultTasksClosedAnswersInner.md) | Варианты ответов на закрытые вопросы | 
**OpenedAnswer** | Pointer to [**NullableSkillVerificationsOpenedAnswer**](SkillVerificationsOpenedAnswer.md) |  | [optional] 
**Question** | **string** | Текст вопроса | 

## Methods

### NewSkillVerificationsTestResultTasks

`func NewSkillVerificationsTestResultTasks(closedAnswers []SkillVerificationsTestResultTasksClosedAnswersInner, question string, ) *SkillVerificationsTestResultTasks`

NewSkillVerificationsTestResultTasks instantiates a new SkillVerificationsTestResultTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillVerificationsTestResultTasksWithDefaults

`func NewSkillVerificationsTestResultTasksWithDefaults() *SkillVerificationsTestResultTasks`

NewSkillVerificationsTestResultTasksWithDefaults instantiates a new SkillVerificationsTestResultTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosedAnswers

`func (o *SkillVerificationsTestResultTasks) GetClosedAnswers() []SkillVerificationsTestResultTasksClosedAnswersInner`

GetClosedAnswers returns the ClosedAnswers field if non-nil, zero value otherwise.

### GetClosedAnswersOk

`func (o *SkillVerificationsTestResultTasks) GetClosedAnswersOk() (*[]SkillVerificationsTestResultTasksClosedAnswersInner, bool)`

GetClosedAnswersOk returns a tuple with the ClosedAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAnswers

`func (o *SkillVerificationsTestResultTasks) SetClosedAnswers(v []SkillVerificationsTestResultTasksClosedAnswersInner)`

SetClosedAnswers sets ClosedAnswers field to given value.


### GetOpenedAnswer

`func (o *SkillVerificationsTestResultTasks) GetOpenedAnswer() SkillVerificationsOpenedAnswer`

GetOpenedAnswer returns the OpenedAnswer field if non-nil, zero value otherwise.

### GetOpenedAnswerOk

`func (o *SkillVerificationsTestResultTasks) GetOpenedAnswerOk() (*SkillVerificationsOpenedAnswer, bool)`

GetOpenedAnswerOk returns a tuple with the OpenedAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAnswer

`func (o *SkillVerificationsTestResultTasks) SetOpenedAnswer(v SkillVerificationsOpenedAnswer)`

SetOpenedAnswer sets OpenedAnswer field to given value.

### HasOpenedAnswer

`func (o *SkillVerificationsTestResultTasks) HasOpenedAnswer() bool`

HasOpenedAnswer returns a boolean if a field has been set.

### SetOpenedAnswerNil

`func (o *SkillVerificationsTestResultTasks) SetOpenedAnswerNil(b bool)`

 SetOpenedAnswerNil sets the value for OpenedAnswer to be an explicit nil

### UnsetOpenedAnswer
`func (o *SkillVerificationsTestResultTasks) UnsetOpenedAnswer()`

UnsetOpenedAnswer ensures that no value is present for OpenedAnswer, not even an explicit nil
### GetQuestion

`func (o *SkillVerificationsTestResultTasks) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *SkillVerificationsTestResultTasks) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *SkillVerificationsTestResultTasks) SetQuestion(v string)`

SetQuestion sets Question field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


