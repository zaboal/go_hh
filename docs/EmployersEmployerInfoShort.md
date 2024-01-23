# EmployersEmployerInfoShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateUrl** | **string** | Ссылка на описание работодателя на сайте | 
**Id** | **string** | Идентификатор работодателя | 
**LogoUrls** | Pointer to [**NullableIncludesLogoUrls**](IncludesLogoUrls.md) |  | [optional] 
**Name** | **string** | Название работодателя | 
**Url** | **string** | URL для получения полного описания работодателя | 

## Methods

### NewEmployersEmployerInfoShort

`func NewEmployersEmployerInfoShort(alternateUrl string, id string, name string, url string, ) *EmployersEmployerInfoShort`

NewEmployersEmployerInfoShort instantiates a new EmployersEmployerInfoShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersEmployerInfoShortWithDefaults

`func NewEmployersEmployerInfoShortWithDefaults() *EmployersEmployerInfoShort`

NewEmployersEmployerInfoShortWithDefaults instantiates a new EmployersEmployerInfoShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *EmployersEmployerInfoShort) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *EmployersEmployerInfoShort) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *EmployersEmployerInfoShort) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *EmployersEmployerInfoShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployersEmployerInfoShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployersEmployerInfoShort) SetId(v string)`

SetId sets Id field to given value.


### GetLogoUrls

`func (o *EmployersEmployerInfoShort) GetLogoUrls() IncludesLogoUrls`

GetLogoUrls returns the LogoUrls field if non-nil, zero value otherwise.

### GetLogoUrlsOk

`func (o *EmployersEmployerInfoShort) GetLogoUrlsOk() (*IncludesLogoUrls, bool)`

GetLogoUrlsOk returns a tuple with the LogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrls

`func (o *EmployersEmployerInfoShort) SetLogoUrls(v IncludesLogoUrls)`

SetLogoUrls sets LogoUrls field to given value.

### HasLogoUrls

`func (o *EmployersEmployerInfoShort) HasLogoUrls() bool`

HasLogoUrls returns a boolean if a field has been set.

### SetLogoUrlsNil

`func (o *EmployersEmployerInfoShort) SetLogoUrlsNil(b bool)`

 SetLogoUrlsNil sets the value for LogoUrls to be an explicit nil

### UnsetLogoUrls
`func (o *EmployersEmployerInfoShort) UnsetLogoUrls()`

UnsetLogoUrls ensures that no value is present for LogoUrls, not even an explicit nil
### GetName

`func (o *EmployersEmployerInfoShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployersEmployerInfoShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployersEmployerInfoShort) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *EmployersEmployerInfoShort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmployersEmployerInfoShort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmployersEmployerInfoShort) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


