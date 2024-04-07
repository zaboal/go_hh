# SkillVerificationsOpenedAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mark** | **string** | Дифференцированная оценка за ответ от работодателя: * &#x60;UNFAIR&#x60; — 0 баллов; * &#x60;FAIR&#x60; — 30 баллов; * &#x60;GOOD&#x60; — 60 баллов; * &#x60;EXCELLENT&#x60; — 100 баллов  | 
**Value** | **string** | Ответ на вопрос | 

## Methods

### NewSkillVerificationsOpenedAnswer

`func NewSkillVerificationsOpenedAnswer(mark string, value string, ) *SkillVerificationsOpenedAnswer`

NewSkillVerificationsOpenedAnswer instantiates a new SkillVerificationsOpenedAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillVerificationsOpenedAnswerWithDefaults

`func NewSkillVerificationsOpenedAnswerWithDefaults() *SkillVerificationsOpenedAnswer`

NewSkillVerificationsOpenedAnswerWithDefaults instantiates a new SkillVerificationsOpenedAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMark

`func (o *SkillVerificationsOpenedAnswer) GetMark() string`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *SkillVerificationsOpenedAnswer) GetMarkOk() (*string, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *SkillVerificationsOpenedAnswer) SetMark(v string)`

SetMark sets Mark field to given value.


### GetValue

`func (o *SkillVerificationsOpenedAnswer) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SkillVerificationsOpenedAnswer) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SkillVerificationsOpenedAnswer) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


