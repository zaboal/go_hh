# VacanciesVacancyEmployer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccreditedItEmployer** | Pointer to **bool** | Флаг, показывающий, прошла ли компания IT аккредитацию | [optional] 
**AlternateUrl** | Pointer to **NullableString** | Ссылка на представление компании на сайте | [optional] 
**Id** | Pointer to **NullableString** | Идентификатор компании | [optional] 
**LogoUrls** | Pointer to [**IncludesLogoUrls**](IncludesLogoUrls.md) |  | [optional] 
**Name** | **string** | Название компании | 
**Trusted** | **bool** | Флаг, показывающий, прошла ли компания проверку на сайте | 
**Url** | Pointer to **NullableString** | URL, на который нужно сделать GET-запрос, чтобы получить информацию о компании | [optional] 
**VacanciesUrl** | Pointer to **NullableString** | Ссылка на поисковую выдачу вакансий данной компании | [optional] 
**Blacklisted** | Pointer to **bool** | Добавлены ли все вакансии работодателя в [список скрытых](#tag/Skrytye-rabotodateli) | [optional] 
**ApplicantServices** | Pointer to [**IncludesEmployerApplicantServices**](IncludesEmployerApplicantServices.md) |  | [optional] 

## Methods

### NewVacanciesVacancyEmployer

`func NewVacanciesVacancyEmployer(name string, trusted bool, ) *VacanciesVacancyEmployer`

NewVacanciesVacancyEmployer instantiates a new VacanciesVacancyEmployer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyEmployerWithDefaults

`func NewVacanciesVacancyEmployerWithDefaults() *VacanciesVacancyEmployer`

NewVacanciesVacancyEmployerWithDefaults instantiates a new VacanciesVacancyEmployer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccreditedItEmployer

`func (o *VacanciesVacancyEmployer) GetAccreditedItEmployer() bool`

GetAccreditedItEmployer returns the AccreditedItEmployer field if non-nil, zero value otherwise.

### GetAccreditedItEmployerOk

`func (o *VacanciesVacancyEmployer) GetAccreditedItEmployerOk() (*bool, bool)`

GetAccreditedItEmployerOk returns a tuple with the AccreditedItEmployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditedItEmployer

`func (o *VacanciesVacancyEmployer) SetAccreditedItEmployer(v bool)`

SetAccreditedItEmployer sets AccreditedItEmployer field to given value.

### HasAccreditedItEmployer

`func (o *VacanciesVacancyEmployer) HasAccreditedItEmployer() bool`

HasAccreditedItEmployer returns a boolean if a field has been set.

### GetAlternateUrl

`func (o *VacanciesVacancyEmployer) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacancyEmployer) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacancyEmployer) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.

### HasAlternateUrl

`func (o *VacanciesVacancyEmployer) HasAlternateUrl() bool`

HasAlternateUrl returns a boolean if a field has been set.

### SetAlternateUrlNil

`func (o *VacanciesVacancyEmployer) SetAlternateUrlNil(b bool)`

 SetAlternateUrlNil sets the value for AlternateUrl to be an explicit nil

### UnsetAlternateUrl
`func (o *VacanciesVacancyEmployer) UnsetAlternateUrl()`

UnsetAlternateUrl ensures that no value is present for AlternateUrl, not even an explicit nil
### GetId

`func (o *VacanciesVacancyEmployer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacancyEmployer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacancyEmployer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacanciesVacancyEmployer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacanciesVacancyEmployer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacanciesVacancyEmployer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLogoUrls

`func (o *VacanciesVacancyEmployer) GetLogoUrls() IncludesLogoUrls`

GetLogoUrls returns the LogoUrls field if non-nil, zero value otherwise.

### GetLogoUrlsOk

`func (o *VacanciesVacancyEmployer) GetLogoUrlsOk() (*IncludesLogoUrls, bool)`

GetLogoUrlsOk returns a tuple with the LogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrls

`func (o *VacanciesVacancyEmployer) SetLogoUrls(v IncludesLogoUrls)`

SetLogoUrls sets LogoUrls field to given value.

### HasLogoUrls

`func (o *VacanciesVacancyEmployer) HasLogoUrls() bool`

HasLogoUrls returns a boolean if a field has been set.

### GetName

`func (o *VacanciesVacancyEmployer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesVacancyEmployer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesVacancyEmployer) SetName(v string)`

SetName sets Name field to given value.


### GetTrusted

`func (o *VacanciesVacancyEmployer) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *VacanciesVacancyEmployer) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *VacanciesVacancyEmployer) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.


### GetUrl

`func (o *VacanciesVacancyEmployer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesVacancyEmployer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesVacancyEmployer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VacanciesVacancyEmployer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VacanciesVacancyEmployer) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VacanciesVacancyEmployer) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVacanciesUrl

`func (o *VacanciesVacancyEmployer) GetVacanciesUrl() string`

GetVacanciesUrl returns the VacanciesUrl field if non-nil, zero value otherwise.

### GetVacanciesUrlOk

`func (o *VacanciesVacancyEmployer) GetVacanciesUrlOk() (*string, bool)`

GetVacanciesUrlOk returns a tuple with the VacanciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacanciesUrl

`func (o *VacanciesVacancyEmployer) SetVacanciesUrl(v string)`

SetVacanciesUrl sets VacanciesUrl field to given value.

### HasVacanciesUrl

`func (o *VacanciesVacancyEmployer) HasVacanciesUrl() bool`

HasVacanciesUrl returns a boolean if a field has been set.

### SetVacanciesUrlNil

`func (o *VacanciesVacancyEmployer) SetVacanciesUrlNil(b bool)`

 SetVacanciesUrlNil sets the value for VacanciesUrl to be an explicit nil

### UnsetVacanciesUrl
`func (o *VacanciesVacancyEmployer) UnsetVacanciesUrl()`

UnsetVacanciesUrl ensures that no value is present for VacanciesUrl, not even an explicit nil
### GetBlacklisted

`func (o *VacanciesVacancyEmployer) GetBlacklisted() bool`

GetBlacklisted returns the Blacklisted field if non-nil, zero value otherwise.

### GetBlacklistedOk

`func (o *VacanciesVacancyEmployer) GetBlacklistedOk() (*bool, bool)`

GetBlacklistedOk returns a tuple with the Blacklisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklisted

`func (o *VacanciesVacancyEmployer) SetBlacklisted(v bool)`

SetBlacklisted sets Blacklisted field to given value.

### HasBlacklisted

`func (o *VacanciesVacancyEmployer) HasBlacklisted() bool`

HasBlacklisted returns a boolean if a field has been set.

### GetApplicantServices

`func (o *VacanciesVacancyEmployer) GetApplicantServices() IncludesEmployerApplicantServices`

GetApplicantServices returns the ApplicantServices field if non-nil, zero value otherwise.

### GetApplicantServicesOk

`func (o *VacanciesVacancyEmployer) GetApplicantServicesOk() (*IncludesEmployerApplicantServices, bool)`

GetApplicantServicesOk returns a tuple with the ApplicantServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantServices

`func (o *VacanciesVacancyEmployer) SetApplicantServices(v IncludesEmployerApplicantServices)`

SetApplicantServices sets ApplicantServices field to given value.

### HasApplicantServices

`func (o *VacanciesVacancyEmployer) HasApplicantServices() bool`

HasApplicantServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


