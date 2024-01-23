# SavedSearchesSavedSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | Дата и время создания | 
**Id** | **string** | Идентификатор поиска | 
**Items** | [**SavedSearchesSavedSearchItemItems**](SavedSearchesSavedSearchItemItems.md) |  | 
**Name** | **string** | Название поиска | 
**NewItems** | [**SavedSearchesSavedSearchItemNewItems**](SavedSearchesSavedSearchItemNewItems.md) |  | 
**Subscription** | **bool** | Статус подписки | 

## Methods

### NewSavedSearchesSavedSearchItem

`func NewSavedSearchesSavedSearchItem(createdAt string, id string, items SavedSearchesSavedSearchItemItems, name string, newItems SavedSearchesSavedSearchItemNewItems, subscription bool, ) *SavedSearchesSavedSearchItem`

NewSavedSearchesSavedSearchItem instantiates a new SavedSearchesSavedSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedSearchesSavedSearchItemWithDefaults

`func NewSavedSearchesSavedSearchItemWithDefaults() *SavedSearchesSavedSearchItem`

NewSavedSearchesSavedSearchItemWithDefaults instantiates a new SavedSearchesSavedSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SavedSearchesSavedSearchItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SavedSearchesSavedSearchItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SavedSearchesSavedSearchItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *SavedSearchesSavedSearchItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedSearchesSavedSearchItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedSearchesSavedSearchItem) SetId(v string)`

SetId sets Id field to given value.


### GetItems

`func (o *SavedSearchesSavedSearchItem) GetItems() SavedSearchesSavedSearchItemItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SavedSearchesSavedSearchItem) GetItemsOk() (*SavedSearchesSavedSearchItemItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SavedSearchesSavedSearchItem) SetItems(v SavedSearchesSavedSearchItemItems)`

SetItems sets Items field to given value.


### GetName

`func (o *SavedSearchesSavedSearchItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedSearchesSavedSearchItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedSearchesSavedSearchItem) SetName(v string)`

SetName sets Name field to given value.


### GetNewItems

`func (o *SavedSearchesSavedSearchItem) GetNewItems() SavedSearchesSavedSearchItemNewItems`

GetNewItems returns the NewItems field if non-nil, zero value otherwise.

### GetNewItemsOk

`func (o *SavedSearchesSavedSearchItem) GetNewItemsOk() (*SavedSearchesSavedSearchItemNewItems, bool)`

GetNewItemsOk returns a tuple with the NewItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewItems

`func (o *SavedSearchesSavedSearchItem) SetNewItems(v SavedSearchesSavedSearchItemNewItems)`

SetNewItems sets NewItems field to given value.


### GetSubscription

`func (o *SavedSearchesSavedSearchItem) GetSubscription() bool`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SavedSearchesSavedSearchItem) GetSubscriptionOk() (*bool, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SavedSearchesSavedSearchItem) SetSubscription(v bool)`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


