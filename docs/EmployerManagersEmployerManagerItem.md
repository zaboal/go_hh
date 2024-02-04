# EmployerManagersEmployerManagerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalPhone** | Pointer to [**EmployerManagersEmployerManagerInfoAdditionalPhone**](EmployerManagersEmployerManagerInfoAdditionalPhone.md) |  | [optional] 
**Area** | [**EmployerManagersArea**](EmployerManagersArea.md) |  | 
**Email** | **string** | Адрес электронной почты менеджера | 
**FirstName** | Pointer to **string** | Имя менеджера | [optional] 
**FullName** | Pointer to **string** | Полное имя менеджера | [optional] 
**Id** | **string** | Идентификатор менеджера | 
**IsMainContactPerson** | Pointer to **bool** | Является ли менеджер основным контактным лицом | [optional] 
**LastName** | Pointer to **string** | Фамилия менеджера | [optional] 
**MiddleName** | Pointer to **string** | Отчество менеджера | [optional] 
**Name** | Pointer to **string** | Полное имя менеджера | [optional] 
**Phone** | Pointer to [**EmployerManagersEmployerManagerInfoPhone**](EmployerManagersEmployerManagerInfoPhone.md) |  | [optional] 
**Position** | **string** | Должность менеджера | 
**VacanciesCount** | Pointer to **NullableFloat32** | Количество опубликованных (активных) вакансий у данного менеджера. &#x60;null&#x60; — если у пользователя нет прав на просмотр вакансий этого менеджера | [optional] 

## Methods

### NewEmployerManagersEmployerManagerItem

`func NewEmployerManagersEmployerManagerItem(area EmployerManagersArea, email string, id string, position string, ) *EmployerManagersEmployerManagerItem`

NewEmployerManagersEmployerManagerItem instantiates a new EmployerManagersEmployerManagerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerManagersEmployerManagerItemWithDefaults

`func NewEmployerManagersEmployerManagerItemWithDefaults() *EmployerManagersEmployerManagerItem`

NewEmployerManagersEmployerManagerItemWithDefaults instantiates a new EmployerManagersEmployerManagerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalPhone

`func (o *EmployerManagersEmployerManagerItem) GetAdditionalPhone() EmployerManagersEmployerManagerInfoAdditionalPhone`

GetAdditionalPhone returns the AdditionalPhone field if non-nil, zero value otherwise.

### GetAdditionalPhoneOk

`func (o *EmployerManagersEmployerManagerItem) GetAdditionalPhoneOk() (*EmployerManagersEmployerManagerInfoAdditionalPhone, bool)`

GetAdditionalPhoneOk returns a tuple with the AdditionalPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPhone

`func (o *EmployerManagersEmployerManagerItem) SetAdditionalPhone(v EmployerManagersEmployerManagerInfoAdditionalPhone)`

SetAdditionalPhone sets AdditionalPhone field to given value.

### HasAdditionalPhone

`func (o *EmployerManagersEmployerManagerItem) HasAdditionalPhone() bool`

HasAdditionalPhone returns a boolean if a field has been set.

### GetArea

`func (o *EmployerManagersEmployerManagerItem) GetArea() EmployerManagersArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *EmployerManagersEmployerManagerItem) GetAreaOk() (*EmployerManagersArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *EmployerManagersEmployerManagerItem) SetArea(v EmployerManagersArea)`

SetArea sets Area field to given value.


### GetEmail

`func (o *EmployerManagersEmployerManagerItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmployerManagersEmployerManagerItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmployerManagersEmployerManagerItem) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *EmployerManagersEmployerManagerItem) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployerManagersEmployerManagerItem) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployerManagersEmployerManagerItem) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *EmployerManagersEmployerManagerItem) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFullName

`func (o *EmployerManagersEmployerManagerItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *EmployerManagersEmployerManagerItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *EmployerManagersEmployerManagerItem) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *EmployerManagersEmployerManagerItem) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetId

`func (o *EmployerManagersEmployerManagerItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployerManagersEmployerManagerItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployerManagersEmployerManagerItem) SetId(v string)`

SetId sets Id field to given value.


### GetIsMainContactPerson

`func (o *EmployerManagersEmployerManagerItem) GetIsMainContactPerson() bool`

GetIsMainContactPerson returns the IsMainContactPerson field if non-nil, zero value otherwise.

### GetIsMainContactPersonOk

`func (o *EmployerManagersEmployerManagerItem) GetIsMainContactPersonOk() (*bool, bool)`

GetIsMainContactPersonOk returns a tuple with the IsMainContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMainContactPerson

`func (o *EmployerManagersEmployerManagerItem) SetIsMainContactPerson(v bool)`

SetIsMainContactPerson sets IsMainContactPerson field to given value.

### HasIsMainContactPerson

`func (o *EmployerManagersEmployerManagerItem) HasIsMainContactPerson() bool`

HasIsMainContactPerson returns a boolean if a field has been set.

### GetLastName

`func (o *EmployerManagersEmployerManagerItem) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployerManagersEmployerManagerItem) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployerManagersEmployerManagerItem) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *EmployerManagersEmployerManagerItem) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *EmployerManagersEmployerManagerItem) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *EmployerManagersEmployerManagerItem) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *EmployerManagersEmployerManagerItem) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *EmployerManagersEmployerManagerItem) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetName

`func (o *EmployerManagersEmployerManagerItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployerManagersEmployerManagerItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployerManagersEmployerManagerItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmployerManagersEmployerManagerItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *EmployerManagersEmployerManagerItem) GetPhone() EmployerManagersEmployerManagerInfoPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EmployerManagersEmployerManagerItem) GetPhoneOk() (*EmployerManagersEmployerManagerInfoPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *EmployerManagersEmployerManagerItem) SetPhone(v EmployerManagersEmployerManagerInfoPhone)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *EmployerManagersEmployerManagerItem) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPosition

`func (o *EmployerManagersEmployerManagerItem) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EmployerManagersEmployerManagerItem) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EmployerManagersEmployerManagerItem) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetVacanciesCount

`func (o *EmployerManagersEmployerManagerItem) GetVacanciesCount() float32`

GetVacanciesCount returns the VacanciesCount field if non-nil, zero value otherwise.

### GetVacanciesCountOk

`func (o *EmployerManagersEmployerManagerItem) GetVacanciesCountOk() (*float32, bool)`

GetVacanciesCountOk returns a tuple with the VacanciesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacanciesCount

`func (o *EmployerManagersEmployerManagerItem) SetVacanciesCount(v float32)`

SetVacanciesCount sets VacanciesCount field to given value.

### HasVacanciesCount

`func (o *EmployerManagersEmployerManagerItem) HasVacanciesCount() bool`

HasVacanciesCount returns a boolean if a field has been set.

### SetVacanciesCountNil

`func (o *EmployerManagersEmployerManagerItem) SetVacanciesCountNil(b bool)`

 SetVacanciesCountNil sets the value for VacanciesCount to be an explicit nil

### UnsetVacanciesCount
`func (o *EmployerManagersEmployerManagerItem) UnsetVacanciesCount()`

UnsetVacanciesCount ensures that no value is present for VacanciesCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


