# EmployerManagersAddEmployerManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalPhone** | Pointer to [**EmployerManagersAddEmployerManagerAdditionalPhone**](EmployerManagersAddEmployerManagerAdditionalPhone.md) |  | [optional] 
**Area** | [**EmployerManagersAreaId**](EmployerManagersAreaId.md) |  | 
**Email** | **string** | Адрес электронной почты менеджера | 
**FirstName** | **string** | Имя менеджера | 
**IsMainContactPerson** | **bool** | Является ли менеджер основным контактным лицом | 
**LastName** | **string** | Фамилия менеджера | 
**ManagerType** | [**EmployerManagersManagerTypeId**](EmployerManagersManagerTypeId.md) |  | 
**MiddleName** | Pointer to **string** | Отчество менеджера | [optional] 
**Permissions** | Pointer to [**[]EmployerManagersPermissions**](EmployerManagersPermissions.md) | Список [прав менеджера](#tag/Menedzhery-rabotodatelya/operation/get-employer-manager-types) | [optional] 
**Phone** | [**EmployerManagersAddEmployerManagerPhone**](EmployerManagersAddEmployerManagerPhone.md) |  | 
**Position** | **string** | Должность менеджера | 

## Methods

### NewEmployerManagersAddEmployerManager

`func NewEmployerManagersAddEmployerManager(area EmployerManagersAreaId, email string, firstName string, isMainContactPerson bool, lastName string, managerType EmployerManagersManagerTypeId, phone EmployerManagersAddEmployerManagerPhone, position string, ) *EmployerManagersAddEmployerManager`

NewEmployerManagersAddEmployerManager instantiates a new EmployerManagersAddEmployerManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerManagersAddEmployerManagerWithDefaults

`func NewEmployerManagersAddEmployerManagerWithDefaults() *EmployerManagersAddEmployerManager`

NewEmployerManagersAddEmployerManagerWithDefaults instantiates a new EmployerManagersAddEmployerManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalPhone

`func (o *EmployerManagersAddEmployerManager) GetAdditionalPhone() EmployerManagersAddEmployerManagerAdditionalPhone`

GetAdditionalPhone returns the AdditionalPhone field if non-nil, zero value otherwise.

### GetAdditionalPhoneOk

`func (o *EmployerManagersAddEmployerManager) GetAdditionalPhoneOk() (*EmployerManagersAddEmployerManagerAdditionalPhone, bool)`

GetAdditionalPhoneOk returns a tuple with the AdditionalPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPhone

`func (o *EmployerManagersAddEmployerManager) SetAdditionalPhone(v EmployerManagersAddEmployerManagerAdditionalPhone)`

SetAdditionalPhone sets AdditionalPhone field to given value.

### HasAdditionalPhone

`func (o *EmployerManagersAddEmployerManager) HasAdditionalPhone() bool`

HasAdditionalPhone returns a boolean if a field has been set.

### GetArea

`func (o *EmployerManagersAddEmployerManager) GetArea() EmployerManagersAreaId`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *EmployerManagersAddEmployerManager) GetAreaOk() (*EmployerManagersAreaId, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *EmployerManagersAddEmployerManager) SetArea(v EmployerManagersAreaId)`

SetArea sets Area field to given value.


### GetEmail

`func (o *EmployerManagersAddEmployerManager) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmployerManagersAddEmployerManager) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmployerManagersAddEmployerManager) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *EmployerManagersAddEmployerManager) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployerManagersAddEmployerManager) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployerManagersAddEmployerManager) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetIsMainContactPerson

`func (o *EmployerManagersAddEmployerManager) GetIsMainContactPerson() bool`

GetIsMainContactPerson returns the IsMainContactPerson field if non-nil, zero value otherwise.

### GetIsMainContactPersonOk

`func (o *EmployerManagersAddEmployerManager) GetIsMainContactPersonOk() (*bool, bool)`

GetIsMainContactPersonOk returns a tuple with the IsMainContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMainContactPerson

`func (o *EmployerManagersAddEmployerManager) SetIsMainContactPerson(v bool)`

SetIsMainContactPerson sets IsMainContactPerson field to given value.


### GetLastName

`func (o *EmployerManagersAddEmployerManager) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployerManagersAddEmployerManager) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployerManagersAddEmployerManager) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetManagerType

`func (o *EmployerManagersAddEmployerManager) GetManagerType() EmployerManagersManagerTypeId`

GetManagerType returns the ManagerType field if non-nil, zero value otherwise.

### GetManagerTypeOk

`func (o *EmployerManagersAddEmployerManager) GetManagerTypeOk() (*EmployerManagersManagerTypeId, bool)`

GetManagerTypeOk returns a tuple with the ManagerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerType

`func (o *EmployerManagersAddEmployerManager) SetManagerType(v EmployerManagersManagerTypeId)`

SetManagerType sets ManagerType field to given value.


### GetMiddleName

`func (o *EmployerManagersAddEmployerManager) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *EmployerManagersAddEmployerManager) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *EmployerManagersAddEmployerManager) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *EmployerManagersAddEmployerManager) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPermissions

`func (o *EmployerManagersAddEmployerManager) GetPermissions() []EmployerManagersPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EmployerManagersAddEmployerManager) GetPermissionsOk() (*[]EmployerManagersPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EmployerManagersAddEmployerManager) SetPermissions(v []EmployerManagersPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EmployerManagersAddEmployerManager) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhone

`func (o *EmployerManagersAddEmployerManager) GetPhone() EmployerManagersAddEmployerManagerPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EmployerManagersAddEmployerManager) GetPhoneOk() (*EmployerManagersAddEmployerManagerPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *EmployerManagersAddEmployerManager) SetPhone(v EmployerManagersAddEmployerManagerPhone)`

SetPhone sets Phone field to given value.


### GetPosition

`func (o *EmployerManagersAddEmployerManager) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EmployerManagersAddEmployerManager) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EmployerManagersAddEmployerManager) SetPosition(v string)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


