# VacanciesVacancyUpgradeListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]VacanciesUpgradeFieldsAction**](VacanciesUpgradeFieldsAction.md) | Список возможных действий | 
**VacancyBillingType** | [**VacanciesUpgradeFieldsBillingTypeFull**](VacanciesUpgradeFieldsBillingTypeFull.md) |  | 
**WithoutAction** | Pointer to [**[]VacanciesVacancyUpgradeListItemWithoutActionInner**](VacanciesVacancyUpgradeListItemWithoutActionInner.md) | Объект с описанием причины, по которой невозможно улучшить вакансию до данного типа. &#x60;Null&#x60;, если массив &#x60;actions&#x60; не пустой | [optional] 

## Methods

### NewVacanciesVacancyUpgradeListItem

`func NewVacanciesVacancyUpgradeListItem(actions []VacanciesUpgradeFieldsAction, vacancyBillingType VacanciesUpgradeFieldsBillingTypeFull, ) *VacanciesVacancyUpgradeListItem`

NewVacanciesVacancyUpgradeListItem instantiates a new VacanciesVacancyUpgradeListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyUpgradeListItemWithDefaults

`func NewVacanciesVacancyUpgradeListItemWithDefaults() *VacanciesVacancyUpgradeListItem`

NewVacanciesVacancyUpgradeListItemWithDefaults instantiates a new VacanciesVacancyUpgradeListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *VacanciesVacancyUpgradeListItem) GetActions() []VacanciesUpgradeFieldsAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *VacanciesVacancyUpgradeListItem) GetActionsOk() (*[]VacanciesUpgradeFieldsAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *VacanciesVacancyUpgradeListItem) SetActions(v []VacanciesUpgradeFieldsAction)`

SetActions sets Actions field to given value.


### GetVacancyBillingType

`func (o *VacanciesVacancyUpgradeListItem) GetVacancyBillingType() VacanciesUpgradeFieldsBillingTypeFull`

GetVacancyBillingType returns the VacancyBillingType field if non-nil, zero value otherwise.

### GetVacancyBillingTypeOk

`func (o *VacanciesVacancyUpgradeListItem) GetVacancyBillingTypeOk() (*VacanciesUpgradeFieldsBillingTypeFull, bool)`

GetVacancyBillingTypeOk returns a tuple with the VacancyBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyBillingType

`func (o *VacanciesVacancyUpgradeListItem) SetVacancyBillingType(v VacanciesUpgradeFieldsBillingTypeFull)`

SetVacancyBillingType sets VacancyBillingType field to given value.


### GetWithoutAction

`func (o *VacanciesVacancyUpgradeListItem) GetWithoutAction() []VacanciesVacancyUpgradeListItemWithoutActionInner`

GetWithoutAction returns the WithoutAction field if non-nil, zero value otherwise.

### GetWithoutActionOk

`func (o *VacanciesVacancyUpgradeListItem) GetWithoutActionOk() (*[]VacanciesVacancyUpgradeListItemWithoutActionInner, bool)`

GetWithoutActionOk returns a tuple with the WithoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutAction

`func (o *VacanciesVacancyUpgradeListItem) SetWithoutAction(v []VacanciesVacancyUpgradeListItemWithoutActionInner)`

SetWithoutAction sets WithoutAction field to given value.

### HasWithoutAction

`func (o *VacanciesVacancyUpgradeListItem) HasWithoutAction() bool`

HasWithoutAction returns a boolean if a field has been set.

### SetWithoutActionNil

`func (o *VacanciesVacancyUpgradeListItem) SetWithoutActionNil(b bool)`

 SetWithoutActionNil sets the value for WithoutAction to be an explicit nil

### UnsetWithoutAction
`func (o *VacanciesVacancyUpgradeListItem) UnsetWithoutAction()`

UnsetWithoutAction ensures that no value is present for WithoutAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


