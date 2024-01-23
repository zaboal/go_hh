# EmployersEmployersBlacklisted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]EmployersEmployersBlacklistedItem**](EmployersEmployersBlacklistedItem.md) | Массив скрытых работодателей | [optional] 
**LimitReached** | Pointer to **bool** | Превышено ли максимальное количество элементов в списке | [optional] 

## Methods

### NewEmployersEmployersBlacklisted

`func NewEmployersEmployersBlacklisted() *EmployersEmployersBlacklisted`

NewEmployersEmployersBlacklisted instantiates a new EmployersEmployersBlacklisted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersEmployersBlacklistedWithDefaults

`func NewEmployersEmployersBlacklistedWithDefaults() *EmployersEmployersBlacklisted`

NewEmployersEmployersBlacklistedWithDefaults instantiates a new EmployersEmployersBlacklisted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EmployersEmployersBlacklisted) GetItems() []EmployersEmployersBlacklistedItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EmployersEmployersBlacklisted) GetItemsOk() (*[]EmployersEmployersBlacklistedItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EmployersEmployersBlacklisted) SetItems(v []EmployersEmployersBlacklistedItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *EmployersEmployersBlacklisted) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimitReached

`func (o *EmployersEmployersBlacklisted) GetLimitReached() bool`

GetLimitReached returns the LimitReached field if non-nil, zero value otherwise.

### GetLimitReachedOk

`func (o *EmployersEmployersBlacklisted) GetLimitReachedOk() (*bool, bool)`

GetLimitReachedOk returns a tuple with the LimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitReached

`func (o *EmployersEmployersBlacklisted) SetLimitReached(v bool)`

SetLimitReached sets LimitReached field to given value.

### HasLimitReached

`func (o *EmployersEmployersBlacklisted) HasLimitReached() bool`

HasLimitReached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


