# ResumesResumeConditionFieldsRequiredCountAndLengthTitle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCount** | Pointer to **NullableFloat32** | Максимальное количество объектов для полей, в которых передается список. Если &#x60;null&#x60; — количество неограниченно | [optional] 
**MinCount** | Pointer to **NullableFloat32** | Минимальное количество объектов для полей,, где необходимо передавать список. Если &#x60;null&#x60; — нижняя граница не определена | [optional] 
**Required** | Pointer to **NullableBool** | Является ли поле необходимым? Для строковых значений поле не должно быть &#x60;null&#x60; или &#x60;\&quot;\&quot;&#x60; | [optional] 
**MaxLength** | Pointer to **NullableFloat32** | Максимальная длина для текстовых полей. Рассчитывается для текста без символов переноса строки (&#x60;  &#x60;). &#x60;null&#x60; — если количество не ограничено | [optional] 
**MinLength** | Pointer to **NullableFloat32** | Минимальная длина для текстовых полей. Рассчитывается для текста без символов переноса строки (&#x60;  &#x60;). &#x60;null&#x60; — если количество не ограничено | [optional] 

## Methods

### NewResumesResumeConditionFieldsRequiredCountAndLengthTitle

`func NewResumesResumeConditionFieldsRequiredCountAndLengthTitle() *ResumesResumeConditionFieldsRequiredCountAndLengthTitle`

NewResumesResumeConditionFieldsRequiredCountAndLengthTitle instantiates a new ResumesResumeConditionFieldsRequiredCountAndLengthTitle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesResumeConditionFieldsRequiredCountAndLengthTitleWithDefaults

`func NewResumesResumeConditionFieldsRequiredCountAndLengthTitleWithDefaults() *ResumesResumeConditionFieldsRequiredCountAndLengthTitle`

NewResumesResumeConditionFieldsRequiredCountAndLengthTitleWithDefaults instantiates a new ResumesResumeConditionFieldsRequiredCountAndLengthTitle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCount

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetMaxCount() float32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetMaxCountOk() (*float32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetMaxCount(v float32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### SetMaxCountNil

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetMaxCountNil(b bool)`

 SetMaxCountNil sets the value for MaxCount to be an explicit nil

### UnsetMaxCount
`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) UnsetMaxCount()`

UnsetMaxCount ensures that no value is present for MaxCount, not even an explicit nil
### GetMinCount

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetMinCount() float32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetMinCountOk() (*float32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetMinCount(v float32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### SetMinCountNil

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetMinCountNil(b bool)`

 SetMinCountNil sets the value for MinCount to be an explicit nil

### UnsetMinCount
`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) UnsetMinCount()`

UnsetMinCount ensures that no value is present for MinCount, not even an explicit nil
### GetRequired

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetMaxLength

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetMaxLength() float32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetMaxLengthOk() (*float32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetMaxLength(v float32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetMinLength

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetMinLength() float32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) GetMinLengthOk() (*float32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetMinLength(v float32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *ResumesResumeConditionFieldsRequiredCountAndLengthTitle) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


