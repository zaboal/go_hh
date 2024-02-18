# NegotiationsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateUrl** | Pointer to **string** | Ссылка на сайт, переход по которой инициирует действие | [optional] 
**DisableReason** | Pointer to **string** | Пояснение, почему действие недоступно | [optional] 
**Enabled** | **bool** | Доступно ли действие | 
**Id** | **string** | Тип действия | 
**Name** | **string** | Описание действия | 

## Methods

### NewNegotiationsAction

`func NewNegotiationsAction(enabled bool, id string, name string, ) *NegotiationsAction`

NewNegotiationsAction instantiates a new NegotiationsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsActionWithDefaults

`func NewNegotiationsActionWithDefaults() *NegotiationsAction`

NewNegotiationsActionWithDefaults instantiates a new NegotiationsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *NegotiationsAction) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *NegotiationsAction) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *NegotiationsAction) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.

### HasAlternateUrl

`func (o *NegotiationsAction) HasAlternateUrl() bool`

HasAlternateUrl returns a boolean if a field has been set.

### GetDisableReason

`func (o *NegotiationsAction) GetDisableReason() string`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *NegotiationsAction) GetDisableReasonOk() (*string, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *NegotiationsAction) SetDisableReason(v string)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *NegotiationsAction) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetEnabled

`func (o *NegotiationsAction) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NegotiationsAction) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NegotiationsAction) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *NegotiationsAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsAction) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NegotiationsAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NegotiationsAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NegotiationsAction) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


