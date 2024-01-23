# ResumeAddResumeRequestAllOfEducation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Additional** | Pointer to [**[]ResumeObjectsEducationAdditional**](ResumeObjectsEducationAdditional.md) | Список куров повышения квалификации | [optional] 
**Attestation** | Pointer to [**[]ResumeObjectsEducationAdditional**](ResumeObjectsEducationAdditional.md) | Список пройденных тестов или экзаменов | [optional] 
**Elementary** | Pointer to [**[]ResumeObjectsEducationElementary**](ResumeObjectsEducationElementary.md) | Среднее образование. Обычно заполняется только при отсутствии высшего образования | [optional] 
**Level** | [**IncludesIdName**](IncludesIdName.md) |  | 
**Primary** | Pointer to [**[]ResumeObjectsEducationPrimary**](ResumeObjectsEducationPrimary.md) | Список образований выше среднего | [optional] 

## Methods

### NewResumeAddResumeRequestAllOfEducation

`func NewResumeAddResumeRequestAllOfEducation(level IncludesIdName, ) *ResumeAddResumeRequestAllOfEducation`

NewResumeAddResumeRequestAllOfEducation instantiates a new ResumeAddResumeRequestAllOfEducation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeAddResumeRequestAllOfEducationWithDefaults

`func NewResumeAddResumeRequestAllOfEducationWithDefaults() *ResumeAddResumeRequestAllOfEducation`

NewResumeAddResumeRequestAllOfEducationWithDefaults instantiates a new ResumeAddResumeRequestAllOfEducation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditional

`func (o *ResumeAddResumeRequestAllOfEducation) GetAdditional() []ResumeObjectsEducationAdditional`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *ResumeAddResumeRequestAllOfEducation) GetAdditionalOk() (*[]ResumeObjectsEducationAdditional, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *ResumeAddResumeRequestAllOfEducation) SetAdditional(v []ResumeObjectsEducationAdditional)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *ResumeAddResumeRequestAllOfEducation) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.

### SetAdditionalNil

`func (o *ResumeAddResumeRequestAllOfEducation) SetAdditionalNil(b bool)`

 SetAdditionalNil sets the value for Additional to be an explicit nil

### UnsetAdditional
`func (o *ResumeAddResumeRequestAllOfEducation) UnsetAdditional()`

UnsetAdditional ensures that no value is present for Additional, not even an explicit nil
### GetAttestation

`func (o *ResumeAddResumeRequestAllOfEducation) GetAttestation() []ResumeObjectsEducationAdditional`

GetAttestation returns the Attestation field if non-nil, zero value otherwise.

### GetAttestationOk

`func (o *ResumeAddResumeRequestAllOfEducation) GetAttestationOk() (*[]ResumeObjectsEducationAdditional, bool)`

GetAttestationOk returns a tuple with the Attestation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestation

`func (o *ResumeAddResumeRequestAllOfEducation) SetAttestation(v []ResumeObjectsEducationAdditional)`

SetAttestation sets Attestation field to given value.

### HasAttestation

`func (o *ResumeAddResumeRequestAllOfEducation) HasAttestation() bool`

HasAttestation returns a boolean if a field has been set.

### SetAttestationNil

`func (o *ResumeAddResumeRequestAllOfEducation) SetAttestationNil(b bool)`

 SetAttestationNil sets the value for Attestation to be an explicit nil

### UnsetAttestation
`func (o *ResumeAddResumeRequestAllOfEducation) UnsetAttestation()`

UnsetAttestation ensures that no value is present for Attestation, not even an explicit nil
### GetElementary

`func (o *ResumeAddResumeRequestAllOfEducation) GetElementary() []ResumeObjectsEducationElementary`

GetElementary returns the Elementary field if non-nil, zero value otherwise.

### GetElementaryOk

`func (o *ResumeAddResumeRequestAllOfEducation) GetElementaryOk() (*[]ResumeObjectsEducationElementary, bool)`

GetElementaryOk returns a tuple with the Elementary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementary

`func (o *ResumeAddResumeRequestAllOfEducation) SetElementary(v []ResumeObjectsEducationElementary)`

SetElementary sets Elementary field to given value.

### HasElementary

`func (o *ResumeAddResumeRequestAllOfEducation) HasElementary() bool`

HasElementary returns a boolean if a field has been set.

### SetElementaryNil

`func (o *ResumeAddResumeRequestAllOfEducation) SetElementaryNil(b bool)`

 SetElementaryNil sets the value for Elementary to be an explicit nil

### UnsetElementary
`func (o *ResumeAddResumeRequestAllOfEducation) UnsetElementary()`

UnsetElementary ensures that no value is present for Elementary, not even an explicit nil
### GetLevel

`func (o *ResumeAddResumeRequestAllOfEducation) GetLevel() IncludesIdName`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ResumeAddResumeRequestAllOfEducation) GetLevelOk() (*IncludesIdName, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ResumeAddResumeRequestAllOfEducation) SetLevel(v IncludesIdName)`

SetLevel sets Level field to given value.


### GetPrimary

`func (o *ResumeAddResumeRequestAllOfEducation) GetPrimary() []ResumeObjectsEducationPrimary`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *ResumeAddResumeRequestAllOfEducation) GetPrimaryOk() (*[]ResumeObjectsEducationPrimary, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *ResumeAddResumeRequestAllOfEducation) SetPrimary(v []ResumeObjectsEducationPrimary)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *ResumeAddResumeRequestAllOfEducation) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *ResumeAddResumeRequestAllOfEducation) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *ResumeAddResumeRequestAllOfEducation) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


