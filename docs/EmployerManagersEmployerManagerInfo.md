# EmployerManagersEmployerManagerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalPhone** | Pointer to [**EmployerManagersEmployerManagerInfoAdditionalPhone**](EmployerManagersEmployerManagerInfoAdditionalPhone.md) |  | [optional] 
**Area** | Pointer to [**EmployerManagersArea**](EmployerManagersArea.md) |  | [optional] 
**Email** | **string** | Адрес электронной почты менеджера | 
**FirstName** | **string** | Имя менеджера | 
**FullName** | Pointer to **string** | Полное имя менеджера | [optional] 
**Id** | **string** | Идентификатор менеджера | 
**IsMainContactPerson** | **bool** | Является ли менеджер основным контактным лицом | 
**LastName** | **string** | Фамилия менеджера | 
**ManagerType** | Pointer to [**EmployerManagersManagerType**](EmployerManagersManagerType.md) |  | [optional] 
**MiddleName** | Pointer to **string** | Отчество менеджера | [optional] 
**Name** | Pointer to **string** | Полное имя менеджера | [optional] 
**Permissions** | [**[]EmployerManagerTypesAvailablePermissions**](EmployerManagerTypesAvailablePermissions.md) | Список [прав менеджера](#tag/Menedzhery-rabotodatelya/operation/get-employer-manager-types) | 
**Phone** | [**EmployerManagersEmployerManagerInfoPhone**](EmployerManagersEmployerManagerInfoPhone.md) |  | 
**Position** | **string** | Должность менеджера | 
**VacanciesCount** | Pointer to **NullableFloat32** | Количество опубликованных (активных) вакансий у данного менеджера. &#x60;null&#x60; — если у пользователя нет прав на просмотр вакансий этого менеджера | [optional] 

## Methods

### NewEmployerManagersEmployerManagerInfo

`func NewEmployerManagersEmployerManagerInfo(email string, firstName string, id string, isMainContactPerson bool, lastName string, permissions []EmployerManagerTypesAvailablePermissions, phone EmployerManagersEmployerManagerInfoPhone, position string, ) *EmployerManagersEmployerManagerInfo`

NewEmployerManagersEmployerManagerInfo instantiates a new EmployerManagersEmployerManagerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerManagersEmployerManagerInfoWithDefaults

`func NewEmployerManagersEmployerManagerInfoWithDefaults() *EmployerManagersEmployerManagerInfo`

NewEmployerManagersEmployerManagerInfoWithDefaults instantiates a new EmployerManagersEmployerManagerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalPhone

`func (o *EmployerManagersEmployerManagerInfo) GetAdditionalPhone() EmployerManagersEmployerManagerInfoAdditionalPhone`

GetAdditionalPhone returns the AdditionalPhone field if non-nil, zero value otherwise.

### GetAdditionalPhoneOk

`func (o *EmployerManagersEmployerManagerInfo) GetAdditionalPhoneOk() (*EmployerManagersEmployerManagerInfoAdditionalPhone, bool)`

GetAdditionalPhoneOk returns a tuple with the AdditionalPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPhone

`func (o *EmployerManagersEmployerManagerInfo) SetAdditionalPhone(v EmployerManagersEmployerManagerInfoAdditionalPhone)`

SetAdditionalPhone sets AdditionalPhone field to given value.

### HasAdditionalPhone

`func (o *EmployerManagersEmployerManagerInfo) HasAdditionalPhone() bool`

HasAdditionalPhone returns a boolean if a field has been set.

### GetArea

`func (o *EmployerManagersEmployerManagerInfo) GetArea() EmployerManagersArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *EmployerManagersEmployerManagerInfo) GetAreaOk() (*EmployerManagersArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *EmployerManagersEmployerManagerInfo) SetArea(v EmployerManagersArea)`

SetArea sets Area field to given value.

### HasArea

`func (o *EmployerManagersEmployerManagerInfo) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetEmail

`func (o *EmployerManagersEmployerManagerInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmployerManagersEmployerManagerInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmployerManagersEmployerManagerInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *EmployerManagersEmployerManagerInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployerManagersEmployerManagerInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployerManagersEmployerManagerInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetFullName

`func (o *EmployerManagersEmployerManagerInfo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *EmployerManagersEmployerManagerInfo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *EmployerManagersEmployerManagerInfo) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *EmployerManagersEmployerManagerInfo) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetId

`func (o *EmployerManagersEmployerManagerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployerManagersEmployerManagerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployerManagersEmployerManagerInfo) SetId(v string)`

SetId sets Id field to given value.


### GetIsMainContactPerson

`func (o *EmployerManagersEmployerManagerInfo) GetIsMainContactPerson() bool`

GetIsMainContactPerson returns the IsMainContactPerson field if non-nil, zero value otherwise.

### GetIsMainContactPersonOk

`func (o *EmployerManagersEmployerManagerInfo) GetIsMainContactPersonOk() (*bool, bool)`

GetIsMainContactPersonOk returns a tuple with the IsMainContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMainContactPerson

`func (o *EmployerManagersEmployerManagerInfo) SetIsMainContactPerson(v bool)`

SetIsMainContactPerson sets IsMainContactPerson field to given value.


### GetLastName

`func (o *EmployerManagersEmployerManagerInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployerManagersEmployerManagerInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployerManagersEmployerManagerInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetManagerType

`func (o *EmployerManagersEmployerManagerInfo) GetManagerType() EmployerManagersManagerType`

GetManagerType returns the ManagerType field if non-nil, zero value otherwise.

### GetManagerTypeOk

`func (o *EmployerManagersEmployerManagerInfo) GetManagerTypeOk() (*EmployerManagersManagerType, bool)`

GetManagerTypeOk returns a tuple with the ManagerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerType

`func (o *EmployerManagersEmployerManagerInfo) SetManagerType(v EmployerManagersManagerType)`

SetManagerType sets ManagerType field to given value.

### HasManagerType

`func (o *EmployerManagersEmployerManagerInfo) HasManagerType() bool`

HasManagerType returns a boolean if a field has been set.

### GetMiddleName

`func (o *EmployerManagersEmployerManagerInfo) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *EmployerManagersEmployerManagerInfo) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *EmployerManagersEmployerManagerInfo) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *EmployerManagersEmployerManagerInfo) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetName

`func (o *EmployerManagersEmployerManagerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployerManagersEmployerManagerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployerManagersEmployerManagerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmployerManagersEmployerManagerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *EmployerManagersEmployerManagerInfo) GetPermissions() []EmployerManagerTypesAvailablePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EmployerManagersEmployerManagerInfo) GetPermissionsOk() (*[]EmployerManagerTypesAvailablePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EmployerManagersEmployerManagerInfo) SetPermissions(v []EmployerManagerTypesAvailablePermissions)`

SetPermissions sets Permissions field to given value.


### GetPhone

`func (o *EmployerManagersEmployerManagerInfo) GetPhone() EmployerManagersEmployerManagerInfoPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EmployerManagersEmployerManagerInfo) GetPhoneOk() (*EmployerManagersEmployerManagerInfoPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *EmployerManagersEmployerManagerInfo) SetPhone(v EmployerManagersEmployerManagerInfoPhone)`

SetPhone sets Phone field to given value.


### GetPosition

`func (o *EmployerManagersEmployerManagerInfo) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EmployerManagersEmployerManagerInfo) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EmployerManagersEmployerManagerInfo) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetVacanciesCount

`func (o *EmployerManagersEmployerManagerInfo) GetVacanciesCount() float32`

GetVacanciesCount returns the VacanciesCount field if non-nil, zero value otherwise.

### GetVacanciesCountOk

`func (o *EmployerManagersEmployerManagerInfo) GetVacanciesCountOk() (*float32, bool)`

GetVacanciesCountOk returns a tuple with the VacanciesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacanciesCount

`func (o *EmployerManagersEmployerManagerInfo) SetVacanciesCount(v float32)`

SetVacanciesCount sets VacanciesCount field to given value.

### HasVacanciesCount

`func (o *EmployerManagersEmployerManagerInfo) HasVacanciesCount() bool`

HasVacanciesCount returns a boolean if a field has been set.

### SetVacanciesCountNil

`func (o *EmployerManagersEmployerManagerInfo) SetVacanciesCountNil(b bool)`

 SetVacanciesCountNil sets the value for VacanciesCount to be an explicit nil

### UnsetVacanciesCount
`func (o *EmployerManagersEmployerManagerInfo) UnsetVacanciesCount()`

UnsetVacanciesCount ensures that no value is present for VacanciesCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


