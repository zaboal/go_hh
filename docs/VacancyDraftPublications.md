# VacancyDraftPublications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingType** | [**VacancyBillingTypeOutput**](VacancyBillingTypeOutput.md) |  | 
**Count** | **float32** | Количество публикаций | 
**PublicationType** | **string** | Тип публикации (справочник [vacancy_billing_type](#tag/Obshie-spravochniki/operation/get-dictionaries)) | 
**VacancyType** | **string** | Тип вакансии (справочник [vacancy_type](#tag/Obshie-spravochniki/operation/get-dictionaries)) | 

## Methods

### NewVacancyDraftPublications

`func NewVacancyDraftPublications(billingType VacancyBillingTypeOutput, count float32, publicationType string, vacancyType string, ) *VacancyDraftPublications`

NewVacancyDraftPublications instantiates a new VacancyDraftPublications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftPublicationsWithDefaults

`func NewVacancyDraftPublicationsWithDefaults() *VacancyDraftPublications`

NewVacancyDraftPublicationsWithDefaults instantiates a new VacancyDraftPublications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingType

`func (o *VacancyDraftPublications) GetBillingType() VacancyBillingTypeOutput`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacancyDraftPublications) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacancyDraftPublications) SetBillingType(v VacancyBillingTypeOutput)`

SetBillingType sets BillingType field to given value.


### SetBillingTypeNil

`func (o *VacancyDraftPublications) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacancyDraftPublications) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetCount

`func (o *VacancyDraftPublications) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VacancyDraftPublications) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VacancyDraftPublications) SetCount(v float32)`

SetCount sets Count field to given value.


### GetPublicationType

`func (o *VacancyDraftPublications) GetPublicationType() string`

GetPublicationType returns the PublicationType field if non-nil, zero value otherwise.

### GetPublicationTypeOk

`func (o *VacancyDraftPublications) GetPublicationTypeOk() (*string, bool)`

GetPublicationTypeOk returns a tuple with the PublicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationType

`func (o *VacancyDraftPublications) SetPublicationType(v string)`

SetPublicationType sets PublicationType field to given value.


### GetVacancyType

`func (o *VacancyDraftPublications) GetVacancyType() string`

GetVacancyType returns the VacancyType field if non-nil, zero value otherwise.

### GetVacancyTypeOk

`func (o *VacancyDraftPublications) GetVacancyTypeOk() (*string, bool)`

GetVacancyTypeOk returns a tuple with the VacancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyType

`func (o *VacancyDraftPublications) SetVacancyType(v string)`

SetVacancyType sets VacancyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


