# MeEmployerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employer** | Pointer to [**MeEmployerProfileCompany**](MeEmployerProfileCompany.md) |  | [optional] 
**PersonalManager** | Pointer to [**MeEmployerProfilePersonalManager**](MeEmployerProfilePersonalManager.md) |  | [optional] 

## Methods

### NewMeEmployerProfile

`func NewMeEmployerProfile() *MeEmployerProfile`

NewMeEmployerProfile instantiates a new MeEmployerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeEmployerProfileWithDefaults

`func NewMeEmployerProfileWithDefaults() *MeEmployerProfile`

NewMeEmployerProfileWithDefaults instantiates a new MeEmployerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployer

`func (o *MeEmployerProfile) GetEmployer() MeEmployerProfileCompany`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *MeEmployerProfile) GetEmployerOk() (*MeEmployerProfileCompany, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *MeEmployerProfile) SetEmployer(v MeEmployerProfileCompany)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *MeEmployerProfile) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetPersonalManager

`func (o *MeEmployerProfile) GetPersonalManager() MeEmployerProfilePersonalManager`

GetPersonalManager returns the PersonalManager field if non-nil, zero value otherwise.

### GetPersonalManagerOk

`func (o *MeEmployerProfile) GetPersonalManagerOk() (*MeEmployerProfilePersonalManager, bool)`

GetPersonalManagerOk returns a tuple with the PersonalManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalManager

`func (o *MeEmployerProfile) SetPersonalManager(v MeEmployerProfilePersonalManager)`

SetPersonalManager sets PersonalManager field to given value.

### HasPersonalManager

`func (o *MeEmployerProfile) HasPersonalManager() bool`

HasPersonalManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


