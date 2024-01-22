# DictionariesAreaItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Areas** | [**[]DictionariesAreaItem**](DictionariesAreaItem.md) | Дочерние регионы | 
**Id** | **string** | Идентификатор региона | 
**Name** | **string** | Название региона | 
**NamePrepositional** | Pointer to **NullableString** | Применимо только для русской локализации.  Название региона в предложном падеже с предлогом \&quot;в\&quot;, например: \&quot;в Москве\&quot;. Возвращается, если в запросе передан параметр &#x60;additional_case&#x3D;prepositional&#x60;  | [optional] 
**ParentId** | Pointer to **NullableString** | Идентификатор родительского региона | [optional] 

## Methods

### NewDictionariesAreaItem

`func NewDictionariesAreaItem(areas []DictionariesAreaItem, id string, name string, ) *DictionariesAreaItem`

NewDictionariesAreaItem instantiates a new DictionariesAreaItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionariesAreaItemWithDefaults

`func NewDictionariesAreaItemWithDefaults() *DictionariesAreaItem`

NewDictionariesAreaItemWithDefaults instantiates a new DictionariesAreaItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreas

`func (o *DictionariesAreaItem) GetAreas() []DictionariesAreaItem`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *DictionariesAreaItem) GetAreasOk() (*[]DictionariesAreaItem, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *DictionariesAreaItem) SetAreas(v []DictionariesAreaItem)`

SetAreas sets Areas field to given value.


### GetId

`func (o *DictionariesAreaItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DictionariesAreaItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DictionariesAreaItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DictionariesAreaItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DictionariesAreaItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DictionariesAreaItem) SetName(v string)`

SetName sets Name field to given value.


### GetNamePrepositional

`func (o *DictionariesAreaItem) GetNamePrepositional() string`

GetNamePrepositional returns the NamePrepositional field if non-nil, zero value otherwise.

### GetNamePrepositionalOk

`func (o *DictionariesAreaItem) GetNamePrepositionalOk() (*string, bool)`

GetNamePrepositionalOk returns a tuple with the NamePrepositional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrepositional

`func (o *DictionariesAreaItem) SetNamePrepositional(v string)`

SetNamePrepositional sets NamePrepositional field to given value.

### HasNamePrepositional

`func (o *DictionariesAreaItem) HasNamePrepositional() bool`

HasNamePrepositional returns a boolean if a field has been set.

### SetNamePrepositionalNil

`func (o *DictionariesAreaItem) SetNamePrepositionalNil(b bool)`

 SetNamePrepositionalNil sets the value for NamePrepositional to be an explicit nil

### UnsetNamePrepositional
`func (o *DictionariesAreaItem) UnsetNamePrepositional()`

UnsetNamePrepositional ensures that no value is present for NamePrepositional, not even an explicit nil
### GetParentId

`func (o *DictionariesAreaItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DictionariesAreaItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DictionariesAreaItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DictionariesAreaItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *DictionariesAreaItem) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *DictionariesAreaItem) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


