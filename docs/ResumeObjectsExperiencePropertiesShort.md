# ResumeObjectsExperiencePropertiesShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to [**NullableIncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**Company** | Pointer to **NullableString** | Название организации | [optional] 
**CompanyId** | Pointer to **NullableString** | Уникальный идентификатор организации | [optional] 
**CompanyUrl** | Pointer to **NullableString** | Сайт компании | [optional] 
**Employer** | Pointer to [**NullableEmployersEmployerInfoShort**](EmployersEmployerInfoShort.md) |  | [optional] 
**End** | Pointer to **NullableString** | Окончание работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**Industries** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Список отраслей компании. Возможные значения приведены в [справочнике индустрий](#tag/Obshie-spravochniki/operation/get-industries) | [optional] 
**Industry** | Pointer to [**NullableResumeObjectsIndustry**](ResumeObjectsIndustry.md) |  | [optional] 
**Position** | Pointer to **string** | Должность | [optional] 
**Start** | Pointer to **string** | Начало работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 

## Methods

### NewResumeObjectsExperiencePropertiesShort

`func NewResumeObjectsExperiencePropertiesShort() *ResumeObjectsExperiencePropertiesShort`

NewResumeObjectsExperiencePropertiesShort instantiates a new ResumeObjectsExperiencePropertiesShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsExperiencePropertiesShortWithDefaults

`func NewResumeObjectsExperiencePropertiesShortWithDefaults() *ResumeObjectsExperiencePropertiesShort`

NewResumeObjectsExperiencePropertiesShortWithDefaults instantiates a new ResumeObjectsExperiencePropertiesShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *ResumeObjectsExperiencePropertiesShort) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeObjectsExperiencePropertiesShort) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeObjectsExperiencePropertiesShort) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeObjectsExperiencePropertiesShort) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeObjectsExperiencePropertiesShort) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCompany

`func (o *ResumeObjectsExperiencePropertiesShort) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ResumeObjectsExperiencePropertiesShort) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ResumeObjectsExperiencePropertiesShort) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ResumeObjectsExperiencePropertiesShort) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ResumeObjectsExperiencePropertiesShort) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCompanyId

`func (o *ResumeObjectsExperiencePropertiesShort) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ResumeObjectsExperiencePropertiesShort) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ResumeObjectsExperiencePropertiesShort) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *ResumeObjectsExperiencePropertiesShort) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *ResumeObjectsExperiencePropertiesShort) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyUrl

`func (o *ResumeObjectsExperiencePropertiesShort) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *ResumeObjectsExperiencePropertiesShort) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *ResumeObjectsExperiencePropertiesShort) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### SetCompanyUrlNil

`func (o *ResumeObjectsExperiencePropertiesShort) SetCompanyUrlNil(b bool)`

 SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil

### UnsetCompanyUrl
`func (o *ResumeObjectsExperiencePropertiesShort) UnsetCompanyUrl()`

UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
### GetEmployer

`func (o *ResumeObjectsExperiencePropertiesShort) GetEmployer() EmployersEmployerInfoShort`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetEmployerOk() (*EmployersEmployerInfoShort, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *ResumeObjectsExperiencePropertiesShort) SetEmployer(v EmployersEmployerInfoShort)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *ResumeObjectsExperiencePropertiesShort) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### SetEmployerNil

`func (o *ResumeObjectsExperiencePropertiesShort) SetEmployerNil(b bool)`

 SetEmployerNil sets the value for Employer to be an explicit nil

### UnsetEmployer
`func (o *ResumeObjectsExperiencePropertiesShort) UnsetEmployer()`

UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
### GetEnd

`func (o *ResumeObjectsExperiencePropertiesShort) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResumeObjectsExperiencePropertiesShort) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ResumeObjectsExperiencePropertiesShort) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ResumeObjectsExperiencePropertiesShort) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ResumeObjectsExperiencePropertiesShort) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIndustries

`func (o *ResumeObjectsExperiencePropertiesShort) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *ResumeObjectsExperiencePropertiesShort) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *ResumeObjectsExperiencePropertiesShort) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetIndustry

`func (o *ResumeObjectsExperiencePropertiesShort) GetIndustry() ResumeObjectsIndustry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetIndustryOk() (*ResumeObjectsIndustry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *ResumeObjectsExperiencePropertiesShort) SetIndustry(v ResumeObjectsIndustry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *ResumeObjectsExperiencePropertiesShort) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### SetIndustryNil

`func (o *ResumeObjectsExperiencePropertiesShort) SetIndustryNil(b bool)`

 SetIndustryNil sets the value for Industry to be an explicit nil

### UnsetIndustry
`func (o *ResumeObjectsExperiencePropertiesShort) UnsetIndustry()`

UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
### GetPosition

`func (o *ResumeObjectsExperiencePropertiesShort) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResumeObjectsExperiencePropertiesShort) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ResumeObjectsExperiencePropertiesShort) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStart

`func (o *ResumeObjectsExperiencePropertiesShort) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResumeObjectsExperiencePropertiesShort) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResumeObjectsExperiencePropertiesShort) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *ResumeObjectsExperiencePropertiesShort) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


