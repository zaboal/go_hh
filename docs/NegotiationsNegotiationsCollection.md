# NegotiationsNegotiationsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Описание коллекции | 
**Id** | **string** | Идентификатор коллекции | 
**Name** | **string** | Название коллекции | 
**Url** | **string** | URL, [GET-запрос на который](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-negotiations) возвращает список откликов/приглашений коллекции  | 
**Counters** | Pointer to [**NegotiationsObjectsEmployerCounters**](NegotiationsObjectsEmployerCounters.md) |  | [optional] 
**OrderTypes** | [**[]NegotiationsNegotiationsOrderTypes**](NegotiationsNegotiationsOrderTypes.md) |  | 

## Methods

### NewNegotiationsNegotiationsCollection

`func NewNegotiationsNegotiationsCollection(description string, id string, name string, url string, orderTypes []NegotiationsNegotiationsOrderTypes, ) *NegotiationsNegotiationsCollection`

NewNegotiationsNegotiationsCollection instantiates a new NegotiationsNegotiationsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsNegotiationsCollectionWithDefaults

`func NewNegotiationsNegotiationsCollectionWithDefaults() *NegotiationsNegotiationsCollection`

NewNegotiationsNegotiationsCollectionWithDefaults instantiates a new NegotiationsNegotiationsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NegotiationsNegotiationsCollection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NegotiationsNegotiationsCollection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NegotiationsNegotiationsCollection) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *NegotiationsNegotiationsCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsNegotiationsCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsNegotiationsCollection) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NegotiationsNegotiationsCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NegotiationsNegotiationsCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NegotiationsNegotiationsCollection) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *NegotiationsNegotiationsCollection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NegotiationsNegotiationsCollection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NegotiationsNegotiationsCollection) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCounters

`func (o *NegotiationsNegotiationsCollection) GetCounters() NegotiationsObjectsEmployerCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *NegotiationsNegotiationsCollection) GetCountersOk() (*NegotiationsObjectsEmployerCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *NegotiationsNegotiationsCollection) SetCounters(v NegotiationsObjectsEmployerCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *NegotiationsNegotiationsCollection) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetOrderTypes

`func (o *NegotiationsNegotiationsCollection) GetOrderTypes() []NegotiationsNegotiationsOrderTypes`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *NegotiationsNegotiationsCollection) GetOrderTypesOk() (*[]NegotiationsNegotiationsOrderTypes, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *NegotiationsNegotiationsCollection) SetOrderTypes(v []NegotiationsNegotiationsOrderTypes)`

SetOrderTypes sets OrderTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


