# ErrorsEmployerAddressErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsEmployerAddressError**](ErrorsEmployerAddressError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsEmployerAddressErrors

`func NewErrorsEmployerAddressErrors(requestId string, errors []ErrorsEmployerAddressError, ) *ErrorsEmployerAddressErrors`

NewErrorsEmployerAddressErrors instantiates a new ErrorsEmployerAddressErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsEmployerAddressErrorsWithDefaults

`func NewErrorsEmployerAddressErrorsWithDefaults() *ErrorsEmployerAddressErrors`

NewErrorsEmployerAddressErrorsWithDefaults instantiates a new ErrorsEmployerAddressErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsEmployerAddressErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsEmployerAddressErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsEmployerAddressErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsEmployerAddressErrors) GetErrors() []ErrorsEmployerAddressError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsEmployerAddressErrors) GetErrorsOk() (*[]ErrorsEmployerAddressError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsEmployerAddressErrors) SetErrors(v []ErrorsEmployerAddressError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


