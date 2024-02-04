# ResumesResumeConditionFieldsRequiredLengthTitleNotInt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLength** | Pointer to **NullableFloat32** | Максимальная длина для текстовых полей. Рассчитывается для текста без символов переноса строки (&#x60;  &#x60;). &#x60;null&#x60; — если количество не ограничено | [optional] 
**MinLength** | Pointer to **NullableFloat32** | Минимальная длина для текстовых полей. Рассчитывается для текста без символов переноса строки (&#x60;  &#x60;). &#x60;null&#x60; — если количество не ограничено | [optional] 
**Required** | Pointer to **NullableBool** | Является ли поле необходимым? Для строковых значений поле не должно быть &#x60;null&#x60; или &#x60;\&quot;\&quot;&#x60; | [optional] 
**NotIn** | Pointer to **[]string** | Список недопустимых значений | [optional] 

## Methods

### NewResumesResumeConditionFieldsRequiredLengthTitleNotInt

`func NewResumesResumeConditionFieldsRequiredLengthTitleNotInt() *ResumesResumeConditionFieldsRequiredLengthTitleNotInt`

NewResumesResumeConditionFieldsRequiredLengthTitleNotInt instantiates a new ResumesResumeConditionFieldsRequiredLengthTitleNotInt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesResumeConditionFieldsRequiredLengthTitleNotIntWithDefaults

`func NewResumesResumeConditionFieldsRequiredLengthTitleNotIntWithDefaults() *ResumesResumeConditionFieldsRequiredLengthTitleNotInt`

NewResumesResumeConditionFieldsRequiredLengthTitleNotIntWithDefaults instantiates a new ResumesResumeConditionFieldsRequiredLengthTitleNotInt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLength

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) GetMaxLength() float32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) GetMaxLengthOk() (*float32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) SetMaxLength(v float32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetMinLength

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) GetMinLength() float32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) GetMinLengthOk() (*float32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) SetMinLength(v float32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetRequired

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetNotIn

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) GetNotIn() []string`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) GetNotInOk() (*[]string, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) SetNotIn(v []string)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.

### SetNotInNil

`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) SetNotInNil(b bool)`

 SetNotInNil sets the value for NotIn to be an explicit nil

### UnsetNotIn
`func (o *ResumesResumeConditionFieldsRequiredLengthTitleNotInt) UnsetNotIn()`

UnsetNotIn ensures that no value is present for NotIn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


