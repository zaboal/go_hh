# MeEmployerProfileManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasAdminRights** | **bool** | Обладает ли текущий менеджер правами администратора | 
**HasMultipleManagerAccounts** | **bool** | Существует ли у пользователя несколько [рабочих аккаунтов](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts) | 
**Id** | **string** | Идентификатор менеджера | 
**IsMainContactPerson** | **bool** | Является ли текущий менеджер главным контактным лицом компании | 
**ManagerSettingsUrl** | **string** | URL, на который нужно сделать GET запрос, чтобы получить [предпочтения менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-settings) | 

## Methods

### NewMeEmployerProfileManager

`func NewMeEmployerProfileManager(hasAdminRights bool, hasMultipleManagerAccounts bool, id string, isMainContactPerson bool, managerSettingsUrl string, ) *MeEmployerProfileManager`

NewMeEmployerProfileManager instantiates a new MeEmployerProfileManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeEmployerProfileManagerWithDefaults

`func NewMeEmployerProfileManagerWithDefaults() *MeEmployerProfileManager`

NewMeEmployerProfileManagerWithDefaults instantiates a new MeEmployerProfileManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasAdminRights

`func (o *MeEmployerProfileManager) GetHasAdminRights() bool`

GetHasAdminRights returns the HasAdminRights field if non-nil, zero value otherwise.

### GetHasAdminRightsOk

`func (o *MeEmployerProfileManager) GetHasAdminRightsOk() (*bool, bool)`

GetHasAdminRightsOk returns a tuple with the HasAdminRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAdminRights

`func (o *MeEmployerProfileManager) SetHasAdminRights(v bool)`

SetHasAdminRights sets HasAdminRights field to given value.


### GetHasMultipleManagerAccounts

`func (o *MeEmployerProfileManager) GetHasMultipleManagerAccounts() bool`

GetHasMultipleManagerAccounts returns the HasMultipleManagerAccounts field if non-nil, zero value otherwise.

### GetHasMultipleManagerAccountsOk

`func (o *MeEmployerProfileManager) GetHasMultipleManagerAccountsOk() (*bool, bool)`

GetHasMultipleManagerAccountsOk returns a tuple with the HasMultipleManagerAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMultipleManagerAccounts

`func (o *MeEmployerProfileManager) SetHasMultipleManagerAccounts(v bool)`

SetHasMultipleManagerAccounts sets HasMultipleManagerAccounts field to given value.


### GetId

`func (o *MeEmployerProfileManager) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeEmployerProfileManager) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeEmployerProfileManager) SetId(v string)`

SetId sets Id field to given value.


### GetIsMainContactPerson

`func (o *MeEmployerProfileManager) GetIsMainContactPerson() bool`

GetIsMainContactPerson returns the IsMainContactPerson field if non-nil, zero value otherwise.

### GetIsMainContactPersonOk

`func (o *MeEmployerProfileManager) GetIsMainContactPersonOk() (*bool, bool)`

GetIsMainContactPersonOk returns a tuple with the IsMainContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMainContactPerson

`func (o *MeEmployerProfileManager) SetIsMainContactPerson(v bool)`

SetIsMainContactPerson sets IsMainContactPerson field to given value.


### GetManagerSettingsUrl

`func (o *MeEmployerProfileManager) GetManagerSettingsUrl() string`

GetManagerSettingsUrl returns the ManagerSettingsUrl field if non-nil, zero value otherwise.

### GetManagerSettingsUrlOk

`func (o *MeEmployerProfileManager) GetManagerSettingsUrlOk() (*string, bool)`

GetManagerSettingsUrlOk returns a tuple with the ManagerSettingsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerSettingsUrl

`func (o *MeEmployerProfileManager) SetManagerSettingsUrl(v string)`

SetManagerSettingsUrl sets ManagerSettingsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


