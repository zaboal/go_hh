# ResumeEmployerFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**ResumeObjectsActions**](ResumeObjectsActions.md) | Дополнительные действия | 
**CanViewFullInfo** | Pointer to **NullableBool** | Наличие права просмотра контактной информации в резюме | [optional] 
**Favorited** | **bool** | Добавлено ли резюме в избранные | 
**JobSearchStatus** | Pointer to [**IncludesIdName**](IncludesIdName.md) | Для получения данных нужно передать параметр &#x60;with_job_search_status&#x3D;true&#x60;. 
Возможные значения перечислены в поле &#x60;job_search_status&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
 | [optional] 
**NegotiationsHistory** | [**ResumeObjectsNegotiationsHistoryForEmployer**](ResumeObjectsNegotiationsHistoryForEmployer.md) | Краткая история откликов/приглашений по резюме | 
**Owner** | [**ResumeObjectsOwner**](ResumeObjectsOwner.md) | Информация о владельце резюме | 
**PaidServices** | [**[]ResumeObjectsEmployerPaidServicesInner**](ResumeObjectsEmployerPaidServicesInner.md) | Платные услуги по резюме для работодателя  Работодателю может быть предложен список платных услуг по резюме.  Например, если полные данные по резюме недоступны, то будет выдано предложение покупки рекомендованной услуги, чтобы такой доступ получить  | 
**Photo** | Pointer to [**NullableResumeObjectsPhotoNoId**](ResumeObjectsPhotoNoId.md) |  | [optional] 
**Portfolio** | [**[]ResumeObjectsPortfolioNoId**](ResumeObjectsPortfolioNoId.md) | Список изображений в портфолио пользователя | 

## Methods

### NewResumeEmployerFields

`func NewResumeEmployerFields(actions ResumeObjectsActions, favorited bool, negotiationsHistory ResumeObjectsNegotiationsHistoryForEmployer, owner ResumeObjectsOwner, paidServices []ResumeObjectsEmployerPaidServicesInner, portfolio []ResumeObjectsPortfolioNoId, ) *ResumeEmployerFields`

NewResumeEmployerFields instantiates a new ResumeEmployerFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeEmployerFieldsWithDefaults

`func NewResumeEmployerFieldsWithDefaults() *ResumeEmployerFields`

NewResumeEmployerFieldsWithDefaults instantiates a new ResumeEmployerFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ResumeEmployerFields) GetActions() ResumeObjectsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeEmployerFields) GetActionsOk() (*ResumeObjectsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeEmployerFields) SetActions(v ResumeObjectsActions)`

SetActions sets Actions field to given value.


### GetCanViewFullInfo

`func (o *ResumeEmployerFields) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeEmployerFields) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeEmployerFields) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeEmployerFields) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeEmployerFields) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeEmployerFields) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetFavorited

`func (o *ResumeEmployerFields) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *ResumeEmployerFields) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *ResumeEmployerFields) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetJobSearchStatus

`func (o *ResumeEmployerFields) GetJobSearchStatus() IncludesIdName`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *ResumeEmployerFields) GetJobSearchStatusOk() (*IncludesIdName, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *ResumeEmployerFields) SetJobSearchStatus(v IncludesIdName)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *ResumeEmployerFields) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.

### GetNegotiationsHistory

`func (o *ResumeEmployerFields) GetNegotiationsHistory() ResumeObjectsNegotiationsHistoryForEmployer`

GetNegotiationsHistory returns the NegotiationsHistory field if non-nil, zero value otherwise.

### GetNegotiationsHistoryOk

`func (o *ResumeEmployerFields) GetNegotiationsHistoryOk() (*ResumeObjectsNegotiationsHistoryForEmployer, bool)`

GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsHistory

`func (o *ResumeEmployerFields) SetNegotiationsHistory(v ResumeObjectsNegotiationsHistoryForEmployer)`

SetNegotiationsHistory sets NegotiationsHistory field to given value.


### GetOwner

`func (o *ResumeEmployerFields) GetOwner() ResumeObjectsOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResumeEmployerFields) GetOwnerOk() (*ResumeObjectsOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResumeEmployerFields) SetOwner(v ResumeObjectsOwner)`

SetOwner sets Owner field to given value.


### GetPaidServices

`func (o *ResumeEmployerFields) GetPaidServices() []ResumeObjectsEmployerPaidServicesInner`

GetPaidServices returns the PaidServices field if non-nil, zero value otherwise.

### GetPaidServicesOk

`func (o *ResumeEmployerFields) GetPaidServicesOk() (*[]ResumeObjectsEmployerPaidServicesInner, bool)`

GetPaidServicesOk returns a tuple with the PaidServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServices

`func (o *ResumeEmployerFields) SetPaidServices(v []ResumeObjectsEmployerPaidServicesInner)`

SetPaidServices sets PaidServices field to given value.


### GetPhoto

`func (o *ResumeEmployerFields) GetPhoto() ResumeObjectsPhotoNoId`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeEmployerFields) GetPhotoOk() (*ResumeObjectsPhotoNoId, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeEmployerFields) SetPhoto(v ResumeObjectsPhotoNoId)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeEmployerFields) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeEmployerFields) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeEmployerFields) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPortfolio

`func (o *ResumeEmployerFields) GetPortfolio() []ResumeObjectsPortfolioNoId`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *ResumeEmployerFields) GetPortfolioOk() (*[]ResumeObjectsPortfolioNoId, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *ResumeEmployerFields) SetPortfolio(v []ResumeObjectsPortfolioNoId)`

SetPortfolio sets Portfolio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


