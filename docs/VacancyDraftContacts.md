# VacancyDraftContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Email | [optional] 
**Name** | Pointer to **NullableString** | Имя менеджера | [optional] 
**Phones** | Pointer to [**[]VacancyDraftPhoneItem**](VacancyDraftPhoneItem.md) | Список телефонов для связи | [optional] 

## Methods

### NewVacancyDraftContacts

`func NewVacancyDraftContacts() *VacancyDraftContacts`

NewVacancyDraftContacts instantiates a new VacancyDraftContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftContactsWithDefaults

`func NewVacancyDraftContactsWithDefaults() *VacancyDraftContacts`

NewVacancyDraftContactsWithDefaults instantiates a new VacancyDraftContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *VacancyDraftContacts) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VacancyDraftContacts) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VacancyDraftContacts) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VacancyDraftContacts) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *VacancyDraftContacts) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *VacancyDraftContacts) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *VacancyDraftContacts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyDraftContacts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyDraftContacts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyDraftContacts) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VacancyDraftContacts) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VacancyDraftContacts) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhones

`func (o *VacancyDraftContacts) GetPhones() []VacancyDraftPhoneItem`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *VacancyDraftContacts) GetPhonesOk() (*[]VacancyDraftPhoneItem, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *VacancyDraftContacts) SetPhones(v []VacancyDraftPhoneItem)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *VacancyDraftContacts) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### SetPhonesNil

`func (o *VacancyDraftContacts) SetPhonesNil(b bool)`

 SetPhonesNil sets the value for Phones to be an explicit nil

### UnsetPhones
`func (o *VacancyDraftContacts) UnsetPhones()`

UnsetPhones ensures that no value is present for Phones, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


