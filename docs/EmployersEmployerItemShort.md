# EmployersEmployerItemShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateUrl** | **string** | Ссылка на описание работодателя на сайте | 
**Id** | **string** | Идентификатор работодателя | 
**LogoUrls** | Pointer to [**NullableIncludesLogoUrls**](IncludesLogoUrls.md) |  | [optional] 
**Name** | **string** | Название работодателя | 
**Url** | **string** | URL для получения полного описания работодателя | 
**VacanciesUrl** | **string** | URL для получения поисковой выдачи с вакансиями данного работодателя | 

## Methods

### NewEmployersEmployerItemShort

`func NewEmployersEmployerItemShort(alternateUrl string, id string, name string, url string, vacanciesUrl string, ) *EmployersEmployerItemShort`

NewEmployersEmployerItemShort instantiates a new EmployersEmployerItemShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersEmployerItemShortWithDefaults

`func NewEmployersEmployerItemShortWithDefaults() *EmployersEmployerItemShort`

NewEmployersEmployerItemShortWithDefaults instantiates a new EmployersEmployerItemShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *EmployersEmployerItemShort) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *EmployersEmployerItemShort) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *EmployersEmployerItemShort) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *EmployersEmployerItemShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployersEmployerItemShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployersEmployerItemShort) SetId(v string)`

SetId sets Id field to given value.


### GetLogoUrls

`func (o *EmployersEmployerItemShort) GetLogoUrls() IncludesLogoUrls`

GetLogoUrls returns the LogoUrls field if non-nil, zero value otherwise.

### GetLogoUrlsOk

`func (o *EmployersEmployerItemShort) GetLogoUrlsOk() (*IncludesLogoUrls, bool)`

GetLogoUrlsOk returns a tuple with the LogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrls

`func (o *EmployersEmployerItemShort) SetLogoUrls(v IncludesLogoUrls)`

SetLogoUrls sets LogoUrls field to given value.

### HasLogoUrls

`func (o *EmployersEmployerItemShort) HasLogoUrls() bool`

HasLogoUrls returns a boolean if a field has been set.

### SetLogoUrlsNil

`func (o *EmployersEmployerItemShort) SetLogoUrlsNil(b bool)`

 SetLogoUrlsNil sets the value for LogoUrls to be an explicit nil

### UnsetLogoUrls
`func (o *EmployersEmployerItemShort) UnsetLogoUrls()`

UnsetLogoUrls ensures that no value is present for LogoUrls, not even an explicit nil
### GetName

`func (o *EmployersEmployerItemShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployersEmployerItemShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployersEmployerItemShort) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *EmployersEmployerItemShort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmployersEmployerItemShort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmployersEmployerItemShort) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVacanciesUrl

`func (o *EmployersEmployerItemShort) GetVacanciesUrl() string`

GetVacanciesUrl returns the VacanciesUrl field if non-nil, zero value otherwise.

### GetVacanciesUrlOk

`func (o *EmployersEmployerItemShort) GetVacanciesUrlOk() (*string, bool)`

GetVacanciesUrlOk returns a tuple with the VacanciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacanciesUrl

`func (o *EmployersEmployerItemShort) SetVacanciesUrl(v string)`

SetVacanciesUrl sets VacanciesUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


