# ResumeObjectsExperienceShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to [**IncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**Company** | Pointer to **NullableString** | Название организации | [optional] 
**CompanyId** | Pointer to **NullableString** | Уникальный идентификатор организации | [optional] 
**CompanyUrl** | Pointer to **NullableString** | Сайт компании | [optional] 
**Employer** | Pointer to [**EmployersEmployerInfoShort**](EmployersEmployerInfoShort.md) |  | [optional] 
**End** | Pointer to **NullableString** | Окончание работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**Industries** | [**[]IncludesIdName**](IncludesIdName.md) | Список отраслей компании. Возможные значения приведены в [справочнике индустрий](#tag/Obshie-spravochniki/operation/get-industries) | 
**Industry** | Pointer to [**ResumeObjectsIndustry**](ResumeObjectsIndustry.md) |  | [optional] 
**Position** | **string** | Должность | 
**Start** | **string** | Начало работы (дата в формате &#x60;ГГГГ-ММ-ДД&#x60;) | 

## Methods

### NewResumeObjectsExperienceShort

`func NewResumeObjectsExperienceShort(industries []IncludesIdName, position string, start string, ) *ResumeObjectsExperienceShort`

NewResumeObjectsExperienceShort instantiates a new ResumeObjectsExperienceShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsExperienceShortWithDefaults

`func NewResumeObjectsExperienceShortWithDefaults() *ResumeObjectsExperienceShort`

NewResumeObjectsExperienceShortWithDefaults instantiates a new ResumeObjectsExperienceShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *ResumeObjectsExperienceShort) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeObjectsExperienceShort) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeObjectsExperienceShort) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeObjectsExperienceShort) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCompany

`func (o *ResumeObjectsExperienceShort) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ResumeObjectsExperienceShort) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ResumeObjectsExperienceShort) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ResumeObjectsExperienceShort) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ResumeObjectsExperienceShort) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ResumeObjectsExperienceShort) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCompanyId

`func (o *ResumeObjectsExperienceShort) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ResumeObjectsExperienceShort) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ResumeObjectsExperienceShort) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ResumeObjectsExperienceShort) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *ResumeObjectsExperienceShort) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *ResumeObjectsExperienceShort) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyUrl

`func (o *ResumeObjectsExperienceShort) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *ResumeObjectsExperienceShort) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *ResumeObjectsExperienceShort) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *ResumeObjectsExperienceShort) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### SetCompanyUrlNil

`func (o *ResumeObjectsExperienceShort) SetCompanyUrlNil(b bool)`

 SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil

### UnsetCompanyUrl
`func (o *ResumeObjectsExperienceShort) UnsetCompanyUrl()`

UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
### GetEmployer

`func (o *ResumeObjectsExperienceShort) GetEmployer() EmployersEmployerInfoShort`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *ResumeObjectsExperienceShort) GetEmployerOk() (*EmployersEmployerInfoShort, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *ResumeObjectsExperienceShort) SetEmployer(v EmployersEmployerInfoShort)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *ResumeObjectsExperienceShort) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetEnd

`func (o *ResumeObjectsExperienceShort) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResumeObjectsExperienceShort) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResumeObjectsExperienceShort) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ResumeObjectsExperienceShort) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ResumeObjectsExperienceShort) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ResumeObjectsExperienceShort) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIndustries

`func (o *ResumeObjectsExperienceShort) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *ResumeObjectsExperienceShort) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *ResumeObjectsExperienceShort) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.


### GetIndustry

`func (o *ResumeObjectsExperienceShort) GetIndustry() ResumeObjectsIndustry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *ResumeObjectsExperienceShort) GetIndustryOk() (*ResumeObjectsIndustry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *ResumeObjectsExperienceShort) SetIndustry(v ResumeObjectsIndustry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *ResumeObjectsExperienceShort) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetPosition

`func (o *ResumeObjectsExperienceShort) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResumeObjectsExperienceShort) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResumeObjectsExperienceShort) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetStart

`func (o *ResumeObjectsExperienceShort) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResumeObjectsExperienceShort) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResumeObjectsExperienceShort) SetStart(v string)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


