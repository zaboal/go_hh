# VacancyDraftContactsWithFullPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** | Email | 
**Name** | **NullableString** | Имя менеджера | 
**Phones** | [**[]VacancyDraftPhoneItemWithFullPhone**](VacancyDraftPhoneItemWithFullPhone.md) | Список телефонов для связи | 

## Methods

### NewVacancyDraftContactsWithFullPhone

`func NewVacancyDraftContactsWithFullPhone(email NullableString, name NullableString, phones []VacancyDraftPhoneItemWithFullPhone, ) *VacancyDraftContactsWithFullPhone`

NewVacancyDraftContactsWithFullPhone instantiates a new VacancyDraftContactsWithFullPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftContactsWithFullPhoneWithDefaults

`func NewVacancyDraftContactsWithFullPhoneWithDefaults() *VacancyDraftContactsWithFullPhone`

NewVacancyDraftContactsWithFullPhoneWithDefaults instantiates a new VacancyDraftContactsWithFullPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *VacancyDraftContactsWithFullPhone) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VacancyDraftContactsWithFullPhone) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VacancyDraftContactsWithFullPhone) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *VacancyDraftContactsWithFullPhone) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *VacancyDraftContactsWithFullPhone) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *VacancyDraftContactsWithFullPhone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyDraftContactsWithFullPhone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyDraftContactsWithFullPhone) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *VacancyDraftContactsWithFullPhone) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VacancyDraftContactsWithFullPhone) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhones

`func (o *VacancyDraftContactsWithFullPhone) GetPhones() []VacancyDraftPhoneItemWithFullPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *VacancyDraftContactsWithFullPhone) GetPhonesOk() (*[]VacancyDraftPhoneItemWithFullPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *VacancyDraftContactsWithFullPhone) SetPhones(v []VacancyDraftPhoneItemWithFullPhone)`

SetPhones sets Phones field to given value.


### SetPhonesNil

`func (o *VacancyDraftContactsWithFullPhone) SetPhonesNil(b bool)`

 SetPhonesNil sets the value for Phones to be an explicit nil

### UnsetPhones
`func (o *VacancyDraftContactsWithFullPhone) UnsetPhones()`

UnsetPhones ensures that no value is present for Phones, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


