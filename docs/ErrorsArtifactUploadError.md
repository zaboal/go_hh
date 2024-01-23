# ErrorsArtifactUploadError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Название поля с ошибкой. Возможные значения: * &#x60;file&#x60; — не указан файл, либо указано несколько. * &#x60;type&#x60; — некорректное значение параметра &#x60;type&#x60;. * &#x60;description&#x60; — слишком длинное описание. * &#x60;limit_exceeded&#x60; — превышено количество артефактов. * &#x60;unknown_format&#x60; — неизвестный формат файла  | 

## Methods

### NewErrorsArtifactUploadError

`func NewErrorsArtifactUploadError(type_ string, value string, ) *ErrorsArtifactUploadError`

NewErrorsArtifactUploadError instantiates a new ErrorsArtifactUploadError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsArtifactUploadErrorWithDefaults

`func NewErrorsArtifactUploadErrorWithDefaults() *ErrorsArtifactUploadError`

NewErrorsArtifactUploadErrorWithDefaults instantiates a new ErrorsArtifactUploadError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsArtifactUploadError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsArtifactUploadError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsArtifactUploadError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsArtifactUploadError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsArtifactUploadError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsArtifactUploadError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


