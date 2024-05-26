# EmployersEmployerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateUrl** | **string** | Ссылка на описание работодателя на сайте | 
**Id** | **string** | Идентификатор работодателя | 
**LogoUrls** | Pointer to [**NullableIncludesLogoUrls**](IncludesLogoUrls.md) |  | [optional] 
**Name** | **string** | Название работодателя | 
**Url** | **string** | URL для получения полного описания работодателя | 
**VacanciesUrl** | **string** | URL для получения поисковой выдачи с вакансиями данного работодателя | 
**OpenVacancies** | **float32** | Количество открытых вакансий у работодателя | 

## Methods

### NewEmployersEmployerItem

`func NewEmployersEmployerItem(alternateUrl string, id string, name string, url string, vacanciesUrl string, openVacancies float32, ) *EmployersEmployerItem`

NewEmployersEmployerItem instantiates a new EmployersEmployerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersEmployerItemWithDefaults

`func NewEmployersEmployerItemWithDefaults() *EmployersEmployerItem`

NewEmployersEmployerItemWithDefaults instantiates a new EmployersEmployerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *EmployersEmployerItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *EmployersEmployerItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *EmployersEmployerItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *EmployersEmployerItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployersEmployerItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployersEmployerItem) SetId(v string)`

SetId sets Id field to given value.


### GetLogoUrls

`func (o *EmployersEmployerItem) GetLogoUrls() IncludesLogoUrls`

GetLogoUrls returns the LogoUrls field if non-nil, zero value otherwise.

### GetLogoUrlsOk

`func (o *EmployersEmployerItem) GetLogoUrlsOk() (*IncludesLogoUrls, bool)`

GetLogoUrlsOk returns a tuple with the LogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrls

`func (o *EmployersEmployerItem) SetLogoUrls(v IncludesLogoUrls)`

SetLogoUrls sets LogoUrls field to given value.

### HasLogoUrls

`func (o *EmployersEmployerItem) HasLogoUrls() bool`

HasLogoUrls returns a boolean if a field has been set.

### SetLogoUrlsNil

`func (o *EmployersEmployerItem) SetLogoUrlsNil(b bool)`

 SetLogoUrlsNil sets the value for LogoUrls to be an explicit nil

### UnsetLogoUrls
`func (o *EmployersEmployerItem) UnsetLogoUrls()`

UnsetLogoUrls ensures that no value is present for LogoUrls, not even an explicit nil
### GetName

`func (o *EmployersEmployerItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployersEmployerItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployersEmployerItem) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *EmployersEmployerItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmployersEmployerItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmployersEmployerItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVacanciesUrl

`func (o *EmployersEmployerItem) GetVacanciesUrl() string`

GetVacanciesUrl returns the VacanciesUrl field if non-nil, zero value otherwise.

### GetVacanciesUrlOk

`func (o *EmployersEmployerItem) GetVacanciesUrlOk() (*string, bool)`

GetVacanciesUrlOk returns a tuple with the VacanciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacanciesUrl

`func (o *EmployersEmployerItem) SetVacanciesUrl(v string)`

SetVacanciesUrl sets VacanciesUrl field to given value.


### GetOpenVacancies

`func (o *EmployersEmployerItem) GetOpenVacancies() float32`

GetOpenVacancies returns the OpenVacancies field if non-nil, zero value otherwise.

### GetOpenVacanciesOk

`func (o *EmployersEmployerItem) GetOpenVacanciesOk() (*float32, bool)`

GetOpenVacanciesOk returns a tuple with the OpenVacancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVacancies

`func (o *EmployersEmployerItem) SetOpenVacancies(v float32)`

SetOpenVacancies sets OpenVacancies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


