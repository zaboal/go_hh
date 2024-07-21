# ResumeObjectsEducationPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Идентификатор | [optional] 
**Name** | **string** | Название учебного заведения | 
**NameId** | Pointer to **NullableString** | Идентификатор учебного заведения | [optional] 
**Organization** | Pointer to **NullableString** | Факультет | [optional] 
**OrganizationId** | Pointer to **NullableString** | Идентификатор факультета | [optional] 
**Result** | Pointer to **NullableString** | Специальность / специализация | [optional] 
**ResultId** | Pointer to **NullableString** | Идентификатор специальности / специализации | [optional] 
**Year** | **float32** | Год окончания | 

## Methods

### NewResumeObjectsEducationPrimary

`func NewResumeObjectsEducationPrimary(name string, year float32, ) *ResumeObjectsEducationPrimary`

NewResumeObjectsEducationPrimary instantiates a new ResumeObjectsEducationPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsEducationPrimaryWithDefaults

`func NewResumeObjectsEducationPrimaryWithDefaults() *ResumeObjectsEducationPrimary`

NewResumeObjectsEducationPrimaryWithDefaults instantiates a new ResumeObjectsEducationPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResumeObjectsEducationPrimary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeObjectsEducationPrimary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeObjectsEducationPrimary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResumeObjectsEducationPrimary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ResumeObjectsEducationPrimary) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResumeObjectsEducationPrimary) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ResumeObjectsEducationPrimary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResumeObjectsEducationPrimary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResumeObjectsEducationPrimary) SetName(v string)`

SetName sets Name field to given value.


### GetNameId

`func (o *ResumeObjectsEducationPrimary) GetNameId() string`

GetNameId returns the NameId field if non-nil, zero value otherwise.

### GetNameIdOk

`func (o *ResumeObjectsEducationPrimary) GetNameIdOk() (*string, bool)`

GetNameIdOk returns a tuple with the NameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameId

`func (o *ResumeObjectsEducationPrimary) SetNameId(v string)`

SetNameId sets NameId field to given value.

### HasNameId

`func (o *ResumeObjectsEducationPrimary) HasNameId() bool`

HasNameId returns a boolean if a field has been set.

### SetNameIdNil

`func (o *ResumeObjectsEducationPrimary) SetNameIdNil(b bool)`

 SetNameIdNil sets the value for NameId to be an explicit nil

### UnsetNameId
`func (o *ResumeObjectsEducationPrimary) UnsetNameId()`

UnsetNameId ensures that no value is present for NameId, not even an explicit nil
### GetOrganization

`func (o *ResumeObjectsEducationPrimary) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResumeObjectsEducationPrimary) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResumeObjectsEducationPrimary) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ResumeObjectsEducationPrimary) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ResumeObjectsEducationPrimary) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ResumeObjectsEducationPrimary) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetOrganizationId

`func (o *ResumeObjectsEducationPrimary) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ResumeObjectsEducationPrimary) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ResumeObjectsEducationPrimary) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ResumeObjectsEducationPrimary) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *ResumeObjectsEducationPrimary) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *ResumeObjectsEducationPrimary) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetResult

`func (o *ResumeObjectsEducationPrimary) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResumeObjectsEducationPrimary) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResumeObjectsEducationPrimary) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ResumeObjectsEducationPrimary) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ResumeObjectsEducationPrimary) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ResumeObjectsEducationPrimary) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetResultId

`func (o *ResumeObjectsEducationPrimary) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *ResumeObjectsEducationPrimary) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *ResumeObjectsEducationPrimary) SetResultId(v string)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *ResumeObjectsEducationPrimary) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### SetResultIdNil

`func (o *ResumeObjectsEducationPrimary) SetResultIdNil(b bool)`

 SetResultIdNil sets the value for ResultId to be an explicit nil

### UnsetResultId
`func (o *ResumeObjectsEducationPrimary) UnsetResultId()`

UnsetResultId ensures that no value is present for ResultId, not even an explicit nil
### GetYear

`func (o *ResumeObjectsEducationPrimary) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ResumeObjectsEducationPrimary) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ResumeObjectsEducationPrimary) SetYear(v float32)`

SetYear sets Year field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


