# VacancyContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Электронная почта. Значение поля должно соответствовать формату email | [optional] 
**Name** | **string** | Имя контакта | 
**Phones** | Pointer to [**[]VacancyPhoneItem**](VacancyPhoneItem.md) | Список телефонов для связи | [optional] 

## Methods

### NewVacancyContacts

`func NewVacancyContacts(name string, ) *VacancyContacts`

NewVacancyContacts instantiates a new VacancyContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyContactsWithDefaults

`func NewVacancyContactsWithDefaults() *VacancyContacts`

NewVacancyContactsWithDefaults instantiates a new VacancyContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *VacancyContacts) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VacancyContacts) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VacancyContacts) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VacancyContacts) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *VacancyContacts) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *VacancyContacts) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *VacancyContacts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyContacts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyContacts) SetName(v string)`

SetName sets Name field to given value.


### GetPhones

`func (o *VacancyContacts) GetPhones() []VacancyPhoneItem`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *VacancyContacts) GetPhonesOk() (*[]VacancyPhoneItem, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *VacancyContacts) SetPhones(v []VacancyPhoneItem)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *VacancyContacts) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### SetPhonesNil

`func (o *VacancyContacts) SetPhonesNil(b bool)`

 SetPhonesNil sets the value for Phones to be an explicit nil

### UnsetPhones
`func (o *VacancyContacts) UnsetPhones()`

UnsetPhones ensures that no value is present for Phones, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


