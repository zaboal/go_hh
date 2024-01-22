# ErrorsNegotiationHideResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип ошибки | 
**Value** | **string** | Причина ошибки: * &#x60;wrong_state&#x60; — действие по отклику/приглашению в данном статусе невозможно * &#x60;no_invitation&#x60; — переписка недоступна, так как в отклике ещё не было приглашения * &#x60;disabled_by_employer&#x60; — возможность переписки по отклику отключена работодателем * &#x60;chat_is_not_ready&#x60; — чат отклика/приглашения еще не готов  | 

## Methods

### NewErrorsNegotiationHideResponseError

`func NewErrorsNegotiationHideResponseError(type_ string, value string, ) *ErrorsNegotiationHideResponseError`

NewErrorsNegotiationHideResponseError instantiates a new ErrorsNegotiationHideResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsNegotiationHideResponseErrorWithDefaults

`func NewErrorsNegotiationHideResponseErrorWithDefaults() *ErrorsNegotiationHideResponseError`

NewErrorsNegotiationHideResponseErrorWithDefaults instantiates a new ErrorsNegotiationHideResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsNegotiationHideResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsNegotiationHideResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsNegotiationHideResponseError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsNegotiationHideResponseError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsNegotiationHideResponseError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsNegotiationHideResponseError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


