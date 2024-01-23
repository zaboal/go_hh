# SuggestsPositionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор должности | 
**ProfessionalRoles** | [**[]SuggestsProfessionalRoleItemWithName**](SuggestsProfessionalRoleItemWithName.md) | Информация о профессиональных ролях, соответствующих должности | 
**Specializations** | Pointer to [**[]SuggestsSpecializationsWithName**](SuggestsSpecializationsWithName.md) | Информация о специализациях, соответствующих должности | [optional] 
**Text** | **string** | Название должности | 

## Methods

### NewSuggestsPositionItem

`func NewSuggestsPositionItem(id string, professionalRoles []SuggestsProfessionalRoleItemWithName, text string, ) *SuggestsPositionItem`

NewSuggestsPositionItem instantiates a new SuggestsPositionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestsPositionItemWithDefaults

`func NewSuggestsPositionItemWithDefaults() *SuggestsPositionItem`

NewSuggestsPositionItemWithDefaults instantiates a new SuggestsPositionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuggestsPositionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuggestsPositionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuggestsPositionItem) SetId(v string)`

SetId sets Id field to given value.


### GetProfessionalRoles

`func (o *SuggestsPositionItem) GetProfessionalRoles() []SuggestsProfessionalRoleItemWithName`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *SuggestsPositionItem) GetProfessionalRolesOk() (*[]SuggestsProfessionalRoleItemWithName, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *SuggestsPositionItem) SetProfessionalRoles(v []SuggestsProfessionalRoleItemWithName)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetSpecializations

`func (o *SuggestsPositionItem) GetSpecializations() []SuggestsSpecializationsWithName`

GetSpecializations returns the Specializations field if non-nil, zero value otherwise.

### GetSpecializationsOk

`func (o *SuggestsPositionItem) GetSpecializationsOk() (*[]SuggestsSpecializationsWithName, bool)`

GetSpecializationsOk returns a tuple with the Specializations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecializations

`func (o *SuggestsPositionItem) SetSpecializations(v []SuggestsSpecializationsWithName)`

SetSpecializations sets Specializations field to given value.

### HasSpecializations

`func (o *SuggestsPositionItem) HasSpecializations() bool`

HasSpecializations returns a boolean if a field has been set.

### GetText

`func (o *SuggestsPositionItem) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SuggestsPositionItem) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SuggestsPositionItem) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


