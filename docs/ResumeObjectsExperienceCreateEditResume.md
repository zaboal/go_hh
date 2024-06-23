# ResumeObjectsExperienceCreateEditResume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to [**IncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**Company** | **NullableString** | Название организации | 
**CompanyId** | Pointer to **NullableString** | Уникальный идентификатор организации | [optional] 
**CompanyUrl** | Pointer to **NullableString** | Сайт компании | [optional] 
**Employer** | Pointer to [**EmployersEmployerInfoShort**](EmployersEmployerInfoShort.md) |  | [optional] 
**End** | Pointer to **NullableString** | Окончание работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**Industries** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Список отраслей компании. Возможные значения приведены в [справочнике индустрий](#tag/Obshie-spravochniki/operation/get-industries) | [optional] 
**Industry** | Pointer to [**ResumeObjectsIndustry**](ResumeObjectsIndustry.md) |  | [optional] 
**Position** | **string** | Должность | 
**Start** | **string** | Начало работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | 
**Description** | **NullableString** | Обязанности, функции, достижения | 

## Methods

### NewResumeObjectsExperienceCreateEditResume

`func NewResumeObjectsExperienceCreateEditResume(company NullableString, position string, start string, description NullableString, ) *ResumeObjectsExperienceCreateEditResume`

NewResumeObjectsExperienceCreateEditResume instantiates a new ResumeObjectsExperienceCreateEditResume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsExperienceCreateEditResumeWithDefaults

`func NewResumeObjectsExperienceCreateEditResumeWithDefaults() *ResumeObjectsExperienceCreateEditResume`

NewResumeObjectsExperienceCreateEditResumeWithDefaults instantiates a new ResumeObjectsExperienceCreateEditResume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *ResumeObjectsExperienceCreateEditResume) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeObjectsExperienceCreateEditResume) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeObjectsExperienceCreateEditResume) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCompany

`func (o *ResumeObjectsExperienceCreateEditResume) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ResumeObjectsExperienceCreateEditResume) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ResumeObjectsExperienceCreateEditResume) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCompanyId

`func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ResumeObjectsExperienceCreateEditResume) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *ResumeObjectsExperienceCreateEditResume) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyUrl

`func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *ResumeObjectsExperienceCreateEditResume) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### SetCompanyUrlNil

`func (o *ResumeObjectsExperienceCreateEditResume) SetCompanyUrlNil(b bool)`

 SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil

### UnsetCompanyUrl
`func (o *ResumeObjectsExperienceCreateEditResume) UnsetCompanyUrl()`

UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
### GetEmployer

`func (o *ResumeObjectsExperienceCreateEditResume) GetEmployer() EmployersEmployerInfoShort`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetEmployerOk() (*EmployersEmployerInfoShort, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *ResumeObjectsExperienceCreateEditResume) SetEmployer(v EmployersEmployerInfoShort)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *ResumeObjectsExperienceCreateEditResume) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetEnd

`func (o *ResumeObjectsExperienceCreateEditResume) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResumeObjectsExperienceCreateEditResume) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ResumeObjectsExperienceCreateEditResume) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ResumeObjectsExperienceCreateEditResume) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ResumeObjectsExperienceCreateEditResume) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIndustries

`func (o *ResumeObjectsExperienceCreateEditResume) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *ResumeObjectsExperienceCreateEditResume) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *ResumeObjectsExperienceCreateEditResume) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetIndustry

`func (o *ResumeObjectsExperienceCreateEditResume) GetIndustry() ResumeObjectsIndustry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetIndustryOk() (*ResumeObjectsIndustry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *ResumeObjectsExperienceCreateEditResume) SetIndustry(v ResumeObjectsIndustry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *ResumeObjectsExperienceCreateEditResume) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetPosition

`func (o *ResumeObjectsExperienceCreateEditResume) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResumeObjectsExperienceCreateEditResume) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetStart

`func (o *ResumeObjectsExperienceCreateEditResume) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResumeObjectsExperienceCreateEditResume) SetStart(v string)`

SetStart sets Start field to given value.


### GetDescription

`func (o *ResumeObjectsExperienceCreateEditResume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResumeObjectsExperienceCreateEditResume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResumeObjectsExperienceCreateEditResume) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ResumeObjectsExperienceCreateEditResume) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResumeObjectsExperienceCreateEditResume) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


