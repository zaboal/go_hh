# ResumesResumeConditionFieldsSiteCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCount** | Pointer to **NullableFloat32** | Максимальное количество объектов для полей, в которых передается список. Если &#x60;null&#x60; — количество неограниченно | [optional] 
**MinCount** | Pointer to **NullableFloat32** | Минимальное количество объектов для полей,, где необходимо передавать список. Если &#x60;null&#x60; — нижняя граница не определена | [optional] 
**Required** | Pointer to **NullableBool** | Является ли поле необходимым? Для строковых значений поле не должно быть &#x60;null&#x60; или &#x60;\&quot;\&quot;&#x60; | [optional] 
**Fields** | Pointer to [**NullableResumesResumeConditionFieldsSiteFields**](ResumesResumeConditionFieldsSiteFields.md) |  | [optional] 

## Methods

### NewResumesResumeConditionFieldsSiteCondition

`func NewResumesResumeConditionFieldsSiteCondition() *ResumesResumeConditionFieldsSiteCondition`

NewResumesResumeConditionFieldsSiteCondition instantiates a new ResumesResumeConditionFieldsSiteCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesResumeConditionFieldsSiteConditionWithDefaults

`func NewResumesResumeConditionFieldsSiteConditionWithDefaults() *ResumesResumeConditionFieldsSiteCondition`

NewResumesResumeConditionFieldsSiteConditionWithDefaults instantiates a new ResumesResumeConditionFieldsSiteCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCount

`func (o *ResumesResumeConditionFieldsSiteCondition) GetMaxCount() float32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *ResumesResumeConditionFieldsSiteCondition) GetMaxCountOk() (*float32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *ResumesResumeConditionFieldsSiteCondition) SetMaxCount(v float32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *ResumesResumeConditionFieldsSiteCondition) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### SetMaxCountNil

`func (o *ResumesResumeConditionFieldsSiteCondition) SetMaxCountNil(b bool)`

 SetMaxCountNil sets the value for MaxCount to be an explicit nil

### UnsetMaxCount
`func (o *ResumesResumeConditionFieldsSiteCondition) UnsetMaxCount()`

UnsetMaxCount ensures that no value is present for MaxCount, not even an explicit nil
### GetMinCount

`func (o *ResumesResumeConditionFieldsSiteCondition) GetMinCount() float32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *ResumesResumeConditionFieldsSiteCondition) GetMinCountOk() (*float32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *ResumesResumeConditionFieldsSiteCondition) SetMinCount(v float32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *ResumesResumeConditionFieldsSiteCondition) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### SetMinCountNil

`func (o *ResumesResumeConditionFieldsSiteCondition) SetMinCountNil(b bool)`

 SetMinCountNil sets the value for MinCount to be an explicit nil

### UnsetMinCount
`func (o *ResumesResumeConditionFieldsSiteCondition) UnsetMinCount()`

UnsetMinCount ensures that no value is present for MinCount, not even an explicit nil
### GetRequired

`func (o *ResumesResumeConditionFieldsSiteCondition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ResumesResumeConditionFieldsSiteCondition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ResumesResumeConditionFieldsSiteCondition) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ResumesResumeConditionFieldsSiteCondition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ResumesResumeConditionFieldsSiteCondition) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ResumesResumeConditionFieldsSiteCondition) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetFields

`func (o *ResumesResumeConditionFieldsSiteCondition) GetFields() ResumesResumeConditionFieldsSiteFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ResumesResumeConditionFieldsSiteCondition) GetFieldsOk() (*ResumesResumeConditionFieldsSiteFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ResumesResumeConditionFieldsSiteCondition) SetFields(v ResumesResumeConditionFieldsSiteFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ResumesResumeConditionFieldsSiteCondition) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *ResumesResumeConditionFieldsSiteCondition) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *ResumesResumeConditionFieldsSiteCondition) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


