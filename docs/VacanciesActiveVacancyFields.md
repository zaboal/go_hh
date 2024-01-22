# VacanciesActiveVacancyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**NullableVacanciesAddress**](VacanciesAddress.md) |  | [optional] 
**CanUpgradeBillingType** | **bool** | Можно ли улучшить биллинговый тип вакансии | 
**Counters** | [**VacancyCountersForActive**](VacancyCountersForActive.md) |  | 
**CreatedAt** | **string** | Дата и время публикации вакансии | 
**ExpiresAt** | **string** | Дата и время окончания публикации вакансии | 
**HasUpdates** | **bool** | Есть ли в откликах/приглашениях по данной вакансии обновления, требующие внимания | 
**Manager** | [**VacancyManagerOutput**](VacancyManagerOutput.md) |  | 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций. Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60;  | [optional] 

## Methods

### NewVacanciesActiveVacancyFields

`func NewVacanciesActiveVacancyFields(canUpgradeBillingType bool, counters VacancyCountersForActive, createdAt string, expiresAt string, hasUpdates bool, manager VacancyManagerOutput, ) *VacanciesActiveVacancyFields`

NewVacanciesActiveVacancyFields instantiates a new VacanciesActiveVacancyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesActiveVacancyFieldsWithDefaults

`func NewVacanciesActiveVacancyFieldsWithDefaults() *VacanciesActiveVacancyFields`

NewVacanciesActiveVacancyFieldsWithDefaults instantiates a new VacanciesActiveVacancyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VacanciesActiveVacancyFields) GetAddress() VacanciesAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesActiveVacancyFields) GetAddressOk() (*VacanciesAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesActiveVacancyFields) SetAddress(v VacanciesAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesActiveVacancyFields) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacanciesActiveVacancyFields) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacanciesActiveVacancyFields) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCanUpgradeBillingType

`func (o *VacanciesActiveVacancyFields) GetCanUpgradeBillingType() bool`

GetCanUpgradeBillingType returns the CanUpgradeBillingType field if non-nil, zero value otherwise.

### GetCanUpgradeBillingTypeOk

`func (o *VacanciesActiveVacancyFields) GetCanUpgradeBillingTypeOk() (*bool, bool)`

GetCanUpgradeBillingTypeOk returns a tuple with the CanUpgradeBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgradeBillingType

`func (o *VacanciesActiveVacancyFields) SetCanUpgradeBillingType(v bool)`

SetCanUpgradeBillingType sets CanUpgradeBillingType field to given value.


### GetCounters

`func (o *VacanciesActiveVacancyFields) GetCounters() VacancyCountersForActive`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *VacanciesActiveVacancyFields) GetCountersOk() (*VacancyCountersForActive, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *VacanciesActiveVacancyFields) SetCounters(v VacancyCountersForActive)`

SetCounters sets Counters field to given value.


### GetCreatedAt

`func (o *VacanciesActiveVacancyFields) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesActiveVacancyFields) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesActiveVacancyFields) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *VacanciesActiveVacancyFields) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VacanciesActiveVacancyFields) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VacanciesActiveVacancyFields) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetHasUpdates

`func (o *VacanciesActiveVacancyFields) GetHasUpdates() bool`

GetHasUpdates returns the HasUpdates field if non-nil, zero value otherwise.

### GetHasUpdatesOk

`func (o *VacanciesActiveVacancyFields) GetHasUpdatesOk() (*bool, bool)`

GetHasUpdatesOk returns a tuple with the HasUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdates

`func (o *VacanciesActiveVacancyFields) SetHasUpdates(v bool)`

SetHasUpdates sets HasUpdates field to given value.


### GetManager

`func (o *VacanciesActiveVacancyFields) GetManager() VacancyManagerOutput`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacanciesActiveVacancyFields) GetManagerOk() (*VacancyManagerOutput, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacanciesActiveVacancyFields) SetManager(v VacancyManagerOutput)`

SetManager sets Manager field to given value.


### GetSortPointDistance

`func (o *VacanciesActiveVacancyFields) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesActiveVacancyFields) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesActiveVacancyFields) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesActiveVacancyFields) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesActiveVacancyFields) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesActiveVacancyFields) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


