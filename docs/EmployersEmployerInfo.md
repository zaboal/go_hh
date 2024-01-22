# EmployersEmployerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccreditedItEmployer** | Pointer to **bool** | Флаг, показывающий, прошел ли работодатель [IT аккредитацию](https://feedback.hh.ru/knowledge-base/article/что-означает-значок-«аккредитована-как-ит-компания») | [optional] 
**AlternateUrl** | **string** | Ссылка на описание работодателя на сайте | 
**ApplicantServices** | Pointer to [**IncludesEmployerApplicantServices**](IncludesEmployerApplicantServices.md) |  | [optional] 
**Area** | [**EmployersEmployerInfoArea**](EmployersEmployerInfoArea.md) |  | 
**BrandedDescription** | Pointer to **NullableString** | Строка с кодом HTML (возможно наличие &#x60;&lt;script/&gt;&#x60; и &#x60;&lt;style/&gt;&#x60;), которая является альтернативой стандартному описанию работодателя. HTML адаптирован для мобильных устройств и корректно отображается без поддержки Javascript.   При этом:  - Контент тянется по ширине на 100% ширины контейнера и умещается без прокрутки в 300px. - Контент рассчитан на то, что он будет вставлен в обвязку, в которую входит название, логотип, сайт и ссылка на вакансии работодателя. - Изображения, которые могут встретиться в таком описании, адаптированы под retina-дисплеи. - Размер шрифта не меньше 12px, размер межстрочного интервала не меньше 16px.  Значение может быть &#x60;null&#x60;, если у работодателя отсутствует индивидуальное описание  | [optional] 
**Branding** | Pointer to [**NullableEmployersBrandingEmployerBranding**](EmployersBrandingEmployerBranding.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Описание работодателя в виде строки с кодом HTML (без &#x60;&lt;script/&gt;&#x60; и &#x60;&lt;style/&gt;&#x60;) | [optional] 
**Id** | **string** | Идентификатор работодателя | 
**Industries** | [**[]IncludesIdName**](IncludesIdName.md) | Список отраслей работодателя. Элементы [справочника индустрий](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-industries) | 
**InsiderInterviews** | [**[]EmployersInsiderInterviews**](EmployersInsiderInterviews.md) | Список интервью | 
**LogoUrls** | Pointer to [**NullableIncludesLogoUrls**](IncludesLogoUrls.md) |  | [optional] 
**Name** | **string** | Название работодателя | 
**OpenVacancies** | Pointer to **NullableFloat32** | Количество открытых вакансий у работодателя | [optional] 
**Relations** | **[]string** | Если работодатель добавлен в черный список, то вернется &#x60;[&#39;blacklisted&#39;]&#x60;, иначе &#x60;[]&#x60; | 
**SiteUrl** | **string** | Адрес сайта работодателя | 
**Trusted** | **bool** | Флаг, показывающий, прошел ли работодатель [проверку на сайте](https://feedback.hh.ru/article/details/id/5951) | 
**Type** | Pointer to **NullableString** | Тип работодателя (прямой работодатель, кадровое агентство и т.п.). Возможные значения описаны в [справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле &#x60;employer_type&#x60;. Возвращает &#x60;null&#x60;, если тип работодателя скрыт | [optional] 
**VacanciesUrl** | **string** | URL для получения поисковой выдачи с вакансиями данного работодателя | 

## Methods

### NewEmployersEmployerInfo

`func NewEmployersEmployerInfo(alternateUrl string, area EmployersEmployerInfoArea, id string, industries []IncludesIdName, insiderInterviews []EmployersInsiderInterviews, name string, relations []*string, siteUrl string, trusted bool, vacanciesUrl string, ) *EmployersEmployerInfo`

NewEmployersEmployerInfo instantiates a new EmployersEmployerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersEmployerInfoWithDefaults

`func NewEmployersEmployerInfoWithDefaults() *EmployersEmployerInfo`

NewEmployersEmployerInfoWithDefaults instantiates a new EmployersEmployerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccreditedItEmployer

`func (o *EmployersEmployerInfo) GetAccreditedItEmployer() bool`

GetAccreditedItEmployer returns the AccreditedItEmployer field if non-nil, zero value otherwise.

### GetAccreditedItEmployerOk

`func (o *EmployersEmployerInfo) GetAccreditedItEmployerOk() (*bool, bool)`

GetAccreditedItEmployerOk returns a tuple with the AccreditedItEmployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditedItEmployer

`func (o *EmployersEmployerInfo) SetAccreditedItEmployer(v bool)`

SetAccreditedItEmployer sets AccreditedItEmployer field to given value.

### HasAccreditedItEmployer

`func (o *EmployersEmployerInfo) HasAccreditedItEmployer() bool`

HasAccreditedItEmployer returns a boolean if a field has been set.

### GetAlternateUrl

`func (o *EmployersEmployerInfo) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *EmployersEmployerInfo) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *EmployersEmployerInfo) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplicantServices

`func (o *EmployersEmployerInfo) GetApplicantServices() IncludesEmployerApplicantServices`

GetApplicantServices returns the ApplicantServices field if non-nil, zero value otherwise.

### GetApplicantServicesOk

`func (o *EmployersEmployerInfo) GetApplicantServicesOk() (*IncludesEmployerApplicantServices, bool)`

GetApplicantServicesOk returns a tuple with the ApplicantServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantServices

`func (o *EmployersEmployerInfo) SetApplicantServices(v IncludesEmployerApplicantServices)`

SetApplicantServices sets ApplicantServices field to given value.

### HasApplicantServices

`func (o *EmployersEmployerInfo) HasApplicantServices() bool`

HasApplicantServices returns a boolean if a field has been set.

### GetArea

`func (o *EmployersEmployerInfo) GetArea() EmployersEmployerInfoArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *EmployersEmployerInfo) GetAreaOk() (*EmployersEmployerInfoArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *EmployersEmployerInfo) SetArea(v EmployersEmployerInfoArea)`

SetArea sets Area field to given value.


### GetBrandedDescription

`func (o *EmployersEmployerInfo) GetBrandedDescription() string`

GetBrandedDescription returns the BrandedDescription field if non-nil, zero value otherwise.

### GetBrandedDescriptionOk

`func (o *EmployersEmployerInfo) GetBrandedDescriptionOk() (*string, bool)`

GetBrandedDescriptionOk returns a tuple with the BrandedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedDescription

`func (o *EmployersEmployerInfo) SetBrandedDescription(v string)`

SetBrandedDescription sets BrandedDescription field to given value.

### HasBrandedDescription

`func (o *EmployersEmployerInfo) HasBrandedDescription() bool`

HasBrandedDescription returns a boolean if a field has been set.

### SetBrandedDescriptionNil

`func (o *EmployersEmployerInfo) SetBrandedDescriptionNil(b bool)`

 SetBrandedDescriptionNil sets the value for BrandedDescription to be an explicit nil

### UnsetBrandedDescription
`func (o *EmployersEmployerInfo) UnsetBrandedDescription()`

UnsetBrandedDescription ensures that no value is present for BrandedDescription, not even an explicit nil
### GetBranding

`func (o *EmployersEmployerInfo) GetBranding() EmployersBrandingEmployerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *EmployersEmployerInfo) GetBrandingOk() (*EmployersBrandingEmployerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *EmployersEmployerInfo) SetBranding(v EmployersBrandingEmployerBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *EmployersEmployerInfo) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### SetBrandingNil

`func (o *EmployersEmployerInfo) SetBrandingNil(b bool)`

 SetBrandingNil sets the value for Branding to be an explicit nil

### UnsetBranding
`func (o *EmployersEmployerInfo) UnsetBranding()`

UnsetBranding ensures that no value is present for Branding, not even an explicit nil
### GetDescription

`func (o *EmployersEmployerInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmployersEmployerInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmployersEmployerInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmployersEmployerInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EmployersEmployerInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EmployersEmployerInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *EmployersEmployerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployersEmployerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployersEmployerInfo) SetId(v string)`

SetId sets Id field to given value.


### GetIndustries

`func (o *EmployersEmployerInfo) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *EmployersEmployerInfo) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *EmployersEmployerInfo) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.


### GetInsiderInterviews

`func (o *EmployersEmployerInfo) GetInsiderInterviews() []EmployersInsiderInterviews`

GetInsiderInterviews returns the InsiderInterviews field if non-nil, zero value otherwise.

### GetInsiderInterviewsOk

`func (o *EmployersEmployerInfo) GetInsiderInterviewsOk() (*[]EmployersInsiderInterviews, bool)`

GetInsiderInterviewsOk returns a tuple with the InsiderInterviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterviews

`func (o *EmployersEmployerInfo) SetInsiderInterviews(v []EmployersInsiderInterviews)`

SetInsiderInterviews sets InsiderInterviews field to given value.


### GetLogoUrls

`func (o *EmployersEmployerInfo) GetLogoUrls() IncludesLogoUrls`

GetLogoUrls returns the LogoUrls field if non-nil, zero value otherwise.

### GetLogoUrlsOk

`func (o *EmployersEmployerInfo) GetLogoUrlsOk() (*IncludesLogoUrls, bool)`

GetLogoUrlsOk returns a tuple with the LogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrls

`func (o *EmployersEmployerInfo) SetLogoUrls(v IncludesLogoUrls)`

SetLogoUrls sets LogoUrls field to given value.

### HasLogoUrls

`func (o *EmployersEmployerInfo) HasLogoUrls() bool`

HasLogoUrls returns a boolean if a field has been set.

### SetLogoUrlsNil

`func (o *EmployersEmployerInfo) SetLogoUrlsNil(b bool)`

 SetLogoUrlsNil sets the value for LogoUrls to be an explicit nil

### UnsetLogoUrls
`func (o *EmployersEmployerInfo) UnsetLogoUrls()`

UnsetLogoUrls ensures that no value is present for LogoUrls, not even an explicit nil
### GetName

`func (o *EmployersEmployerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployersEmployerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployersEmployerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOpenVacancies

`func (o *EmployersEmployerInfo) GetOpenVacancies() float32`

GetOpenVacancies returns the OpenVacancies field if non-nil, zero value otherwise.

### GetOpenVacanciesOk

`func (o *EmployersEmployerInfo) GetOpenVacanciesOk() (*float32, bool)`

GetOpenVacanciesOk returns a tuple with the OpenVacancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVacancies

`func (o *EmployersEmployerInfo) SetOpenVacancies(v float32)`

SetOpenVacancies sets OpenVacancies field to given value.

### HasOpenVacancies

`func (o *EmployersEmployerInfo) HasOpenVacancies() bool`

HasOpenVacancies returns a boolean if a field has been set.

### SetOpenVacanciesNil

`func (o *EmployersEmployerInfo) SetOpenVacanciesNil(b bool)`

 SetOpenVacanciesNil sets the value for OpenVacancies to be an explicit nil

### UnsetOpenVacancies
`func (o *EmployersEmployerInfo) UnsetOpenVacancies()`

UnsetOpenVacancies ensures that no value is present for OpenVacancies, not even an explicit nil
### GetRelations

`func (o *EmployersEmployerInfo) GetRelations() []*string`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *EmployersEmployerInfo) GetRelationsOk() (*[]*string, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *EmployersEmployerInfo) SetRelations(v []*string)`

SetRelations sets Relations field to given value.


### GetSiteUrl

`func (o *EmployersEmployerInfo) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *EmployersEmployerInfo) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *EmployersEmployerInfo) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.


### GetTrusted

`func (o *EmployersEmployerInfo) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *EmployersEmployerInfo) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *EmployersEmployerInfo) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.


### GetType

`func (o *EmployersEmployerInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmployersEmployerInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmployersEmployerInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmployersEmployerInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *EmployersEmployerInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EmployersEmployerInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVacanciesUrl

`func (o *EmployersEmployerInfo) GetVacanciesUrl() string`

GetVacanciesUrl returns the VacanciesUrl field if non-nil, zero value otherwise.

### GetVacanciesUrlOk

`func (o *EmployersEmployerInfo) GetVacanciesUrlOk() (*string, bool)`

GetVacanciesUrlOk returns a tuple with the VacanciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacanciesUrl

`func (o *EmployersEmployerInfo) SetVacanciesUrl(v string)`

SetVacanciesUrl sets VacanciesUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


