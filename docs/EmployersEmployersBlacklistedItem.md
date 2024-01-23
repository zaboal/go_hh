# EmployersEmployersBlacklistedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateUrl** | Pointer to **NullableString** | Ссылка на представление компании на сайте | [optional] 
**Id** | Pointer to **NullableString** | Идентификатор компании | [optional] 
**LogoUrls** | Pointer to [**NullableIncludesLogoUrls**](IncludesLogoUrls.md) |  | [optional] 
**Name** | **string** | Название компании | 
**OpenVacancies** | **float32** | Количество открытых вакансий у работодателя | 
**Url** | Pointer to **NullableString** | URL, на который нужно сделать GET-запрос, чтобы получить информацию о компании | [optional] 
**VacanciesUrl** | Pointer to **NullableString** | Ссылка на поисковую выдачу вакансий данной компании | [optional] 

## Methods

### NewEmployersEmployersBlacklistedItem

`func NewEmployersEmployersBlacklistedItem(name string, openVacancies float32, ) *EmployersEmployersBlacklistedItem`

NewEmployersEmployersBlacklistedItem instantiates a new EmployersEmployersBlacklistedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersEmployersBlacklistedItemWithDefaults

`func NewEmployersEmployersBlacklistedItemWithDefaults() *EmployersEmployersBlacklistedItem`

NewEmployersEmployersBlacklistedItemWithDefaults instantiates a new EmployersEmployersBlacklistedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *EmployersEmployersBlacklistedItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *EmployersEmployersBlacklistedItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *EmployersEmployersBlacklistedItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.

### HasAlternateUrl

`func (o *EmployersEmployersBlacklistedItem) HasAlternateUrl() bool`

HasAlternateUrl returns a boolean if a field has been set.

### SetAlternateUrlNil

`func (o *EmployersEmployersBlacklistedItem) SetAlternateUrlNil(b bool)`

 SetAlternateUrlNil sets the value for AlternateUrl to be an explicit nil

### UnsetAlternateUrl
`func (o *EmployersEmployersBlacklistedItem) UnsetAlternateUrl()`

UnsetAlternateUrl ensures that no value is present for AlternateUrl, not even an explicit nil
### GetId

`func (o *EmployersEmployersBlacklistedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployersEmployersBlacklistedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployersEmployersBlacklistedItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployersEmployersBlacklistedItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EmployersEmployersBlacklistedItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EmployersEmployersBlacklistedItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLogoUrls

`func (o *EmployersEmployersBlacklistedItem) GetLogoUrls() IncludesLogoUrls`

GetLogoUrls returns the LogoUrls field if non-nil, zero value otherwise.

### GetLogoUrlsOk

`func (o *EmployersEmployersBlacklistedItem) GetLogoUrlsOk() (*IncludesLogoUrls, bool)`

GetLogoUrlsOk returns a tuple with the LogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrls

`func (o *EmployersEmployersBlacklistedItem) SetLogoUrls(v IncludesLogoUrls)`

SetLogoUrls sets LogoUrls field to given value.

### HasLogoUrls

`func (o *EmployersEmployersBlacklistedItem) HasLogoUrls() bool`

HasLogoUrls returns a boolean if a field has been set.

### SetLogoUrlsNil

`func (o *EmployersEmployersBlacklistedItem) SetLogoUrlsNil(b bool)`

 SetLogoUrlsNil sets the value for LogoUrls to be an explicit nil

### UnsetLogoUrls
`func (o *EmployersEmployersBlacklistedItem) UnsetLogoUrls()`

UnsetLogoUrls ensures that no value is present for LogoUrls, not even an explicit nil
### GetName

`func (o *EmployersEmployersBlacklistedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployersEmployersBlacklistedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployersEmployersBlacklistedItem) SetName(v string)`

SetName sets Name field to given value.


### GetOpenVacancies

`func (o *EmployersEmployersBlacklistedItem) GetOpenVacancies() float32`

GetOpenVacancies returns the OpenVacancies field if non-nil, zero value otherwise.

### GetOpenVacanciesOk

`func (o *EmployersEmployersBlacklistedItem) GetOpenVacanciesOk() (*float32, bool)`

GetOpenVacanciesOk returns a tuple with the OpenVacancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVacancies

`func (o *EmployersEmployersBlacklistedItem) SetOpenVacancies(v float32)`

SetOpenVacancies sets OpenVacancies field to given value.


### GetUrl

`func (o *EmployersEmployersBlacklistedItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmployersEmployersBlacklistedItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmployersEmployersBlacklistedItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EmployersEmployersBlacklistedItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EmployersEmployersBlacklistedItem) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EmployersEmployersBlacklistedItem) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVacanciesUrl

`func (o *EmployersEmployersBlacklistedItem) GetVacanciesUrl() string`

GetVacanciesUrl returns the VacanciesUrl field if non-nil, zero value otherwise.

### GetVacanciesUrlOk

`func (o *EmployersEmployersBlacklistedItem) GetVacanciesUrlOk() (*string, bool)`

GetVacanciesUrlOk returns a tuple with the VacanciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacanciesUrl

`func (o *EmployersEmployersBlacklistedItem) SetVacanciesUrl(v string)`

SetVacanciesUrl sets VacanciesUrl field to given value.

### HasVacanciesUrl

`func (o *EmployersEmployersBlacklistedItem) HasVacanciesUrl() bool`

HasVacanciesUrl returns a boolean if a field has been set.

### SetVacanciesUrlNil

`func (o *EmployersEmployersBlacklistedItem) SetVacanciesUrlNil(b bool)`

 SetVacanciesUrlNil sets the value for VacanciesUrl to be an explicit nil

### UnsetVacanciesUrl
`func (o *EmployersEmployersBlacklistedItem) UnsetVacanciesUrl()`

UnsetVacanciesUrl ensures that no value is present for VacanciesUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


