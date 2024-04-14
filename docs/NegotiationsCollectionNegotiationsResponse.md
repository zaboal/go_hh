# NegotiationsCollectionNegotiationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderedBy** | [**IncludesIdName**](IncludesIdName.md) | Применяемый тип сортировки | 
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Items** | [**[]NegotiationsCollectionNegotiationsItemsInner**](NegotiationsCollectionNegotiationsItemsInner.md) |  | 

## Methods

### NewNegotiationsCollectionNegotiationsResponse

`func NewNegotiationsCollectionNegotiationsResponse(orderedBy IncludesIdName, found float32, page float32, pages float32, perPage float32, items []NegotiationsCollectionNegotiationsItemsInner, ) *NegotiationsCollectionNegotiationsResponse`

NewNegotiationsCollectionNegotiationsResponse instantiates a new NegotiationsCollectionNegotiationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsCollectionNegotiationsResponseWithDefaults

`func NewNegotiationsCollectionNegotiationsResponseWithDefaults() *NegotiationsCollectionNegotiationsResponse`

NewNegotiationsCollectionNegotiationsResponseWithDefaults instantiates a new NegotiationsCollectionNegotiationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderedBy

`func (o *NegotiationsCollectionNegotiationsResponse) GetOrderedBy() IncludesIdName`

GetOrderedBy returns the OrderedBy field if non-nil, zero value otherwise.

### GetOrderedByOk

`func (o *NegotiationsCollectionNegotiationsResponse) GetOrderedByOk() (*IncludesIdName, bool)`

GetOrderedByOk returns a tuple with the OrderedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedBy

`func (o *NegotiationsCollectionNegotiationsResponse) SetOrderedBy(v IncludesIdName)`

SetOrderedBy sets OrderedBy field to given value.


### GetFound

`func (o *NegotiationsCollectionNegotiationsResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *NegotiationsCollectionNegotiationsResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *NegotiationsCollectionNegotiationsResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *NegotiationsCollectionNegotiationsResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *NegotiationsCollectionNegotiationsResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *NegotiationsCollectionNegotiationsResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *NegotiationsCollectionNegotiationsResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *NegotiationsCollectionNegotiationsResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *NegotiationsCollectionNegotiationsResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *NegotiationsCollectionNegotiationsResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *NegotiationsCollectionNegotiationsResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *NegotiationsCollectionNegotiationsResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetItems

`func (o *NegotiationsCollectionNegotiationsResponse) GetItems() []NegotiationsCollectionNegotiationsItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NegotiationsCollectionNegotiationsResponse) GetItemsOk() (*[]NegotiationsCollectionNegotiationsItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NegotiationsCollectionNegotiationsResponse) SetItems(v []NegotiationsCollectionNegotiationsItemsInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


