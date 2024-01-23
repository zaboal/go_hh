# MeEmployerProfilePersonalManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email персонального менеджера | 
**FirstName** | **string** | Имя персонального менеджера | 
**Id** | **string** | Идентификатор персонального менеджера | 
**IsAvailable** | **bool** | Доступен ли менеджер в данный момент | 
**LastName** | **string** | Фамилия персонального менеджера | 
**PhotoUrls** | Pointer to [**NullableMeEmployerProfilePersonalManagerPhotoUrls**](MeEmployerProfilePersonalManagerPhotoUrls.md) |  | [optional] 
**Unavailable** | Pointer to [**NullableMeEmployerProfilePersonalManagerUnavailable**](MeEmployerProfilePersonalManagerUnavailable.md) |  | [optional] 

## Methods

### NewMeEmployerProfilePersonalManager

`func NewMeEmployerProfilePersonalManager(email string, firstName string, id string, isAvailable bool, lastName string, ) *MeEmployerProfilePersonalManager`

NewMeEmployerProfilePersonalManager instantiates a new MeEmployerProfilePersonalManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeEmployerProfilePersonalManagerWithDefaults

`func NewMeEmployerProfilePersonalManagerWithDefaults() *MeEmployerProfilePersonalManager`

NewMeEmployerProfilePersonalManagerWithDefaults instantiates a new MeEmployerProfilePersonalManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *MeEmployerProfilePersonalManager) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MeEmployerProfilePersonalManager) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MeEmployerProfilePersonalManager) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *MeEmployerProfilePersonalManager) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MeEmployerProfilePersonalManager) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MeEmployerProfilePersonalManager) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetId

`func (o *MeEmployerProfilePersonalManager) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeEmployerProfilePersonalManager) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeEmployerProfilePersonalManager) SetId(v string)`

SetId sets Id field to given value.


### GetIsAvailable

`func (o *MeEmployerProfilePersonalManager) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *MeEmployerProfilePersonalManager) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *MeEmployerProfilePersonalManager) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.


### GetLastName

`func (o *MeEmployerProfilePersonalManager) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MeEmployerProfilePersonalManager) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MeEmployerProfilePersonalManager) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPhotoUrls

`func (o *MeEmployerProfilePersonalManager) GetPhotoUrls() MeEmployerProfilePersonalManagerPhotoUrls`

GetPhotoUrls returns the PhotoUrls field if non-nil, zero value otherwise.

### GetPhotoUrlsOk

`func (o *MeEmployerProfilePersonalManager) GetPhotoUrlsOk() (*MeEmployerProfilePersonalManagerPhotoUrls, bool)`

GetPhotoUrlsOk returns a tuple with the PhotoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrls

`func (o *MeEmployerProfilePersonalManager) SetPhotoUrls(v MeEmployerProfilePersonalManagerPhotoUrls)`

SetPhotoUrls sets PhotoUrls field to given value.

### HasPhotoUrls

`func (o *MeEmployerProfilePersonalManager) HasPhotoUrls() bool`

HasPhotoUrls returns a boolean if a field has been set.

### SetPhotoUrlsNil

`func (o *MeEmployerProfilePersonalManager) SetPhotoUrlsNil(b bool)`

 SetPhotoUrlsNil sets the value for PhotoUrls to be an explicit nil

### UnsetPhotoUrls
`func (o *MeEmployerProfilePersonalManager) UnsetPhotoUrls()`

UnsetPhotoUrls ensures that no value is present for PhotoUrls, not even an explicit nil
### GetUnavailable

`func (o *MeEmployerProfilePersonalManager) GetUnavailable() MeEmployerProfilePersonalManagerUnavailable`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *MeEmployerProfilePersonalManager) GetUnavailableOk() (*MeEmployerProfilePersonalManagerUnavailable, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *MeEmployerProfilePersonalManager) SetUnavailable(v MeEmployerProfilePersonalManagerUnavailable)`

SetUnavailable sets Unavailable field to given value.

### HasUnavailable

`func (o *MeEmployerProfilePersonalManager) HasUnavailable() bool`

HasUnavailable returns a boolean if a field has been set.

### SetUnavailableNil

`func (o *MeEmployerProfilePersonalManager) SetUnavailableNil(b bool)`

 SetUnavailableNil sets the value for Unavailable to be an explicit nil

### UnsetUnavailable
`func (o *MeEmployerProfilePersonalManager) UnsetUnavailable()`

UnsetUnavailable ensures that no value is present for Unavailable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


