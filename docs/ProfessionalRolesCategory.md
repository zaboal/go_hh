# ProfessionalRolesCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор категории профессиональной роли | 
**Name** | **NullableString** | Имя категории профессиональной роли | 
**Roles** | [**[]ProfessionalRolesRole**](ProfessionalRolesRole.md) | Список профессиональных ролей, входящих в эту категорию  | 

## Methods

### NewProfessionalRolesCategory

`func NewProfessionalRolesCategory(id string, name NullableString, roles []ProfessionalRolesRole, ) *ProfessionalRolesCategory`

NewProfessionalRolesCategory instantiates a new ProfessionalRolesCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfessionalRolesCategoryWithDefaults

`func NewProfessionalRolesCategoryWithDefaults() *ProfessionalRolesCategory`

NewProfessionalRolesCategoryWithDefaults instantiates a new ProfessionalRolesCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfessionalRolesCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfessionalRolesCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfessionalRolesCategory) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProfessionalRolesCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfessionalRolesCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfessionalRolesCategory) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProfessionalRolesCategory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProfessionalRolesCategory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRoles

`func (o *ProfessionalRolesCategory) GetRoles() []ProfessionalRolesRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ProfessionalRolesCategory) GetRolesOk() (*[]ProfessionalRolesRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ProfessionalRolesCategory) SetRoles(v []ProfessionalRolesRole)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


