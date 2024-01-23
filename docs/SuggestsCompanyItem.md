# SuggestsCompanyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор организации | 
**LogoUrls** | Pointer to [**SuggestsLogoUrl**](SuggestsLogoUrl.md) |  | [optional] 
**Text** | **string** | Название организации | 
**Url** | Pointer to **string** | Сайт организации | [optional] 
**Area** | [**SuggestsArea**](SuggestsArea.md) |  | 
**Industries** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Сферы деятельности | [optional] 

## Methods

### NewSuggestsCompanyItem

`func NewSuggestsCompanyItem(id string, text string, area SuggestsArea, ) *SuggestsCompanyItem`

NewSuggestsCompanyItem instantiates a new SuggestsCompanyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestsCompanyItemWithDefaults

`func NewSuggestsCompanyItemWithDefaults() *SuggestsCompanyItem`

NewSuggestsCompanyItemWithDefaults instantiates a new SuggestsCompanyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuggestsCompanyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuggestsCompanyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuggestsCompanyItem) SetId(v string)`

SetId sets Id field to given value.


### GetLogoUrls

`func (o *SuggestsCompanyItem) GetLogoUrls() SuggestsLogoUrl`

GetLogoUrls returns the LogoUrls field if non-nil, zero value otherwise.

### GetLogoUrlsOk

`func (o *SuggestsCompanyItem) GetLogoUrlsOk() (*SuggestsLogoUrl, bool)`

GetLogoUrlsOk returns a tuple with the LogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrls

`func (o *SuggestsCompanyItem) SetLogoUrls(v SuggestsLogoUrl)`

SetLogoUrls sets LogoUrls field to given value.

### HasLogoUrls

`func (o *SuggestsCompanyItem) HasLogoUrls() bool`

HasLogoUrls returns a boolean if a field has been set.

### GetText

`func (o *SuggestsCompanyItem) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SuggestsCompanyItem) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SuggestsCompanyItem) SetText(v string)`

SetText sets Text field to given value.


### GetUrl

`func (o *SuggestsCompanyItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SuggestsCompanyItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SuggestsCompanyItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SuggestsCompanyItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetArea

`func (o *SuggestsCompanyItem) GetArea() SuggestsArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *SuggestsCompanyItem) GetAreaOk() (*SuggestsArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *SuggestsCompanyItem) SetArea(v SuggestsArea)`

SetArea sets Area field to given value.


### GetIndustries

`func (o *SuggestsCompanyItem) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *SuggestsCompanyItem) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *SuggestsCompanyItem) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *SuggestsCompanyItem) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


