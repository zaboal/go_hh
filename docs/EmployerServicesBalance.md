# EmployerServicesBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actual** | **int32** | Текущее значение баланса | 
**Initial** | **int32** | Значение баланса на момент активации услуги | 

## Methods

### NewEmployerServicesBalance

`func NewEmployerServicesBalance(actual int32, initial int32, ) *EmployerServicesBalance`

NewEmployerServicesBalance instantiates a new EmployerServicesBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerServicesBalanceWithDefaults

`func NewEmployerServicesBalanceWithDefaults() *EmployerServicesBalance`

NewEmployerServicesBalanceWithDefaults instantiates a new EmployerServicesBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActual

`func (o *EmployerServicesBalance) GetActual() int32`

GetActual returns the Actual field if non-nil, zero value otherwise.

### GetActualOk

`func (o *EmployerServicesBalance) GetActualOk() (*int32, bool)`

GetActualOk returns a tuple with the Actual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActual

`func (o *EmployerServicesBalance) SetActual(v int32)`

SetActual sets Actual field to given value.


### GetInitial

`func (o *EmployerServicesBalance) GetInitial() int32`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *EmployerServicesBalance) GetInitialOk() (*int32, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *EmployerServicesBalance) SetInitial(v int32)`

SetInitial sets Initial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


