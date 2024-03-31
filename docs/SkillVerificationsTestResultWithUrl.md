# SkillVerificationsTestResultWithUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mark** | Pointer to **string** | Дифференцированная оценка за тест:  * &#x60;UNFAIR&#x60; — от 0 до 14 баллов; * &#x60;FAIR&#x60; — от 15 до 44 баллов; * &#x60;GOOD&#x60; — от 45 до 79 баллов; * &#x60;EXCELLENT&#x60; — от 80 до 100 баллов  | [optional] 
**Score** | **int32** | Результат прохождения теста в баллах (от 0 до 100) | 
**AlternateUrl** | **string** | Ссылка на результат теста на сайте | 
**Url** | **string** | Ссылка на результат теста | 

## Methods

### NewSkillVerificationsTestResultWithUrl

`func NewSkillVerificationsTestResultWithUrl(score int32, alternateUrl string, url string, ) *SkillVerificationsTestResultWithUrl`

NewSkillVerificationsTestResultWithUrl instantiates a new SkillVerificationsTestResultWithUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillVerificationsTestResultWithUrlWithDefaults

`func NewSkillVerificationsTestResultWithUrlWithDefaults() *SkillVerificationsTestResultWithUrl`

NewSkillVerificationsTestResultWithUrlWithDefaults instantiates a new SkillVerificationsTestResultWithUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMark

`func (o *SkillVerificationsTestResultWithUrl) GetMark() string`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *SkillVerificationsTestResultWithUrl) GetMarkOk() (*string, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *SkillVerificationsTestResultWithUrl) SetMark(v string)`

SetMark sets Mark field to given value.

### HasMark

`func (o *SkillVerificationsTestResultWithUrl) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetScore

`func (o *SkillVerificationsTestResultWithUrl) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SkillVerificationsTestResultWithUrl) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SkillVerificationsTestResultWithUrl) SetScore(v int32)`

SetScore sets Score field to given value.


### GetAlternateUrl

`func (o *SkillVerificationsTestResultWithUrl) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *SkillVerificationsTestResultWithUrl) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *SkillVerificationsTestResultWithUrl) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetUrl

`func (o *SkillVerificationsTestResultWithUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SkillVerificationsTestResultWithUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SkillVerificationsTestResultWithUrl) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


