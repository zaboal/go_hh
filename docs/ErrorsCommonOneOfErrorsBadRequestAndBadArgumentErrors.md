# ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsCommonBadRequestError**](ErrorsCommonBadRequestError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors

`func NewErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors(requestId string, errors []ErrorsCommonBadRequestError, ) *ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors`

NewErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors instantiates a new ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrorsWithDefaults

`func NewErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrorsWithDefaults() *ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors`

NewErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrorsWithDefaults instantiates a new ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors) GetErrors() []ErrorsCommonBadRequestError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors) GetErrorsOk() (*[]ErrorsCommonBadRequestError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsCommonOneOfErrorsBadRequestAndBadArgumentErrors) SetErrors(v []ErrorsCommonBadRequestError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


