# ManagerAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAccountId** | **string** | Идентификатор текущего рабочего аккаунта менеджера. Совпадает со значением переданного в заголовке &#x60;X-Manager-Account-Id&#x60;  | 
**IsPrimaryAccountBlocked** | **bool** | Заблокирован ли главный рабочий аккаунт менеджера | 
**Items** | [**[]ManagerAccount**](ManagerAccount.md) | Список рабочих аккаунтов менеджера | 
**PrimaryAccountId** | **string** | Идентификатор главного рабочего аккаунта менеджера | 

## Methods

### NewManagerAccounts

`func NewManagerAccounts(currentAccountId string, isPrimaryAccountBlocked bool, items []ManagerAccount, primaryAccountId string, ) *ManagerAccounts`

NewManagerAccounts instantiates a new ManagerAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagerAccountsWithDefaults

`func NewManagerAccountsWithDefaults() *ManagerAccounts`

NewManagerAccountsWithDefaults instantiates a new ManagerAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAccountId

`func (o *ManagerAccounts) GetCurrentAccountId() string`

GetCurrentAccountId returns the CurrentAccountId field if non-nil, zero value otherwise.

### GetCurrentAccountIdOk

`func (o *ManagerAccounts) GetCurrentAccountIdOk() (*string, bool)`

GetCurrentAccountIdOk returns a tuple with the CurrentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAccountId

`func (o *ManagerAccounts) SetCurrentAccountId(v string)`

SetCurrentAccountId sets CurrentAccountId field to given value.


### GetIsPrimaryAccountBlocked

`func (o *ManagerAccounts) GetIsPrimaryAccountBlocked() bool`

GetIsPrimaryAccountBlocked returns the IsPrimaryAccountBlocked field if non-nil, zero value otherwise.

### GetIsPrimaryAccountBlockedOk

`func (o *ManagerAccounts) GetIsPrimaryAccountBlockedOk() (*bool, bool)`

GetIsPrimaryAccountBlockedOk returns a tuple with the IsPrimaryAccountBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryAccountBlocked

`func (o *ManagerAccounts) SetIsPrimaryAccountBlocked(v bool)`

SetIsPrimaryAccountBlocked sets IsPrimaryAccountBlocked field to given value.


### GetItems

`func (o *ManagerAccounts) GetItems() []ManagerAccount`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ManagerAccounts) GetItemsOk() (*[]ManagerAccount, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ManagerAccounts) SetItems(v []ManagerAccount)`

SetItems sets Items field to given value.


### GetPrimaryAccountId

`func (o *ManagerAccounts) GetPrimaryAccountId() string`

GetPrimaryAccountId returns the PrimaryAccountId field if non-nil, zero value otherwise.

### GetPrimaryAccountIdOk

`func (o *ManagerAccounts) GetPrimaryAccountIdOk() (*string, bool)`

GetPrimaryAccountIdOk returns a tuple with the PrimaryAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountId

`func (o *ManagerAccounts) SetPrimaryAccountId(v string)`

SetPrimaryAccountId sets PrimaryAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


