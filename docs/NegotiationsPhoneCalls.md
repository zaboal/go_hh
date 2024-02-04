# NegotiationsPhoneCalls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]NegotiationsPhoneCallItem**](NegotiationsPhoneCallItem.md) | Список звонков | 
**PickedUpPhoneByOpponent** | **bool** | Ответил ли абонент соискателю хотя бы один раз | 

## Methods

### NewNegotiationsPhoneCalls

`func NewNegotiationsPhoneCalls(items []NegotiationsPhoneCallItem, pickedUpPhoneByOpponent bool, ) *NegotiationsPhoneCalls`

NewNegotiationsPhoneCalls instantiates a new NegotiationsPhoneCalls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsPhoneCallsWithDefaults

`func NewNegotiationsPhoneCallsWithDefaults() *NegotiationsPhoneCalls`

NewNegotiationsPhoneCallsWithDefaults instantiates a new NegotiationsPhoneCalls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *NegotiationsPhoneCalls) GetItems() []NegotiationsPhoneCallItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NegotiationsPhoneCalls) GetItemsOk() (*[]NegotiationsPhoneCallItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NegotiationsPhoneCalls) SetItems(v []NegotiationsPhoneCallItem)`

SetItems sets Items field to given value.


### GetPickedUpPhoneByOpponent

`func (o *NegotiationsPhoneCalls) GetPickedUpPhoneByOpponent() bool`

GetPickedUpPhoneByOpponent returns the PickedUpPhoneByOpponent field if non-nil, zero value otherwise.

### GetPickedUpPhoneByOpponentOk

`func (o *NegotiationsPhoneCalls) GetPickedUpPhoneByOpponentOk() (*bool, bool)`

GetPickedUpPhoneByOpponentOk returns a tuple with the PickedUpPhoneByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickedUpPhoneByOpponent

`func (o *NegotiationsPhoneCalls) SetPickedUpPhoneByOpponent(v bool)`

SetPickedUpPhoneByOpponent sets PickedUpPhoneByOpponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


