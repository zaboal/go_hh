# ErrorsNegotiationTopicsStateChangeDeniedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Ошибки при переводе откликов по вакансии:   * &#x60;application_denied&#x60; — ошибка доступа к отклику. Может возникнуть в случае перевода нескольких откликов в другой статус, если,  как минимум один из откликов принадлежит другой вакансии  | 

## Methods

### NewErrorsNegotiationTopicsStateChangeDeniedError

`func NewErrorsNegotiationTopicsStateChangeDeniedError(type_ string, value string, ) *ErrorsNegotiationTopicsStateChangeDeniedError`

NewErrorsNegotiationTopicsStateChangeDeniedError instantiates a new ErrorsNegotiationTopicsStateChangeDeniedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsNegotiationTopicsStateChangeDeniedErrorWithDefaults

`func NewErrorsNegotiationTopicsStateChangeDeniedErrorWithDefaults() *ErrorsNegotiationTopicsStateChangeDeniedError`

NewErrorsNegotiationTopicsStateChangeDeniedErrorWithDefaults instantiates a new ErrorsNegotiationTopicsStateChangeDeniedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsNegotiationTopicsStateChangeDeniedError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsNegotiationTopicsStateChangeDeniedError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsNegotiationTopicsStateChangeDeniedError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsNegotiationTopicsStateChangeDeniedError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsNegotiationTopicsStateChangeDeniedError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsNegotiationTopicsStateChangeDeniedError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


