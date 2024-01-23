# VacanciesUpgradeFieldsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartId** | Pointer to **NullableInt32** | Идентификатор заказа, ожидающего активации. Возвращается только для действий с &#x60;actions.type&#x3D;activate&#x60; | [optional] 
**Price** | Pointer to [**NullableVacanciesUpgradeFieldsPrice**](VacanciesUpgradeFieldsPrice.md) |  | [optional] 
**Type** | **string** | Тип действия:  * &#x60;direct_upgrade&#x60; — публикации вакансий данного типа есть на счету. Вы можете изменить тип вакансии. * &#x60;activate&#x60; — публикации вакансий данного типа есть в неактивированных заказах. Перейдите по ссылке, указанной в поле &#x60;actions.url&#x60;, и активируйте заказ. После этого станет доступно улучшение вакансии. * &#x60;buy&#x60; — нет доступных публикаций вакансий данного типа. Перейдите по ссылке, указанной в поле &#x60;actions.url&#x60;, чтобы перейти к покупке публикаций нужного типа  | 
**Url** | Pointer to **NullableString** | Ссылка на действие | [optional] 

## Methods

### NewVacanciesUpgradeFieldsAction

`func NewVacanciesUpgradeFieldsAction(type_ string, ) *VacanciesUpgradeFieldsAction`

NewVacanciesUpgradeFieldsAction instantiates a new VacanciesUpgradeFieldsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesUpgradeFieldsActionWithDefaults

`func NewVacanciesUpgradeFieldsActionWithDefaults() *VacanciesUpgradeFieldsAction`

NewVacanciesUpgradeFieldsActionWithDefaults instantiates a new VacanciesUpgradeFieldsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartId

`func (o *VacanciesUpgradeFieldsAction) GetCartId() int32`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *VacanciesUpgradeFieldsAction) GetCartIdOk() (*int32, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *VacanciesUpgradeFieldsAction) SetCartId(v int32)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *VacanciesUpgradeFieldsAction) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *VacanciesUpgradeFieldsAction) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *VacanciesUpgradeFieldsAction) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetPrice

`func (o *VacanciesUpgradeFieldsAction) GetPrice() VacanciesUpgradeFieldsPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *VacanciesUpgradeFieldsAction) GetPriceOk() (*VacanciesUpgradeFieldsPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *VacanciesUpgradeFieldsAction) SetPrice(v VacanciesUpgradeFieldsPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *VacanciesUpgradeFieldsAction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *VacanciesUpgradeFieldsAction) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *VacanciesUpgradeFieldsAction) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetType

`func (o *VacanciesUpgradeFieldsAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesUpgradeFieldsAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesUpgradeFieldsAction) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *VacanciesUpgradeFieldsAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesUpgradeFieldsAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesUpgradeFieldsAction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VacanciesUpgradeFieldsAction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VacanciesUpgradeFieldsAction) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VacanciesUpgradeFieldsAction) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


