# ErrorsSavedSearchNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | Pointer to **string** | Название поля запроса с ошибкой. Возможные значения: * &#x60;saved_search_not_found&#x60; — автопоиск не найден или не принадлежит текущему пользователю. * &#x60;manager_not_found&#x60; — менеджер не найден  | [optional] 

## Methods

### NewErrorsSavedSearchNotFoundError

`func NewErrorsSavedSearchNotFoundError(type_ string, ) *ErrorsSavedSearchNotFoundError`

NewErrorsSavedSearchNotFoundError instantiates a new ErrorsSavedSearchNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsSavedSearchNotFoundErrorWithDefaults

`func NewErrorsSavedSearchNotFoundErrorWithDefaults() *ErrorsSavedSearchNotFoundError`

NewErrorsSavedSearchNotFoundErrorWithDefaults instantiates a new ErrorsSavedSearchNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsSavedSearchNotFoundError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsSavedSearchNotFoundError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsSavedSearchNotFoundError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsSavedSearchNotFoundError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsSavedSearchNotFoundError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsSavedSearchNotFoundError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorsSavedSearchNotFoundError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


