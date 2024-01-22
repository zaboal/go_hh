# ErrorsVacancyError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | Pointer to **float32** | Количество найденных дублей вакансии | [optional] 
**Items** | Pointer to [**[]ErrorsVacancyErrorItemsInner**](ErrorsVacancyErrorItemsInner.md) | Массив идентификаторов вакансий-дублей | [optional] 
**Reason** | Pointer to **string** | Причина ошибки | [optional] 
**Type** | **string** | Тип ошибки | 
**Value** | **string** | Наименование поля с ошибкой | 

## Methods

### NewErrorsVacancyError

`func NewErrorsVacancyError(type_ string, value string, ) *ErrorsVacancyError`

NewErrorsVacancyError instantiates a new ErrorsVacancyError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyErrorWithDefaults

`func NewErrorsVacancyErrorWithDefaults() *ErrorsVacancyError`

NewErrorsVacancyErrorWithDefaults instantiates a new ErrorsVacancyError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *ErrorsVacancyError) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *ErrorsVacancyError) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *ErrorsVacancyError) SetFound(v float32)`

SetFound sets Found field to given value.

### HasFound

`func (o *ErrorsVacancyError) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetItems

`func (o *ErrorsVacancyError) GetItems() []ErrorsVacancyErrorItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ErrorsVacancyError) GetItemsOk() (*[]ErrorsVacancyErrorItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ErrorsVacancyError) SetItems(v []ErrorsVacancyErrorItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *ErrorsVacancyError) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetReason

`func (o *ErrorsVacancyError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorsVacancyError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorsVacancyError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorsVacancyError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetType

`func (o *ErrorsVacancyError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsVacancyError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsVacancyError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsVacancyError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsVacancyError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsVacancyError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


