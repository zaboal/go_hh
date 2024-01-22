# IncludesEducationalInstitutionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acronym** | Pointer to **NullableString** | Сокращенное название учебного заведения | [optional] 
**Area** | [**IncludesEducationalInstitutionItemArea**](IncludesEducationalInstitutionItemArea.md) |  | 
**Id** | **string** | Идентификатор учебного заведения | 
**Synonyms** | Pointer to **NullableString** | Альтернативное название учебного заведения | [optional] 
**Text** | **string** | Полное название учебного заведения | 

## Methods

### NewIncludesEducationalInstitutionItem

`func NewIncludesEducationalInstitutionItem(area IncludesEducationalInstitutionItemArea, id string, text string, ) *IncludesEducationalInstitutionItem`

NewIncludesEducationalInstitutionItem instantiates a new IncludesEducationalInstitutionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludesEducationalInstitutionItemWithDefaults

`func NewIncludesEducationalInstitutionItemWithDefaults() *IncludesEducationalInstitutionItem`

NewIncludesEducationalInstitutionItemWithDefaults instantiates a new IncludesEducationalInstitutionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcronym

`func (o *IncludesEducationalInstitutionItem) GetAcronym() string`

GetAcronym returns the Acronym field if non-nil, zero value otherwise.

### GetAcronymOk

`func (o *IncludesEducationalInstitutionItem) GetAcronymOk() (*string, bool)`

GetAcronymOk returns a tuple with the Acronym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcronym

`func (o *IncludesEducationalInstitutionItem) SetAcronym(v string)`

SetAcronym sets Acronym field to given value.

### HasAcronym

`func (o *IncludesEducationalInstitutionItem) HasAcronym() bool`

HasAcronym returns a boolean if a field has been set.

### SetAcronymNil

`func (o *IncludesEducationalInstitutionItem) SetAcronymNil(b bool)`

 SetAcronymNil sets the value for Acronym to be an explicit nil

### UnsetAcronym
`func (o *IncludesEducationalInstitutionItem) UnsetAcronym()`

UnsetAcronym ensures that no value is present for Acronym, not even an explicit nil
### GetArea

`func (o *IncludesEducationalInstitutionItem) GetArea() IncludesEducationalInstitutionItemArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *IncludesEducationalInstitutionItem) GetAreaOk() (*IncludesEducationalInstitutionItemArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *IncludesEducationalInstitutionItem) SetArea(v IncludesEducationalInstitutionItemArea)`

SetArea sets Area field to given value.


### GetId

`func (o *IncludesEducationalInstitutionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncludesEducationalInstitutionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncludesEducationalInstitutionItem) SetId(v string)`

SetId sets Id field to given value.


### GetSynonyms

`func (o *IncludesEducationalInstitutionItem) GetSynonyms() string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *IncludesEducationalInstitutionItem) GetSynonymsOk() (*string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *IncludesEducationalInstitutionItem) SetSynonyms(v string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *IncludesEducationalInstitutionItem) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### SetSynonymsNil

`func (o *IncludesEducationalInstitutionItem) SetSynonymsNil(b bool)`

 SetSynonymsNil sets the value for Synonyms to be an explicit nil

### UnsetSynonyms
`func (o *IncludesEducationalInstitutionItem) UnsetSynonyms()`

UnsetSynonyms ensures that no value is present for Synonyms, not even an explicit nil
### GetText

`func (o *IncludesEducationalInstitutionItem) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *IncludesEducationalInstitutionItem) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *IncludesEducationalInstitutionItem) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


