# NegotiationsMessagesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Items** | Pointer to [**[]NegotiationsMessagesGet**](NegotiationsMessagesGet.md) | Список сообщений | [optional] 

## Methods

### NewNegotiationsMessagesGetResponse

`func NewNegotiationsMessagesGetResponse(found float32, page float32, pages float32, perPage float32, ) *NegotiationsMessagesGetResponse`

NewNegotiationsMessagesGetResponse instantiates a new NegotiationsMessagesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsMessagesGetResponseWithDefaults

`func NewNegotiationsMessagesGetResponseWithDefaults() *NegotiationsMessagesGetResponse`

NewNegotiationsMessagesGetResponseWithDefaults instantiates a new NegotiationsMessagesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *NegotiationsMessagesGetResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *NegotiationsMessagesGetResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *NegotiationsMessagesGetResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *NegotiationsMessagesGetResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *NegotiationsMessagesGetResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *NegotiationsMessagesGetResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *NegotiationsMessagesGetResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *NegotiationsMessagesGetResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *NegotiationsMessagesGetResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *NegotiationsMessagesGetResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *NegotiationsMessagesGetResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *NegotiationsMessagesGetResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetItems

`func (o *NegotiationsMessagesGetResponse) GetItems() []NegotiationsMessagesGet`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NegotiationsMessagesGetResponse) GetItemsOk() (*[]NegotiationsMessagesGet, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NegotiationsMessagesGetResponse) SetItems(v []NegotiationsMessagesGet)`

SetItems sets Items field to given value.

### HasItems

`func (o *NegotiationsMessagesGetResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


