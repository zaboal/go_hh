# ResumeObjectsEducationAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Идентификатор | [optional] 
**Name** | **string** | Название курса / теста | 
**Organization** | **string** | Организация, проводившая курс / тест | 
**Result** | Pointer to **NullableString** | Специальность / специализация | [optional] 
**Year** | **float32** | Год окончания / сдачи | 

## Methods

### NewResumeObjectsEducationAdditional

`func NewResumeObjectsEducationAdditional(name string, organization string, year float32, ) *ResumeObjectsEducationAdditional`

NewResumeObjectsEducationAdditional instantiates a new ResumeObjectsEducationAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsEducationAdditionalWithDefaults

`func NewResumeObjectsEducationAdditionalWithDefaults() *ResumeObjectsEducationAdditional`

NewResumeObjectsEducationAdditionalWithDefaults instantiates a new ResumeObjectsEducationAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResumeObjectsEducationAdditional) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeObjectsEducationAdditional) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeObjectsEducationAdditional) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResumeObjectsEducationAdditional) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ResumeObjectsEducationAdditional) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResumeObjectsEducationAdditional) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ResumeObjectsEducationAdditional) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResumeObjectsEducationAdditional) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResumeObjectsEducationAdditional) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *ResumeObjectsEducationAdditional) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResumeObjectsEducationAdditional) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResumeObjectsEducationAdditional) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetResult

`func (o *ResumeObjectsEducationAdditional) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResumeObjectsEducationAdditional) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResumeObjectsEducationAdditional) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ResumeObjectsEducationAdditional) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ResumeObjectsEducationAdditional) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ResumeObjectsEducationAdditional) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetYear

`func (o *ResumeObjectsEducationAdditional) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ResumeObjectsEducationAdditional) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ResumeObjectsEducationAdditional) SetYear(v float32)`

SetYear sets Year field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


