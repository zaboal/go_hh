# VacancyDraftAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Адрес из [списка доступных адресов работодателя](#tag/Adresa-rabotodatelya/operation/get-employer-addresses) | [optional] 
**ShowMetroOnly** | Pointer to **NullableBool** | Показывать только метро для указанного адреса | [optional] 

## Methods

### NewVacancyDraftAddress

`func NewVacancyDraftAddress() *VacancyDraftAddress`

NewVacancyDraftAddress instantiates a new VacancyDraftAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftAddressWithDefaults

`func NewVacancyDraftAddressWithDefaults() *VacancyDraftAddress`

NewVacancyDraftAddressWithDefaults instantiates a new VacancyDraftAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyDraftAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyDraftAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyDraftAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacancyDraftAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacancyDraftAddress) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacancyDraftAddress) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetShowMetroOnly

`func (o *VacancyDraftAddress) GetShowMetroOnly() bool`

GetShowMetroOnly returns the ShowMetroOnly field if non-nil, zero value otherwise.

### GetShowMetroOnlyOk

`func (o *VacancyDraftAddress) GetShowMetroOnlyOk() (*bool, bool)`

GetShowMetroOnlyOk returns a tuple with the ShowMetroOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMetroOnly

`func (o *VacancyDraftAddress) SetShowMetroOnly(v bool)`

SetShowMetroOnly sets ShowMetroOnly field to given value.

### HasShowMetroOnly

`func (o *VacancyDraftAddress) HasShowMetroOnly() bool`

HasShowMetroOnly returns a boolean if a field has been set.

### SetShowMetroOnlyNil

`func (o *VacancyDraftAddress) SetShowMetroOnlyNil(b bool)`

 SetShowMetroOnlyNil sets the value for ShowMetroOnly to be an explicit nil

### UnsetShowMetroOnly
`func (o *VacancyDraftAddress) UnsetShowMetroOnly()`

UnsetShowMetroOnly ensures that no value is present for ShowMetroOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


