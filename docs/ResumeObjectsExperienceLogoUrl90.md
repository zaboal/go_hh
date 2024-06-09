# ResumeObjectsExperienceLogoUrl90

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to [**IncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**Company** | Pointer to **NullableString** | Название организации | [optional] 
**CompanyId** | Pointer to **NullableString** | Уникальный идентификатор организации | [optional] 
**CompanyUrl** | Pointer to **NullableString** | Сайт компании | [optional] 
**Description** | Pointer to **NullableString** | Обязанности, функции, достижения | [optional] 
**Employer** | Pointer to [**EmployersEmployerInfoShortLogoUrl90**](EmployersEmployerInfoShortLogoUrl90.md) |  | [optional] 
**End** | Pointer to **NullableString** | Окончание работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**Industries** | [**[]IncludesIdName**](IncludesIdName.md) | Список отраслей компании. Возможные значения приведены в [справочнике индустрий](#tag/Obshie-spravochniki/operation/get-industries) | 
**Industry** | Pointer to [**ResumeObjectsIndustry**](ResumeObjectsIndustry.md) |  | [optional] 
**Position** | **string** | Должность | 
**Start** | **string** | Начало работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | 

## Methods

### NewResumeObjectsExperienceLogoUrl90

`func NewResumeObjectsExperienceLogoUrl90(industries []IncludesIdName, position string, start string, ) *ResumeObjectsExperienceLogoUrl90`

NewResumeObjectsExperienceLogoUrl90 instantiates a new ResumeObjectsExperienceLogoUrl90 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsExperienceLogoUrl90WithDefaults

`func NewResumeObjectsExperienceLogoUrl90WithDefaults() *ResumeObjectsExperienceLogoUrl90`

NewResumeObjectsExperienceLogoUrl90WithDefaults instantiates a new ResumeObjectsExperienceLogoUrl90 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *ResumeObjectsExperienceLogoUrl90) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeObjectsExperienceLogoUrl90) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeObjectsExperienceLogoUrl90) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCompany

`func (o *ResumeObjectsExperienceLogoUrl90) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ResumeObjectsExperienceLogoUrl90) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ResumeObjectsExperienceLogoUrl90) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ResumeObjectsExperienceLogoUrl90) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ResumeObjectsExperienceLogoUrl90) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCompanyId

`func (o *ResumeObjectsExperienceLogoUrl90) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ResumeObjectsExperienceLogoUrl90) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ResumeObjectsExperienceLogoUrl90) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *ResumeObjectsExperienceLogoUrl90) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *ResumeObjectsExperienceLogoUrl90) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyUrl

`func (o *ResumeObjectsExperienceLogoUrl90) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *ResumeObjectsExperienceLogoUrl90) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *ResumeObjectsExperienceLogoUrl90) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### SetCompanyUrlNil

`func (o *ResumeObjectsExperienceLogoUrl90) SetCompanyUrlNil(b bool)`

 SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil

### UnsetCompanyUrl
`func (o *ResumeObjectsExperienceLogoUrl90) UnsetCompanyUrl()`

UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
### GetDescription

`func (o *ResumeObjectsExperienceLogoUrl90) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResumeObjectsExperienceLogoUrl90) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResumeObjectsExperienceLogoUrl90) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResumeObjectsExperienceLogoUrl90) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResumeObjectsExperienceLogoUrl90) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmployer

`func (o *ResumeObjectsExperienceLogoUrl90) GetEmployer() EmployersEmployerInfoShortLogoUrl90`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetEmployerOk() (*EmployersEmployerInfoShortLogoUrl90, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *ResumeObjectsExperienceLogoUrl90) SetEmployer(v EmployersEmployerInfoShortLogoUrl90)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *ResumeObjectsExperienceLogoUrl90) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetEnd

`func (o *ResumeObjectsExperienceLogoUrl90) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResumeObjectsExperienceLogoUrl90) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ResumeObjectsExperienceLogoUrl90) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ResumeObjectsExperienceLogoUrl90) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ResumeObjectsExperienceLogoUrl90) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIndustries

`func (o *ResumeObjectsExperienceLogoUrl90) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *ResumeObjectsExperienceLogoUrl90) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.


### GetIndustry

`func (o *ResumeObjectsExperienceLogoUrl90) GetIndustry() ResumeObjectsIndustry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetIndustryOk() (*ResumeObjectsIndustry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *ResumeObjectsExperienceLogoUrl90) SetIndustry(v ResumeObjectsIndustry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *ResumeObjectsExperienceLogoUrl90) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetPosition

`func (o *ResumeObjectsExperienceLogoUrl90) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResumeObjectsExperienceLogoUrl90) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetStart

`func (o *ResumeObjectsExperienceLogoUrl90) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResumeObjectsExperienceLogoUrl90) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResumeObjectsExperienceLogoUrl90) SetStart(v string)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


