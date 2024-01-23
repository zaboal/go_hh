# ErrorsCommonBadArgumentErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsCommonBadArgumentError**](ErrorsCommonBadArgumentError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsCommonBadArgumentErrors

`func NewErrorsCommonBadArgumentErrors(requestId string, errors []ErrorsCommonBadArgumentError, ) *ErrorsCommonBadArgumentErrors`

NewErrorsCommonBadArgumentErrors instantiates a new ErrorsCommonBadArgumentErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadArgumentErrorsWithDefaults

`func NewErrorsCommonBadArgumentErrorsWithDefaults() *ErrorsCommonBadArgumentErrors`

NewErrorsCommonBadArgumentErrorsWithDefaults instantiates a new ErrorsCommonBadArgumentErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsCommonBadArgumentErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsCommonBadArgumentErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsCommonBadArgumentErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsCommonBadArgumentErrors) GetErrors() []ErrorsCommonBadArgumentError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsCommonBadArgumentErrors) GetErrorsOk() (*[]ErrorsCommonBadArgumentError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsCommonBadArgumentErrors) SetErrors(v []ErrorsCommonBadArgumentError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


