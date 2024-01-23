# VacancyContactsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallTrackingEnabled** | Pointer to **NullableBool** | Флаг подключения виртуального номера | [optional] 
**Email** | Pointer to **NullableString** | Электронная почта. Значение поля должно соответствовать формату email | [optional] 
**Name** | Pointer to **NullableString** | Имя контакта | [optional] 
**Phones** | Pointer to [**[]VacancyPhoneItem**](VacancyPhoneItem.md) | Список телефонов для связи | [optional] 

## Methods

### NewVacancyContactsOutput

`func NewVacancyContactsOutput() *VacancyContactsOutput`

NewVacancyContactsOutput instantiates a new VacancyContactsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyContactsOutputWithDefaults

`func NewVacancyContactsOutputWithDefaults() *VacancyContactsOutput`

NewVacancyContactsOutputWithDefaults instantiates a new VacancyContactsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallTrackingEnabled

`func (o *VacancyContactsOutput) GetCallTrackingEnabled() bool`

GetCallTrackingEnabled returns the CallTrackingEnabled field if non-nil, zero value otherwise.

### GetCallTrackingEnabledOk

`func (o *VacancyContactsOutput) GetCallTrackingEnabledOk() (*bool, bool)`

GetCallTrackingEnabledOk returns a tuple with the CallTrackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTrackingEnabled

`func (o *VacancyContactsOutput) SetCallTrackingEnabled(v bool)`

SetCallTrackingEnabled sets CallTrackingEnabled field to given value.

### HasCallTrackingEnabled

`func (o *VacancyContactsOutput) HasCallTrackingEnabled() bool`

HasCallTrackingEnabled returns a boolean if a field has been set.

### SetCallTrackingEnabledNil

`func (o *VacancyContactsOutput) SetCallTrackingEnabledNil(b bool)`

 SetCallTrackingEnabledNil sets the value for CallTrackingEnabled to be an explicit nil

### UnsetCallTrackingEnabled
`func (o *VacancyContactsOutput) UnsetCallTrackingEnabled()`

UnsetCallTrackingEnabled ensures that no value is present for CallTrackingEnabled, not even an explicit nil
### GetEmail

`func (o *VacancyContactsOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VacancyContactsOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VacancyContactsOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VacancyContactsOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *VacancyContactsOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *VacancyContactsOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *VacancyContactsOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyContactsOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyContactsOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyContactsOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VacancyContactsOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VacancyContactsOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhones

`func (o *VacancyContactsOutput) GetPhones() []VacancyPhoneItem`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *VacancyContactsOutput) GetPhonesOk() (*[]VacancyPhoneItem, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *VacancyContactsOutput) SetPhones(v []VacancyPhoneItem)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *VacancyContactsOutput) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### SetPhonesNil

`func (o *VacancyContactsOutput) SetPhonesNil(b bool)`

 SetPhonesNil sets the value for Phones to be an explicit nil

### UnsetPhones
`func (o *VacancyContactsOutput) UnsetPhones()`

UnsetPhones ensures that no value is present for Phones, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


