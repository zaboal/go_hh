# VacanciesEmployerPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccreditedItEmployer** | Pointer to **bool** | Флаг, показывающий, прошла ли компания IT аккредитацию | [optional] 
**AlternateUrl** | Pointer to **NullableString** | Ссылка на представление компании на сайте | [optional] 
**Id** | Pointer to **NullableString** | Идентификатор компании | [optional] 
**LogoUrls** | Pointer to [**NullableIncludesLogoUrls**](IncludesLogoUrls.md) |  | [optional] 
**Name** | **string** | Название компании | 
**Trusted** | **bool** | Флаг, показывающий, прошла ли компания проверку на сайте | 
**Url** | Pointer to **NullableString** | URL, на который нужно сделать GET-запрос, чтобы получить информацию о компании | [optional] 
**VacanciesUrl** | Pointer to **NullableString** | Ссылка на поисковую выдачу вакансий данной компании | [optional] 

## Methods

### NewVacanciesEmployerPublic

`func NewVacanciesEmployerPublic(name string, trusted bool, ) *VacanciesEmployerPublic`

NewVacanciesEmployerPublic instantiates a new VacanciesEmployerPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesEmployerPublicWithDefaults

`func NewVacanciesEmployerPublicWithDefaults() *VacanciesEmployerPublic`

NewVacanciesEmployerPublicWithDefaults instantiates a new VacanciesEmployerPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccreditedItEmployer

`func (o *VacanciesEmployerPublic) GetAccreditedItEmployer() bool`

GetAccreditedItEmployer returns the AccreditedItEmployer field if non-nil, zero value otherwise.

### GetAccreditedItEmployerOk

`func (o *VacanciesEmployerPublic) GetAccreditedItEmployerOk() (*bool, bool)`

GetAccreditedItEmployerOk returns a tuple with the AccreditedItEmployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditedItEmployer

`func (o *VacanciesEmployerPublic) SetAccreditedItEmployer(v bool)`

SetAccreditedItEmployer sets AccreditedItEmployer field to given value.

### HasAccreditedItEmployer

`func (o *VacanciesEmployerPublic) HasAccreditedItEmployer() bool`

HasAccreditedItEmployer returns a boolean if a field has been set.

### GetAlternateUrl

`func (o *VacanciesEmployerPublic) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesEmployerPublic) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesEmployerPublic) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.

### HasAlternateUrl

`func (o *VacanciesEmployerPublic) HasAlternateUrl() bool`

HasAlternateUrl returns a boolean if a field has been set.

### SetAlternateUrlNil

`func (o *VacanciesEmployerPublic) SetAlternateUrlNil(b bool)`

 SetAlternateUrlNil sets the value for AlternateUrl to be an explicit nil

### UnsetAlternateUrl
`func (o *VacanciesEmployerPublic) UnsetAlternateUrl()`

UnsetAlternateUrl ensures that no value is present for AlternateUrl, not even an explicit nil
### GetId

`func (o *VacanciesEmployerPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesEmployerPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesEmployerPublic) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacanciesEmployerPublic) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacanciesEmployerPublic) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacanciesEmployerPublic) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLogoUrls

`func (o *VacanciesEmployerPublic) GetLogoUrls() IncludesLogoUrls`

GetLogoUrls returns the LogoUrls field if non-nil, zero value otherwise.

### GetLogoUrlsOk

`func (o *VacanciesEmployerPublic) GetLogoUrlsOk() (*IncludesLogoUrls, bool)`

GetLogoUrlsOk returns a tuple with the LogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrls

`func (o *VacanciesEmployerPublic) SetLogoUrls(v IncludesLogoUrls)`

SetLogoUrls sets LogoUrls field to given value.

### HasLogoUrls

`func (o *VacanciesEmployerPublic) HasLogoUrls() bool`

HasLogoUrls returns a boolean if a field has been set.

### SetLogoUrlsNil

`func (o *VacanciesEmployerPublic) SetLogoUrlsNil(b bool)`

 SetLogoUrlsNil sets the value for LogoUrls to be an explicit nil

### UnsetLogoUrls
`func (o *VacanciesEmployerPublic) UnsetLogoUrls()`

UnsetLogoUrls ensures that no value is present for LogoUrls, not even an explicit nil
### GetName

`func (o *VacanciesEmployerPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesEmployerPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesEmployerPublic) SetName(v string)`

SetName sets Name field to given value.


### GetTrusted

`func (o *VacanciesEmployerPublic) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *VacanciesEmployerPublic) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *VacanciesEmployerPublic) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.


### GetUrl

`func (o *VacanciesEmployerPublic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesEmployerPublic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesEmployerPublic) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VacanciesEmployerPublic) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VacanciesEmployerPublic) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VacanciesEmployerPublic) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVacanciesUrl

`func (o *VacanciesEmployerPublic) GetVacanciesUrl() string`

GetVacanciesUrl returns the VacanciesUrl field if non-nil, zero value otherwise.

### GetVacanciesUrlOk

`func (o *VacanciesEmployerPublic) GetVacanciesUrlOk() (*string, bool)`

GetVacanciesUrlOk returns a tuple with the VacanciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacanciesUrl

`func (o *VacanciesEmployerPublic) SetVacanciesUrl(v string)`

SetVacanciesUrl sets VacanciesUrl field to given value.

### HasVacanciesUrl

`func (o *VacanciesEmployerPublic) HasVacanciesUrl() bool`

HasVacanciesUrl returns a boolean if a field has been set.

### SetVacanciesUrlNil

`func (o *VacanciesEmployerPublic) SetVacanciesUrlNil(b bool)`

 SetVacanciesUrlNil sets the value for VacanciesUrl to be an explicit nil

### UnsetVacanciesUrl
`func (o *VacanciesEmployerPublic) UnsetVacanciesUrl()`

UnsetVacanciesUrl ensures that no value is present for VacanciesUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


