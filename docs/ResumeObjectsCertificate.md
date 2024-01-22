# ResumeObjectsCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievedAt** | Pointer to **string** | Дата получения (в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**Owner** | Pointer to **NullableString** | На кого выдан сертификат. Возвращается только для сертификатов с &#x60;type &#x3D; microsoft&#x60; | [optional] 
**Title** | Pointer to **string** | Название сертификата | [optional] 
**Type** | Pointer to **string** | Тип сертификата. Доступные значения:  * &#x60;custom&#x60;; * &#x60;microsoft&#x60;  | [optional] 
**Url** | Pointer to **NullableString** | Ссылка на страницу с описанием сертификата | [optional] 

## Methods

### NewResumeObjectsCertificate

`func NewResumeObjectsCertificate() *ResumeObjectsCertificate`

NewResumeObjectsCertificate instantiates a new ResumeObjectsCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsCertificateWithDefaults

`func NewResumeObjectsCertificateWithDefaults() *ResumeObjectsCertificate`

NewResumeObjectsCertificateWithDefaults instantiates a new ResumeObjectsCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievedAt

`func (o *ResumeObjectsCertificate) GetAchievedAt() string`

GetAchievedAt returns the AchievedAt field if non-nil, zero value otherwise.

### GetAchievedAtOk

`func (o *ResumeObjectsCertificate) GetAchievedAtOk() (*string, bool)`

GetAchievedAtOk returns a tuple with the AchievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievedAt

`func (o *ResumeObjectsCertificate) SetAchievedAt(v string)`

SetAchievedAt sets AchievedAt field to given value.

### HasAchievedAt

`func (o *ResumeObjectsCertificate) HasAchievedAt() bool`

HasAchievedAt returns a boolean if a field has been set.

### GetOwner

`func (o *ResumeObjectsCertificate) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResumeObjectsCertificate) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResumeObjectsCertificate) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ResumeObjectsCertificate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ResumeObjectsCertificate) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ResumeObjectsCertificate) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTitle

`func (o *ResumeObjectsCertificate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeObjectsCertificate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeObjectsCertificate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResumeObjectsCertificate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ResumeObjectsCertificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResumeObjectsCertificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResumeObjectsCertificate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResumeObjectsCertificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *ResumeObjectsCertificate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumeObjectsCertificate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumeObjectsCertificate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ResumeObjectsCertificate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ResumeObjectsCertificate) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ResumeObjectsCertificate) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


