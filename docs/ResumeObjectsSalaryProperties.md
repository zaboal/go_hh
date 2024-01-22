# ResumeObjectsSalaryProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat32** | Сумма | [optional] 
**Currency** | Pointer to **NullableString** | Идентификатор валюты. Возможные значения перечислены в массиве &#x60;currency&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 

## Methods

### NewResumeObjectsSalaryProperties

`func NewResumeObjectsSalaryProperties() *ResumeObjectsSalaryProperties`

NewResumeObjectsSalaryProperties instantiates a new ResumeObjectsSalaryProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsSalaryPropertiesWithDefaults

`func NewResumeObjectsSalaryPropertiesWithDefaults() *ResumeObjectsSalaryProperties`

NewResumeObjectsSalaryPropertiesWithDefaults instantiates a new ResumeObjectsSalaryProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResumeObjectsSalaryProperties) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResumeObjectsSalaryProperties) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResumeObjectsSalaryProperties) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResumeObjectsSalaryProperties) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ResumeObjectsSalaryProperties) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ResumeObjectsSalaryProperties) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *ResumeObjectsSalaryProperties) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ResumeObjectsSalaryProperties) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ResumeObjectsSalaryProperties) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ResumeObjectsSalaryProperties) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ResumeObjectsSalaryProperties) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ResumeObjectsSalaryProperties) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


