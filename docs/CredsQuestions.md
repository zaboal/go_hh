# CredsQuestions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Описание вопроса | [optional] 
**IsActive** | Pointer to **bool** | Показан ли вопрос изначально, актуально для динамических вопросов | [optional] 
**PossibleAnswers** | Pointer to **[]string** | Возможные ответы на вопрос, гарантировано придут в поле answers | [optional] 
**QuestionId** | Pointer to **string** | Идентификатор вопроса (совпадает с ключом объекта) | [optional] 
**QuestionTitle** | Pointer to **string** | Текст вопроса отображаемый на форме | [optional] 
**QuestionType** | Pointer to **string** | Возможность мульти выбора ответов на данный вопрос \&quot;single_choice\&quot; / \&quot;multi_select\&quot; | [optional] 
**Required** | Pointer to **bool** | Обязателен ли вопрос для получения ответа | [optional] 
**SkipTitleAtView** | Pointer to **bool** | Пропускать ли текст вопроса на просмотре, если false - ответы внутри placeholder, если true - просто перечисляем без текста вопроса | [optional] 
**ViewTitle** | Pointer to **NullableString** | Текст вопроса на просмотре | [optional] 

## Methods

### NewCredsQuestions

`func NewCredsQuestions() *CredsQuestions`

NewCredsQuestions instantiates a new CredsQuestions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredsQuestionsWithDefaults

`func NewCredsQuestionsWithDefaults() *CredsQuestions`

NewCredsQuestionsWithDefaults instantiates a new CredsQuestions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CredsQuestions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredsQuestions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredsQuestions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredsQuestions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CredsQuestions) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CredsQuestions) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsActive

`func (o *CredsQuestions) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CredsQuestions) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CredsQuestions) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CredsQuestions) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPossibleAnswers

`func (o *CredsQuestions) GetPossibleAnswers() []string`

GetPossibleAnswers returns the PossibleAnswers field if non-nil, zero value otherwise.

### GetPossibleAnswersOk

`func (o *CredsQuestions) GetPossibleAnswersOk() (*[]string, bool)`

GetPossibleAnswersOk returns a tuple with the PossibleAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleAnswers

`func (o *CredsQuestions) SetPossibleAnswers(v []string)`

SetPossibleAnswers sets PossibleAnswers field to given value.

### HasPossibleAnswers

`func (o *CredsQuestions) HasPossibleAnswers() bool`

HasPossibleAnswers returns a boolean if a field has been set.

### GetQuestionId

`func (o *CredsQuestions) GetQuestionId() string`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *CredsQuestions) GetQuestionIdOk() (*string, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *CredsQuestions) SetQuestionId(v string)`

SetQuestionId sets QuestionId field to given value.

### HasQuestionId

`func (o *CredsQuestions) HasQuestionId() bool`

HasQuestionId returns a boolean if a field has been set.

### GetQuestionTitle

`func (o *CredsQuestions) GetQuestionTitle() string`

GetQuestionTitle returns the QuestionTitle field if non-nil, zero value otherwise.

### GetQuestionTitleOk

`func (o *CredsQuestions) GetQuestionTitleOk() (*string, bool)`

GetQuestionTitleOk returns a tuple with the QuestionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionTitle

`func (o *CredsQuestions) SetQuestionTitle(v string)`

SetQuestionTitle sets QuestionTitle field to given value.

### HasQuestionTitle

`func (o *CredsQuestions) HasQuestionTitle() bool`

HasQuestionTitle returns a boolean if a field has been set.

### GetQuestionType

`func (o *CredsQuestions) GetQuestionType() string`

GetQuestionType returns the QuestionType field if non-nil, zero value otherwise.

### GetQuestionTypeOk

`func (o *CredsQuestions) GetQuestionTypeOk() (*string, bool)`

GetQuestionTypeOk returns a tuple with the QuestionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionType

`func (o *CredsQuestions) SetQuestionType(v string)`

SetQuestionType sets QuestionType field to given value.

### HasQuestionType

`func (o *CredsQuestions) HasQuestionType() bool`

HasQuestionType returns a boolean if a field has been set.

### GetRequired

`func (o *CredsQuestions) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CredsQuestions) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CredsQuestions) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CredsQuestions) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSkipTitleAtView

`func (o *CredsQuestions) GetSkipTitleAtView() bool`

GetSkipTitleAtView returns the SkipTitleAtView field if non-nil, zero value otherwise.

### GetSkipTitleAtViewOk

`func (o *CredsQuestions) GetSkipTitleAtViewOk() (*bool, bool)`

GetSkipTitleAtViewOk returns a tuple with the SkipTitleAtView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTitleAtView

`func (o *CredsQuestions) SetSkipTitleAtView(v bool)`

SetSkipTitleAtView sets SkipTitleAtView field to given value.

### HasSkipTitleAtView

`func (o *CredsQuestions) HasSkipTitleAtView() bool`

HasSkipTitleAtView returns a boolean if a field has been set.

### GetViewTitle

`func (o *CredsQuestions) GetViewTitle() string`

GetViewTitle returns the ViewTitle field if non-nil, zero value otherwise.

### GetViewTitleOk

`func (o *CredsQuestions) GetViewTitleOk() (*string, bool)`

GetViewTitleOk returns a tuple with the ViewTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTitle

`func (o *CredsQuestions) SetViewTitle(v string)`

SetViewTitle sets ViewTitle field to given value.

### HasViewTitle

`func (o *CredsQuestions) HasViewTitle() bool`

HasViewTitle returns a boolean if a field has been set.

### SetViewTitleNil

`func (o *CredsQuestions) SetViewTitleNil(b bool)`

 SetViewTitleNil sets the value for ViewTitle to be an explicit nil

### UnsetViewTitle
`func (o *CredsQuestions) UnsetViewTitle()`

UnsetViewTitle ensures that no value is present for ViewTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


