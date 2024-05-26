# VacanciesMatchVacancyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**NullableVacanciesAddress**](VacanciesAddress.md) |  | [optional] 
**CanInvite** | **bool** | Можно ли пригласить соискателя на данную вакансию | 
**CreatedAt** | **string** | Дата и время публикации вакансии | 
**EmployerNegotiationsState** | [**NullableIncludesIdName**](IncludesIdName.md) |  | 
**Manager** | [**VacancyManagerOutput**](VacancyManagerOutput.md) |  | 
**NegotiationsActions** | [**[]VacancyNegotiationActions**](VacancyNegotiationActions.md) | Действия для [создания отклика](#tag/Otklikipriglasheniya-rabotodatelya/operation/invite-applicant-to-vacancy). Если создать отклик невозможно (например, нет нужных услуг), то вернется пустой массив | 
**NegotiationsState** | [**NullableIncludesIdName**](IncludesIdName.md) |  | 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций. Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60;  | [optional] 
**Templates** | Pointer to [**[]VacancyTemplates**](VacancyTemplates.md) | Шаблоны писем | [optional] 

## Methods

### NewVacanciesMatchVacancyFields

`func NewVacanciesMatchVacancyFields(canInvite bool, createdAt string, employerNegotiationsState NullableIncludesIdName, manager VacancyManagerOutput, negotiationsActions []VacancyNegotiationActions, negotiationsState NullableIncludesIdName, ) *VacanciesMatchVacancyFields`

NewVacanciesMatchVacancyFields instantiates a new VacanciesMatchVacancyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesMatchVacancyFieldsWithDefaults

`func NewVacanciesMatchVacancyFieldsWithDefaults() *VacanciesMatchVacancyFields`

NewVacanciesMatchVacancyFieldsWithDefaults instantiates a new VacanciesMatchVacancyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VacanciesMatchVacancyFields) GetAddress() VacanciesAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesMatchVacancyFields) GetAddressOk() (*VacanciesAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesMatchVacancyFields) SetAddress(v VacanciesAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesMatchVacancyFields) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacanciesMatchVacancyFields) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacanciesMatchVacancyFields) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCanInvite

`func (o *VacanciesMatchVacancyFields) GetCanInvite() bool`

GetCanInvite returns the CanInvite field if non-nil, zero value otherwise.

### GetCanInviteOk

`func (o *VacanciesMatchVacancyFields) GetCanInviteOk() (*bool, bool)`

GetCanInviteOk returns a tuple with the CanInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInvite

`func (o *VacanciesMatchVacancyFields) SetCanInvite(v bool)`

SetCanInvite sets CanInvite field to given value.


### GetCreatedAt

`func (o *VacanciesMatchVacancyFields) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesMatchVacancyFields) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesMatchVacancyFields) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmployerNegotiationsState

`func (o *VacanciesMatchVacancyFields) GetEmployerNegotiationsState() IncludesIdName`

GetEmployerNegotiationsState returns the EmployerNegotiationsState field if non-nil, zero value otherwise.

### GetEmployerNegotiationsStateOk

`func (o *VacanciesMatchVacancyFields) GetEmployerNegotiationsStateOk() (*IncludesIdName, bool)`

GetEmployerNegotiationsStateOk returns a tuple with the EmployerNegotiationsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerNegotiationsState

`func (o *VacanciesMatchVacancyFields) SetEmployerNegotiationsState(v IncludesIdName)`

SetEmployerNegotiationsState sets EmployerNegotiationsState field to given value.


### SetEmployerNegotiationsStateNil

`func (o *VacanciesMatchVacancyFields) SetEmployerNegotiationsStateNil(b bool)`

 SetEmployerNegotiationsStateNil sets the value for EmployerNegotiationsState to be an explicit nil

### UnsetEmployerNegotiationsState
`func (o *VacanciesMatchVacancyFields) UnsetEmployerNegotiationsState()`

UnsetEmployerNegotiationsState ensures that no value is present for EmployerNegotiationsState, not even an explicit nil
### GetManager

`func (o *VacanciesMatchVacancyFields) GetManager() VacancyManagerOutput`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacanciesMatchVacancyFields) GetManagerOk() (*VacancyManagerOutput, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacanciesMatchVacancyFields) SetManager(v VacancyManagerOutput)`

SetManager sets Manager field to given value.


### GetNegotiationsActions

`func (o *VacanciesMatchVacancyFields) GetNegotiationsActions() []VacancyNegotiationActions`

GetNegotiationsActions returns the NegotiationsActions field if non-nil, zero value otherwise.

### GetNegotiationsActionsOk

`func (o *VacanciesMatchVacancyFields) GetNegotiationsActionsOk() (*[]VacancyNegotiationActions, bool)`

GetNegotiationsActionsOk returns a tuple with the NegotiationsActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsActions

`func (o *VacanciesMatchVacancyFields) SetNegotiationsActions(v []VacancyNegotiationActions)`

SetNegotiationsActions sets NegotiationsActions field to given value.


### GetNegotiationsState

`func (o *VacanciesMatchVacancyFields) GetNegotiationsState() IncludesIdName`

GetNegotiationsState returns the NegotiationsState field if non-nil, zero value otherwise.

### GetNegotiationsStateOk

`func (o *VacanciesMatchVacancyFields) GetNegotiationsStateOk() (*IncludesIdName, bool)`

GetNegotiationsStateOk returns a tuple with the NegotiationsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsState

`func (o *VacanciesMatchVacancyFields) SetNegotiationsState(v IncludesIdName)`

SetNegotiationsState sets NegotiationsState field to given value.


### SetNegotiationsStateNil

`func (o *VacanciesMatchVacancyFields) SetNegotiationsStateNil(b bool)`

 SetNegotiationsStateNil sets the value for NegotiationsState to be an explicit nil

### UnsetNegotiationsState
`func (o *VacanciesMatchVacancyFields) UnsetNegotiationsState()`

UnsetNegotiationsState ensures that no value is present for NegotiationsState, not even an explicit nil
### GetSortPointDistance

`func (o *VacanciesMatchVacancyFields) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesMatchVacancyFields) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesMatchVacancyFields) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesMatchVacancyFields) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesMatchVacancyFields) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesMatchVacancyFields) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
### GetTemplates

`func (o *VacanciesMatchVacancyFields) GetTemplates() []VacancyTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *VacanciesMatchVacancyFields) GetTemplatesOk() (*[]VacancyTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *VacanciesMatchVacancyFields) SetTemplates(v []VacancyTemplates)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *VacanciesMatchVacancyFields) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


