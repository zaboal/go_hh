# VacanciesArgumentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Argument** | **string** | Параметр поиска вакансии | 
**ClusterGroup** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**DisableUrl** | **string** | URL поиска вакансий, который получится, если перестать учитывать в поиске данный параметр | 
**HexColor** | Pointer to **NullableString** | Цвет линии в HEX-формате &#x60;RRGGBB&#x60; (от &#x60;000000&#x60; до &#x60;FFFFFF&#x60;). Возвращается только для аргумента &#x60;metro&#x60; | [optional] 
**MetroType** | Pointer to **NullableString** | Станция или линия метро (&#x60;station&#x60;/&#x60;line&#x60;). Возвращается только для аргумента &#x60;metro&#x60; | [optional] 
**Name** | Pointer to **NullableString** | Название значения | [optional] 
**Value** | **string** | Значение параметра | 
**ValueDescription** | Pointer to **NullableString** | Описание параметра | [optional] 

## Methods

### NewVacanciesArgumentItem

`func NewVacanciesArgumentItem(argument string, disableUrl string, value string, ) *VacanciesArgumentItem`

NewVacanciesArgumentItem instantiates a new VacanciesArgumentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesArgumentItemWithDefaults

`func NewVacanciesArgumentItemWithDefaults() *VacanciesArgumentItem`

NewVacanciesArgumentItemWithDefaults instantiates a new VacanciesArgumentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgument

`func (o *VacanciesArgumentItem) GetArgument() string`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *VacanciesArgumentItem) GetArgumentOk() (*string, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *VacanciesArgumentItem) SetArgument(v string)`

SetArgument sets Argument field to given value.


### GetClusterGroup

`func (o *VacanciesArgumentItem) GetClusterGroup() IncludesIdName`

GetClusterGroup returns the ClusterGroup field if non-nil, zero value otherwise.

### GetClusterGroupOk

`func (o *VacanciesArgumentItem) GetClusterGroupOk() (*IncludesIdName, bool)`

GetClusterGroupOk returns a tuple with the ClusterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroup

`func (o *VacanciesArgumentItem) SetClusterGroup(v IncludesIdName)`

SetClusterGroup sets ClusterGroup field to given value.

### HasClusterGroup

`func (o *VacanciesArgumentItem) HasClusterGroup() bool`

HasClusterGroup returns a boolean if a field has been set.

### SetClusterGroupNil

`func (o *VacanciesArgumentItem) SetClusterGroupNil(b bool)`

 SetClusterGroupNil sets the value for ClusterGroup to be an explicit nil

### UnsetClusterGroup
`func (o *VacanciesArgumentItem) UnsetClusterGroup()`

UnsetClusterGroup ensures that no value is present for ClusterGroup, not even an explicit nil
### GetDisableUrl

`func (o *VacanciesArgumentItem) GetDisableUrl() string`

GetDisableUrl returns the DisableUrl field if non-nil, zero value otherwise.

### GetDisableUrlOk

`func (o *VacanciesArgumentItem) GetDisableUrlOk() (*string, bool)`

GetDisableUrlOk returns a tuple with the DisableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUrl

`func (o *VacanciesArgumentItem) SetDisableUrl(v string)`

SetDisableUrl sets DisableUrl field to given value.


### GetHexColor

`func (o *VacanciesArgumentItem) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *VacanciesArgumentItem) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *VacanciesArgumentItem) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.

### HasHexColor

`func (o *VacanciesArgumentItem) HasHexColor() bool`

HasHexColor returns a boolean if a field has been set.

### SetHexColorNil

`func (o *VacanciesArgumentItem) SetHexColorNil(b bool)`

 SetHexColorNil sets the value for HexColor to be an explicit nil

### UnsetHexColor
`func (o *VacanciesArgumentItem) UnsetHexColor()`

UnsetHexColor ensures that no value is present for HexColor, not even an explicit nil
### GetMetroType

`func (o *VacanciesArgumentItem) GetMetroType() string`

GetMetroType returns the MetroType field if non-nil, zero value otherwise.

### GetMetroTypeOk

`func (o *VacanciesArgumentItem) GetMetroTypeOk() (*string, bool)`

GetMetroTypeOk returns a tuple with the MetroType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroType

`func (o *VacanciesArgumentItem) SetMetroType(v string)`

SetMetroType sets MetroType field to given value.

### HasMetroType

`func (o *VacanciesArgumentItem) HasMetroType() bool`

HasMetroType returns a boolean if a field has been set.

### SetMetroTypeNil

`func (o *VacanciesArgumentItem) SetMetroTypeNil(b bool)`

 SetMetroTypeNil sets the value for MetroType to be an explicit nil

### UnsetMetroType
`func (o *VacanciesArgumentItem) UnsetMetroType()`

UnsetMetroType ensures that no value is present for MetroType, not even an explicit nil
### GetName

`func (o *VacanciesArgumentItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesArgumentItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesArgumentItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacanciesArgumentItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VacanciesArgumentItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VacanciesArgumentItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *VacanciesArgumentItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VacanciesArgumentItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VacanciesArgumentItem) SetValue(v string)`

SetValue sets Value field to given value.


### GetValueDescription

`func (o *VacanciesArgumentItem) GetValueDescription() string`

GetValueDescription returns the ValueDescription field if non-nil, zero value otherwise.

### GetValueDescriptionOk

`func (o *VacanciesArgumentItem) GetValueDescriptionOk() (*string, bool)`

GetValueDescriptionOk returns a tuple with the ValueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDescription

`func (o *VacanciesArgumentItem) SetValueDescription(v string)`

SetValueDescription sets ValueDescription field to given value.

### HasValueDescription

`func (o *VacanciesArgumentItem) HasValueDescription() bool`

HasValueDescription returns a boolean if a field has been set.

### SetValueDescriptionNil

`func (o *VacanciesArgumentItem) SetValueDescriptionNil(b bool)`

 SetValueDescriptionNil sets the value for ValueDescription to be an explicit nil

### UnsetValueDescription
`func (o *VacanciesArgumentItem) UnsetValueDescription()`

UnsetValueDescription ensures that no value is present for ValueDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


