# ResumesAccessTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoHideTimeOptions** | Pointer to [**[]ResumesAutoHideTimeOptions**](ResumesAutoHideTimeOptions.md) | Варианты времени автоскрытия резюме при неактивности пользователя. Возвращается только для пользователей rabota.by | [optional] 
**Items** | [**[]ResumesAccessTypesItem**](ResumesAccessTypesItem.md) | Доступные типы видимости резюме | 

## Methods

### NewResumesAccessTypes

`func NewResumesAccessTypes(items []ResumesAccessTypesItem, ) *ResumesAccessTypes`

NewResumesAccessTypes instantiates a new ResumesAccessTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesAccessTypesWithDefaults

`func NewResumesAccessTypesWithDefaults() *ResumesAccessTypes`

NewResumesAccessTypesWithDefaults instantiates a new ResumesAccessTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoHideTimeOptions

`func (o *ResumesAccessTypes) GetAutoHideTimeOptions() []ResumesAutoHideTimeOptions`

GetAutoHideTimeOptions returns the AutoHideTimeOptions field if non-nil, zero value otherwise.

### GetAutoHideTimeOptionsOk

`func (o *ResumesAccessTypes) GetAutoHideTimeOptionsOk() (*[]ResumesAutoHideTimeOptions, bool)`

GetAutoHideTimeOptionsOk returns a tuple with the AutoHideTimeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoHideTimeOptions

`func (o *ResumesAccessTypes) SetAutoHideTimeOptions(v []ResumesAutoHideTimeOptions)`

SetAutoHideTimeOptions sets AutoHideTimeOptions field to given value.

### HasAutoHideTimeOptions

`func (o *ResumesAccessTypes) HasAutoHideTimeOptions() bool`

HasAutoHideTimeOptions returns a boolean if a field has been set.

### SetAutoHideTimeOptionsNil

`func (o *ResumesAccessTypes) SetAutoHideTimeOptionsNil(b bool)`

 SetAutoHideTimeOptionsNil sets the value for AutoHideTimeOptions to be an explicit nil

### UnsetAutoHideTimeOptions
`func (o *ResumesAccessTypes) UnsetAutoHideTimeOptions()`

UnsetAutoHideTimeOptions ensures that no value is present for AutoHideTimeOptions, not even an explicit nil
### GetItems

`func (o *ResumesAccessTypes) GetItems() []ResumesAccessTypesItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResumesAccessTypes) GetItemsOk() (*[]ResumesAccessTypesItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResumesAccessTypes) SetItems(v []ResumesAccessTypesItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


