# EmployerServicesEmployerServiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatedAt** | **string** | Время активации услуги (в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**Balance** | Pointer to [**EmployerServicesBalance**](EmployerServicesBalance.md) |  | [optional] 
**ExpiresAt** | **string** | Время истечения услуги (в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**Id** | **string** | Идентификатор услуги | 
**ServiceType** | [**EmployerServicesServiceType**](EmployerServicesServiceType.md) |  | 

## Methods

### NewEmployerServicesEmployerServiceItem

`func NewEmployerServicesEmployerServiceItem(activatedAt string, expiresAt string, id string, serviceType EmployerServicesServiceType, ) *EmployerServicesEmployerServiceItem`

NewEmployerServicesEmployerServiceItem instantiates a new EmployerServicesEmployerServiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerServicesEmployerServiceItemWithDefaults

`func NewEmployerServicesEmployerServiceItemWithDefaults() *EmployerServicesEmployerServiceItem`

NewEmployerServicesEmployerServiceItemWithDefaults instantiates a new EmployerServicesEmployerServiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatedAt

`func (o *EmployerServicesEmployerServiceItem) GetActivatedAt() string`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *EmployerServicesEmployerServiceItem) GetActivatedAtOk() (*string, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *EmployerServicesEmployerServiceItem) SetActivatedAt(v string)`

SetActivatedAt sets ActivatedAt field to given value.


### GetBalance

`func (o *EmployerServicesEmployerServiceItem) GetBalance() EmployerServicesBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *EmployerServicesEmployerServiceItem) GetBalanceOk() (*EmployerServicesBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *EmployerServicesEmployerServiceItem) SetBalance(v EmployerServicesBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *EmployerServicesEmployerServiceItem) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetExpiresAt

`func (o *EmployerServicesEmployerServiceItem) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EmployerServicesEmployerServiceItem) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EmployerServicesEmployerServiceItem) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *EmployerServicesEmployerServiceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployerServicesEmployerServiceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployerServicesEmployerServiceItem) SetId(v string)`

SetId sets Id field to given value.


### GetServiceType

`func (o *EmployerServicesEmployerServiceItem) GetServiceType() EmployerServicesServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *EmployerServicesEmployerServiceItem) GetServiceTypeOk() (*EmployerServicesServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *EmployerServicesEmployerServiceItem) SetServiceType(v EmployerServicesServiceType)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


