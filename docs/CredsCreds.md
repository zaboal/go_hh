# CredsCreds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**map[string]CredsAnswers**](CredsAnswers.md) |  | [optional] 
**QuestionToAnswerMap** | Pointer to **map[string][]string** |  | [optional] 
**Questions** | Pointer to [**map[string]CredsQuestions**](CredsQuestions.md) |  | [optional] 

## Methods

### NewCredsCreds

`func NewCredsCreds() *CredsCreds`

NewCredsCreds instantiates a new CredsCreds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredsCredsWithDefaults

`func NewCredsCredsWithDefaults() *CredsCreds`

NewCredsCredsWithDefaults instantiates a new CredsCreds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *CredsCreds) GetAnswers() map[string]CredsAnswers`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *CredsCreds) GetAnswersOk() (*map[string]CredsAnswers, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *CredsCreds) SetAnswers(v map[string]CredsAnswers)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *CredsCreds) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetQuestionToAnswerMap

`func (o *CredsCreds) GetQuestionToAnswerMap() map[string][]string`

GetQuestionToAnswerMap returns the QuestionToAnswerMap field if non-nil, zero value otherwise.

### GetQuestionToAnswerMapOk

`func (o *CredsCreds) GetQuestionToAnswerMapOk() (*map[string][]string, bool)`

GetQuestionToAnswerMapOk returns a tuple with the QuestionToAnswerMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionToAnswerMap

`func (o *CredsCreds) SetQuestionToAnswerMap(v map[string][]string)`

SetQuestionToAnswerMap sets QuestionToAnswerMap field to given value.

### HasQuestionToAnswerMap

`func (o *CredsCreds) HasQuestionToAnswerMap() bool`

HasQuestionToAnswerMap returns a boolean if a field has been set.

### GetQuestions

`func (o *CredsCreds) GetQuestions() map[string]CredsQuestions`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *CredsCreds) GetQuestionsOk() (*map[string]CredsQuestions, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *CredsCreds) SetQuestions(v map[string]CredsQuestions)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *CredsCreds) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


