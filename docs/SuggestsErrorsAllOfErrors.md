# SuggestsErrorsAllOfErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Поле, в котором допущена ошибка.  Возможные значения: * &#x60;locale&#x60; — указан неподдерживаемый язык * &#x60;text&#x60; — искомый текст должен быть длиной от 2 до 30000 символов * &#x60;area_id&#x60; - указан не валидный идентификатор  | [optional] 

## Methods

### NewSuggestsErrorsAllOfErrors

`func NewSuggestsErrorsAllOfErrors() *SuggestsErrorsAllOfErrors`

NewSuggestsErrorsAllOfErrors instantiates a new SuggestsErrorsAllOfErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestsErrorsAllOfErrorsWithDefaults

`func NewSuggestsErrorsAllOfErrorsWithDefaults() *SuggestsErrorsAllOfErrors`

NewSuggestsErrorsAllOfErrorsWithDefaults instantiates a new SuggestsErrorsAllOfErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SuggestsErrorsAllOfErrors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SuggestsErrorsAllOfErrors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SuggestsErrorsAllOfErrors) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SuggestsErrorsAllOfErrors) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *SuggestsErrorsAllOfErrors) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SuggestsErrorsAllOfErrors) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SuggestsErrorsAllOfErrors) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SuggestsErrorsAllOfErrors) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


