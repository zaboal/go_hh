# ProfessionalRolesRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptIncompleteResumes** | **bool** | На роль принимаются отклики неполным резюме | 
**Id** | **string** | Идентификатор профессиональной роли | 
**IsDefault** | **bool** | Дефолтная роль | 
**Name** | **string** | Имя профессиональной роли | 
**SearchDeprecated** | Pointer to **bool** | Наличие запрета на использование в поиске при составлении поискового запроса | [optional] 
**SelectDeprecated** | Pointer to **bool** | Наличие запрета на использование при создании новых сущностей (резюме или вакансии) | [optional] 

## Methods

### NewProfessionalRolesRole

`func NewProfessionalRolesRole(acceptIncompleteResumes bool, id string, isDefault bool, name string, ) *ProfessionalRolesRole`

NewProfessionalRolesRole instantiates a new ProfessionalRolesRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfessionalRolesRoleWithDefaults

`func NewProfessionalRolesRoleWithDefaults() *ProfessionalRolesRole`

NewProfessionalRolesRoleWithDefaults instantiates a new ProfessionalRolesRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptIncompleteResumes

`func (o *ProfessionalRolesRole) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *ProfessionalRolesRole) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *ProfessionalRolesRole) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.


### GetId

`func (o *ProfessionalRolesRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfessionalRolesRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfessionalRolesRole) SetId(v string)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *ProfessionalRolesRole) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProfessionalRolesRole) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProfessionalRolesRole) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetName

`func (o *ProfessionalRolesRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfessionalRolesRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfessionalRolesRole) SetName(v string)`

SetName sets Name field to given value.


### GetSearchDeprecated

`func (o *ProfessionalRolesRole) GetSearchDeprecated() bool`

GetSearchDeprecated returns the SearchDeprecated field if non-nil, zero value otherwise.

### GetSearchDeprecatedOk

`func (o *ProfessionalRolesRole) GetSearchDeprecatedOk() (*bool, bool)`

GetSearchDeprecatedOk returns a tuple with the SearchDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDeprecated

`func (o *ProfessionalRolesRole) SetSearchDeprecated(v bool)`

SetSearchDeprecated sets SearchDeprecated field to given value.

### HasSearchDeprecated

`func (o *ProfessionalRolesRole) HasSearchDeprecated() bool`

HasSearchDeprecated returns a boolean if a field has been set.

### GetSelectDeprecated

`func (o *ProfessionalRolesRole) GetSelectDeprecated() bool`

GetSelectDeprecated returns the SelectDeprecated field if non-nil, zero value otherwise.

### GetSelectDeprecatedOk

`func (o *ProfessionalRolesRole) GetSelectDeprecatedOk() (*bool, bool)`

GetSelectDeprecatedOk returns a tuple with the SelectDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectDeprecated

`func (o *ProfessionalRolesRole) SetSelectDeprecated(v bool)`

SetSelectDeprecated sets SelectDeprecated field to given value.

### HasSelectDeprecated

`func (o *ProfessionalRolesRole) HasSelectDeprecated() bool`

HasSelectDeprecated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


