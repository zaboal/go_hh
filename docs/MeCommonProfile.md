# MeCommonProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **NullableString** | Тип авторизации | 
**IsAdmin** | **bool** | Является ли текущий пользователь администратором сайта | 
**IsApplicant** | **bool** | Является ли текущий пользователь соискателем | 
**IsApplication** | **bool** | Является ли авторизованный клиент приложением | 
**IsEmployer** | **bool** | Является ли текущий пользователь менеджером | 
**IsEmployerIntegration** | **bool** | Является ли текущий пользователь работодателем | 

## Methods

### NewMeCommonProfile

`func NewMeCommonProfile(authType NullableString, isAdmin bool, isApplicant bool, isApplication bool, isEmployer bool, isEmployerIntegration bool, ) *MeCommonProfile`

NewMeCommonProfile instantiates a new MeCommonProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeCommonProfileWithDefaults

`func NewMeCommonProfileWithDefaults() *MeCommonProfile`

NewMeCommonProfileWithDefaults instantiates a new MeCommonProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *MeCommonProfile) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *MeCommonProfile) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *MeCommonProfile) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### SetAuthTypeNil

`func (o *MeCommonProfile) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *MeCommonProfile) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetIsAdmin

`func (o *MeCommonProfile) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *MeCommonProfile) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *MeCommonProfile) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetIsApplicant

`func (o *MeCommonProfile) GetIsApplicant() bool`

GetIsApplicant returns the IsApplicant field if non-nil, zero value otherwise.

### GetIsApplicantOk

`func (o *MeCommonProfile) GetIsApplicantOk() (*bool, bool)`

GetIsApplicantOk returns a tuple with the IsApplicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApplicant

`func (o *MeCommonProfile) SetIsApplicant(v bool)`

SetIsApplicant sets IsApplicant field to given value.


### GetIsApplication

`func (o *MeCommonProfile) GetIsApplication() bool`

GetIsApplication returns the IsApplication field if non-nil, zero value otherwise.

### GetIsApplicationOk

`func (o *MeCommonProfile) GetIsApplicationOk() (*bool, bool)`

GetIsApplicationOk returns a tuple with the IsApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApplication

`func (o *MeCommonProfile) SetIsApplication(v bool)`

SetIsApplication sets IsApplication field to given value.


### GetIsEmployer

`func (o *MeCommonProfile) GetIsEmployer() bool`

GetIsEmployer returns the IsEmployer field if non-nil, zero value otherwise.

### GetIsEmployerOk

`func (o *MeCommonProfile) GetIsEmployerOk() (*bool, bool)`

GetIsEmployerOk returns a tuple with the IsEmployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmployer

`func (o *MeCommonProfile) SetIsEmployer(v bool)`

SetIsEmployer sets IsEmployer field to given value.


### GetIsEmployerIntegration

`func (o *MeCommonProfile) GetIsEmployerIntegration() bool`

GetIsEmployerIntegration returns the IsEmployerIntegration field if non-nil, zero value otherwise.

### GetIsEmployerIntegrationOk

`func (o *MeCommonProfile) GetIsEmployerIntegrationOk() (*bool, bool)`

GetIsEmployerIntegrationOk returns a tuple with the IsEmployerIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmployerIntegration

`func (o *MeCommonProfile) SetIsEmployerIntegration(v bool)`

SetIsEmployerIntegration sets IsEmployerIntegration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


