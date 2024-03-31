# SkillVerificationsTestResultNano

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mark** | Pointer to **string** | Дифференцированная оценка за тест:  * &#x60;UNFAIR&#x60; — от 0 до 14 баллов; * &#x60;FAIR&#x60; — от 15 до 44 баллов; * &#x60;GOOD&#x60; — от 45 до 79 баллов; * &#x60;EXCELLENT&#x60; — от 80 до 100 баллов  | [optional] 
**Score** | **int32** | Результат прохождения теста в баллах (от 0 до 100) | 

## Methods

### NewSkillVerificationsTestResultNano

`func NewSkillVerificationsTestResultNano(score int32, ) *SkillVerificationsTestResultNano`

NewSkillVerificationsTestResultNano instantiates a new SkillVerificationsTestResultNano object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillVerificationsTestResultNanoWithDefaults

`func NewSkillVerificationsTestResultNanoWithDefaults() *SkillVerificationsTestResultNano`

NewSkillVerificationsTestResultNanoWithDefaults instantiates a new SkillVerificationsTestResultNano object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMark

`func (o *SkillVerificationsTestResultNano) GetMark() string`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *SkillVerificationsTestResultNano) GetMarkOk() (*string, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *SkillVerificationsTestResultNano) SetMark(v string)`

SetMark sets Mark field to given value.

### HasMark

`func (o *SkillVerificationsTestResultNano) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetScore

`func (o *SkillVerificationsTestResultNano) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SkillVerificationsTestResultNano) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SkillVerificationsTestResultNano) SetScore(v int32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


