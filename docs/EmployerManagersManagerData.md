# EmployerManagersManagerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalPhone** | Pointer to [**EmployerManagersEmployerManagerItemAdditionalPhone**](EmployerManagersEmployerManagerItemAdditionalPhone.md) |  | [optional] 
**Permissions** | Pointer to [**[]EmployerManagerTypesAvailablePermissions**](EmployerManagerTypesAvailablePermissions.md) | Список прав, которые можно дать данному типу менеджера | [optional] 
**Phone** | Pointer to [**EmployerManagersEmployerManagerItemPhone**](EmployerManagersEmployerManagerItemPhone.md) |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 

## Methods

### NewEmployerManagersManagerData

`func NewEmployerManagersManagerData() *EmployerManagersManagerData`

NewEmployerManagersManagerData instantiates a new EmployerManagersManagerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerManagersManagerDataWithDefaults

`func NewEmployerManagersManagerDataWithDefaults() *EmployerManagersManagerData`

NewEmployerManagersManagerDataWithDefaults instantiates a new EmployerManagersManagerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalPhone

`func (o *EmployerManagersManagerData) GetAdditionalPhone() EmployerManagersEmployerManagerItemAdditionalPhone`

GetAdditionalPhone returns the AdditionalPhone field if non-nil, zero value otherwise.

### GetAdditionalPhoneOk

`func (o *EmployerManagersManagerData) GetAdditionalPhoneOk() (*EmployerManagersEmployerManagerItemAdditionalPhone, bool)`

GetAdditionalPhoneOk returns a tuple with the AdditionalPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPhone

`func (o *EmployerManagersManagerData) SetAdditionalPhone(v EmployerManagersEmployerManagerItemAdditionalPhone)`

SetAdditionalPhone sets AdditionalPhone field to given value.

### HasAdditionalPhone

`func (o *EmployerManagersManagerData) HasAdditionalPhone() bool`

HasAdditionalPhone returns a boolean if a field has been set.

### GetPermissions

`func (o *EmployerManagersManagerData) GetPermissions() []EmployerManagerTypesAvailablePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EmployerManagersManagerData) GetPermissionsOk() (*[]EmployerManagerTypesAvailablePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EmployerManagersManagerData) SetPermissions(v []EmployerManagerTypesAvailablePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EmployerManagersManagerData) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhone

`func (o *EmployerManagersManagerData) GetPhone() EmployerManagersEmployerManagerItemPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EmployerManagersManagerData) GetPhoneOk() (*EmployerManagersEmployerManagerItemPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *EmployerManagersManagerData) SetPhone(v EmployerManagersEmployerManagerItemPhone)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *EmployerManagersManagerData) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPosition

`func (o *EmployerManagersManagerData) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EmployerManagersManagerData) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EmployerManagersManagerData) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *EmployerManagersManagerData) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


