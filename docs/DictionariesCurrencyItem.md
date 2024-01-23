# DictionariesCurrencyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abbr** | **string** | Краткое обозначение | 
**Code** | **string** | Код | 
**Default** | **bool** | Используется ли в качестве валюты по умолчанию? | 
**InUse** | **bool** | Можно ли использовать на данном хосте? | 
**Name** | **string** | Название | 
**Rate** | **float32** | Курс по отношению к рублю | 

## Methods

### NewDictionariesCurrencyItem

`func NewDictionariesCurrencyItem(abbr string, code string, default_ bool, inUse bool, name string, rate float32, ) *DictionariesCurrencyItem`

NewDictionariesCurrencyItem instantiates a new DictionariesCurrencyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionariesCurrencyItemWithDefaults

`func NewDictionariesCurrencyItemWithDefaults() *DictionariesCurrencyItem`

NewDictionariesCurrencyItemWithDefaults instantiates a new DictionariesCurrencyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbr

`func (o *DictionariesCurrencyItem) GetAbbr() string`

GetAbbr returns the Abbr field if non-nil, zero value otherwise.

### GetAbbrOk

`func (o *DictionariesCurrencyItem) GetAbbrOk() (*string, bool)`

GetAbbrOk returns a tuple with the Abbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbr

`func (o *DictionariesCurrencyItem) SetAbbr(v string)`

SetAbbr sets Abbr field to given value.


### GetCode

`func (o *DictionariesCurrencyItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DictionariesCurrencyItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DictionariesCurrencyItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetDefault

`func (o *DictionariesCurrencyItem) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DictionariesCurrencyItem) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DictionariesCurrencyItem) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetInUse

`func (o *DictionariesCurrencyItem) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *DictionariesCurrencyItem) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *DictionariesCurrencyItem) SetInUse(v bool)`

SetInUse sets InUse field to given value.


### GetName

`func (o *DictionariesCurrencyItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DictionariesCurrencyItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DictionariesCurrencyItem) SetName(v string)`

SetName sets Name field to given value.


### GetRate

`func (o *DictionariesCurrencyItem) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *DictionariesCurrencyItem) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *DictionariesCurrencyItem) SetRate(v float32)`

SetRate sets Rate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


