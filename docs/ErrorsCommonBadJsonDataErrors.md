# ErrorsCommonBadJsonDataErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsCommonBadJsonDataError**](ErrorsCommonBadJsonDataError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsCommonBadJsonDataErrors

`func NewErrorsCommonBadJsonDataErrors(requestId string, errors []ErrorsCommonBadJsonDataError, ) *ErrorsCommonBadJsonDataErrors`

NewErrorsCommonBadJsonDataErrors instantiates a new ErrorsCommonBadJsonDataErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadJsonDataErrorsWithDefaults

`func NewErrorsCommonBadJsonDataErrorsWithDefaults() *ErrorsCommonBadJsonDataErrors`

NewErrorsCommonBadJsonDataErrorsWithDefaults instantiates a new ErrorsCommonBadJsonDataErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsCommonBadJsonDataErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsCommonBadJsonDataErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsCommonBadJsonDataErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsCommonBadJsonDataErrors) GetErrors() []ErrorsCommonBadJsonDataError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsCommonBadJsonDataErrors) GetErrorsOk() (*[]ErrorsCommonBadJsonDataError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsCommonBadJsonDataErrors) SetErrors(v []ErrorsCommonBadJsonDataError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


