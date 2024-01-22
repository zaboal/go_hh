# SuggestsErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**BadArgument** | Pointer to **string** |  | [optional] 
**BadArguments** | Pointer to [**[]SuggestsErrorsAllOfBadArguments**](SuggestsErrorsAllOfBadArguments.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Errors** | [**[]SuggestsErrorsAllOfErrors**](SuggestsErrorsAllOfErrors.md) |  | 

## Methods

### NewSuggestsErrors

`func NewSuggestsErrors(requestId string, errors []SuggestsErrorsAllOfErrors, ) *SuggestsErrors`

NewSuggestsErrors instantiates a new SuggestsErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestsErrorsWithDefaults

`func NewSuggestsErrorsWithDefaults() *SuggestsErrors`

NewSuggestsErrorsWithDefaults instantiates a new SuggestsErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SuggestsErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SuggestsErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SuggestsErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetBadArgument

`func (o *SuggestsErrors) GetBadArgument() string`

GetBadArgument returns the BadArgument field if non-nil, zero value otherwise.

### GetBadArgumentOk

`func (o *SuggestsErrors) GetBadArgumentOk() (*string, bool)`

GetBadArgumentOk returns a tuple with the BadArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArgument

`func (o *SuggestsErrors) SetBadArgument(v string)`

SetBadArgument sets BadArgument field to given value.

### HasBadArgument

`func (o *SuggestsErrors) HasBadArgument() bool`

HasBadArgument returns a boolean if a field has been set.

### GetBadArguments

`func (o *SuggestsErrors) GetBadArguments() []SuggestsErrorsAllOfBadArguments`

GetBadArguments returns the BadArguments field if non-nil, zero value otherwise.

### GetBadArgumentsOk

`func (o *SuggestsErrors) GetBadArgumentsOk() (*[]SuggestsErrorsAllOfBadArguments, bool)`

GetBadArgumentsOk returns a tuple with the BadArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArguments

`func (o *SuggestsErrors) SetBadArguments(v []SuggestsErrorsAllOfBadArguments)`

SetBadArguments sets BadArguments field to given value.

### HasBadArguments

`func (o *SuggestsErrors) HasBadArguments() bool`

HasBadArguments returns a boolean if a field has been set.

### GetDescription

`func (o *SuggestsErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SuggestsErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SuggestsErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SuggestsErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *SuggestsErrors) GetErrors() []SuggestsErrorsAllOfErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SuggestsErrors) GetErrorsOk() (*[]SuggestsErrorsAllOfErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SuggestsErrors) SetErrors(v []SuggestsErrorsAllOfErrors)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


