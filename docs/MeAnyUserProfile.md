# MeAnyUserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Email текущего пользователя | [optional] 
**FirstName** | **string** | Имя текущего пользователя | 
**Id** | **string** | Идентификатор текущего пользователя | 
**IsAnonymous** | Pointer to **bool** |  | [optional] 
**LastName** | **string** | Фамилия текущего пользователя | 
**MidName** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **NullableString** | Отчество текущего пользователя | [optional] 
**Phone** | Pointer to **NullableString** | Телефон текущего пользователя | [optional] 

## Methods

### NewMeAnyUserProfile

`func NewMeAnyUserProfile(firstName string, id string, lastName string, ) *MeAnyUserProfile`

NewMeAnyUserProfile instantiates a new MeAnyUserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeAnyUserProfileWithDefaults

`func NewMeAnyUserProfileWithDefaults() *MeAnyUserProfile`

NewMeAnyUserProfileWithDefaults instantiates a new MeAnyUserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *MeAnyUserProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MeAnyUserProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MeAnyUserProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MeAnyUserProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MeAnyUserProfile) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MeAnyUserProfile) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFirstName

`func (o *MeAnyUserProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MeAnyUserProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MeAnyUserProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetId

`func (o *MeAnyUserProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeAnyUserProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeAnyUserProfile) SetId(v string)`

SetId sets Id field to given value.


### GetIsAnonymous

`func (o *MeAnyUserProfile) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *MeAnyUserProfile) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *MeAnyUserProfile) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.

### HasIsAnonymous

`func (o *MeAnyUserProfile) HasIsAnonymous() bool`

HasIsAnonymous returns a boolean if a field has been set.

### GetLastName

`func (o *MeAnyUserProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MeAnyUserProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MeAnyUserProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMidName

`func (o *MeAnyUserProfile) GetMidName() string`

GetMidName returns the MidName field if non-nil, zero value otherwise.

### GetMidNameOk

`func (o *MeAnyUserProfile) GetMidNameOk() (*string, bool)`

GetMidNameOk returns a tuple with the MidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidName

`func (o *MeAnyUserProfile) SetMidName(v string)`

SetMidName sets MidName field to given value.

### HasMidName

`func (o *MeAnyUserProfile) HasMidName() bool`

HasMidName returns a boolean if a field has been set.

### GetMiddleName

`func (o *MeAnyUserProfile) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *MeAnyUserProfile) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *MeAnyUserProfile) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *MeAnyUserProfile) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *MeAnyUserProfile) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *MeAnyUserProfile) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPhone

`func (o *MeAnyUserProfile) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MeAnyUserProfile) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MeAnyUserProfile) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MeAnyUserProfile) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MeAnyUserProfile) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MeAnyUserProfile) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


