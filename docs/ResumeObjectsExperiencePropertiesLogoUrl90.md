# ResumeObjectsExperiencePropertiesLogoUrl90

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to [**NullableIncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**Company** | Pointer to **NullableString** | Название организации | [optional] 
**CompanyId** | Pointer to **NullableString** | Уникальный идентификатор организации | [optional] 
**CompanyUrl** | Pointer to **NullableString** | Сайт компании | [optional] 
**Description** | Pointer to **NullableString** | Обязанности, функции, достижения | [optional] 
**Employer** | Pointer to [**NullableEmployersEmployerInfoShortLogoUrl90**](EmployersEmployerInfoShortLogoUrl90.md) |  | [optional] 
**End** | Pointer to **NullableString** | Окончание работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**Industries** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Список отраслей компании. Возможные значения приведены в [справочнике индустрий](#tag/Obshie-spravochniki/operation/get-industries) | [optional] 
**Industry** | Pointer to [**NullableResumeObjectsIndustry**](ResumeObjectsIndustry.md) |  | [optional] 
**Position** | Pointer to **string** | Должность | [optional] 
**Start** | Pointer to **string** | Начало работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 

## Methods

### NewResumeObjectsExperiencePropertiesLogoUrl90

`func NewResumeObjectsExperiencePropertiesLogoUrl90() *ResumeObjectsExperiencePropertiesLogoUrl90`

NewResumeObjectsExperiencePropertiesLogoUrl90 instantiates a new ResumeObjectsExperiencePropertiesLogoUrl90 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsExperiencePropertiesLogoUrl90WithDefaults

`func NewResumeObjectsExperiencePropertiesLogoUrl90WithDefaults() *ResumeObjectsExperiencePropertiesLogoUrl90`

NewResumeObjectsExperiencePropertiesLogoUrl90WithDefaults instantiates a new ResumeObjectsExperiencePropertiesLogoUrl90 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCompany

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCompanyId

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyUrl

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### SetCompanyUrlNil

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetCompanyUrlNil(b bool)`

 SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil

### UnsetCompanyUrl
`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) UnsetCompanyUrl()`

UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
### GetDescription

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmployer

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetEmployer() EmployersEmployerInfoShortLogoUrl90`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetEmployerOk() (*EmployersEmployerInfoShortLogoUrl90, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetEmployer(v EmployersEmployerInfoShortLogoUrl90)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### SetEmployerNil

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetEmployerNil(b bool)`

 SetEmployerNil sets the value for Employer to be an explicit nil

### UnsetEmployer
`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) UnsetEmployer()`

UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
### GetEnd

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIndustries

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetIndustry

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetIndustry() ResumeObjectsIndustry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetIndustryOk() (*ResumeObjectsIndustry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetIndustry(v ResumeObjectsIndustry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### SetIndustryNil

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetIndustryNil(b bool)`

 SetIndustryNil sets the value for Industry to be an explicit nil

### UnsetIndustry
`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) UnsetIndustry()`

UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
### GetPosition

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStart

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *ResumeObjectsExperiencePropertiesLogoUrl90) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


