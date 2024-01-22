# VacanciesAvailableVacancyTypeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailablePublicationsCount** | **float32** | Общее количество публикаций, доступных данному менеджеру | 
**Description** | **string** | Описание типа публикации | 
**Name** | **string** | Название типа публикации | 
**Publications** | [**[]VacancyPublication**](VacancyPublication.md) | Список регионов, где может быть опубликована вакансия и количество публикаций, доступных работодателю | 
**VacancyBillingType** | [**VacanciesAvailableVacancyTypeItemVacancyBillingType**](VacanciesAvailableVacancyTypeItemVacancyBillingType.md) |  | 
**VacancyTypes** | [**[]VacancyType**](VacancyType.md) | Список типов вакансии | 

## Methods

### NewVacanciesAvailableVacancyTypeItem

`func NewVacanciesAvailableVacancyTypeItem(availablePublicationsCount float32, description string, name string, publications []VacancyPublication, vacancyBillingType VacanciesAvailableVacancyTypeItemVacancyBillingType, vacancyTypes []VacancyType, ) *VacanciesAvailableVacancyTypeItem`

NewVacanciesAvailableVacancyTypeItem instantiates a new VacanciesAvailableVacancyTypeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesAvailableVacancyTypeItemWithDefaults

`func NewVacanciesAvailableVacancyTypeItemWithDefaults() *VacanciesAvailableVacancyTypeItem`

NewVacanciesAvailableVacancyTypeItemWithDefaults instantiates a new VacanciesAvailableVacancyTypeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailablePublicationsCount

`func (o *VacanciesAvailableVacancyTypeItem) GetAvailablePublicationsCount() float32`

GetAvailablePublicationsCount returns the AvailablePublicationsCount field if non-nil, zero value otherwise.

### GetAvailablePublicationsCountOk

`func (o *VacanciesAvailableVacancyTypeItem) GetAvailablePublicationsCountOk() (*float32, bool)`

GetAvailablePublicationsCountOk returns a tuple with the AvailablePublicationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePublicationsCount

`func (o *VacanciesAvailableVacancyTypeItem) SetAvailablePublicationsCount(v float32)`

SetAvailablePublicationsCount sets AvailablePublicationsCount field to given value.


### GetDescription

`func (o *VacanciesAvailableVacancyTypeItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacanciesAvailableVacancyTypeItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacanciesAvailableVacancyTypeItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *VacanciesAvailableVacancyTypeItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesAvailableVacancyTypeItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesAvailableVacancyTypeItem) SetName(v string)`

SetName sets Name field to given value.


### GetPublications

`func (o *VacanciesAvailableVacancyTypeItem) GetPublications() []VacancyPublication`

GetPublications returns the Publications field if non-nil, zero value otherwise.

### GetPublicationsOk

`func (o *VacanciesAvailableVacancyTypeItem) GetPublicationsOk() (*[]VacancyPublication, bool)`

GetPublicationsOk returns a tuple with the Publications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublications

`func (o *VacanciesAvailableVacancyTypeItem) SetPublications(v []VacancyPublication)`

SetPublications sets Publications field to given value.


### GetVacancyBillingType

`func (o *VacanciesAvailableVacancyTypeItem) GetVacancyBillingType() VacanciesAvailableVacancyTypeItemVacancyBillingType`

GetVacancyBillingType returns the VacancyBillingType field if non-nil, zero value otherwise.

### GetVacancyBillingTypeOk

`func (o *VacanciesAvailableVacancyTypeItem) GetVacancyBillingTypeOk() (*VacanciesAvailableVacancyTypeItemVacancyBillingType, bool)`

GetVacancyBillingTypeOk returns a tuple with the VacancyBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyBillingType

`func (o *VacanciesAvailableVacancyTypeItem) SetVacancyBillingType(v VacanciesAvailableVacancyTypeItemVacancyBillingType)`

SetVacancyBillingType sets VacancyBillingType field to given value.


### GetVacancyTypes

`func (o *VacanciesAvailableVacancyTypeItem) GetVacancyTypes() []VacancyType`

GetVacancyTypes returns the VacancyTypes field if non-nil, zero value otherwise.

### GetVacancyTypesOk

`func (o *VacanciesAvailableVacancyTypeItem) GetVacancyTypesOk() (*[]VacancyType, bool)`

GetVacancyTypesOk returns a tuple with the VacancyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyTypes

`func (o *VacanciesAvailableVacancyTypeItem) SetVacancyTypes(v []VacancyType)`

SetVacancyTypes sets VacancyTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


