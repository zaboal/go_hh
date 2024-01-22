# ErrorsNegotiationHideResponseErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsNegotiationHideResponseError**](ErrorsNegotiationHideResponseError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsNegotiationHideResponseErrors

`func NewErrorsNegotiationHideResponseErrors(requestId string, errors []ErrorsNegotiationHideResponseError, ) *ErrorsNegotiationHideResponseErrors`

NewErrorsNegotiationHideResponseErrors instantiates a new ErrorsNegotiationHideResponseErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsNegotiationHideResponseErrorsWithDefaults

`func NewErrorsNegotiationHideResponseErrorsWithDefaults() *ErrorsNegotiationHideResponseErrors`

NewErrorsNegotiationHideResponseErrorsWithDefaults instantiates a new ErrorsNegotiationHideResponseErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsNegotiationHideResponseErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsNegotiationHideResponseErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsNegotiationHideResponseErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsNegotiationHideResponseErrors) GetErrors() []ErrorsNegotiationHideResponseError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsNegotiationHideResponseErrors) GetErrorsOk() (*[]ErrorsNegotiationHideResponseError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsNegotiationHideResponseErrors) SetErrors(v []ErrorsNegotiationHideResponseError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


