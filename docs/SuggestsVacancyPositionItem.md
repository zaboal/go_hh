# SuggestsVacancyPositionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор должности | 
**ProfessionalRoles** | [**[]SuggestsProfessionalRoleItemWithName**](SuggestsProfessionalRoleItemWithName.md) | Информация о профессиональных ролях, соответствующих должности | 
**Text** | **string** | Название должности | 

## Methods

### NewSuggestsVacancyPositionItem

`func NewSuggestsVacancyPositionItem(id string, professionalRoles []SuggestsProfessionalRoleItemWithName, text string, ) *SuggestsVacancyPositionItem`

NewSuggestsVacancyPositionItem instantiates a new SuggestsVacancyPositionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestsVacancyPositionItemWithDefaults

`func NewSuggestsVacancyPositionItemWithDefaults() *SuggestsVacancyPositionItem`

NewSuggestsVacancyPositionItemWithDefaults instantiates a new SuggestsVacancyPositionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuggestsVacancyPositionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuggestsVacancyPositionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuggestsVacancyPositionItem) SetId(v string)`

SetId sets Id field to given value.


### GetProfessionalRoles

`func (o *SuggestsVacancyPositionItem) GetProfessionalRoles() []SuggestsProfessionalRoleItemWithName`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *SuggestsVacancyPositionItem) GetProfessionalRolesOk() (*[]SuggestsProfessionalRoleItemWithName, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *SuggestsVacancyPositionItem) SetProfessionalRoles(v []SuggestsProfessionalRoleItemWithName)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetText

`func (o *SuggestsVacancyPositionItem) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SuggestsVacancyPositionItem) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SuggestsVacancyPositionItem) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


