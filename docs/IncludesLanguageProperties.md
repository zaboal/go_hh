# IncludesLanguageProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор | 
**Name** | **string** | Название | 
**Level** | Pointer to [**IncludesIdName**](IncludesIdName.md) | Уровень владения. Возможные значения элементов приведены в поле &#x60;language_level&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 

## Methods

### NewIncludesLanguageProperties

`func NewIncludesLanguageProperties(id string, name string, ) *IncludesLanguageProperties`

NewIncludesLanguageProperties instantiates a new IncludesLanguageProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludesLanguagePropertiesWithDefaults

`func NewIncludesLanguagePropertiesWithDefaults() *IncludesLanguageProperties`

NewIncludesLanguagePropertiesWithDefaults instantiates a new IncludesLanguageProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IncludesLanguageProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncludesLanguageProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncludesLanguageProperties) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IncludesLanguageProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncludesLanguageProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncludesLanguageProperties) SetName(v string)`

SetName sets Name field to given value.


### GetLevel

`func (o *IncludesLanguageProperties) GetLevel() IncludesIdName`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *IncludesLanguageProperties) GetLevelOk() (*IncludesIdName, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *IncludesLanguageProperties) SetLevel(v IncludesIdName)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *IncludesLanguageProperties) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


