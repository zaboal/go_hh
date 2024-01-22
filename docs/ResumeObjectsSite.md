# ResumeObjectsSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**IncludesIdName**](IncludesIdName.md) | Тип профиля. Элемент справочника [resume_contacts_site_type](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**Url** | Pointer to **NullableString** | Ссылка на профиль или идентификатор | [optional] 

## Methods

### NewResumeObjectsSite

`func NewResumeObjectsSite() *ResumeObjectsSite`

NewResumeObjectsSite instantiates a new ResumeObjectsSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsSiteWithDefaults

`func NewResumeObjectsSiteWithDefaults() *ResumeObjectsSite`

NewResumeObjectsSiteWithDefaults instantiates a new ResumeObjectsSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResumeObjectsSite) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResumeObjectsSite) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResumeObjectsSite) SetType(v IncludesIdName)`

SetType sets Type field to given value.

### HasType

`func (o *ResumeObjectsSite) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *ResumeObjectsSite) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumeObjectsSite) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumeObjectsSite) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ResumeObjectsSite) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ResumeObjectsSite) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ResumeObjectsSite) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


