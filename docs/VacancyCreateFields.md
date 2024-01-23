# VacancyCreateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | [**VacancyArea**](VacancyArea.md) |  | 
**BillingType** | [**VacancyBillingType**](VacancyBillingType.md) |  | 
**Description** | **string** | Описание в html, не менее 200 символов | 
**DriverLicenseTypes** | Pointer to [**[]VacancyDriverLicenceTypeItem**](VacancyDriverLicenceTypeItem.md) | Список требуемых категорий водительских прав | [optional] 
**Manager** | Pointer to [**NullableVacancyManager**](VacancyManager.md) |  | [optional] 
**Name** | **string** | Название | 
**PreviousId** | Pointer to **NullableString** | Если этот параметр передан, то у новой вакансии дополнительно будет создана связь с предыдущей вакансией (поле previous_id). Этот параметр не влияет на другие и не связан с ними, их всё равно необходимо передавать.  Должен быть равен только ID архивной вакансии. ID архивной вакансии можно получить, запросив [список архивных вакансий](#tag/Upravlenie-vakansiyami/operation/get-archived-vacancies) &lt;a name&#x3D;&#39;previous_id&#39;&gt;&lt;/a&gt;  | [optional] 
**Type** | [**VacancyType**](VacancyType.md) |  | 

## Methods

### NewVacancyCreateFields

`func NewVacancyCreateFields(area VacancyArea, billingType VacancyBillingType, description string, name string, type_ VacancyType, ) *VacancyCreateFields`

NewVacancyCreateFields instantiates a new VacancyCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyCreateFieldsWithDefaults

`func NewVacancyCreateFieldsWithDefaults() *VacancyCreateFields`

NewVacancyCreateFieldsWithDefaults instantiates a new VacancyCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *VacancyCreateFields) GetArea() VacancyArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacancyCreateFields) GetAreaOk() (*VacancyArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacancyCreateFields) SetArea(v VacancyArea)`

SetArea sets Area field to given value.


### GetBillingType

`func (o *VacancyCreateFields) GetBillingType() VacancyBillingType`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacancyCreateFields) GetBillingTypeOk() (*VacancyBillingType, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacancyCreateFields) SetBillingType(v VacancyBillingType)`

SetBillingType sets BillingType field to given value.


### GetDescription

`func (o *VacancyCreateFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyCreateFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyCreateFields) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDriverLicenseTypes

`func (o *VacancyCreateFields) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacancyCreateFields) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacancyCreateFields) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *VacancyCreateFields) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### GetManager

`func (o *VacancyCreateFields) GetManager() VacancyManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacancyCreateFields) GetManagerOk() (*VacancyManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacancyCreateFields) SetManager(v VacancyManager)`

SetManager sets Manager field to given value.

### HasManager

`func (o *VacancyCreateFields) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *VacancyCreateFields) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *VacancyCreateFields) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetName

`func (o *VacancyCreateFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyCreateFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyCreateFields) SetName(v string)`

SetName sets Name field to given value.


### GetPreviousId

`func (o *VacancyCreateFields) GetPreviousId() string`

GetPreviousId returns the PreviousId field if non-nil, zero value otherwise.

### GetPreviousIdOk

`func (o *VacancyCreateFields) GetPreviousIdOk() (*string, bool)`

GetPreviousIdOk returns a tuple with the PreviousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousId

`func (o *VacancyCreateFields) SetPreviousId(v string)`

SetPreviousId sets PreviousId field to given value.

### HasPreviousId

`func (o *VacancyCreateFields) HasPreviousId() bool`

HasPreviousId returns a boolean if a field has been set.

### SetPreviousIdNil

`func (o *VacancyCreateFields) SetPreviousIdNil(b bool)`

 SetPreviousIdNil sets the value for PreviousId to be an explicit nil

### UnsetPreviousId
`func (o *VacancyCreateFields) UnsetPreviousId()`

UnsetPreviousId ensures that no value is present for PreviousId, not even an explicit nil
### GetType

`func (o *VacancyCreateFields) GetType() VacancyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacancyCreateFields) GetTypeOk() (*VacancyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacancyCreateFields) SetType(v VacancyType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


