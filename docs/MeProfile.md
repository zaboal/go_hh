# MeProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **NullableString** | Тип авторизации | 
**IsAdmin** | **bool** | Является ли текущий пользователь администратором сайта | 
**IsApplicant** | **bool** | Является ли текущий пользователь соискателем | 
**IsApplication** | **bool** | Является ли авторизованный клиент приложением | 
**IsEmployer** | **bool** | Является ли текущий пользователь менеджером | 
**IsEmployerIntegration** | **bool** | Является ли текущий пользователь работодателем | 
**Email** | Pointer to **NullableString** | Email текущего пользователя | [optional] 
**FirstName** | **string** | Имя текущего пользователя | 
**Id** | **string** | Идентификатор текущего пользователя | 
**IsAnonymous** | Pointer to **bool** |  | [optional] 
**LastName** | **string** | Фамилия текущего пользователя | 
**MidName** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **NullableString** | Отчество текущего пользователя | [optional] 
**Phone** | Pointer to **NullableString** | Телефон текущего пользователя | [optional] 
**Counters** | [**MeApplicantProfileCounters**](MeApplicantProfileCounters.md) |  | 
**Employer** | Pointer to [**MeEmployerProfileCompany**](MeEmployerProfileCompany.md) |  | [optional] 
**IsInSearch** | **map[string]interface{}** |  | 
**Manager** | Pointer to [**MeEmployerProfileManager**](MeEmployerProfileManager.md) |  | [optional] 
**NegotiationsUrl** | **map[string]interface{}** |  | 
**PersonalManager** | Pointer to [**MeEmployerProfilePersonalManager**](MeEmployerProfilePersonalManager.md) |  | [optional] 
**ProfileVideos** | Pointer to [**ProfileVideosList**](ProfileVideosList.md) |  | [optional] 
**ResumesUrl** | **map[string]interface{}** |  | 
**UserStatuses** | Pointer to [**UserStatusesApplicant**](UserStatusesApplicant.md) |  | [optional] 

## Methods

### NewMeProfile

`func NewMeProfile(authType NullableString, isAdmin bool, isApplicant bool, isApplication bool, isEmployer bool, isEmployerIntegration bool, firstName string, id string, lastName string, counters MeApplicantProfileCounters, isInSearch map[string]interface{}, negotiationsUrl map[string]interface{}, resumesUrl map[string]interface{}, ) *MeProfile`

NewMeProfile instantiates a new MeProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeProfileWithDefaults

`func NewMeProfileWithDefaults() *MeProfile`

NewMeProfileWithDefaults instantiates a new MeProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *MeProfile) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *MeProfile) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *MeProfile) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### SetAuthTypeNil

`func (o *MeProfile) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *MeProfile) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetIsAdmin

`func (o *MeProfile) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *MeProfile) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *MeProfile) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetIsApplicant

`func (o *MeProfile) GetIsApplicant() bool`

GetIsApplicant returns the IsApplicant field if non-nil, zero value otherwise.

### GetIsApplicantOk

`func (o *MeProfile) GetIsApplicantOk() (*bool, bool)`

GetIsApplicantOk returns a tuple with the IsApplicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApplicant

`func (o *MeProfile) SetIsApplicant(v bool)`

SetIsApplicant sets IsApplicant field to given value.


### GetIsApplication

`func (o *MeProfile) GetIsApplication() bool`

GetIsApplication returns the IsApplication field if non-nil, zero value otherwise.

### GetIsApplicationOk

`func (o *MeProfile) GetIsApplicationOk() (*bool, bool)`

GetIsApplicationOk returns a tuple with the IsApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApplication

`func (o *MeProfile) SetIsApplication(v bool)`

SetIsApplication sets IsApplication field to given value.


### GetIsEmployer

`func (o *MeProfile) GetIsEmployer() bool`

GetIsEmployer returns the IsEmployer field if non-nil, zero value otherwise.

### GetIsEmployerOk

`func (o *MeProfile) GetIsEmployerOk() (*bool, bool)`

GetIsEmployerOk returns a tuple with the IsEmployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmployer

`func (o *MeProfile) SetIsEmployer(v bool)`

SetIsEmployer sets IsEmployer field to given value.


### GetIsEmployerIntegration

`func (o *MeProfile) GetIsEmployerIntegration() bool`

GetIsEmployerIntegration returns the IsEmployerIntegration field if non-nil, zero value otherwise.

### GetIsEmployerIntegrationOk

`func (o *MeProfile) GetIsEmployerIntegrationOk() (*bool, bool)`

GetIsEmployerIntegrationOk returns a tuple with the IsEmployerIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmployerIntegration

`func (o *MeProfile) SetIsEmployerIntegration(v bool)`

SetIsEmployerIntegration sets IsEmployerIntegration field to given value.


### GetEmail

`func (o *MeProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MeProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MeProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MeProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MeProfile) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MeProfile) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFirstName

`func (o *MeProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MeProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MeProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetId

`func (o *MeProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeProfile) SetId(v string)`

SetId sets Id field to given value.


### GetIsAnonymous

`func (o *MeProfile) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *MeProfile) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *MeProfile) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.

### HasIsAnonymous

`func (o *MeProfile) HasIsAnonymous() bool`

HasIsAnonymous returns a boolean if a field has been set.

### GetLastName

`func (o *MeProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MeProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MeProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMidName

`func (o *MeProfile) GetMidName() string`

GetMidName returns the MidName field if non-nil, zero value otherwise.

### GetMidNameOk

`func (o *MeProfile) GetMidNameOk() (*string, bool)`

GetMidNameOk returns a tuple with the MidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidName

`func (o *MeProfile) SetMidName(v string)`

SetMidName sets MidName field to given value.

### HasMidName

`func (o *MeProfile) HasMidName() bool`

HasMidName returns a boolean if a field has been set.

### GetMiddleName

`func (o *MeProfile) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *MeProfile) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *MeProfile) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *MeProfile) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *MeProfile) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *MeProfile) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPhone

`func (o *MeProfile) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MeProfile) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MeProfile) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MeProfile) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MeProfile) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MeProfile) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCounters

`func (o *MeProfile) GetCounters() MeApplicantProfileCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *MeProfile) GetCountersOk() (*MeApplicantProfileCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *MeProfile) SetCounters(v MeApplicantProfileCounters)`

SetCounters sets Counters field to given value.


### GetEmployer

`func (o *MeProfile) GetEmployer() MeEmployerProfileCompany`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *MeProfile) GetEmployerOk() (*MeEmployerProfileCompany, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *MeProfile) SetEmployer(v MeEmployerProfileCompany)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *MeProfile) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetIsInSearch

`func (o *MeProfile) GetIsInSearch() map[string]interface{}`

GetIsInSearch returns the IsInSearch field if non-nil, zero value otherwise.

### GetIsInSearchOk

`func (o *MeProfile) GetIsInSearchOk() (*map[string]interface{}, bool)`

GetIsInSearchOk returns a tuple with the IsInSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInSearch

`func (o *MeProfile) SetIsInSearch(v map[string]interface{})`

SetIsInSearch sets IsInSearch field to given value.


### SetIsInSearchNil

`func (o *MeProfile) SetIsInSearchNil(b bool)`

 SetIsInSearchNil sets the value for IsInSearch to be an explicit nil

### UnsetIsInSearch
`func (o *MeProfile) UnsetIsInSearch()`

UnsetIsInSearch ensures that no value is present for IsInSearch, not even an explicit nil
### GetManager

`func (o *MeProfile) GetManager() MeEmployerProfileManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *MeProfile) GetManagerOk() (*MeEmployerProfileManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *MeProfile) SetManager(v MeEmployerProfileManager)`

SetManager sets Manager field to given value.

### HasManager

`func (o *MeProfile) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetNegotiationsUrl

`func (o *MeProfile) GetNegotiationsUrl() map[string]interface{}`

GetNegotiationsUrl returns the NegotiationsUrl field if non-nil, zero value otherwise.

### GetNegotiationsUrlOk

`func (o *MeProfile) GetNegotiationsUrlOk() (*map[string]interface{}, bool)`

GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsUrl

`func (o *MeProfile) SetNegotiationsUrl(v map[string]interface{})`

SetNegotiationsUrl sets NegotiationsUrl field to given value.


### SetNegotiationsUrlNil

`func (o *MeProfile) SetNegotiationsUrlNil(b bool)`

 SetNegotiationsUrlNil sets the value for NegotiationsUrl to be an explicit nil

### UnsetNegotiationsUrl
`func (o *MeProfile) UnsetNegotiationsUrl()`

UnsetNegotiationsUrl ensures that no value is present for NegotiationsUrl, not even an explicit nil
### GetPersonalManager

`func (o *MeProfile) GetPersonalManager() MeEmployerProfilePersonalManager`

GetPersonalManager returns the PersonalManager field if non-nil, zero value otherwise.

### GetPersonalManagerOk

`func (o *MeProfile) GetPersonalManagerOk() (*MeEmployerProfilePersonalManager, bool)`

GetPersonalManagerOk returns a tuple with the PersonalManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalManager

`func (o *MeProfile) SetPersonalManager(v MeEmployerProfilePersonalManager)`

SetPersonalManager sets PersonalManager field to given value.

### HasPersonalManager

`func (o *MeProfile) HasPersonalManager() bool`

HasPersonalManager returns a boolean if a field has been set.

### GetProfileVideos

`func (o *MeProfile) GetProfileVideos() ProfileVideosList`

GetProfileVideos returns the ProfileVideos field if non-nil, zero value otherwise.

### GetProfileVideosOk

`func (o *MeProfile) GetProfileVideosOk() (*ProfileVideosList, bool)`

GetProfileVideosOk returns a tuple with the ProfileVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileVideos

`func (o *MeProfile) SetProfileVideos(v ProfileVideosList)`

SetProfileVideos sets ProfileVideos field to given value.

### HasProfileVideos

`func (o *MeProfile) HasProfileVideos() bool`

HasProfileVideos returns a boolean if a field has been set.

### GetResumesUrl

`func (o *MeProfile) GetResumesUrl() map[string]interface{}`

GetResumesUrl returns the ResumesUrl field if non-nil, zero value otherwise.

### GetResumesUrlOk

`func (o *MeProfile) GetResumesUrlOk() (*map[string]interface{}, bool)`

GetResumesUrlOk returns a tuple with the ResumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumesUrl

`func (o *MeProfile) SetResumesUrl(v map[string]interface{})`

SetResumesUrl sets ResumesUrl field to given value.


### SetResumesUrlNil

`func (o *MeProfile) SetResumesUrlNil(b bool)`

 SetResumesUrlNil sets the value for ResumesUrl to be an explicit nil

### UnsetResumesUrl
`func (o *MeProfile) UnsetResumesUrl()`

UnsetResumesUrl ensures that no value is present for ResumesUrl, not even an explicit nil
### GetUserStatuses

`func (o *MeProfile) GetUserStatuses() UserStatusesApplicant`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MeProfile) GetUserStatusesOk() (*UserStatusesApplicant, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatuses

`func (o *MeProfile) SetUserStatuses(v UserStatusesApplicant)`

SetUserStatuses sets UserStatuses field to given value.

### HasUserStatuses

`func (o *MeProfile) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


